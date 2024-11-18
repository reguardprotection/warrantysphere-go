# WarrantySphere SDK

## Overview

### Available Operations

* [HealthControllerHealthCheck](#healthcontrollerhealthcheck) - Health Check

## HealthControllerHealthCheck

Health Check

### Example Usage

```go
package main

import(
	"github.com/reguardprotection/warrantysphere"
	"context"
	"log"
)

func main() {
    s := warrantysphere.New()

    ctx := context.Background()
    res, err := s.HealthControllerHealthCheck(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
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

**[*operations.HealthControllerHealthCheckResponse](../../models/operations/healthcontrollerhealthcheckresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |