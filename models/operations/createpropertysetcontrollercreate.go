// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CreatePropertySetControllerCreateResponse struct {
	HTTPMeta                         components.HTTPMetadata `json:"-"`
	CreatePropertySetCommandResponse *components.CreatePropertySetCommandResponse
}

func (o *CreatePropertySetControllerCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatePropertySetControllerCreateResponse) GetCreatePropertySetCommandResponse() *components.CreatePropertySetCommandResponse {
	if o == nil {
		return nil
	}
	return o.CreatePropertySetCommandResponse
}