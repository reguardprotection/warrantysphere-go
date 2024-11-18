// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CreateTransactionControllerCreateTransactionResponse struct {
	HTTPMeta                  components.HTTPMetadata `json:"-"`
	CreateTransactionResponse *components.CreateTransactionResponse
}

func (o *CreateTransactionControllerCreateTransactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTransactionControllerCreateTransactionResponse) GetCreateTransactionResponse() *components.CreateTransactionResponse {
	if o == nil {
		return nil
	}
	return o.CreateTransactionResponse
}