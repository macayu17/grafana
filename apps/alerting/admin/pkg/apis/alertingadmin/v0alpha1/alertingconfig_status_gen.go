// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type AlertingConfigstatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State AlertingConfigStatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewAlertingConfigstatusOperatorState creates a new AlertingConfigstatusOperatorState object.
func NewAlertingConfigstatusOperatorState() *AlertingConfigstatusOperatorState {
	return &AlertingConfigstatusOperatorState{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigstatusOperatorState.
func (AlertingConfigstatusOperatorState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigstatusOperatorState"
}

// +k8s:openapi-gen=true
type AlertingConfigStatus struct {
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]AlertingConfigstatusOperatorState `json:"operatorStates,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewAlertingConfigStatus creates a new AlertingConfigStatus object.
func NewAlertingConfigStatus() *AlertingConfigStatus {
	return &AlertingConfigStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigStatus.
func (AlertingConfigStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigStatus"
}

// +k8s:openapi-gen=true
type AlertingConfigStatusOperatorStateState string

const (
	AlertingConfigStatusOperatorStateStateSuccess    AlertingConfigStatusOperatorStateState = "success"
	AlertingConfigStatusOperatorStateStateInProgress AlertingConfigStatusOperatorStateState = "in_progress"
	AlertingConfigStatusOperatorStateStateFailed     AlertingConfigStatusOperatorStateState = "failed"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigStatusOperatorStateState.
func (AlertingConfigStatusOperatorStateState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigStatusOperatorStateState"
}
