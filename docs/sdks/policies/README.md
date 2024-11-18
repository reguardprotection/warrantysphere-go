# Policies
(*Policies*)

## Overview

### Available Operations

* [ListPoliciesQueryControllerListPolicies](#listpoliciesquerycontrollerlistpolicies) - List Policies
* [CreatePolicyCommandControllerCreatePolicy](#createpolicycommandcontrollercreatepolicy) - Create Policy
* [ExportPoliciesQueryControllerExportPolicies](#exportpoliciesquerycontrollerexportpolicies) - Export Policies
* [RetrievePolicyQueryControllerRetrievePolicy](#retrievepolicyquerycontrollerretrievepolicy) - Retrieve Policy
* [UpdatePolicyCommandControllerUpdatePolicy](#updatepolicycommandcontrollerupdatepolicy) - Update Policy
* [DeletePolicyCommandControllerDelete](#deletepolicycommandcontrollerdelete) - Delete Policy
* [UpdatePolicyDocumentCommandControllerUpdatePolicyDocument](#updatepolicydocumentcommandcontrollerupdatepolicydocument) - Update Policy Document
* [UpdatePolicyEmailCommandControllerUpdate](#updatepolicyemailcommandcontrollerupdate) - Update Policy Email
* [UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleStep](#updateclaimlifecyclestepcommandcontrollerupdateclaimlifecyclestep) - Update Policy Claim Lifecycle Steps
* [PublishPolicyCommandControllerPublish](#publishpolicycommandcontrollerpublish) - Publish Policy
* [ArchivePolicyCommandControllerArchive](#archivepolicycommandcontrollerarchive) - Archive Policy
* [DuplicatePolicyCommandControllerDuplicate](#duplicatepolicycommandcontrollerduplicate) - Duplicate
* [MigratePolicyFromSandboxCommandControllerMigrateSandbox](#migratepolicyfromsandboxcommandcontrollermigratesandbox) - Migrate Sandbox Policy to Live
* [UpdatePolicyCancelationRuleCommandControllerUpdate](#updatepolicycancelationrulecommandcontrollerupdate) - Update Cancelation Rule
* [DeletePolicyCancelationRuleCommandControllerUpdate](#deletepolicycancelationrulecommandcontrollerupdate) - Delete Cancelation Rule
* [CreatePolicyCancelationRuleCommandControllerCreate](#createpolicycancelationrulecommandcontrollercreate) - Create Cancelation Rule
* [PolicyQuoteControllerProvisionWarranty](#policyquotecontrollerprovisionwarranty) - Quote

## ListPoliciesQueryControllerListPolicies

List Policies

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
    res, err := s.Policies.ListPoliciesQueryControllerListPolicies(ctx, operations.ListPoliciesQueryControllerListPoliciesRequest{
        Search: warrantysphere.String("100 Day Inspection"),
        PolicyIds: warrantysphere.String("pol_a9dd3c31cb70434189a93fc3b5aad30b"),
        PropertySetIds: warrantysphere.String("prpset_692f3b83dd394d118ef8d3bc85800f04"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPoliciesQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.ListPoliciesQueryControllerListPoliciesRequest](../../models/operations/listpoliciesquerycontrollerlistpoliciesrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.ListPoliciesQueryControllerListPoliciesResponse](../../models/operations/listpoliciesquerycontrollerlistpoliciesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreatePolicyCommandControllerCreatePolicy

Create Policy

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
    res, err := s.Policies.CreatePolicyCommandControllerCreatePolicy(ctx, components.CreatePolicyCommandRequest{
        Title: "My Policy",
        Icon: "file-shield-02",
        Tagline: "<value>",
        PolicyNumber: "HGP-20249453",
        ClaimWaitingPeriod: warrantysphere.Float64(30),
        PropertySetID: "prpset_4536e5cfe0904730a26dfc30a392573d",
        ClaimLifecycleSteps: []components.ClaimLifecycleStepDto{
            components.ClaimLifecycleStepDto{
                ID: warrantysphere.String("cls_ccb73f5636c04add8e91c3929353f94b"),
                Order: warrantysphere.Float64(0),
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: "What has caused you to open this claim?",
            },
            components.ClaimLifecycleStepDto{
                ID: warrantysphere.String("cls_ccb73f5636c04add8e91c3929353f94b"),
                Order: warrantysphere.Float64(0),
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: "What has caused you to open this claim?",
            },
            components.ClaimLifecycleStepDto{
                ID: warrantysphere.String("cls_ccb73f5636c04add8e91c3929353f94b"),
                Order: warrantysphere.Float64(0),
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: "What has caused you to open this claim?",
            },
        },
        Coverages: []components.CreateCoverageDto{
            components.CreateCoverageDto{
                Order: 0,
                Title: "Electric and appliances",
                Description: "This covers your electrical service panels, and various appliances in your home.",
                Group: "Plumbing",
                ReferenceID: warrantysphere.String("<id>"),
            },
            components.CreateCoverageDto{
                Order: 0,
                Title: "Electric and appliances",
                Description: "This covers your electrical service panels, and various appliances in your home.",
                Group: "Plumbing",
                ReferenceID: warrantysphere.String("<id>"),
            },
        },
        TermsAndConditions: components.CreateDocumentDto{
            Title: "<value>",
        },
        CoverageSummary: components.CreateDocumentDto{
            Title: "<value>",
        },
        Terms: []components.CreatePolicyTermDto{
            components.CreatePolicyTermDto{
                Duration: 426,
                Title: "<value>",
                Order: 2484.58,
            },
            components.CreatePolicyTermDto{
                Duration: 426,
                Title: "<value>",
                Order: 1209.36,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePolicyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.CreatePolicyCommandRequest](../../models/components/createpolicycommandrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreatePolicyCommandControllerCreatePolicyResponse](../../models/operations/createpolicycommandcontrollercreatepolicyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ExportPoliciesQueryControllerExportPolicies

Export Policies

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
    res, err := s.Policies.ExportPoliciesQueryControllerExportPolicies(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExportPoliciesQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `limit`                                                  | **float64*                                               | :heavy_minus_sign:                                       | Maximum number of records to be exported                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ExportPoliciesQueryControllerExportPoliciesResponse](../../models/operations/exportpoliciesquerycontrollerexportpoliciesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrievePolicyQueryControllerRetrievePolicy

Retrieve Policy

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
    res, err := s.Policies.RetrievePolicyQueryControllerRetrievePolicy(ctx, "pol_00cf0518737942c6946995372a198b3a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy.                         | pol_00cf0518737942c6946995372a198b3a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrievePolicyQueryControllerRetrievePolicyResponse](../../models/operations/retrievepolicyquerycontrollerretrievepolicyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdatePolicyCommandControllerUpdatePolicy

Update Policy

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
    res, err := s.Policies.UpdatePolicyCommandControllerUpdatePolicy(ctx, "pol_cf20c82919da426493d48fc4ac4aa2ca", components.UpdatePolicyCommandRequest{
        CanCreateNewVersion: true,
        Terms: []components.UpdatePolicyTermDto{
            components.UpdatePolicyTermDto{
                Duration: 426,
                Title: "<value>",
                Order: 6751.61,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdatePolicyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `policyID`                                                                                     | *string*                                                                                       | :heavy_check_mark:                                                                             | Unique identifier of the policy.                                                               | pol_cf20c82919da426493d48fc4ac4aa2ca                                                           |
| `updatePolicyCommandRequest`                                                                   | [components.UpdatePolicyCommandRequest](../../models/components/updatepolicycommandrequest.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.UpdatePolicyCommandControllerUpdatePolicyResponse](../../models/operations/updatepolicycommandcontrollerupdatepolicyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeletePolicyCommandControllerDelete

Delete Policy

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
    res, err := s.Policies.DeletePolicyCommandControllerDelete(ctx, "pol_5ed72337bf6348fb8f262cc10a97223b")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeletePolicyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy.                         | pol_5ed72337bf6348fb8f262cc10a97223b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePolicyCommandControllerDeleteResponse](../../models/operations/deletepolicycommandcontrollerdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdatePolicyDocumentCommandControllerUpdatePolicyDocument

Update Policy Document

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
    res, err := s.Policies.UpdatePolicyDocumentCommandControllerUpdatePolicyDocument(ctx, "poldoc_2004a2e541d64513b92dd01be9b4c388", "pol_a95f2c5a90214df2acacbaf2b5c00d04", components.UpdatePolicyDocumentCommandRequest{
        CanCreateNewVersion: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdatePolicyDocumentCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `documentID`                                                                                                   | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Unique identifier of the document.                                                                             | poldoc_2004a2e541d64513b92dd01be9b4c388                                                                        |
| `policyID`                                                                                                     | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Unique identifier of the policy.                                                                               | pol_a95f2c5a90214df2acacbaf2b5c00d04                                                                           |
| `updatePolicyDocumentCommandRequest`                                                                           | [components.UpdatePolicyDocumentCommandRequest](../../models/components/updatepolicydocumentcommandrequest.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentResponse](../../models/operations/updatepolicydocumentcommandcontrollerupdatepolicydocumentresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdatePolicyEmailCommandControllerUpdate

Update Policy Email

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
    res, err := s.Policies.UpdatePolicyEmailCommandControllerUpdate(ctx, "polemail_26be6c1c2c4c4f898de4cd7942012fdd", "pol_cdbcff5e43bc4b2b85d18209086cd9d5", components.UpdatePolicyEmailCommandRequest{
        Type: components.UpdatePolicyEmailCommandRequestTypeWarrantyProvisioned,
        Subject: warrantysphere.String("<value>"),
        Content: warrantysphere.String("<value>"),
        Byline: warrantysphere.String("<value>"),
        InstructionImage: warrantysphere.String("<value>"),
        Header: warrantysphere.String("<value>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdatePolicyEmailCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `emailID`                                                                                                | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Unique identifier of the email.                                                                          | polemail_26be6c1c2c4c4f898de4cd7942012fdd                                                                |
| `policyID`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Unique identifier of the policy.                                                                         | pol_cdbcff5e43bc4b2b85d18209086cd9d5                                                                     |
| `updatePolicyEmailCommandRequest`                                                                        | [components.UpdatePolicyEmailCommandRequest](../../models/components/updatepolicyemailcommandrequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.UpdatePolicyEmailCommandControllerUpdateResponse](../../models/operations/updatepolicyemailcommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleStep

Update Policy Claim Lifecycle Steps

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
    res, err := s.Policies.UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleStep(ctx, "pol_884729829e484f9388085cd9a7e6b72a", components.UpdateClaimLifecycleStepCommandRequest{
        Steps: []components.UpdateClaimLifecycleStepDtoPUBLIC{
            components.UpdateClaimLifecycleStepDtoPUBLIC{
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: warrantysphere.String("What has caused you to open this claim?"),
            },
            components.UpdateClaimLifecycleStepDtoPUBLIC{
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: warrantysphere.String("What has caused you to open this claim?"),
            },
            components.UpdateClaimLifecycleStepDtoPUBLIC{
                Icon: "info-circle",
                Title: "Describe your issue",
                Description: warrantysphere.String("What has caused you to open this claim?"),
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClaimLifecycleStepCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `policyID`                                                                                                             | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Unique identifier of the policy.                                                                                       | pol_884729829e484f9388085cd9a7e6b72a                                                                                   |
| `updateClaimLifecycleStepCommandRequest`                                                                               | [components.UpdateClaimLifecycleStepCommandRequest](../../models/components/updateclaimlifecyclestepcommandrequest.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `xWspherePermissionKeys`                                                                                               | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | List of permission keys                                                                                                |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleStepResponse](../../models/operations/updateclaimlifecyclestepcommandcontrollerupdateclaimlifecyclestepresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## PublishPolicyCommandControllerPublish

Publish Policy

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
    res, err := s.Policies.PublishPolicyCommandControllerPublish(ctx, "pol_e712290b8b7744f59ad7ac14da3490b5")
    if err != nil {
        log.Fatal(err)
    }
    if res.PublishPolicyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the policy.                     | pol_e712290b8b7744f59ad7ac14da3490b5                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PublishPolicyCommandControllerPublishResponse](../../models/operations/publishpolicycommandcontrollerpublishresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ArchivePolicyCommandControllerArchive

Archive Policy

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
    res, err := s.Policies.ArchivePolicyCommandControllerArchive(ctx, "pol_a1a9da00d340424d9fff6c8063ff73c8")
    if err != nil {
        log.Fatal(err)
    }
    if res.ArchivePolicyCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy.                         | pol_a1a9da00d340424d9fff6c8063ff73c8                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ArchivePolicyCommandControllerArchiveResponse](../../models/operations/archivepolicycommandcontrollerarchiveresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DuplicatePolicyCommandControllerDuplicate

Duplicate

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
    res, err := s.Policies.DuplicatePolicyCommandControllerDuplicate(ctx, "pol_27081af6055d491cbf54b3d3790f835a", components.DuplicatePolicyCommandRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `policyID`                                                                                           | *string*                                                                                             | :heavy_check_mark:                                                                                   | Unique identifier of the policy the new policy will be created from.                                 | pol_27081af6055d491cbf54b3d3790f835a                                                                 |
| `duplicatePolicyCommandRequest`                                                                      | [components.DuplicatePolicyCommandRequest](../../models/components/duplicatepolicycommandrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.DuplicatePolicyCommandControllerDuplicateResponse](../../models/operations/duplicatepolicycommandcontrollerduplicateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## MigratePolicyFromSandboxCommandControllerMigrateSandbox

Migrate Sandbox Policy to Live

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
    res, err := s.Policies.MigratePolicyFromSandboxCommandControllerMigrateSandbox(ctx, "pol_51e5f39f81bc46659a810cc2111c8204")
    if err != nil {
        log.Fatal(err)
    }
    if res.MigratePolicyFromSandboxToLiveCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the sandbox policy.                 | pol_51e5f39f81bc46659a810cc2111c8204                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.MigratePolicyFromSandboxCommandControllerMigrateSandboxResponse](../../models/operations/migratepolicyfromsandboxcommandcontrollermigratesandboxresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdatePolicyCancelationRuleCommandControllerUpdate

Update Cancelation Rule

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
    res, err := s.Policies.UpdatePolicyCancelationRuleCommandControllerUpdate(ctx, "pol_b3182dd757b14e05a005a02c892216e3", "US", "CA", components.UpdatePolicyCancelationRuleCommandRequest{
        AdminFee: warrantysphere.Float64(0),
        DeductClaims: true,
        RuleText: warrantysphere.String("In the state of California, the cancelation fee cannot be charged if the policy is not..."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdatePolicyCancelationRuleCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |                                                                                                                              |
| `policyID`                                                                                                                   | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | Unique identifier of the policy.                                                                                             | pol_b3182dd757b14e05a005a02c892216e3                                                                                         |
| `country`                                                                                                                    | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The country of the policy.                                                                                                   | US                                                                                                                           |
| `state`                                                                                                                      | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The state of the policy.                                                                                                     | CA                                                                                                                           |
| `updatePolicyCancelationRuleCommandRequest`                                                                                  | [components.UpdatePolicyCancelationRuleCommandRequest](../../models/components/updatepolicycancelationrulecommandrequest.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |                                                                                                                              |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |                                                                                                                              |

### Response

**[*operations.UpdatePolicyCancelationRuleCommandControllerUpdateResponse](../../models/operations/updatepolicycancelationrulecommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeletePolicyCancelationRuleCommandControllerUpdate

Delete Cancelation Rule

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
    res, err := s.Policies.DeletePolicyCancelationRuleCommandControllerUpdate(ctx, "pol_73a1aa6a696f4a589b3696eb6263c580", "US", "CA")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeletePolicyCancelationRuleCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy.                         | pol_73a1aa6a696f4a589b3696eb6263c580                     |
| `country`                                                | *string*                                                 | :heavy_check_mark:                                       | The country of the policy.                               | US                                                       |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | The state of the policy.                                 | CA                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePolicyCancelationRuleCommandControllerUpdateResponse](../../models/operations/deletepolicycancelationrulecommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreatePolicyCancelationRuleCommandControllerCreate

Create Cancelation Rule

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
    res, err := s.Policies.CreatePolicyCancelationRuleCommandControllerCreate(ctx, "<id>", components.CreatePolicyCancelationRuleCommandRequest{
        Country: "US",
        State: "CA",
        AdminFee: warrantysphere.Float64(0),
        DeductClaims: true,
        RuleText: warrantysphere.String("In the state of California, the cancelation fee cannot be charged if the policy is not..."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePolicyCancelationRuleCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `policyID`                                                                                                                   | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The unique identifier of the policy.                                                                                         |
| `createPolicyCancelationRuleCommandRequest`                                                                                  | [components.CreatePolicyCancelationRuleCommandRequest](../../models/components/createpolicycancelationrulecommandrequest.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CreatePolicyCancelationRuleCommandControllerCreateResponse](../../models/operations/createpolicycancelationrulecommandcontrollercreateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## PolicyQuoteControllerProvisionWarranty

Quote

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
    res, err := s.Policies.PolicyQuoteControllerProvisionWarranty(ctx, "pol_a3be1a473b55403582f93f070fec0b30", components.PolicyQuoteRequestBody{
        DistributorID: warrantysphere.String("dist_752b6169cf9648698602b10e85c17876"),
        Distributor: &components.DistributorDto{
            Name: "<value>",
            ParentID: warrantysphere.String("dist_3c060b448ec24879b1be2c9ff5179f73"),
        },
        CustomerID: warrantysphere.String("cus_95fb1cc0a3044be582a553b49b491f70"),
        ShippingAddress: components.AddressDto{
            Line1: "100 Holliday St",
            Line2: nil,
            Zip: "21202",
            City: "Baltimore",
            State: "MD",
            Country: "US",
            County: warrantysphere.String("Berkshire"),
        },
        AssetID: warrantysphere.String("ast_9ae3735979b34a91a250bf25a5c902d0"),
        PropertyValues: components.PolicyQuoteRequestBodyPropertyValues{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyQuoteResponseBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `policyID`                                                                             | *string*                                                                               | :heavy_check_mark:                                                                     | ID of the policy for which to request the quote                                        | pol_a3be1a473b55403582f93f070fec0b30                                                   |
| `policyQuoteRequestBody`                                                               | [components.PolicyQuoteRequestBody](../../models/components/policyquoterequestbody.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.PolicyQuoteControllerProvisionWarrantyResponse](../../models/operations/policyquotecontrollerprovisionwarrantyresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |