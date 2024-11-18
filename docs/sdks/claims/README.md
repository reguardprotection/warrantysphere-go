# Claims
(*Claims*)

## Overview

### Available Operations

* [ListClaimsQueryControllerList](#listclaimsquerycontrollerlist) - List Claims
* [OpenClaimCommandControllerOpen](#openclaimcommandcontrolleropen) - Open Claim
* [UpdateClaimCommandControllerUpdate](#updateclaimcommandcontrollerupdate) - Update Claim
* [UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleSteps](#updateclaimlifecyclestepcommandcontrollerupdateclaimlifecyclesteps) - Update Claim Lifecycle Steps
* [ExpireClaimCommandControllerExpire](#expireclaimcommandcontrollerexpire) - Expire Claim
* [CancelClaimCommandControllerCancel](#cancelclaimcommandcontrollercancel) - Cancel Claim
* [CloseClaimCommandControllerClose](#closeclaimcommandcontrollerclose) - Close Claim

## ListClaimsQueryControllerList

List Claims

### Example Usage

```go
package main

import(
	"os"
	"github.com/reguardprotection/warrantysphere"
	"context"
	"github.com/reguardprotection/warrantysphere/models/operations"
	"log"
)

func main() {
    s := warrantysphere.New(
        warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
    )

    ctx := context.Background()
    res, err := s.Claims.ListClaimsQueryControllerList(ctx, operations.ListClaimsQueryControllerListRequest{
        CoverageID: warrantysphere.String("cov_b68af62c1d14489e88c4600b7795aa79"),
        WarrantyID: warrantysphere.String("wrt_97d857c5a5d44b9bb4acf32a7ee15bfa"),
        CustomerID: warrantysphere.String("cus_5f229898e62843f684ce7d6cae4c6609"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClaimsQueryResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListClaimsQueryControllerListRequest](../../models/operations/listclaimsquerycontrollerlistrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListClaimsQueryControllerListResponse](../../models/operations/listclaimsquerycontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## OpenClaimCommandControllerOpen

Open Claim

### Example Usage

```go
package main

import(
	"os"
	"github.com/reguardprotection/warrantysphere"
	"context"
	"github.com/reguardprotection/warrantysphere/types"
	"github.com/reguardprotection/warrantysphere/models/components"
	"log"
)

func main() {
    s := warrantysphere.New(
        warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
    )

    ctx := context.Background()
    res, err := s.Claims.OpenClaimCommandControllerOpen(ctx, components.OpenClaimCommandRequest{
        WarrantyUniqueID: "wrt_e13a9356fb1948a6ab4bd231eba4247b",
        ClaimDate: types.MustNewTimeFromString("2024-11-18T15:05:52.056Z"),
        Complaint: "The roof is broken",
        Amount: &components.AmountDto{
            Standard: warrantysphere.Float64(25000),
            Goodwill: warrantysphere.Float64(25000),
            Total: warrantysphere.Float64(50000),
        },
        CancellationReason: "Customer changed his mind",
        Items: []components.CreateClaimItemCommandRequest{
            components.CreateClaimItemCommandRequest{
                Description: warrantysphere.String("Roof problem"),
                CoverageID: warrantysphere.String("cov_2904ef7369d844b39bcc0f5bef84b3ca"),
                Amount: &components.AmountDto{
                    Standard: warrantysphere.Float64(25000),
                    Goodwill: warrantysphere.Float64(25000),
                    Total: warrantysphere.Float64(50000),
                },
                Cure: &components.CureDto{
                    Title: warrantysphere.String("Water Bucket"),
                    Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
                    EstimatedCost: warrantysphere.Float64(5000),
                },
                Resolution: &components.ResolutionDto{
                    Title: warrantysphere.String("Water Bucket"),
                    Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
                    ActualCost: warrantysphere.Float64(5000),
                },
                AdjudicationReason: warrantysphere.String("Water was delivered to stop the roof from being on fire."),
            },
            components.CreateClaimItemCommandRequest{
                Description: warrantysphere.String("Roof problem"),
                CoverageID: warrantysphere.String("cov_2904ef7369d844b39bcc0f5bef84b3ca"),
                Amount: &components.AmountDto{
                    Standard: warrantysphere.Float64(25000),
                    Goodwill: warrantysphere.Float64(25000),
                    Total: warrantysphere.Float64(50000),
                },
                Cure: &components.CureDto{
                    Title: warrantysphere.String("Water Bucket"),
                    Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
                    EstimatedCost: warrantysphere.Float64(5000),
                },
                Resolution: &components.ResolutionDto{
                    Title: warrantysphere.String("Water Bucket"),
                    Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
                    ActualCost: warrantysphere.Float64(5000),
                },
                AdjudicationReason: warrantysphere.String("Water was delivered to stop the roof from being on fire."),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OpenClaimCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.OpenClaimCommandRequest](../../models/components/openclaimcommandrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.OpenClaimCommandControllerOpenResponse](../../models/operations/openclaimcommandcontrolleropenresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimCommandControllerUpdate

Update Claim

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
    res, err := s.Claims.UpdateClaimCommandControllerUpdate(ctx, "clm_e0805873f4f3403ca29656697aead013", components.UpdateClaimCommandRequest{
        Complaint: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClaimCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `claimID`                                                                                    | *string*                                                                                     | :heavy_check_mark:                                                                           | Unique identifier of the claim associated with the claim item.                               | clm_e0805873f4f3403ca29656697aead013                                                         |
| `updateClaimCommandRequest`                                                                  | [components.UpdateClaimCommandRequest](../../models/components/updateclaimcommandrequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.UpdateClaimCommandControllerUpdateResponse](../../models/operations/updateclaimcommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleSteps

Update Claim Lifecycle Steps

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
    res, err := s.Claims.UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleSteps(ctx, "clm_a45943b9c4b24685b1d6603c8b030e1d", components.UpdateClaimLifecycleStepCommandRequest{
        Steps: []components.UpdateClaimLifecycleStepDtoPUBLIC{
            components.UpdateClaimLifecycleStepDtoPUBLIC{
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: warrantysphere.String("What has caused you to open this claim?"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClaimLifecycleStepCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `claimID`                                                                                                              | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Unique identifier of the claim.                                                                                        | clm_a45943b9c4b24685b1d6603c8b030e1d                                                                                   |
| `updateClaimLifecycleStepCommandRequest`                                                                               | [components.UpdateClaimLifecycleStepCommandRequest](../../models/components/updateclaimlifecyclestepcommandrequest.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleStepsResponse](../../models/operations/updateclaimlifecyclestepcommandcontrollerupdateclaimlifecyclestepsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ExpireClaimCommandControllerExpire

Expire Claim

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
    res, err := s.Claims.ExpireClaimCommandControllerExpire(ctx, "clm_235d5d318b9c4110b0181318a27bd583")
    if err != nil {
        log.Fatal(err)
    }
    if res.ExpireClaimCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `claimID`                                                | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the claim.                          | clm_235d5d318b9c4110b0181318a27bd583                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ExpireClaimCommandControllerExpireResponse](../../models/operations/expireclaimcommandcontrollerexpireresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CancelClaimCommandControllerCancel

Cancel Claim

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
    res, err := s.Claims.CancelClaimCommandControllerCancel(ctx, "clm_b6529b1d37a24b56a4f0044fa01fcecb", components.CancelClaimCommandRequestBody{
        CancellationReason: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelClaimCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `claimID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim                                                                       | clm_b6529b1d37a24b56a4f0044fa01fcecb                                                                 |
| `cancelClaimCommandRequestBody`                                                                      | [components.CancelClaimCommandRequestBody](../../models/components/cancelclaimcommandrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CancelClaimCommandControllerCancelResponse](../../models/operations/cancelclaimcommandcontrollercancelresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CloseClaimCommandControllerClose

Close Claim

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
    res, err := s.Claims.CloseClaimCommandControllerClose(ctx, "clm_79dce970bc004712b59c4c0f138fa9d5")
    if err != nil {
        log.Fatal(err)
    }
    if res.CloseClaimCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `claimID`                                                      | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim associated with the claim item. | clm_79dce970bc004712b59c4c0f138fa9d5                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |

### Response

**[*operations.CloseClaimCommandControllerCloseResponse](../../models/operations/closeclaimcommandcontrollercloseresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |