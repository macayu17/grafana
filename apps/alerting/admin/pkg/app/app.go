package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grafana/grafana-app-sdk/app"
	"github.com/grafana/grafana-app-sdk/resource"
	"github.com/grafana/grafana-app-sdk/simple"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/grafana/grafana/apps/alerting/admin/pkg/apis/alertingadmin/v0alpha1"
	"github.com/grafana/grafana/apps/alerting/admin/pkg/app/config"
)

// statusSingletonName is the fixed object name for singleton-per-org status
// resources in this app. Each kind has at most one instance per namespace.
const statusSingletonName = "default"

func New(cfg app.Config) (app.App, error) {
	runtimeConfig, _ := cfg.SpecificConfig.(config.RuntimeConfig)

	simpleConfig := simple.AppConfig{
		Name:       "alerting.admin",
		KubeConfig: cfg.KubeConfig,
		VersionedCustomRoutes: map[string]simple.AppVersionRouteHandlers{
			"v0alpha1": {
				{
					Namespaced: true,
					Path:       "/summary",
					Method:     "GET",
				}: newAggregateSummaryHandler(runtimeConfig.ClientGenerator),
			},
		},
		ManagedKinds: []simple.AppManagedKind{
			{Kind: v0alpha1.ConfigKind()},
			{Kind: v0alpha1.ExternalAlertmanagerSyncKind()},
		},
	}

	a, err := simple.NewApp(simpleConfig)
	if err != nil {
		return nil, err
	}

	if err := a.ValidateManifest(cfg.ManifestData); err != nil {
		return nil, err
	}

	return a, nil
}

// newAggregateSummaryHandler returns a handler for the /summary route that
// fans out across status kinds in this group and returns each kind's
// singleton (when present) as a composite grouped by area of concern,
// mirroring the structure of Config.spec. Sub-keys absent from the response
// mean "no singleton observed for this concern yet" — clients should treat
// absence as "not yet observed" rather than as an empty status.
//
// To add a new status kind: build its client here, GET its singleton, and
// place the resulting .Status payload at the response path that matches
// where the kind's settings live on Config.spec
// (e.g. alertmanager → externalSync → ExternalAlertmanagerSync.status).
func newAggregateSummaryHandler(clientGenerator resource.ClientGenerator) simple.AppCustomRouteHandler {
	return func(ctx context.Context, writer app.CustomRouteResponseWriter, request *app.CustomRouteRequest) error {
		alertmanager := map[string]interface{}{}

		if clientGenerator != nil {
			eams, err := v0alpha1.NewExternalAlertmanagerSyncClientFromGenerator(clientGenerator)
			if err == nil {
				obj, getErr := eams.Get(ctx, resource.Identifier{Namespace: request.ResourceIdentifier.Namespace, Name: statusSingletonName})
				if getErr == nil {
					alertmanager["externalSync"] = obj.Status
				} else if !k8serrors.IsNotFound(getErr) {
					return getErr
				}
			}
		}

		body := map[string]interface{}{}
		if len(alertmanager) > 0 {
			body["alertmanager"] = alertmanager
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		return json.NewEncoder(writer).Encode(map[string]interface{}{"body": body})
	}
}
