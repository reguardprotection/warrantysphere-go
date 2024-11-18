// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdateCustomerControllerUpdateCustomerRequest struct {
	// The unique identifier of the customer.
	CustomerID            string                           `pathParam:"style=simple,explode=false,name=customerId"`
	UpdateCustomerRequest components.UpdateCustomerRequest `request:"mediaType=application/json"`
}

func (o *UpdateCustomerControllerUpdateCustomerRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *UpdateCustomerControllerUpdateCustomerRequest) GetUpdateCustomerRequest() components.UpdateCustomerRequest {
	if o == nil {
		return components.UpdateCustomerRequest{}
	}
	return o.UpdateCustomerRequest
}

type UpdateCustomerControllerUpdateCustomerResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	UpdateCustomerResponseBody *components.UpdateCustomerResponseBody
}

func (o *UpdateCustomerControllerUpdateCustomerResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateCustomerControllerUpdateCustomerResponse) GetUpdateCustomerResponseBody() *components.UpdateCustomerResponseBody {
	if o == nil {
		return nil
	}
	return o.UpdateCustomerResponseBody
}