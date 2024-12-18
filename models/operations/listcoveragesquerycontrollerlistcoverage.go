// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ListCoveragesQueryControllerListCoverageRequest struct {
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page  *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Unique identifier of the policy.
	PolicyID string `queryParam:"style=form,explode=true,name=policyId"`
}

func (l ListCoveragesQueryControllerListCoverageRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCoveragesQueryControllerListCoverageRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCoveragesQueryControllerListCoverageRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCoveragesQueryControllerListCoverageRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListCoveragesQueryControllerListCoverageRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

type ListCoveragesQueryControllerListCoverageResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	ListCoveragesQueryResponse *components.ListCoveragesQueryResponse
}

func (o *ListCoveragesQueryControllerListCoverageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListCoveragesQueryControllerListCoverageResponse) GetListCoveragesQueryResponse() *components.ListCoveragesQueryResponse {
	if o == nil {
		return nil
	}
	return o.ListCoveragesQueryResponse
}
