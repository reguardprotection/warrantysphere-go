# UpdateWebhookResponse


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `DistributorIds`                                                                                         | []*string*                                                                                               | :heavy_minus_sign:                                                                                       | Endpoint will only fire if event contains matching distributor ID(s). Empty array will match all events. |                                                                                                          |
| `URL`                                                                                                    | **string*                                                                                                | :heavy_minus_sign:                                                                                       | URL to update endpoint to point to                                                                       | https://www.website.com                                                                                  |
| `Description`                                                                                            | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Description of endpoint                                                                                  |                                                                                                          |
| `Disabled`                                                                                               | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | If endpoint is disabled                                                                                  | false                                                                                                    |
| `EventTypes`                                                                                             | []*string*                                                                                               | :heavy_minus_sign:                                                                                       | Events to trigger the endpoint. Empty array will trigger all events                                      | [<br/>"user.updated"<br/>]                                                                               |
| `CreatedAt`                                                                                              | [time.Time](https://pkg.go.dev/time#Time)                                                                | :heavy_check_mark:                                                                                       | N/A                                                                                                      | 2024-11-18 15:05:52.193 +0000 UTC                                                                        |
| `ID`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Id of endpoint to update                                                                                 | ep_1234567890                                                                                            |