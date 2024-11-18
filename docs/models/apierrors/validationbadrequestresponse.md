# ValidationBadRequestResponse


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Message`                                                                                               | []*string*                                                                                              | :heavy_check_mark:                                                                                      | Detailed validation errors                                                                              | [<br/>"\u003cproperty\u003e should not be null or undefined",<br/>"\u003cproperty\u003e should be an integer"<br/>] |
| `StatusCode`                                                                                            | *float64*                                                                                               | :heavy_check_mark:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `Error`                                                                                                 | [apierrors.Error](../../models/apierrors/error.md)                                                      | :heavy_check_mark:                                                                                      | N/A                                                                                                     |                                                                                                         |