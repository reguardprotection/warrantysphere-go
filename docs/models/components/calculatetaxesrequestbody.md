# CalculateTaxesRequestBody


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Amount`                                                             | *int64*                                                              | :heavy_check_mark:                                                   | The amount of the transaction in cents. For example, $10.00 is 1000. |
| `Address`                                                            | [components.AddressDto](../../models/components/addressdto.md)       | :heavy_check_mark:                                                   | The address used to calculate taxes.                                 |