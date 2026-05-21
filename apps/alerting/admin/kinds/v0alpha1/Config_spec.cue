package v0alpha1

// Config is the per-org alerting admin config — a singleton resource carrying
// admin-controllable settings for the alerting stack. Future admin toggles
// (e.g. simplified-routing) land here as siblings on .spec.
ConfigSpec: {
	// UID of the Mimir/Cortex Alertmanager datasource to sync configuration
	// from. Empty (omitted) means no per-org sync is configured. The
	// operator-level unified_alerting.external_alertmanager_uid ini setting
	// still wins over this when set — runtime observation of which source is
	// active lives on the ExternalAlertmanagerSync resource (status.origin).
	externalAlertmanagerUid?: string
}
