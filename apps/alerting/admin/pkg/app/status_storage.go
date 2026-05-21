package app

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/grafana/grafana-app-sdk/resource"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/grafana/grafana/apps/alerting/admin/pkg/apis/alertingadmin/v0alpha1"
)

// statusSingletonName is the only object name the synthetic AlertingStatus
// kind exposes per namespace. Mirrors the convention used for the other
// singleton-per-org kinds in this group.
const statusSingletonName = "default"

// statusStorage implements a synthetic rest.Storage for the AlertingStatus
// kind. Reads compose .status from the other status kinds in this group;
// writes are rejected. The kind is not backed by apistore.
//
// The storage is registered via customStorageWrapper, which replaces the
// auto-generated app-sdk storage for the AlertingStatus GroupVersionResource.
type statusStorage struct {
	clientGenerator resource.ClientGenerator

	gr             schema.GroupResource
	tableConverter rest.TableConvertor

	clientOnce sync.Once
	client     *v0alpha1.ExternalAlertmanagerSyncClient
	clientErr  error
}

var (
	_ rest.Scoper               = (*statusStorage)(nil)
	_ rest.SingularNameProvider = (*statusStorage)(nil)
	_ rest.Getter               = (*statusStorage)(nil)
	_ rest.Lister               = (*statusStorage)(nil)
	_ rest.Storage              = (*statusStorage)(nil)
	_ rest.TableConvertor       = (*statusStorage)(nil)
)

func NewStatusStorage(clientGenerator resource.ClientGenerator) *statusStorage {
	gr := schema.GroupResource{
		Group:    v0alpha1.APIGroup,
		Resource: strings.ToLower(v0alpha1.AlertingStatusKind().Plural()),
	}
	return &statusStorage{
		clientGenerator: clientGenerator,
		gr:              gr,
		tableConverter:  rest.NewDefaultTableConvertor(gr),
	}
}

// resolveClient lazily constructs the ExternalAlertmanagerSyncClient on
// first request. Eager construction at InstallAPIs time would race the
// apiserver bring-up (same reason the ngalert syncer constructs lazily;
// see external_am_syncer.go).
func (s *statusStorage) resolveClient() (*v0alpha1.ExternalAlertmanagerSyncClient, error) {
	s.clientOnce.Do(func() {
		if s.clientGenerator == nil {
			s.clientErr = fmt.Errorf("status storage missing ClientGenerator")
			return
		}
		c, err := v0alpha1.NewExternalAlertmanagerSyncClientFromGenerator(s.clientGenerator)
		if err != nil {
			s.clientErr = fmt.Errorf("build ExternalAlertmanagerSync client: %w", err)
			return
		}
		s.client = c
	})
	return s.client, s.clientErr
}

func (s *statusStorage) New() runtime.Object {
	return v0alpha1.AlertingStatusKind().ZeroValue()
}

func (s *statusStorage) NewList() runtime.Object {
	return v0alpha1.AlertingStatusKind().ZeroListValue()
}

func (s *statusStorage) Destroy() {}

func (s *statusStorage) NamespaceScoped() bool {
	return true
}

func (s *statusStorage) GetSingularName() string {
	return strings.ToLower(v0alpha1.AlertingStatusKind().Kind())
}

func (s *statusStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return s.tableConverter.ConvertToTable(ctx, object, tableOptions)
}

// Get returns the synthesized AlertingStatus singleton for the request
// namespace. Any name other than the singleton returns NotFound — there is
// only one AlertingStatus per namespace by construction.
func (s *statusStorage) Get(ctx context.Context, name string, _ *metav1.GetOptions) (runtime.Object, error) {
	ns := apirequest.NamespaceValue(ctx)
	if ns == "" {
		return nil, apierrors.NewBadRequest("missing namespace in request")
	}
	if name != statusSingletonName {
		return nil, apierrors.NewNotFound(s.gr, name)
	}
	return s.synthesize(ctx, ns)
}

// List returns an AlertingStatusList containing the single synthesized
// AlertingStatus for the request namespace.
func (s *statusStorage) List(ctx context.Context, _ *internalversion.ListOptions) (runtime.Object, error) {
	ns := apirequest.NamespaceValue(ctx)
	if ns == "" {
		return nil, apierrors.NewBadRequest("missing namespace in request")
	}
	item, err := s.synthesize(ctx, ns)
	if err != nil {
		return nil, err
	}
	list := v0alpha1.AlertingStatusKind().ZeroListValue().(*v0alpha1.AlertingStatusList)
	list.TypeMeta = metav1.TypeMeta{
		APIVersion: v0alpha1.APIGroup + "/" + v0alpha1.APIVersion,
		Kind:       v0alpha1.AlertingStatusKind().Kind() + "List",
	}
	list.Items = []v0alpha1.AlertingStatus{*item}
	return list, nil
}

// synthesize composes an AlertingStatus for the given namespace by fetching
// the other status kinds' singletons and mapping their .status payloads
// onto AlertingStatus.status under the area sub-tree defined in
// alertingStatus_status.cue.
//
// Missing observation kinds (NotFound) leave the corresponding area sub-key
// absent rather than emitting a zero-valued object; clients should treat
// absence as not yet observed.
func (s *statusStorage) synthesize(ctx context.Context, namespace string) (*v0alpha1.AlertingStatus, error) {
	out := &v0alpha1.AlertingStatus{
		ObjectMeta: metav1.ObjectMeta{
			Name:      statusSingletonName,
			Namespace: namespace,
		},
	}
	out.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   v0alpha1.APIGroup,
		Version: v0alpha1.APIVersion,
		Kind:    v0alpha1.AlertingStatusKind().Kind(),
	})

	client, err := s.resolveClient()
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}

	eams, err := client.Get(ctx, resource.Identifier{Namespace: namespace, Name: statusSingletonName})
	switch {
	case apierrors.IsNotFound(err):
		// no observation for this concern yet — leave alertmanager absent
	case err != nil:
		return nil, err
	default:
		out.Status.Alertmanager = &v0alpha1.AlertingStatusAlertingStatusAlertmanager{
			ExternalSync: projectExternalSync(&eams.Status),
		}
	}

	return out, nil
}

// projectExternalSync maps an ExternalAlertmanagerSync .status onto the
// AlertingStatus kind's equivalent sub-tree. The shapes mirror each other
// but they're separate generated types because each kind owns its own gen
// surface.
func projectExternalSync(src *v0alpha1.ExternalAlertmanagerSyncStatus) *v0alpha1.AlertingStatusExternalAlertmanagerSyncStatus {
	if src == nil {
		return nil
	}
	out := &v0alpha1.AlertingStatusExternalAlertmanagerSyncStatus{
		ObservedGeneration: src.ObservedGeneration,
		DatasourceUid:      src.DatasourceUid,
		LastSuccessAt:      src.LastSuccessAt,
	}
	if src.Origin != nil {
		o := v0alpha1.AlertingStatusExternalAlertmanagerSyncStatusOrigin(*src.Origin)
		out.Origin = &o
	}
	if len(src.Conditions) > 0 {
		out.Conditions = make([]v0alpha1.AlertingStatusCondition, len(src.Conditions))
		for i, c := range src.Conditions {
			out.Conditions[i] = v0alpha1.AlertingStatusCondition{
				Type:               c.Type,
				Status:             v0alpha1.AlertingStatusConditionStatus(c.Status),
				LastTransitionTime: c.LastTransitionTime,
				Reason:             c.Reason,
				Message:            c.Message,
				ObservedGeneration: c.ObservedGeneration,
			}
		}
	}
	return out
}
