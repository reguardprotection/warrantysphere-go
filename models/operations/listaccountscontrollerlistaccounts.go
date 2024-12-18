// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

type AccountTypes string

const (
	AccountTypesBusiness   AccountTypes = "BUSINESS"
	AccountTypesIndividual AccountTypes = "INDIVIDUAL"
)

func (e AccountTypes) ToPointer() *AccountTypes {
	return &e
}
func (e *AccountTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUSINESS":
		fallthrough
	case "INDIVIDUAL":
		*e = AccountTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountTypes: %v", v)
	}
}

type ListAccountsControllerListAccountsRequest struct {
	Limit        *int64         `default:"10" queryParam:"style=form,explode=true,name=limit"`
	Page         *int64         `default:"1" queryParam:"style=form,explode=true,name=page"`
	AccountTypes []AccountTypes `queryParam:"style=form,explode=true,name=accountTypes"`
}

func (l ListAccountsControllerListAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsControllerListAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountsControllerListAccountsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountsControllerListAccountsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAccountsControllerListAccountsRequest) GetAccountTypes() []AccountTypes {
	if o == nil {
		return nil
	}
	return o.AccountTypes
}

type ListAccountsControllerListAccountsResponse struct {
	HTTPMeta             components.HTTPMetadata `json:"-"`
	ListAccountsResponse *components.ListAccountsResponse
}

func (o *ListAccountsControllerListAccountsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountsControllerListAccountsResponse) GetListAccountsResponse() *components.ListAccountsResponse {
	if o == nil {
		return nil
	}
	return o.ListAccountsResponse
}
