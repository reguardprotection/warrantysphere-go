// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

type CreateDatePropertyRequest struct {
	// The visible name of the property
	Name string `json:"name"`
	// The key of this property. Only lowercase and no special characters or spaces
	ValueKey string `json:"valueKey"`
	// The formId
	FormID *string `default:"null" json:"formId"`
	// The extra configuration of this property
	Config DatePropertyConfig `json:"config"`
}

func (c CreateDatePropertyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateDatePropertyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateDatePropertyRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateDatePropertyRequest) GetValueKey() string {
	if o == nil {
		return ""
	}
	return o.ValueKey
}

func (o *CreateDatePropertyRequest) GetFormID() *string {
	if o == nil {
		return nil
	}
	return o.FormID
}

func (o *CreateDatePropertyRequest) GetConfig() DatePropertyConfig {
	if o == nil {
		return DatePropertyConfig{}
	}
	return o.Config
}
