// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

// AccountAggregateStatus - Activation status of the distributor.
type AccountAggregateStatus struct {
}

// AccountAggregateDistributor - Distributor that is linked to the account
type AccountAggregateDistributor struct {
	// Unique identifier of the distributor
	ID string `json:"id"`
	// User-defined reference ID.
	ReferenceID *string `json:"referenceId"`
	// Datetime when the distributor was created.
	Created time.Time `json:"created"`
	// Datetime when the distributor was last modified.
	Modified time.Time `json:"modified"`
	Name     string    `json:"name"`
	Email    *string   `json:"email"`
	Phone    *string   `json:"phone"`
	// Activation status of the distributor.
	Status AccountAggregateStatus `json:"status"`
	// Parent distributor of this distributor.
	Parent *DistributorAggregate `json:"parent,omitempty"`
	// Unique identifier of the parent distributor of the distributor.
	ParentID *string `json:"parentId"`
	// The computed permission key based off of the parent.
	PermissionKey string `json:"permissionKey"`
	// Parent distributor of this distributor.
	OnboardingLink *DistributorOnboardingLinkAggregate `json:"onboardingLink,omitempty"`
	// Unique identifier of the parent distributor of the distributor.
	OnboardingLinkID *string `json:"onboardingLinkId"`
	// Activation date of the Distributor.
	ActivatedDate time.Time `json:"activatedDate"`
	// Deactivation date of the Distributor.
	DeactivatedDate time.Time `json:"deactivatedDate"`
	// Archived date of the Distributor.
	ArchivedDate time.Time `json:"archivedDate"`
}

func (a AccountAggregateDistributor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountAggregateDistributor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountAggregateDistributor) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountAggregateDistributor) GetReferenceID() *string {
	if o == nil {
		return nil
	}
	return o.ReferenceID
}

func (o *AccountAggregateDistributor) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *AccountAggregateDistributor) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}

func (o *AccountAggregateDistributor) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AccountAggregateDistributor) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AccountAggregateDistributor) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *AccountAggregateDistributor) GetStatus() AccountAggregateStatus {
	if o == nil {
		return AccountAggregateStatus{}
	}
	return o.Status
}

func (o *AccountAggregateDistributor) GetParent() *DistributorAggregate {
	if o == nil {
		return nil
	}
	return o.Parent
}

func (o *AccountAggregateDistributor) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *AccountAggregateDistributor) GetPermissionKey() string {
	if o == nil {
		return ""
	}
	return o.PermissionKey
}

func (o *AccountAggregateDistributor) GetOnboardingLink() *DistributorOnboardingLinkAggregate {
	if o == nil {
		return nil
	}
	return o.OnboardingLink
}

func (o *AccountAggregateDistributor) GetOnboardingLinkID() *string {
	if o == nil {
		return nil
	}
	return o.OnboardingLinkID
}

func (o *AccountAggregateDistributor) GetActivatedDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ActivatedDate
}

func (o *AccountAggregateDistributor) GetDeactivatedDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.DeactivatedDate
}

func (o *AccountAggregateDistributor) GetArchivedDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ArchivedDate
}

