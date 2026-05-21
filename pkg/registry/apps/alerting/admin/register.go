package admin

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	restclient "k8s.io/client-go/rest"

	"github.com/grafana/grafana-app-sdk/app"
	appsdkapiserver "github.com/grafana/grafana-app-sdk/k8s/apiserver"
	"github.com/grafana/grafana-app-sdk/resource"
	"github.com/grafana/grafana-app-sdk/simple"

	alertingadminv0alpha1 "github.com/grafana/grafana/apps/alerting/admin/pkg/apis/alertingadmin/v0alpha1"
	"github.com/grafana/grafana/apps/alerting/admin/pkg/apis/manifestdata"
	adminApp "github.com/grafana/grafana/apps/alerting/admin/pkg/app"
	adminAppConfig "github.com/grafana/grafana/apps/alerting/admin/pkg/app/config"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/ngalert"
	"github.com/grafana/grafana/pkg/setting"
)

var (
	_ appsdkapiserver.AppInstaller = (*AppInstaller)(nil)
)

type AppInstaller struct {
	appsdkapiserver.AppInstaller
	clientGenerator resource.ClientGenerator
}

// GetAuthorizer permits all requests for now. An admin-only RBAC policy
// belongs here before the API graduates.
func (a *AppInstaller) GetAuthorizer() authorizer.Authorizer {
	return authorizer.AuthorizerFunc(
		func(ctx context.Context, a authorizer.Attributes) (authorizer.Decision, string, error) {
			return authorizer.DecisionAllow, "", nil
		},
	)
}

// InstallAPIs swaps the auto-generated storage for the synthetic
// AlertingStatus kind with our composing statusStorage. The AlertingStatus
// kind is read-only; it has no apistore backing — Get/List call into the
// other status kinds' singletons and assemble the response on the fly. See
// apps/alerting/admin/pkg/app/status_storage.go for the implementation,
// and apps/plugins/pkg/app/storage_wrapper.go for the prior-art pattern
// this mirrors.
func (a *AppInstaller) InstallAPIs(server appsdkapiserver.GenericAPIServer, restOptsGetter generic.RESTOptionsGetter) error {
	statusGVR := alertingadminv0alpha1.AlertingStatusKind().GroupVersionResource()
	wrapped := adminApp.NewCustomStorageWrapper(server, map[schema.GroupVersionResource]rest.Storage{
		statusGVR: adminApp.NewStatusStorage(a.clientGenerator),
	})
	return a.AppInstaller.InstallAPIs(wrapped, restOptsGetter)
}

func RegisterAppInstaller(
	cfg *setting.Cfg,
	ng *ngalert.AlertNG,
	clientGenerator resource.ClientGenerator,
) (*AppInstaller, error) {
	if ng != nil && ng.IsDisabled() {
		log.New("app-registry").Info("Skipping Kubernetes Alerting Admin apiserver (admin.alerting.grafana.app): Unified Alerting is disabled")
		return nil, nil
	}

	return NewAppInstaller(clientGenerator)
}

func NewAppInstaller(clientGenerator resource.ClientGenerator) (*AppInstaller, error) {
	installer := &AppInstaller{clientGenerator: clientGenerator}

	localManifest := manifestdata.LocalManifest()
	runtimeConfig := adminAppConfig.RuntimeConfig{ClientGenerator: clientGenerator}

	provider := simple.NewAppProvider(localManifest, runtimeConfig, adminApp.New)

	appConfig := app.Config{
		KubeConfig:     restclient.Config{},
		ManifestData:   *localManifest.ManifestData,
		SpecificConfig: runtimeConfig,
	}

	i, err := appsdkapiserver.NewDefaultAppInstaller(provider, appConfig, &manifestdata.GoTypeAssociator{})
	if err != nil {
		return nil, err
	}
	installer.AppInstaller = i
	return installer, nil
}
