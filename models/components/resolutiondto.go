// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ResolutionDto struct {
	// Title of the cure associated with the claim item.
	Title *string `json:"title"`
	// Description of the cure aassociated with the claim item.
	Description *string `json:"description"`
	// Actual amount (in cents) of the cure.
	ActualCost *float64 `json:"actualCost"`
}

func (o *ResolutionDto) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ResolutionDto) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ResolutionDto) GetActualCost() *float64 {
	if o == nil {
		return nil
	}
	return o.ActualCost
}
