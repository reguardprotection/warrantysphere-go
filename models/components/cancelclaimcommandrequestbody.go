// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CancelClaimCommandRequestBody struct {
	// Reason for cancellation of the claim
	CancellationReason string `json:"cancellationReason"`
}

func (o *CancelClaimCommandRequestBody) GetCancellationReason() string {
	if o == nil {
		return ""
	}
	return o.CancellationReason
}