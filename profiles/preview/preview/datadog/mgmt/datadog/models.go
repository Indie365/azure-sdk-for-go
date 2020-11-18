// +build go1.9

// Copyright 2020 Microsoft Corporation
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

package datadog

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/datadog/mgmt/2020-02-01-preview/datadog"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type LiftrResourceCategories = original.LiftrResourceCategories

const (
	MonitorLogs LiftrResourceCategories = original.MonitorLogs
	Unknown     LiftrResourceCategories = original.Unknown
)

type ManagedIdentityTypes = original.ManagedIdentityTypes

const (
	SystemAssigned ManagedIdentityTypes = original.SystemAssigned
	UserAssigned   ManagedIdentityTypes = original.UserAssigned
)

type MarketplaceSubscriptionStatus = original.MarketplaceSubscriptionStatus

const (
	Active    MarketplaceSubscriptionStatus = original.Active
	Suspended MarketplaceSubscriptionStatus = original.Suspended
)

type MonitoringStatus = original.MonitoringStatus

const (
	Disabled MonitoringStatus = original.Disabled
	Enabled  MonitoringStatus = original.Enabled
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted     ProvisioningState = original.Accepted
	Canceled     ProvisioningState = original.Canceled
	Creating     ProvisioningState = original.Creating
	Deleted      ProvisioningState = original.Deleted
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	NotSpecified ProvisioningState = original.NotSpecified
	Succeeded    ProvisioningState = original.Succeeded
	Updating     ProvisioningState = original.Updating
)

type SingleSignOnStates = original.SingleSignOnStates

const (
	Disable  SingleSignOnStates = original.Disable
	Enable   SingleSignOnStates = original.Enable
	Existing SingleSignOnStates = original.Existing
	Initial  SingleSignOnStates = original.Initial
)

type TagAction = original.TagAction

const (
	Exclude TagAction = original.Exclude
	Include TagAction = original.Include
)

type APIKey = original.APIKey
type APIKeyListResponse = original.APIKeyListResponse
type APIKeyListResponseIterator = original.APIKeyListResponseIterator
type APIKeyListResponsePage = original.APIKeyListResponsePage
type APIKeysClient = original.APIKeysClient
type BaseClient = original.BaseClient
type ErrorResponseBody = original.ErrorResponseBody
type FilteringTag = original.FilteringTag
type Host = original.Host
type HostListResponse = original.HostListResponse
type HostListResponseIterator = original.HostListResponseIterator
type HostListResponsePage = original.HostListResponsePage
type HostMetadata = original.HostMetadata
type HostsClient = original.HostsClient
type IdentityProperties = original.IdentityProperties
type InstallMethod = original.InstallMethod
type LinkedResource = original.LinkedResource
type LinkedResourceListResponse = original.LinkedResourceListResponse
type LinkedResourceListResponseIterator = original.LinkedResourceListResponseIterator
type LinkedResourceListResponsePage = original.LinkedResourceListResponsePage
type LinkedResourcesClient = original.LinkedResourcesClient
type LogRules = original.LogRules
type LogsAgent = original.LogsAgent
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
type MonitoredResourcesClient = original.MonitoredResourcesClient
type MonitoringTagRules = original.MonitoringTagRules
type MonitoringTagRulesListResponse = original.MonitoringTagRulesListResponse
type MonitoringTagRulesListResponseIterator = original.MonitoringTagRulesListResponseIterator
type MonitoringTagRulesListResponsePage = original.MonitoringTagRulesListResponsePage
type MonitoringTagRulesProperties = original.MonitoringTagRulesProperties
type MonitorsClient = original.MonitorsClient
type MonitorsCreateFuture = original.MonitorsCreateFuture
type MonitorsDeleteFuture = original.MonitorsDeleteFuture
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type OrganizationProperties = original.OrganizationProperties
type RefreshSetPasswordClient = original.RefreshSetPasswordClient
type ResourceProviderDefaultErrorResponse = original.ResourceProviderDefaultErrorResponse
type ResourceSku = original.ResourceSku
type SetPasswordLink = original.SetPasswordLink
type SingleSignOnConfigurationsClient = original.SingleSignOnConfigurationsClient
type SingleSignOnConfigurationsCreateOrUpdateFuture = original.SingleSignOnConfigurationsCreateOrUpdateFuture
type SingleSignOnProperties = original.SingleSignOnProperties
type SingleSignOnResource = original.SingleSignOnResource
type SingleSignOnResourceListResponse = original.SingleSignOnResourceListResponse
type SingleSignOnResourceListResponseIterator = original.SingleSignOnResourceListResponseIterator
type SingleSignOnResourceListResponsePage = original.SingleSignOnResourceListResponsePage
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
func NewAPIKeysClient(subscriptionID string) APIKeysClient {
	return original.NewAPIKeysClient(subscriptionID)
}
func NewAPIKeysClientWithBaseURI(baseURI string, subscriptionID string) APIKeysClient {
	return original.NewAPIKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewHostListResponseIterator(page HostListResponsePage) HostListResponseIterator {
	return original.NewHostListResponseIterator(page)
}
func NewHostListResponsePage(cur HostListResponse, getNextPage func(context.Context, HostListResponse) (HostListResponse, error)) HostListResponsePage {
	return original.NewHostListResponsePage(cur, getNextPage)
}
func NewHostsClient(subscriptionID string) HostsClient {
	return original.NewHostsClient(subscriptionID)
}
func NewHostsClientWithBaseURI(baseURI string, subscriptionID string) HostsClient {
	return original.NewHostsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedResourceListResponseIterator(page LinkedResourceListResponsePage) LinkedResourceListResponseIterator {
	return original.NewLinkedResourceListResponseIterator(page)
}
func NewLinkedResourceListResponsePage(cur LinkedResourceListResponse, getNextPage func(context.Context, LinkedResourceListResponse) (LinkedResourceListResponse, error)) LinkedResourceListResponsePage {
	return original.NewLinkedResourceListResponsePage(cur, getNextPage)
}
func NewLinkedResourcesClient(subscriptionID string) LinkedResourcesClient {
	return original.NewLinkedResourcesClient(subscriptionID)
}
func NewLinkedResourcesClientWithBaseURI(baseURI string, subscriptionID string) LinkedResourcesClient {
	return original.NewLinkedResourcesClientWithBaseURI(baseURI, subscriptionID)
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
func NewMonitoredResourcesClient(subscriptionID string) MonitoredResourcesClient {
	return original.NewMonitoredResourcesClient(subscriptionID)
}
func NewMonitoredResourcesClientWithBaseURI(baseURI string, subscriptionID string) MonitoredResourcesClient {
	return original.NewMonitoredResourcesClientWithBaseURI(baseURI, subscriptionID)
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
func NewRefreshSetPasswordClient(subscriptionID string) RefreshSetPasswordClient {
	return original.NewRefreshSetPasswordClient(subscriptionID)
}
func NewRefreshSetPasswordClientWithBaseURI(baseURI string, subscriptionID string) RefreshSetPasswordClient {
	return original.NewRefreshSetPasswordClientWithBaseURI(baseURI, subscriptionID)
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
func NewTagRulesClient(subscriptionID string) TagRulesClient {
	return original.NewTagRulesClient(subscriptionID)
}
func NewTagRulesClientWithBaseURI(baseURI string, subscriptionID string) TagRulesClient {
	return original.NewTagRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
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
