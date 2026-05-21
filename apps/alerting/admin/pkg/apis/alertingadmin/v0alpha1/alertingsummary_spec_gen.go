// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type AlertingSummarySpec struct {
	// Reserved. Writers are rejected by storage; readers should ignore.
	Reserved *string `json:"reserved,omitempty"`
}

// NewAlertingSummarySpec creates a new AlertingSummarySpec object.
func NewAlertingSummarySpec() *AlertingSummarySpec {
	return &AlertingSummarySpec{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingSummarySpec.
func (AlertingSummarySpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingSummarySpec"
}
