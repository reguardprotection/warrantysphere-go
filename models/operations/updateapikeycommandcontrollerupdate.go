// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdateAPIKeyCommandControllerUpdateRequest struct {
	// The id of the key to update
	ID                  string                         `pathParam:"style=simple,explode=false,name=id"`
	UpdateAPIKeyRequest components.UpdateAPIKeyRequest `request:"mediaType=application/json"`
}

func (o *UpdateAPIKeyCommandControllerUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAPIKeyCommandControllerUpdateRequest) GetUpdateAPIKeyRequest() components.UpdateAPIKeyRequest {
	if o == nil {
		return components.UpdateAPIKeyRequest{}
	}
	return o.UpdateAPIKeyRequest
}

type UpdateAPIKeyCommandControllerUpdateResponse struct {
	HTTPMeta             components.HTTPMetadata `json:"-"`
	UpdateAPIKeyResponse *components.UpdateAPIKeyResponse
}

func (o *UpdateAPIKeyCommandControllerUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAPIKeyCommandControllerUpdateResponse) GetUpdateAPIKeyResponse() *components.UpdateAPIKeyResponse {
	if o == nil {
		return nil
	}
	return o.UpdateAPIKeyResponse
}