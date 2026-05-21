package kinds

import (
	"github.com/grafana/grafana/apps/alerting/admin/kinds/v0alpha1"
)

// AlertingStatus is a synthetic k8s kind exposing the per-namespace observation
// state across the admin app's status-carrying kinds. The resource is read-only
// and not backed by apistore — a custom rest.Storage in
// pkg/app/status_storage.go synthesizes it on Get and List by composing the
// status payloads of the other kinds in this group.
//
// Spec is empty (observation-only kind). The status field carries the
// area-grouped observation payload, mirroring AlertingConfig.spec structure.
alertingStatusKind: {
	kind:       "AlertingStatus"
	pluralName: "AlertingStatuses"
}

alertingStatusv0alpha1: alertingStatusKind & {
	schema: {
		spec:   v0alpha1.AlertingStatusSpec
		status: v0alpha1.AlertingStatusStatus
	}
}
