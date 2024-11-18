// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateCoverageCommandResponse struct {
	// Unique identifier of the coverage.
	ID string `json:"id"`
	// Unique identifier of the policy the coverage belongs to
	PolicyID string `json:"policyId"`
}

func (o *UpdateCoverageCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateCoverageCommandResponse) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}