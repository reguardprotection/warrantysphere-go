# PoliciesPlans
(*PoliciesPlans*)

## Overview

### Available Operations

* [ExportPlanCoverageConfigsQueryControllerExportPolicies](#exportplancoverageconfigsquerycontrollerexportpolicies) - Export Plan Coverage Configuration
* [ListPolicyPlansQueryControllerListPolicyPlans](#listpolicyplansquerycontrollerlistpolicyplans) - List Plans
* [CreatePolicyPlanCommandControllerCreatePlan](#createpolicyplancommandcontrollercreateplan) - Create Plan
* [RetrievePolicyPlanQueryControllerRetrievePolicyPlan](#retrievepolicyplanquerycontrollerretrievepolicyplan) - Retrieve Plan
* [UpdatePolicyPlanCommandControllerUpdate](#updatepolicyplancommandcontrollerupdate) - Update Plan
* [DeletePolicyPlanCommandControllerDelete](#deletepolicyplancommandcontrollerdelete) - Delete Plan
* [SetVisibilityCommandControllerCreatePlan](#setvisibilitycommandcontrollercreateplan) - Set Plan Visibility

## ExportPlanCoverageConfigsQueryControllerExportPolicies

Export Plan Coverage Configuration

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
    res, err := s.PoliciesPlans.ExportPlanCoverageConfigsQueryControllerExportPolicies(ctx, "policy_af37e94fa3f94a3792a673bab32a3dda", "pln_af37e94fa3f94a3792a673bab32a3dda")
    if err != nil {
        log.Fatal(err)
    }
    if res.ExportPlanCoverageConfigsQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy.                         | policy_af37e94fa3f94a3792a673bab32a3dda                  |
| `planID`                                                 | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the plan.                           | pln_af37e94fa3f94a3792a673bab32a3dda                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ExportPlanCoverageConfigsQueryControllerExportPoliciesResponse](../../models/operations/exportplancoverageconfigsquerycontrollerexportpoliciesresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## ListPolicyPlansQueryControllerListPolicyPlans

List Plans

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
    res, err := s.PoliciesPlans.ListPolicyPlansQueryControllerListPolicyPlans(ctx, "pol_0fe54a7bf4ae405592011ac6a0aac011", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPolicyPlansQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the policy                          | pol_0fe54a7bf4ae405592011ac6a0aac011                     |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListPolicyPlansQueryControllerListPolicyPlansResponse](../../models/operations/listpolicyplansquerycontrollerlistpolicyplansresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## CreatePolicyPlanCommandControllerCreatePlan

Create Plan

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
    res, err := s.PoliciesPlans.CreatePolicyPlanCommandControllerCreatePlan(ctx, "pol_ca1ef5585b7144c8a7d1519074b9d294", components.CreatePolicyPlanCommandRequest{
        CanCreateNewVersion: true,
        Order: 109.76,
        Title: "My Plan",
        Description: warrantysphere.String("This is my plan"),
        Visible: true,
        Conditions: []components.ConditionGroupDto{
            components.ConditionGroupDto{
                Comparison: components.ComparisonAnd,
                Conditions: [][]string{
                    []string{
                        "<value>",
                        "<value>",
                    },
                    []string{
                        "<value>",
                    },
                },
            },
            components.ConditionGroupDto{
                Comparison: components.ComparisonAnd,
                Conditions: [][]string{
                    []string{
                        "<value>",
                    },
                },
            },
            components.ConditionGroupDto{
                Comparison: components.ComparisonAnd,
                Conditions: [][]string{
                    []string{
                        "<value>",
                        "<value>",
                    },
                    []string{
                        "<value>",
                        "<value>",
                        "<value>",
                    },
                    []string{
                        "<value>",
                        "<value>",
                        "<value>",
                    },
                },
            },
        },
        Addons: []components.AddonDto{
            components.AddonDto{
                ID: "<id>",
                Title: "Electrical Group",
                Description: "Electrical appliances addon.",
            },
            components.AddonDto{
                ID: "<id>",
                Title: "Electrical Group",
                Description: "Electrical appliances addon.",
            },
        },
        LiabilityGroups: []components.LiabilityGroupDto{
            components.LiabilityGroupDto{
                ID: "<id>",
                Title: "Electrical Liability Group",
                Description: "Liability group for electrical appliances.",
                Limit: 100,
                LimitPropertyPath: "address.city",
            },
        },
        CoverageConfigs: []components.PolicyTermCoverageConfigDto{
            components.PolicyTermCoverageConfigDto{
                TermID: "<id>",
                Coverages: []components.CoverageConfigDto{
                    components.CoverageConfigDto{
                        CoverageID: "<id>",
                        Requirement: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Deductible: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityLimit: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Price: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityGroups: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                    },
                },
            },
            components.PolicyTermCoverageConfigDto{
                TermID: "<id>",
                Coverages: []components.CoverageConfigDto{
                    components.CoverageConfigDto{
                        CoverageID: "<id>",
                        Requirement: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Deductible: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityLimit: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Price: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityGroups: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                    },
                    components.CoverageConfigDto{
                        CoverageID: "<id>",
                        Requirement: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Deductible: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityLimit: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Price: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                                "<value>",
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityGroups: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                    },
                },
            },
            components.PolicyTermCoverageConfigDto{
                TermID: "<id>",
                Coverages: []components.CoverageConfigDto{
                    components.CoverageConfigDto{
                        CoverageID: "<id>",
                        Requirement: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Deductible: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityLimit: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        Price: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                        LiabilityGroups: components.RuleDto{
                            Blocks: []string{
                                "<value>",
                            },
                            DefaultValue: components.DefaultValue{},
                            ValuePropertyPath: "address.city",
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePolicyPlanCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `policyID`                                                                                             | *string*                                                                                               | :heavy_check_mark:                                                                                     | Identifier of the linked policy.                                                                       | pol_ca1ef5585b7144c8a7d1519074b9d294                                                                   |
| `createPolicyPlanCommandRequest`                                                                       | [components.CreatePolicyPlanCommandRequest](../../models/components/createpolicyplancommandrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.CreatePolicyPlanCommandControllerCreatePlanResponse](../../models/operations/createpolicyplancommandcontrollercreateplanresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## RetrievePolicyPlanQueryControllerRetrievePolicyPlan

Retrieve Plan

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
    res, err := s.PoliciesPlans.RetrievePolicyPlanQueryControllerRetrievePolicyPlan(ctx, "policy_7d1a92d4818748e58158926ca09706bc", "pln_7d1a92d4818748e58158926ca09706bc")
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyPlanAggregate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `policyID`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the wpolicylan.                     | policy_7d1a92d4818748e58158926ca09706bc                  |
| `planID`                                                 | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the plan.                           | pln_7d1a92d4818748e58158926ca09706bc                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RetrievePolicyPlanQueryControllerRetrievePolicyPlanResponse](../../models/operations/retrievepolicyplanquerycontrollerretrievepolicyplanresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## UpdatePolicyPlanCommandControllerUpdate

Update Plan

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
    res, err := s.PoliciesPlans.UpdatePolicyPlanCommandControllerUpdate(ctx, "policy_7d1a92d4818748e58158926ca09706bc", "pln_01be00dd09f14bfba816cf242f653af8", components.UpdatePolicyPlanCommandRequest{
        CanCreateNewVersion: false,
        Conditions: &components.ConditionGroupDto{
            Comparison: components.ComparisonAnd,
            Conditions: [][]string{
                []string{
                    "<value>",
                },
            },
        },
        LiabilityLimitGroup: []components.LiabilityGroupDto{
            components.LiabilityGroupDto{
                ID: "<id>",
                Title: "Electrical Liability Group",
                Description: "Liability group for electrical appliances.",
                Limit: 100,
                LimitPropertyPath: "address.city",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdatePolicyPlanCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `policyID`                                                                                             | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the wpolicylan.                                                                   | policy_7d1a92d4818748e58158926ca09706bc                                                                |
| `planID`                                                                                               | *string*                                                                                               | :heavy_check_mark:                                                                                     | Unique identifier of the plan.                                                                         | pln_01be00dd09f14bfba816cf242f653af8                                                                   |
| `updatePolicyPlanCommandRequest`                                                                       | [components.UpdatePolicyPlanCommandRequest](../../models/components/updatepolicyplancommandrequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.UpdatePolicyPlanCommandControllerUpdateResponse](../../models/operations/updatepolicyplancommandcontrollerupdateresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DeletePolicyPlanCommandControllerDelete

Delete Plan

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
    res, err := s.PoliciesPlans.DeletePolicyPlanCommandControllerDelete(ctx, "policy_7d1a92d4818748e58158926ca09706bc", "pln_854460b42e504fa08140388e12f0aba3", components.DeletePolicyPlanCommandRequestBody{
        CanCreateNewVersion: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeletePolicyPlanCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `policyID`                                                                                                     | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Unique identifier of the wpolicylan.                                                                           | policy_7d1a92d4818748e58158926ca09706bc                                                                        |
| `planID`                                                                                                       | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Unique identifier of the policies plan.                                                                        | pln_854460b42e504fa08140388e12f0aba3                                                                           |
| `deletePolicyPlanCommandRequestBody`                                                                           | [components.DeletePolicyPlanCommandRequestBody](../../models/components/deletepolicyplancommandrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.DeletePolicyPlanCommandControllerDeleteResponse](../../models/operations/deletepolicyplancommandcontrollerdeleteresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## SetVisibilityCommandControllerCreatePlan

Set Plan Visibility

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
    res, err := s.PoliciesPlans.SetVisibilityCommandControllerCreatePlan(ctx, "policy_af37e94fa3f94a3792a673bab32a3dda", "pol_80a7f0d1b8834244ab60fbd29408aa90", components.SetPlanVisibilityCommandRequest{
        Visible: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SetPlanVisibilityCommandResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `policyID`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Unique identifier of the policy.                                                                         | policy_af37e94fa3f94a3792a673bab32a3dda                                                                  |
| `planID`                                                                                                 | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Identifier of the linked plan.                                                                           | pol_80a7f0d1b8834244ab60fbd29408aa90                                                                     |
| `setPlanVisibilityCommandRequest`                                                                        | [components.SetPlanVisibilityCommandRequest](../../models/components/setplanvisibilitycommandrequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.SetVisibilityCommandControllerCreatePlanResponse](../../models/operations/setvisibilitycommandcontrollercreateplanresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ValidationBadRequestResponse | 400                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |