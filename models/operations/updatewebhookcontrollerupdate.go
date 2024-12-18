// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdateWebhookControllerUpdateRequest struct {
	// Id of endpoint to update
	EndpointID        string                       `pathParam:"style=simple,explode=false,name=endpointId"`
	UpdateWebhookBody components.UpdateWebhookBody `request:"mediaType=application/json"`
}

func (o *UpdateWebhookControllerUpdateRequest) GetEndpointID() string {
	if o == nil {
		return ""
	}
	return o.EndpointID
}

func (o *UpdateWebhookControllerUpdateRequest) GetUpdateWebhookBody() components.UpdateWebhookBody {
	if o == nil {
		return components.UpdateWebhookBody{}
	}
	return o.UpdateWebhookBody
}

type UpdateWebhookControllerUpdateResponse struct {
	HTTPMeta              components.HTTPMetadata `json:"-"`
	UpdateWebhookResponse *components.UpdateWebhookResponse
}

func (o *UpdateWebhookControllerUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateWebhookControllerUpdateResponse) GetUpdateWebhookResponse() *components.UpdateWebhookResponse {
	if o == nil {
		return nil
	}
	return o.UpdateWebhookResponse
}
