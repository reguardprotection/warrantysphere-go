// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/reguardprotection/warrantysphere/internal/utils"
	"time"
)

// AccountType - The type of bank account used for the eCheck.Net transaction
type AccountType struct {
}

// BankAccount - Bank account details of the payment method (only present when the payment method is for a bank account)
type BankAccount struct {
	// The type of bank account used for the eCheck.Net transaction
	AccountType AccountType `json:"accountType"`
	// The masked bank account number
	AccountNumber string `json:"accountNumber"`
	// Name of the person who holds the bank account
	NameOnAccount string `json:"nameOnAccount"`
	// The name of the bank
	BankName string `json:"bankName"`
}

func (o *BankAccount) GetAccountType() AccountType {
	if o == nil {
		return AccountType{}
	}
	return o.AccountType
}

func (o *BankAccount) GetAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.AccountNumber
}

func (o *BankAccount) GetNameOnAccount() string {
	if o == nil {
		return ""
	}
	return o.NameOnAccount
}

func (o *BankAccount) GetBankName() string {
	if o == nil {
		return ""
	}
	return o.BankName
}

// CardType - Type of credit card
type CardType string

const (
	CardTypeVisa            CardType = "Visa"
	CardTypeMasterCard      CardType = "MasterCard"
	CardTypeAmericanExpress CardType = "AmericanExpress"
	CardTypeDiscover        CardType = "Discover"
	CardTypeJcb             CardType = "JCB"
	CardTypeDinersClub      CardType = "DinersClub"
)

func (e CardType) ToPointer() *CardType {
	return &e
}
func (e *CardType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Visa":
		fallthrough
	case "MasterCard":
		fallthrough
	case "AmericanExpress":
		fallthrough
	case "Discover":
		fallthrough
	case "JCB":
		fallthrough
	case "DinersClub":
		*e = CardType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardType: %v", v)
	}
}

// CreditCard - Credit card details of the payment method (only present when the payment method is for a credit card)
type CreditCard struct {
	// The customer's masked credit card number
	CardNumber string `json:"cardNumber"`
	// The expiration date for the customer's credit card (7 characters)
	ExpirationDate string `json:"expirationDate"`
	// Type of credit card
	CardType CardType `json:"cardType"`
}

func (o *CreditCard) GetCardNumber() string {
	if o == nil {
		return ""
	}
	return o.CardNumber
}

func (o *CreditCard) GetExpirationDate() string {
	if o == nil {
		return ""
	}
	return o.ExpirationDate
}

func (o *CreditCard) GetCardType() CardType {
	if o == nil {
		return CardType("")
	}
	return o.CardType
}

type PaymentMethodAggregate struct {
	// Unique ID of the payment method
	ID string `json:"id"`
	// Unique ID of the payment method in its associated gateway
	ExternalID *string `json:"externalId"`
	// Unique ID of the payment method's user
	UserID string `json:"userId"`
	// User details of the payment method
	User *UserAggregate `json:"user,omitempty"`
	// Unique ID of the payment method's associated gateway
	Gateway string `json:"gateway"`
	// Datetime when the object was created.
	Created time.Time `json:"created"`
	// Datetime when the object was last modified.
	Modified time.Time `json:"modified"`
	// Billing address of the payment method
	BillingAddress AddressAggregate `json:"billingAddress"`
	// Billing details of the payment method
	BillingDetails BillingDetailsAggregate `json:"billingDetails"`
	// Bank account details of the payment method (only present when the payment method is for a bank account)
	BankAccount *BankAccount `json:"bankAccount"`
	// Credit card details of the payment method (only present when the payment method is for a credit card)
	CreditCard *CreditCard `json:"creditCard"`
}

func (p PaymentMethodAggregate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PaymentMethodAggregate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PaymentMethodAggregate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PaymentMethodAggregate) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *PaymentMethodAggregate) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *PaymentMethodAggregate) GetUser() *UserAggregate {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *PaymentMethodAggregate) GetGateway() string {
	if o == nil {
		return ""
	}
	return o.Gateway
}

func (o *PaymentMethodAggregate) GetCreated() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Created
}

func (o *PaymentMethodAggregate) GetModified() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Modified
}

func (o *PaymentMethodAggregate) GetBillingAddress() AddressAggregate {
	if o == nil {
		return AddressAggregate{}
	}
	return o.BillingAddress
}

func (o *PaymentMethodAggregate) GetBillingDetails() BillingDetailsAggregate {
	if o == nil {
		return BillingDetailsAggregate{}
	}
	return o.BillingDetails
}

func (o *PaymentMethodAggregate) GetBankAccount() *BankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *PaymentMethodAggregate) GetCreditCard() *CreditCard {
	if o == nil {
		return nil
	}
	return o.CreditCard
}
