# StatusChangeAggregate


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `From`                                             | [components.From](../../models/components/from.md) | :heavy_check_mark:                                 | Status of the payment before the change            |
| `To`                                               | [components.To](../../models/components/to.md)     | :heavy_check_mark:                                 | Status of the payment after the change             |
| `Date`                                             | [time.Time](https://pkg.go.dev/time#Time)          | :heavy_check_mark:                                 | Datetime when the status change occured            |