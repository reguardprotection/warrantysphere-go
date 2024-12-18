// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DistributorDto struct {
	// User defined unique identifier for the distributor.
	ReferenceID *string `json:"referenceId,omitempty"`
	// Full name of the distributor.
	Name string `json:"name"`
	// Email address of the distributor.
	Email *string `json:"email,omitempty"`
	// Phone number of the distributor. Unique across distributors.
	Phone *string `json:"phone,omitempty"`
	// Unique identifier of the distributor under which this new distributor will be found.
	ParentID   *string `json:"parentId,omitempty"`
	ParentName *string `json:"parentName,omitempty"`
}

func (o *DistributorDto) GetReferenceID() *string {
	if o == nil {
		return nil
	}
	return o.ReferenceID
}

func (o *DistributorDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DistributorDto) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *DistributorDto) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *DistributorDto) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *DistributorDto) GetParentName() *string {
	if o == nil {
		return nil
	}
	return o.ParentName
}
