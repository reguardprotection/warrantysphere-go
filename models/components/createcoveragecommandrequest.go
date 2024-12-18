// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateCoverageCommandRequest struct {
	// Whether or not the command should create a new draft policy if required.
	CanCreateNewVersion bool `json:"canCreateNewVersion"`
	// Coverage order
	Order *float64 `json:"order,omitempty"`
	// Unique identifier of the policy this coverage is linked to.
	PolicyID string `json:"policyId"`
	// User-defined reference Id.
	ReferenceID *string `json:"referenceId,omitempty"`
	// Title of coverage
	Title string `json:"title"`
	// Description of the coverage being provided
	Description string `json:"description"`
	// The group the coverage is categorized in
	Group string `json:"group"`
	// Details about what is included in the coverage
	Inclusions *string `json:"inclusions"`
	// Details about what is excluded in the coverage
	Exclusions *string `json:"exclusions"`
}

func (o *CreateCoverageCommandRequest) GetCanCreateNewVersion() bool {
	if o == nil {
		return false
	}
	return o.CanCreateNewVersion
}

func (o *CreateCoverageCommandRequest) GetOrder() *float64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *CreateCoverageCommandRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *CreateCoverageCommandRequest) GetReferenceID() *string {
	if o == nil {
		return nil
	}
	return o.ReferenceID
}

func (o *CreateCoverageCommandRequest) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *CreateCoverageCommandRequest) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CreateCoverageCommandRequest) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *CreateCoverageCommandRequest) GetInclusions() *string {
	if o == nil {
		return nil
	}
	return o.Inclusions
}

func (o *CreateCoverageCommandRequest) GetExclusions() *string {
	if o == nil {
		return nil
	}
	return o.Exclusions
}
