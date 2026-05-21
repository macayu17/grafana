// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncstatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State ExternalAlertmanagerSyncStatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewExternalAlertmanagerSyncstatusOperatorState creates a new ExternalAlertmanagerSyncstatusOperatorState object.
func NewExternalAlertmanagerSyncstatusOperatorState() *ExternalAlertmanagerSyncstatusOperatorState {
	return &ExternalAlertmanagerSyncstatusOperatorState{}
}

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncstatusOperatorState.
func (ExternalAlertmanagerSyncstatusOperatorState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncstatusOperatorState"
}

// Condition mirrors metav1.Condition. Declared inline because the app-sdk
// codegen in this repo does not yet have a built-in path for referencing the
// k8s metav1.Condition type from CUE. Field semantics are k8s-standard:
//   - status flips between True/False/Unknown.
//   - lastTransitionTime advances only when status flips.
//   - reason is a PascalCase machine-readable enum (e.g. "SyncSucceeded",
//     "MimirFetchFailed"); see SyncReason in the syncer.
//   - message is human-readable detail.
//   - observedGeneration records the spec.generation this condition evaluation
//     reflects, when applicable.
//
// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncCondition struct {
	Type   string                                  `json:"type"`
	Status ExternalAlertmanagerSyncConditionStatus `json:"status"`
	// RFC3339
	LastTransitionTime string  `json:"lastTransitionTime"`
	Reason             string  `json:"reason"`
	Message            *string `json:"message,omitempty"`
	ObservedGeneration *int64  `json:"observedGeneration,omitempty"`
}

// NewExternalAlertmanagerSyncCondition creates a new ExternalAlertmanagerSyncCondition object.
func NewExternalAlertmanagerSyncCondition() *ExternalAlertmanagerSyncCondition {
	return &ExternalAlertmanagerSyncCondition{}
}

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncCondition.
func (ExternalAlertmanagerSyncCondition) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncCondition"
}

// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncStatus struct {
	// observedGeneration is the spec.generation last evaluated by the syncer.
	// Always 0 today since the spec is empty; carried for forward compatibility
	// with the conditions pattern.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// UID actually used on the last sync attempt. May differ from
	// `AlertingConfig.spec.alertmanager.externalSync.datasourceUid` immediately after
	// a spec change, until the next tick. When `origin = "ini"`, this is the
	// grafana.ini override value.
	DatasourceUid *string `json:"datasourceUid,omitempty"`
	// Which source supplied datasourceUid on the last run:
	//   - "api": value from AlertingConfig.spec.alertmanager.externalSync.datasourceUid
	//     (set by an admin via the k8s API).
	//   - "ini": grafana.ini override (`[unified_alerting]
	//     external_alertmanager_uid`), set by the server operator. Wins over
	//     api when both are present.
	// This field lets clients see when an ini override is in effect without
	// having to know the precedence rule.
	Origin *ExternalAlertmanagerSyncStatusOrigin `json:"origin,omitempty"`
	// Unix epoch seconds of the most recent successful sync. Preserved across
	// failure streaks — answers "even though it's broken now, when did it
	// last work?". Not derivable from `Synced.lastTransitionTime` (which
	// marks when the current state was entered, not when success last held).
	// Omitted when no sync has ever succeeded for this org.
	LastSuccessAt *int64 `json:"lastSuccessAt,omitempty"`
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]ExternalAlertmanagerSyncstatusOperatorState `json:"operatorStates,omitempty"`
	// Standard k8s-style condition list. v1 carries one type:
	//   - Synced: True after a successful sync, False after a failed attempt,
	//     Unknown until the first attempt has run.
	// Future state dimensions land here as additional condition types.
	Conditions []ExternalAlertmanagerSyncCondition `json:"conditions,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewExternalAlertmanagerSyncStatus creates a new ExternalAlertmanagerSyncStatus object.
func NewExternalAlertmanagerSyncStatus() *ExternalAlertmanagerSyncStatus {
	return &ExternalAlertmanagerSyncStatus{}
}

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncStatus.
func (ExternalAlertmanagerSyncStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncStatus"
}

// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncStatusOperatorStateState string

const (
	ExternalAlertmanagerSyncStatusOperatorStateStateSuccess    ExternalAlertmanagerSyncStatusOperatorStateState = "success"
	ExternalAlertmanagerSyncStatusOperatorStateStateInProgress ExternalAlertmanagerSyncStatusOperatorStateState = "in_progress"
	ExternalAlertmanagerSyncStatusOperatorStateStateFailed     ExternalAlertmanagerSyncStatusOperatorStateState = "failed"
)

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncStatusOperatorStateState.
func (ExternalAlertmanagerSyncStatusOperatorStateState) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncStatusOperatorStateState"
}

// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncConditionStatus string

const (
	ExternalAlertmanagerSyncConditionStatusTrue    ExternalAlertmanagerSyncConditionStatus = "True"
	ExternalAlertmanagerSyncConditionStatusFalse   ExternalAlertmanagerSyncConditionStatus = "False"
	ExternalAlertmanagerSyncConditionStatusUnknown ExternalAlertmanagerSyncConditionStatus = "Unknown"
)

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncConditionStatus.
func (ExternalAlertmanagerSyncConditionStatus) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncConditionStatus"
}

// +k8s:openapi-gen=true
type ExternalAlertmanagerSyncStatusOrigin string

const (
	ExternalAlertmanagerSyncStatusOriginApi ExternalAlertmanagerSyncStatusOrigin = "api"
	ExternalAlertmanagerSyncStatusOriginIni ExternalAlertmanagerSyncStatusOrigin = "ini"
)

// OpenAPIModelName returns the OpenAPI model name for ExternalAlertmanagerSyncStatusOrigin.
func (ExternalAlertmanagerSyncStatusOrigin) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.ExternalAlertmanagerSyncStatusOrigin"
}
