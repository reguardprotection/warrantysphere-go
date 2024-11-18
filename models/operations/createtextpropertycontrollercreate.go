// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CreateTextPropertyControllerCreateResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	CreatePropertyCommandResponse *components.CreatePropertyCommandResponse
}

func (o *CreateTextPropertyControllerCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTextPropertyControllerCreateResponse) GetCreatePropertyCommandResponse() *components.CreatePropertyCommandResponse {
	if o == nil {
		return nil
	}
	return o.CreatePropertyCommandResponse
}