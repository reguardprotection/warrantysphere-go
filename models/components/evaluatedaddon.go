// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EvaluatedAddon struct {
	// The unique identifier of the addon
	AddonID string `json:"addonId"`
	// The addon details
	Addon     AddonDto            `json:"addon"`
	Coverages []EvaluatedCoverage `json:"coverages"`
	// The price before taxes of the addon and its coverages
	Subtotal float64 `json:"subtotal"`
	// The taxes of the addon and its coverages
	Taxes []EvaluatedTax `json:"taxes"`
	// The price including taxes of the addon and its coverages
	Total float64 `json:"total"`
}

func (o *EvaluatedAddon) GetAddonID() string {
	if o == nil {
		return ""
	}
	return o.AddonID
}

func (o *EvaluatedAddon) GetAddon() AddonDto {
	if o == nil {
		return AddonDto{}
	}
	return o.Addon
}

func (o *EvaluatedAddon) GetCoverages() []EvaluatedCoverage {
	if o == nil {
		return []EvaluatedCoverage{}
	}
	return o.Coverages
}

func (o *EvaluatedAddon) GetSubtotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.Subtotal
}

func (o *EvaluatedAddon) GetTaxes() []EvaluatedTax {
	if o == nil {
		return []EvaluatedTax{}
	}
	return o.Taxes
}

func (o *EvaluatedAddon) GetTotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.Total
}
