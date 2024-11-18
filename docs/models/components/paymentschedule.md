# PaymentSchedule


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ID`                                               | *string*                                           | :heavy_check_mark:                                 | Unique identifier of the payment schedule          |                                                    |
| `Installments`                                     | *float64*                                          | :heavy_check_mark:                                 | In how many installments is the payment being paid | 1                                                  |
| `Rate`                                             | *float64*                                          | :heavy_check_mark:                                 | Percentage rate                                    |                                                    |
| `Title`                                            | *string*                                           | :heavy_check_mark:                                 | Specify a type of payment                          | One Time Payment                                   |