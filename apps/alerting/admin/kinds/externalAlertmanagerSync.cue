package kinds

import (
	"github.com/grafana/grafana/apps/alerting/admin/kinds/v0alpha1"
)

externalAlertmanagerSyncKind: {
	kind:       "ExternalAlertmanagerSync"
	pluralName: "ExternalAlertmanagerSyncs"
}

externalAlertmanagerSyncv0alpha1: externalAlertmanagerSyncKind & {
	schema: {
		spec:   v0alpha1.ExternalAlertmanagerSyncSpec
		status: v0alpha1.ExternalAlertmanagerSyncStatus
	}
}
