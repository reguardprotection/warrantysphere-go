// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

type RetrievePropertyQueryResponseType struct {
}

type Config struct {
}

type RetrievePropertyQueryResponse struct {
	// Unique identifier of the property.
	ID       string                            `json:"id"`
	Type     RetrievePropertyQueryResponseType `json:"type"`
	Name     string                            `json:"name"`
	ValueKey string                            `json:"valueKey"`
	FormID   string                            `json:"formId"`
	Config   Config                            `json:"config"`
	Created  time.Time                         `json:"created"`
	Modified time.Time                         `json:"modified"`
}

func (r RetrievePropertyQueryResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrievePropertyQueryResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrievePropertyQueryResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrievePropertyQueryResponse) GetType() RetrievePropertyQueryResponseType {
	if o == nil {
		return RetrievePropertyQueryResponseType{}
	}
	return o.Type
}

func (o *RetrievePropertyQueryResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RetrievePropertyQueryResponse) GetValueKey() string {
	if o == nil {
		return ""
	}
	return o.ValueKey
}

func (o *RetrievePropertyQueryResponse) GetFormID() string {
	if o == nil {
		return ""
	}
	return o.FormID
}

func (o *RetrievePropertyQueryResponse) GetConfig() Config {
	if o == nil {
		return Config{}
	}
	return o.Config
}

func (o *RetrievePropertyQueryResponse) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *RetrievePropertyQueryResponse) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}
