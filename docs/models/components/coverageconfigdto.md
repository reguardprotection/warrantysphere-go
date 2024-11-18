# CoverageConfigDto


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `CoverageID`                                             | *string*                                                 | :heavy_check_mark:                                       | Unique identifier of the coverage config.                |
| `Requirement`                                            | [components.RuleDto](../../models/components/ruledto.md) | :heavy_check_mark:                                       | The requirement rule for the coverage.                   |
| `Deductible`                                             | [components.RuleDto](../../models/components/ruledto.md) | :heavy_check_mark:                                       | The deductible rule for the coverage.                    |
| `LiabilityLimit`                                         | [components.RuleDto](../../models/components/ruledto.md) | :heavy_check_mark:                                       | The liability limit rule for the coverage.               |
| `Price`                                                  | [components.RuleDto](../../models/components/ruledto.md) | :heavy_check_mark:                                       | The retail price rule for the coverage.                  |
| `LiabilityGroups`                                        | [components.RuleDto](../../models/components/ruledto.md) | :heavy_check_mark:                                       | The liability groups rule for the coverage.              |