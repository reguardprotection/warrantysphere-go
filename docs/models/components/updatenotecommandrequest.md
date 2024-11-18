# UpdateNoteCommandRequest


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Title`                                   | **string*                                 | :heavy_minus_sign:                        | The title of the note.                    |                                           |
| `Content`                                 | **string*                                 | :heavy_minus_sign:                        | The content associated with the note.     |                                           |
| `Feed`                                    | [][]*string*                              | :heavy_minus_sign:                        | The feeds of users who can view the note. | [<br/>"staff",<br/>"customer"<br/>]       |