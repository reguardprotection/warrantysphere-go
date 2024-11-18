# BankAccount

Bank account details of the payment method (only present when the payment method is for a bank account)


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AccountType`                                                    | [components.AccountType](../../models/components/accounttype.md) | :heavy_check_mark:                                               | The type of bank account used for the eCheck.Net transaction     |
| `AccountNumber`                                                  | *string*                                                         | :heavy_check_mark:                                               | The masked bank account number                                   |
| `NameOnAccount`                                                  | *string*                                                         | :heavy_check_mark:                                               | Name of the person who holds the bank account                    |
| `BankName`                                                       | *string*                                                         | :heavy_check_mark:                                               | The name of the bank                                             |