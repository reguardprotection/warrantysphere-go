# PortalConfig
(*PortalConfig*)

## Overview

### Available Operations

* [RetrievePortalConfigControllerUpdatePortalConfig](#retrieveportalconfigcontrollerupdateportalconfig) - Retrieve Portal Configuration
* [UpdatePortalConfigControllerUpdatePortalConfig](#updateportalconfigcontrollerupdateportalconfig) - Update Portal Configuration

## RetrievePortalConfigControllerUpdatePortalConfig

Retrieve Portal Configuration

### Example Usage

```go
package main

import(
	"os"
	"github.com/reguardprotection/warrantysphere"
	"context"
	"log"
)

func main() {
    s := warrantysphere.New(
        warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
    )

    ctx := context.Background()
    res, err := s.PortalConfig.RetrievePortalConfigControllerUpdatePortalConfig(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrievePortalConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrievePortalConfigControllerUpdatePortalConfigResponse](../../models/operations/retrieveportalconfigcontrollerupdateportalconfigresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdatePortalConfigControllerUpdatePortalConfig

Update Portal Configuration

### Example Usage

```go
package main

import(
	"os"
	"github.com/reguardprotection/warrantysphere"
	"context"
	"github.com/reguardprotection/warrantysphere/models/components"
	"log"
)

func main() {
    s := warrantysphere.New(
        warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
    )

    ctx := context.Background()
    res, err := s.PortalConfig.UpdatePortalConfigControllerUpdatePortalConfig(ctx, components.UpdatePortalConfigRequest{
        BrandName: "<value>",
        TagLine: "<value>",
        ProductName: "<value>",
        PrColor: "<value>",
        AcColor: "<value>",
        Logo: "<value>",
        Avatar: "https://picsum.photos/seed/VErXvlUez7/745/377",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdatePortalConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.UpdatePortalConfigRequest](../../models/components/updateportalconfigrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdatePortalConfigControllerUpdatePortalConfigResponse](../../models/operations/updateportalconfigcontrollerupdateportalconfigresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |