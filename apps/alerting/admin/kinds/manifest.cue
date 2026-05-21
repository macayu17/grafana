package kinds

import (
	"github.com/grafana/grafana/apps/alerting/admin/kinds/v0alpha1"
)

manifest: {
	appName:       "alerting-admin"
	groupOverride: "admin.alerting.grafana.app"
	versions: {
		"v0alpha1": {
			codegen: {
				ts: {enabled: false}
				go: {enabled: true}
			}
			kinds: [
				alertingConfigv0alpha1,
				externalAlertmanagerSyncv0alpha1,
				alertingStatusv0alpha1,
			]
		}
	}
	roles: {}
}

// AlertingConfig kind: per-org alerting admin config (singleton). Inlined
// here rather than in a separate file because a separate `AlertingConfig.cue`
// collides with the SDK config selector file `config.cue` on case-insensitive
// filesystems.
alertingConfigKind: {
	kind:       "AlertingConfig"
	pluralName: "AlertingConfigs"
}

alertingConfigv0alpha1: alertingConfigKind & {
	schema: {
		spec: v0alpha1.AlertingConfigSpec
	}
}
