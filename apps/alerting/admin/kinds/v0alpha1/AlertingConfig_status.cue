package v0alpha1

// AlertingConfigStatus reports the runtime observation of admin alerting
// concerns for an org. Written by the sync workers and other controllers
// that own fields on spec; clients read only.
//
// Shape: conditions at the top level, auxiliary state nested by area of
// concern to mirror spec. Conditions stay top-level because:
//   - k8s convention: meta.SetStatusCondition, kubectl wait
//     --for=condition=, and controller-runtime helpers all expect
//     .status.conditions.
//   - Condition Type names disambiguate concerns (e.g. "Synced" for
//     external sync; future concerns add their own types).
//
// State is modelled with the standard k8s Conditions FSM pattern. Status
// writes only happen on condition transitions (managed by
// meta.SetStatusCondition — LastTransitionTime advances only when Status
// flips) so the resource's history budget is preserved for spec audit.
//
// Tick-by-tick liveness ("when did the last attempt happen, was it
// successful") is observable via metrics rather than this resource —
// see ExternalAMConfigSyncTotal / ExternalAMConfigSyncDuration.
// Heartbeats belong on metrics; status reports state transitions.
AlertingConfigStatus: {
	// observedGeneration is the spec.generation last evaluated by the
	// controllers writing this status. Carried for forward compatibility
	// with the conditions pattern.
	observedGeneration?: int

	// alertmanager groups runtime observations for the per-org alerting
	// stack. Sub-trees mirror spec.alertmanager structure 1:1 so spec and
	// auxiliary status read symmetrically.
	alertmanager?: {
		// externalSync carries the observation context for the external
		// Alertmanager configuration sync worker. Conditions about this
		// concern live at .status.conditions[type=Synced], not here.
		externalSync?: {
			// datasourceUid is the UID actually used on the last sync
			// attempt. May differ from
			// spec.alertmanager.externalSync.datasourceUid immediately
			// after a spec change, until the next tick. When
			// `origin = "ini"`, this is the grafana.ini override value.
			datasourceUid?: string

			// origin records which source supplied datasourceUid on the
			// last run:
			//   - "api": value from spec.alertmanager.externalSync.datasourceUid
			//     (set by an admin via the k8s API).
			//   - "ini": grafana.ini override (`[unified_alerting]
			//     external_alertmanager_uid`), set by the server operator.
			//     Wins over api when both are present.
			origin?: "api" | "ini"
		}
	}

	// Standard k8s-style condition list. Each binary-state concern owns
	// one condition type. Current types:
	//   - Synced: True after a successful external Alertmanager sync,
	//     False after a failed attempt, Unknown until the first attempt
	//     has run.
	// Future state dimensions land here as additional condition types.
	conditions?: [...#Condition]
}

// Condition mirrors metav1.Condition. Declared inline because the app-sdk
// codegen in this repo does not yet have a built-in path for referencing
// the k8s metav1.Condition type from CUE. Field semantics are k8s-standard:
//   - status flips between True/False/Unknown.
//   - lastTransitionTime advances only when status flips (managed by
//     meta.SetStatusCondition in the sync worker).
//   - reason is a PascalCase machine-readable enum (e.g. "SyncSucceeded",
//     "MimirFetchFailed"); see SyncReason in the syncer.
//   - message is human-readable detail.
//   - observedGeneration records the spec.generation this condition
//     evaluation reflects, when applicable.
#Condition: {
	type:                string
	status:              "True" | "False" | "Unknown"
	lastTransitionTime:  string // RFC3339
	reason:              string
	message?:            string
	observedGeneration?: int
}
