// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package billing

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/billing/mgmt/2018-11-01-preview/billing"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccountType = original.AccountType

const (
	AccountTypeEnrollment   AccountType = original.AccountTypeEnrollment
	AccountTypeOrganization AccountType = original.AccountTypeOrganization
)

type EligibleProductType = original.EligibleProductType

const (
	AzureReservation          EligibleProductType = original.AzureReservation
	DevTestAzureSubscription  EligibleProductType = original.DevTestAzureSubscription
	StandardAzureSubscription EligibleProductType = original.StandardAzureSubscription
)

type Frequency = original.Frequency

const (
	Monthly    Frequency = original.Monthly
	OneTime    Frequency = original.OneTime
	UsageBased Frequency = original.UsageBased
)

type Kind = original.Kind

const (
	CreditNote Kind = original.CreditNote
	Invoice    Kind = original.Invoice
	Receipt    Kind = original.Receipt
	VoidNote   Kind = original.VoidNote
)

type PaymentMethodType = original.PaymentMethodType

const (
	ChequeWire PaymentMethodType = original.ChequeWire
	Credits    PaymentMethodType = original.Credits
)

type ProductStatusType = original.ProductStatusType

const (
	Active    ProductStatusType = original.Active
	AutoRenew ProductStatusType = original.AutoRenew
	Cancelled ProductStatusType = original.Cancelled
	Disabled  ProductStatusType = original.Disabled
	Expired   ProductStatusType = original.Expired
	Expiring  ProductStatusType = original.Expiring
	Inactive  ProductStatusType = original.Inactive
	PastDue   ProductStatusType = original.PastDue
)

type ProductTransferStatus = original.ProductTransferStatus

const (
	Completed  ProductTransferStatus = original.Completed
	Failed     ProductTransferStatus = original.Failed
	InProgress ProductTransferStatus = original.InProgress
	NotStarted ProductTransferStatus = original.NotStarted
)

type ProductType = original.ProductType

const (
	ProductTypeAzureReservation  ProductType = original.ProductTypeAzureReservation
	ProductTypeAzureSubscription ProductType = original.ProductTypeAzureSubscription
)

type ReservationType = original.ReservationType

const (
	Purchase    ReservationType = original.Purchase
	UsageCharge ReservationType = original.UsageCharge
)

type Status = original.Status

const (
	StatusDue     Status = original.StatusDue
	StatusPaid    Status = original.StatusPaid
	StatusPastDue Status = original.StatusPastDue
	StatusVoid    Status = original.StatusVoid
)

type SubscriptionStatusType = original.SubscriptionStatusType

const (
	SubscriptionStatusTypeAbandoned SubscriptionStatusType = original.SubscriptionStatusTypeAbandoned
	SubscriptionStatusTypeActive    SubscriptionStatusType = original.SubscriptionStatusTypeActive
	SubscriptionStatusTypeDeleted   SubscriptionStatusType = original.SubscriptionStatusTypeDeleted
	SubscriptionStatusTypeInactive  SubscriptionStatusType = original.SubscriptionStatusTypeInactive
	SubscriptionStatusTypeWarning   SubscriptionStatusType = original.SubscriptionStatusTypeWarning
)

type TransactionTypeKind = original.TransactionTypeKind

const (
	All         TransactionTypeKind = original.All
	Reservation TransactionTypeKind = original.Reservation
)

type TransferStatus = original.TransferStatus

const (
	TransferStatusCanceled            TransferStatus = original.TransferStatusCanceled
	TransferStatusCompleted           TransferStatus = original.TransferStatusCompleted
	TransferStatusCompletedWithErrors TransferStatus = original.TransferStatusCompletedWithErrors
	TransferStatusDeclined            TransferStatus = original.TransferStatusDeclined
	TransferStatusFailed              TransferStatus = original.TransferStatusFailed
	TransferStatusInProgress          TransferStatus = original.TransferStatusInProgress
	TransferStatusPending             TransferStatus = original.TransferStatusPending
)

type UpdateAutoRenew = original.UpdateAutoRenew

const (
	False UpdateAutoRenew = original.False
	True  UpdateAutoRenew = original.True
)

