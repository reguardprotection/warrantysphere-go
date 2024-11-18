# APIKeyAggregate


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ID`                                                   | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Created`                                              | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |
| `Modified`                                             | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |
| `Expired`                                              | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |
| `Name`                                                 | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `TenantID`                                             | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Tenant`                                               | [components.Tenant](../../models/components/tenant.md) | :heavy_check_mark:                                     | N/A                                                    |
| `Key`                                                  | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Mode`                                                 | [components.Mode](../../models/components/mode.md)     | :heavy_check_mark:                                     | N/A                                                    |
| `Enabled`                                              | *bool*                                                 | :heavy_check_mark:                                     | N/A                                                    |
| `LastUsed`                                             | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |
| `Details`                                              | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `LastRolled`                                           | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |
| `ServiceAccount`                                       | *bool*                                                 | :heavy_check_mark:                                     | N/A                                                    |