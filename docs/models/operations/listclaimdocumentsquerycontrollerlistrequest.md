# ListClaimDocumentsQueryControllerListRequest


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Limit`                                       | **int64*                                      | :heavy_minus_sign:                            | N/A                                           |                                               |
| `Page`                                        | **int64*                                      | :heavy_minus_sign:                            | N/A                                           |                                               |
| `Search`                                      | **string*                                     | :heavy_minus_sign:                            | Value used to search documents by title       |                                               |
| `Feed`                                        | []*string*                                    | :heavy_minus_sign:                            | The feeds of users who can view the document. | [<br/>"staff",<br/>"customer"<br/>]           |
| `ClaimID`                                     | *string*                                      | :heavy_check_mark:                            | The unique identifier of the claim.           | clm_2510484397dd43c08e1a0ff557cb28f5          |