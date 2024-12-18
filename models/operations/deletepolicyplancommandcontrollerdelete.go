// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type DeletePolicyPlanCommandControllerDeleteRequest struct {
	// Unique identifier of the wpolicylan.
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
	// Unique identifier of the policies plan.
	PlanID                             string                                        `pathParam:"style=simple,explode=false,name=planId"`
	DeletePolicyPlanCommandRequestBody components.DeletePolicyPlanCommandRequestBody `request:"mediaType=application/json"`
}

func (o *DeletePolicyPlanCommandControllerDeleteRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *DeletePolicyPlanCommandControllerDeleteRequest) GetPlanID() string {
	if o == nil {
		return ""
	}
	return o.PlanID
}

func (o *DeletePolicyPlanCommandControllerDeleteRequest) GetDeletePolicyPlanCommandRequestBody() components.DeletePolicyPlanCommandRequestBody {
	if o == nil {
		return components.DeletePolicyPlanCommandRequestBody{}
	}
	return o.DeletePolicyPlanCommandRequestBody
}

type DeletePolicyPlanCommandControllerDeleteResponse struct {
	HTTPMeta                        components.HTTPMetadata `json:"-"`
	DeletePolicyPlanCommandResponse *components.DeletePolicyPlanCommandResponse
}

func (o *DeletePolicyPlanCommandControllerDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeletePolicyPlanCommandControllerDeleteResponse) GetDeletePolicyPlanCommandResponse() *components.DeletePolicyPlanCommandResponse {
	if o == nil {
		return nil
	}
	return o.DeletePolicyPlanCommandResponse
}
