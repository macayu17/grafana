// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type AlertingConfigSpec struct {
	// alertmanager groups admin settings for the per-org alerting stack.
	Alertmanager *AlertingConfigV0alpha1SpecAlertmanager `json:"alertmanager,omitempty"`
}

// NewAlertingConfigSpec creates a new AlertingConfigSpec object.
func NewAlertingConfigSpec() *AlertingConfigSpec {
	return &AlertingConfigSpec{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigSpec.
func (AlertingConfigSpec) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigSpec"
}

// +k8s:openapi-gen=true
type AlertingConfigV0alpha1SpecAlertmanagerExternalSync struct {
	// datasourceUid is the UID of the Mimir/Cortex Alertmanager
	// datasource to sync configuration from. Empty (omitted) means
	// no per-org sync is configured. The operator-level
	// unified_alerting.external_alertmanager_uid ini setting still
	// wins over this when set — runtime observation of which source
	// is active lives on the ExternalAlertmanagerSync resource
	// (status.origin).
	DatasourceUid *string `json:"datasourceUid,omitempty"`
}

// NewAlertingConfigV0alpha1SpecAlertmanagerExternalSync creates a new AlertingConfigV0alpha1SpecAlertmanagerExternalSync object.
func NewAlertingConfigV0alpha1SpecAlertmanagerExternalSync() *AlertingConfigV0alpha1SpecAlertmanagerExternalSync {
	return &AlertingConfigV0alpha1SpecAlertmanagerExternalSync{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigV0alpha1SpecAlertmanagerExternalSync.
func (AlertingConfigV0alpha1SpecAlertmanagerExternalSync) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigV0alpha1SpecAlertmanagerExternalSync"
}

// +k8s:openapi-gen=true
type AlertingConfigV0alpha1SpecAlertmanager struct {
	// externalSync groups admin settings for the external Alertmanager
	// configuration sync worker. The worker fetches the alertmanager
	// configuration from a Mimir/Cortex datasource and merges it into the
	// org's local alertmanager configuration on each MAM sync tick.
	ExternalSync *AlertingConfigV0alpha1SpecAlertmanagerExternalSync `json:"externalSync,omitempty"`
}

// NewAlertingConfigV0alpha1SpecAlertmanager creates a new AlertingConfigV0alpha1SpecAlertmanager object.
func NewAlertingConfigV0alpha1SpecAlertmanager() *AlertingConfigV0alpha1SpecAlertmanager {
	return &AlertingConfigV0alpha1SpecAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigV0alpha1SpecAlertmanager.
func (AlertingConfigV0alpha1SpecAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigV0alpha1SpecAlertmanager"
}
