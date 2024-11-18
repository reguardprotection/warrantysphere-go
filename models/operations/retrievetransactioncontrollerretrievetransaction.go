// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type RetrieveTransactionControllerRetrieveTransactionRequest struct {
	// Unique identifier of the transaction
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RetrieveTransactionControllerRetrieveTransactionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RetrieveTransactionControllerRetrieveTransactionResponse struct {
	HTTPMeta             components.HTTPMetadata `json:"-"`
	TransactionAggregate *components.TransactionAggregate
}

func (o *RetrieveTransactionControllerRetrieveTransactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveTransactionControllerRetrieveTransactionResponse) GetTransactionAggregate() *components.TransactionAggregate {
	if o == nil {
		return nil
	}
	return o.TransactionAggregate
}
