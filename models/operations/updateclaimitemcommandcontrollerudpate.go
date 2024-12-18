// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdateClaimItemCommandControllerUdpateRequest struct {
	// Unique identifier of the claim item.
	ItemID string `pathParam:"style=simple,explode=false,name=itemId"`
	// Unique identifier of the claim associated with the claim item.
	ClaimID                       string                                   `pathParam:"style=simple,explode=false,name=claimId"`
	UpdateClaimItemCommandRequest components.UpdateClaimItemCommandRequest `request:"mediaType=application/json"`
}

func (o *UpdateClaimItemCommandControllerUdpateRequest) GetItemID() string {
	if o == nil {
		return ""
	}
	return o.ItemID
}

func (o *UpdateClaimItemCommandControllerUdpateRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *UpdateClaimItemCommandControllerUdpateRequest) GetUpdateClaimItemCommandRequest() components.UpdateClaimItemCommandRequest {
	if o == nil {
		return components.UpdateClaimItemCommandRequest{}
	}
	return o.UpdateClaimItemCommandRequest
}

type UpdateClaimItemCommandControllerUdpateResponse struct {
	HTTPMeta                       components.HTTPMetadata `json:"-"`
	UpdateClaimItemCommandResponse *components.UpdateClaimItemCommandResponse
}

func (o *UpdateClaimItemCommandControllerUdpateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateClaimItemCommandControllerUdpateResponse) GetUpdateClaimItemCommandResponse() *components.UpdateClaimItemCommandResponse {
	if o == nil {
		return nil
	}
	return o.UpdateClaimItemCommandResponse
}
