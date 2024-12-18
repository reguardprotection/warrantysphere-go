// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type CalculateTaxesControllerCalculateTaxesResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	CalculateTaxesResponseBody *components.CalculateTaxesResponseBody
}

func (o *CalculateTaxesControllerCalculateTaxesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CalculateTaxesControllerCalculateTaxesResponse) GetCalculateTaxesResponseBody() *components.CalculateTaxesResponseBody {
	if o == nil {
		return nil
	}
	return o.CalculateTaxesResponseBody
}
