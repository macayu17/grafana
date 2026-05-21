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
				configv0alpha1,
				externalAlertmanagerSyncv0alpha1,
				summaryv0alpha1,
			]
		}
	}
	roles: {}
}

// Config kind: per-org alerting admin config (singleton). Inlined here rather
// than in a separate file because a separate `Config.cue` collides with the
// SDK config selector file `config.cue` on case-insensitive filesystems.
configKind: {
	kind:       "Config"
	pluralName: "Configs"
}

configv0alpha1: configKind & {
	schema: {
		spec: v0alpha1.ConfigSpec
	}
}
