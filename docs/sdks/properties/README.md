# Properties
(*Properties*)

## Overview

### Available Operations

* [RetrievePropertyControllerRetrieve](#retrievepropertycontrollerretrieve) - Retrieve Property
* [ListPropertiesControllerRetrieve](#listpropertiescontrollerretrieve) - List Properties
* [CreateBooleanPropertyControllerCreate](#createbooleanpropertycontrollercreate) - Create Boolean Property
* [CreateDatePropertyControllerCreate](#createdatepropertycontrollercreate) - Create Date Property
* [CreateNumberPropertyControllerCreate](#createnumberpropertycontrollercreate) - Create Number Property
* [CreateObjectPropertyControllerCreate](#createobjectpropertycontrollercreate) - Create Object Property
* [CreateTextPropertyControllerCreate](#createtextpropertycontrollercreate) - Create Text Property

## RetrievePropertyControllerRetrieve

Retrieve Property

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
    res, err := s.Properties.RetrievePropertyControllerRetrieve(ctx, "prop_5ff656ee9b1349289722a9cf13da93e3")
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrievePropertyQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `propertyID`                                             | *string*                                                 | :heavy_check_mark:                                       | The unique id of the property                            | prop_5ff656ee9b1349289722a9cf13da93e3                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrievePropertyControllerRetrieveResponse](../../models/operations/retrievepropertycontrollerretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ListPropertiesControllerRetrieve

List Properties

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
    res, err := s.Properties.ListPropertiesControllerRetrieve(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPropertiesQueryResponse != nil {
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
| `search`                                                 | **string*                                                | :heavy_minus_sign:                                       | Search by name of the property                           |
| `type_`                                                  | [*operations.Type](../../models/operations/type.md)      | :heavy_minus_sign:                                       | Search by type of the property                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListPropertiesControllerRetrieveResponse](../../models/operations/listpropertiescontrollerretrieveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateBooleanPropertyControllerCreate

Create Boolean Property

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
    res, err := s.Properties.CreateBooleanPropertyControllerCreate(ctx, components.CreateBooleanPropertyRequest{
        Name: "My Property",
        ValueKey: "my-property-key",
        FormID: warrantysphere.String("address"),
        Config: components.CreateBooleanPropertyRequestConfig{
            Format: components.CreateBooleanPropertyRequestFormatCheckbox.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePropertyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.CreateBooleanPropertyRequest](../../models/components/createbooleanpropertyrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateBooleanPropertyControllerCreateResponse](../../models/operations/createbooleanpropertycontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateDatePropertyControllerCreate

Create Date Property

### Example Usage

```go
package main

import(
	"os"
	"github.com/reguardprotection/warrantysphere"
	"context"
	"github.com/reguardprotection/warrantysphere/models/components"
	"github.com/reguardprotection/warrantysphere/types"
	"log"
)

func main() {
    s := warrantysphere.New(
        warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
    )

    ctx := context.Background()
    res, err := s.Properties.CreateDatePropertyControllerCreate(ctx, components.CreateDatePropertyRequest{
        Name: "My Property",
        ValueKey: "my-property-key",
        FormID: warrantysphere.String("address"),
        Config: components.DatePropertyConfig{
            Format: components.FormatDatetime.ToPointer(),
            Min: types.MustNewTimeFromString("2024-11-18T15:05:52.153Z"),
            Max: types.MustNewTimeFromString("2024-11-18T15:05:52.153Z"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePropertyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateDatePropertyRequest](../../models/components/createdatepropertyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateDatePropertyControllerCreateResponse](../../models/operations/createdatepropertycontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateNumberPropertyControllerCreate

Create Number Property

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
    res, err := s.Properties.CreateNumberPropertyControllerCreate(ctx, components.CreateNumberPropertyRequest{
        Name: "My Property",
        ValueKey: "my-property-key",
        FormID: warrantysphere.String("address"),
        Config: components.CreateNumberPropertyRequestConfig{
            Format: components.CreateNumberPropertyRequestFormatNumber.ToPointer(),
            Min: warrantysphere.Float64(0),
            Max: warrantysphere.Float64(100),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePropertyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateNumberPropertyRequest](../../models/components/createnumberpropertyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateNumberPropertyControllerCreateResponse](../../models/operations/createnumberpropertycontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateObjectPropertyControllerCreate

Create Object Property

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
    res, err := s.Properties.CreateObjectPropertyControllerCreate(ctx, components.CreateObjectPropertyRequest{
        Name: "My Property",
        ValueKey: "my-property-key",
        FormID: warrantysphere.String("address"),
        Config: components.CreateObjectPropertyRequestConfig{},
        Properties: []components.CreateObjectPropertyAssignmentRequest{
            components.CreateObjectPropertyAssignmentRequest{
                PropertyID: "prop_9c615cb8295d45d3a8d3641b2e9ffa6f",
                Required: false,
                Order: 158.96,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePropertyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateObjectPropertyRequest](../../models/components/createobjectpropertyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateObjectPropertyControllerCreateResponse](../../models/operations/createobjectpropertycontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateTextPropertyControllerCreate

Create Text Property

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
    res, err := s.Properties.CreateTextPropertyControllerCreate(ctx, components.CreateTextPropertyRequest{
        Name: "My Property",
        ValueKey: "my-property-key",
        FormID: warrantysphere.String("address"),
        Config: components.CreateTextPropertyRequestConfig{
            Format: components.CreateTextPropertyRequestFormatEmail.ToPointer(),
            AllowedValues: []string{
                "value1",
                "value2",
            },
            MinLength: warrantysphere.Float64(0),
            MaxLength: warrantysphere.Float64(100),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePropertyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateTextPropertyRequest](../../models/components/createtextpropertyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateTextPropertyControllerCreateResponse](../../models/operations/createtextpropertycontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |