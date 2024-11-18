// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateDistributorsSettingsCommandRequest struct {
	// Setting allowing users to be registered from distributor onboarding.
	CanRegisterUsers        *bool            `json:"canRegisterUsers,omitempty"`
	SelfServeInvoicePackage SelfServePackage `json:"selfServeInvoicePackage"`
}

func (o *UpdateDistributorsSettingsCommandRequest) GetCanRegisterUsers() *bool {
	if o == nil {
		return nil
	}
	return o.CanRegisterUsers
}

func (o *UpdateDistributorsSettingsCommandRequest) GetSelfServeInvoicePackage() SelfServePackage {
	if o == nil {
		return SelfServePackage{}
	}
	return o.SelfServeInvoicePackage
}