// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type OpenClaimCommandResponse struct {
	// Unique identifier of the claim.
	ID string `json:"id"`
}

func (o *OpenClaimCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
