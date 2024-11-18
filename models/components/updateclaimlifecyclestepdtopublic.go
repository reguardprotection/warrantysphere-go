// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

// UpdateClaimLifecycleStepDtoPUBLICStatus - The default status of the step seen by the customer when a claim is opened.
type UpdateClaimLifecycleStepDtoPUBLICStatus string

const (
	UpdateClaimLifecycleStepDtoPUBLICStatusNotStartedYet   UpdateClaimLifecycleStepDtoPUBLICStatus = "NOT_STARTED_YET"
	UpdateClaimLifecycleStepDtoPUBLICStatusDone            UpdateClaimLifecycleStepDtoPUBLICStatus = "DONE"
	UpdateClaimLifecycleStepDtoPUBLICStatusInProgress      UpdateClaimLifecycleStepDtoPUBLICStatus = "IN_PROGRESS"
	UpdateClaimLifecycleStepDtoPUBLICStatusPendingCustomer UpdateClaimLifecycleStepDtoPUBLICStatus = "PENDING_CUSTOMER"
	UpdateClaimLifecycleStepDtoPUBLICStatusUnsuccessful    UpdateClaimLifecycleStepDtoPUBLICStatus = "UNSUCCESSFUL"
	UpdateClaimLifecycleStepDtoPUBLICStatusSkipped         UpdateClaimLifecycleStepDtoPUBLICStatus = "SKIPPED"
)

func (e UpdateClaimLifecycleStepDtoPUBLICStatus) ToPointer() *UpdateClaimLifecycleStepDtoPUBLICStatus {
	return &e
}
func (e *UpdateClaimLifecycleStepDtoPUBLICStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NOT_STARTED_YET":
		fallthrough
	case "DONE":
		fallthrough
	case "IN_PROGRESS":
		fallthrough
	case "PENDING_CUSTOMER":
		fallthrough
	case "UNSUCCESSFUL":
		fallthrough
	case "SKIPPED":
		*e = UpdateClaimLifecycleStepDtoPUBLICStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateClaimLifecycleStepDtoPUBLICStatus: %v", v)
	}
}

type UpdateClaimLifecycleStepDtoPUBLIC struct {
	// The icon used for the claim lifecycle step.
	Icon string `json:"icon"`
	// The title of the claim lifecycle step.
	Title string `json:"title"`
	// The description of the claim lifecycle step.
	Description *string `json:"description,omitempty"`
	// The default status of the step seen by the customer when a claim is opened.
	Status *UpdateClaimLifecycleStepDtoPUBLICStatus `default:"NOT_STARTED_YET" json:"status"`
	// Wether or not this step is optional for the customer.
	Optional *bool `default:"false" json:"optional"`
	// Wether or not this step is visible to the customer.
	Visible *bool `default:"true" json:"visible"`
}

func (u UpdateClaimLifecycleStepDtoPUBLIC) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateClaimLifecycleStepDtoPUBLIC) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateClaimLifecycleStepDtoPUBLIC) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *UpdateClaimLifecycleStepDtoPUBLIC) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *UpdateClaimLifecycleStepDtoPUBLIC) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateClaimLifecycleStepDtoPUBLIC) GetStatus() *UpdateClaimLifecycleStepDtoPUBLICStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UpdateClaimLifecycleStepDtoPUBLIC) GetOptional() *bool {
	if o == nil {
		return nil
	}
	return o.Optional
}

func (o *UpdateClaimLifecycleStepDtoPUBLIC) GetVisible() *bool {
	if o == nil {
		return nil
	}
	return o.Visible
}