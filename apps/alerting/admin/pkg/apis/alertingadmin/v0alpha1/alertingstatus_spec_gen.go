// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type AlertingStatusSpec struct {
	// Reserved. Writers are rejected by storage; readers should ignore.
	Reserved *string `json:"reserved,omitempty"`
}

// NewAlertingStatusSpec creates a new AlertingStatusSpec object.
func NewAlertingStatusSpec() *AlertingStatusSpec {
	return &AlertingStatusSpec{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingStatusSpec.
func (AlertingStatusSpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingStatusSpec"
}
