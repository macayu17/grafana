package kinds

import (
	"github.com/grafana/grafana/apps/alerting/admin/kinds/v0alpha1"
)

// AlertingSummary is a synthetic k8s kind exposing a per-namespace aggregate
// of observation data across the admin app's status kinds. The resource is
// read-only and not backed by storage — a custom rest.Storage in
// pkg/app/summary_storage.go synthesizes it on Get and List by composing
// the status payloads of the other kinds in this group.
//
// Spec is empty (observation-only kind). Status carries the area-grouped
// observation payload, mirroring AlertingConfig.spec structure.
alertingSummaryKind: {
	kind:       "AlertingSummary"
	pluralName: "AlertingSummaries"
}

alertingSummaryv0alpha1: alertingSummaryKind & {
	schema: {
		spec:   v0alpha1.AlertingSummarySpec
		status: v0alpha1.AlertingSummaryStatus
	}
}
