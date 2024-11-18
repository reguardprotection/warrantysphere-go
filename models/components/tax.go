// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type TaxType string

const (
	TaxTypeCity    TaxType = "city"
	TaxTypeCountry TaxType = "country"
	TaxTypeState   TaxType = "state"
	TaxTypeCounty  TaxType = "county"
	TaxTypeGst     TaxType = "gst"
	TaxTypeQst     TaxType = "qst"
	TaxTypePst     TaxType = "pst"
	TaxTypeOther   TaxType = "other"
)

func (e TaxType) ToPointer() *TaxType {
	return &e
}
func (e *TaxType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "city":
		fallthrough
	case "country":
		fallthrough
	case "state":
		fallthrough
	case "county":
		fallthrough
	case "gst":
		fallthrough
	case "qst":
		fallthrough
	case "pst":
		fallthrough
	case "other":
		*e = TaxType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxType: %v", v)
	}
}

type Tax struct {
	Type    TaxType `json:"type"`
	Label   string  `json:"label"`
	Amount  int64   `json:"amount"`
	Rate    int64   `json:"rate"`
	City    string  `json:"city"`
	County  *string `json:"county"`
	State   string  `json:"state"`
	Country string  `json:"country"`
}

func (o *Tax) GetType() TaxType {
	if o == nil {
		return TaxType("")
	}
	return o.Type
}

func (o *Tax) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *Tax) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

func (o *Tax) GetRate() int64 {
	if o == nil {
		return 0
	}
	return o.Rate
}

func (o *Tax) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *Tax) GetCounty() *string {
	if o == nil {
		return nil
	}
	return o.County
}

func (o *Tax) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *Tax) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}