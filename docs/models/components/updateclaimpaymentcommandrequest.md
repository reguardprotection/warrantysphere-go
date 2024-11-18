# UpdateClaimPaymentCommandRequest


## Fields

| Field                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Title`                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                 | The title of the payment.                                                                                                                                                                                                                                          |
| `Description`                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                 | Long description of the payment                                                                                                                                                                                                                                    |
| `Destination`                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                 | Unique ID of the payment in its payment method's associated gateway                                                                                                                                                                                                |
| `Amount`                                                                                                                                                                                                                                                           | **int64*                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                 | The monetary amount of the deductible to be paid for any claim.<br/>                  A positive integer representing how much to charge in the smallest currency unit (e.g. 100 cents to charge $1.00).<br/>                  If not 0, the minimum amount is 50 ($0.50). |