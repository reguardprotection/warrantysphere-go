// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DeletePolicyCommandResponse struct {
	// Unique identifier of the policy.
	ID string `json:"id"`
}

func (o *DeletePolicyCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
