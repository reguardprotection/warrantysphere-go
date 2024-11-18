# PoliciesCoverages
(*PoliciesCoverages*)

## Overview

### Available Operations

* [ListCoveragesQueryControllerListCoverage](#listcoveragesquerycontrollerlistcoverage) - List Policy Coverages
* [CreateCoverageControllerCreateCoverage](#createcoveragecontrollercreatecoverage) - Create Policy Coverage
* [ExportCoveragesQueryControllerExportCoverages](#exportcoveragesquerycontrollerexportcoverages) - Export Policy Coverages
* [RetrieveCoverageQueryControllerRetrieveCoverage](#retrievecoveragequerycontrollerretrievecoverage) - Retrieve Policy Coverage
* [DeleteCoverageControllerDeleteCoverageCommandResponse](#deletecoveragecontrollerdeletecoveragecommandresponse) - Delete Policy Coverage
* [UpdateCoverageCommandControllerUpdate](#updatecoveragecommandcontrollerupdate) - Update Policy Coverage

## ListCoveragesQueryControllerListCoverage

List Policy Coverages

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
    res, err := s.PoliciesCoverages.ListCoveragesQueryControllerListCoverage(ctx, "pol_cd4166aa6df44372b4ee2fd12c163129", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCoveragesQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy.                         | pol_cd4166aa6df44372b4ee2fd12c163129                     |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListCoveragesQueryControllerListCoverageResponse](../../models/operations/listcoveragesquerycontrollerlistcoverageresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateCoverageControllerCreateCoverage

Create Policy Coverage

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
    res, err := s.PoliciesCoverages.CreateCoverageControllerCreateCoverage(ctx, components.CreateCoverageCommandRequest{
        CanCreateNewVersion: true,
        PolicyID: "pol_e8dc8a65d91747f097d2c2beef2a1fc5",
        Title: "<value>",
        Description: "admired quizzically upon",
        Group: "<value>",
        Inclusions: warrantysphere.String("<value>"),
        Exclusions: warrantysphere.String("<value>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCoverageCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.CreateCoverageCommandRequest](../../models/components/createcoveragecommandrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateCoverageControllerCreateCoverageResponse](../../models/operations/createcoveragecontrollercreatecoverageresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ExportCoveragesQueryControllerExportCoverages

Export Policy Coverages

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
    res, err := s.PoliciesCoverages.ExportCoveragesQueryControllerExportCoverages(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExportCoveragesQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `policyID`                                                        | *string*                                                          | :heavy_check_mark:                                                | Unique identifier of ht epolicy whos coverages are being exported |
| `limit`                                                           | **float64*                                                        | :heavy_minus_sign:                                                | Maximum number of records to be exported                          |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.ExportCoveragesQueryControllerExportCoveragesResponse](../../models/operations/exportcoveragesquerycontrollerexportcoveragesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveCoverageQueryControllerRetrieveCoverage

Retrieve Policy Coverage

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
    res, err := s.PoliciesCoverages.RetrieveCoverageQueryControllerRetrieveCoverage(ctx, "cov_6d270da63e334badb6ad96077784381b")
    if err != nil {
        log.Fatal(err)
    }
    if res.CoverageAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `coverageID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the coverage.                       | cov_6d270da63e334badb6ad96077784381b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrieveCoverageQueryControllerRetrieveCoverageResponse](../../models/operations/retrievecoveragequerycontrollerretrievecoverageresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteCoverageControllerDeleteCoverageCommandResponse

Delete Policy Coverage

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
    res, err := s.PoliciesCoverages.DeleteCoverageControllerDeleteCoverageCommandResponse(ctx, "<id>", components.DeleteCoverageCommandRequest{
        CanCreateNewVersion: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCoverageCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `coverageID`                                                                                       | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `deleteCoverageCommandRequest`                                                                     | [components.DeleteCoverageCommandRequest](../../models/components/deletecoveragecommandrequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.DeleteCoverageControllerDeleteCoverageCommandResponseResponse](../../models/operations/deletecoveragecontrollerdeletecoveragecommandresponseresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateCoverageCommandControllerUpdate

Update Policy Coverage

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
    res, err := s.PoliciesCoverages.UpdateCoverageCommandControllerUpdate(ctx, "cov_d90e10180e144afb8a7c021f12ed2685", components.UpdateCoverageCommandRequest{
        CanCreateNewVersion: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateCoverageCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `coverageID`                                                                                       | *string*                                                                                           | :heavy_check_mark:                                                                                 | Unique id of the coverage                                                                          | cov_d90e10180e144afb8a7c021f12ed2685                                                               |
| `updateCoverageCommandRequest`                                                                     | [components.UpdateCoverageCommandRequest](../../models/components/updatecoveragecommandrequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.UpdateCoverageCommandControllerUpdateResponse](../../models/operations/updatecoveragecommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |