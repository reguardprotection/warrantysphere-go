// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ListPaymentsQueryControllerListPaymentsRequest struct {
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page  *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
}

func (l ListPaymentsQueryControllerListPaymentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListPaymentsQueryControllerListPaymentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListPaymentsQueryControllerListPaymentsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListPaymentsQueryControllerListPaymentsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type ListPaymentsQueryControllerListPaymentsResponse struct {
	HTTPMeta             components.HTTPMetadata `json:"-"`
	ListPaymentsResponse *components.ListPaymentsResponse
}

func (o *ListPaymentsQueryControllerListPaymentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListPaymentsQueryControllerListPaymentsResponse) GetListPaymentsResponse() *components.ListPaymentsResponse {
	if o == nil {
		return nil
	}
	return o.ListPaymentsResponse
}
