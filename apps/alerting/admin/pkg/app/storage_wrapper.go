package app

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	genericserver "k8s.io/apiserver/pkg/server"

	appsdkapiserver "github.com/grafana/grafana-app-sdk/k8s/apiserver"
)

// customStorageWrapper sits between the app-sdk's GenericAPIServer and
// the AppInstaller's InstallAPIs call, swapping the auto-generated
// rest.Storage for selected GroupVersionResources before the apiserver
// wires them up. Mirrors apps/plugins/pkg/app/storage_wrapper.go — we
// only need it for the synthetic Summary kind here.
type customStorageWrapper struct {
	wrapped appsdkapiserver.GenericAPIServer
	replace map[schema.GroupVersionResource]rest.Storage
}

var _ appsdkapiserver.GenericAPIServer = (*customStorageWrapper)(nil)

// NewCustomStorageWrapper wraps a GenericAPIServer with a per-GVR storage
// override map. The wrapper is consumed by AppInstaller.InstallAPIs in
// pkg/registry/apps/alerting/admin/register.go.
func NewCustomStorageWrapper(wrapped appsdkapiserver.GenericAPIServer, replace map[schema.GroupVersionResource]rest.Storage) appsdkapiserver.GenericAPIServer {
	return &customStorageWrapper{wrapped: wrapped, replace: replace}
}

func (c *customStorageWrapper) InstallAPIGroup(apiGroupInfo *genericserver.APIGroupInfo) error {
	if apiGroupInfo == nil || apiGroupInfo.VersionedResourcesStorageMap == nil {
		return fmt.Errorf("apiGroupInfo cannot be nil")
	}
	for gvr, storage := range c.replace {
		if _, ok := apiGroupInfo.VersionedResourcesStorageMap[gvr.Version]; !ok {
			apiGroupInfo.VersionedResourcesStorageMap[gvr.Version] = map[string]rest.Storage{}
		}
		apiGroupInfo.VersionedResourcesStorageMap[gvr.Version][gvr.Resource] = storage
	}
	return c.wrapped.InstallAPIGroup(apiGroupInfo)
}

// RegisteredWebServices is required by the GenericAPIServer interface but
// unused on this path — apiserver routes are wired via InstallAPIGroup.
func (c *customStorageWrapper) RegisteredWebServices() []*restful.WebService {
	return []*restful.WebService{}
}
