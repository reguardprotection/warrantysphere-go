// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type VoidTransactionResponse struct {
	// Unique identifier of the transaction
	ID string `json:"id"`
}

func (o *VoidTransactionResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
