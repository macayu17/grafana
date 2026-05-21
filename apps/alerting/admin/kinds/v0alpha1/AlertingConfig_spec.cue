package v0alpha1

// AlertingConfig is the per-org alerting admin config — a singleton resource
// carrying admin-controllable settings for the alerting stack. Structure
// groups fields by area of concern (top level) → feature (second level) →
// knob (leaf). New admin toggles land under the appropriate area sub-object;
// the structure lets related knobs cluster naturally as the surface grows.
AlertingConfigSpec: {
	// alertmanager groups admin settings for the per-org alerting stack.
	alertmanager?: {
		// externalSync groups admin settings for the external Alertmanager
		// configuration sync worker. The worker fetches the alertmanager
		// configuration from a Mimir/Cortex datasource and merges it into the
		// org's local alertmanager configuration on each MAM sync tick.
		externalSync?: {
			// datasourceUid is the UID of the Mimir/Cortex Alertmanager
			// datasource to sync configuration from. Empty (omitted) means
			// no per-org sync is configured. The operator-level
			// unified_alerting.external_alertmanager_uid ini setting still
			// wins over this when set — runtime observation of which source
			// is active lives on the ExternalAlertmanagerSync resource
			// (status.origin).
			datasourceUid?: string
		}
	}
}
