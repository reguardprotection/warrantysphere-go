// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

// CreateTextPropertyRequestFormat - The format of the property
type CreateTextPropertyRequestFormat string

const (
	CreateTextPropertyRequestFormatText  CreateTextPropertyRequestFormat = "text"
	CreateTextPropertyRequestFormatBlock CreateTextPropertyRequestFormat = "block"
	CreateTextPropertyRequestFormatEmail CreateTextPropertyRequestFormat = "email"
	CreateTextPropertyRequestFormatPhone CreateTextPropertyRequestFormat = "phone"
)

func (e CreateTextPropertyRequestFormat) ToPointer() *CreateTextPropertyRequestFormat {
	return &e
}
func (e *CreateTextPropertyRequestFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "block":
		fallthrough
	case "email":
		fallthrough
	case "phone":
		*e = CreateTextPropertyRequestFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTextPropertyRequestFormat: %v", v)
	}
}

// CreateTextPropertyRequestConfig - The extra configuration of this property
type CreateTextPropertyRequestConfig struct {
	// The format of the property
	Format *CreateTextPropertyRequestFormat `json:"format"`
	// An array that represents what options the property value is allowed to be
	AllowedValues []string `json:"allowedValues"`
	// The minimum characters allowed
	MinLength *float64 `json:"minLength,omitempty"`
	// The maximum characters allowed
	MaxLength *float64 `json:"maxLength,omitempty"`
}

func (o *CreateTextPropertyRequestConfig) GetFormat() *CreateTextPropertyRequestFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *CreateTextPropertyRequestConfig) GetAllowedValues() []string {
	if o == nil {
		return nil
	}
	return o.AllowedValues
}

func (o *CreateTextPropertyRequestConfig) GetMinLength() *float64 {
	if o == nil {
		return nil
	}
	return o.MinLength
}

func (o *CreateTextPropertyRequestConfig) GetMaxLength() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxLength
}

type CreateTextPropertyRequest struct {
	// The visible name of the property
	Name string `json:"name"`
	// The key of this property. Only lowercase and no special characters or spaces
	ValueKey string `json:"valueKey"`
	// The formId
	FormID *string `default:"null" json:"formId"`
	// The extra configuration of this property
	Config CreateTextPropertyRequestConfig `json:"config"`
}

func (c CreateTextPropertyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTextPropertyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTextPropertyRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTextPropertyRequest) GetValueKey() string {
	if o == nil {
		return ""
	}
	return o.ValueKey
}

func (o *CreateTextPropertyRequest) GetFormID() *string {
	if o == nil {
		return nil
	}
	return o.FormID
}

func (o *CreateTextPropertyRequest) GetConfig() CreateTextPropertyRequestConfig {
	if o == nil {
		return CreateTextPropertyRequestConfig{}
	}
	return o.Config
}