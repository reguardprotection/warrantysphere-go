// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type RetrievePolicyPlanQueryControllerRetrievePolicyPlanRequest struct {
	// Unique identifier of the wpolicylan.
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
	// Unique identifier of the plan.
	PlanID string `pathParam:"style=simple,explode=false,name=planId"`
}

func (o *RetrievePolicyPlanQueryControllerRetrievePolicyPlanRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *RetrievePolicyPlanQueryControllerRetrievePolicyPlanRequest) GetPlanID() string {
	if o == nil {
		return ""
	}
	return o.PlanID
}

type RetrievePolicyPlanQueryControllerRetrievePolicyPlanResponse struct {
	HTTPMeta            components.HTTPMetadata `json:"-"`
	PolicyPlanAggregate *components.PolicyPlanAggregate
}

func (o *RetrievePolicyPlanQueryControllerRetrievePolicyPlanResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrievePolicyPlanQueryControllerRetrievePolicyPlanResponse) GetPolicyPlanAggregate() *components.PolicyPlanAggregate {
	if o == nil {
		return nil
	}
	return o.PolicyPlanAggregate
}