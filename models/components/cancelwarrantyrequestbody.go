// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

type CancelWarrantyRequestBody struct {
	// Effective date of the cancellation.
	CancellationDate *time.Time `json:"cancellationDate,omitempty"`
}

func (c CancelWarrantyRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CancelWarrantyRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CancelWarrantyRequestBody) GetCancellationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.CancellationDate
}
