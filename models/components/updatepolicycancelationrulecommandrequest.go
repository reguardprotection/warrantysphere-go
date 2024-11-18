// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdatePolicyCancelationRuleCommandRequest struct {
	AdminFee     *float64 `json:"adminFee,omitempty"`
	DeductClaims bool     `json:"deductClaims"`
	RuleText     *string  `json:"ruleText,omitempty"`
}

func (o *UpdatePolicyCancelationRuleCommandRequest) GetAdminFee() *float64 {
	if o == nil {
		return nil
	}
	return o.AdminFee
}

func (o *UpdatePolicyCancelationRuleCommandRequest) GetDeductClaims() bool {
	if o == nil {
		return false
	}
	return o.DeductClaims
}

func (o *UpdatePolicyCancelationRuleCommandRequest) GetRuleText() *string {
	if o == nil {
		return nil
	}
	return o.RuleText
}