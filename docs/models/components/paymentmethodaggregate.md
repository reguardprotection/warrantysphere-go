# PaymentMethodAggregate


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                    | *string*                                                                                                | :heavy_check_mark:                                                                                      | Unique ID of the payment method                                                                         | pym_6d199f45d74b4f3eb7b691a3aeac5316                                                                    |
| `ExternalID`                                                                                            | *string*                                                                                                | :heavy_check_mark:                                                                                      | Unique ID of the payment method in its associated gateway                                               |                                                                                                         |
| `UserID`                                                                                                | *string*                                                                                                | :heavy_check_mark:                                                                                      | Unique ID of the payment method's user                                                                  | usr_306974d3d4eb42628c2ba57a14ba4da8                                                                    |
| `User`                                                                                                  | [*components.UserAggregate](../../models/components/useraggregate.md)                                   | :heavy_minus_sign:                                                                                      | User details of the payment method                                                                      |                                                                                                         |
| `Gateway`                                                                                               | *string*                                                                                                | :heavy_check_mark:                                                                                      | Unique ID of the payment method's associated gateway                                                    |                                                                                                         |
| `Created`                                                                                               | [time.Time](https://pkg.go.dev/time#Time)                                                               | :heavy_check_mark:                                                                                      | Datetime when the object was created.                                                                   | 2024-11-18 15:05:48.791 +0000 UTC                                                                       |
| `Modified`                                                                                              | [time.Time](https://pkg.go.dev/time#Time)                                                               | :heavy_check_mark:                                                                                      | Datetime when the object was last modified.                                                             | 2024-11-18 15:05:48.791 +0000 UTC                                                                       |
| `BillingAddress`                                                                                        | [components.AddressAggregate](../../models/components/addressaggregate.md)                              | :heavy_check_mark:                                                                                      | Billing address of the payment method                                                                   |                                                                                                         |
| `BillingDetails`                                                                                        | [components.BillingDetailsAggregate](../../models/components/billingdetailsaggregate.md)                | :heavy_check_mark:                                                                                      | Billing details of the payment method                                                                   |                                                                                                         |
| `BankAccount`                                                                                           | [components.BankAccount](../../models/components/bankaccount.md)                                        | :heavy_check_mark:                                                                                      | Bank account details of the payment method (only present when the payment method is for a bank account) |                                                                                                         |
| `CreditCard`                                                                                            | [components.CreditCard](../../models/components/creditcard.md)                                          | :heavy_check_mark:                                                                                      | Credit card details of the payment method (only present when the payment method is for a credit card)   |                                                                                                         |