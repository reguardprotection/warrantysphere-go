# DocumentMetadata


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Filename`                                               | *string*                                                 | :heavy_check_mark:                                       | The original name of the file                            |
| `Extension`                                              | *string*                                                 | :heavy_check_mark:                                       | The type of file                                         |
| `Size`                                                   | *float64*                                                | :heavy_check_mark:                                       | The file size                                            |
| `Expires`                                                | *float64*                                                | :heavy_check_mark:                                       | The expiry date of the download url, converted to number |
| `BucketName`                                             | *string*                                                 | :heavy_check_mark:                                       | The name of the GCP storage bucket                       |
| `IsPublic`                                               | *bool*                                                   | :heavy_check_mark:                                       | The access level of the GCP storage bucket               |
| `StoragePath`                                            | *string*                                                 | :heavy_check_mark:                                       | The path of the stored file in GCP                       |