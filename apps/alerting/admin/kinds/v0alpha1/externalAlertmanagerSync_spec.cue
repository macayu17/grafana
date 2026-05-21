package v0alpha1

// Status resources under the admin.alerting.grafana.app group carry runtime
// observations only. ExternalAlertmanagerSync's spec is intentionally a no-op:
// admin-controllable configuration for external Alertmanager sync lives on
// Config.spec.externalAlertmanagerUid in the same group.
//
// The k8s machinery and app-sdk codegen require every kind to declare a spec
// type with at least one field. We carry a single reserved placeholder field
// to satisfy that requirement; do not add anything meaningful here — new
// admin-configurable settings belong on Config.spec.
ExternalAlertmanagerSyncSpec: {
	// Reserved. Always omitted by writers; ignored by readers.
	reserved?: string
}

// ExternalAlertmanagerSyncStatus reports the runtime state of the external
// Alertmanager configuration sync worker for an org. Written by the sync
// worker; clients read only.
//
// State is modelled with the standard k8s Conditions FSM pattern: the
// boolean "is sync currently working?" lives on the `Synced` condition,
// whose lastTransitionTime advances only when the boolean flips. Auxiliary
// fields below carry context (which UID was attempted, where it came from)
// and one piece of historical state (lastSuccessAt) that can't be recovered
// from a single condition.
//
// Tick-by-tick liveness (when the last attempt happened) is observable via
// the syncer's metrics rather than this resource — see
// ExternalAMConfigSyncTotal / ExternalAMConfigSyncDuration.
ExternalAlertmanagerSyncStatus: {
	// observedGeneration is the spec.generation last evaluated by the syncer.
	// Always 0 today since the spec is empty; carried for forward compatibility
	// with the conditions pattern.
	observedGeneration?: int

	// UID actually used on the last sync attempt. May differ from
	// `Config.spec.externalAlertmanagerUid` immediately after a spec change,
	// until the next tick. When `origin = "ini"`, this is the grafana.ini
	// override value.
	datasourceUid?: string

	// Which source supplied datasourceUid on the last run:
	//   - "api": value from Config.spec.externalAlertmanagerUid (set by an
	//     admin via the k8s API).
	//   - "ini": grafana.ini override (`[unified_alerting]
	//     external_alertmanager_uid`), set by the server operator. Wins over
	//     api when both are present.
	// This field lets clients see when an ini override is in effect without
	// having to know the precedence rule.
	origin?: "api" | "ini"

	// Unix epoch seconds of the most recent successful sync. Preserved across
	// failure streaks — answers "even though it's broken now, when did it
	// last work?". Not derivable from `Synced.lastTransitionTime` (which
	// marks when the current state was entered, not when success last held).
	// Omitted when no sync has ever succeeded for this org.
	lastSuccessAt?: int

	// Standard k8s-style condition list. v1 carries one type:
	//   - Synced: True after a successful sync, False after a failed attempt,
	//     Unknown until the first attempt has run.
	// Future state dimensions land here as additional condition types.
	conditions?: [...#Condition]
}

// Condition mirrors metav1.Condition. Declared inline because the app-sdk
// codegen in this repo does not yet have a built-in path for referencing the
// k8s metav1.Condition type from CUE. Field semantics are k8s-standard:
//   - status flips between True/False/Unknown.
//   - lastTransitionTime advances only when status flips.
//   - reason is a PascalCase machine-readable enum (e.g. "SyncSucceeded",
//     "MimirFetchFailed"); see SyncReason in the syncer.
//   - message is human-readable detail.
//   - observedGeneration records the spec.generation this condition evaluation
//     reflects, when applicable.
#Condition: {
	type:               string
	status:             "True" | "False" | "Unknown"
	lastTransitionTime: string // RFC3339
	reason:             string
	message?:           string
	observedGeneration?: int
}
