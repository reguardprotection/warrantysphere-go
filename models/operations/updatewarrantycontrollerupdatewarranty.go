// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdateWarrantyControllerUpdateWarrantyRequest struct {
	// Unique identifier of the warranty.
	WarrantyID                   string                                  `pathParam:"style=simple,explode=false,name=warrantyId"`
	UpdateWarrantyCommandRequest components.UpdateWarrantyCommandRequest `request:"mediaType=application/json"`
}

func (o *UpdateWarrantyControllerUpdateWarrantyRequest) GetWarrantyID() string {
	if o == nil {
		return ""
	}
	return o.WarrantyID
}

func (o *UpdateWarrantyControllerUpdateWarrantyRequest) GetUpdateWarrantyCommandRequest() components.UpdateWarrantyCommandRequest {
	if o == nil {
		return components.UpdateWarrantyCommandRequest{}
	}
	return o.UpdateWarrantyCommandRequest
}

type UpdateWarrantyControllerUpdateWarrantyResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	UpdateWarrantyCommandResponse *components.UpdateWarrantyCommandResponse
}

func (o *UpdateWarrantyControllerUpdateWarrantyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateWarrantyControllerUpdateWarrantyResponse) GetUpdateWarrantyCommandResponse() *components.UpdateWarrantyCommandResponse {
	if o == nil {
		return nil
	}
	return o.UpdateWarrantyCommandResponse
}
