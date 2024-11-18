// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PolicyPlanCoverageFeeAggregatePolicy - Policy linked to the plan coverage fee.
type PolicyPlanCoverageFeeAggregatePolicy struct {
}

// Plan linked to the plan coverage fee.
type Plan struct {
}

// Term linked to the plan coverage fee.
type Term struct {
}

// Coverage linked to the plan coverage fee.
type Coverage struct {
}

// FeeType - Type of fee for the plan coverage fee.
type FeeType struct {
}

// Rule - Condition rule used to determine alternate values.
type Rule struct {
}

type PolicyPlanCoverageFeeAggregate struct {
	// Unique identifier of the plan coverage fee.
	ID string `json:"id"`
	// Unique identifier of the policy.
	PolicyID string `json:"policyId"`
	// Policy linked to the plan coverage fee.
	Policy *PolicyPlanCoverageFeeAggregatePolicy `json:"policy,omitempty"`
	// Unique identifier of the plan.
	PlanID string `json:"planId"`
	// Plan linked to the plan coverage fee.
	Plan *Plan `json:"plan,omitempty"`
	// Unique identifier of the term.
	TermID string `json:"termId"`
	// Term linked to the plan coverage fee.
	Term *Term `json:"term,omitempty"`
	// Unique identifier of the coverage.
	CoverageID string `json:"coverageId"`
	// Coverage linked to the plan coverage fee.
	Coverage *Coverage `json:"coverage,omitempty"`
	// Title of the plan coverage fee
	Title string `json:"title"`
	// Type of fee for the plan coverage fee.
	FeeType FeeType `json:"feeType"`
	// Condition rule used to determine alternate values.
	Rule Rule `json:"rule"`
}

func (o *PolicyPlanCoverageFeeAggregate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PolicyPlanCoverageFeeAggregate) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *PolicyPlanCoverageFeeAggregate) GetPolicy() *PolicyPlanCoverageFeeAggregatePolicy {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *PolicyPlanCoverageFeeAggregate) GetPlanID() string {
	if o == nil {
		return ""
	}
	return o.PlanID
}

func (o *PolicyPlanCoverageFeeAggregate) GetPlan() *Plan {
	if o == nil {
		return nil
	}
	return o.Plan
}

func (o *PolicyPlanCoverageFeeAggregate) GetTermID() string {
	if o == nil {
		return ""
	}
	return o.TermID
}

func (o *PolicyPlanCoverageFeeAggregate) GetTerm() *Term {
	if o == nil {
		return nil
	}
	return o.Term
}

func (o *PolicyPlanCoverageFeeAggregate) GetCoverageID() string {
	if o == nil {
		return ""
	}
	return o.CoverageID
}

func (o *PolicyPlanCoverageFeeAggregate) GetCoverage() *Coverage {
	if o == nil {
		return nil
	}
	return o.Coverage
}

func (o *PolicyPlanCoverageFeeAggregate) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *PolicyPlanCoverageFeeAggregate) GetFeeType() FeeType {
	if o == nil {
		return FeeType{}
	}
	return o.FeeType
}

func (o *PolicyPlanCoverageFeeAggregate) GetRule() Rule {
	if o == nil {
		return Rule{}
	}
	return o.Rule
}
