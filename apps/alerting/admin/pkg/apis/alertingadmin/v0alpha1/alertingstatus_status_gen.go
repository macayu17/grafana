// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// AlertingStatusExternalAlertmanagerSync carries the observation for the
// external Alertmanager configuration sync worker. The aggregate handler
// builds it by copying AlertingConfig.status.externalAlertmanagerSync
// (auxiliary fields) and projecting
// AlertingConfig.status.conditions[type=ExternalAlertmanagerSynced] into
// conditions.
// +k8s:openapi-gen=true
type AlertingStatusAlertingStatusExternalAlertmanagerSync struct {
	DatasourceUid *string                                                     `json:"datasourceUid,omitempty"`
	Origin        *AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin `json:"origin,omitempty"`
	Conditions    []AlertingStatusCondition                                   `json:"conditions,omitempty"`
}

// NewAlertingStatusAlertingStatusExternalAlertmanagerSync creates a new AlertingStatusAlertingStatusExternalAlertmanagerSync object.
func NewAlertingStatusAlertingStatusExternalAlertmanagerSync() *AlertingStatusAlertingStatusExternalAlertmanagerSync {
	return &AlertingStatusAlertingStatusExternalAlertmanagerSync{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusAlertingStatusExternalAlertmanagerSync.
func (AlertingStatusAlertingStatusExternalAlertmanagerSync) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSync"
}

// Condition mirrors metav1.Condition. Declared inline because the app-sdk
// codegen in this repo does not yet have a built-in path for referencing
// the k8s metav1.Condition type from CUE. Field semantics are k8s-standard:
//   - status flips between True/False/Unknown.
//   - lastTransitionTime advances only when status flips (managed by the
//     hand-rolled equivalent of meta.SetStatusCondition in the sync
//     worker, since AlertingConfigCondition is a codegen-emitted type
//     distinct from metav1.Condition).
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
	// externalAlertmanagerSync is the observation payload for the external
	// Alertmanager configuration sync worker. Mirrors the same key on
	// AlertingConfig.status, plus the matching ExternalAlertmanagerSynced
	// condition routed from AlertingConfig.status.conditions[].
	ExternalAlertmanagerSync *AlertingStatusAlertingStatusExternalAlertmanagerSync `json:"externalAlertmanagerSync,omitempty"`
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
type AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin string

const (
	AlertingStatusAlertingStatusExternalAlertmanagerSyncOriginApi AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin = "api"
	AlertingStatusAlertingStatusExternalAlertmanagerSyncOriginIni AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin.
func (AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusAlertingStatusExternalAlertmanagerSyncOrigin"
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
