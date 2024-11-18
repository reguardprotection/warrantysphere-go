// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdateClaimCommandControllerUpdateRequest struct {
	// Unique identifier of the claim associated with the claim item.
	ClaimID                   string                               `pathParam:"style=simple,explode=false,name=claimId"`
	UpdateClaimCommandRequest components.UpdateClaimCommandRequest `request:"mediaType=application/json"`
}

func (o *UpdateClaimCommandControllerUpdateRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *UpdateClaimCommandControllerUpdateRequest) GetUpdateClaimCommandRequest() components.UpdateClaimCommandRequest {
	if o == nil {
		return components.UpdateClaimCommandRequest{}
	}
	return o.UpdateClaimCommandRequest
}

type UpdateClaimCommandControllerUpdateResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	UpdateClaimCommandResponse *components.UpdateClaimCommandResponse
}

func (o *UpdateClaimCommandControllerUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateClaimCommandControllerUpdateResponse) GetUpdateClaimCommandResponse() *components.UpdateClaimCommandResponse {
	if o == nil {
		return nil
	}
	return o.UpdateClaimCommandResponse
}
