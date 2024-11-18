# Warranties
(*Warranties*)

## Overview

### Available Operations

* [ListWarrantiesControllerListWarranties](#listwarrantiescontrollerlistwarranties) - List Warranties
* [ProvisionWarrantyControllerProvisionWarranty](#provisionwarrantycontrollerprovisionwarranty) - Provision Warranty
* [UpdateWarrantyControllerUpdateWarranty](#updatewarrantycontrollerupdatewarranty) - Update
* [DeleteWarrantyControllerDeleteWarranty](#deletewarrantycontrollerdeletewarranty) - Delete
* [FinalizeWarrantyControllerUpdateWarranty](#finalizewarrantycontrollerupdatewarranty) - Finalize
* [CancelWarrantyControllerCancelWarranty](#cancelwarrantycontrollercancelwarranty) - Cancel Warranty
* [ActivateWarrantyControllerActivateWarranty](#activatewarrantycontrolleractivatewarranty) - Activate Warranty

## ListWarrantiesControllerListWarranties

List Warranties

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
    res, err := s.Warranties.ListWarrantiesControllerListWarranties(ctx, operations.ListWarrantiesControllerListWarrantiesRequest{
        Status: warrantysphere.String("DRAFT,ACTIVATED,CANCELLED,REGISTERED,EXPIRED"),
        CustomerID: warrantysphere.String("cus_b25956d6771c4a8195013f884e46a2ac"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWarrantiesResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.ListWarrantiesControllerListWarrantiesRequest](../../models/operations/listwarrantiescontrollerlistwarrantiesrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.ListWarrantiesControllerListWarrantiesResponse](../../models/operations/listwarrantiescontrollerlistwarrantiesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ProvisionWarrantyControllerProvisionWarranty

Provision Warranty

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
    res, err := s.Warranties.ProvisionWarrantyControllerProvisionWarranty(ctx, components.ProvisionWarrantyRequestBody{
        TermStart: types.MustNewTimeFromString("2024-11-18T15:05:52.186Z"),
        PolicyID: "pol_8af4643df753461ebb3aca5ba24440f3",
        PlanID: "pln_be2c7a23b5f840778b068f050f783378",
        PaymentScheduleID: warrantysphere.String("psc_208641996dc24681a1ef887977702a2b"),
        AddonIds: []string{
            "add_c170a1b01d4a4020963a904d33db84f5",
        },
        DistributorID: warrantysphere.String("dist_ab27878a4e764f09bbefb0143c80ce78"),
        Distributor: &components.DistributorDto{
            Name: "<value>",
            ParentID: warrantysphere.String("dist_3c060b448ec24879b1be2c9ff5179f73"),
        },
        CustomerID: warrantysphere.String("cus_5d3ef41dbbf14fb983a876bb29db94b8"),
        PaymentMethodID: warrantysphere.String("pym_4508c9368dba47bbab1766bc2b8e541e"),
        AssetID: warrantysphere.String("ast_621f08b69db74b3d8c9f76319bbf71e7"),
        PropertyValues: components.ProvisionWarrantyRequestBodyPropertyValues{},
        Fees: []components.WarrantyCoverageFeeDto{
            components.WarrantyCoverageFeeDto{
                PlanCoverageFeeID: "<id>",
                CoverageID: "<id>",
                Title: "Transaction Fee",
                FeeType: components.WarrantyCoverageFeeDtoFeeType{},
            },
        },
        Metadata: &components.ProvisionWarrantyRequestBodyMetadata{},
        ShippingAddress: components.AddressDto{
            Line1: "100 Holliday St",
            Line2: nil,
            Zip: "21202",
            City: "Baltimore",
            State: "MD",
            Country: "US",
            County: warrantysphere.String("Tayside"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProvisionWarrantyResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.ProvisionWarrantyRequestBody](../../models/components/provisionwarrantyrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ProvisionWarrantyControllerProvisionWarrantyResponse](../../models/operations/provisionwarrantycontrollerprovisionwarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateWarrantyControllerUpdateWarranty

Update

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
    res, err := s.Warranties.UpdateWarrantyControllerUpdateWarranty(ctx, "wrt_5d7a0b59aff14571b6eeb562a650d05e", components.UpdateWarrantyCommandRequest{
        TermStart: types.MustNewTimeFromString("2024-11-18T15:05:52.19Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateWarrantyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `warrantyID`                                                                                       | *string*                                                                                           | :heavy_check_mark:                                                                                 | Unique identifier of the warranty.                                                                 | wrt_5d7a0b59aff14571b6eeb562a650d05e                                                               |
| `updateWarrantyCommandRequest`                                                                     | [components.UpdateWarrantyCommandRequest](../../models/components/updatewarrantycommandrequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.UpdateWarrantyControllerUpdateWarrantyResponse](../../models/operations/updatewarrantycontrollerupdatewarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteWarrantyControllerDeleteWarranty

Delete

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
    res, err := s.Warranties.DeleteWarrantyControllerDeleteWarranty(ctx, "wrt_e636d07cf21647a78bfbf971e52ea196")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteWarrantyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `warrantyID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the warranty.                       | wrt_e636d07cf21647a78bfbf971e52ea196                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteWarrantyControllerDeleteWarrantyResponse](../../models/operations/deletewarrantycontrollerdeletewarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## FinalizeWarrantyControllerUpdateWarranty

Finalize

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
    res, err := s.Warranties.FinalizeWarrantyControllerUpdateWarranty(ctx, "wrt_a8bdbbe14091498f8ab58a18d605ee78", components.FinalizeWarrantyCommandRequest{
        TermStart: types.MustNewTimeFromString("2024-11-18T15:05:52.184Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FinalizeWarrantyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `warrantyID`                                                                                           | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the warranty.                                                                     | wrt_a8bdbbe14091498f8ab58a18d605ee78                                                                   |
| `finalizeWarrantyCommandRequest`                                                                       | [components.FinalizeWarrantyCommandRequest](../../models/components/finalizewarrantycommandrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.FinalizeWarrantyControllerUpdateWarrantyResponse](../../models/operations/finalizewarrantycontrollerupdatewarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CancelWarrantyControllerCancelWarranty

Cancel Warranty

### Example Usage

```go
package main

import(
	"os"
	"github.com/reguardprotection/warrantysphere"
	"context"
	"github.com/reguardprotection/warrantysphere/models/operations"
	"github.com/reguardprotection/warrantysphere/types"
	"github.com/reguardprotection/warrantysphere/models/components"
	"log"
)

func main() {
    s := warrantysphere.New(
        warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
    )

    ctx := context.Background()
    res, err := s.Warranties.CancelWarrantyControllerCancelWarranty(ctx, "wrt_fe8834b790d842db87f8d59d1d94c181", operations.QueryParamIDFieldID, components.CancelWarrantyRequestBody{
        CancellationDate: types.MustNewTimeFromString("2024-11-18T15:05:52.182Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelWarrantyResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `uniqueID`                                                                                   | *string*                                                                                     | :heavy_check_mark:                                                                           | Unique ID of the warranty.                                                                   | wrt_fe8834b790d842db87f8d59d1d94c181                                                         |
| `idField`                                                                                    | [operations.QueryParamIDField](../../models/operations/queryparamidfield.md)                 | :heavy_check_mark:                                                                           | Field used to identify the warranty.                                                         |                                                                                              |
| `cancelWarrantyRequestBody`                                                                  | [components.CancelWarrantyRequestBody](../../models/components/cancelwarrantyrequestbody.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.CancelWarrantyControllerCancelWarrantyResponse](../../models/operations/cancelwarrantycontrollercancelwarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ActivateWarrantyControllerActivateWarranty

Activate Warranty

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
    res, err := s.Warranties.ActivateWarrantyControllerActivateWarranty(ctx, "wrt_538726e37a774dffb7f32c6756d5f3e4", components.ActivateWarrantyCommandRequestBody{
        TransactionCloseDate: types.MustTimeFromString("2024-11-18T15:05:52.174Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ActivateWarrantyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `warrantyID`                                                                                                   | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Unique id of the warranty                                                                                      | wrt_538726e37a774dffb7f32c6756d5f3e4                                                                           |
| `activateWarrantyCommandRequestBody`                                                                           | [components.ActivateWarrantyCommandRequestBody](../../models/components/activatewarrantycommandrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.ActivateWarrantyControllerActivateWarrantyResponse](../../models/operations/activatewarrantycontrolleractivatewarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |