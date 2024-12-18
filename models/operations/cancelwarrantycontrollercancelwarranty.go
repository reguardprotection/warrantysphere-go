// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"github.com/reguardprotection/warrantysphere/models/components"
)

// QueryParamIDField - Field used to identify the warranty.
type QueryParamIDField string

const (
	QueryParamIDFieldID             QueryParamIDField = "id"
	QueryParamIDFieldReferenceID    QueryParamIDField = "referenceId"
	QueryParamIDFieldContractNumber QueryParamIDField = "contractNumber"
)

func (e QueryParamIDField) ToPointer() *QueryParamIDField {
	return &e
}
func (e *QueryParamIDField) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "referenceId":
		fallthrough
	case "contractNumber":
		*e = QueryParamIDField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamIDField: %v", v)
	}
}

type CancelWarrantyControllerCancelWarrantyRequest struct {
	// Unique ID of the warranty.
	UniqueID string `pathParam:"style=simple,explode=false,name=uniqueId"`
	// Field used to identify the warranty.
	IDField                   QueryParamIDField                    `default:"id" queryParam:"style=form,explode=true,name=idField"`
	CancelWarrantyRequestBody components.CancelWarrantyRequestBody `request:"mediaType=application/json"`
}

func (c CancelWarrantyControllerCancelWarrantyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CancelWarrantyControllerCancelWarrantyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CancelWarrantyControllerCancelWarrantyRequest) GetUniqueID() string {
	if o == nil {
		return ""
	}
	return o.UniqueID
}

func (o *CancelWarrantyControllerCancelWarrantyRequest) GetIDField() QueryParamIDField {
	if o == nil {
		return QueryParamIDField("")
	}
	return o.IDField
}

func (o *CancelWarrantyControllerCancelWarrantyRequest) GetCancelWarrantyRequestBody() components.CancelWarrantyRequestBody {
	if o == nil {
		return components.CancelWarrantyRequestBody{}
	}
	return o.CancelWarrantyRequestBody
}

type CancelWarrantyControllerCancelWarrantyResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	CancelWarrantyResponseBody *components.CancelWarrantyResponseBody
}

func (o *CancelWarrantyControllerCancelWarrantyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CancelWarrantyControllerCancelWarrantyResponse) GetCancelWarrantyResponseBody() *components.CancelWarrantyResponseBody {
	if o == nil {
		return nil
	}
	return o.CancelWarrantyResponseBody
}
