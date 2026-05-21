package v0alpha1

// AlertingStatusStatus is the area-grouped observation payload returned on
// the AlertingStatus synthetic kind. Sub-keys mirror AlertingConfig.spec
// one-to-one and bundle each concern's auxiliary state alongside its
// condition for an at-a-glance view:
//
//   AlertingConfig.spec.alertmanager.externalSync.datasourceUid     (intent)
//   AlertingStatus.status.alertmanager.externalSync.{...}           (observation)
//
// The aggregate handler in pkg/app/status_storage.go composes this by:
//   - copying AlertingConfig.status.alertmanager.<area> sub-trees
//   - filtering AlertingConfig.status.conditions[] by Type and routing
//     each into the matching area's conditions list.
//
// Sub-trees are omitted when no corresponding observation exists yet;
// absence means not yet observed, not empty.
AlertingStatusStatus: {
	// alertmanager groups observations for the per-org alerting stack.
	alertmanager?: AlertingStatusAlertmanager
}

// AlertingStatusAlertmanager groups runtime observations for the alerting
// stack. Top-level keys mirror feature groupings on
// AlertingConfig.spec.alertmanager.
AlertingStatusAlertmanager: {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors
	// AlertingConfig.spec.alertmanager.externalSync, plus the matching
	// Synced condition routed from AlertingConfig.status.conditions[].
	externalSync?: AlertingStatusExternalSync
}

// AlertingStatusExternalSync carries the observation for the external
// Alertmanager configuration sync worker. The aggregate handler builds it
// by copying AlertingConfig.status.alertmanager.externalSync (auxiliary
// fields) and projecting AlertingConfig.status.conditions[type=Synced]
// into conditions.
AlertingStatusExternalSync: {
	datasourceUid?: string
	origin?:        "api" | "ini"
	conditions?: [...#Condition]
}
