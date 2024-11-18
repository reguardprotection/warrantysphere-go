# RetrieveAssetQueryResponse


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                       | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Unique identifier of the asset.                                                                            | ast_3ab14ca0b2c042f28605b299d8a4c79e                                                                       |
| `ReferenceID`                                                                                              | *string*                                                                                                   | :heavy_check_mark:                                                                                         | User-defined reference ID.                                                                                 |                                                                                                            |
| `Name`                                                                                                     | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Display name of the asset.                                                                                 |                                                                                                            |
| `Description`                                                                                              | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Short description of the asset.                                                                            |                                                                                                            |
| `PropertySetID`                                                                                            | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The unique identifier of the property set                                                                  | prpset_3acbb08d61da43e2a0817341f004b31a                                                                    |
| `PropertySet`                                                                                              | [components.PropertySetAggregate](../../models/components/propertysetaggregate.md)                         | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `CustomerID`                                                                                               | *string*                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `Customer`                                                                                                 | [components.CustomerAggregate](../../models/components/customeraggregate.md)                               | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `Submissions`                                                                                              | []*string*                                                                                                 | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `Values`                                                                                                   | [components.RetrieveAssetQueryResponseValues](../../models/components/retrieveassetqueryresponsevalues.md) | :heavy_check_mark:                                                                                         | The values of the most recent asset submission                                                             |                                                                                                            |