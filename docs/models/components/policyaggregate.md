# PolicyAggregate


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ID`                                                                                      | *string*                                                                                  | :heavy_check_mark:                                                                        | Unique identifier of the policy.                                                          | pol_22bcc406229547ef8df17a04de309cc5                                                      |
| `ReferenceID`                                                                             | *string*                                                                                  | :heavy_check_mark:                                                                        | User-defined reference ID.                                                                |                                                                                           |
| `Created`                                                                                 | [time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_check_mark:                                                                        | Datetime when the policy was created.                                                     | 2024-11-18 15:05:48.724 +0000 UTC                                                         |
| `Modified`                                                                                | [time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_check_mark:                                                                        | Datetime when the policy was last modified.                                               | 2024-11-18 15:05:48.724 +0000 UTC                                                         |
| `Deleted`                                                                                 | [time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_check_mark:                                                                        | Datetime when the policy was deleted.                                                     | 2024-11-18 15:05:48.724 +0000 UTC                                                         |
| `PublishedDate`                                                                           | [time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_check_mark:                                                                        | N/A                                                                                       | 2024-11-18 15:05:48.724 +0000 UTC                                                         |
| `PolicyNumber`                                                                            | *string*                                                                                  | :heavy_check_mark:                                                                        | Formatted unique identifier for the policy.                                               | HGP-20248041                                                                              |
| `SucceededByID`                                                                           | **string*                                                                                 | :heavy_minus_sign:                                                                        | Pointer to the policy version that comes after this one                                   | pol_0336c83ebc6348d2a1896a89c8cd4d83                                                      |
| `SucceededBy`                                                                             | [*components.SucceededBy](../../models/components/succeededby.md)                         | :heavy_minus_sign:                                                                        | Linked policy succeeding this one.                                                        |                                                                                           |
| `Title`                                                                                   | *string*                                                                                  | :heavy_check_mark:                                                                        | Descriptive title of the policy. Used in the Customer Portal application.                 |                                                                                           |
| `FriendlyTitle`                                                                           | *string*                                                                                  | :heavy_check_mark:                                                                        | Friendly title of the policy. Used in policy checkout.                                    |                                                                                           |
| `Icon`                                                                                    | *string*                                                                                  | :heavy_check_mark:                                                                        | Displayed to customers and distributors when selecting a policy at checkout.              |                                                                                           |
| `Tagline`                                                                                 | *string*                                                                                  | :heavy_check_mark:                                                                        | A tagline that provides a brief description of the policy.                                |                                                                                           |
| `Description`                                                                             | *string*                                                                                  | :heavy_check_mark:                                                                        | A description that provides a detailed description of the policy.                         |                                                                                           |
| `Validity`                                                                                | *string*                                                                                  | :heavy_check_mark:                                                                        | Text rule for the validity period of the provisioned warranties.                          |                                                                                           |
| `Status`                                                                                  | [*components.PolicyAggregateStatus](../../models/components/policyaggregatestatus.md)     | :heavy_minus_sign:                                                                        | Current status of the policy.                                                             |                                                                                           |
| `PropertySetID`                                                                           | *string*                                                                                  | :heavy_check_mark:                                                                        | The unique identifier of the property set                                                 | prpset_a3067331f7fb4efda4ffa21abd9938f2                                                   |
| `PropertySet`                                                                             | [components.PropertySetAggregate](../../models/components/propertysetaggregate.md)        | :heavy_check_mark:                                                                        | N/A                                                                                       |                                                                                           |
| `ClaimWaitingPeriod`                                                                      | *float64*                                                                                 | :heavy_check_mark:                                                                        | Duration (in days) until a customer can make a claim for this policy.                     |                                                                                           |
| `TermsAndConditionsID`                                                                    | *string*                                                                                  | :heavy_check_mark:                                                                        | Unique identifier for the linked Terms & Conditions document for this policy.             |                                                                                           |
| `TermsAndConditions`                                                                      | [*components.PolicyDocumentAggregate](../../models/components/policydocumentaggregate.md) | :heavy_minus_sign:                                                                        | Linked Terms & Conditions document for this policy.                                       |                                                                                           |
| `CoverageSummaryID`                                                                       | *string*                                                                                  | :heavy_check_mark:                                                                        | Unique identifier for the linked Coverage Summary document for this policy.               |                                                                                           |
| `CoverageSummary`                                                                         | [*components.PolicyDocumentAggregate](../../models/components/policydocumentaggregate.md) | :heavy_minus_sign:                                                                        | Linked Coverage Summary document for this policy.                                         |                                                                                           |
| `WelcomeEmailDelayHours`                                                                  | *float64*                                                                                 | :heavy_check_mark:                                                                        | Number of hours the warranty provisioned email is delayed.                                |                                                                                           |
| `Coverages`                                                                               | [][][components.Coverages](../../models/components/coverages.md)                          | :heavy_minus_sign:                                                                        | Coverages included in the policy.                                                         |                                                                                           |
| `Emails`                                                                                  | [][components.PolicyEmailAggregate](../../models/components/policyemailaggregate.md)      | :heavy_minus_sign:                                                                        | Emails included in the policy.                                                            |                                                                                           |
| `CancelationRules`                                                                        | []*string*                                                                                | :heavy_minus_sign:                                                                        | Cancelation rules for the policy.                                                         |                                                                                           |
| `Plans`                                                                                   | [][][components.Plans](../../models/components/plans.md)                                  | :heavy_minus_sign:                                                                        | Plans included in the policy.                                                             |                                                                                           |
| `ClaimLifecycleSteps`                                                                     | [][][components.ClaimLifecycleSteps](../../models/components/claimlifecyclesteps.md)      | :heavy_minus_sign:                                                                        | The claim lifecycle steps used for this policy.                                           |                                                                                           |
| `Terms`                                                                                   | [][components.PolicyTerm](../../models/components/policyterm.md)                          | :heavy_check_mark:                                                                        | Allowed durations and payment schedules for this policy.                                  |                                                                                           |
| `RequiresWarrantyActivation`                                                              | **bool*                                                                                   | :heavy_minus_sign:                                                                        | Whether or not provisioned warranties require manual activation from the customer.        |                                                                                           |