// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

type UpdatePolicyCommandRequest struct {
	// Whether or not the command should create a new draft policy if required.
	CanCreateNewVersion bool `json:"canCreateNewVersion"`
	// Title of the policy.
	Title *string `json:"title,omitempty"`
	// Friendly title of the policy.
	FriendlyTitle *string `json:"friendlyTitle,omitempty"`
	// Icon of the policy.
	Icon *string `json:"icon,omitempty"`
	// Tagline of the policy.
	Tagline *string `json:"tagline,omitempty"`
	// Description of the policy.
	Description *string `json:"description,omitempty"`
	// The available terms for the policy.
	Terms []UpdatePolicyTermDto `json:"terms,omitempty"`
	// Duration (in days) until a claim can be made for the policy.
	ClaimWaitingPeriod *float64 `json:"claimWaitingPeriod,omitempty"`
	// Whether or not provisioned warranties require manual activation from the customer.
	RequiresWarrantyActivation *bool `default:"false" json:"requiresWarrantyActivation"`
}

func (u UpdatePolicyCommandRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdatePolicyCommandRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdatePolicyCommandRequest) GetCanCreateNewVersion() bool {
	if o == nil {
		return false
	}
	return o.CanCreateNewVersion
}

func (o *UpdatePolicyCommandRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UpdatePolicyCommandRequest) GetFriendlyTitle() *string {
	if o == nil {
		return nil
	}
	return o.FriendlyTitle
}

func (o *UpdatePolicyCommandRequest) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *UpdatePolicyCommandRequest) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *UpdatePolicyCommandRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdatePolicyCommandRequest) GetTerms() []UpdatePolicyTermDto {
	if o == nil {
		return nil
	}
	return o.Terms
}

func (o *UpdatePolicyCommandRequest) GetClaimWaitingPeriod() *float64 {
	if o == nil {
		return nil
	}
	return o.ClaimWaitingPeriod
}

func (o *UpdatePolicyCommandRequest) GetRequiresWarrantyActivation() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresWarrantyActivation
}
