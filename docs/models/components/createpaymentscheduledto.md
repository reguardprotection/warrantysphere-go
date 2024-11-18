# CreatePaymentScheduleDto


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ID`                                                      | *string*                                                  | :heavy_check_mark:                                        | Unique identifier of the payment schedule.                |
| `Title`                                                   | *string*                                                  | :heavy_check_mark:                                        | Title of the payment schedule.                            |
| `Rate`                                                    | *float64*                                                 | :heavy_check_mark:                                        | Percentage rate at which each installment is adjusted by. |
| `Installments`                                            | *float64*                                                 | :heavy_check_mark:                                        | Number of installments payment will be made in.           |