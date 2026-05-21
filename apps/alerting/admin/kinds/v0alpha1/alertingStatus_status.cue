package v0alpha1

// AlertingStatusStatus is the area-grouped observation payload returned on
// the AlertingStatus kind. Sub-keys mirror AlertingConfig.spec one-to-one so
// spec and status read symmetrically:
//
//   AlertingConfig.spec.alertmanager.externalSync.datasourceUid     (intent)
//   AlertingStatus.status.alertmanager.externalSync.{...}           (observation)
//
// Sub-trees are omitted when no corresponding observation exists yet;
// absence means not yet observed, not empty.
AlertingStatusStatus: {
	// alertmanager groups observations for the per-org alerting stack.
	alertmanager?: AlertingStatusAlertmanager
}

// AlertingStatusAlertmanager groups runtime observations for the alerting
// stack. Top-level keys mirror feature groupings on AlertingConfig.spec.alertmanager.
AlertingStatusAlertmanager: {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors AlertingConfig.spec.alertmanager.externalSync.
	externalSync?: ExternalAlertmanagerSyncStatus
}
