# ClaimsItems
(*ClaimsItems*)

## Overview

### Available Operations

* [ListClaimItemsQueryControllerList](#listclaimitemsquerycontrollerlist) - List Claim Items
* [CreateClaimItemCommandControllerCreate](#createclaimitemcommandcontrollercreate) - Create Claim Item
* [RetrieveClaimItemQueryControllerRetrieve](#retrieveclaimitemquerycontrollerretrieve) - Retrieve Claim Item
* [UpdateClaimItemCommandControllerUdpate](#updateclaimitemcommandcontrollerudpate) - Update Claim Item
* [DeleteClaimItemCommandControllerDelete](#deleteclaimitemcommandcontrollerdelete) - Delete Claim Item
* [ApproveClaimItemCommandControllerApprove](#approveclaimitemcommandcontrollerapprove) - Approve Claim Item
* [RejectClaimItemCommandControllerReject](#rejectclaimitemcommandcontrollerreject) - Reject Claim Item
* [ResolveClaimItemCommandControllerResolve](#resolveclaimitemcommandcontrollerresolve) - Resolve Claim Item
* [RedraftClaimItemCommandControllerRedraft](#redraftclaimitemcommandcontrollerredraft) - Redraft Claim Item

## ListClaimItemsQueryControllerList

List Claim Items

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
    res, err := s.ClaimsItems.ListClaimItemsQueryControllerList(ctx, operations.ListClaimItemsQueryControllerListRequest{
        CoverageID: warrantysphere.String("cov_0c3bd888961446bd8770bf79b8c7fb05"),
        ClaimID: "clm_85647ba40550480cb903ff44f082f4d6",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClaimItemsQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListClaimItemsQueryControllerListRequest](../../models/operations/listclaimitemsquerycontrollerlistrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListClaimItemsQueryControllerListResponse](../../models/operations/listclaimitemsquerycontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateClaimItemCommandControllerCreate

Create Claim Item

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
    res, err := s.ClaimsItems.CreateClaimItemCommandControllerCreate(ctx, "clm_2cb1cd568683458cb39a9f8c41b72569", components.CreateClaimItemCommandRequest{
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `claimID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim associated with the claim item.                                       | clm_2cb1cd568683458cb39a9f8c41b72569                                                                 |
| `createClaimItemCommandRequest`                                                                      | [components.CreateClaimItemCommandRequest](../../models/components/createclaimitemcommandrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CreateClaimItemCommandControllerCreateResponse](../../models/operations/createclaimitemcommandcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveClaimItemQueryControllerRetrieve

Retrieve Claim Item

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
    res, err := s.ClaimsItems.RetrieveClaimItemQueryControllerRetrieve(ctx, "clmitm_b6d9820b7a874b4c943abc1f7aa27f3f", "clm_593f63daaf46412da362d9ac064a860b")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClaimItemAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `itemID`                                                       | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim item.                           | clmitm_b6d9820b7a874b4c943abc1f7aa27f3f                        |
| `claimID`                                                      | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim associated with the claim item. | clm_593f63daaf46412da362d9ac064a860b                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |

### Response

**[*operations.RetrieveClaimItemQueryControllerRetrieveResponse](../../models/operations/retrieveclaimitemquerycontrollerretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimItemCommandControllerUdpate

Update Claim Item

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
    res, err := s.ClaimsItems.UpdateClaimItemCommandControllerUdpate(ctx, "clmitm_553a4ced4b4b4926b4681c1fad84b609", "clm_1b0fb39a0d4040b693e5d1a62a76db4f", components.UpdateClaimItemCommandRequest{
        CoverageID: warrantysphere.String("cov_61d2f61c88504e0187b3087808bcbef9"),
        Description: warrantysphere.String("Roof problem"),
        Amount: &components.UpdateClaimItemCommandRequestAmount{
            Standard: warrantysphere.Float64(25000),
            Goodwill: warrantysphere.Float64(25000),
            Total: warrantysphere.Float64(50000),
        },
        Cure: &components.UpdateClaimItemCommandRequestCure{
            Title: warrantysphere.String("Water Bucket"),
            Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
            EstimatedCost: warrantysphere.Float64(5000),
        },
        Resolution: &components.UpdateClaimItemCommandRequestResolution{
            Title: warrantysphere.String("Water Bucket"),
            Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
            ActualCost: warrantysphere.Float64(5000),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `itemID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim item.                                                                 | clmitm_553a4ced4b4b4926b4681c1fad84b609                                                              |
| `claimID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim associated with the claim item.                                       | clm_1b0fb39a0d4040b693e5d1a62a76db4f                                                                 |
| `updateClaimItemCommandRequest`                                                                      | [components.UpdateClaimItemCommandRequest](../../models/components/updateclaimitemcommandrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.UpdateClaimItemCommandControllerUdpateResponse](../../models/operations/updateclaimitemcommandcontrollerudpateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteClaimItemCommandControllerDelete

Delete Claim Item

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
    res, err := s.ClaimsItems.DeleteClaimItemCommandControllerDelete(ctx, "I made a typo", "clmitm_e5956287507648eb96f45a553b375d33", "clm_55bdfeb7268c4d71ad253e0170007d48")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `deletedReason`                                                | *string*                                                       | :heavy_check_mark:                                             | Reason given as to why the claim item was deleted.             | I made a typo                                                  |
| `itemID`                                                       | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim item.                           | clmitm_e5956287507648eb96f45a553b375d33                        |
| `claimID`                                                      | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim associated with the claim item. | clm_55bdfeb7268c4d71ad253e0170007d48                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |

### Response

**[*operations.DeleteClaimItemCommandControllerDeleteResponse](../../models/operations/deleteclaimitemcommandcontrollerdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ApproveClaimItemCommandControllerApprove

Approve Claim Item

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
    res, err := s.ClaimsItems.ApproveClaimItemCommandControllerApprove(ctx, "clmitm_d7a13e937b5c4ba5bbb0705de249af23", "clm_a4e840f7986a405e8d27af2f68da137e", components.ApproveClaimItemCommandRequest{
        Reason: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ApproveClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `itemID`                                                                                               | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the claim item.                                                                   | clmitm_d7a13e937b5c4ba5bbb0705de249af23                                                                |
| `claimID`                                                                                              | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the claim associated with the claim item.                                         | clm_a4e840f7986a405e8d27af2f68da137e                                                                   |
| `approveClaimItemCommandRequest`                                                                       | [components.ApproveClaimItemCommandRequest](../../models/components/approveclaimitemcommandrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.ApproveClaimItemCommandControllerApproveResponse](../../models/operations/approveclaimitemcommandcontrollerapproveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RejectClaimItemCommandControllerReject

Reject Claim Item

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
    res, err := s.ClaimsItems.RejectClaimItemCommandControllerReject(ctx, "clmitm_bf51d32fdcbf4f0180c42fba1a2bef50", "clm_11cc7bed707047d2b064a189d24d0670", components.RejectClaimItemCommandRequest{
        Reason: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RejectClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `itemID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim item.                                                                 | clmitm_bf51d32fdcbf4f0180c42fba1a2bef50                                                              |
| `claimID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim associated with the claim item.                                       | clm_11cc7bed707047d2b064a189d24d0670                                                                 |
| `rejectClaimItemCommandRequest`                                                                      | [components.RejectClaimItemCommandRequest](../../models/components/rejectclaimitemcommandrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.RejectClaimItemCommandControllerRejectResponse](../../models/operations/rejectclaimitemcommandcontrollerrejectresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ResolveClaimItemCommandControllerResolve

Resolve Claim Item

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
    res, err := s.ClaimsItems.ResolveClaimItemCommandControllerResolve(ctx, "clmitm_a3114ae3941f4b2b8bf5107168da44a5", "clm_3e34a5e56d8743b1a485b42482de20b4", components.ResolveClaimItemCommandRequest{
        Amount: &components.AmountDto{
            Standard: warrantysphere.Float64(25000),
            Goodwill: warrantysphere.Float64(25000),
            Total: warrantysphere.Float64(50000),
        },
        Resolution: components.ResolutionDto{
            Title: warrantysphere.String("Water Bucket"),
            Description: warrantysphere.String("Water bucket was used to douse the roof and put out the fire."),
            ActualCost: warrantysphere.Float64(5000),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResolveClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `itemID`                                                                                               | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the claim item.                                                                   | clmitm_a3114ae3941f4b2b8bf5107168da44a5                                                                |
| `claimID`                                                                                              | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the claim associated with the claim item.                                         | clm_3e34a5e56d8743b1a485b42482de20b4                                                                   |
| `resolveClaimItemCommandRequest`                                                                       | [components.ResolveClaimItemCommandRequest](../../models/components/resolveclaimitemcommandrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.ResolveClaimItemCommandControllerResolveResponse](../../models/operations/resolveclaimitemcommandcontrollerresolveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RedraftClaimItemCommandControllerRedraft

Redraft Claim Item

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
    res, err := s.ClaimsItems.RedraftClaimItemCommandControllerRedraft(ctx, "clmitm_daa32275199a4cccb645e05aaa623ec9", "clm_2d7f565c331a4d66a36a84f73f478e88", components.RedraftClaimItemCommandRequest{
        Reason: warrantysphere.String("<value>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RedraftClaimItemCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `itemID`                                                                                               | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the claim item.                                                                   | clmitm_daa32275199a4cccb645e05aaa623ec9                                                                |
| `claimID`                                                                                              | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the claim associated with the claim item.                                         | clm_2d7f565c331a4d66a36a84f73f478e88                                                                   |
| `redraftClaimItemCommandRequest`                                                                       | [components.RedraftClaimItemCommandRequest](../../models/components/redraftclaimitemcommandrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.RedraftClaimItemCommandControllerRedraftResponse](../../models/operations/redraftclaimitemcommandcontrollerredraftresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |