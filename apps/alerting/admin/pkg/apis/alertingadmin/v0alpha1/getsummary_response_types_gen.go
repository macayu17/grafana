// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type GetSummaryResponse struct {
	Body map[string]interface{} `json:"body"`
}

// NewGetSummaryResponse creates a new GetSummaryResponse object.
func NewGetSummaryResponse() *GetSummaryResponse {
	return &GetSummaryResponse{
		Body: map[string]interface{}{},
	}
}

// OpenAPIModelName returns the OpenAPI model name for GetSummaryResponse.
func (GetSummaryResponse) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.GetSummaryResponse"
}
