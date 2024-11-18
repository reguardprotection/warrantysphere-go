# DocumentsProcessorsRuns
(*DocumentsProcessorsRuns*)

## Overview

### Available Operations

* [DocumentProcessorRunsList](#documentprocessorrunslist) - List Processor Runs
* [DocumentProcessorRunsCreate](#documentprocessorrunscreate) - Create Processor Run
* [DocumentProcessorRunsRetrieve](#documentprocessorrunsretrieve) - Retrieve Processor Run
* [DocumentProcessorRunsDelete](#documentprocessorrunsdelete) - Delete Processor Run

## DocumentProcessorRunsList

Returns a paginated list of document processing runs where the `items` array contains up to `limit` runs. The runs are returned sorted by creation date, with the most recent runs appearing first. Passing optional `processorId` or `documentId` parameters will result in filtering the results to runs of the given processor and/or document. Each entry in the array is a separate run object, if no more runs are available, the resulting array will be empty.

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
    res, err := s.DocumentsProcessorsRuns.DocumentProcessorRunsList(ctx, operations.DocumentProcessorRunsListRequest{
        DocumentID: warrantysphere.String("sitm_db6bd84e069841e3ac22779a0135550c"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDocumentProcessorRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DocumentProcessorRunsListRequest](../../models/operations/documentprocessorrunslistrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.DocumentProcessorRunsListResponse](../../models/operations/documentprocessorrunslistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DocumentProcessorRunsCreate

Create/initiate a new run of the specified processor on the given document.

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
    res, err := s.DocumentsProcessorsRuns.DocumentProcessorRunsCreate(ctx, components.CreateDocumentProcessorRunBody{
        ProcessorID: "<id>",
        DocumentID: "sitm_fc65f856b12d45a7afc18aa300d0ecd2",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDocumentProcessorRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.CreateDocumentProcessorRunBody](../../models/components/createdocumentprocessorrunbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.DocumentProcessorRunsCreateResponse](../../models/operations/documentprocessorrunscreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DocumentProcessorRunsRetrieve

Returns the processing run for a valid identifier.

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
    res, err := s.DocumentsProcessorsRuns.DocumentProcessorRunsRetrieve(ctx, "sipr_8807f52854b84ed6a4b5b6ae9206d293", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageItemProcessorResultAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |                                                                         |
| `runID`                                                                 | *string*                                                                | :heavy_check_mark:                                                      | Unique identifier of the document processor run.                        | sipr_8807f52854b84ed6a4b5b6ae9206d293                                   |
| `refreshDownloadURL`                                                    | **bool*                                                                 | :heavy_minus_sign:                                                      | Boolean flag indicating if the run's `downloadURL` should be refreshed. |                                                                         |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |                                                                         |

### Response

**[*operations.DocumentProcessorRunsRetrieveResponse](../../models/operations/documentprocessorrunsretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DocumentProcessorRunsDelete

Permanently deletes a processing run and its related output data. It cannot be undone.

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
    res, err := s.DocumentsProcessorsRuns.DocumentProcessorRunsDelete(ctx, "sipr_80d99491271946c684375a051ca9e850")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteDocumentProcessorRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `runID`                                                  | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the document processor run.         | sipr_80d99491271946c684375a051ca9e850                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DocumentProcessorRunsDeleteResponse](../../models/operations/documentprocessorrunsdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |