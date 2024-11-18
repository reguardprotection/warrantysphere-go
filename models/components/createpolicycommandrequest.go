// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
)

type CreatePolicyCommandRequest struct {
	// Title of the policy.
	Title string `json:"title"`
	// Friendly title of the policy.
	FriendlyTitle *string `json:"friendlyTitle,omitempty"`
	// Icon of the policy.
	Icon string `json:"icon"`
	// Tagline of the policy.
	Tagline string `json:"tagline"`
	// Description of the policy.
	Description *string `json:"description,omitempty"`
	// Identifier for this policy which is also used to keep track of policy versions.
	PolicyNumber string `json:"policyNumber"`
	// Number of days after the warranty is activated before a claim can be opened.
	ClaimWaitingPeriod *float64 `json:"claimWaitingPeriod,omitempty"`
	// The unique identifier of the property set
	PropertySetID string `json:"propertySetId"`
	// The claim lifecycle steps used for this policy.
	ClaimLifecycleSteps []ClaimLifecycleStepDto `json:"claimLifecycleSteps,omitempty"`
	// Coverages linked to this policy.
	Coverages []CreateCoverageDto `json:"coverages,omitempty"`
	// 'Terms & Conditions' document linked to this policy.
	TermsAndConditions CreateDocumentDto `json:"termsAndConditions"`
	// 'Coverage Summary' document linked to this policy.
	CoverageSummary CreateDocumentDto `json:"coverageSummary"`
	// Terms available on the policy.
	Terms []CreatePolicyTermDto `json:"terms"`
	// Whether or not provisioned warranties require manual activation from the customer.
	RequiresWarrantyActivation *bool `default:"false" json:"requiresWarrantyActivation"`
}

func (c CreatePolicyCommandRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePolicyCommandRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePolicyCommandRequest) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *CreatePolicyCommandRequest) GetFriendlyTitle() *string {
	if o == nil {
		return nil
	}
	return o.FriendlyTitle
}

func (o *CreatePolicyCommandRequest) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *CreatePolicyCommandRequest) GetTagline() string {
	if o == nil {
		return ""
	}
	return o.Tagline
}

func (o *CreatePolicyCommandRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreatePolicyCommandRequest) GetPolicyNumber() string {
	if o == nil {
		return ""
	}
	return o.PolicyNumber
}

func (o *CreatePolicyCommandRequest) GetClaimWaitingPeriod() *float64 {
	if o == nil {
		return nil
	}
	return o.ClaimWaitingPeriod
}

func (o *CreatePolicyCommandRequest) GetPropertySetID() string {
	if o == nil {
		return ""
	}
	return o.PropertySetID
}

func (o *CreatePolicyCommandRequest) GetClaimLifecycleSteps() []ClaimLifecycleStepDto {
	if o == nil {
		return nil
	}
	return o.ClaimLifecycleSteps
}

func (o *CreatePolicyCommandRequest) GetCoverages() []CreateCoverageDto {
	if o == nil {
		return nil
	}
	return o.Coverages
}

func (o *CreatePolicyCommandRequest) GetTermsAndConditions() CreateDocumentDto {
	if o == nil {
		return CreateDocumentDto{}
	}
	return o.TermsAndConditions
}

func (o *CreatePolicyCommandRequest) GetCoverageSummary() CreateDocumentDto {
	if o == nil {
		return CreateDocumentDto{}
	}
	return o.CoverageSummary
}

func (o *CreatePolicyCommandRequest) GetTerms() []CreatePolicyTermDto {
	if o == nil {
		return []CreatePolicyTermDto{}
	}
	return o.Terms
}

func (o *CreatePolicyCommandRequest) GetRequiresWarrantyActivation() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresWarrantyActivation
}