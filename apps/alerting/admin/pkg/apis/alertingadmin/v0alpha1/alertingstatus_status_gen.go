// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// AlertingStatusAlertmanager groups runtime observations for the alerting
// stack. Top-level keys mirror feature groupings on AlertingConfig.spec.alertmanager.
// +k8s:openapi-gen=true
type AlertingStatusAlertingStatusAlertmanager struct {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors AlertingConfig.spec.alertmanager.externalSync.
	ExternalSync *AlertingStatusExternalAlertmanagerSyncStatus `json:"externalSync,omitempty"`
}

// NewAlertingStatusAlertingStatusAlertmanager creates a new AlertingStatusAlertingStatusAlertmanager object.
func NewAlertingStatusAlertingStatusAlertmanager() *AlertingStatusAlertingStatusAlertmanager {
	return &AlertingStatusAlertingStatusAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusAlertingStatusAlertmanager.
func (AlertingStatusAlertingStatusAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusAlertingStatusAlertmanager"
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
type AlertingStatusExternalAlertmanagerSyncStatus struct {
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
	Origin *AlertingStatusExternalAlertmanagerSyncStatusOrigin `json:"origin,omitempty"`
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
	Conditions []AlertingStatusCondition `json:"conditions,omitempty"`
}

// NewAlertingStatusExternalAlertmanagerSyncStatus creates a new AlertingStatusExternalAlertmanagerSyncStatus object.
func NewAlertingStatusExternalAlertmanagerSyncStatus() *AlertingStatusExternalAlertmanagerSyncStatus {
	return &AlertingStatusExternalAlertmanagerSyncStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusExternalAlertmanagerSyncStatus.
func (AlertingStatusExternalAlertmanagerSyncStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusExternalAlertmanagerSyncStatus"
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
type AlertingStatusCondition struct {
	Type   string                        `json:"type"`
	Status AlertingStatusConditionStatus `json:"status"`
	// RFC3339
	LastTransitionTime string  `json:"lastTransitionTime"`
	Reason             string  `json:"reason"`
	Message            *string `json:"message,omitempty"`
	ObservedGeneration *int64  `json:"observedGeneration,omitempty"`
}

// NewAlertingStatusCondition creates a new AlertingStatusCondition object.
func NewAlertingStatusCondition() *AlertingStatusCondition {
	return &AlertingStatusCondition{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusCondition.
func (AlertingStatusCondition) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusCondition"
}

// +k8s:openapi-gen=true
type AlertingStatusstatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State AlertingStatusStatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewAlertingStatusstatusOperatorState creates a new AlertingStatusstatusOperatorState object.
func NewAlertingStatusstatusOperatorState() *AlertingStatusstatusOperatorState {
	return &AlertingStatusstatusOperatorState{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusstatusOperatorState.
func (AlertingStatusstatusOperatorState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusstatusOperatorState"
}

// +k8s:openapi-gen=true
type AlertingStatusStatus struct {
	// alertmanager groups observations for the per-org alerting stack.
	Alertmanager *AlertingStatusAlertingStatusAlertmanager `json:"alertmanager,omitempty"`
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]AlertingStatusstatusOperatorState `json:"operatorStates,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewAlertingStatusStatus creates a new AlertingStatusStatus object.
func NewAlertingStatusStatus() *AlertingStatusStatus {
	return &AlertingStatusStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusStatus.
func (AlertingStatusStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusStatus"
}

// +k8s:openapi-gen=true
type AlertingStatusExternalAlertmanagerSyncStatusOrigin string

const (
	AlertingStatusExternalAlertmanagerSyncStatusOriginApi AlertingStatusExternalAlertmanagerSyncStatusOrigin = "api"
	AlertingStatusExternalAlertmanagerSyncStatusOriginIni AlertingStatusExternalAlertmanagerSyncStatusOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusExternalAlertmanagerSyncStatusOrigin.
func (AlertingStatusExternalAlertmanagerSyncStatusOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusExternalAlertmanagerSyncStatusOrigin"
}

// +k8s:openapi-gen=true
type AlertingStatusConditionStatus string

const (
	AlertingStatusConditionStatusTrue    AlertingStatusConditionStatus = "True"
	AlertingStatusConditionStatusFalse   AlertingStatusConditionStatus = "False"
	AlertingStatusConditionStatusUnknown AlertingStatusConditionStatus = "Unknown"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusConditionStatus.
func (AlertingStatusConditionStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusConditionStatus"
}

// +k8s:openapi-gen=true
type AlertingStatusStatusOperatorStateState string

const (
	AlertingStatusStatusOperatorStateStateSuccess    AlertingStatusStatusOperatorStateState = "success"
	AlertingStatusStatusOperatorStateStateInProgress AlertingStatusStatusOperatorStateState = "in_progress"
	AlertingStatusStatusOperatorStateStateFailed     AlertingStatusStatusOperatorStateState = "failed"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusStatusOperatorStateState.
func (AlertingStatusStatusOperatorStateState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusStatusOperatorStateState"
}
