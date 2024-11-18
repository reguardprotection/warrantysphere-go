// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ListClaimItemsQueryControllerListRequest struct {
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page  *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Unique identifier of the coverage used to filter the claim items.
	CoverageID *string `queryParam:"style=form,explode=true,name=coverageId"`
	// Search used to filter claim items based on their coverage title.
	Search *string `queryParam:"style=form,explode=true,name=search"`
	// Option to include deleted claim item in the search.
	WithDeleted *bool `queryParam:"style=form,explode=true,name=withDeleted"`
	// Unique identifier of the claim used to filter the claim items.
	ClaimID string `pathParam:"style=simple,explode=false,name=claimId"`
	// List of permission keys
	XWspherePermissionKeys *string `header:"style=simple,explode=false,name=X-WSPHERE-PERMISSION-KEYS"`
}

func (l ListClaimItemsQueryControllerListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListClaimItemsQueryControllerListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListClaimItemsQueryControllerListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListClaimItemsQueryControllerListRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListClaimItemsQueryControllerListRequest) GetCoverageID() *string {
	if o == nil {
		return nil
	}
	return o.CoverageID
}

func (o *ListClaimItemsQueryControllerListRequest) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *ListClaimItemsQueryControllerListRequest) GetWithDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.WithDeleted
}

func (o *ListClaimItemsQueryControllerListRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *ListClaimItemsQueryControllerListRequest) GetXWspherePermissionKeys() *string {
	if o == nil {
		return nil
	}
	return o.XWspherePermissionKeys
}

type ListClaimItemsQueryControllerListResponse struct {
	HTTPMeta                    components.HTTPMetadata `json:"-"`
	ListClaimItemsQueryResponse *components.ListClaimItemsQueryResponse
}

func (o *ListClaimItemsQueryControllerListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListClaimItemsQueryControllerListResponse) GetListClaimItemsQueryResponse() *components.ListClaimItemsQueryResponse {
	if o == nil {
		return nil
	}
	return o.ListClaimItemsQueryResponse
}
