// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdatePolicyCancelationRuleCommandResponse struct {
	Country string `json:"country"`
	State   string `json:"state"`
}

func (o *UpdatePolicyCancelationRuleCommandResponse) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *UpdatePolicyCancelationRuleCommandResponse) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}