// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RetryPaymentRequestBody struct {
	// The unique identifier of the payment method.
	PaymentMethodID *string `json:"paymentMethodId,omitempty"`
}

func (o *RetryPaymentRequestBody) GetPaymentMethodID() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}
