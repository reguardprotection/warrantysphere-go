# CreateAPIKeyRequest


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `UserID`                                   | *string*                                   | :heavy_check_mark:                         | N/A                                        | usr_92ab811892594373a0fae275c777acbe       |
| `Name`                                     | **string*                                  | :heavy_minus_sign:                         | N/A                                        |                                            |
| `Expired`                                  | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Datetime the api key is set to expire      | 2024-11-18 15:05:52.164 +0000 UTC          |