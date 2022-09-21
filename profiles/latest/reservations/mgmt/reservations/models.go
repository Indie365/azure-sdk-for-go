//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package reservations

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/reservations/mgmt/2022-03-01/reservations"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppliedScopeType = original.AppliedScopeType

const (
	AppliedScopeTypeShared AppliedScopeType = original.AppliedScopeTypeShared
	AppliedScopeTypeSingle AppliedScopeType = original.AppliedScopeTypeSingle
)

type CalculateExchangeOperationResultStatus = original.CalculateExchangeOperationResultStatus

const (
	CalculateExchangeOperationResultStatusCancelled CalculateExchangeOperationResultStatus = original.CalculateExchangeOperationResultStatusCancelled
	CalculateExchangeOperationResultStatusFailed    CalculateExchangeOperationResultStatus = original.CalculateExchangeOperationResultStatusFailed
	CalculateExchangeOperationResultStatusPending   CalculateExchangeOperationResultStatus = original.CalculateExchangeOperationResultStatusPending
	CalculateExchangeOperationResultStatusSucceeded CalculateExchangeOperationResultStatus = original.CalculateExchangeOperationResultStatusSucceeded
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DisplayProvisioningState = original.DisplayProvisioningState

const (
	DisplayProvisioningStateCancelled  DisplayProvisioningState = original.DisplayProvisioningStateCancelled
	DisplayProvisioningStateExpired    DisplayProvisioningState = original.DisplayProvisioningStateExpired
	DisplayProvisioningStateExpiring   DisplayProvisioningState = original.DisplayProvisioningStateExpiring
	DisplayProvisioningStateFailed     DisplayProvisioningState = original.DisplayProvisioningStateFailed
	DisplayProvisioningStatePending    DisplayProvisioningState = original.DisplayProvisioningStatePending
	DisplayProvisioningStateProcessing DisplayProvisioningState = original.DisplayProvisioningStateProcessing
	DisplayProvisioningStateSucceeded  DisplayProvisioningState = original.DisplayProvisioningStateSucceeded
)

type ErrorResponseCode = original.ErrorResponseCode

const (
	ErrorResponseCodeActivateQuoteFailed                           ErrorResponseCode = original.ErrorResponseCodeActivateQuoteFailed
	ErrorResponseCodeAppliedScopesNotAssociatedWithCommerceAccount ErrorResponseCode = original.ErrorResponseCodeAppliedScopesNotAssociatedWithCommerceAccount
	ErrorResponseCodeAppliedScopesSameAsExisting                   ErrorResponseCode = original.ErrorResponseCodeAppliedScopesSameAsExisting
	ErrorResponseCodeAuthorizationFailed                           ErrorResponseCode = original.ErrorResponseCodeAuthorizationFailed
	ErrorResponseCodeBadRequest                                    ErrorResponseCode = original.ErrorResponseCodeBadRequest
	ErrorResponseCodeBillingCustomerInputError                     ErrorResponseCode = original.ErrorResponseCodeBillingCustomerInputError
	ErrorResponseCodeBillingError                                  ErrorResponseCode = original.ErrorResponseCodeBillingError
	ErrorResponseCodeBillingPaymentInstrumentHardError             ErrorResponseCode = original.ErrorResponseCodeBillingPaymentInstrumentHardError
	ErrorResponseCodeBillingPaymentInstrumentSoftError             ErrorResponseCode = original.ErrorResponseCodeBillingPaymentInstrumentSoftError
	ErrorResponseCodeBillingScopeIDCannotBeChanged                 ErrorResponseCode = original.ErrorResponseCodeBillingScopeIDCannotBeChanged
	ErrorResponseCodeBillingTransientError                         ErrorResponseCode = original.ErrorResponseCodeBillingTransientError
	ErrorResponseCodeCalculatePriceFailed                          ErrorResponseCode = original.ErrorResponseCodeCalculatePriceFailed
	ErrorResponseCodeCapacityUpdateScopesFailed                    ErrorResponseCode = original.ErrorResponseCodeCapacityUpdateScopesFailed
	ErrorResponseCodeClientCertificateThumbprintNotSet             ErrorResponseCode = original.ErrorResponseCodeClientCertificateThumbprintNotSet
	ErrorResponseCodeCreateQuoteFailed                             ErrorResponseCode = original.ErrorResponseCodeCreateQuoteFailed
	ErrorResponseCodeForbidden                                     ErrorResponseCode = original.ErrorResponseCodeForbidden
	ErrorResponseCodeFulfillmentConfigurationError                 ErrorResponseCode = original.ErrorResponseCodeFulfillmentConfigurationError
	ErrorResponseCodeFulfillmentError                              ErrorResponseCode = original.ErrorResponseCodeFulfillmentError
	ErrorResponseCodeFulfillmentOutOfStockError                    ErrorResponseCode = original.ErrorResponseCodeFulfillmentOutOfStockError
	ErrorResponseCodeFulfillmentTransientError                     ErrorResponseCode = original.ErrorResponseCodeFulfillmentTransientError
	ErrorResponseCodeHTTPMethodNotSupported                        ErrorResponseCode = original.ErrorResponseCodeHTTPMethodNotSupported
	ErrorResponseCodeInternalServerError                           ErrorResponseCode = original.ErrorResponseCodeInternalServerError
	ErrorResponseCodeInvalidAccessToken                            ErrorResponseCode = original.ErrorResponseCodeInvalidAccessToken
	ErrorResponseCodeInvalidFulfillmentRequestParameters           ErrorResponseCode = original.ErrorResponseCodeInvalidFulfillmentRequestParameters
	ErrorResponseCodeInvalidHealthCheckType                        ErrorResponseCode = original.ErrorResponseCodeInvalidHealthCheckType
	ErrorResponseCodeInvalidLocationID                             ErrorResponseCode = original.ErrorResponseCodeInvalidLocationID
	ErrorResponseCodeInvalidRefundQuantity                         ErrorResponseCode = original.ErrorResponseCodeInvalidRefundQuantity
	ErrorResponseCodeInvalidRequestContent                         ErrorResponseCode = original.ErrorResponseCodeInvalidRequestContent
	ErrorResponseCodeInvalidRequestURI                             ErrorResponseCode = original.ErrorResponseCodeInvalidRequestURI
	ErrorResponseCodeInvalidReservationID                          ErrorResponseCode = original.ErrorResponseCodeInvalidReservationID
	ErrorResponseCodeInvalidReservationOrderID                     ErrorResponseCode = original.ErrorResponseCodeInvalidReservationOrderID
	ErrorResponseCodeInvalidSingleAppliedScopesCount               ErrorResponseCode = original.ErrorResponseCodeInvalidSingleAppliedScopesCount
	ErrorResponseCodeInvalidSubscriptionID                         ErrorResponseCode = original.ErrorResponseCodeInvalidSubscriptionID
	ErrorResponseCodeInvalidTenantID                               ErrorResponseCode = original.ErrorResponseCodeInvalidTenantID
	ErrorResponseCodeMissingAppliedScopesForSingle                 ErrorResponseCode = original.ErrorResponseCodeMissingAppliedScopesForSingle
	ErrorResponseCodeMissingTenantID                               ErrorResponseCode = original.ErrorResponseCodeMissingTenantID
	ErrorResponseCodeNonsupportedAccountID                         ErrorResponseCode = original.ErrorResponseCodeNonsupportedAccountID
	ErrorResponseCodeNotSpecified                                  ErrorResponseCode = original.ErrorResponseCodeNotSpecified
	ErrorResponseCodeNotSupportedCountry                           ErrorResponseCode = original.ErrorResponseCodeNotSupportedCountry
	ErrorResponseCodeNoValidReservationsToReRate                   ErrorResponseCode = original.ErrorResponseCodeNoValidReservationsToReRate
	ErrorResponseCodeOperationCannotBePerformedInCurrentState      ErrorResponseCode = original.ErrorResponseCodeOperationCannotBePerformedInCurrentState
	ErrorResponseCodeOperationFailed                               ErrorResponseCode = original.ErrorResponseCodeOperationFailed
	ErrorResponseCodePatchValuesSameAsExisting                     ErrorResponseCode = original.ErrorResponseCodePatchValuesSameAsExisting
	ErrorResponseCodePaymentInstrumentNotFound                     ErrorResponseCode = original.ErrorResponseCodePaymentInstrumentNotFound
	ErrorResponseCodePurchaseError                                 ErrorResponseCode = original.ErrorResponseCodePurchaseError
	ErrorResponseCodeRefundLimitExceeded                           ErrorResponseCode = original.ErrorResponseCodeRefundLimitExceeded
	ErrorResponseCodeReRateOnlyAllowedForEA                        ErrorResponseCode = original.ErrorResponseCodeReRateOnlyAllowedForEA
	ErrorResponseCodeReservationIDNotInReservationOrder            ErrorResponseCode = original.ErrorResponseCodeReservationIDNotInReservationOrder
	ErrorResponseCodeReservationOrderCreationFailed                ErrorResponseCode = original.ErrorResponseCodeReservationOrderCreationFailed
	ErrorResponseCodeReservationOrderIDAlreadyExists               ErrorResponseCode = original.ErrorResponseCodeReservationOrderIDAlreadyExists
	ErrorResponseCodeReservationOrderNotEnabled                    ErrorResponseCode = original.ErrorResponseCodeReservationOrderNotEnabled
	ErrorResponseCodeReservationOrderNotFound                      ErrorResponseCode = original.ErrorResponseCodeReservationOrderNotFound
	ErrorResponseCodeRiskCheckFailed                               ErrorResponseCode = original.ErrorResponseCodeRiskCheckFailed
	ErrorResponseCodeRoleAssignmentCreationFailed                  ErrorResponseCode = original.ErrorResponseCodeRoleAssignmentCreationFailed
	ErrorResponseCodeSelfServiceRefundNotSupported                 ErrorResponseCode = original.ErrorResponseCodeSelfServiceRefundNotSupported
	ErrorResponseCodeServerTimeout                                 ErrorResponseCode = original.ErrorResponseCodeServerTimeout
	ErrorResponseCodeUnauthenticatedRequestsThrottled              ErrorResponseCode = original.ErrorResponseCodeUnauthenticatedRequestsThrottled
	ErrorResponseCodeUnsupportedReservationTerm                    ErrorResponseCode = original.ErrorResponseCodeUnsupportedReservationTerm
)

type ExchangeOperationResultStatus = original.ExchangeOperationResultStatus

const (
	ExchangeOperationResultStatusCancelled        ExchangeOperationResultStatus = original.ExchangeOperationResultStatusCancelled
	ExchangeOperationResultStatusFailed           ExchangeOperationResultStatus = original.ExchangeOperationResultStatusFailed
	ExchangeOperationResultStatusPendingPurchases ExchangeOperationResultStatus = original.ExchangeOperationResultStatusPendingPurchases
	ExchangeOperationResultStatusPendingRefunds   ExchangeOperationResultStatus = original.ExchangeOperationResultStatusPendingRefunds
	ExchangeOperationResultStatusSucceeded        ExchangeOperationResultStatus = original.ExchangeOperationResultStatusSucceeded
)

type InstanceFlexibility = original.InstanceFlexibility

const (
	InstanceFlexibilityOff InstanceFlexibility = original.InstanceFlexibilityOff
	InstanceFlexibilityOn  InstanceFlexibility = original.InstanceFlexibilityOn
)

type Kind = original.Kind

const (
	KindMicrosoftCompute Kind = original.KindMicrosoftCompute
)

type Location = original.Location

const (
	LocationAustraliaeast      Location = original.LocationAustraliaeast
	LocationAustraliasoutheast Location = original.LocationAustraliasoutheast
	LocationBrazilsouth        Location = original.LocationBrazilsouth
	LocationCanadacentral      Location = original.LocationCanadacentral
	LocationCanadaeast         Location = original.LocationCanadaeast
	LocationCentralindia       Location = original.LocationCentralindia
	LocationCentralus          Location = original.LocationCentralus
	LocationEastasia           Location = original.LocationEastasia
	LocationEastus             Location = original.LocationEastus
	LocationEastus2            Location = original.LocationEastus2
	LocationJapaneast          Location = original.LocationJapaneast
	LocationJapanwest          Location = original.LocationJapanwest
	LocationNorthcentralus     Location = original.LocationNorthcentralus
	LocationNortheurope        Location = original.LocationNortheurope
	LocationSouthcentralus     Location = original.LocationSouthcentralus
	LocationSoutheastasia      Location = original.LocationSoutheastasia
	LocationSouthindia         Location = original.LocationSouthindia
	LocationUksouth            Location = original.LocationUksouth
	LocationUkwest             Location = original.LocationUkwest
	LocationWestcentralus      Location = original.LocationWestcentralus
	LocationWesteurope         Location = original.LocationWesteurope
	LocationWestindia          Location = original.LocationWestindia
	LocationWestus             Location = original.LocationWestus
	LocationWestus2            Location = original.LocationWestus2
)

type OperationStatus = original.OperationStatus

const (
	OperationStatusCancelled OperationStatus = original.OperationStatusCancelled
	OperationStatusFailed    OperationStatus = original.OperationStatusFailed
	OperationStatusPending   OperationStatus = original.OperationStatusPending
	OperationStatusSucceeded OperationStatus = original.OperationStatusSucceeded
)

type PaymentStatus = original.PaymentStatus

const (
	PaymentStatusCancelled PaymentStatus = original.PaymentStatusCancelled
	PaymentStatusFailed    PaymentStatus = original.PaymentStatusFailed
	PaymentStatusScheduled PaymentStatus = original.PaymentStatusScheduled
	PaymentStatusSucceeded PaymentStatus = original.PaymentStatusSucceeded
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateBillingFailed         ProvisioningState = original.ProvisioningStateBillingFailed
	ProvisioningStateCancelled             ProvisioningState = original.ProvisioningStateCancelled
	ProvisioningStateConfirmedBilling      ProvisioningState = original.ProvisioningStateConfirmedBilling
	ProvisioningStateConfirmedResourceHold ProvisioningState = original.ProvisioningStateConfirmedResourceHold
	ProvisioningStateCreated               ProvisioningState = original.ProvisioningStateCreated
	ProvisioningStateCreating              ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateExpired               ProvisioningState = original.ProvisioningStateExpired
	ProvisioningStateFailed                ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMerged                ProvisioningState = original.ProvisioningStateMerged
	ProvisioningStatePendingBilling        ProvisioningState = original.ProvisioningStatePendingBilling
	ProvisioningStatePendingResourceHold   ProvisioningState = original.ProvisioningStatePendingResourceHold
	ProvisioningStateSplit                 ProvisioningState = original.ProvisioningStateSplit
	ProvisioningStateSucceeded             ProvisioningState = original.ProvisioningStateSucceeded
)

type QuotaRequestState = original.QuotaRequestState

const (
	QuotaRequestStateAccepted   QuotaRequestState = original.QuotaRequestStateAccepted
	QuotaRequestStateFailed     QuotaRequestState = original.QuotaRequestStateFailed
	QuotaRequestStateInProgress QuotaRequestState = original.QuotaRequestStateInProgress
	QuotaRequestStateInvalid    QuotaRequestState = original.QuotaRequestStateInvalid
	QuotaRequestStateSucceeded  QuotaRequestState = original.QuotaRequestStateSucceeded
)

type ReservationBillingPlan = original.ReservationBillingPlan

const (
	ReservationBillingPlanMonthly ReservationBillingPlan = original.ReservationBillingPlanMonthly
	ReservationBillingPlanUpfront ReservationBillingPlan = original.ReservationBillingPlanUpfront
)

type ReservationTerm = original.ReservationTerm

const (
	ReservationTermP1Y ReservationTerm = original.ReservationTermP1Y
	ReservationTermP3Y ReservationTerm = original.ReservationTermP3Y
	ReservationTermP5Y ReservationTerm = original.ReservationTermP5Y
)

type ReservedResourceType = original.ReservedResourceType

const (
	ReservedResourceTypeAppService             ReservedResourceType = original.ReservedResourceTypeAppService
	ReservedResourceTypeAVS                    ReservedResourceType = original.ReservedResourceTypeAVS
	ReservedResourceTypeAzureDataExplorer      ReservedResourceType = original.ReservedResourceTypeAzureDataExplorer
	ReservedResourceTypeAzureFiles             ReservedResourceType = original.ReservedResourceTypeAzureFiles
	ReservedResourceTypeBlockBlob              ReservedResourceType = original.ReservedResourceTypeBlockBlob
	ReservedResourceTypeCosmosDb               ReservedResourceType = original.ReservedResourceTypeCosmosDb
	ReservedResourceTypeDatabricks             ReservedResourceType = original.ReservedResourceTypeDatabricks
	ReservedResourceTypeDataFactory            ReservedResourceType = original.ReservedResourceTypeDataFactory
	ReservedResourceTypeDedicatedHost          ReservedResourceType = original.ReservedResourceTypeDedicatedHost
	ReservedResourceTypeManagedDisk            ReservedResourceType = original.ReservedResourceTypeManagedDisk
	ReservedResourceTypeMariaDb                ReservedResourceType = original.ReservedResourceTypeMariaDb
	ReservedResourceTypeMySQL                  ReservedResourceType = original.ReservedResourceTypeMySQL
	ReservedResourceTypeNetAppStorage          ReservedResourceType = original.ReservedResourceTypeNetAppStorage
	ReservedResourceTypePostgreSQL             ReservedResourceType = original.ReservedResourceTypePostgreSQL
	ReservedResourceTypeRedHat                 ReservedResourceType = original.ReservedResourceTypeRedHat
	ReservedResourceTypeRedHatOsa              ReservedResourceType = original.ReservedResourceTypeRedHatOsa
	ReservedResourceTypeRedisCache             ReservedResourceType = original.ReservedResourceTypeRedisCache
	ReservedResourceTypeSapHana                ReservedResourceType = original.ReservedResourceTypeSapHana
	ReservedResourceTypeSQLAzureHybridBenefit  ReservedResourceType = original.ReservedResourceTypeSQLAzureHybridBenefit
	ReservedResourceTypeSQLDatabases           ReservedResourceType = original.ReservedResourceTypeSQLDatabases
	ReservedResourceTypeSQLDataWarehouse       ReservedResourceType = original.ReservedResourceTypeSQLDataWarehouse
	ReservedResourceTypeSQLEdge                ReservedResourceType = original.ReservedResourceTypeSQLEdge
	ReservedResourceTypeSuseLinux              ReservedResourceType = original.ReservedResourceTypeSuseLinux
	ReservedResourceTypeVirtualMachines        ReservedResourceType = original.ReservedResourceTypeVirtualMachines
	ReservedResourceTypeVirtualMachineSoftware ReservedResourceType = original.ReservedResourceTypeVirtualMachineSoftware
	ReservedResourceTypeVMwareCloudSimple      ReservedResourceType = original.ReservedResourceTypeVMwareCloudSimple
)

type ResourceType = original.ResourceType

const (
	ResourceTypeDedicated       ResourceType = original.ResourceTypeDedicated
	ResourceTypeLowPriority     ResourceType = original.ResourceTypeLowPriority
	ResourceTypeServiceSpecific ResourceType = original.ResourceTypeServiceSpecific
	ResourceTypeShared          ResourceType = original.ResourceTypeShared
	ResourceTypeStandard        ResourceType = original.ResourceTypeStandard
)

type StatusCode = original.StatusCode

const (
	StatusCodeActive                 StatusCode = original.StatusCodeActive
	StatusCodeExpired                StatusCode = original.StatusCodeExpired
	StatusCodeMerged                 StatusCode = original.StatusCodeMerged
	StatusCodeNone                   StatusCode = original.StatusCodeNone
	StatusCodePaymentInstrumentError StatusCode = original.StatusCodePaymentInstrumentError
	StatusCodePending                StatusCode = original.StatusCodePending
	StatusCodeProcessing             StatusCode = original.StatusCodeProcessing
	StatusCodePurchaseError          StatusCode = original.StatusCodePurchaseError
	StatusCodeSplit                  StatusCode = original.StatusCodeSplit
	StatusCodeSucceeded              StatusCode = original.StatusCodeSucceeded
)

type UserFriendlyAppliedScopeType = original.UserFriendlyAppliedScopeType

const (
	UserFriendlyAppliedScopeTypeManagementGroup UserFriendlyAppliedScopeType = original.UserFriendlyAppliedScopeTypeManagementGroup
	UserFriendlyAppliedScopeTypeNone            UserFriendlyAppliedScopeType = original.UserFriendlyAppliedScopeTypeNone
	UserFriendlyAppliedScopeTypeResourceGroup   UserFriendlyAppliedScopeType = original.UserFriendlyAppliedScopeTypeResourceGroup
	UserFriendlyAppliedScopeTypeShared          UserFriendlyAppliedScopeType = original.UserFriendlyAppliedScopeTypeShared
	UserFriendlyAppliedScopeTypeSingle          UserFriendlyAppliedScopeType = original.UserFriendlyAppliedScopeTypeSingle
)

type UserFriendlyRenewState = original.UserFriendlyRenewState

const (
	UserFriendlyRenewStateNotApplicable UserFriendlyRenewState = original.UserFriendlyRenewStateNotApplicable
	UserFriendlyRenewStateNotRenewed    UserFriendlyRenewState = original.UserFriendlyRenewStateNotRenewed
	UserFriendlyRenewStateOff           UserFriendlyRenewState = original.UserFriendlyRenewStateOff
	UserFriendlyRenewStateOn            UserFriendlyRenewState = original.UserFriendlyRenewStateOn
	UserFriendlyRenewStateRenewed       UserFriendlyRenewState = original.UserFriendlyRenewStateRenewed
)

type AppliedReservationList = original.AppliedReservationList
type AppliedReservations = original.AppliedReservations
type AppliedReservationsProperties = original.AppliedReservationsProperties
type AvailableScopeProperties = original.AvailableScopeProperties
type AvailableScopeRequest = original.AvailableScopeRequest
type AvailableScopeRequestProperties = original.AvailableScopeRequestProperties
type BaseClient = original.BaseClient
type BillingInformation = original.BillingInformation
type CalculateExchangeClient = original.CalculateExchangeClient
type CalculateExchangeOperationResultResponse = original.CalculateExchangeOperationResultResponse
type CalculateExchangePostFuture = original.CalculateExchangePostFuture
type CalculateExchangeRequest = original.CalculateExchangeRequest
type CalculateExchangeRequestProperties = original.CalculateExchangeRequestProperties
type CalculateExchangeResponseProperties = original.CalculateExchangeResponseProperties
type CalculatePriceResponse = original.CalculatePriceResponse
type CalculatePriceResponseProperties = original.CalculatePriceResponseProperties
type CalculatePriceResponsePropertiesBillingCurrencyTotal = original.CalculatePriceResponsePropertiesBillingCurrencyTotal
type CalculatePriceResponsePropertiesPricingCurrencyTotal = original.CalculatePriceResponsePropertiesPricingCurrencyTotal
type CalculateRefundClient = original.CalculateRefundClient
type CalculateRefundRequest = original.CalculateRefundRequest
type CalculateRefundRequestProperties = original.CalculateRefundRequestProperties
type CalculateRefundResponse = original.CalculateRefundResponse
type Catalog = original.Catalog
type CatalogMsrp = original.CatalogMsrp
type ChangeDirectoryRequest = original.ChangeDirectoryRequest
type ChangeDirectoryResponse = original.ChangeDirectoryResponse
type ChangeDirectoryResult = original.ChangeDirectoryResult
type Client = original.Client
type CreateGenericQuotaRequestParameters = original.CreateGenericQuotaRequestParameters
type CurrentQuotaLimit = original.CurrentQuotaLimit
type CurrentQuotaLimitBase = original.CurrentQuotaLimitBase
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type ExceptionResponse = original.ExceptionResponse
type ExchangeClient = original.ExchangeClient
type ExchangeOperationResultResponse = original.ExchangeOperationResultResponse
type ExchangePolicyError = original.ExchangePolicyError
type ExchangePolicyErrors = original.ExchangePolicyErrors
type ExchangePostFuture = original.ExchangePostFuture
type ExchangeRequest = original.ExchangeRequest
type ExchangeRequestProperties = original.ExchangeRequestProperties
type ExchangeResponseProperties = original.ExchangeResponseProperties
type ExtendedErrorInfo = original.ExtendedErrorInfo
type ExtendedStatusInfo = original.ExtendedStatusInfo
type List = original.List
type ListCatalog = original.ListCatalog
type ListIterator = original.ListIterator
type ListPage = original.ListPage
type ListResponse = original.ListResponse
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type MergeProperties = original.MergeProperties
type MergePropertiesType = original.MergePropertiesType
type MergeRequest = original.MergeRequest
type OperationClient = original.OperationClient
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationResponse = original.OperationResponse
type OperationResultError = original.OperationResultError
type OrderBillingPlanInformation = original.OrderBillingPlanInformation
type OrderClient = original.OrderClient
type OrderList = original.OrderList
type OrderListIterator = original.OrderListIterator
type OrderListPage = original.OrderListPage
type OrderProperties = original.OrderProperties
type OrderPurchaseFuture = original.OrderPurchaseFuture
type OrderResponse = original.OrderResponse
type Patch = original.Patch
type PatchProperties = original.PatchProperties
type PatchPropertiesRenewProperties = original.PatchPropertiesRenewProperties
type PaymentDetail = original.PaymentDetail
type Price = original.Price
type Properties = original.Properties
type PropertiesUtilization = original.PropertiesUtilization
type PurchaseRequest = original.PurchaseRequest
type PurchaseRequestProperties = original.PurchaseRequestProperties
type PurchaseRequestPropertiesReservedResourceProperties = original.PurchaseRequestPropertiesReservedResourceProperties
type QuotaClient = original.QuotaClient
type QuotaCreateOrUpdateFuture = original.QuotaCreateOrUpdateFuture
type QuotaLimits = original.QuotaLimits
type QuotaLimitsIterator = original.QuotaLimitsIterator
type QuotaLimitsPage = original.QuotaLimitsPage
type QuotaLimitsResponse = original.QuotaLimitsResponse
type QuotaProperties = original.QuotaProperties
type QuotaRequestDetails = original.QuotaRequestDetails
type QuotaRequestDetailsList = original.QuotaRequestDetailsList
type QuotaRequestDetailsListIterator = original.QuotaRequestDetailsListIterator
type QuotaRequestDetailsListPage = original.QuotaRequestDetailsListPage
type QuotaRequestOneResourceProperties = original.QuotaRequestOneResourceProperties
type QuotaRequestOneResourceSubmitResponse = original.QuotaRequestOneResourceSubmitResponse
type QuotaRequestProperties = original.QuotaRequestProperties
type QuotaRequestStatusClient = original.QuotaRequestStatusClient
type QuotaRequestStatusDetails = original.QuotaRequestStatusDetails
type QuotaRequestSubmitResponse = original.QuotaRequestSubmitResponse
type QuotaRequestSubmitResponse201 = original.QuotaRequestSubmitResponse201
type QuotaUpdateFuture = original.QuotaUpdateFuture
type RefundBillingInformation = original.RefundBillingInformation
type RefundPolicyError = original.RefundPolicyError
type RefundPolicyResult = original.RefundPolicyResult
type RefundPolicyResultProperty = original.RefundPolicyResultProperty
type RefundRequest = original.RefundRequest
type RefundRequestProperties = original.RefundRequestProperties
type RefundResponse = original.RefundResponse
type RefundResponseProperties = original.RefundResponseProperties
type RenewPropertiesResponse = original.RenewPropertiesResponse
type RenewPropertiesResponseBillingCurrencyTotal = original.RenewPropertiesResponseBillingCurrencyTotal
type RenewPropertiesResponsePricingCurrencyTotal = original.RenewPropertiesResponsePricingCurrencyTotal
type ReservationAvailableScopesFuture = original.ReservationAvailableScopesFuture
type ReservationMergeFuture = original.ReservationMergeFuture
type ReservationUpdateFuture = original.ReservationUpdateFuture
type ResourceName = original.ResourceName
type Response = original.Response
type ReturnClient = original.ReturnClient
type ScopeProperties = original.ScopeProperties
type ServiceError = original.ServiceError
type ServiceErrorDetail = original.ServiceErrorDetail
type SetObject = original.SetObject
type SkuCapability = original.SkuCapability
type SkuName = original.SkuName
type SkuProperty = original.SkuProperty
type SkuRestriction = original.SkuRestriction
type SplitFuture = original.SplitFuture
type SplitProperties = original.SplitProperties
type SplitPropertiesType = original.SplitPropertiesType
type SplitRequest = original.SplitRequest
type SubRequest = original.SubRequest
type SubscriptionScopeProperties = original.SubscriptionScopeProperties
type Summary = original.Summary
type SystemData = original.SystemData
type ToExchange = original.ToExchange
type ToPurchaseCalculateExchange = original.ToPurchaseCalculateExchange
type ToPurchaseExchange = original.ToPurchaseExchange
type ToReturn = original.ToReturn
type ToReturnForExchange = original.ToReturnForExchange
type UtilizationAggregates = original.UtilizationAggregates

func New() BaseClient {
	return original.New()
}
func NewCalculateExchangeClient() CalculateExchangeClient {
	return original.NewCalculateExchangeClient()
}
func NewCalculateExchangeClientWithBaseURI(baseURI string) CalculateExchangeClient {
	return original.NewCalculateExchangeClientWithBaseURI(baseURI)
}
func NewCalculateRefundClient() CalculateRefundClient {
	return original.NewCalculateRefundClient()
}
func NewCalculateRefundClientWithBaseURI(baseURI string) CalculateRefundClient {
	return original.NewCalculateRefundClientWithBaseURI(baseURI)
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewExchangeClient() ExchangeClient {
	return original.NewExchangeClient()
}
func NewExchangeClientWithBaseURI(baseURI string) ExchangeClient {
	return original.NewExchangeClientWithBaseURI(baseURI)
}
func NewListIterator(page ListPage) ListIterator {
	return original.NewListIterator(page)
}
func NewListPage(cur List, getNextPage func(context.Context, List) (List, error)) ListPage {
	return original.NewListPage(cur, getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOrderClient() OrderClient {
	return original.NewOrderClient()
}
func NewOrderClientWithBaseURI(baseURI string) OrderClient {
	return original.NewOrderClientWithBaseURI(baseURI)
}
func NewOrderListIterator(page OrderListPage) OrderListIterator {
	return original.NewOrderListIterator(page)
}
func NewOrderListPage(cur OrderList, getNextPage func(context.Context, OrderList) (OrderList, error)) OrderListPage {
	return original.NewOrderListPage(cur, getNextPage)
}
func NewQuotaClient() QuotaClient {
	return original.NewQuotaClient()
}
func NewQuotaClientWithBaseURI(baseURI string) QuotaClient {
	return original.NewQuotaClientWithBaseURI(baseURI)
}
func NewQuotaLimitsIterator(page QuotaLimitsPage) QuotaLimitsIterator {
	return original.NewQuotaLimitsIterator(page)
}
func NewQuotaLimitsPage(cur QuotaLimits, getNextPage func(context.Context, QuotaLimits) (QuotaLimits, error)) QuotaLimitsPage {
	return original.NewQuotaLimitsPage(cur, getNextPage)
}
func NewQuotaRequestDetailsListIterator(page QuotaRequestDetailsListPage) QuotaRequestDetailsListIterator {
	return original.NewQuotaRequestDetailsListIterator(page)
}
func NewQuotaRequestDetailsListPage(cur QuotaRequestDetailsList, getNextPage func(context.Context, QuotaRequestDetailsList) (QuotaRequestDetailsList, error)) QuotaRequestDetailsListPage {
	return original.NewQuotaRequestDetailsListPage(cur, getNextPage)
}
func NewQuotaRequestStatusClient() QuotaRequestStatusClient {
	return original.NewQuotaRequestStatusClient()
}
func NewQuotaRequestStatusClientWithBaseURI(baseURI string) QuotaRequestStatusClient {
	return original.NewQuotaRequestStatusClientWithBaseURI(baseURI)
}
func NewReturnClient() ReturnClient {
	return original.NewReturnClient()
}
func NewReturnClientWithBaseURI(baseURI string) ReturnClient {
	return original.NewReturnClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return original.PossibleAppliedScopeTypeValues()
}
func PossibleCalculateExchangeOperationResultStatusValues() []CalculateExchangeOperationResultStatus {
	return original.PossibleCalculateExchangeOperationResultStatusValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDisplayProvisioningStateValues() []DisplayProvisioningState {
	return original.PossibleDisplayProvisioningStateValues()
}
func PossibleErrorResponseCodeValues() []ErrorResponseCode {
	return original.PossibleErrorResponseCodeValues()
}
func PossibleExchangeOperationResultStatusValues() []ExchangeOperationResultStatus {
	return original.PossibleExchangeOperationResultStatusValues()
}
func PossibleInstanceFlexibilityValues() []InstanceFlexibility {
	return original.PossibleInstanceFlexibilityValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLocationValues() []Location {
	return original.PossibleLocationValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func PossiblePaymentStatusValues() []PaymentStatus {
	return original.PossiblePaymentStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleQuotaRequestStateValues() []QuotaRequestState {
	return original.PossibleQuotaRequestStateValues()
}
func PossibleReservationBillingPlanValues() []ReservationBillingPlan {
	return original.PossibleReservationBillingPlanValues()
}
func PossibleReservationTermValues() []ReservationTerm {
	return original.PossibleReservationTermValues()
}
func PossibleReservedResourceTypeValues() []ReservedResourceType {
	return original.PossibleReservedResourceTypeValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleStatusCodeValues() []StatusCode {
	return original.PossibleStatusCodeValues()
}
func PossibleUserFriendlyAppliedScopeTypeValues() []UserFriendlyAppliedScopeType {
	return original.PossibleUserFriendlyAppliedScopeTypeValues()
}
func PossibleUserFriendlyRenewStateValues() []UserFriendlyRenewState {
	return original.PossibleUserFriendlyRenewStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
