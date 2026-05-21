package v0alpha1

// SummaryStatus is the area-grouped observation payload returned on the
// Summary kind. Sub-keys mirror Config.spec one-to-one so spec and status
// read symmetrically:
//
//   Config.spec.alertmanager.externalSync.datasourceUid     (intent)
//   Summary.status.alertmanager.externalSync.{...}          (observation)
//
// Sub-trees are omitted when no corresponding observation exists yet;
// absence means not yet observed, not empty.
SummaryStatus: {
	// alertmanager groups observations for the per-org alerting stack.
	alertmanager?: SummaryAlertmanager
}

// SummaryAlertmanager groups runtime observations for the alerting stack.
// Top-level keys mirror feature groupings on Config.spec.alertmanager.
SummaryAlertmanager: {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors Config.spec.alertmanager.externalSync.
	externalSync?: ExternalAlertmanagerSyncStatus
}
