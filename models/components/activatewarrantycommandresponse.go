// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ActivateWarrantyCommandResponse struct {
	// Unique identifier of the warranty.
	ID string `json:"id"`
}

func (o *ActivateWarrantyCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
