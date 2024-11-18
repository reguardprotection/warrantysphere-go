# Distributors
(*Distributors*)

## Overview

### Available Operations

* [ListDistributorsQueryControllerListDistributors](#listdistributorsquerycontrollerlistdistributors) - List Distributors
* [CreateDistributorCommandControllerCreateDistributor](#createdistributorcommandcontrollercreatedistributor) - Create Distributor
* [RetrieveDistributorQueryControllerRetrieveDistributor](#retrievedistributorquerycontrollerretrievedistributor) - Retrieve Distributor
* [UpdateDistributorCommandControllerUpdateDistributor](#updatedistributorcommandcontrollerupdatedistributor) - Update Distributor
* [ArchiveDistributorCommandControllerArchive](#archivedistributorcommandcontrollerarchive) - Archive Distributor
* [DeactivateDistributorCommandControllerDeactivate](#deactivatedistributorcommandcontrollerdeactivate) - Deactivate Distributor
* [ReactivateDistributorCommandControllerArchive](#reactivatedistributorcommandcontrollerarchive) - Reactivate Distributor
* [CreateDistributorOnboardingLinkCommandControllerGenerateDistributorOnboardingLink](#createdistributoronboardinglinkcommandcontrollergeneratedistributoronboardinglink) - Create Link

## ListDistributorsQueryControllerListDistributors

List Distributors

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
    res, err := s.Distributors.ListDistributorsQueryControllerListDistributors(ctx, operations.ListDistributorsQueryControllerListDistributorsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDistributorsQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.ListDistributorsQueryControllerListDistributorsRequest](../../models/operations/listdistributorsquerycontrollerlistdistributorsrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |
| `opts`                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.ListDistributorsQueryControllerListDistributorsResponse](../../models/operations/listdistributorsquerycontrollerlistdistributorsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateDistributorCommandControllerCreateDistributor

Create Distributor

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
    res, err := s.Distributors.CreateDistributorCommandControllerCreateDistributor(ctx, components.CreateDistributorCommandRequest{
        Name: "My Distributor",
        Email: warrantysphere.String("mycompany@mycompany.com"),
        Phone: warrantysphere.String("(647) 426-4565"),
        ParentID: warrantysphere.String("dist_874e0837f66a46d5888f59fdb9e2471b"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDistributorCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [components.CreateDistributorCommandRequest](../../models/components/createdistributorcommandrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateDistributorCommandControllerCreateDistributorResponse](../../models/operations/createdistributorcommandcontrollercreatedistributorresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveDistributorQueryControllerRetrieveDistributor

Retrieve Distributor

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
    res, err := s.Distributors.RetrieveDistributorQueryControllerRetrieveDistributor(ctx, "dist_32cc514683634fc6be690b0e7c17f3f2")
    if err != nil {
        log.Fatal(err)
    }
    if res.DistributorAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `distributorID`                                          | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the distributor.                    | dist_32cc514683634fc6be690b0e7c17f3f2                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrieveDistributorQueryControllerRetrieveDistributorResponse](../../models/operations/retrievedistributorquerycontrollerretrievedistributorresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateDistributorCommandControllerUpdateDistributor

Update Distributor

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
    res, err := s.Distributors.UpdateDistributorCommandControllerUpdateDistributor(ctx, "<id>", components.UpdateDistributorCommandRequest{
        Name: "My distributor",
        Email: warrantysphere.String("mycompany@mycompany.com"),
        Phone: warrantysphere.String("(647) 426-4565"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateDistributorCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `distributorID`                                                                                          | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Unique identifier of the distributor.                                                                    |
| `updateDistributorCommandRequest`                                                                        | [components.UpdateDistributorCommandRequest](../../models/components/updatedistributorcommandrequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateDistributorCommandControllerUpdateDistributorResponse](../../models/operations/updatedistributorcommandcontrollerupdatedistributorresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ArchiveDistributorCommandControllerArchive

Archive Distributor

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
    res, err := s.Distributors.ArchiveDistributorCommandControllerArchive(ctx, "dist_ae2f4c11fef941e9ac625c32c7acff87")
    if err != nil {
        log.Fatal(err)
    }
    if res.ArchiveDistributorCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `distributorID`                                          | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the distributor.                    | dist_ae2f4c11fef941e9ac625c32c7acff87                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ArchiveDistributorCommandControllerArchiveResponse](../../models/operations/archivedistributorcommandcontrollerarchiveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeactivateDistributorCommandControllerDeactivate

Deactivate Distributor

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
    res, err := s.Distributors.DeactivateDistributorCommandControllerDeactivate(ctx, "dist_18696dba81d141c5aee6c50a5d0db60b")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeactivateDistributorCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `distributorID`                                          | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the distributor.                    | dist_18696dba81d141c5aee6c50a5d0db60b                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeactivateDistributorCommandControllerDeactivateResponse](../../models/operations/deactivatedistributorcommandcontrollerdeactivateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ReactivateDistributorCommandControllerArchive

Reactivate Distributor

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
    res, err := s.Distributors.ReactivateDistributorCommandControllerArchive(ctx, "dist_1dbba250f6c4405e9699fbe9d8e48d72")
    if err != nil {
        log.Fatal(err)
    }
    if res.ReactivateDistributorCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `distributorID`                                          | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the distributor.                    | dist_1dbba250f6c4405e9699fbe9d8e48d72                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ReactivateDistributorCommandControllerArchiveResponse](../../models/operations/reactivatedistributorcommandcontrollerarchiveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateDistributorOnboardingLinkCommandControllerGenerateDistributorOnboardingLink

Create Link

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
    res, err := s.Distributors.CreateDistributorOnboardingLinkCommandControllerGenerateDistributorOnboardingLink(ctx, components.CreateDistributorOnboardingLinkCommandRequest{
        Roles: []string{
            "['dist_Viewer']",
        },
        CanCreateUser: true,
        Expiry: 12,
        DistributorID: "dist_5944a84222c34e6ebdbb843b307fe0ac",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DistributorOnboardingLinkAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [components.CreateDistributorOnboardingLinkCommandRequest](../../models/components/createdistributoronboardinglinkcommandrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.CreateDistributorOnboardingLinkCommandControllerGenerateDistributorOnboardingLinkResponse](../../models/operations/createdistributoronboardinglinkcommandcontrollergeneratedistributoronboardinglinkresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |