# ClaimsDocuments
(*ClaimsDocuments*)

## Overview

### Available Operations

* [ListClaimDocumentsQueryControllerList](#listclaimdocumentsquerycontrollerlist) - List Claim Documents
* [CreateClaimDocumentControllerCreate](#createclaimdocumentcontrollercreate) - Create Claim Document
* [RetrieveClaimDocumentQueryControllerRetrieve](#retrieveclaimdocumentquerycontrollerretrieve) - Retrieve Claim Document
* [UpdateClaimDocumentControllerUpdate](#updateclaimdocumentcontrollerupdate) - Update Claim Document
* [DeleteClaimDocumentControllerDelete](#deleteclaimdocumentcontrollerdelete) - Delete Claim Document

## ListClaimDocumentsQueryControllerList

List Claim Documents

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
    res, err := s.ClaimsDocuments.ListClaimDocumentsQueryControllerList(ctx, operations.ListClaimDocumentsQueryControllerListRequest{
        Feed: []string{
            "staff",
            "customer",
        },
        ClaimID: "clm_2510484397dd43c08e1a0ff557cb28f5",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDocumentsQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.ListClaimDocumentsQueryControllerListRequest](../../models/operations/listclaimdocumentsquerycontrollerlistrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.ListClaimDocumentsQueryControllerListResponse](../../models/operations/listclaimdocumentsquerycontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateClaimDocumentControllerCreate

Create Claim Document

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

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.ClaimsDocuments.CreateClaimDocumentControllerCreate(ctx, "clm_9442d597203f436f96ff7f1bc05913b1", components.CreateClaimDocumentBody{
        File: components.File{
            FileName: "example.file",
            Content: content,
        },
        Title: "<value>",
        Feed: [][]string{
            []string{
                "staff",
            },
            []string{
                "customer",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClaimDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `claimID`                                                                                | *string*                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the claim.                                                      | clm_9442d597203f436f96ff7f1bc05913b1                                                     |
| `createClaimDocumentBody`                                                                | [components.CreateClaimDocumentBody](../../models/components/createclaimdocumentbody.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreateClaimDocumentControllerCreateResponse](../../models/operations/createclaimdocumentcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveClaimDocumentQueryControllerRetrieve

Retrieve Claim Document

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
    res, err := s.ClaimsDocuments.RetrieveClaimDocumentQueryControllerRetrieve(ctx, "clm_7b5611a3aac24dfc8da9f287b3504cc9", "doc_eb20cca041084e60afa0713274448de5")
    if err != nil {
        log.Fatal(err)
    }
    if res.DocumentAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `claimID`                                                | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the claim.                      | clm_7b5611a3aac24dfc8da9f287b3504cc9                     |
| `documentID`                                             | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the document.                   | doc_eb20cca041084e60afa0713274448de5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrieveClaimDocumentQueryControllerRetrieveResponse](../../models/operations/retrieveclaimdocumentquerycontrollerretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimDocumentControllerUpdate

Update Claim Document

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
    res, err := s.ClaimsDocuments.UpdateClaimDocumentControllerUpdate(ctx, "clm_15048084b48b42b4938047ea79163d2a", "sitm_123213e8201440c1a82b4b41695ffc07", components.UpdateClaimDocumentBody{
        Feed: [][]string{
            []string{
                "staff",
            },
            []string{
                "customer",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClaimDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `claimID`                                                                                | *string*                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the claim.                                                      | clm_15048084b48b42b4938047ea79163d2a                                                     |
| `documentID`                                                                             | *string*                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the document.                                                   | sitm_123213e8201440c1a82b4b41695ffc07                                                    |
| `updateClaimDocumentBody`                                                                | [components.UpdateClaimDocumentBody](../../models/components/updateclaimdocumentbody.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.UpdateClaimDocumentControllerUpdateResponse](../../models/operations/updateclaimdocumentcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteClaimDocumentControllerDelete

Delete Claim Document

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
    res, err := s.ClaimsDocuments.DeleteClaimDocumentControllerDelete(ctx, "clm_e4323cfac2ee48249f4d74f67dd43e91", "sitm_2570a0fab55c4edf8c3fef82943e521d")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteDocumentCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `claimID`                                                | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the claim.                      | clm_e4323cfac2ee48249f4d74f67dd43e91                     |
| `documentID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the document.                       | sitm_2570a0fab55c4edf8c3fef82943e521d                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteClaimDocumentControllerDeleteResponse](../../models/operations/deleteclaimdocumentcontrollerdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |