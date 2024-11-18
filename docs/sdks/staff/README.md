# Staff
(*Staff*)

## Overview

### Available Operations

* [ListStaffQueryControllerListStaff](#liststaffquerycontrollerliststaff) - List Staff
* [RetrieveStaffQueryControllerRetrieveStaff](#retrievestaffquerycontrollerretrievestaff) - Retrieve Staff
* [RetrieveStaffPermissionsQueryControllerRetrieveStaffPermissions](#retrievestaffpermissionsquerycontrollerretrievestaffpermissions) - Retrieve Staff Roles
* [InviteStaffAndGrantRolesCommandControllerInviteStaffAndGrantRoles](#invitestaffandgrantrolescommandcontrollerinvitestaffandgrantroles) - Invite Staff and Grant Roles
* [UpdateStaffControllerUpdateStaff](#updatestaffcontrollerupdatestaff) - Update Staff
* [CreateRoleAssignmentCommandControllerCreate](#createroleassignmentcommandcontrollercreate) - Grant Staff Role
* [RemoveRoleAssignmentCommandControllerRemove](#removeroleassignmentcommandcontrollerremove) - Remove Staff Role
* [DeactivateStaffControllerDeactivateStaff](#deactivatestaffcontrollerdeactivatestaff) - Deactivate Staff

## ListStaffQueryControllerListStaff

List Staff

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
    res, err := s.Staff.ListStaffQueryControllerListStaff(ctx, operations.ListStaffQueryControllerListStaffRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListStaffQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListStaffQueryControllerListStaffRequest](../../models/operations/liststaffquerycontrollerliststaffrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ListStaffQueryControllerListStaffResponse](../../models/operations/liststaffquerycontrollerliststaffresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveStaffQueryControllerRetrieveStaff

Retrieve Staff

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
    res, err := s.Staff.RetrieveStaffQueryControllerRetrieveStaff(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.StaffAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `uniqueID`                                                | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `idField`                                                 | [*operations.IDField](../../models/operations/idfield.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.RetrieveStaffQueryControllerRetrieveStaffResponse](../../models/operations/retrievestaffquerycontrollerretrievestaffresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveStaffPermissionsQueryControllerRetrieveStaffPermissions

Retrieve Staff Roles

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
    res, err := s.Staff.RetrieveStaffPermissionsQueryControllerRetrieveStaffPermissions(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveStaffPermissionsQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `staffID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveStaffPermissionsQueryControllerRetrieveStaffPermissionsResponse](../../models/operations/retrievestaffpermissionsquerycontrollerretrievestaffpermissionsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## InviteStaffAndGrantRolesCommandControllerInviteStaffAndGrantRoles

Invite Staff and Grant Roles

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
    res, err := s.Staff.InviteStaffAndGrantRolesCommandControllerInviteStaffAndGrantRoles(ctx, components.InviteStaffAndGrantRolesRequest{
        Name: warrantysphere.String("Victor Solv"),
        Email: warrantysphere.String("vsolv@scalarsolutions.com"),
        Phone: warrantysphere.String("(160) 217-6634"),
        Roles: []string{
            "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StaffAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [components.InviteStaffAndGrantRolesRequest](../../models/components/invitestaffandgrantrolesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.InviteStaffAndGrantRolesCommandControllerInviteStaffAndGrantRolesResponse](../../models/operations/invitestaffandgrantrolescommandcontrollerinvitestaffandgrantrolesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateStaffControllerUpdateStaff

Update Staff

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
    res, err := s.Staff.UpdateStaffControllerUpdateStaff(ctx, components.UpdateStaffRequest{
        ID: "stf_17c32a9741f2448b89b16315233b0ba9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateStaffResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.UpdateStaffRequest](../../models/components/updatestaffrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateStaffControllerUpdateStaffResponse](../../models/operations/updatestaffcontrollerupdatestaffresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateRoleAssignmentCommandControllerCreate

Grant Staff Role

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
    res, err := s.Staff.CreateRoleAssignmentCommandControllerCreate(ctx, components.CreateRoleAssignmentRequest{
        StaffID: "stf_a6bd32a7a09247f2b5b187a1aeba3b25",
        RoleID: "rol_213f8a2eb8f743489aea8c03aff84317",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateRoleAssignmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateRoleAssignmentRequest](../../models/components/createroleassignmentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateRoleAssignmentCommandControllerCreateResponse](../../models/operations/createroleassignmentcommandcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RemoveRoleAssignmentCommandControllerRemove

Remove Staff Role

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
    res, err := s.Staff.RemoveRoleAssignmentCommandControllerRemove(ctx, "stf_f20878ba3ffa44198469ea27c2691ca6", "rol_616bdc9948cc43139c0984de2676c582", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveRoleAssignmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `staffID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | stf_f20878ba3ffa44198469ea27c2691ca6                     |
| `roleID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | rol_616bdc9948cc43139c0984de2676c582                     |
| `permissionKey`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RemoveRoleAssignmentCommandControllerRemoveResponse](../../models/operations/removeroleassignmentcommandcontrollerremoveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeactivateStaffControllerDeactivateStaff

Deactivate Staff

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
    res, err := s.Staff.DeactivateStaffControllerDeactivateStaff(ctx, "stf_bdb2c38e62ba47f08589bb2c02071003")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeactivateStaffResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | stf_bdb2c38e62ba47f08589bb2c02071003                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeactivateStaffControllerDeactivateStaffResponse](../../models/operations/deactivatestaffcontrollerdeactivatestaffresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |