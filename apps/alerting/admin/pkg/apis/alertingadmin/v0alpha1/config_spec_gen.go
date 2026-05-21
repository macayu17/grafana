// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type ConfigSpec struct {
	// alertmanager groups admin settings for the per-org alerting stack.
	Alertmanager *ConfigV0alpha1SpecAlertmanager `json:"alertmanager,omitempty"`
}

// NewConfigSpec creates a new ConfigSpec object.
func NewConfigSpec() *ConfigSpec {
	return &ConfigSpec{}
}

// OpenAPIModelName returns the OpenAPI model name for ConfigSpec.
func (ConfigSpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigSpec"
}

// +k8s:openapi-gen=true
type ConfigV0alpha1SpecAlertmanagerExternalSync struct {
	// datasourceUid is the UID of the Mimir/Cortex Alertmanager
	// datasource to sync configuration from. Empty (omitted) means
	// no per-org sync is configured. The operator-level
	// unified_alerting.external_alertmanager_uid ini setting still
	// wins over this when set — runtime observation of which source
	// is active lives on the ExternalAlertmanagerSync resource
	// (status.origin).
	DatasourceUid *string `json:"datasourceUid,omitempty"`
}

// NewConfigV0alpha1SpecAlertmanagerExternalSync creates a new ConfigV0alpha1SpecAlertmanagerExternalSync object.
func NewConfigV0alpha1SpecAlertmanagerExternalSync() *ConfigV0alpha1SpecAlertmanagerExternalSync {
	return &ConfigV0alpha1SpecAlertmanagerExternalSync{}
}

// OpenAPIModelName returns the OpenAPI model name for ConfigV0alpha1SpecAlertmanagerExternalSync.
func (ConfigV0alpha1SpecAlertmanagerExternalSync) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigV0alpha1SpecAlertmanagerExternalSync"
}

// +k8s:openapi-gen=true
type ConfigV0alpha1SpecAlertmanager struct {
	// externalSync groups admin settings for the external Alertmanager
	// configuration sync worker. The worker fetches the alertmanager
	// configuration from a Mimir/Cortex datasource and merges it into the
	// org's local alertmanager configuration on each MAM sync tick.
	ExternalSync *ConfigV0alpha1SpecAlertmanagerExternalSync `json:"externalSync,omitempty"`
}

// NewConfigV0alpha1SpecAlertmanager creates a new ConfigV0alpha1SpecAlertmanager object.
func NewConfigV0alpha1SpecAlertmanager() *ConfigV0alpha1SpecAlertmanager {
	return &ConfigV0alpha1SpecAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for ConfigV0alpha1SpecAlertmanager.
func (ConfigV0alpha1SpecAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ConfigV0alpha1SpecAlertmanager"
}
