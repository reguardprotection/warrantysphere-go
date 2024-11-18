// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type RetrievePolicyQueryControllerRetrievePolicyRequest struct {
	// Unique identifier of the policy.
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
}

func (o *RetrievePolicyQueryControllerRetrievePolicyRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

type RetrievePolicyQueryControllerRetrievePolicyResponse struct {
	HTTPMeta        components.HTTPMetadata `json:"-"`
	PolicyAggregate *components.PolicyAggregate
}

func (o *RetrievePolicyQueryControllerRetrievePolicyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrievePolicyQueryControllerRetrievePolicyResponse) GetPolicyAggregate() *components.PolicyAggregate {
	if o == nil {
		return nil
	}
	return o.PolicyAggregate
}
