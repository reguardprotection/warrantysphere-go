// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CreateNumberPropertyControllerCreateResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	CreatePropertyCommandResponse *components.CreatePropertyCommandResponse
}

func (o *CreateNumberPropertyControllerCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateNumberPropertyControllerCreateResponse) GetCreatePropertyCommandResponse() *components.CreatePropertyCommandResponse {
	if o == nil {
		return nil
	}
	return o.CreatePropertyCommandResponse
}
