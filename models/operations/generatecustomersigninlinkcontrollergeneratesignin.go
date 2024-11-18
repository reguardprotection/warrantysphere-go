// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type GenerateCustomerSignInLinkControllerGenerateSignInRequest struct {
	// Unique identifier of the customer
	UniqueID                         string                                      `pathParam:"style=simple,explode=false,name=uniqueId"`
	GenerateSignInLinkCommandRequest components.GenerateSignInLinkCommandRequest `request:"mediaType=application/json"`
}

func (o *GenerateCustomerSignInLinkControllerGenerateSignInRequest) GetUniqueID() string {
	if o == nil {
		return ""
	}
	return o.UniqueID
}

func (o *GenerateCustomerSignInLinkControllerGenerateSignInRequest) GetGenerateSignInLinkCommandRequest() components.GenerateSignInLinkCommandRequest {
	if o == nil {
		return components.GenerateSignInLinkCommandRequest{}
	}
	return o.GenerateSignInLinkCommandRequest
}

type GenerateCustomerSignInLinkControllerGenerateSignInResponse struct {
	HTTPMeta                          components.HTTPMetadata `json:"-"`
	GenerateSignInLinkCommandResponse *components.GenerateSignInLinkCommandResponse
}

func (o *GenerateCustomerSignInLinkControllerGenerateSignInResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenerateCustomerSignInLinkControllerGenerateSignInResponse) GetGenerateSignInLinkCommandResponse() *components.GenerateSignInLinkCommandResponse {
	if o == nil {
		return nil
	}
	return o.GenerateSignInLinkCommandResponse
}