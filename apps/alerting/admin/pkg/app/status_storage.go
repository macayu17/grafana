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

// conditionTypeExternalAlertmanagerSynced mirrors the condition Type
// written by the sync worker onto AlertingConfig.status.conditions[]. The
// synthetic AlertingStatus projection routes this condition into
// status.externalAlertmanagerSync.conditions so each feature sub-tree
// carries the relevant condition alongside its auxiliary fields. Future
// features add their own condition types here.
const conditionTypeExternalAlertmanagerSynced = "ExternalAlertmanagerSynced"

// statusStorage implements a synthetic rest.Storage for the AlertingStatus
// kind. Reads compose .status by fetching the source-of-truth kinds in this
// group (currently AlertingConfig — the sync worker writes its observation
// onto AlertingConfig.status) and projecting them into the area-grouped
// AlertingStatus shape. Writes are rejected; the kind is not backed by
// apistore.
//
// The storage is registered via customStorageWrapper, which replaces the
// auto-generated app-sdk storage for the AlertingStatus GroupVersionResource.
type statusStorage struct {
	clientGenerator resource.ClientGenerator

	gr             schema.GroupResource
	tableConverter rest.TableConvertor

	clientOnce sync.Once
	client     *v0alpha1.AlertingConfigClient
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

// resolveClient lazily constructs the AlertingConfigClient on first
// request. Eager construction at InstallAPIs time would race the apiserver
// bring-up (same reason the ngalert syncer constructs lazily; see
// external_am_syncer.go).
func (s *statusStorage) resolveClient() (*v0alpha1.AlertingConfigClient, error) {
	s.clientOnce.Do(func() {
		if s.clientGenerator == nil {
			s.clientErr = fmt.Errorf("status storage missing ClientGenerator")
			return
		}
		c, err := v0alpha1.NewAlertingConfigClientFromGenerator(s.clientGenerator)
		if err != nil {
			s.clientErr = fmt.Errorf("build AlertingConfig client: %w", err)
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

// synthesize composes an AlertingStatus for the given namespace by reading
// AlertingConfig.status and projecting its area sub-trees plus relevant
// conditions into the AlertingStatus shape defined in
// alertingStatus_status.cue.
//
// When no AlertingConfig exists yet (NotFound) the alertmanager area is
// omitted rather than emitted as an empty object; clients should treat
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

	cfg, err := client.Get(ctx, resource.Identifier{Namespace: namespace, Name: statusSingletonName})
	switch {
	case apierrors.IsNotFound(err):
		// no config (and therefore no observation) for this org yet —
		// leave feature sub-trees absent
	case err != nil:
		return nil, err
	default:
		if es := projectExternalSync(&cfg.Status); es != nil {
			out.Status.ExternalAlertmanagerSync = es
		}
	}

	return out, nil
}

// projectExternalSync builds the externalAlertmanagerSync sub-tree on
// AlertingStatus by copying the matching auxiliary fields from
// AlertingConfig.status.externalAlertmanagerSync and filtering
// AlertingConfig.status.conditions[] for the
// ExternalAlertmanagerSynced condition. Returns nil when there is nothing
// to report (no aux fields and no matching condition).
func projectExternalSync(src *v0alpha1.AlertingConfigStatus) *v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSync {
	var out *v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSync
	ensure := func() *v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSync {
		if out == nil {
			out = &v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSync{}
		}
		return out
	}

	if src.ExternalAlertmanagerSync != nil {
		es := src.ExternalAlertmanagerSync
		if es.DatasourceUid != nil {
			ensure().DatasourceUid = es.DatasourceUid
		}
		if es.Origin != nil {
			o := v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin(*es.Origin)
			ensure().Origin = &o
		}
	}

	for _, c := range src.Conditions {
		if c.Type != conditionTypeExternalAlertmanagerSynced {
			continue
		}
		ensure().Conditions = append(out.Conditions, v0alpha1.AlertingStatusCondition{
			Type:               c.Type,
			Status:             v0alpha1.AlertingStatusConditionStatus(c.Status),
			LastTransitionTime: c.LastTransitionTime,
			Reason:             c.Reason,
			Message:            c.Message,
			ObservedGeneration: c.ObservedGeneration,
		})
	}

	return out
}
