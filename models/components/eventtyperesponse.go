// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

type EventTypeResponse struct {
	// Description of event type
	Description *string `default:"null" json:"description"`
	// Name of event type
	Name *string `default:"null" json:"name"`
}

func (e EventTypeResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventTypeResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EventTypeResponse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *EventTypeResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
