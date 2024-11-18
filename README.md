# github.com/reguardprotection/warrantysphere

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/reguardprotection/warrantysphere* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/reguardprotection/warrantysphere&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/reguard/reguard). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/reguardprotection/warrantysphere
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/reguardprotection/warrantysphere"
	"log"
)

func main() {
	s := warrantysphere.New()

	ctx := context.Background()
	res, err := s.HealthControllerHealthCheck(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountingAccounts](docs/sdks/accountingaccounts/README.md)

* [ListAccountsControllerListAccounts](docs/sdks/accountingaccounts/README.md#listaccountscontrollerlistaccounts) - List Accounts
* [CreateAccountControllerCreateAccount](docs/sdks/accountingaccounts/README.md#createaccountcontrollercreateaccount) - Create Account
* [RetrieveAccountControllerRetrieveAccount](docs/sdks/accountingaccounts/README.md#retrieveaccountcontrollerretrieveaccount) - Retrieve Account
* [UpdateAccountControllerUpdateAccount](docs/sdks/accountingaccounts/README.md#updateaccountcontrollerupdateaccount) - Update Account

### [AccountingTransactions](docs/sdks/accountingtransactions/README.md)

* [ListTransactionsControllerListTransactions](docs/sdks/accountingtransactions/README.md#listtransactionscontrollerlisttransactions) - List Transactions
* [CreateTransactionControllerCreateTransaction](docs/sdks/accountingtransactions/README.md#createtransactioncontrollercreatetransaction) - Create Transaction
* [RetrieveTransactionControllerRetrieveTransaction](docs/sdks/accountingtransactions/README.md#retrievetransactioncontrollerretrievetransaction) - Retrieve Transaction
* [VoidTransactionControllerVoidTransaction](docs/sdks/accountingtransactions/README.md#voidtransactioncontrollervoidtransaction) - Void Transaction

### [Activities](docs/sdks/activities/README.md)

* [ListActivitiesQueryControllerList](docs/sdks/activities/README.md#listactivitiesquerycontrollerlist) - List Activities

### [Assets](docs/sdks/assets/README.md)

* [ListAssetsQueryControllerListAssets](docs/sdks/assets/README.md#listassetsquerycontrollerlistassets) - List Assets
* [CreateAssetCommandControllerCreate](docs/sdks/assets/README.md#createassetcommandcontrollercreate) - Create Asset
* [RetrieveAssetQueryControllerRetrieveAsset](docs/sdks/assets/README.md#retrieveassetquerycontrollerretrieveasset) - Retrieve Asset
* [UpdateAssetCommandControllerUpdate](docs/sdks/assets/README.md#updateassetcommandcontrollerupdate) - Update Asset
* [DeleteAssetCommandControllerDelete](docs/sdks/assets/README.md#deleteassetcommandcontrollerdelete) - Delete Asset

### [Claims](docs/sdks/claims/README.md)

* [ListClaimsQueryControllerList](docs/sdks/claims/README.md#listclaimsquerycontrollerlist) - List Claims
* [OpenClaimCommandControllerOpen](docs/sdks/claims/README.md#openclaimcommandcontrolleropen) - Open Claim
* [UpdateClaimCommandControllerUpdate](docs/sdks/claims/README.md#updateclaimcommandcontrollerupdate) - Update Claim
* [UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleSteps](docs/sdks/claims/README.md#updateclaimlifecyclestepcommandcontrollerupdateclaimlifecyclesteps) - Update Claim Lifecycle Steps
* [ExpireClaimCommandControllerExpire](docs/sdks/claims/README.md#expireclaimcommandcontrollerexpire) - Expire Claim
* [CancelClaimCommandControllerCancel](docs/sdks/claims/README.md#cancelclaimcommandcontrollercancel) - Cancel Claim
* [CloseClaimCommandControllerClose](docs/sdks/claims/README.md#closeclaimcommandcontrollerclose) - Close Claim

### [ClaimsDocuments](docs/sdks/claimsdocuments/README.md)

* [ListClaimDocumentsQueryControllerList](docs/sdks/claimsdocuments/README.md#listclaimdocumentsquerycontrollerlist) - List Claim Documents
* [CreateClaimDocumentControllerCreate](docs/sdks/claimsdocuments/README.md#createclaimdocumentcontrollercreate) - Create Claim Document
* [RetrieveClaimDocumentQueryControllerRetrieve](docs/sdks/claimsdocuments/README.md#retrieveclaimdocumentquerycontrollerretrieve) - Retrieve Claim Document
* [UpdateClaimDocumentControllerUpdate](docs/sdks/claimsdocuments/README.md#updateclaimdocumentcontrollerupdate) - Update Claim Document
* [DeleteClaimDocumentControllerDelete](docs/sdks/claimsdocuments/README.md#deleteclaimdocumentcontrollerdelete) - Delete Claim Document

### [ClaimsItems](docs/sdks/claimsitems/README.md)

* [ListClaimItemsQueryControllerList](docs/sdks/claimsitems/README.md#listclaimitemsquerycontrollerlist) - List Claim Items
* [CreateClaimItemCommandControllerCreate](docs/sdks/claimsitems/README.md#createclaimitemcommandcontrollercreate) - Create Claim Item
* [RetrieveClaimItemQueryControllerRetrieve](docs/sdks/claimsitems/README.md#retrieveclaimitemquerycontrollerretrieve) - Retrieve Claim Item
* [UpdateClaimItemCommandControllerUdpate](docs/sdks/claimsitems/README.md#updateclaimitemcommandcontrollerudpate) - Update Claim Item
* [DeleteClaimItemCommandControllerDelete](docs/sdks/claimsitems/README.md#deleteclaimitemcommandcontrollerdelete) - Delete Claim Item
* [ApproveClaimItemCommandControllerApprove](docs/sdks/claimsitems/README.md#approveclaimitemcommandcontrollerapprove) - Approve Claim Item
* [RejectClaimItemCommandControllerReject](docs/sdks/claimsitems/README.md#rejectclaimitemcommandcontrollerreject) - Reject Claim Item
* [ResolveClaimItemCommandControllerResolve](docs/sdks/claimsitems/README.md#resolveclaimitemcommandcontrollerresolve) - Resolve Claim Item
* [RedraftClaimItemCommandControllerRedraft](docs/sdks/claimsitems/README.md#redraftclaimitemcommandcontrollerredraft) - Redraft Claim Item

### [ClaimsNotes](docs/sdks/claimsnotes/README.md)

* [ListClaimNotesQueryControllerList](docs/sdks/claimsnotes/README.md#listclaimnotesquerycontrollerlist) - List Claim Notes
* [CreateClaimNoteCommandControllerCreate](docs/sdks/claimsnotes/README.md#createclaimnotecommandcontrollercreate) - Create Claim Note
* [RetrieveClaimNoteQueryControllerRetrieve](docs/sdks/claimsnotes/README.md#retrieveclaimnotequerycontrollerretrieve) - Retrieve Claim Note
* [UpdateClaimNoteCommandControllerUdpate](docs/sdks/claimsnotes/README.md#updateclaimnotecommandcontrollerudpate) - Update Claim Note
* [DeleteClaimNoteCommandControllerDelete](docs/sdks/claimsnotes/README.md#deleteclaimnotecommandcontrollerdelete) - Delete Claim Note

### [ClaimsPayments](docs/sdks/claimspayments/README.md)

* [ListClaimPaymentsQueryControllerList](docs/sdks/claimspayments/README.md#listclaimpaymentsquerycontrollerlist) - List Claim Payments
* [CreateClaimPaymentCommandControllerCreate](docs/sdks/claimspayments/README.md#createclaimpaymentcommandcontrollercreate) - Create Claim Payment
* [RetrieveClaimPaymentQueryControllerRetrieve](docs/sdks/claimspayments/README.md#retrieveclaimpaymentquerycontrollerretrieve) - Retrieve Claim Payment
* [UpdateClaimPaymentCommandControllerUpdate](docs/sdks/claimspayments/README.md#updateclaimpaymentcommandcontrollerupdate) - Update Claim Payment
* [IssueClaimPaymentCommandControllerUpdate](docs/sdks/claimspayments/README.md#issueclaimpaymentcommandcontrollerupdate) - Issue Claim Payment
* [VoidClaimPaymentCommandControllerUpdate](docs/sdks/claimspayments/README.md#voidclaimpaymentcommandcontrollerupdate) - Void Claim Payment

### [Customers](docs/sdks/customers/README.md)

* [ListCustomersControllerList](docs/sdks/customers/README.md#listcustomerscontrollerlist) - List Customers
* [CreateCustomerControllerCreateCustomer](docs/sdks/customers/README.md#createcustomercontrollercreatecustomer) - Create Customer
* [GenerateCustomerSignInLinkControllerGenerateSignIn](docs/sdks/customers/README.md#generatecustomersigninlinkcontrollergeneratesignin) - Generate Customer Sign In Link
* [UpdateCustomerControllerUpdateCustomer](docs/sdks/customers/README.md#updatecustomercontrollerupdatecustomer) - Update Customer
* [DeleteCustomerControllerDeleteCustomer](docs/sdks/customers/README.md#deletecustomercontrollerdeletecustomer) - Delete Customer

### [Distributors](docs/sdks/distributors/README.md)

* [ListDistributorsQueryControllerListDistributors](docs/sdks/distributors/README.md#listdistributorsquerycontrollerlistdistributors) - List Distributors
* [CreateDistributorCommandControllerCreateDistributor](docs/sdks/distributors/README.md#createdistributorcommandcontrollercreatedistributor) - Create Distributor
* [RetrieveDistributorQueryControllerRetrieveDistributor](docs/sdks/distributors/README.md#retrievedistributorquerycontrollerretrievedistributor) - Retrieve Distributor
* [UpdateDistributorCommandControllerUpdateDistributor](docs/sdks/distributors/README.md#updatedistributorcommandcontrollerupdatedistributor) - Update Distributor
* [ArchiveDistributorCommandControllerArchive](docs/sdks/distributors/README.md#archivedistributorcommandcontrollerarchive) - Archive Distributor
* [DeactivateDistributorCommandControllerDeactivate](docs/sdks/distributors/README.md#deactivatedistributorcommandcontrollerdeactivate) - Deactivate Distributor
* [ReactivateDistributorCommandControllerArchive](docs/sdks/distributors/README.md#reactivatedistributorcommandcontrollerarchive) - Reactivate Distributor
* [CreateDistributorOnboardingLinkCommandControllerGenerateDistributorOnboardingLink](docs/sdks/distributors/README.md#createdistributoronboardinglinkcommandcontrollergeneratedistributoronboardinglink) - Create Link

### [Documents](docs/sdks/documents/README.md)

* [DocumentsList](docs/sdks/documents/README.md#documentslist) - List Documents
* [DocumentsUpload](docs/sdks/documents/README.md#documentsupload) - Upload Document
* [DocumentsRetrieve](docs/sdks/documents/README.md#documentsretrieve) - Retrieve Document
* [DocumentsUpdate](docs/sdks/documents/README.md#documentsupdate) - Update Document
* [DocumentsDelete](docs/sdks/documents/README.md#documentsdelete) - Delete Document

### [DocumentsProcessors](docs/sdks/documentsprocessors/README.md)

* [DocumentProcessorsList](docs/sdks/documentsprocessors/README.md#documentprocessorslist) - List Processors
* [DocumentProcessorsRetrieve](docs/sdks/documentsprocessors/README.md#documentprocessorsretrieve) - Retrieve Processor

### [DocumentsProcessorsRuns](docs/sdks/documentsprocessorsruns/README.md)

* [DocumentProcessorRunsList](docs/sdks/documentsprocessorsruns/README.md#documentprocessorrunslist) - List Processor Runs
* [DocumentProcessorRunsCreate](docs/sdks/documentsprocessorsruns/README.md#documentprocessorrunscreate) - Create Processor Run
* [DocumentProcessorRunsRetrieve](docs/sdks/documentsprocessorsruns/README.md#documentprocessorrunsretrieve) - Retrieve Processor Run
* [DocumentProcessorRunsDelete](docs/sdks/documentsprocessorsruns/README.md#documentprocessorrunsdelete) - Delete Processor Run

### [Notes](docs/sdks/notes/README.md)

* [ListNotesQueryControllerList](docs/sdks/notes/README.md#listnotesquerycontrollerlist) - List Notes
* [CreateNoteCommandControllerCreate](docs/sdks/notes/README.md#createnotecommandcontrollercreate) - Create Note
* [RetrieveNoteQueryControllerRetrieve](docs/sdks/notes/README.md#retrievenotequerycontrollerretrieve) - Retrieve Note
* [UpdateNoteCommandControllerUpdate](docs/sdks/notes/README.md#updatenotecommandcontrollerupdate) - Update Note
* [DeleteNoteCommandControllerDelete](docs/sdks/notes/README.md#deletenotecommandcontrollerdelete) - Delete Note

### [OrganizationSettings](docs/sdks/organizationsettings/README.md)

* [RetrieveDistributorsSettingsQueryControllerRetrieveDistributor](docs/sdks/organizationsettings/README.md#retrievedistributorssettingsquerycontrollerretrievedistributor) - Retrieve Organization Settings
* [UpdateDistributorsSettingsQueryControllerRetrieveDistributor](docs/sdks/organizationsettings/README.md#updatedistributorssettingsquerycontrollerretrievedistributor) - Update Organization Settings

### [Policies](docs/sdks/policies/README.md)

* [ListPoliciesQueryControllerListPolicies](docs/sdks/policies/README.md#listpoliciesquerycontrollerlistpolicies) - List Policies
* [CreatePolicyCommandControllerCreatePolicy](docs/sdks/policies/README.md#createpolicycommandcontrollercreatepolicy) - Create Policy
* [ExportPoliciesQueryControllerExportPolicies](docs/sdks/policies/README.md#exportpoliciesquerycontrollerexportpolicies) - Export Policies
* [RetrievePolicyQueryControllerRetrievePolicy](docs/sdks/policies/README.md#retrievepolicyquerycontrollerretrievepolicy) - Retrieve Policy
* [UpdatePolicyCommandControllerUpdatePolicy](docs/sdks/policies/README.md#updatepolicycommandcontrollerupdatepolicy) - Update Policy
* [DeletePolicyCommandControllerDelete](docs/sdks/policies/README.md#deletepolicycommandcontrollerdelete) - Delete Policy
* [UpdatePolicyDocumentCommandControllerUpdatePolicyDocument](docs/sdks/policies/README.md#updatepolicydocumentcommandcontrollerupdatepolicydocument) - Update Policy Document
* [UpdatePolicyEmailCommandControllerUpdate](docs/sdks/policies/README.md#updatepolicyemailcommandcontrollerupdate) - Update Policy Email
* [UpdateClaimLifecycleStepCommandControllerUpdateClaimLifecycleStep](docs/sdks/policies/README.md#updateclaimlifecyclestepcommandcontrollerupdateclaimlifecyclestep) - Update Policy Claim Lifecycle Steps
* [PublishPolicyCommandControllerPublish](docs/sdks/policies/README.md#publishpolicycommandcontrollerpublish) - Publish Policy
* [ArchivePolicyCommandControllerArchive](docs/sdks/policies/README.md#archivepolicycommandcontrollerarchive) - Archive Policy
* [DuplicatePolicyCommandControllerDuplicate](docs/sdks/policies/README.md#duplicatepolicycommandcontrollerduplicate) - Duplicate
* [MigratePolicyFromSandboxCommandControllerMigrateSandbox](docs/sdks/policies/README.md#migratepolicyfromsandboxcommandcontrollermigratesandbox) - Migrate Sandbox Policy to Live
* [UpdatePolicyCancelationRuleCommandControllerUpdate](docs/sdks/policies/README.md#updatepolicycancelationrulecommandcontrollerupdate) - Update Cancelation Rule
* [DeletePolicyCancelationRuleCommandControllerUpdate](docs/sdks/policies/README.md#deletepolicycancelationrulecommandcontrollerupdate) - Delete Cancelation Rule
* [CreatePolicyCancelationRuleCommandControllerCreate](docs/sdks/policies/README.md#createpolicycancelationrulecommandcontrollercreate) - Create Cancelation Rule
* [PolicyQuoteControllerProvisionWarranty](docs/sdks/policies/README.md#policyquotecontrollerprovisionwarranty) - Quote

### [PoliciesCoverages](docs/sdks/policiescoverages/README.md)

* [ListCoveragesQueryControllerListCoverage](docs/sdks/policiescoverages/README.md#listcoveragesquerycontrollerlistcoverage) - List Policy Coverages
* [CreateCoverageControllerCreateCoverage](docs/sdks/policiescoverages/README.md#createcoveragecontrollercreatecoverage) - Create Policy Coverage
* [ExportCoveragesQueryControllerExportCoverages](docs/sdks/policiescoverages/README.md#exportcoveragesquerycontrollerexportcoverages) - Export Policy Coverages
* [RetrieveCoverageQueryControllerRetrieveCoverage](docs/sdks/policiescoverages/README.md#retrievecoveragequerycontrollerretrievecoverage) - Retrieve Policy Coverage
* [DeleteCoverageControllerDeleteCoverageCommandResponse](docs/sdks/policiescoverages/README.md#deletecoveragecontrollerdeletecoveragecommandresponse) - Delete Policy Coverage
* [UpdateCoverageCommandControllerUpdate](docs/sdks/policiescoverages/README.md#updatecoveragecommandcontrollerupdate) - Update Policy Coverage

### [PoliciesPlans](docs/sdks/policiesplans/README.md)

* [ExportPlanCoverageConfigsQueryControllerExportPolicies](docs/sdks/policiesplans/README.md#exportplancoverageconfigsquerycontrollerexportpolicies) - Export Plan Coverage Configuration
* [ListPolicyPlansQueryControllerListPolicyPlans](docs/sdks/policiesplans/README.md#listpolicyplansquerycontrollerlistpolicyplans) - List Plans
* [CreatePolicyPlanCommandControllerCreatePlan](docs/sdks/policiesplans/README.md#createpolicyplancommandcontrollercreateplan) - Create Plan
* [RetrievePolicyPlanQueryControllerRetrievePolicyPlan](docs/sdks/policiesplans/README.md#retrievepolicyplanquerycontrollerretrievepolicyplan) - Retrieve Plan
* [UpdatePolicyPlanCommandControllerUpdate](docs/sdks/policiesplans/README.md#updatepolicyplancommandcontrollerupdate) - Update Plan
* [DeletePolicyPlanCommandControllerDelete](docs/sdks/policiesplans/README.md#deletepolicyplancommandcontrollerdelete) - Delete Plan
* [SetVisibilityCommandControllerCreatePlan](docs/sdks/policiesplans/README.md#setvisibilitycommandcontrollercreateplan) - Set Plan Visibility

### [PortalConfig](docs/sdks/portalconfig/README.md)

* [RetrievePortalConfigControllerUpdatePortalConfig](docs/sdks/portalconfig/README.md#retrieveportalconfigcontrollerupdateportalconfig) - Retrieve Portal Configuration
* [UpdatePortalConfigControllerUpdatePortalConfig](docs/sdks/portalconfig/README.md#updateportalconfigcontrollerupdateportalconfig) - Update Portal Configuration

### [Properties](docs/sdks/properties/README.md)

* [RetrievePropertyControllerRetrieve](docs/sdks/properties/README.md#retrievepropertycontrollerretrieve) - Retrieve Property
* [ListPropertiesControllerRetrieve](docs/sdks/properties/README.md#listpropertiescontrollerretrieve) - List Properties
* [CreateBooleanPropertyControllerCreate](docs/sdks/properties/README.md#createbooleanpropertycontrollercreate) - Create Boolean Property
* [CreateDatePropertyControllerCreate](docs/sdks/properties/README.md#createdatepropertycontrollercreate) - Create Date Property
* [CreateNumberPropertyControllerCreate](docs/sdks/properties/README.md#createnumberpropertycontrollercreate) - Create Number Property
* [CreateObjectPropertyControllerCreate](docs/sdks/properties/README.md#createobjectpropertycontrollercreate) - Create Object Property
* [CreateTextPropertyControllerCreate](docs/sdks/properties/README.md#createtextpropertycontrollercreate) - Create Text Property

### [PropertySets](docs/sdks/propertysets/README.md)

* [RetrievePropertySetControllerRetrieve](docs/sdks/propertysets/README.md#retrievepropertysetcontrollerretrieve) - Retrieve Property Set
* [ListPropertySetsControllerRetrieve](docs/sdks/propertysets/README.md#listpropertysetscontrollerretrieve) - List Property Sets
* [CreatePropertySetControllerCreate](docs/sdks/propertysets/README.md#createpropertysetcontrollercreate) - Create Property Set
* [RetrievePropertySetJSONControllerRetrieve](docs/sdks/propertysets/README.md#retrievepropertysetjsoncontrollerretrieve) - Retrieve Property Set JSON

### [Sales](docs/sdks/sales/README.md)

* [ListPaymentsQueryControllerListPayments](docs/sdks/sales/README.md#listpaymentsquerycontrollerlistpayments) - List Payments
* [ListPaymentsQueryControllerExportPayments](docs/sdks/sales/README.md#listpaymentsquerycontrollerexportpayments) - Export Payments (CSV)
* [RetrievePaymentQueryControllerRetrievePayment](docs/sdks/sales/README.md#retrievepaymentquerycontrollerretrievepayment) - Retrieve Payment
* [ResendPaymentReminderCommandControllerResendPaymentReminder](docs/sdks/sales/README.md#resendpaymentremindercommandcontrollerresendpaymentreminder) - Resend Payment Reminder
* [RetryPaymentCommandControllerRetryPayment](docs/sdks/sales/README.md#retrypaymentcommandcontrollerretrypayment) - Retry Payment

### [Staff](docs/sdks/staff/README.md)

* [ListStaffQueryControllerListStaff](docs/sdks/staff/README.md#liststaffquerycontrollerliststaff) - List Staff
* [RetrieveStaffQueryControllerRetrieveStaff](docs/sdks/staff/README.md#retrievestaffquerycontrollerretrievestaff) - Retrieve Staff
* [RetrieveStaffPermissionsQueryControllerRetrieveStaffPermissions](docs/sdks/staff/README.md#retrievestaffpermissionsquerycontrollerretrievestaffpermissions) - Retrieve Staff Roles
* [InviteStaffAndGrantRolesCommandControllerInviteStaffAndGrantRoles](docs/sdks/staff/README.md#invitestaffandgrantrolescommandcontrollerinvitestaffandgrantroles) - Invite Staff and Grant Roles
* [UpdateStaffControllerUpdateStaff](docs/sdks/staff/README.md#updatestaffcontrollerupdatestaff) - Update Staff
* [CreateRoleAssignmentCommandControllerCreate](docs/sdks/staff/README.md#createroleassignmentcommandcontrollercreate) - Grant Staff Role
* [RemoveRoleAssignmentCommandControllerRemove](docs/sdks/staff/README.md#removeroleassignmentcommandcontrollerremove) - Remove Staff Role
* [DeactivateStaffControllerDeactivateStaff](docs/sdks/staff/README.md#deactivatestaffcontrollerdeactivatestaff) - Deactivate Staff

### [StaffAPIKey](docs/sdks/staffapikey/README.md)

* [ListAPIKeysQueryControllerListKeys](docs/sdks/staffapikey/README.md#listapikeysquerycontrollerlistkeys) - List Staff API Keys
* [CreateAPIKeyCommandControllerCreate](docs/sdks/staffapikey/README.md#createapikeycommandcontrollercreate) - Create Staff API Key
* [RollAPIKeyCommandControllerRoll](docs/sdks/staffapikey/README.md#rollapikeycommandcontrollerroll) - Roll Staff API Key
* [UpdateAPIKeyCommandControllerUpdate](docs/sdks/staffapikey/README.md#updateapikeycommandcontrollerupdate) - Update Staff API Key
* [DisableAPIKeyCommandControllerDisable](docs/sdks/staffapikey/README.md#disableapikeycommandcontrollerdisable) - Disable Staff API Key

### [Taxes](docs/sdks/taxes/README.md)

* [CalculateTaxesControllerCalculateTaxes](docs/sdks/taxes/README.md#calculatetaxescontrollercalculatetaxes) - Calculate Taxes

### [Warranties](docs/sdks/warranties/README.md)

* [ListWarrantiesControllerListWarranties](docs/sdks/warranties/README.md#listwarrantiescontrollerlistwarranties) - List Warranties
* [ProvisionWarrantyControllerProvisionWarranty](docs/sdks/warranties/README.md#provisionwarrantycontrollerprovisionwarranty) - Provision Warranty
* [UpdateWarrantyControllerUpdateWarranty](docs/sdks/warranties/README.md#updatewarrantycontrollerupdatewarranty) - Update
* [DeleteWarrantyControllerDeleteWarranty](docs/sdks/warranties/README.md#deletewarrantycontrollerdeletewarranty) - Delete
* [FinalizeWarrantyControllerUpdateWarranty](docs/sdks/warranties/README.md#finalizewarrantycontrollerupdatewarranty) - Finalize
* [CancelWarrantyControllerCancelWarranty](docs/sdks/warranties/README.md#cancelwarrantycontrollercancelwarranty) - Cancel Warranty
* [ActivateWarrantyControllerActivateWarranty](docs/sdks/warranties/README.md#activatewarrantycontrolleractivatewarranty) - Activate Warranty

### [WarrantySphere SDK](docs/sdks/warrantysphere/README.md)

* [HealthControllerHealthCheck](docs/sdks/warrantysphere/README.md#healthcontrollerhealthcheck) - Health Check

### [Webhooks](docs/sdks/webhooks/README.md)

* [ListWebhookEventsControllerList](docs/sdks/webhooks/README.md#listwebhookeventscontrollerlist) - List Webhook Events
* [ListWebhooksControllerList](docs/sdks/webhooks/README.md#listwebhookscontrollerlist) - List Webhooks
* [CreateWebhookControllerCreate](docs/sdks/webhooks/README.md#createwebhookcontrollercreate) - Create Webhook
* [DeleteWebhookControllerDelete](docs/sdks/webhooks/README.md#deletewebhookcontrollerdelete) - Delete Webhook
* [UpdateWebhookControllerUpdate](docs/sdks/webhooks/README.md#updatewebhookcontrollerupdate) - Update Webhook

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/reguardprotection/warrantysphere"
	"github.com/reguardprotection/warrantysphere/retry"
	"log"
	"models/operations"
)

func main() {
	s := warrantysphere.New()

	ctx := context.Background()
	res, err := s.HealthControllerHealthCheck(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/reguardprotection/warrantysphere"
	"github.com/reguardprotection/warrantysphere/retry"
	"log"
)

func main() {
	s := warrantysphere.New(
		warrantysphere.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	ctx := context.Background()
	res, err := s.HealthControllerHealthCheck(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `ListAccountsControllerListAccounts` function may return the following errors:

| Error Type                             | Status Code | Content Type     |
| -------------------------------------- | ----------- | ---------------- |
| apierrors.ValidationBadRequestResponse | 400         | application/json |
| apierrors.APIError                     | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/reguardprotection/warrantysphere"
	"github.com/reguardprotection/warrantysphere/models/apierrors"
	"log"
	"os"
)

func main() {
	s := warrantysphere.New(
		warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
	)

	ctx := context.Background()
	res, err := s.AccountingAccounts.ListAccountsControllerListAccounts(ctx, nil, nil, nil)
	if err != nil {

		var e *apierrors.ValidationBadRequestResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/reguardprotection/warrantysphere"
	"log"
)

func main() {
	s := warrantysphere.New(
		warrantysphere.WithServerURL("https://api.warrantysphere.com"),
	)

	ctx := context.Background()
	res, err := s.HealthControllerHealthCheck(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  | Environment Variable     |
| -------- | ------ | ------- | ------------------------ |
| `APIKey` | apiKey | API key | `WARRANTYSPHERE_API_KEY` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/reguardprotection/warrantysphere"
	"log"
	"os"
)

func main() {
	s := warrantysphere.New(
		warrantysphere.WithSecurity(os.Getenv("WARRANTYSPHERE_API_KEY")),
	)

	ctx := context.Background()
	res, err := s.HealthControllerHealthCheck(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/reguardprotection/warrantysphere&utm_campaign=go)
