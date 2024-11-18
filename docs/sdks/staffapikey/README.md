# StaffAPIKey
(*StaffAPIKey*)

## Overview

### Available Operations

* [ListAPIKeysQueryControllerListKeys](#listapikeysquerycontrollerlistkeys) - List Staff API Keys
* [CreateAPIKeyCommandControllerCreate](#createapikeycommandcontrollercreate) - Create Staff API Key
* [RollAPIKeyCommandControllerRoll](#rollapikeycommandcontrollerroll) - Roll Staff API Key
* [UpdateAPIKeyCommandControllerUpdate](#updateapikeycommandcontrollerupdate) - Update Staff API Key
* [DisableAPIKeyCommandControllerDisable](#disableapikeycommandcontrollerdisable) - Disable Staff API Key

## ListAPIKeysQueryControllerListKeys

List Staff API Keys

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
    res, err := s.StaffAPIKey.ListAPIKeysQueryControllerListKeys(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPIKeysQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListAPIKeysQueryControllerListKeysResponse](../../models/operations/listapikeysquerycontrollerlistkeysresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateAPIKeyCommandControllerCreate

Create Staff API Key

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
    res, err := s.StaffAPIKey.CreateAPIKeyCommandControllerCreate(ctx, components.CreateAPIKeyRequest{
        UserID: "usr_92ab811892594373a0fae275c777acbe",
        Expired: types.MustNewTimeFromString("2024-11-18T15:05:52.164Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAPIKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.CreateAPIKeyRequest](../../models/components/createapikeyrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateAPIKeyCommandControllerCreateResponse](../../models/operations/createapikeycommandcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RollAPIKeyCommandControllerRoll

Roll Staff API Key

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
    res, err := s.StaffAPIKey.RollAPIKeyCommandControllerRoll(ctx, "api_76fe03697b044c66b168876beb7f3923")
    if err != nil {
        log.Fatal(err)
    }
    if res.RollAPIKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The id of the key to roll                                | api_76fe03697b044c66b168876beb7f3923                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RollAPIKeyCommandControllerRollResponse](../../models/operations/rollapikeycommandcontrollerrollresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateAPIKeyCommandControllerUpdate

Update Staff API Key

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
    res, err := s.StaffAPIKey.UpdateAPIKeyCommandControllerUpdate(ctx, "api_6b7e1e84b0204d0d80e3e72039f00eb6", components.UpdateAPIKeyRequest{
        Expired: types.MustNewTimeFromString("2024-11-18T15:05:52.166Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAPIKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `id`                                                                             | *string*                                                                         | :heavy_check_mark:                                                               | The id of the key to update                                                      | api_6b7e1e84b0204d0d80e3e72039f00eb6                                             |
| `updateAPIKeyRequest`                                                            | [components.UpdateAPIKeyRequest](../../models/components/updateapikeyrequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.UpdateAPIKeyCommandControllerUpdateResponse](../../models/operations/updateapikeycommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DisableAPIKeyCommandControllerDisable

Disable Staff API Key

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
    res, err := s.StaffAPIKey.DisableAPIKeyCommandControllerDisable(ctx, "api_3122056849e54f02896a231f2508f8a7")
    if err != nil {
        log.Fatal(err)
    }
    if res.DisableAPIKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The id of the key to disable                             | api_3122056849e54f02896a231f2508f8a7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DisableAPIKeyCommandControllerDisableResponse](../../models/operations/disableapikeycommandcontrollerdisableresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |