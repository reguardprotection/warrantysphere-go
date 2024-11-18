// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

type UpdateWarrantyCommandRequest struct {
	// Date when the warranty's term starts.
	TermStart     *time.Time `json:"termStart,omitempty"`
	DistributorID *string    `json:"distributorId,omitempty"`
}

func (u UpdateWarrantyCommandRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateWarrantyCommandRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateWarrantyCommandRequest) GetTermStart() *time.Time {
	if o == nil {
		return nil
	}
	return o.TermStart
}

func (o *UpdateWarrantyCommandRequest) GetDistributorID() *string {
	if o == nil {
		return nil
	}
	return o.DistributorID
}