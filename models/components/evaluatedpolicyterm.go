// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EvaluatedPolicyTerm struct {
	// The policy term
	Term          PolicyTerm          `json:"term"`
	BaseCoverages []EvaluatedCoverage `json:"baseCoverages"`
	// The price before taxes of the base plan and coverages
	BaseSubtotal float64 `json:"baseSubtotal"`
	// The taxes of the base plan and its coverages
	BaseTaxes []EvaluatedTax `json:"baseTaxes"`
	// The price including taxes of the base plan and coverages
	BaseTotal float64          `json:"baseTotal"`
	Addons    []EvaluatedAddon `json:"addons"`
}

func (o *EvaluatedPolicyTerm) GetTerm() PolicyTerm {
	if o == nil {
		return PolicyTerm{}
	}
	return o.Term
}

func (o *EvaluatedPolicyTerm) GetBaseCoverages() []EvaluatedCoverage {
	if o == nil {
		return []EvaluatedCoverage{}
	}
	return o.BaseCoverages
}

func (o *EvaluatedPolicyTerm) GetBaseSubtotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.BaseSubtotal
}

func (o *EvaluatedPolicyTerm) GetBaseTaxes() []EvaluatedTax {
	if o == nil {
		return []EvaluatedTax{}
	}
	return o.BaseTaxes
}

func (o *EvaluatedPolicyTerm) GetBaseTotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.BaseTotal
}

func (o *EvaluatedPolicyTerm) GetAddons() []EvaluatedAddon {
	if o == nil {
		return []EvaluatedAddon{}
	}
	return o.Addons
}