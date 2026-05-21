// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type ConfigstatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State ConfigStatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewConfigstatusOperatorState creates a new ConfigstatusOperatorState object.
func NewConfigstatusOperatorState() *ConfigstatusOperatorState {
	return &ConfigstatusOperatorState{}
}

// OpenAPIModelName returns the OpenAPI model name for ConfigstatusOperatorState.
func (ConfigstatusOperatorState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigstatusOperatorState"
}

// +k8s:openapi-gen=true
type ConfigStatus struct {
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]ConfigstatusOperatorState `json:"operatorStates,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewConfigStatus creates a new ConfigStatus object.
func NewConfigStatus() *ConfigStatus {
	return &ConfigStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for ConfigStatus.
func (ConfigStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigStatus"
}

// +k8s:openapi-gen=true
type ConfigStatusOperatorStateState string

const (
	ConfigStatusOperatorStateStateSuccess    ConfigStatusOperatorStateState = "success"
	ConfigStatusOperatorStateStateInProgress ConfigStatusOperatorStateState = "in_progress"
	ConfigStatusOperatorStateStateFailed     ConfigStatusOperatorStateState = "failed"
)

// OpenAPIModelName returns the OpenAPI model name for ConfigStatusOperatorStateState.
func (ConfigStatusOperatorStateState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigStatusOperatorStateState"
}
