# UpdatePolicyDocumentCommandRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CanCreateNewVersion`                                                    | *bool*                                                                   | :heavy_check_mark:                                                       | Whether or not the command should create a new draft policy if required. |
| `LegalIntroduction`                                                      | **string*                                                                | :heavy_minus_sign:                                                       | Legal introduction of the document.                                      |
| `WelcomeIntroduction`                                                    | **string*                                                                | :heavy_minus_sign:                                                       | Welcome introduction of the document.                                    |
| `Definitions`                                                            | **string*                                                                | :heavy_minus_sign:                                                       | Definitions of the document.                                             |