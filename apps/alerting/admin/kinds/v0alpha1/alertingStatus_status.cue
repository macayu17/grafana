package v0alpha1

// AlertingStatusStatus is the per-feature observation payload returned on
// the AlertingStatus synthetic kind. Sub-keys mirror AlertingConfig.spec /
// AlertingConfig.status one-to-one so the aggregate view reads like a
// snapshot of the underlying feature config + observation:
//
//   AlertingConfig.spec.externalAlertmanagerSync.datasourceUid     (intent)
//   AlertingStatus.status.externalAlertmanagerSync.{...}           (observation)
//
// The aggregate handler in pkg/app/status_storage.go composes this by:
//   - copying AlertingConfig.status.<feature> sub-trees
//   - filtering AlertingConfig.status.conditions[] by Type and routing
//     each into the matching feature's conditions list.
//
// Sub-trees are omitted when no corresponding observation exists yet;
// absence means not yet observed, not empty.
AlertingStatusStatus: {
	// externalAlertmanagerSync is the observation payload for the external
	// Alertmanager configuration sync worker. Mirrors the same key on
	// AlertingConfig.status, plus the matching ExternalAlertmanagerSynced
	// condition routed from AlertingConfig.status.conditions[].
	externalAlertmanagerSync?: AlertingStatusExternalAlertmanagerSync
}

// AlertingStatusExternalAlertmanagerSync carries the observation for the
// external Alertmanager configuration sync worker. The aggregate handler
// builds it by copying AlertingConfig.status.externalAlertmanagerSync
// (auxiliary fields) and projecting
// AlertingConfig.status.conditions[type=ExternalAlertmanagerSynced] into
// conditions.
AlertingStatusExternalAlertmanagerSync: {
	datasourceUid?: string
	origin?:        "api" | "ini"
	conditions?: [...#Condition]
}
