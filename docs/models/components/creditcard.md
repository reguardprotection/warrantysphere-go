# CreditCard

Credit card details of the payment method (only present when the payment method is for a credit card)


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CardNumber`                                                      | *string*                                                          | :heavy_check_mark:                                                | The customer's masked credit card number                          |
| `ExpirationDate`                                                  | *string*                                                          | :heavy_check_mark:                                                | The expiration date for the customer's credit card (7 characters) |
| `CardType`                                                        | [components.CardType](../../models/components/cardtype.md)        | :heavy_check_mark:                                                | Type of credit card                                               |