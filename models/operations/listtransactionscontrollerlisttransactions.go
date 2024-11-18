// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ListTransactionsControllerListTransactionsRequest struct {
	Limit *int64 `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page  *int64 `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Unique identifier of the source or destination account of the transactions
	AccountID string `queryParam:"style=form,explode=true,name=accountId"`
}

func (l ListTransactionsControllerListTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTransactionsControllerListTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTransactionsControllerListTransactionsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListTransactionsControllerListTransactionsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListTransactionsControllerListTransactionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type ListTransactionsControllerListTransactionsResponse struct {
	HTTPMeta                 components.HTTPMetadata `json:"-"`
	ListTransactionsResponse *components.ListTransactionsResponse
}

func (o *ListTransactionsControllerListTransactionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTransactionsControllerListTransactionsResponse) GetListTransactionsResponse() *components.ListTransactionsResponse {
	if o == nil {
		return nil
	}
	return o.ListTransactionsResponse
}