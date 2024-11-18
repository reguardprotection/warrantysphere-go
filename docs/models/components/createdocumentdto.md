# CreateDocumentDto


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Title`                                                                   | *string*                                                                  | :heavy_check_mark:                                                        | Title of the document. Either 'Terms & Conditions' or 'Coverage Summary'. |
| `LegalIntroduction`                                                       | **string*                                                                 | :heavy_minus_sign:                                                        | Legal introduction for this document.                                     |
| `WelcomeIntroduction`                                                     | **string*                                                                 | :heavy_minus_sign:                                                        | Welcome introduction for this document.                                   |
| `Definitions`                                                             | **string*                                                                 | :heavy_minus_sign:                                                        | Definitions for this document.                                            |
| `UploadFile`                                                              | **bool*                                                                   | :heavy_minus_sign:                                                        | Document pdf file if a file is to be uploaded.                            |