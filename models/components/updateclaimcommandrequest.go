// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateClaimCommandRequest struct {
	// Reason as to why the customer opened the claim.
	Complaint string `json:"complaint"`
}

func (o *UpdateClaimCommandRequest) GetComplaint() string {
	if o == nil {
		return ""
	}
	return o.Complaint
}
