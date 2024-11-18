# DeleteClaimItemCommandControllerDeleteRequest


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `DeletedReason`                                                | *string*                                                       | :heavy_check_mark:                                             | Reason given as to why the claim item was deleted.             | I made a typo                                                  |
| `ItemID`                                                       | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim item.                           | clmitm_e5956287507648eb96f45a553b375d33                        |
| `ClaimID`                                                      | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim associated with the claim item. | clm_55bdfeb7268c4d71ad253e0170007d48                           |