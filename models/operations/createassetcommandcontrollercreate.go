// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CreateAssetCommandControllerCreateResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	CreateAssetCommandResponse *components.CreateAssetCommandResponse
}

func (o *CreateAssetCommandControllerCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateAssetCommandControllerCreateResponse) GetCreateAssetCommandResponse() *components.CreateAssetCommandResponse {
	if o == nil {
		return nil
	}
	return o.CreateAssetCommandResponse
}
