// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type ConfigSpec struct {
	// UID of the Mimir/Cortex Alertmanager datasource to sync configuration
	// from. Empty (omitted) means no per-org sync is configured. The
	// operator-level unified_alerting.external_alertmanager_uid ini setting
	// still wins over this when set — runtime observation of which source is
	// active lives on the ExternalAlertmanagerSync resource (status.origin).
	ExternalAlertmanagerUid *string `json:"externalAlertmanagerUid,omitempty"`
}

// NewConfigSpec creates a new ConfigSpec object.
func NewConfigSpec() *ConfigSpec {
	return &ConfigSpec{}
}

// OpenAPIModelName returns the OpenAPI model name for ConfigSpec.
func (ConfigSpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigSpec"
}
