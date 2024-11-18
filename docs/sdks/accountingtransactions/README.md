# AccountingTransactions
(*AccountingTransactions*)

## Overview

### Available Operations

* [ListTransactionsControllerListTransactions](#listtransactionscontrollerlisttransactions) - List Transactions
* [CreateTransactionControllerCreateTransaction](#createtransactioncontrollercreatetransaction) - Create Transaction
* [RetrieveTransactionControllerRetrieveTransaction](#retrievetransactioncontrollerretrievetransaction) - Retrieve Transaction
* [VoidTransactionControllerVoidTransaction](#voidtransactioncontrollervoidtransaction) - Void Transaction

## ListTransactionsControllerListTransactions

List Transactions

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
    res, err := s.AccountingTransactions.ListTransactionsControllerListTransactions(ctx, "acc_b73b3097cbaf496b81348c501cade41a", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `accountID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | Unique identifier of the source or destination account of the transactions | acc_b73b3097cbaf496b81348c501cade41a                                       |
| `limit`                                                                    | **int64*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |                                                                            |
| `page`                                                                     | **int64*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.ListTransactionsControllerListTransactionsResponse](../../models/operations/listtransactionscontrollerlisttransactionsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreateTransactionControllerCreateTransaction

Create Transaction

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
    res, err := s.AccountingTransactions.CreateTransactionControllerCreateTransaction(ctx, components.CreateTransactionDto{
        Title: "<value>",
        Amount: 17765,
        SourceAccountID: warrantysphere.String("acc_84ed43ad7a89469c8fda85bf20e8f711"),
        DestinationAccountID: warrantysphere.String("acc_f1a75cf424d546e2bffed5132c39e302"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.CreateTransactionDto](../../models/components/createtransactiondto.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateTransactionControllerCreateTransactionResponse](../../models/operations/createtransactioncontrollercreatetransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrieveTransactionControllerRetrieveTransaction

Retrieve Transaction

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
    res, err := s.AccountingTransactions.RetrieveTransactionControllerRetrieveTransaction(ctx, "txn_8b81e7a785474cee9042be0f09dc9ca1")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the transaction                     | txn_8b81e7a785474cee9042be0f09dc9ca1                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrieveTransactionControllerRetrieveTransactionResponse](../../models/operations/retrievetransactioncontrollerretrievetransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## VoidTransactionControllerVoidTransaction

Void Transaction

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
    res, err := s.AccountingTransactions.VoidTransactionControllerVoidTransaction(ctx, "txn_bd4e0ae4999b4b959b6c0d336eaf72aa")
    if err != nil {
        log.Fatal(err)
    }
    if res.VoidTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the transaction                     | txn_bd4e0ae4999b4b959b6c0d336eaf72aa                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.VoidTransactionControllerVoidTransactionResponse](../../models/operations/voidtransactioncontrollervoidtransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |