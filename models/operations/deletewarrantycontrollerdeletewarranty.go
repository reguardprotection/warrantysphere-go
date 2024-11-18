// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type DeleteWarrantyControllerDeleteWarrantyRequest struct {
	// Unique identifier of the warranty.
	WarrantyID string `pathParam:"style=simple,explode=false,name=warrantyId"`
}

func (o *DeleteWarrantyControllerDeleteWarrantyRequest) GetWarrantyID() string {
	if o == nil {
		return ""
	}
	return o.WarrantyID
}

type DeleteWarrantyControllerDeleteWarrantyResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	DeleteWarrantyCommandResponse *components.DeleteWarrantyCommandResponse
}

func (o *DeleteWarrantyControllerDeleteWarrantyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteWarrantyControllerDeleteWarrantyResponse) GetDeleteWarrantyCommandResponse() *components.DeleteWarrantyCommandResponse {
	if o == nil {
		return nil
	}
	return o.DeleteWarrantyCommandResponse
}