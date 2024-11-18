# Contract

Warranty contract


## Fields

| Field                                                                                                                                                                                     | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               | Example                                                                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                                                                                      | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Unique identifier of the document.                                                                                                                                                        | sitm_3e7932a1cad24916be0c4291a496d78c                                                                                                                                                     |
| `ReferenceID`                                                                                                                                                                             | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | User-defined unique ID for the document (for reference only).                                                                                                                             |                                                                                                                                                                                           |
| `Created`                                                                                                                                                                                 | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                        | Datetime when the object was created.                                                                                                                                                     | 2024-11-18 15:05:47.525 +0000 UTC                                                                                                                                                         |
| `CreatedByID`                                                                                                                                                                             | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Unique ID of the user who created/uploaded this document.                                                                                                                                 | usr_f3df21c8fa9b4176bc1601f61393caa3                                                                                                                                                      |
| `CreatedBy`                                                                                                                                                                               | [*components.WarrantyAggregateCreatedBy](../../models/components/warrantyaggregatecreatedby.md)                                                                                           | :heavy_minus_sign:                                                                                                                                                                        | Data of the user who created/uploaded this document.                                                                                                                                      |                                                                                                                                                                                           |
| `Modified`                                                                                                                                                                                | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                        | Datetime when the object was last modified.                                                                                                                                               | 2024-11-18 15:05:47.525 +0000 UTC                                                                                                                                                         |
| `ModifiedByID`                                                                                                                                                                            | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Unique ID of the user who last modified this document (title, description, etc.).                                                                                                         | usr_0f6a9d822ca54a92a084d28afda1011d                                                                                                                                                      |
| `ModifiedBy`                                                                                                                                                                              | [*components.WarrantyAggregateModifiedBy](../../models/components/warrantyaggregatemodifiedby.md)                                                                                         | :heavy_minus_sign:                                                                                                                                                                        | Data of the user who last modified this document (title, description, etc.).                                                                                                              |                                                                                                                                                                                           |
| `Deleted`                                                                                                                                                                                 | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                        | Datetime when the object was deleted.                                                                                                                                                     | 2024-11-18 15:05:47.525 +0000 UTC                                                                                                                                                         |
| `DeletedByID`                                                                                                                                                                             | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Unique ID of the user who deleted this document.                                                                                                                                          | usr_befa289e6bd44359ad17706c435e32ca                                                                                                                                                      |
| `DeletedBy`                                                                                                                                                                               | [*components.WarrantyAggregateDeletedBy](../../models/components/warrantyaggregatedeletedby.md)                                                                                           | :heavy_minus_sign:                                                                                                                                                                        | Data of the user who deleted this document.                                                                                                                                               |                                                                                                                                                                                           |
| `Title`                                                                                                                                                                                   | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Title of the document.                                                                                                                                                                    |                                                                                                                                                                                           |
| `Description`                                                                                                                                                                             | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Short text description of the document.                                                                                                                                                   |                                                                                                                                                                                           |
| `MimeType`                                                                                                                                                                                | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Mime type of the document                                                                                                                                                                 | application/json                                                                                                                                                                          |
| `Metadata`                                                                                                                                                                                | [components.WarrantyAggregateMetadata](../../models/components/warrantyaggregatemetadata.md)                                                                                              | :heavy_check_mark:                                                                                                                                                                        | JSON object storing abitrary user-defined metadata.<br/><br/>Some elements may be automatically populated by the system or various document processors.                                   | {<br/>"externalId": 123<br/>}                                                                                                                                                             |
| `Extension`                                                                                                                                                                               | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | N/A                                                                                                                                                                                       | pdf                                                                                                                                                                                       |
| `OriginalFilename`                                                                                                                                                                        | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | N/A                                                                                                                                                                                       |                                                                                                                                                                                           |
| `PrivateFile`                                                                                                                                                                             | **bool*                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                        | Boolean flag indicating if a file may or may not be accessed with a public URL. Private files can only be accessed when their downloadURL includes an authentication token (short-lived). |                                                                                                                                                                                           |
| `ExternalID`                                                                                                                                                                              | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | System-defined unique identifier.                                                                                                                                                         |                                                                                                                                                                                           |
| `DownloadURL`                                                                                                                                                                             | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | Download URL of the document. Private document downloadURLs include a short-lived authentication token.                                                                                   |                                                                                                                                                                                           |
| `ProcessingResults`                                                                                                                                                                       | [][components.StorageItemProcessorResultAggregate](../../models/components/storageitemprocessorresultaggregate.md)                                                                        | :heavy_minus_sign:                                                                                                                                                                        | List of processing results for this document (when applicable).                                                                                                                           |                                                                                                                                                                                           |
| `Feed`                                                                                                                                                                                    | [][components.WarrantyAggregateFeed](../../models/components/warrantyaggregatefeed.md)                                                                                                    | :heavy_check_mark:                                                                                                                                                                        | The feeds of users who can view the document                                                                                                                                              |                                                                                                                                                                                           |