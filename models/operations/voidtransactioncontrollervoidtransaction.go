// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type VoidTransactionControllerVoidTransactionRequest struct {
	// Unique identifier of the transaction
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *VoidTransactionControllerVoidTransactionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type VoidTransactionControllerVoidTransactionResponse struct {
	HTTPMeta                components.HTTPMetadata `json:"-"`
	VoidTransactionResponse *components.VoidTransactionResponse
}

func (o *VoidTransactionControllerVoidTransactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VoidTransactionControllerVoidTransactionResponse) GetVoidTransactionResponse() *components.VoidTransactionResponse {
	if o == nil {
		return nil
	}
	return o.VoidTransactionResponse
}
