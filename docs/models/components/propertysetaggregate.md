# PropertySetAggregate


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `ID`                                        | *string*                                    | :heavy_check_mark:                          | Unique identifier for a property set.       | prpset_2735ecab9685461181a18352e500bcd2     |
| `Name`                                      | *string*                                    | :heavy_check_mark:                          | Display name of the property set.           |                                             |
| `Icon`                                      | *string*                                    | :heavy_check_mark:                          | Icon of the property set.                   |                                             |
| `Properties`                                | []*string*                                  | :heavy_check_mark:                          | List of properties in the property set.     |                                             |
| `Created`                                   | [time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                          | Datetime when the object was created.       | 2024-11-18 15:05:48.705 +0000 UTC           |
| `Modified`                                  | [time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                          | Datetime when the object was last modified. | 2024-11-18 15:05:48.705 +0000 UTC           |