// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type GetStatusResponse struct {
	Body map[string]interface{} `json:"body"`
}

// NewGetStatusResponse creates a new GetStatusResponse object.
func NewGetStatusResponse() *GetStatusResponse {
	return &GetStatusResponse{
		Body: map[string]interface{}{},
	}
}

// OpenAPIModelName returns the OpenAPI model name for GetStatusResponse.
func (GetStatusResponse) OpenAPIModelName() string {
	return "com.github.grafana.grafana.apps.alerting.admin.pkg.apis.alertingadmin.v0alpha1.GetStatusResponse"
}
