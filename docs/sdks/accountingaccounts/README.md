# AccountingAccounts
(*AccountingAccounts*)

## Overview

### Available Operations

* [ListAccountsControllerListAccounts](#listaccountscontrollerlistaccounts) - List Accounts
* [CreateAccountControllerCreateAccount](#createaccountcontrollercreateaccount) - Create Account
* [RetrieveAccountControllerRetrieveAccount](#retrieveaccountcontrollerretrieveaccount) - Retrieve Account
* [UpdateAccountControllerUpdateAccount](#updateaccountcontrollerupdateaccount) - Update Account

## ListAccountsControllerListAccounts

List Accounts

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
    res, err := s.AccountingAccounts.ListAccountsControllerListAccounts(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAccountsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `limit`                                                              | **int64*                                                             | :heavy_minus_sign:                                                   | N/A                                                                  |
| `page`                                                               | **int64*                                                             | :heavy_minus_sign:                                                   | N/A                                                                  |
| `accountTypes`                                                       | [][operations.AccountTypes](../../models/operations/accounttypes.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.ListAccountsControllerListAccountsResponse](../../models/operations/listaccountscontrollerlistaccountsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateAccountControllerCreateAccount

Create Account

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
    res, err := s.AccountingAccounts.CreateAccountControllerCreateAccount(ctx, components.CreateAccountDto{
        Name: "<value>",
        Description: warrantysphere.String("peninsula apud mostly experienced veto likewise gosh who"),
        Owner: components.CreateCreateAccountDtoOwnerAccountIndividualOwnerDto(
            components.AccountIndividualOwnerDto{
                Type: components.TypeIndividual,
                Name: "<value>",
                Email: "Valentina.Schmeler78@hotmail.com",
                Phone: warrantysphere.String("601-213-0255 x537"),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.CreateAccountDto](../../models/components/createaccountdto.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateAccountControllerCreateAccountResponse](../../models/operations/createaccountcontrollercreateaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveAccountControllerRetrieveAccount

Retrieve Account

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
    res, err := s.AccountingAccounts.RetrieveAccountControllerRetrieveAccount(ctx, "acc_94087ec6370c4a33841c5b43fff3ee52")
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the account                         | acc_94087ec6370c4a33841c5b43fff3ee52                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrieveAccountControllerRetrieveAccountResponse](../../models/operations/retrieveaccountcontrollerretrieveaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateAccountControllerUpdateAccount

Update Account

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
    res, err := s.AccountingAccounts.UpdateAccountControllerUpdateAccount(ctx, "acc_9c1cae58c2b44693945ef10eb1481e9d", components.UpdateAccountDto{
        Owner: components.CreateUpdateAccountDtoOwnerAccountBusinessOwnerDto(
            components.AccountBusinessOwnerDto{
                Type: components.AccountBusinessOwnerDtoTypeBusiness,
                Name: "<value>",
                Email: "Nels40@gmail.com",
                Phone: warrantysphere.String("1-443-545-8419 x43422"),
                BusinessType: warrantysphere.String("<value>"),
                Address: components.AddressDto{
                    Line1: "100 Holliday St",
                    Line2: nil,
                    Zip: "21202",
                    City: "Baltimore",
                    State: "MD",
                    Country: "US",
                    County: nil,
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | Unique identifier of the account                                           | acc_9c1cae58c2b44693945ef10eb1481e9d                                       |
| `updateAccountDto`                                                         | [components.UpdateAccountDto](../../models/components/updateaccountdto.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.UpdateAccountControllerUpdateAccountResponse](../../models/operations/updateaccountcontrollerupdateaccountresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |