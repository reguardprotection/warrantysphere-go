// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CancelClaimCommandControllerCancelRequest struct {
	// Unique identifier of the claim
	ClaimID                       string                                   `pathParam:"style=simple,explode=false,name=claimId"`
	CancelClaimCommandRequestBody components.CancelClaimCommandRequestBody `request:"mediaType=application/json"`
}

func (o *CancelClaimCommandControllerCancelRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *CancelClaimCommandControllerCancelRequest) GetCancelClaimCommandRequestBody() components.CancelClaimCommandRequestBody {
	if o == nil {
		return components.CancelClaimCommandRequestBody{}
	}
	return o.CancelClaimCommandRequestBody
}

type CancelClaimCommandControllerCancelResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	CancelClaimCommandResponse *components.CancelClaimCommandResponse
}

func (o *CancelClaimCommandControllerCancelResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CancelClaimCommandControllerCancelResponse) GetCancelClaimCommandResponse() *components.CancelClaimCommandResponse {
	if o == nil {
		return nil
	}
	return o.CancelClaimCommandResponse
}
