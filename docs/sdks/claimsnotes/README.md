# ClaimsNotes
(*ClaimsNotes*)

## Overview

### Available Operations

* [ListClaimNotesQueryControllerList](#listclaimnotesquerycontrollerlist) - List Claim Notes
* [CreateClaimNoteCommandControllerCreate](#createclaimnotecommandcontrollercreate) - Create Claim Note
* [RetrieveClaimNoteQueryControllerRetrieve](#retrieveclaimnotequerycontrollerretrieve) - Retrieve Claim Note
* [UpdateClaimNoteCommandControllerUdpate](#updateclaimnotecommandcontrollerudpate) - Update Claim Note
* [DeleteClaimNoteCommandControllerDelete](#deleteclaimnotecommandcontrollerdelete) - Delete Claim Note

## ListClaimNotesQueryControllerList

List Claim Notes

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
    res, err := s.ClaimsNotes.ListClaimNotesQueryControllerList(ctx, operations.ListClaimNotesQueryControllerListRequest{
        ClaimID: "clm_585de9387a98426484ef27910ae1f42c",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClaimNotesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListClaimNotesQueryControllerListRequest](../../models/operations/listclaimnotesquerycontrollerlistrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListClaimNotesQueryControllerListResponse](../../models/operations/listclaimnotesquerycontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateClaimNoteCommandControllerCreate

Create Claim Note

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
    res, err := s.ClaimsNotes.CreateClaimNoteCommandControllerCreate(ctx, "clm_f8fc6f4447bc4bb3bb2e893433136655", components.CreateClaimNoteCommandRequest{
        Content: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClaimNoteCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `claimID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim associated with the note.                                             | clm_f8fc6f4447bc4bb3bb2e893433136655                                                                 |
| `createClaimNoteCommandRequest`                                                                      | [components.CreateClaimNoteCommandRequest](../../models/components/createclaimnotecommandrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CreateClaimNoteCommandControllerCreateResponse](../../models/operations/createclaimnotecommandcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveClaimNoteQueryControllerRetrieve

Retrieve Claim Note

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
    res, err := s.ClaimsNotes.RetrieveClaimNoteQueryControllerRetrieve(ctx, "clm_d8775afaa6274570a2ba85a2e5449ebc", "clm_adf9f0cab1234ec3b785d2e933f0d9ea")
    if err != nil {
        log.Fatal(err)
    }
    if res.NoteAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `claimID`                                                      | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim used to filter the claim notes. | clm_d8775afaa6274570a2ba85a2e5449ebc                           |
| `noteID`                                                       | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim used to filter the claim notes. | clm_adf9f0cab1234ec3b785d2e933f0d9ea                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |

### Response

**[*operations.RetrieveClaimNoteQueryControllerRetrieveResponse](../../models/operations/retrieveclaimnotequerycontrollerretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimNoteCommandControllerUdpate

Update Claim Note

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
    res, err := s.ClaimsNotes.UpdateClaimNoteCommandControllerUdpate(ctx, "note_23a0e637e29d40ce98a37d0f444bc9dd", "clm_2b54dddecc3e4cc7a8d9fc43a187353d", components.UpdateClaimNoteCommandRequest{
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
    if res.UpdateClaimNoteCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `noteID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the note.                                                                       | note_23a0e637e29d40ce98a37d0f444bc9dd                                                                |
| `claimID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the claim associated with the note.                                             | clm_2b54dddecc3e4cc7a8d9fc43a187353d                                                                 |
| `updateClaimNoteCommandRequest`                                                                      | [components.UpdateClaimNoteCommandRequest](../../models/components/updateclaimnotecommandrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.UpdateClaimNoteCommandControllerUdpateResponse](../../models/operations/updateclaimnotecommandcontrollerudpateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteClaimNoteCommandControllerDelete

Delete Claim Note

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
    res, err := s.ClaimsNotes.DeleteClaimNoteCommandControllerDelete(ctx, "note_0cfcd2cece9f4f4da450cabd6abea441", "clm_84032817236348bdbd8b21cafd526ce4")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteClaimNoteCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `noteID`                                                       | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim note.                           | note_0cfcd2cece9f4f4da450cabd6abea441                          |
| `claimID`                                                      | *string*                                                       | :heavy_check_mark:                                             | Unique identifier of the claim associated with the claim note. | clm_84032817236348bdbd8b21cafd526ce4                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |

### Response

**[*operations.DeleteClaimNoteCommandControllerDeleteResponse](../../models/operations/deleteclaimnotecommandcontrollerdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |