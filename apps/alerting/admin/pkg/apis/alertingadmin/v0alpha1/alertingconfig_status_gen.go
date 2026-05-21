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

// Condition mirrors metav1.Condition. Declared inline because the app-sdk
// codegen in this repo does not yet have a built-in path for referencing
// the k8s metav1.Condition type from CUE. Field semantics are k8s-standard:
//   - status flips between True/False/Unknown.
//   - lastTransitionTime advances only when status flips (managed by
//     meta.SetStatusCondition in the sync worker).
//   - reason is a PascalCase machine-readable enum (e.g. "SyncSucceeded",
//     "MimirFetchFailed"); see SyncReason in the syncer.
//   - message is human-readable detail.
//   - observedGeneration records the spec.generation this condition
//     evaluation reflects, when applicable.
//
// +k8s:openapi-gen=true
type AlertingConfigCondition struct {
	Type   string                        `json:"type"`
	Status AlertingConfigConditionStatus `json:"status"`
	// RFC3339
	LastTransitionTime string  `json:"lastTransitionTime"`
	Reason             string  `json:"reason"`
	Message            *string `json:"message,omitempty"`
	ObservedGeneration *int64  `json:"observedGeneration,omitempty"`
}

// NewAlertingConfigCondition creates a new AlertingConfigCondition object.
func NewAlertingConfigCondition() *AlertingConfigCondition {
	return &AlertingConfigCondition{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigCondition.
func (AlertingConfigCondition) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigCondition"
}

// +k8s:openapi-gen=true
type AlertingConfigStatus struct {
	// observedGeneration is the spec.generation last evaluated by the
	// controllers writing this status. Carried for forward compatibility
	// with the conditions pattern.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// alertmanager groups runtime observations for the per-org alerting
	// stack. Sub-trees mirror spec.alertmanager structure 1:1 so spec and
	// auxiliary status read symmetrically.
	Alertmanager *AlertingConfigV0alpha1StatusAlertmanager `json:"alertmanager,omitempty"`
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]AlertingConfigstatusOperatorState `json:"operatorStates,omitempty"`
	// Standard k8s-style condition list. Each binary-state concern owns
	// one condition type. Current types:
	//   - Synced: True after a successful external Alertmanager sync,
	//     False after a failed attempt, Unknown until the first attempt
	//     has run.
	// Future state dimensions land here as additional condition types.
	Conditions []AlertingConfigCondition `json:"conditions,omitempty"`
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
type AlertingConfigV0alpha1StatusAlertmanagerExternalSync struct {
	// datasourceUid is the UID actually used on the last sync
	// attempt. May differ from
	// spec.alertmanager.externalSync.datasourceUid immediately
	// after a spec change, until the next tick. When
	// `origin = "ini"`, this is the grafana.ini override value.
	DatasourceUid *string `json:"datasourceUid,omitempty"`
	// origin records which source supplied datasourceUid on the
	// last run:
	//   - "api": value from spec.alertmanager.externalSync.datasourceUid
	//     (set by an admin via the k8s API).
	//   - "ini": grafana.ini override (`[unified_alerting]
	//     external_alertmanager_uid`), set by the server operator.
	//     Wins over api when both are present.
	Origin *AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin `json:"origin,omitempty"`
}

// NewAlertingConfigV0alpha1StatusAlertmanagerExternalSync creates a new AlertingConfigV0alpha1StatusAlertmanagerExternalSync object.
func NewAlertingConfigV0alpha1StatusAlertmanagerExternalSync() *AlertingConfigV0alpha1StatusAlertmanagerExternalSync {
	return &AlertingConfigV0alpha1StatusAlertmanagerExternalSync{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigV0alpha1StatusAlertmanagerExternalSync.
func (AlertingConfigV0alpha1StatusAlertmanagerExternalSync) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigV0alpha1StatusAlertmanagerExternalSync"
}

// +k8s:openapi-gen=true
type AlertingConfigV0alpha1StatusAlertmanager struct {
	// externalSync carries the observation context for the external
	// Alertmanager configuration sync worker. Conditions about this
	// concern live at .status.conditions[type=Synced], not here.
	ExternalSync *AlertingConfigV0alpha1StatusAlertmanagerExternalSync `json:"externalSync,omitempty"`
}

// NewAlertingConfigV0alpha1StatusAlertmanager creates a new AlertingConfigV0alpha1StatusAlertmanager object.
func NewAlertingConfigV0alpha1StatusAlertmanager() *AlertingConfigV0alpha1StatusAlertmanager {
	return &AlertingConfigV0alpha1StatusAlertmanager{}
}

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigV0alpha1StatusAlertmanager.
func (AlertingConfigV0alpha1StatusAlertmanager) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigV0alpha1StatusAlertmanager"
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

// +k8s:openapi-gen=true
type AlertingConfigConditionStatus string

const (
	AlertingConfigConditionStatusTrue    AlertingConfigConditionStatus = "True"
	AlertingConfigConditionStatusFalse   AlertingConfigConditionStatus = "False"
	AlertingConfigConditionStatusUnknown AlertingConfigConditionStatus = "Unknown"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigConditionStatus.
func (AlertingConfigConditionStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigConditionStatus"
}

// +k8s:openapi-gen=true
type AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin string

const (
	AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOriginApi AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin = "api"
	AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOriginIni AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin.
func (AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.AlertingConfigV0alpha1StatusAlertmanagerExternalSyncOrigin"
}