type AcceptTransferProperties = original.AcceptTransferProperties
type AcceptTransferRequest = original.AcceptTransferRequest
type Account = original.Account
type AccountListResult = original.AccountListResult
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsUpdateFuture = original.AccountsUpdateFuture
type Address = original.Address
type Agreement = original.Agreement
type AgreementListResult = original.AgreementListResult
type AgreementProperties = original.AgreementProperties
type AgreementsClient = original.AgreementsClient
type Amount = original.Amount
type AvailableBalance = original.AvailableBalance
type AvailableBalanceProperties = original.AvailableBalanceProperties
type AvailableBalancesClient = original.AvailableBalancesClient
type BaseClient = original.BaseClient
type Department = original.Department
type DepartmentListResult = original.DepartmentListResult
type DepartmentProperties = original.DepartmentProperties
type DepartmentsClient = original.DepartmentsClient
type DetailedTransferStatus = original.DetailedTransferStatus
type DownloadProperties = original.DownloadProperties
type DownloadURL = original.DownloadURL
type EnabledAzureSKUs = original.EnabledAzureSKUs
type Enrollment = original.Enrollment
type EnrollmentAccount = original.EnrollmentAccount
type EnrollmentAccountContext = original.EnrollmentAccountContext
type EnrollmentAccountListResult = original.EnrollmentAccountListResult
type EnrollmentAccountProperties = original.EnrollmentAccountProperties
type EnrollmentAccountsClient = original.EnrollmentAccountsClient
type EnrollmentPolicies = original.EnrollmentPolicies
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type InitiateTransferProperties = original.InitiateTransferProperties
type InitiateTransferRequest = original.InitiateTransferRequest
type InvoiceListResult = original.InvoiceListResult
type InvoiceSection = original.InvoiceSection
type InvoiceSectionListResult = original.InvoiceSectionListResult
type InvoiceSectionProperties = original.InvoiceSectionProperties
type InvoiceSectionsClient = original.InvoiceSectionsClient
type InvoiceSectionsCreateFuture = original.InvoiceSectionsCreateFuture
type InvoiceSectionsUpdateFuture = original.InvoiceSectionsUpdateFuture
type InvoiceSummary = original.InvoiceSummary
type InvoiceSummaryProperties = original.InvoiceSummaryProperties
type InvoicesClient = original.InvoicesClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type OperationsClient = original.OperationsClient
type Participants = original.Participants
type PaymentMethod = original.PaymentMethod
type PaymentMethodProperties = original.PaymentMethodProperties
type PaymentMethodsClient = original.PaymentMethodsClient
type PaymentMethodsListResult = original.PaymentMethodsListResult
type PaymentMethodsListResultIterator = original.PaymentMethodsListResultIterator
type PaymentMethodsListResultPage = original.PaymentMethodsListResultPage
type PaymentProperties = original.PaymentProperties
type PermissionsClient = original.PermissionsClient
type PermissionsListResult = original.PermissionsListResult
type PermissionsProperties = original.PermissionsProperties
type PoliciesClient = original.PoliciesClient
type Policy = original.Policy
type PolicyProperties = original.PolicyProperties
type PriceSheetClient = original.PriceSheetClient
type PriceSheetDownloadFuture = original.PriceSheetDownloadFuture
type ProductDetails = original.ProductDetails
type ProductSummary = original.ProductSummary
type ProductSummaryProperties = original.ProductSummaryProperties
type ProductsClient = original.ProductsClient
type ProductsListResult = original.ProductsListResult
type ProductsListResultIterator = original.ProductsListResultIterator
type ProductsListResultPage = original.ProductsListResultPage
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileProperties = original.ProfileProperties
type ProfilesClient = original.ProfilesClient
type ProfilesUpdateFuture = original.ProfilesUpdateFuture
type Property = original.Property
type PropertyClient = original.PropertyClient
type PropertySummary = original.PropertySummary
type RecipientTransferDetails = original.RecipientTransferDetails
type RecipientTransferDetailsListResult = original.RecipientTransferDetailsListResult
type RecipientTransferDetailsListResultIterator = original.RecipientTransferDetailsListResultIterator
type RecipientTransferDetailsListResultPage = original.RecipientTransferDetailsListResultPage
type RecipientTransferProperties = original.RecipientTransferProperties
type RecipientTransfersClient = original.RecipientTransfersClient
type Resource = original.Resource
type RoleAssignment = original.RoleAssignment
type RoleAssignmentListResult = original.RoleAssignmentListResult
type RoleAssignmentPayload = original.RoleAssignmentPayload
type RoleAssignmentProperties = original.RoleAssignmentProperties
type RoleAssignmentsClient = original.RoleAssignmentsClient
type RoleDefinition = original.RoleDefinition
type RoleDefinitionListResult = original.RoleDefinitionListResult
type RoleDefinitionProperties = original.RoleDefinitionProperties
type RoleDefinitionsClient = original.RoleDefinitionsClient
type SubscriptionProperties = original.SubscriptionProperties
type SubscriptionSummary = original.SubscriptionSummary
type SubscriptionsClient = original.SubscriptionsClient
type SubscriptionsListResult = original.SubscriptionsListResult
type SubscriptionsListResultIterator = original.SubscriptionsListResultIterator
type SubscriptionsListResultPage = original.SubscriptionsListResultPage
type SubscriptionsTransferFuture = original.SubscriptionsTransferFuture
type TransactionsClient = original.TransactionsClient
type TransactionsListResult = original.TransactionsListResult
type TransactionsListResultIterator = original.TransactionsListResultIterator
type TransactionsListResultPage = original.TransactionsListResultPage
type TransactionsSummary = original.TransactionsSummary
type TransactionsSummaryProperties = original.TransactionsSummaryProperties
type TransferBillingSubscriptionRequest = original.TransferBillingSubscriptionRequest
type TransferBillingSubscriptionRequestProperties = original.TransferBillingSubscriptionRequestProperties
type TransferBillingSubscriptionResult = original.TransferBillingSubscriptionResult
type TransferBillingSubscriptionResultProperties = original.TransferBillingSubscriptionResultProperties
type TransferDetails = original.TransferDetails
type TransferDetailsListResult = original.TransferDetailsListResult
type TransferDetailsListResultIterator = original.TransferDetailsListResultIterator
type TransferDetailsListResultPage = original.TransferDetailsListResultPage
type TransferProductRequestProperties = original.TransferProductRequestProperties
type TransferProperties = original.TransferProperties
type TransfersClient = original.TransfersClient
type UpdateAutoRenewOperationSummary = original.UpdateAutoRenewOperationSummary
type UpdateAutoRenewOperationSummaryProperties = original.UpdateAutoRenewOperationSummaryProperties
type UpdateAutoRenewRequest = original.UpdateAutoRenewRequest

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAgreementsClient(subscriptionID string) AgreementsClient {
	return original.NewAgreementsClient(subscriptionID)
}
func NewAgreementsClientWithBaseURI(baseURI string, subscriptionID string) AgreementsClient {
	return original.NewAgreementsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableBalancesClient(subscriptionID string) AvailableBalancesClient {
	return original.NewAvailableBalancesClient(subscriptionID)
}
func NewAvailableBalancesClientWithBaseURI(baseURI string, subscriptionID string) AvailableBalancesClient {
	return original.NewAvailableBalancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDepartmentsClient(subscriptionID string) DepartmentsClient {
	return original.NewDepartmentsClient(subscriptionID)
}
func NewDepartmentsClientWithBaseURI(baseURI string, subscriptionID string) DepartmentsClient {
	return original.NewDepartmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnrollmentAccountsClient(subscriptionID string) EnrollmentAccountsClient {
	return original.NewEnrollmentAccountsClient(subscriptionID)
}
func NewEnrollmentAccountsClientWithBaseURI(baseURI string, subscriptionID string) EnrollmentAccountsClient {
	return original.NewEnrollmentAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoiceSectionsClient(subscriptionID string) InvoiceSectionsClient {
	return original.NewInvoiceSectionsClient(subscriptionID)
}
func NewInvoiceSectionsClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsClient {
	return original.NewInvoiceSectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicesClient(subscriptionID string) InvoicesClient {
	return original.NewInvoicesClient(subscriptionID)
}
func NewInvoicesClientWithBaseURI(baseURI string, subscriptionID string) InvoicesClient {
	return original.NewInvoicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPaymentMethodsClient(subscriptionID string) PaymentMethodsClient {
	return original.NewPaymentMethodsClient(subscriptionID)
}
func NewPaymentMethodsClientWithBaseURI(baseURI string, subscriptionID string) PaymentMethodsClient {
	return original.NewPaymentMethodsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPaymentMethodsListResultIterator(page PaymentMethodsListResultPage) PaymentMethodsListResultIterator {
	return original.NewPaymentMethodsListResultIterator(page)
}
func NewPaymentMethodsListResultPage(getNextPage func(context.Context, PaymentMethodsListResult) (PaymentMethodsListResult, error)) PaymentMethodsListResultPage {
	return original.NewPaymentMethodsListResultPage(getNextPage)
}
func NewPermissionsClient(subscriptionID string) PermissionsClient {
	return original.NewPermissionsClient(subscriptionID)
}
func NewPermissionsClientWithBaseURI(baseURI string, subscriptionID string) PermissionsClient {
	return original.NewPermissionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClient(subscriptionID)
}
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClientWithBaseURI(baseURI, subscriptionID)
}
func NewProductsClient(subscriptionID string) ProductsClient {
	return original.NewProductsClient(subscriptionID)
}
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return original.NewProductsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProductsListResultIterator(page ProductsListResultPage) ProductsListResultIterator {
	return original.NewProductsListResultIterator(page)
}
func NewProductsListResultPage(getNextPage func(context.Context, ProductsListResult) (ProductsListResult, error)) ProductsListResultPage {
	return original.NewProductsListResultPage(getNextPage)
}
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPropertyClient(subscriptionID string) PropertyClient {
	return original.NewPropertyClient(subscriptionID)
}
func NewPropertyClientWithBaseURI(baseURI string, subscriptionID string) PropertyClient {
	return original.NewPropertyClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecipientTransferDetailsListResultIterator(page RecipientTransferDetailsListResultPage) RecipientTransferDetailsListResultIterator {
	return original.NewRecipientTransferDetailsListResultIterator(page)
}
func NewRecipientTransferDetailsListResultPage(getNextPage func(context.Context, RecipientTransferDetailsListResult) (RecipientTransferDetailsListResult, error)) RecipientTransferDetailsListResultPage {
	return original.NewRecipientTransferDetailsListResultPage(getNextPage)
}
func NewRecipientTransfersClient(subscriptionID string) RecipientTransfersClient {
	return original.NewRecipientTransfersClient(subscriptionID)
}
func NewRecipientTransfersClientWithBaseURI(baseURI string, subscriptionID string) RecipientTransfersClient {
	return original.NewRecipientTransfersClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleAssignmentsClient(subscriptionID string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClient(subscriptionID)
}
func NewRoleAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentsClient {
	return original.NewRoleAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoleDefinitionsClient(subscriptionID string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClient(subscriptionID)
}
func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) RoleDefinitionsClient {
	return original.NewRoleDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsListResultIterator(page SubscriptionsListResultPage) SubscriptionsListResultIterator {
	return original.NewSubscriptionsListResultIterator(page)
}
func NewSubscriptionsListResultPage(getNextPage func(context.Context, SubscriptionsListResult) (SubscriptionsListResult, error)) SubscriptionsListResultPage {
	return original.NewSubscriptionsListResultPage(getNextPage)
}
func NewTransactionsClient(subscriptionID string) TransactionsClient {
	return original.NewTransactionsClient(subscriptionID)
}
func NewTransactionsClientWithBaseURI(baseURI string, subscriptionID string) TransactionsClient {
	return original.NewTransactionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransactionsListResultIterator(page TransactionsListResultPage) TransactionsListResultIterator {
	return original.NewTransactionsListResultIterator(page)
}
func NewTransactionsListResultPage(getNextPage func(context.Context, TransactionsListResult) (TransactionsListResult, error)) TransactionsListResultPage {
	return original.NewTransactionsListResultPage(getNextPage)
}
func NewTransferDetailsListResultIterator(page TransferDetailsListResultPage) TransferDetailsListResultIterator {
	return original.NewTransferDetailsListResultIterator(page)
}
func NewTransferDetailsListResultPage(getNextPage func(context.Context, TransferDetailsListResult) (TransferDetailsListResult, error)) TransferDetailsListResultPage {
	return original.NewTransferDetailsListResultPage(getNextPage)
}
func NewTransfersClient(subscriptionID string) TransfersClient {
	return original.NewTransfersClient(subscriptionID)
}
func NewTransfersClientWithBaseURI(baseURI string, subscriptionID string) TransfersClient {
	return original.NewTransfersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccountTypeValues() []AccountType {
	return original.PossibleAccountTypeValues()
}
func PossibleEligibleProductTypeValues() []EligibleProductType {
	return original.PossibleEligibleProductTypeValues()
}
func PossibleFrequencyValues() []Frequency {
	return original.PossibleFrequencyValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossiblePaymentMethodTypeValues() []PaymentMethodType {
	return original.PossiblePaymentMethodTypeValues()
}
func PossibleProductStatusTypeValues() []ProductStatusType {
	return original.PossibleProductStatusTypeValues()
}
func PossibleProductTransferStatusValues() []ProductTransferStatus {
	return original.PossibleProductTransferStatusValues()
}
func PossibleProductTypeValues() []ProductType {
	return original.PossibleProductTypeValues()
}
func PossibleReservationTypeValues() []ReservationType {
	return original.PossibleReservationTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleSubscriptionStatusTypeValues() []SubscriptionStatusType {
	return original.PossibleSubscriptionStatusTypeValues()
}
func PossibleTransactionTypeKindValues() []TransactionTypeKind {
	return original.PossibleTransactionTypeKindValues()
}
func PossibleTransferStatusValues() []TransferStatus {
	return original.PossibleTransferStatusValues()
}
func PossibleUpdateAutoRenewValues() []UpdateAutoRenew {
	return original.PossibleUpdateAutoRenewValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
