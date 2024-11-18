# Activities
(*Activities*)

## Overview

### Available Operations

* [ListActivitiesQueryControllerList](#listactivitiesquerycontrollerlist) - List Activities

## ListActivitiesQueryControllerList

List Activities

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
    res, err := s.Activities.ListActivitiesQueryControllerList(ctx, []operations.ListActivitiesQueryControllerListQueryParamFeed{
        operations.ListActivitiesQueryControllerListQueryParamFeedCustomer,
        operations.ListActivitiesQueryControllerListQueryParamFeedStaff,
    }, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListActivitiesQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `feed`                                                                                                                                     | [][operations.ListActivitiesQueryControllerListQueryParamFeed](../../models/operations/listactivitiesquerycontrollerlistqueryparamfeed.md) | :heavy_check_mark:                                                                                                                         | The feeds of users who can view the activity.                                                                                              |
| `limit`                                                                                                                                    | **int64*                                                                                                                                   | :heavy_minus_sign:                                                                                                                         | N/A                                                                                                                                        |
| `page`                                                                                                                                     | **int64*                                                                                                                                   | :heavy_minus_sign:                                                                                                                         | N/A                                                                                                                                        |
| `triggeredByID`                                                                                                                            | **string*                                                                                                                                  | :heavy_minus_sign:                                                                                                                         | The unique identifier of the user who triggered the activity (null if system).                                                             |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.ListActivitiesQueryControllerListResponse](../../models/operations/listactivitiesquerycontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |