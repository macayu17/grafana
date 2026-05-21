package v0alpha1

// AlertingSummaryStatus is the area-grouped observation payload returned on
// the AlertingSummary kind. Sub-keys mirror AlertingConfig.spec one-to-one
// so spec and status read symmetrically:
//
//   AlertingConfig.spec.alertmanager.externalSync.datasourceUid     (intent)
//   AlertingSummary.status.alertmanager.externalSync.{...}          (observation)
//
// Sub-trees are omitted when no corresponding observation exists yet;
// absence means not yet observed, not empty.
AlertingSummaryStatus: {
	// alertmanager groups observations for the per-org alerting stack.
	alertmanager?: AlertingSummaryAlertmanager
}

// AlertingSummaryAlertmanager groups runtime observations for the alerting
// stack. Top-level keys mirror feature groupings on AlertingConfig.spec.alertmanager.
AlertingSummaryAlertmanager: {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors AlertingConfig.spec.alertmanager.externalSync.
	externalSync?: ExternalAlertmanagerSyncStatus
}
