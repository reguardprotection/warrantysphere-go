# Sales
(*Sales*)

## Overview

### Available Operations

* [ListPaymentsQueryControllerListPayments](#listpaymentsquerycontrollerlistpayments) - List Payments
* [ListPaymentsQueryControllerExportPayments](#listpaymentsquerycontrollerexportpayments) - Export Payments (CSV)
* [RetrievePaymentQueryControllerRetrievePayment](#retrievepaymentquerycontrollerretrievepayment) - Retrieve Payment
* [ResendPaymentReminderCommandControllerResendPaymentReminder](#resendpaymentremindercommandcontrollerresendpaymentreminder) - Resend Payment Reminder
* [RetryPaymentCommandControllerRetryPayment](#retrypaymentcommandcontrollerretrypayment) - Retry Payment

## ListPaymentsQueryControllerListPayments

List Payments

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
    res, err := s.Sales.ListPaymentsQueryControllerListPayments(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPaymentsResponse != nil {
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

**[*operations.ListPaymentsQueryControllerListPaymentsResponse](../../models/operations/listpaymentsquerycontrollerlistpaymentsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ListPaymentsQueryControllerExportPayments

Export Payments (CSV)

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
    res, err := s.Sales.ListPaymentsQueryControllerExportPayments(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.ListPaymentsQueryControllerExportPaymentsResponse](../../models/operations/listpaymentsquerycontrollerexportpaymentsresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrievePaymentQueryControllerRetrievePayment

Retrieve Payment

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
    res, err := s.Sales.RetrievePaymentQueryControllerRetrievePayment(ctx, "py_4e4a449e4e5547b680fb65079e30e9ea")
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrievePaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the payment.                    | py_4e4a449e4e5547b680fb65079e30e9ea                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrievePaymentQueryControllerRetrievePaymentResponse](../../models/operations/retrievepaymentquerycontrollerretrievepaymentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ResendPaymentReminderCommandControllerResendPaymentReminder

Resend Payment Reminder

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
    res, err := s.Sales.ResendPaymentReminderCommandControllerResendPaymentReminder(ctx, "py_b1d33c1e07bc49899967025271d0ed50")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResendPaymentReminderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `paymentID`                                              | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the payment.                    | py_b1d33c1e07bc49899967025271d0ed50                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ResendPaymentReminderCommandControllerResendPaymentReminderResponse](../../models/operations/resendpaymentremindercommandcontrollerresendpaymentreminderresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetryPaymentCommandControllerRetryPayment

Retry Payment

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
    res, err := s.Sales.RetryPaymentCommandControllerRetryPayment(ctx, "py_de5eef16ece040d7b285079ecd3de867", components.RetryPaymentRequestBody{
        PaymentMethodID: warrantysphere.String("pym_9a20e26b8a7f4a38a5520f52c98a8620"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetryPaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `paymentID`                                                                              | *string*                                                                                 | :heavy_check_mark:                                                                       | The unique identifier of the payment.                                                    | py_de5eef16ece040d7b285079ecd3de867                                                      |
| `retryPaymentRequestBody`                                                                | [components.RetryPaymentRequestBody](../../models/components/retrypaymentrequestbody.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.RetryPaymentCommandControllerRetryPaymentResponse](../../models/operations/retrypaymentcommandcontrollerretrypaymentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |