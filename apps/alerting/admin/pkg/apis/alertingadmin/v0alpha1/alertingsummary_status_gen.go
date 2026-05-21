// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// AlertingSummaryAlertmanager groups runtime observations for the alerting
// stack. Top-level keys mirror feature groupings on AlertingConfig.spec.alertmanager.
// +k8s:openapi-gen=true
type AlertingSummaryAlertingSummaryAlertmanager struct {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors AlertingConfig.spec.alertmanager.externalSync.
	ExternalSync *AlertingSummaryExternalAlertmanagerSyncStatus `json:"externalSync,omitempty"`
}

// NewAlertingSummaryAlertingSummaryAlertmanager creates a new AlertingSummaryAlertingSummaryAlertmanager object.
func NewAlertingSummaryAlertingSummaryAlertmanager() *AlertingSummaryAlertingSummaryAlertmanager {
	return &AlertingSummaryAlertingSummaryAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryAlertingSummaryAlertmanager.
func (AlertingSummaryAlertingSummaryAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryAlertingSummaryAlertmanager"
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
type AlertingSummaryExternalAlertmanagerSyncStatus struct {
	// observedGeneration is the spec.generation last evaluated by the syncer.
	// Always 0 today since the spec is empty; carried for forward compatibility
	// with the conditions pattern.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// UID actually used on the last sync attempt. May differ from
	// `AlertingConfig.spec.alertmanager.externalSync.datasourceUid` immediately after
	// a spec change, until the next tick. When `origin = "ini"`, this is the
	// grafana.ini override value.
	DatasourceUid *string `json:"datasourceUid,omitempty"`
	// Which source supplied datasourceUid on the last run:
	//   - "api": value from AlertingConfig.spec.alertmanager.externalSync.datasourceUid
	//     (set by an admin via the k8s API).
	//   - "ini": grafana.ini override (`[unified_alerting]
	//     external_alertmanager_uid`), set by the server operator. Wins over
	//     api when both are present.
	// This field lets clients see when an ini override is in effect without
	// having to know the precedence rule.
	Origin *AlertingSummaryExternalAlertmanagerSyncStatusOrigin `json:"origin,omitempty"`
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
	Conditions []AlertingSummaryCondition `json:"conditions,omitempty"`
}

// NewAlertingSummaryExternalAlertmanagerSyncStatus creates a new AlertingSummaryExternalAlertmanagerSyncStatus object.
func NewAlertingSummaryExternalAlertmanagerSyncStatus() *AlertingSummaryExternalAlertmanagerSyncStatus {
	return &AlertingSummaryExternalAlertmanagerSyncStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryExternalAlertmanagerSyncStatus.
func (AlertingSummaryExternalAlertmanagerSyncStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryExternalAlertmanagerSyncStatus"
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
type AlertingSummaryCondition struct {
	Type   string                         `json:"type"`
	Status AlertingSummaryConditionStatus `json:"status"`
	// RFC3339
	LastTransitionTime string  `json:"lastTransitionTime"`
	Reason             string  `json:"reason"`
	Message            *string `json:"message,omitempty"`
	ObservedGeneration *int64  `json:"observedGeneration,omitempty"`
}

// NewAlertingSummaryCondition creates a new AlertingSummaryCondition object.
func NewAlertingSummaryCondition() *AlertingSummaryCondition {
	return &AlertingSummaryCondition{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryCondition.
func (AlertingSummaryCondition) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryCondition"
}

// +k8s:openapi-gen=true
type AlertingSummarystatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State AlertingSummaryStatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewAlertingSummarystatusOperatorState creates a new AlertingSummarystatusOperatorState object.
func NewAlertingSummarystatusOperatorState() *AlertingSummarystatusOperatorState {
	return &AlertingSummarystatusOperatorState{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingSummarystatusOperatorState.
func (AlertingSummarystatusOperatorState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummarystatusOperatorState"
}

// +k8s:openapi-gen=true
type AlertingSummaryStatus struct {
	// alertmanager groups observations for the per-org alerting stack.
	Alertmanager *AlertingSummaryAlertingSummaryAlertmanager `json:"alertmanager,omitempty"`
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]AlertingSummarystatusOperatorState `json:"operatorStates,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewAlertingSummaryStatus creates a new AlertingSummaryStatus object.
func NewAlertingSummaryStatus() *AlertingSummaryStatus {
	return &AlertingSummaryStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryStatus.
func (AlertingSummaryStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryStatus"
}

// +k8s:openapi-gen=true
type AlertingSummaryExternalAlertmanagerSyncStatusOrigin string

const (
	AlertingSummaryExternalAlertmanagerSyncStatusOriginApi AlertingSummaryExternalAlertmanagerSyncStatusOrigin = "api"
	AlertingSummaryExternalAlertmanagerSyncStatusOriginIni AlertingSummaryExternalAlertmanagerSyncStatusOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryExternalAlertmanagerSyncStatusOrigin.
func (AlertingSummaryExternalAlertmanagerSyncStatusOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryExternalAlertmanagerSyncStatusOrigin"
}

// +k8s:openapi-gen=true
type AlertingSummaryConditionStatus string

const (
	AlertingSummaryConditionStatusTrue    AlertingSummaryConditionStatus = "True"
	AlertingSummaryConditionStatusFalse   AlertingSummaryConditionStatus = "False"
	AlertingSummaryConditionStatusUnknown AlertingSummaryConditionStatus = "Unknown"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryConditionStatus.
func (AlertingSummaryConditionStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryConditionStatus"
}

// +k8s:openapi-gen=true
type AlertingSummaryStatusOperatorStateState string

const (
	AlertingSummaryStatusOperatorStateStateSuccess    AlertingSummaryStatusOperatorStateState = "success"
	AlertingSummaryStatusOperatorStateStateInProgress AlertingSummaryStatusOperatorStateState = "in_progress"
	AlertingSummaryStatusOperatorStateStateFailed     AlertingSummaryStatusOperatorStateState = "failed"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingSummaryStatusOperatorStateState.
func (AlertingSummaryStatusOperatorStateState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummaryStatusOperatorStateState"
}
