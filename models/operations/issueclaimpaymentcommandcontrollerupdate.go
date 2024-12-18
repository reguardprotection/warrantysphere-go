// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type IssueClaimPaymentCommandControllerUpdateRequest struct {
	// Unique ID of the payment
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
	// Unique identifier of the claim associated with the claim payment.
	ClaimID string `pathParam:"style=simple,explode=false,name=claimId"`
}

func (o *IssueClaimPaymentCommandControllerUpdateRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

func (o *IssueClaimPaymentCommandControllerUpdateRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

type IssueClaimPaymentCommandControllerUpdateResponse struct {
	HTTPMeta                         components.HTTPMetadata `json:"-"`
	IssueClaimPaymentCommandResponse *components.IssueClaimPaymentCommandResponse
}

func (o *IssueClaimPaymentCommandControllerUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IssueClaimPaymentCommandControllerUpdateResponse) GetIssueClaimPaymentCommandResponse() *components.IssueClaimPaymentCommandResponse {
	if o == nil {
		return nil
	}
	return o.IssueClaimPaymentCommandResponse
}
