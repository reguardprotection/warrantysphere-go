// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdatePolicyPlanCommandResponse struct {
	// Unique identifier of the plan
	ID string `json:"id"`
	// Unique identifier of the policy the plan belongs to
	PolicyID string `json:"policyId"`
}

func (o *UpdatePolicyPlanCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdatePolicyPlanCommandResponse) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}
