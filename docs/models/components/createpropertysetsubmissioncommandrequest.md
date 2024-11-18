# CreatePropertySetSubmissionCommandRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `SetID`                                                                  | *string*                                                                 | :heavy_check_mark:                                                       | The unique identifier of the property set                                | prpset_07a62dce89ea4ba79b03127329673159                                  |
| `PropertyValues`                                                         | [components.PropertyValues](../../models/components/propertyvalues.md)   | :heavy_check_mark:                                                       | The json object containing all the values of the property set properties | {<br/>"square_feet": 2000<br/>}                                          |