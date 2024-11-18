// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ListPolicyPlansQueryControllerListPolicyPlansRequest struct {
	// Unique identifier of the policy
	PolicyID string `pathParam:"style=simple,explode=false,name=policyId"`
	Limit    *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page     *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (l ListPolicyPlansQueryControllerListPolicyPlansRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPolicyPlansQueryControllerListPolicyPlansRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPolicyPlansQueryControllerListPolicyPlansRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *ListPolicyPlansQueryControllerListPolicyPlansRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListPolicyPlansQueryControllerListPolicyPlansRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type ListPolicyPlansQueryControllerListPolicyPlansResponse struct {
	HTTPMeta                     components.HTTPMetadata `json:"-"`
	ListPolicyPlansQueryResponse *components.ListPolicyPlansQueryResponse
}

func (o *ListPolicyPlansQueryControllerListPolicyPlansResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListPolicyPlansQueryControllerListPolicyPlansResponse) GetListPolicyPlansQueryResponse() *components.ListPolicyPlansQueryResponse {
	if o == nil {
		return nil
	}
	return o.ListPolicyPlansQueryResponse
}