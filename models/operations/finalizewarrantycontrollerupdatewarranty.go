// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type FinalizeWarrantyControllerUpdateWarrantyRequest struct {
	// Unique identifier of the warranty.
	WarrantyID                     string                                    `pathParam:"style=simple,explode=false,name=warrantyId"`
	FinalizeWarrantyCommandRequest components.FinalizeWarrantyCommandRequest `request:"mediaType=application/json"`
}

func (o *FinalizeWarrantyControllerUpdateWarrantyRequest) GetWarrantyID() string {
	if o == nil {
		return ""
	}
	return o.WarrantyID
}

func (o *FinalizeWarrantyControllerUpdateWarrantyRequest) GetFinalizeWarrantyCommandRequest() components.FinalizeWarrantyCommandRequest {
	if o == nil {
		return components.FinalizeWarrantyCommandRequest{}
	}
	return o.FinalizeWarrantyCommandRequest
}

type FinalizeWarrantyControllerUpdateWarrantyResponse struct {
	HTTPMeta                        components.HTTPMetadata `json:"-"`
	FinalizeWarrantyCommandResponse *components.FinalizeWarrantyCommandResponse
}

func (o *FinalizeWarrantyControllerUpdateWarrantyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FinalizeWarrantyControllerUpdateWarrantyResponse) GetFinalizeWarrantyCommandResponse() *components.FinalizeWarrantyCommandResponse {
	if o == nil {
		return nil
	}
	return o.FinalizeWarrantyCommandResponse
}
