package app

import (
	"github.com/grafana/grafana-app-sdk/app"
	"github.com/grafana/grafana-app-sdk/simple"

	"github.com/grafana/grafana/apps/alerting/admin/pkg/apis/alertingadmin/v0alpha1"
	"github.com/grafana/grafana/apps/alerting/admin/pkg/app/config"
)

func New(cfg app.Config) (app.App, error) {
	_, _ = cfg.SpecificConfig.(config.RuntimeConfig)

	simpleConfig := simple.AppConfig{
		Name:       "alerting.admin",
		KubeConfig: cfg.KubeConfig,
		ManagedKinds: []simple.AppManagedKind{
			{Kind: v0alpha1.ConfigKind()},
			{Kind: v0alpha1.ExternalAlertmanagerSyncKind()},
			{Kind: v0alpha1.SummaryKind()},
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
