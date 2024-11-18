// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type SetVisibilityCommandControllerCreatePlanRequest struct {
	// Unique identifier of the policy.
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
	// Identifier of the linked plan.
	PlanID                          string                                     `pathParam:"style=simple,explode=false,name=planId"`
	SetPlanVisibilityCommandRequest components.SetPlanVisibilityCommandRequest `request:"mediaType=application/json"`
}

func (o *SetVisibilityCommandControllerCreatePlanRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *SetVisibilityCommandControllerCreatePlanRequest) GetPlanID() string {
	if o == nil {
		return ""
	}
	return o.PlanID
}

func (o *SetVisibilityCommandControllerCreatePlanRequest) GetSetPlanVisibilityCommandRequest() components.SetPlanVisibilityCommandRequest {
	if o == nil {
		return components.SetPlanVisibilityCommandRequest{}
	}
	return o.SetPlanVisibilityCommandRequest
}

type SetVisibilityCommandControllerCreatePlanResponse struct {
	HTTPMeta                         components.HTTPMetadata `json:"-"`
	SetPlanVisibilityCommandResponse *components.SetPlanVisibilityCommandResponse
}

func (o *SetVisibilityCommandControllerCreatePlanResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetVisibilityCommandControllerCreatePlanResponse) GetSetPlanVisibilityCommandResponse() *components.SetPlanVisibilityCommandResponse {
	if o == nil {
		return nil
	}
	return o.SetPlanVisibilityCommandResponse
}
