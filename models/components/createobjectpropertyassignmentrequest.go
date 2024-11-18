// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

type CreateObjectPropertyAssignmentRequest struct {
	// The unique id of the property being assigned to the object
	PropertyID string `json:"propertyId"`
	// Whether or not this property is required
	Required bool `json:"required"`
	// The order that this property appears in the object
	Order float64 `json:"order"`
	// Whether or not this property is visible when showing the property values in the UI
	Hidden *bool `default:"false" json:"hidden"`
}

func (c CreateObjectPropertyAssignmentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateObjectPropertyAssignmentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateObjectPropertyAssignmentRequest) GetPropertyID() string {
	if o == nil {
		return ""
	}
	return o.PropertyID
}

func (o *CreateObjectPropertyAssignmentRequest) GetRequired() bool {
	if o == nil {
		return false
	}
	return o.Required
}

func (o *CreateObjectPropertyAssignmentRequest) GetOrder() float64 {
	if o == nil {
		return 0.0
	}
	return o.Order
}

func (o *CreateObjectPropertyAssignmentRequest) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}