# UpdatePolicyTermDto


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ID`                                                                                         | **string*                                                                                    | :heavy_minus_sign:                                                                           | Unique identifier of the existing policy term. Leave empty for new policy terms.             |                                                                                              |
| `Duration`                                                                                   | *float64*                                                                                    | :heavy_check_mark:                                                                           | Term duration (in days) of the policy term.                                                  | 426                                                                                          |
| `Title`                                                                                      | *string*                                                                                     | :heavy_check_mark:                                                                           | Friendly name of the policy term (visible in checkout).                                      |                                                                                              |
| `PaymentSchedules`                                                                           | [][components.CreatePaymentScheduleDto](../../models/components/createpaymentscheduledto.md) | :heavy_minus_sign:                                                                           | Payment schedules available on the policy term.                                              |                                                                                              |
| `Order`                                                                                      | *float64*                                                                                    | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |