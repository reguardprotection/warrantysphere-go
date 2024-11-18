# CreateClaimDocumentBody


## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `File`                                                                                                            | [components.File](../../models/components/file.md)                                                                | :heavy_check_mark:                                                                                                | Binary data of the file to be uploaded.                                                                           |                                                                                                                   |
| `ReferenceID`                                                                                                     | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | User-defined unique identifier to be associated with the document.                                                |                                                                                                                   |
| `Title`                                                                                                           | *string*                                                                                                          | :heavy_check_mark:                                                                                                | Title of the document.                                                                                            |                                                                                                                   |
| `Description`                                                                                                     | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | Description of the document.                                                                                      |                                                                                                                   |
| `Extension`                                                                                                       | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | File extension to be used when storing the file.<br/><br/>If not provided, the file will be stored without any extension. |                                                                                                                   |
| `Feed`                                                                                                            | [][]*string*                                                                                                      | :heavy_minus_sign:                                                                                                | The feeds of users who can view the document.                                                                     | [<br/>"staff",<br/>"customer"<br/>]                                                                               |