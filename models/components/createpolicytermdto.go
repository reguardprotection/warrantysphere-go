// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreatePolicyTermDto struct {
	// Term duration (in days) of the policy term.
	Duration float64 `json:"duration"`
	// Friendly name of the policy term (visible in checkout).
	Title string `json:"title"`
	// Payment schedules available on the policy term.
	PaymentSchedules []CreatePaymentScheduleDto `json:"paymentSchedules,omitempty"`
	Order            float64                    `json:"order"`
}

func (o *CreatePolicyTermDto) GetDuration() float64 {
	if o == nil {
		return 0.0
	}
	return o.Duration
}

func (o *CreatePolicyTermDto) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *CreatePolicyTermDto) GetPaymentSchedules() []CreatePaymentScheduleDto {
	if o == nil {
		return nil
	}
	return o.PaymentSchedules
}

func (o *CreatePolicyTermDto) GetOrder() float64 {
	if o == nil {
		return 0.0
	}
	return o.Order
}
