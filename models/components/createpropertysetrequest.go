// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreatePropertySetRequest struct {
	// The name of the property set
	Name string `json:"name"`
	// The icon of the property set
	Icon string `json:"icon"`
	// The properties and their configuration of this property set
	Properties []CreatePropertyAssignment `json:"properties"`
}

func (o *CreatePropertySetRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreatePropertySetRequest) GetIcon() string {
	if o == nil {
		return ""
	}
	return o.Icon
}

func (o *CreatePropertySetRequest) GetProperties() []CreatePropertyAssignment {
	if o == nil {
		return []CreatePropertyAssignment{}
	}
	return o.Properties
}
