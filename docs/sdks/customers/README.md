# Customers
(*Customers*)

## Overview

### Available Operations

* [ListCustomersControllerList](#listcustomerscontrollerlist) - List Customers
* [CreateCustomerControllerCreateCustomer](#createcustomercontrollercreatecustomer) - Create Customer
* [GenerateCustomerSignInLinkControllerGenerateSignIn](#generatecustomersigninlinkcontrollergeneratesignin) - Generate Customer Sign In Link
* [UpdateCustomerControllerUpdateCustomer](#updatecustomercontrollerupdatecustomer) - Update Customer
* [DeleteCustomerControllerDeleteCustomer](#deletecustomercontrollerdeletecustomer) - Delete Customer

## ListCustomersControllerList

List Customers

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
    res, err := s.Customers.ListCustomersControllerList(ctx, operations.ListCustomersControllerListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCustomersResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListCustomersControllerListRequest](../../models/operations/listcustomerscontrollerlistrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListCustomersControllerListResponse](../../models/operations/listcustomerscontrollerlistresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateCustomerControllerCreateCustomer

Create Customer

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
    res, err := s.Customers.CreateCustomerControllerCreateCustomer(ctx, components.CreateCustomerRequest{
        Name: "Victor Solv",
        Email: "vsolv@scalarsolutions.com",
        Phone: "(160) 217-6634",
        Address: &components.CreateCustomerRequestAddress{
            Line1: "100 Holliday St",
            Line2: nil,
            Zip: "21202",
            City: "Baltimore",
            State: "MD",
            Country: "US",
            County: warrantysphere.String("Johnson County"),
        },
        Metadata: warrantysphere.String("{\"Language Preference\":\"English\"}"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCustomerResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateCustomerRequest](../../models/components/createcustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateCustomerControllerCreateCustomerResponse](../../models/operations/createcustomercontrollercreatecustomerresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## GenerateCustomerSignInLinkControllerGenerateSignIn

Generate Customer Sign In Link

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
    res, err := s.Customers.GenerateCustomerSignInLinkControllerGenerateSignIn(ctx, "<id>", components.GenerateSignInLinkCommandRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.GenerateSignInLinkCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `uniqueID`                                                                                                 | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Unique identifier of the customer                                                                          |
| `generateSignInLinkCommandRequest`                                                                         | [components.GenerateSignInLinkCommandRequest](../../models/components/generatesigninlinkcommandrequest.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GenerateCustomerSignInLinkControllerGenerateSignInResponse](../../models/operations/generatecustomersigninlinkcontrollergeneratesigninresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateCustomerControllerUpdateCustomer

Update Customer

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
    res, err := s.Customers.UpdateCustomerControllerUpdateCustomer(ctx, "cus_52328cbccf7645e99b2826eeb22d03cc", components.UpdateCustomerRequest{
        Name: warrantysphere.String("Victor Solv"),
        Email: warrantysphere.String("vsolv@scalarsolutions.com"),
        Phone: warrantysphere.String("(160) 217-6634"),
        Metadata: warrantysphere.String("{\"Language Preference\":\"English\"}"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateCustomerResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `customerID`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | The unique identifier of the customer.                                               | cus_52328cbccf7645e99b2826eeb22d03cc                                                 |
| `updateCustomerRequest`                                                              | [components.UpdateCustomerRequest](../../models/components/updatecustomerrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.UpdateCustomerControllerUpdateCustomerResponse](../../models/operations/updatecustomercontrollerupdatecustomerresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeleteCustomerControllerDeleteCustomer

Delete Customer

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
    res, err := s.Customers.DeleteCustomerControllerDeleteCustomer(ctx, "cus_0dc7262c26024a17bb247f4ba7a8bc32")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCustomerResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `customerID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the customer.                       | cus_0dc7262c26024a17bb247f4ba7a8bc32                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCustomerControllerDeleteCustomerResponse](../../models/operations/deletecustomercontrollerdeletecustomerresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |