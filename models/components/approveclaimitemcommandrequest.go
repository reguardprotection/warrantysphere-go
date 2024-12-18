// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ApproveClaimItemCommandRequest struct {
	// Reason provided as to why the claim item was approved.
	Reason string `json:"reason"`
}

func (o *ApproveClaimItemCommandRequest) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}
