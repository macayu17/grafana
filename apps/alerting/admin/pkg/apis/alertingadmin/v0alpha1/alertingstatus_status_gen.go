// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// AlertingStatusAlertmanager groups runtime observations for the alerting
// stack. Top-level keys mirror feature groupings on
// AlertingConfig.spec.alertmanager.
// +k8s:openapi-gen=true
type AlertingStatusAlertingStatusAlertmanager struct {
	// externalSync is the observation payload for the external Alertmanager
	// configuration sync worker. Mirrors
	// AlertingConfig.spec.alertmanager.externalSync, plus the matching
	// Synced condition routed from AlertingConfig.status.conditions[].
	ExternalSync *AlertingStatusAlertingStatusExternalSync `json:"externalSync,omitempty"`
}

// NewAlertingStatusAlertingStatusAlertmanager creates a new AlertingStatusAlertingStatusAlertmanager object.
func NewAlertingStatusAlertingStatusAlertmanager() *AlertingStatusAlertingStatusAlertmanager {
	return &AlertingStatusAlertingStatusAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusAlertingStatusAlertmanager.
func (AlertingStatusAlertingStatusAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusAlertingStatusAlertmanager"
}

// AlertingStatusExternalSync carries the observation for the external
// Alertmanager configuration sync worker. The aggregate handler builds it
// by copying AlertingConfig.status.alertmanager.externalSync (auxiliary
// fields) and projecting AlertingConfig.status.conditions[type=Synced]
// into conditions.
// +k8s:openapi-gen=true
type AlertingStatusAlertingStatusExternalSync struct {
	DatasourceUid *string                                         `json:"datasourceUid,omitempty"`
	Origin        *AlertingStatusAlertingStatusExternalSyncOrigin `json:"origin,omitempty"`
	Conditions    []AlertingStatusCondition                       `json:"conditions,omitempty"`
}

// NewAlertingStatusAlertingStatusExternalSync creates a new AlertingStatusAlertingStatusExternalSync object.
func NewAlertingStatusAlertingStatusExternalSync() *AlertingStatusAlertingStatusExternalSync {
	return &AlertingStatusAlertingStatusExternalSync{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusAlertingStatusExternalSync.
func (AlertingStatusAlertingStatusExternalSync) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusAlertingStatusExternalSync"
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
type AlertingStatusAlertingStatusExternalSyncOrigin string

const (
	AlertingStatusAlertingStatusExternalSyncOriginApi AlertingStatusAlertingStatusExternalSyncOrigin = "api"
	AlertingStatusAlertingStatusExternalSyncOriginIni AlertingStatusAlertingStatusExternalSyncOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusAlertingStatusExternalSyncOrigin.
func (AlertingStatusAlertingStatusExternalSyncOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusAlertingStatusExternalSyncOrigin"
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
