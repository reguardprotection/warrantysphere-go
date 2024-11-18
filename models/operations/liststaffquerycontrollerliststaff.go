// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ListStaffQueryControllerListStaffRequest struct {
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page  *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Search string used to match listed staff's names or emails.
	Search *string `queryParam:"style=form,explode=true,name=search"`
	// Unique identifiers of the roles the the listed staff must have.
	RoleIds *string `queryParam:"style=form,explode=true,name=roleIds"`
	// Unique identifiers of the distributor the the listed staff is associated to.
	DistributorID *string `queryParam:"style=form,explode=true,name=distributorId"`
}

func (l ListStaffQueryControllerListStaffRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListStaffQueryControllerListStaffRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListStaffQueryControllerListStaffRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListStaffQueryControllerListStaffRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListStaffQueryControllerListStaffRequest) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *ListStaffQueryControllerListStaffRequest) GetRoleIds() *string {
	if o == nil {
		return nil
	}
	return o.RoleIds
}

func (o *ListStaffQueryControllerListStaffRequest) GetDistributorID() *string {
	if o == nil {
		return nil
	}
	return o.DistributorID
}

type ListStaffQueryControllerListStaffResponse struct {
	HTTPMeta               components.HTTPMetadata `json:"-"`
	ListStaffQueryResponse *components.ListStaffQueryResponse
}

func (o *ListStaffQueryControllerListStaffResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListStaffQueryControllerListStaffResponse) GetListStaffQueryResponse() *components.ListStaffQueryResponse {
	if o == nil {
		return nil
	}
	return o.ListStaffQueryResponse
}