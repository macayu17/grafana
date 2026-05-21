// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncSpec struct {
	// Reserved. Always omitted by writers; ignored by readers.
	Reserved *string `json:"reserved,omitempty"`
}

// NewExternalAlertmanagerSyncSpec creates a new ExternalAlertmanagerSyncSpec object.
func NewExternalAlertmanagerSyncSpec() *ExternalAlertmanagerSyncSpec {
	return &ExternalAlertmanagerSyncSpec{}
}

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncSpec.
func (ExternalAlertmanagerSyncSpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncSpec"
}