// AccountAggregateAddress - Shipping address of the customer
type AccountAggregateAddress struct {
	Line1   string  `json:"line1"`
	Line2   *string `json:"line2"`
	Zip     string  `json:"zip"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Country string  `json:"country"`
}

func (o *AccountAggregateAddress) GetLine1() string {
	if o == nil {
		return ""
	}
	return o.Line1
}

func (o *AccountAggregateAddress) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *AccountAggregateAddress) GetZip() string {
	if o == nil {
		return ""
	}
	return o.Zip
}

func (o *AccountAggregateAddress) GetCity() string {
	if o == nil {
		return ""
	}
	return o.City
}

func (o *AccountAggregateAddress) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *AccountAggregateAddress) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

// Customer that is linked to the account
type Customer struct {
	// Unique identifier for the customer.
	ID string `json:"id"`
	// User-defined reference ID.
	ReferenceID *string `json:"referenceId"`
	// The customer's full name.
	Name string `json:"name"`
	// The customer's email address.
	Email string `json:"email"`
	// The customer's phone number.
	Phone *string `json:"phone"`
	// Unique identifier for the user.
	UserID string `json:"userId"`
	// User linked to the customer.
	User *UserAggregate `json:"user,omitempty"`
	// Datetime when the object was created.
	Created time.Time `json:"created"`
	// Datetime when the object was last modified.
	Modified time.Time `json:"modified"`
	// Shipping address of the customer
	Address *AccountAggregateAddress `json:"address"`
	// JSON object of metadata related to the customer.
	Metadata *string `json:"metadata"`
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Customer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Customer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Customer) GetReferenceID() *string {
	if o == nil {
		return nil
	}
	return o.ReferenceID
}

func (o *Customer) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Customer) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *Customer) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Customer) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *Customer) GetUser() *UserAggregate {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Customer) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *Customer) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}

func (o *Customer) GetAddress() *AccountAggregateAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Customer) GetMetadata() *string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type OwnerType string

const (
	OwnerTypeBusiness OwnerType = "BUSINESS"
)

func (e OwnerType) ToPointer() *OwnerType {
	return &e
}
func (e *OwnerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUSINESS":
		*e = OwnerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OwnerType: %v", v)
	}
}

type BusinessOwnerAggregate struct {
	Type OwnerType `json:"type"`
	// Name of the account owner
	Name string `json:"name"`
	// Email address of the account owner
	Email string `json:"email"`
	// Phone number of the account owner
	Phone *string `json:"phone"`
	// Business type of the account owner
	BusinessType *string `json:"businessType"`
	// Physical address of the account owner's business
	Address AddressAggregate `json:"address"`
}

func (o *BusinessOwnerAggregate) GetType() OwnerType {
	if o == nil {
		return OwnerType("")
	}
	return o.Type
}

func (o *BusinessOwnerAggregate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *BusinessOwnerAggregate) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *BusinessOwnerAggregate) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *BusinessOwnerAggregate) GetBusinessType() *string {
	if o == nil {
		return nil
	}
	return o.BusinessType
}

func (o *BusinessOwnerAggregate) GetAddress() AddressAggregate {
	if o == nil {
		return AddressAggregate{}
	}
	return o.Address
}

type IndividualOwnerAggregateOwnerType string

const (
	IndividualOwnerAggregateOwnerTypeIndividual IndividualOwnerAggregateOwnerType = "INDIVIDUAL"
)

func (e IndividualOwnerAggregateOwnerType) ToPointer() *IndividualOwnerAggregateOwnerType {
	return &e
}
func (e *IndividualOwnerAggregateOwnerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INDIVIDUAL":
		*e = IndividualOwnerAggregateOwnerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IndividualOwnerAggregateOwnerType: %v", v)
	}
}

type IndividualOwnerAggregate struct {
	Type IndividualOwnerAggregateOwnerType `json:"type"`
	// Name of the account owner
	Name string `json:"name"`
	// Email address of the account owner
	Email string `json:"email"`
	// Phone number of the account owner
	Phone *string `json:"phone"`
}

func (o *IndividualOwnerAggregate) GetType() IndividualOwnerAggregateOwnerType {
	if o == nil {
		return IndividualOwnerAggregateOwnerType("")
	}
	return o.Type
}

func (o *IndividualOwnerAggregate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *IndividualOwnerAggregate) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *IndividualOwnerAggregate) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

type OwnerUnionType string

const (
	OwnerUnionTypeIndividualOwnerAggregate OwnerUnionType = "IndividualOwnerAggregate"
	OwnerUnionTypeBusinessOwnerAggregate   OwnerUnionType = "BusinessOwnerAggregate"
)

// Owner - Account owner information
type Owner struct {
	IndividualOwnerAggregate *IndividualOwnerAggregate
	BusinessOwnerAggregate   *BusinessOwnerAggregate

	Type OwnerUnionType
}

func CreateOwnerIndividualOwnerAggregate(individualOwnerAggregate IndividualOwnerAggregate) Owner {
	typ := OwnerUnionTypeIndividualOwnerAggregate

	return Owner{
		IndividualOwnerAggregate: &individualOwnerAggregate,
		Type:                     typ,
	}
}

func CreateOwnerBusinessOwnerAggregate(businessOwnerAggregate BusinessOwnerAggregate) Owner {
	typ := OwnerUnionTypeBusinessOwnerAggregate

	return Owner{
		BusinessOwnerAggregate: &businessOwnerAggregate,
		Type:                   typ,
	}
}

func (u *Owner) UnmarshalJSON(data []byte) error {

	var individualOwnerAggregate IndividualOwnerAggregate = IndividualOwnerAggregate{}
	if err := utils.UnmarshalJSON(data, &individualOwnerAggregate, "", true, true); err == nil {
		u.IndividualOwnerAggregate = &individualOwnerAggregate
		u.Type = OwnerUnionTypeIndividualOwnerAggregate
		return nil
	}

	var businessOwnerAggregate BusinessOwnerAggregate = BusinessOwnerAggregate{}
	if err := utils.UnmarshalJSON(data, &businessOwnerAggregate, "", true, true); err == nil {
		u.BusinessOwnerAggregate = &businessOwnerAggregate
		u.Type = OwnerUnionTypeBusinessOwnerAggregate
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Owner", string(data))
}

func (u Owner) MarshalJSON() ([]byte, error) {
	if u.IndividualOwnerAggregate != nil {
		return utils.MarshalJSON(u.IndividualOwnerAggregate, "", true)
	}

	if u.BusinessOwnerAggregate != nil {
		return utils.MarshalJSON(u.BusinessOwnerAggregate, "", true)
	}

	return nil, errors.New("could not marshal union type Owner: all fields are null")
}

type AccountAggregate struct {
	// Unique identifier for the account.
	ID string `json:"id"`
	// Account name
	Name string `json:"name"`
	// Account description
	Description *string `json:"description"`
	// Account currency
	Currency string `json:"currency"`
	// Unique ID of a distributor that is linked to the account
	DistributorID *string `json:"distributorId"`
	// Distributor that is linked to the account
	Distributor *AccountAggregateDistributor `json:"distributor,omitempty"`
	// Unique ID of a customer that is linked to the account
	CustomerID *string `json:"customerId"`
	// Customer that is linked to the account
	Customer *Customer `json:"customer,omitempty"`
	// Datetime when the object was created.
	Created time.Time `json:"created"`
	// Datetime when the object was last modified.
	Modified time.Time `json:"modified"`
	// Account owner information
	Owner *Owner `json:"owner"`
}

func (a AccountAggregate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountAggregate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountAggregate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountAggregate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AccountAggregate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountAggregate) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *AccountAggregate) GetDistributorID() *string {
	if o == nil {
		return nil
	}
	return o.DistributorID
}

func (o *AccountAggregate) GetDistributor() *AccountAggregateDistributor {
	if o == nil {
		return nil
	}
	return o.Distributor
}

func (o *AccountAggregate) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *AccountAggregate) GetCustomer() *Customer {
	if o == nil {
		return nil
	}
	return o.Customer
}

func (o *AccountAggregate) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *AccountAggregate) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}

func (o *AccountAggregate) GetOwner() *Owner {
	if o == nil {
		return nil
	}
	return o.Owner
}
