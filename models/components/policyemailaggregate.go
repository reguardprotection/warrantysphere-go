// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

// PolicyEmailAggregatePolicy - The policy the email is linked to
type PolicyEmailAggregatePolicy struct {
}

type PolicyEmailAggregateType string

const (
	PolicyEmailAggregateTypeWarrantyProvisioned PolicyEmailAggregateType = "Warranty Provisioned"
)

func (e PolicyEmailAggregateType) ToPointer() *PolicyEmailAggregateType {
	return &e
}
func (e *PolicyEmailAggregateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Warranty Provisioned":
		*e = PolicyEmailAggregateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyEmailAggregateType: %v", v)
	}
}

type PolicyEmailAggregate struct {
	// Unique identifier of the policy email.
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	// Unique identifier of the policy.
	PolicyID string `json:"policyId"`
	// The policy the email is linked to
	Policy PolicyEmailAggregatePolicy `json:"policy"`
	Type   PolicyEmailAggregateType   `json:"type"`
	// Email  subject.
	Subject *string `default:"null" json:"subject"`
	// Email body.
	Content *string `default:"null" json:"content"`
	// Email byline.
	Byline *string `default:"null" json:"byline"`
	// Email instruction image.
	InstructionImage *string `default:"null" json:"instructionImage"`
	// Datetime when the policy email created.
	Created time.Time `json:"created"`
	// Datetime when the policy email modified.
	Modified time.Time `json:"modified"`
}

func (p PolicyEmailAggregate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyEmailAggregate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PolicyEmailAggregate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PolicyEmailAggregate) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *PolicyEmailAggregate) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *PolicyEmailAggregate) GetPolicy() PolicyEmailAggregatePolicy {
	if o == nil {
		return PolicyEmailAggregatePolicy{}
	}
	return o.Policy
}

func (o *PolicyEmailAggregate) GetType() PolicyEmailAggregateType {
	if o == nil {
		return PolicyEmailAggregateType("")
	}
	return o.Type
}

func (o *PolicyEmailAggregate) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *PolicyEmailAggregate) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *PolicyEmailAggregate) GetByline() *string {
	if o == nil {
		return nil
	}
	return o.Byline
}

func (o *PolicyEmailAggregate) GetInstructionImage() *string {
	if o == nil {
		return nil
	}
	return o.InstructionImage
}

func (o *PolicyEmailAggregate) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *PolicyEmailAggregate) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}
