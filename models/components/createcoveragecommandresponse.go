// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateCoverageCommandResponse struct {
	// Unique identifier of the coverage.
	ID string `json:"id"`
	// Unique identifier of the policy the coverage belongs to
	PolicyID string `json:"policyId"`
}

func (o *CreateCoverageCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateCoverageCommandResponse) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}
