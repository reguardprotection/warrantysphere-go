# OrganizationSettings
(*OrganizationSettings*)

## Overview

### Available Operations

* [RetrieveDistributorsSettingsQueryControllerRetrieveDistributor](#retrievedistributorssettingsquerycontrollerretrievedistributor) - Retrieve Organization Settings
* [UpdateDistributorsSettingsQueryControllerRetrieveDistributor](#updatedistributorssettingsquerycontrollerretrievedistributor) - Update Organization Settings

## RetrieveDistributorsSettingsQueryControllerRetrieveDistributor

Retrieve Organization Settings

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
    res, err := s.OrganizationSettings.RetrieveDistributorsSettingsQueryControllerRetrieveDistributor(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveDistributorsSettingsCommandResponse != nil {
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

**[*operations.RetrieveDistributorsSettingsQueryControllerRetrieveDistributorResponse](../../models/operations/retrievedistributorssettingsquerycontrollerretrievedistributorresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateDistributorsSettingsQueryControllerRetrieveDistributor

Update Organization Settings

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
    res, err := s.OrganizationSettings.UpdateDistributorsSettingsQueryControllerRetrieveDistributor(ctx, components.UpdateDistributorsSettingsCommandRequest{
        SelfServeInvoicePackage: components.SelfServePackage{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveDistributorsSettingsCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [components.UpdateDistributorsSettingsCommandRequest](../../models/components/updatedistributorssettingscommandrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdateDistributorsSettingsQueryControllerRetrieveDistributorResponse](../../models/operations/updatedistributorssettingsquerycontrollerretrievedistributorresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |