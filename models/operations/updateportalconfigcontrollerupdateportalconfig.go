// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdatePortalConfigControllerUpdatePortalConfigResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	UpdatePortalConfigResponse *components.UpdatePortalConfigResponse
}

func (o *UpdatePortalConfigControllerUpdatePortalConfigResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePortalConfigControllerUpdatePortalConfigResponse) GetUpdatePortalConfigResponse() *components.UpdatePortalConfigResponse {
	if o == nil {
		return nil
	}
	return o.UpdatePortalConfigResponse
}