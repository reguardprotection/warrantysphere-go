// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CreateClaimPaymentCommandControllerCreateRequest struct {
	// Unique identifier of the claim associated with the claim payment.
	ClaimID                          string                                      `pathParam:"style=simple,explode=false,name=claimId"`
	CreateClaimPaymentCommandRequest components.CreateClaimPaymentCommandRequest `request:"mediaType=application/json"`
}

func (o *CreateClaimPaymentCommandControllerCreateRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *CreateClaimPaymentCommandControllerCreateRequest) GetCreateClaimPaymentCommandRequest() components.CreateClaimPaymentCommandRequest {
	if o == nil {
		return components.CreateClaimPaymentCommandRequest{}
	}
	return o.CreateClaimPaymentCommandRequest
}

type CreateClaimPaymentCommandControllerCreateResponse struct {
	HTTPMeta                          components.HTTPMetadata `json:"-"`
	CreateClaimPaymentCommandResponse *components.CreateClaimPaymentCommandResponse
}

func (o *CreateClaimPaymentCommandControllerCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateClaimPaymentCommandControllerCreateResponse) GetCreateClaimPaymentCommandResponse() *components.CreateClaimPaymentCommandResponse {
	if o == nil {
		return nil
	}
	return o.CreateClaimPaymentCommandResponse
}