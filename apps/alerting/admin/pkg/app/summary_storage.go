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

// summarySingletonName is the only object name the synthetic Summary kind
// exposes per namespace. Mirrors the convention used for the other
// singleton-per-org kinds in this group.
const summarySingletonName = "default"

// summaryStorage implements a synthetic rest.Storage for the Summary kind.
// Reads compose .status from the other status kinds in this group; writes
// are rejected. The kind is not backed by apistore.
//
// The storage is registered via customStorageWrapper, which replaces the
// auto-generated app-sdk storage for the Summary GroupVersionResource.
type summaryStorage struct {
	clientGenerator resource.ClientGenerator

	gr             schema.GroupResource
	tableConverter rest.TableConvertor

	clientOnce sync.Once
	client     *v0alpha1.ExternalAlertmanagerSyncClient
	clientErr  error
}

var (
	_ rest.Scoper               = (*summaryStorage)(nil)
	_ rest.SingularNameProvider = (*summaryStorage)(nil)
	_ rest.Getter               = (*summaryStorage)(nil)
	_ rest.Lister               = (*summaryStorage)(nil)
	_ rest.Storage              = (*summaryStorage)(nil)
	_ rest.TableConvertor       = (*summaryStorage)(nil)
)

func NewSummaryStorage(clientGenerator resource.ClientGenerator) *summaryStorage {
	gr := schema.GroupResource{
		Group:    v0alpha1.APIGroup,
		Resource: strings.ToLower(v0alpha1.SummaryKind().Plural()),
	}
	return &summaryStorage{
		clientGenerator: clientGenerator,
		gr:              gr,
		tableConverter:  rest.NewDefaultTableConvertor(gr),
	}
}

// resolveClient lazily constructs the ExternalAlertmanagerSyncClient on
// first request. Eager construction at InstallAPIs time would race the
// apiserver bring-up (same reason the ngalert syncer constructs lazily;
// see external_am_syncer.go).
func (s *summaryStorage) resolveClient() (*v0alpha1.ExternalAlertmanagerSyncClient, error) {
	s.clientOnce.Do(func() {
		if s.clientGenerator == nil {
			s.clientErr = fmt.Errorf("summary storage missing ClientGenerator")
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

func (s *summaryStorage) New() runtime.Object {
	return v0alpha1.SummaryKind().ZeroValue()
}

func (s *summaryStorage) NewList() runtime.Object {
	return v0alpha1.SummaryKind().ZeroListValue()
}

func (s *summaryStorage) Destroy() {}

func (s *summaryStorage) NamespaceScoped() bool {
	return true
}

func (s *summaryStorage) GetSingularName() string {
	return strings.ToLower(v0alpha1.SummaryKind().Kind())
}

func (s *summaryStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return s.tableConverter.ConvertToTable(ctx, object, tableOptions)
}

// Get returns the synthesized Summary singleton for the request namespace.
// Any name other than the singleton returns NotFound — there is only one
// Summary per namespace by construction.
func (s *summaryStorage) Get(ctx context.Context, name string, _ *metav1.GetOptions) (runtime.Object, error) {
	ns := apirequest.NamespaceValue(ctx)
	if ns == "" {
		return nil, apierrors.NewBadRequest("missing namespace in request")
	}
	if name != summarySingletonName {
		return nil, apierrors.NewNotFound(s.gr, name)
	}
	return s.synthesize(ctx, ns)
}

// List returns a SummaryList containing the single synthesized Summary
// for the request namespace.
func (s *summaryStorage) List(ctx context.Context, _ *internalversion.ListOptions) (runtime.Object, error) {
	ns := apirequest.NamespaceValue(ctx)
	if ns == "" {
		return nil, apierrors.NewBadRequest("missing namespace in request")
	}
	item, err := s.synthesize(ctx, ns)
	if err != nil {
		return nil, err
	}
	list := v0alpha1.SummaryKind().ZeroListValue().(*v0alpha1.SummaryList)
	list.TypeMeta = metav1.TypeMeta{
		APIVersion: v0alpha1.APIGroup + "/" + v0alpha1.APIVersion,
		Kind:       v0alpha1.SummaryKind().Kind() + "List",
	}
	list.Items = []v0alpha1.Summary{*item}
	return list, nil
}

// synthesize composes a Summary for the given namespace by fetching the
// other status kinds' singletons and mapping their .status payloads onto
// Summary.status under the area sub-tree defined in summary_status.cue.
//
// Missing observation kinds (NotFound) leave the corresponding area
// sub-key absent rather than emitting a zero-valued object; clients should
// treat absence as not yet observed.
func (s *summaryStorage) synthesize(ctx context.Context, namespace string) (*v0alpha1.Summary, error) {
	out := &v0alpha1.Summary{
		ObjectMeta: metav1.ObjectMeta{
			Name:      summarySingletonName,
			Namespace: namespace,
		},
	}
	out.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   v0alpha1.APIGroup,
		Version: v0alpha1.APIVersion,
		Kind:    v0alpha1.SummaryKind().Kind(),
	})

	client, err := s.resolveClient()
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}

	eams, err := client.Get(ctx, resource.Identifier{Namespace: namespace, Name: summarySingletonName})
	switch {
	case apierrors.IsNotFound(err):
		// no observation for this concern yet — leave alertmanager absent
	case err != nil:
		return nil, err
	default:
		out.Status.Alertmanager = &v0alpha1.SummarySummaryAlertmanager{
			ExternalSync: projectExternalSync(&eams.Status),
		}
	}

	return out, nil
}

// projectExternalSync maps an ExternalAlertmanagerSync .status onto the
// Summary kind's equivalent sub-tree. The shapes mirror each other but
// they're separate generated types because each kind owns its own gen
// surface.
func projectExternalSync(src *v0alpha1.ExternalAlertmanagerSyncStatus) *v0alpha1.SummaryExternalAlertmanagerSyncStatus {
	if src == nil {
		return nil
	}
	out := &v0alpha1.SummaryExternalAlertmanagerSyncStatus{
		ObservedGeneration: src.ObservedGeneration,
		DatasourceUid:      src.DatasourceUid,
		LastSuccessAt:      src.LastSuccessAt,
	}
	if src.Origin != nil {
		o := v0alpha1.SummaryExternalAlertmanagerSyncStatusOrigin(*src.Origin)
		out.Origin = &o
	}
	if len(src.Conditions) > 0 {
		out.Conditions = make([]v0alpha1.SummaryCondition, len(src.Conditions))
		for i, c := range src.Conditions {
			out.Conditions[i] = v0alpha1.SummaryCondition{
				Type:               c.Type,
				Status:             v0alpha1.SummaryConditionStatus(c.Status),
				LastTransitionTime: c.LastTransitionTime,
				Reason:             c.Reason,
				Message:            c.Message,
				ObservedGeneration: c.ObservedGeneration,
			}
		}
	}
	return out
}
