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

type UpdateAutoRenew = original.UpdateAutoRenew

const (
	False UpdateAutoRenew = original.False
	True  UpdateAutoRenew = original.True
)

type Account = original.Account
type AccountListResult = original.AccountListResult
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsWithCreateInvoiceSectionPermissionClient = original.AccountsWithCreateInvoiceSectionPermissionClient
type Address = original.Address
type Amount = original.Amount
type AvailableBalance = original.AvailableBalance
type AvailableBalanceByBillingProfileClient = original.AvailableBalanceByBillingProfileClient
type AvailableBalanceProperties = original.AvailableBalanceProperties
type BaseClient = original.BaseClient
type Department = original.Department
type DepartmentListResult = original.DepartmentListResult
type DepartmentProperties = original.DepartmentProperties
type DepartmentsByBillingAccountIDClient = original.DepartmentsByBillingAccountIDClient
type DepartmentsClient = original.DepartmentsClient
type DownloadProperties = original.DownloadProperties
type DownloadURL = original.DownloadURL
type EnabledAzureSKUs = original.EnabledAzureSKUs
type Enrollment = original.Enrollment
type EnrollmentAccount = original.EnrollmentAccount
type EnrollmentAccountContext = original.EnrollmentAccountContext
type EnrollmentAccountListResult = original.EnrollmentAccountListResult
type EnrollmentAccountProperties = original.EnrollmentAccountProperties
type EnrollmentAccountsByBillingAccountIDClient = original.EnrollmentAccountsByBillingAccountIDClient
type EnrollmentAccountsClient = original.EnrollmentAccountsClient
type EnrollmentPolicies = original.EnrollmentPolicies
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type InvoiceClient = original.InvoiceClient
type InvoiceListResult = original.InvoiceListResult
type InvoicePricesheetClient = original.InvoicePricesheetClient
type InvoicePricesheetDownloadFuture = original.InvoicePricesheetDownloadFuture
type InvoiceSection = original.InvoiceSection
type InvoiceSectionListResult = original.InvoiceSectionListResult
type InvoiceSectionProperties = original.InvoiceSectionProperties
type InvoiceSectionsByBillingAccountIDClient = original.InvoiceSectionsByBillingAccountIDClient
type InvoiceSectionsClient = original.InvoiceSectionsClient
type InvoiceSectionsCreateFuture = original.InvoiceSectionsCreateFuture
type InvoiceSectionsUpdateFuture = original.InvoiceSectionsUpdateFuture
type InvoiceSectionsWithCreateSubscriptionPermissionClient = original.InvoiceSectionsWithCreateSubscriptionPermissionClient
type InvoiceSummary = original.InvoiceSummary
type InvoiceSummaryProperties = original.InvoiceSummaryProperties
type InvoicesByBillingAccountClient = original.InvoicesByBillingAccountClient
type InvoicesByBillingProfileClient = original.InvoicesByBillingProfileClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type OperationsClient = original.OperationsClient
type PaymentMethod = original.PaymentMethod
type PaymentMethodProperties = original.PaymentMethodProperties
type PaymentMethodsByBillingProfileClient = original.PaymentMethodsByBillingProfileClient
type PaymentMethodsListResult = original.PaymentMethodsListResult
type PaymentMethodsListResultIterator = original.PaymentMethodsListResultIterator
type PaymentMethodsListResultPage = original.PaymentMethodsListResultPage
type PaymentProperties = original.PaymentProperties
type Policy = original.Policy
type PolicyClient = original.PolicyClient
type PolicyProperties = original.PolicyProperties
type ProductSummary = original.ProductSummary
type ProductSummaryProperties = original.ProductSummaryProperties
type ProductsByBillingAccountClient = original.ProductsByBillingAccountClient
type ProductsByBillingSubscriptionsClient = original.ProductsByBillingSubscriptionsClient
type ProductsByInvoiceSectionClient = original.ProductsByInvoiceSectionClient
type ProductsClient = original.ProductsClient
type ProductsListResult = original.ProductsListResult
type ProductsListResultIterator = original.ProductsListResultIterator
type ProductsListResultPage = original.ProductsListResultPage
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileProperties = original.ProfileProperties
type ProfilesByBillingAccountIDClient = original.ProfilesByBillingAccountIDClient
type ProfilesClient = original.ProfilesClient
type ProfilesUpdateFuture = original.ProfilesUpdateFuture
type Property = original.Property
type PropertyClient = original.PropertyClient
type Resource = original.Resource
type SubscriptionClient = original.SubscriptionClient
type SubscriptionProperties = original.SubscriptionProperties
type SubscriptionSummary = original.SubscriptionSummary
type SubscriptionTransferFuture = original.SubscriptionTransferFuture
type SubscriptionsByBillingProfileClient = original.SubscriptionsByBillingProfileClient
type SubscriptionsByInvoiceSectionClient = original.SubscriptionsByInvoiceSectionClient
type SubscriptionsListResult = original.SubscriptionsListResult
type SubscriptionsListResultIterator = original.SubscriptionsListResultIterator
type SubscriptionsListResultPage = original.SubscriptionsListResultPage
type TransactionsByBillingAccountClient = original.TransactionsByBillingAccountClient
type TransactionsListResult = original.TransactionsListResult
type TransactionsListResultIterator = original.TransactionsListResultIterator
type TransactionsListResultPage = original.TransactionsListResultPage
type TransactionsSummary = original.TransactionsSummary
type TransactionsSummaryProperties = original.TransactionsSummaryProperties
type TransferBillingSubscriptionRequest = original.TransferBillingSubscriptionRequest
type TransferBillingSubscriptionRequestProperties = original.TransferBillingSubscriptionRequestProperties
type TransferBillingSubscriptionResult = original.TransferBillingSubscriptionResult
type TransferBillingSubscriptionResultProperties = original.TransferBillingSubscriptionResultProperties
type TransferProductRequestProperties = original.TransferProductRequestProperties
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
func NewAccountsWithCreateInvoiceSectionPermissionClient(subscriptionID string) AccountsWithCreateInvoiceSectionPermissionClient {
	return original.NewAccountsWithCreateInvoiceSectionPermissionClient(subscriptionID)
}
func NewAccountsWithCreateInvoiceSectionPermissionClientWithBaseURI(baseURI string, subscriptionID string) AccountsWithCreateInvoiceSectionPermissionClient {
	return original.NewAccountsWithCreateInvoiceSectionPermissionClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableBalanceByBillingProfileClient(subscriptionID string) AvailableBalanceByBillingProfileClient {
	return original.NewAvailableBalanceByBillingProfileClient(subscriptionID)
}
func NewAvailableBalanceByBillingProfileClientWithBaseURI(baseURI string, subscriptionID string) AvailableBalanceByBillingProfileClient {
	return original.NewAvailableBalanceByBillingProfileClientWithBaseURI(baseURI, subscriptionID)
}
func NewDepartmentsByBillingAccountIDClient(subscriptionID string) DepartmentsByBillingAccountIDClient {
	return original.NewDepartmentsByBillingAccountIDClient(subscriptionID)
}
func NewDepartmentsByBillingAccountIDClientWithBaseURI(baseURI string, subscriptionID string) DepartmentsByBillingAccountIDClient {
	return original.NewDepartmentsByBillingAccountIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewDepartmentsClient(subscriptionID string) DepartmentsClient {
	return original.NewDepartmentsClient(subscriptionID)
}
func NewDepartmentsClientWithBaseURI(baseURI string, subscriptionID string) DepartmentsClient {
	return original.NewDepartmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnrollmentAccountsByBillingAccountIDClient(subscriptionID string) EnrollmentAccountsByBillingAccountIDClient {
	return original.NewEnrollmentAccountsByBillingAccountIDClient(subscriptionID)
}
func NewEnrollmentAccountsByBillingAccountIDClientWithBaseURI(baseURI string, subscriptionID string) EnrollmentAccountsByBillingAccountIDClient {
	return original.NewEnrollmentAccountsByBillingAccountIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewEnrollmentAccountsClient(subscriptionID string) EnrollmentAccountsClient {
	return original.NewEnrollmentAccountsClient(subscriptionID)
}
func NewEnrollmentAccountsClientWithBaseURI(baseURI string, subscriptionID string) EnrollmentAccountsClient {
	return original.NewEnrollmentAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoiceClient(subscriptionID string) InvoiceClient {
	return original.NewInvoiceClient(subscriptionID)
}
func NewInvoiceClientWithBaseURI(baseURI string, subscriptionID string) InvoiceClient {
	return original.NewInvoiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicePricesheetClient(subscriptionID string) InvoicePricesheetClient {
	return original.NewInvoicePricesheetClient(subscriptionID)
}
func NewInvoicePricesheetClientWithBaseURI(baseURI string, subscriptionID string) InvoicePricesheetClient {
	return original.NewInvoicePricesheetClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoiceSectionsByBillingAccountIDClient(subscriptionID string) InvoiceSectionsByBillingAccountIDClient {
	return original.NewInvoiceSectionsByBillingAccountIDClient(subscriptionID)
}
func NewInvoiceSectionsByBillingAccountIDClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsByBillingAccountIDClient {
	return original.NewInvoiceSectionsByBillingAccountIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoiceSectionsClient(subscriptionID string) InvoiceSectionsClient {
	return original.NewInvoiceSectionsClient(subscriptionID)
}
func NewInvoiceSectionsClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsClient {
	return original.NewInvoiceSectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoiceSectionsWithCreateSubscriptionPermissionClient(subscriptionID string) InvoiceSectionsWithCreateSubscriptionPermissionClient {
	return original.NewInvoiceSectionsWithCreateSubscriptionPermissionClient(subscriptionID)
}
func NewInvoiceSectionsWithCreateSubscriptionPermissionClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsWithCreateSubscriptionPermissionClient {
	return original.NewInvoiceSectionsWithCreateSubscriptionPermissionClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicesByBillingAccountClient(subscriptionID string) InvoicesByBillingAccountClient {
	return original.NewInvoicesByBillingAccountClient(subscriptionID)
}
func NewInvoicesByBillingAccountClientWithBaseURI(baseURI string, subscriptionID string) InvoicesByBillingAccountClient {
	return original.NewInvoicesByBillingAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicesByBillingProfileClient(subscriptionID string) InvoicesByBillingProfileClient {
	return original.NewInvoicesByBillingProfileClient(subscriptionID)
}
func NewInvoicesByBillingProfileClientWithBaseURI(baseURI string, subscriptionID string) InvoicesByBillingProfileClient {
	return original.NewInvoicesByBillingProfileClientWithBaseURI(baseURI, subscriptionID)
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
func NewPaymentMethodsByBillingProfileClient(subscriptionID string) PaymentMethodsByBillingProfileClient {
	return original.NewPaymentMethodsByBillingProfileClient(subscriptionID)
}
func NewPaymentMethodsByBillingProfileClientWithBaseURI(baseURI string, subscriptionID string) PaymentMethodsByBillingProfileClient {
	return original.NewPaymentMethodsByBillingProfileClientWithBaseURI(baseURI, subscriptionID)
}
func NewPaymentMethodsListResultIterator(page PaymentMethodsListResultPage) PaymentMethodsListResultIterator {
	return original.NewPaymentMethodsListResultIterator(page)
}
func NewPaymentMethodsListResultPage(getNextPage func(context.Context, PaymentMethodsListResult) (PaymentMethodsListResult, error)) PaymentMethodsListResultPage {
	return original.NewPaymentMethodsListResultPage(getNextPage)
}
func NewPolicyClient(subscriptionID string) PolicyClient {
	return original.NewPolicyClient(subscriptionID)
}
func NewPolicyClientWithBaseURI(baseURI string, subscriptionID string) PolicyClient {
	return original.NewPolicyClientWithBaseURI(baseURI, subscriptionID)
}
func NewProductsByBillingAccountClient(subscriptionID string) ProductsByBillingAccountClient {
	return original.NewProductsByBillingAccountClient(subscriptionID)
}
func NewProductsByBillingAccountClientWithBaseURI(baseURI string, subscriptionID string) ProductsByBillingAccountClient {
	return original.NewProductsByBillingAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewProductsByBillingSubscriptionsClient(subscriptionID string) ProductsByBillingSubscriptionsClient {
	return original.NewProductsByBillingSubscriptionsClient(subscriptionID)
}
func NewProductsByBillingSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) ProductsByBillingSubscriptionsClient {
	return original.NewProductsByBillingSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProductsByInvoiceSectionClient(subscriptionID string) ProductsByInvoiceSectionClient {
	return original.NewProductsByInvoiceSectionClient(subscriptionID)
}
func NewProductsByInvoiceSectionClientWithBaseURI(baseURI string, subscriptionID string) ProductsByInvoiceSectionClient {
	return original.NewProductsByInvoiceSectionClientWithBaseURI(baseURI, subscriptionID)
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
func NewProfilesByBillingAccountIDClient(subscriptionID string) ProfilesByBillingAccountIDClient {
	return original.NewProfilesByBillingAccountIDClient(subscriptionID)
}
func NewProfilesByBillingAccountIDClientWithBaseURI(baseURI string, subscriptionID string) ProfilesByBillingAccountIDClient {
	return original.NewProfilesByBillingAccountIDClientWithBaseURI(baseURI, subscriptionID)
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
func NewSubscriptionClient(subscriptionID string) SubscriptionClient {
	return original.NewSubscriptionClient(subscriptionID)
}
func NewSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionClient {
	return original.NewSubscriptionClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsByBillingProfileClient(subscriptionID string) SubscriptionsByBillingProfileClient {
	return original.NewSubscriptionsByBillingProfileClient(subscriptionID)
}
func NewSubscriptionsByBillingProfileClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsByBillingProfileClient {
	return original.NewSubscriptionsByBillingProfileClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsByInvoiceSectionClient(subscriptionID string) SubscriptionsByInvoiceSectionClient {
	return original.NewSubscriptionsByInvoiceSectionClient(subscriptionID)
}
func NewSubscriptionsByInvoiceSectionClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsByInvoiceSectionClient {
	return original.NewSubscriptionsByInvoiceSectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsListResultIterator(page SubscriptionsListResultPage) SubscriptionsListResultIterator {
	return original.NewSubscriptionsListResultIterator(page)
}
func NewSubscriptionsListResultPage(getNextPage func(context.Context, SubscriptionsListResult) (SubscriptionsListResult, error)) SubscriptionsListResultPage {
	return original.NewSubscriptionsListResultPage(getNextPage)
}
func NewTransactionsByBillingAccountClient(subscriptionID string) TransactionsByBillingAccountClient {
	return original.NewTransactionsByBillingAccountClient(subscriptionID)
}
func NewTransactionsByBillingAccountClientWithBaseURI(baseURI string, subscriptionID string) TransactionsByBillingAccountClient {
	return original.NewTransactionsByBillingAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransactionsListResultIterator(page TransactionsListResultPage) TransactionsListResultIterator {
	return original.NewTransactionsListResultIterator(page)
}
func NewTransactionsListResultPage(getNextPage func(context.Context, TransactionsListResult) (TransactionsListResult, error)) TransactionsListResultPage {
	return original.NewTransactionsListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccountTypeValues() []AccountType {
	return original.PossibleAccountTypeValues()
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
func PossibleUpdateAutoRenewValues() []UpdateAutoRenew {
	return original.PossibleUpdateAutoRenewValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
