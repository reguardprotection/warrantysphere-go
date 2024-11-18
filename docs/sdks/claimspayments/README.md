# ClaimsPayments
(*ClaimsPayments*)

## Overview

### Available Operations

* [ListClaimPaymentsQueryControllerList](#listclaimpaymentsquerycontrollerlist) - List Claim Payments
* [CreateClaimPaymentCommandControllerCreate](#createclaimpaymentcommandcontrollercreate) - Create Claim Payment
* [RetrieveClaimPaymentQueryControllerRetrieve](#retrieveclaimpaymentquerycontrollerretrieve) - Retrieve Claim Payment
* [UpdateClaimPaymentCommandControllerUpdate](#updateclaimpaymentcommandcontrollerupdate) - Update Claim Payment
* [IssueClaimPaymentCommandControllerUpdate](#issueclaimpaymentcommandcontrollerupdate) - Issue Claim Payment
* [VoidClaimPaymentCommandControllerUpdate](#voidclaimpaymentcommandcontrollerupdate) - Void Claim Payment

## ListClaimPaymentsQueryControllerList

List Claim Payments

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
    res, err := s.ClaimsPayments.ListClaimPaymentsQueryControllerList(ctx, operations.ListClaimPaymentsQueryControllerListRequest{
        ClaimID: "clm_7cdaae0579a24740827e287f128d0095",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClaimPaymentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.ListClaimPaymentsQueryControllerListRequest](../../models/operations/listclaimpaymentsquerycontrollerlistrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.ListClaimPaymentsQueryControllerListResponse](../../models/operations/listclaimpaymentsquerycontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateClaimPaymentCommandControllerCreate

Create Claim Payment

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
    res, err := s.ClaimsPayments.CreateClaimPaymentCommandControllerCreate(ctx, "clm_86c2b0e07c2f445f805ea1e46b774407", components.CreateClaimPaymentCommandRequest{
        Amount: 671794,
        CustomerID: "cus_c3cb4d002a8f41ba93d59c6abad4baf7",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClaimPaymentCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `claimID`                                                                                                  | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Unique identifier of the claim associated with the claim payment.                                          | clm_86c2b0e07c2f445f805ea1e46b774407                                                                       |
| `createClaimPaymentCommandRequest`                                                                         | [components.CreateClaimPaymentCommandRequest](../../models/components/createclaimpaymentcommandrequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.CreateClaimPaymentCommandControllerCreateResponse](../../models/operations/createclaimpaymentcommandcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveClaimPaymentQueryControllerRetrieve

Retrieve Claim Payment

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
    res, err := s.ClaimsPayments.RetrieveClaimPaymentQueryControllerRetrieve(ctx, "clm_dff189f468cd462a89581439cf1ac1cb", "clmpy_d5526f62a28f41bb92340a4606b28740")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClaimPaymentAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |                                                                     |
| `claimID`                                                           | *string*                                                            | :heavy_check_mark:                                                  | Unique identifier of the claim used to filter the claim payments.   | clm_dff189f468cd462a89581439cf1ac1cb                                |
| `paymentID`                                                         | *string*                                                            | :heavy_check_mark:                                                  | Unique identifier of the payment used to filter the claim payments. | clmpy_d5526f62a28f41bb92340a4606b28740                              |
| `opts`                                                              | [][operations.Option](../../models/operations/option.md)            | :heavy_minus_sign:                                                  | The options for this request.                                       |                                                                     |

### Response

**[*operations.RetrieveClaimPaymentQueryControllerRetrieveResponse](../../models/operations/retrieveclaimpaymentquerycontrollerretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimPaymentCommandControllerUpdate

Update Claim Payment

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
    res, err := s.ClaimsPayments.UpdateClaimPaymentCommandControllerUpdate(ctx, "clmpy_986e5e4c7aca4f8b9e6d388a0bdbce1f", "clm_199ee14228e04b1d8fddeaf8a69ae53a", components.UpdateClaimPaymentCommandRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClaimPaymentCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `paymentID`                                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Unique ID of the payment                                                                                   | clmpy_986e5e4c7aca4f8b9e6d388a0bdbce1f                                                                     |
| `claimID`                                                                                                  | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Unique identifier of the claim associated with the claim payment.                                          | clm_199ee14228e04b1d8fddeaf8a69ae53a                                                                       |
| `updateClaimPaymentCommandRequest`                                                                         | [components.UpdateClaimPaymentCommandRequest](../../models/components/updateclaimpaymentcommandrequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.UpdateClaimPaymentCommandControllerUpdateResponse](../../models/operations/updateclaimpaymentcommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## IssueClaimPaymentCommandControllerUpdate

Issue Claim Payment

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
    res, err := s.ClaimsPayments.IssueClaimPaymentCommandControllerUpdate(ctx, "clmpy_02c39d48032f4d4091f31db459b941df", "clm_c603911f8c2c49febf4ed67fa3af8786")
    if err != nil {
        log.Fatal(err)
    }
    if res.IssueClaimPaymentCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |                                                                   |
| `paymentID`                                                       | *string*                                                          | :heavy_check_mark:                                                | Unique ID of the payment                                          | clmpy_02c39d48032f4d4091f31db459b941df                            |
| `claimID`                                                         | *string*                                                          | :heavy_check_mark:                                                | Unique identifier of the claim associated with the claim payment. | clm_c603911f8c2c49febf4ed67fa3af8786                              |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |                                                                   |

### Response

**[*operations.IssueClaimPaymentCommandControllerUpdateResponse](../../models/operations/issueclaimpaymentcommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## VoidClaimPaymentCommandControllerUpdate

Void Claim Payment

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
    res, err := s.ClaimsPayments.VoidClaimPaymentCommandControllerUpdate(ctx, "clmpy_097936f1afcc4e5cba8b1c8fe94128c7", "clm_47d1bbc6a16747edb050fde5e511cc3e")
    if err != nil {
        log.Fatal(err)
    }
    if res.VoidClaimPaymentCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |                                                                   |
| `paymentID`                                                       | *string*                                                          | :heavy_check_mark:                                                | Unique ID of the payment                                          | clmpy_097936f1afcc4e5cba8b1c8fe94128c7                            |
| `claimID`                                                         | *string*                                                          | :heavy_check_mark:                                                | Unique identifier of the claim associated with the claim payment. | clm_47d1bbc6a16747edb050fde5e511cc3e                              |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |                                                                   |

### Response

**[*operations.VoidClaimPaymentCommandControllerUpdateResponse](../../models/operations/voidclaimpaymentcommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |