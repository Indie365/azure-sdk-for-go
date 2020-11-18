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

package account

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DataLakeStoreAccountState = original.DataLakeStoreAccountState

const (
	Active    DataLakeStoreAccountState = original.Active
	Suspended DataLakeStoreAccountState = original.Suspended
)

type DataLakeStoreAccountStatus = original.DataLakeStoreAccountStatus

const (
	Canceled   DataLakeStoreAccountStatus = original.Canceled
	Creating   DataLakeStoreAccountStatus = original.Creating
	Deleted    DataLakeStoreAccountStatus = original.Deleted
	Deleting   DataLakeStoreAccountStatus = original.Deleting
	Failed     DataLakeStoreAccountStatus = original.Failed
	Patching   DataLakeStoreAccountStatus = original.Patching
	Resuming   DataLakeStoreAccountStatus = original.Resuming
	Running    DataLakeStoreAccountStatus = original.Running
	Succeeded  DataLakeStoreAccountStatus = original.Succeeded
	Suspending DataLakeStoreAccountStatus = original.Suspending
	Undeleting DataLakeStoreAccountStatus = original.Undeleting
)

type EncryptionConfigType = original.EncryptionConfigType

const (
	ServiceManaged EncryptionConfigType = original.ServiceManaged
	UserManaged    EncryptionConfigType = original.UserManaged
)

type EncryptionProvisioningState = original.EncryptionProvisioningState

const (
	EncryptionProvisioningStateCreating  EncryptionProvisioningState = original.EncryptionProvisioningStateCreating
	EncryptionProvisioningStateSucceeded EncryptionProvisioningState = original.EncryptionProvisioningStateSucceeded
)

type EncryptionState = original.EncryptionState

const (
	Disabled EncryptionState = original.Disabled
	Enabled  EncryptionState = original.Enabled
)

type FirewallAllowAzureIpsState = original.FirewallAllowAzureIpsState

const (
	FirewallAllowAzureIpsStateDisabled FirewallAllowAzureIpsState = original.FirewallAllowAzureIpsStateDisabled
	FirewallAllowAzureIpsStateEnabled  FirewallAllowAzureIpsState = original.FirewallAllowAzureIpsStateEnabled
)

type FirewallState = original.FirewallState

const (
	FirewallStateDisabled FirewallState = original.FirewallStateDisabled
	FirewallStateEnabled  FirewallState = original.FirewallStateEnabled
)

type OperationOrigin = original.OperationOrigin

const (
	System     OperationOrigin = original.System
	User       OperationOrigin = original.User
	Usersystem OperationOrigin = original.Usersystem
)

type SubscriptionState = original.SubscriptionState

const (
	SubscriptionStateDeleted      SubscriptionState = original.SubscriptionStateDeleted
	SubscriptionStateRegistered   SubscriptionState = original.SubscriptionStateRegistered
	SubscriptionStateSuspended    SubscriptionState = original.SubscriptionStateSuspended
	SubscriptionStateUnregistered SubscriptionState = original.SubscriptionStateUnregistered
	SubscriptionStateWarned       SubscriptionState = original.SubscriptionStateWarned
)

type TierType = original.TierType

const (
	Commitment100TB TierType = original.Commitment100TB
	Commitment10TB  TierType = original.Commitment10TB
	Commitment1PB   TierType = original.Commitment1PB
	Commitment1TB   TierType = original.Commitment1TB
	Commitment500TB TierType = original.Commitment500TB
	Commitment5PB   TierType = original.Commitment5PB
	Consumption     TierType = original.Consumption
)

type TrustedIDProviderState = original.TrustedIDProviderState

const (
	TrustedIDProviderStateDisabled TrustedIDProviderState = original.TrustedIDProviderStateDisabled
	TrustedIDProviderStateEnabled  TrustedIDProviderState = original.TrustedIDProviderStateEnabled
)

type UsageUnit = original.UsageUnit

const (
	Bytes           UsageUnit = original.Bytes
	BytesPerSecond  UsageUnit = original.BytesPerSecond
	Count           UsageUnit = original.Count
	CountsPerSecond UsageUnit = original.CountsPerSecond
	Percent         UsageUnit = original.Percent
	Seconds         UsageUnit = original.Seconds
)

type AccountsClient = original.AccountsClient
type AccountsCreateFutureType = original.AccountsCreateFutureType
type AccountsDeleteFutureType = original.AccountsDeleteFutureType
type AccountsUpdateFutureType = original.AccountsUpdateFutureType
type BaseClient = original.BaseClient
type CapabilityInformation = original.CapabilityInformation
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CreateDataLakeStoreAccountParameters = original.CreateDataLakeStoreAccountParameters
type CreateDataLakeStoreAccountProperties = original.CreateDataLakeStoreAccountProperties
type CreateFirewallRuleWithAccountParameters = original.CreateFirewallRuleWithAccountParameters
type CreateOrUpdateFirewallRuleParameters = original.CreateOrUpdateFirewallRuleParameters
type CreateOrUpdateFirewallRuleProperties = original.CreateOrUpdateFirewallRuleProperties
type CreateOrUpdateTrustedIDProviderParameters = original.CreateOrUpdateTrustedIDProviderParameters
type CreateOrUpdateTrustedIDProviderProperties = original.CreateOrUpdateTrustedIDProviderProperties
type CreateOrUpdateVirtualNetworkRuleParameters = original.CreateOrUpdateVirtualNetworkRuleParameters
type CreateOrUpdateVirtualNetworkRuleProperties = original.CreateOrUpdateVirtualNetworkRuleProperties
type CreateTrustedIDProviderWithAccountParameters = original.CreateTrustedIDProviderWithAccountParameters
type CreateVirtualNetworkRuleWithAccountParameters = original.CreateVirtualNetworkRuleWithAccountParameters
type DataLakeStoreAccount = original.DataLakeStoreAccount
type DataLakeStoreAccountBasic = original.DataLakeStoreAccountBasic
type DataLakeStoreAccountListResult = original.DataLakeStoreAccountListResult
type DataLakeStoreAccountListResultIterator = original.DataLakeStoreAccountListResultIterator
type DataLakeStoreAccountListResultPage = original.DataLakeStoreAccountListResultPage
type DataLakeStoreAccountProperties = original.DataLakeStoreAccountProperties
type DataLakeStoreAccountPropertiesBasic = original.DataLakeStoreAccountPropertiesBasic
type EncryptionConfig = original.EncryptionConfig
type EncryptionIdentity = original.EncryptionIdentity
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleListResultIterator = original.FirewallRuleListResultIterator
type FirewallRuleListResultPage = original.FirewallRuleListResultPage
type FirewallRuleProperties = original.FirewallRuleProperties
type FirewallRulesClient = original.FirewallRulesClient
type KeyVaultMetaInfo = original.KeyVaultMetaInfo
type LocationsClient = original.LocationsClient
type NameAvailabilityInformation = original.NameAvailabilityInformation
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type SubResource = original.SubResource
type TrustedIDProvider = original.TrustedIDProvider
type TrustedIDProviderListResult = original.TrustedIDProviderListResult
type TrustedIDProviderListResultIterator = original.TrustedIDProviderListResultIterator
type TrustedIDProviderListResultPage = original.TrustedIDProviderListResultPage
type TrustedIDProviderProperties = original.TrustedIDProviderProperties
type TrustedIDProvidersClient = original.TrustedIDProvidersClient
type UpdateDataLakeStoreAccountParameters = original.UpdateDataLakeStoreAccountParameters
type UpdateDataLakeStoreAccountProperties = original.UpdateDataLakeStoreAccountProperties
type UpdateEncryptionConfig = original.UpdateEncryptionConfig
type UpdateFirewallRuleParameters = original.UpdateFirewallRuleParameters
type UpdateFirewallRuleProperties = original.UpdateFirewallRuleProperties
type UpdateFirewallRuleWithAccountParameters = original.UpdateFirewallRuleWithAccountParameters
type UpdateKeyVaultMetaInfo = original.UpdateKeyVaultMetaInfo
type UpdateTrustedIDProviderParameters = original.UpdateTrustedIDProviderParameters
type UpdateTrustedIDProviderProperties = original.UpdateTrustedIDProviderProperties
type UpdateTrustedIDProviderWithAccountParameters = original.UpdateTrustedIDProviderWithAccountParameters
type UpdateVirtualNetworkRuleParameters = original.UpdateVirtualNetworkRuleParameters
type UpdateVirtualNetworkRuleProperties = original.UpdateVirtualNetworkRuleProperties
type UpdateVirtualNetworkRuleWithAccountParameters = original.UpdateVirtualNetworkRuleWithAccountParameters
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type VirtualNetworkRule = original.VirtualNetworkRule
type VirtualNetworkRuleListResult = original.VirtualNetworkRuleListResult
type VirtualNetworkRuleListResultIterator = original.VirtualNetworkRuleListResultIterator
type VirtualNetworkRuleListResultPage = original.VirtualNetworkRuleListResultPage
type VirtualNetworkRuleProperties = original.VirtualNetworkRuleProperties
type VirtualNetworkRulesClient = original.VirtualNetworkRulesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataLakeStoreAccountListResultIterator(page DataLakeStoreAccountListResultPage) DataLakeStoreAccountListResultIterator {
	return original.NewDataLakeStoreAccountListResultIterator(page)
}
func NewDataLakeStoreAccountListResultPage(cur DataLakeStoreAccountListResult, getNextPage func(context.Context, DataLakeStoreAccountListResult) (DataLakeStoreAccountListResult, error)) DataLakeStoreAccountListResultPage {
	return original.NewDataLakeStoreAccountListResultPage(cur, getNextPage)
}
func NewFirewallRuleListResultIterator(page FirewallRuleListResultPage) FirewallRuleListResultIterator {
	return original.NewFirewallRuleListResultIterator(page)
}
func NewFirewallRuleListResultPage(cur FirewallRuleListResult, getNextPage func(context.Context, FirewallRuleListResult) (FirewallRuleListResult, error)) FirewallRuleListResultPage {
	return original.NewFirewallRuleListResultPage(cur, getNextPage)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTrustedIDProviderListResultIterator(page TrustedIDProviderListResultPage) TrustedIDProviderListResultIterator {
	return original.NewTrustedIDProviderListResultIterator(page)
}
func NewTrustedIDProviderListResultPage(cur TrustedIDProviderListResult, getNextPage func(context.Context, TrustedIDProviderListResult) (TrustedIDProviderListResult, error)) TrustedIDProviderListResultPage {
	return original.NewTrustedIDProviderListResultPage(cur, getNextPage)
}
func NewTrustedIDProvidersClient(subscriptionID string) TrustedIDProvidersClient {
	return original.NewTrustedIDProvidersClient(subscriptionID)
}
func NewTrustedIDProvidersClientWithBaseURI(baseURI string, subscriptionID string) TrustedIDProvidersClient {
	return original.NewTrustedIDProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkRuleListResultIterator(page VirtualNetworkRuleListResultPage) VirtualNetworkRuleListResultIterator {
	return original.NewVirtualNetworkRuleListResultIterator(page)
}
func NewVirtualNetworkRulesClient(subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClient(subscriptionID)
}
func NewVirtualNetworkRulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkRulesClient {
	return original.NewVirtualNetworkRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDataLakeStoreAccountStateValues() []DataLakeStoreAccountState {
	return original.PossibleDataLakeStoreAccountStateValues()
}
func PossibleDataLakeStoreAccountStatusValues() []DataLakeStoreAccountStatus {
	return original.PossibleDataLakeStoreAccountStatusValues()
}
func PossibleEncryptionConfigTypeValues() []EncryptionConfigType {
	return original.PossibleEncryptionConfigTypeValues()
}
func PossibleEncryptionProvisioningStateValues() []EncryptionProvisioningState {
	return original.PossibleEncryptionProvisioningStateValues()
}
func PossibleEncryptionStateValues() []EncryptionState {
	return original.PossibleEncryptionStateValues()
}
func PossibleFirewallAllowAzureIpsStateValues() []FirewallAllowAzureIpsState {
	return original.PossibleFirewallAllowAzureIpsStateValues()
}
func PossibleFirewallStateValues() []FirewallState {
	return original.PossibleFirewallStateValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleSubscriptionStateValues() []SubscriptionState {
	return original.PossibleSubscriptionStateValues()
}
func PossibleTierTypeValues() []TierType {
	return original.PossibleTierTypeValues()
}
func PossibleTrustedIDProviderStateValues() []TrustedIDProviderState {
	return original.PossibleTrustedIDProviderStateValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
