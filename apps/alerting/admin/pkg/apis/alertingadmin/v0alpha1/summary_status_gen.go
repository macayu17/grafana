// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// SummaryAlertmanager groups runtime observations for the alerting stack.
// Top-level keys mirror feature groupings on Config.spec.alertmanager.
// +k8s:openapi-gen=true
type SummarySummaryAlertmanager struct {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors Config.spec.alertmanager.externalSync.
	ExternalSync *SummaryExternalAlertmanagerSyncStatus `json:"externalSync,omitempty"`
}

// NewSummarySummaryAlertmanager creates a new SummarySummaryAlertmanager object.
func NewSummarySummaryAlertmanager() *SummarySummaryAlertmanager {
	return &SummarySummaryAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for SummarySummaryAlertmanager.
func (SummarySummaryAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummarySummaryAlertmanager"
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
// +k8s:openapi-gen=true
type SummaryExternalAlertmanagerSyncStatus struct {
	// observedGeneration is the spec.generation last evaluated by the syncer.
	// Always 0 today since the spec is empty; carried for forward compatibility
	// with the conditions pattern.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// UID actually used on the last sync attempt. May differ from
	// `Config.spec.alertmanager.externalSync.datasourceUid` immediately after
	// a spec change, until the next tick. When `origin = "ini"`, this is the
	// grafana.ini override value.
	DatasourceUid *string `json:"datasourceUid,omitempty"`
	// Which source supplied datasourceUid on the last run:
	//   - "api": value from Config.spec.alertmanager.externalSync.datasourceUid
	//     (set by an admin via the k8s API).
	//   - "ini": grafana.ini override (`[unified_alerting]
	//     external_alertmanager_uid`), set by the server operator. Wins over
	//     api when both are present.
	// This field lets clients see when an ini override is in effect without
	// having to know the precedence rule.
	Origin *SummaryExternalAlertmanagerSyncStatusOrigin `json:"origin,omitempty"`
	// Unix epoch seconds of the most recent successful sync. Preserved across
	// failure streaks — answers "even though it's broken now, when did it
	// last work?". Not derivable from `Synced.lastTransitionTime` (which
	// marks when the current state was entered, not when success last held).
	// Omitted when no sync has ever succeeded for this org.
	LastSuccessAt *int64 `json:"lastSuccessAt,omitempty"`
	// Standard k8s-style condition list. v1 carries one type:
	//   - Synced: True after a successful sync, False after a failed attempt,
	//     Unknown until the first attempt has run.
	// Future state dimensions land here as additional condition types.
	Conditions []SummaryCondition `json:"conditions,omitempty"`
}

// NewSummaryExternalAlertmanagerSyncStatus creates a new SummaryExternalAlertmanagerSyncStatus object.
func NewSummaryExternalAlertmanagerSyncStatus() *SummaryExternalAlertmanagerSyncStatus {
	return &SummaryExternalAlertmanagerSyncStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for SummaryExternalAlertmanagerSyncStatus.
func (SummaryExternalAlertmanagerSyncStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummaryExternalAlertmanagerSyncStatus"
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
//
// +k8s:openapi-gen=true
type SummaryCondition struct {
	Type   string                 `json:"type"`
	Status SummaryConditionStatus `json:"status"`
	// RFC3339
	LastTransitionTime string  `json:"lastTransitionTime"`
	Reason             string  `json:"reason"`
	Message            *string `json:"message,omitempty"`
	ObservedGeneration *int64  `json:"observedGeneration,omitempty"`
}

// NewSummaryCondition creates a new SummaryCondition object.
func NewSummaryCondition() *SummaryCondition {
	return &SummaryCondition{}
}

// OpenAPIModelName returns the OpenAPI model name for SummaryCondition.
func (SummaryCondition) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummaryCondition"
}

// +k8s:openapi-gen=true
type SummarystatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State SummaryStatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewSummarystatusOperatorState creates a new SummarystatusOperatorState object.
func NewSummarystatusOperatorState() *SummarystatusOperatorState {
	return &SummarystatusOperatorState{}
}

// OpenAPIModelName returns the OpenAPI model name for SummarystatusOperatorState.
func (SummarystatusOperatorState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummarystatusOperatorState"
}

// +k8s:openapi-gen=true
type SummaryStatus struct {
	// alertmanager groups observations for the per-org alerting stack.
	Alertmanager *SummarySummaryAlertmanager `json:"alertmanager,omitempty"`
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]SummarystatusOperatorState `json:"operatorStates,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewSummaryStatus creates a new SummaryStatus object.
func NewSummaryStatus() *SummaryStatus {
	return &SummaryStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for SummaryStatus.
func (SummaryStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummaryStatus"
}

// +k8s:openapi-gen=true
type SummaryExternalAlertmanagerSyncStatusOrigin string

const (
	SummaryExternalAlertmanagerSyncStatusOriginApi SummaryExternalAlertmanagerSyncStatusOrigin = "api"
	SummaryExternalAlertmanagerSyncStatusOriginIni SummaryExternalAlertmanagerSyncStatusOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for SummaryExternalAlertmanagerSyncStatusOrigin.
func (SummaryExternalAlertmanagerSyncStatusOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummaryExternalAlertmanagerSyncStatusOrigin"
}

// +k8s:openapi-gen=true
type SummaryConditionStatus string

const (
	SummaryConditionStatusTrue    SummaryConditionStatus = "True"
	SummaryConditionStatusFalse   SummaryConditionStatus = "False"
	SummaryConditionStatusUnknown SummaryConditionStatus = "Unknown"
)

// OpenAPIModelName returns the OpenAPI model name for SummaryConditionStatus.
func (SummaryConditionStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummaryConditionStatus"
}

// +k8s:openapi-gen=true
type SummaryStatusOperatorStateState string

const (
	SummaryStatusOperatorStateStateSuccess    SummaryStatusOperatorStateState = "success"
	SummaryStatusOperatorStateStateInProgress SummaryStatusOperatorStateState = "in_progress"
	SummaryStatusOperatorStateStateFailed     SummaryStatusOperatorStateState = "failed"
)

// OpenAPIModelName returns the OpenAPI model name for SummaryStatusOperatorStateState.
func (SummaryStatusOperatorStateState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummaryStatusOperatorStateState"
}
