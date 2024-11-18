// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AddressAggregate struct {
	Line1   string  `json:"line1"`
	Line2   *string `json:"line2"`
	Zip     string  `json:"zip"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Country string  `json:"country"`
}

func (o *AddressAggregate) GetLine1() string {
	if o == nil {
		return ""
	}
	return o.Line1
}

func (o *AddressAggregate) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *AddressAggregate) GetZip() string {
	if o == nil {
		return ""
	}
	return o.Zip
}

func (o *AddressAggregate) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *AddressAggregate) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *AddressAggregate) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}