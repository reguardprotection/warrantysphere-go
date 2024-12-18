// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type DisableAPIKeyCommandControllerDisableRequest struct {
	// The id of the key to disable
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DisableAPIKeyCommandControllerDisableRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DisableAPIKeyCommandControllerDisableResponse struct {
	HTTPMeta              components.HTTPMetadata `json:"-"`
	DisableAPIKeyResponse *components.DisableAPIKeyResponse
}

func (o *DisableAPIKeyCommandControllerDisableResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DisableAPIKeyCommandControllerDisableResponse) GetDisableAPIKeyResponse() *components.DisableAPIKeyResponse {
	if o == nil {
		return nil
	}
	return o.DisableAPIKeyResponse
}
