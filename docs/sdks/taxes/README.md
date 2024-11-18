# Taxes
(*Taxes*)

## Overview

### Available Operations

* [CalculateTaxesControllerCalculateTaxes](#calculatetaxescontrollercalculatetaxes) - Calculate Taxes

## CalculateTaxesControllerCalculateTaxes

Shows the sales tax that should be collected for a given address and amount.

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
    res, err := s.Taxes.CalculateTaxesControllerCalculateTaxes(ctx, components.CalculateTaxesRequestBody{
        Amount: 449445,
        Address: components.AddressDto{
            Line1: "100 Holliday St",
            Line2: nil,
            Zip: "21202",
            City: "Baltimore",
            State: "MD",
            Country: "US",
            County: warrantysphere.String("Greater Manchester"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalculateTaxesResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CalculateTaxesRequestBody](../../models/components/calculatetaxesrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CalculateTaxesControllerCalculateTaxesResponse](../../models/operations/calculatetaxescontrollercalculatetaxesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |