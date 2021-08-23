//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package datadog

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/datadog/mgmt/2021-03-01/datadog"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type LiftrResourceCategories = original.LiftrResourceCategories

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = original.LiftrResourceCategoriesMonitorLogs
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = original.LiftrResourceCategoriesUnknown
)

type ManagedIdentityTypes = original.ManagedIdentityTypes

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = original.ManagedIdentityTypesSystemAssigned
	ManagedIdentityTypesUserAssigned   ManagedIdentityTypes = original.ManagedIdentityTypesUserAssigned
)

type MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatus

const (
	MarketplaceSubscriptionStatusActive       MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatusActive
	MarketplaceSubscriptionStatusProvisioning MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatusProvisioning
	MarketplaceSubscriptionStatusSuspended    MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatusSuspended
	MarketplaceSubscriptionStatusUnsubscribed MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatusUnsubscribed
)

type MonitoringStatus = original.MonitoringStatus

const (
	MonitoringStatusDisabled MonitoringStatus = original.MonitoringStatusDisabled
	MonitoringStatusEnabled  MonitoringStatus = original.MonitoringStatusEnabled
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateNotSpecified ProvisioningState = original.ProvisioningStateNotSpecified
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type SingleSignOnStates = original.SingleSignOnStates

const (
	SingleSignOnStatesDisable  SingleSignOnStates = original.SingleSignOnStatesDisable
	SingleSignOnStatesEnable   SingleSignOnStates = original.SingleSignOnStatesEnable
	SingleSignOnStatesExisting SingleSignOnStates = original.SingleSignOnStatesExisting
	SingleSignOnStatesInitial  SingleSignOnStates = original.SingleSignOnStatesInitial
)

type TagAction = original.TagAction

const (
	TagActionExclude TagAction = original.TagActionExclude
	TagActionInclude TagAction = original.TagActionInclude
)

type APIKey = original.APIKey
type APIKeyListResponse = original.APIKeyListResponse
type APIKeyListResponseIterator = original.APIKeyListResponseIterator
type APIKeyListResponsePage = original.APIKeyListResponsePage
type AgreementProperties = original.AgreementProperties
type AgreementResource = original.AgreementResource
type AgreementResourceListResponse = original.AgreementResourceListResponse
type AgreementResourceListResponseIterator = original.AgreementResourceListResponseIterator
type AgreementResourceListResponsePage = original.AgreementResourceListResponsePage
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type FilteringTag = original.FilteringTag
type Host = original.Host
type HostListResponse = original.HostListResponse
type HostListResponseIterator = original.HostListResponseIterator
type HostListResponsePage = original.HostListResponsePage
type HostMetadata = original.HostMetadata
type IdentityProperties = original.IdentityProperties
type InstallMethod = original.InstallMethod
type LinkedResource = original.LinkedResource
type LinkedResourceListResponse = original.LinkedResourceListResponse
type LinkedResourceListResponseIterator = original.LinkedResourceListResponseIterator
type LinkedResourceListResponsePage = original.LinkedResourceListResponsePage
type LogRules = original.LogRules
type LogsAgent = original.LogsAgent
type MarketplaceAgreementsClient = original.MarketplaceAgreementsClient
type MetricRules = original.MetricRules
type MonitorProperties = original.MonitorProperties
type MonitorResource = original.MonitorResource
type MonitorResourceListResponse = original.MonitorResourceListResponse
type MonitorResourceListResponseIterator = original.MonitorResourceListResponseIterator
type MonitorResourceListResponsePage = original.MonitorResourceListResponsePage
type MonitorResourceUpdateParameters = original.MonitorResourceUpdateParameters
type MonitorUpdateProperties = original.MonitorUpdateProperties
type MonitoredResource = original.MonitoredResource
type MonitoredResourceListResponse = original.MonitoredResourceListResponse
type MonitoredResourceListResponseIterator = original.MonitoredResourceListResponseIterator
type MonitoredResourceListResponsePage = original.MonitoredResourceListResponsePage
type MonitoringTagRules = original.MonitoringTagRules
type MonitoringTagRulesListResponse = original.MonitoringTagRulesListResponse
type MonitoringTagRulesListResponseIterator = original.MonitoringTagRulesListResponseIterator
type MonitoringTagRulesListResponsePage = original.MonitoringTagRulesListResponsePage
type MonitoringTagRulesProperties = original.MonitoringTagRulesProperties
type MonitorsClient = original.MonitorsClient
type MonitorsCreateFuture = original.MonitorsCreateFuture
type MonitorsDeleteFuture = original.MonitorsDeleteFuture
type MonitorsUpdateFuture = original.MonitorsUpdateFuture
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type OrganizationProperties = original.OrganizationProperties
type ResourceSku = original.ResourceSku
type SetPasswordLink = original.SetPasswordLink
type SingleSignOnConfigurationsClient = original.SingleSignOnConfigurationsClient
type SingleSignOnConfigurationsCreateOrUpdateFuture = original.SingleSignOnConfigurationsCreateOrUpdateFuture
type SingleSignOnProperties = original.SingleSignOnProperties
type SingleSignOnResource = original.SingleSignOnResource
type SingleSignOnResourceListResponse = original.SingleSignOnResourceListResponse
type SingleSignOnResourceListResponseIterator = original.SingleSignOnResourceListResponseIterator
type SingleSignOnResourceListResponsePage = original.SingleSignOnResourceListResponsePage
type SystemData = original.SystemData
type TagRulesClient = original.TagRulesClient
type UserInfo = original.UserInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAPIKeyListResponseIterator(page APIKeyListResponsePage) APIKeyListResponseIterator {
	return original.NewAPIKeyListResponseIterator(page)
}
func NewAPIKeyListResponsePage(cur APIKeyListResponse, getNextPage func(context.Context, APIKeyListResponse) (APIKeyListResponse, error)) APIKeyListResponsePage {
	return original.NewAPIKeyListResponsePage(cur, getNextPage)
}
func NewAgreementResourceListResponseIterator(page AgreementResourceListResponsePage) AgreementResourceListResponseIterator {
	return original.NewAgreementResourceListResponseIterator(page)
}
func NewAgreementResourceListResponsePage(cur AgreementResourceListResponse, getNextPage func(context.Context, AgreementResourceListResponse) (AgreementResourceListResponse, error)) AgreementResourceListResponsePage {
	return original.NewAgreementResourceListResponsePage(cur, getNextPage)
}
func NewHostListResponseIterator(page HostListResponsePage) HostListResponseIterator {
	return original.NewHostListResponseIterator(page)
}
func NewHostListResponsePage(cur HostListResponse, getNextPage func(context.Context, HostListResponse) (HostListResponse, error)) HostListResponsePage {
	return original.NewHostListResponsePage(cur, getNextPage)
}
func NewLinkedResourceListResponseIterator(page LinkedResourceListResponsePage) LinkedResourceListResponseIterator {
	return original.NewLinkedResourceListResponseIterator(page)
}
func NewLinkedResourceListResponsePage(cur LinkedResourceListResponse, getNextPage func(context.Context, LinkedResourceListResponse) (LinkedResourceListResponse, error)) LinkedResourceListResponsePage {
	return original.NewLinkedResourceListResponsePage(cur, getNextPage)
}
func NewMarketplaceAgreementsClient(subscriptionID string) MarketplaceAgreementsClient {
	return original.NewMarketplaceAgreementsClient(subscriptionID)
}
func NewMarketplaceAgreementsClientWithBaseURI(baseURI string, subscriptionID string) MarketplaceAgreementsClient {
	return original.NewMarketplaceAgreementsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMonitorResourceListResponseIterator(page MonitorResourceListResponsePage) MonitorResourceListResponseIterator {
	return original.NewMonitorResourceListResponseIterator(page)
}
func NewMonitorResourceListResponsePage(cur MonitorResourceListResponse, getNextPage func(context.Context, MonitorResourceListResponse) (MonitorResourceListResponse, error)) MonitorResourceListResponsePage {
	return original.NewMonitorResourceListResponsePage(cur, getNextPage)
}
func NewMonitoredResourceListResponseIterator(page MonitoredResourceListResponsePage) MonitoredResourceListResponseIterator {
	return original.NewMonitoredResourceListResponseIterator(page)
}
func NewMonitoredResourceListResponsePage(cur MonitoredResourceListResponse, getNextPage func(context.Context, MonitoredResourceListResponse) (MonitoredResourceListResponse, error)) MonitoredResourceListResponsePage {
	return original.NewMonitoredResourceListResponsePage(cur, getNextPage)
}
func NewMonitoringTagRulesListResponseIterator(page MonitoringTagRulesListResponsePage) MonitoringTagRulesListResponseIterator {
	return original.NewMonitoringTagRulesListResponseIterator(page)
}
func NewMonitoringTagRulesListResponsePage(cur MonitoringTagRulesListResponse, getNextPage func(context.Context, MonitoringTagRulesListResponse) (MonitoringTagRulesListResponse, error)) MonitoringTagRulesListResponsePage {
	return original.NewMonitoringTagRulesListResponsePage(cur, getNextPage)
}
func NewMonitorsClient(subscriptionID string) MonitorsClient {
	return original.NewMonitorsClient(subscriptionID)
}
func NewMonitorsClientWithBaseURI(baseURI string, subscriptionID string) MonitorsClient {
	return original.NewMonitorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSingleSignOnConfigurationsClient(subscriptionID string) SingleSignOnConfigurationsClient {
	return original.NewSingleSignOnConfigurationsClient(subscriptionID)
}
func NewSingleSignOnConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SingleSignOnConfigurationsClient {
	return original.NewSingleSignOnConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSingleSignOnResourceListResponseIterator(page SingleSignOnResourceListResponsePage) SingleSignOnResourceListResponseIterator {
	return original.NewSingleSignOnResourceListResponseIterator(page)
}
func NewSingleSignOnResourceListResponsePage(cur SingleSignOnResourceListResponse, getNextPage func(context.Context, SingleSignOnResourceListResponse) (SingleSignOnResourceListResponse, error)) SingleSignOnResourceListResponsePage {
	return original.NewSingleSignOnResourceListResponsePage(cur, getNextPage)
}
func NewTagRulesClient(subscriptionID string) TagRulesClient {
	return original.NewTagRulesClient(subscriptionID)
}
func NewTagRulesClientWithBaseURI(baseURI string, subscriptionID string) TagRulesClient {
	return original.NewTagRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return original.PossibleLiftrResourceCategoriesValues()
}
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return original.PossibleManagedIdentityTypesValues()
}
func PossibleMarketplaceSubscriptionStatusValues() []MarketplaceSubscriptionStatus {
	return original.PossibleMarketplaceSubscriptionStatusValues()
}
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return original.PossibleMonitoringStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSingleSignOnStatesValues() []SingleSignOnStates {
	return original.PossibleSingleSignOnStatesValues()
}
func PossibleTagActionValues() []TagAction {
	return original.PossibleTagActionValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
