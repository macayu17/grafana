// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type SummarySpec struct {
	// Reserved. Writers are rejected by storage; readers should ignore.
	Reserved *string `json:"reserved,omitempty"`
}

// NewSummarySpec creates a new SummarySpec object.
func NewSummarySpec() *SummarySpec {
	return &SummarySpec{}
}

// OpenAPIModelName returns the OpenAPI model name for SummarySpec.
func (SummarySpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.SummarySpec"
}
