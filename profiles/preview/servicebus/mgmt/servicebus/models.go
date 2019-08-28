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

package servicebus

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type DefaultAction = original.DefaultAction

const (
	Allow DefaultAction = original.Allow
	Deny  DefaultAction = original.Deny
)

type EncodingCaptureDescription = original.EncodingCaptureDescription

const (
	Avro        EncodingCaptureDescription = original.Avro
	AvroDeflate EncodingCaptureDescription = original.AvroDeflate
)

type EntityStatus = original.EntityStatus

const (
	Active          EntityStatus = original.Active
	Creating        EntityStatus = original.Creating
	Deleting        EntityStatus = original.Deleting
	Disabled        EntityStatus = original.Disabled
	ReceiveDisabled EntityStatus = original.ReceiveDisabled
	Renaming        EntityStatus = original.Renaming
	Restoring       EntityStatus = original.Restoring
	SendDisabled    EntityStatus = original.SendDisabled
	Unknown         EntityStatus = original.Unknown
)

type FilterType = original.FilterType

const (
	FilterTypeCorrelationFilter FilterType = original.FilterTypeCorrelationFilter
	FilterTypeSQLFilter         FilterType = original.FilterTypeSQLFilter
)

type KeyType = original.KeyType

const (
	PrimaryKey   KeyType = original.PrimaryKey
	SecondaryKey KeyType = original.SecondaryKey
)

type NameSpaceType = original.NameSpaceType

const (
	EventHub        NameSpaceType = original.EventHub
	Messaging       NameSpaceType = original.Messaging
	Mixed           NameSpaceType = original.Mixed
	NotificationHub NameSpaceType = original.NotificationHub
	Relay           NameSpaceType = original.Relay
)

type NetworkRuleIPAction = original.NetworkRuleIPAction

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = original.NetworkRuleIPActionAllow
)

type ProvisioningStateDR = original.ProvisioningStateDR

const (
	Accepted  ProvisioningStateDR = original.Accepted
	Failed    ProvisioningStateDR = original.Failed
	Succeeded ProvisioningStateDR = original.Succeeded
)

type RoleDisasterRecovery = original.RoleDisasterRecovery

const (
	Primary               RoleDisasterRecovery = original.Primary
	PrimaryNotReplicating RoleDisasterRecovery = original.PrimaryNotReplicating
	Secondary             RoleDisasterRecovery = original.Secondary
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type UnavailableReason = original.UnavailableReason

const (
	InvalidName                           UnavailableReason = original.InvalidName
	NameInLockdown                        UnavailableReason = original.NameInLockdown
	NameInUse                             UnavailableReason = original.NameInUse
	None                                  UnavailableReason = original.None
	SubscriptionIsDisabled                UnavailableReason = original.SubscriptionIsDisabled
	TooManyNamespaceInCurrentSubscription UnavailableReason = original.TooManyNamespaceInCurrentSubscription
)

type AccessKeys = original.AccessKeys
type Action = original.Action
type ArmDisasterRecovery = original.ArmDisasterRecovery
type ArmDisasterRecoveryListResult = original.ArmDisasterRecoveryListResult
type ArmDisasterRecoveryListResultIterator = original.ArmDisasterRecoveryListResultIterator
type ArmDisasterRecoveryListResultPage = original.ArmDisasterRecoveryListResultPage
type ArmDisasterRecoveryProperties = original.ArmDisasterRecoveryProperties
type AuthorizationRuleProperties = original.AuthorizationRuleProperties
type BaseClient = original.BaseClient
type CaptureDescription = original.CaptureDescription
type CheckNameAvailability = original.CheckNameAvailability
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CorrelationFilter = original.CorrelationFilter
type Destination = original.Destination
type DestinationProperties = original.DestinationProperties
type DisasterRecoveryConfigsClient = original.DisasterRecoveryConfigsClient
type ErrorResponse = original.ErrorResponse
type EventHubListResult = original.EventHubListResult
type EventHubListResultIterator = original.EventHubListResultIterator
type EventHubListResultPage = original.EventHubListResultPage
type EventHubsClient = original.EventHubsClient
type Eventhub = original.Eventhub
type EventhubProperties = original.EventhubProperties
type MessageCountDetails = original.MessageCountDetails
type MigrationConfigListResult = original.MigrationConfigListResult
type MigrationConfigListResultIterator = original.MigrationConfigListResultIterator
type MigrationConfigListResultPage = original.MigrationConfigListResultPage
type MigrationConfigProperties = original.MigrationConfigProperties
type MigrationConfigPropertiesProperties = original.MigrationConfigPropertiesProperties
type MigrationConfigsClient = original.MigrationConfigsClient
type MigrationConfigsCreateAndStartMigrationFuture = original.MigrationConfigsCreateAndStartMigrationFuture
type NWRuleSetIPRules = original.NWRuleSetIPRules
type NWRuleSetVirtualNetworkRules = original.NWRuleSetVirtualNetworkRules
type NamespacesClient = original.NamespacesClient
type NamespacesCreateOrUpdateFuture = original.NamespacesCreateOrUpdateFuture
type NamespacesDeleteFuture = original.NamespacesDeleteFuture
type NetworkRuleSet = original.NetworkRuleSet
type NetworkRuleSetListResult = original.NetworkRuleSetListResult
type NetworkRuleSetProperties = original.NetworkRuleSetProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PremiumMessagingRegions = original.PremiumMessagingRegions
type PremiumMessagingRegionsClient = original.PremiumMessagingRegionsClient
type PremiumMessagingRegionsListResult = original.PremiumMessagingRegionsListResult
type PremiumMessagingRegionsListResultIterator = original.PremiumMessagingRegionsListResultIterator
type PremiumMessagingRegionsListResultPage = original.PremiumMessagingRegionsListResultPage
type PremiumMessagingRegionsProperties = original.PremiumMessagingRegionsProperties
type QueuesClient = original.QueuesClient
type RegenerateAccessKeyParameters = original.RegenerateAccessKeyParameters
type RegionsClient = original.RegionsClient
type Resource = original.Resource
type ResourceNamespacePatch = original.ResourceNamespacePatch
type Rule = original.Rule
type RuleListResult = original.RuleListResult
type RuleListResultIterator = original.RuleListResultIterator
type RuleListResultPage = original.RuleListResultPage
type Ruleproperties = original.Ruleproperties
type RulesClient = original.RulesClient
type SBAuthorizationRule = original.SBAuthorizationRule
type SBAuthorizationRuleListResult = original.SBAuthorizationRuleListResult
type SBAuthorizationRuleListResultIterator = original.SBAuthorizationRuleListResultIterator
type SBAuthorizationRuleListResultPage = original.SBAuthorizationRuleListResultPage
type SBAuthorizationRuleProperties = original.SBAuthorizationRuleProperties
type SBNamespace = original.SBNamespace
type SBNamespaceListResult = original.SBNamespaceListResult
type SBNamespaceListResultIterator = original.SBNamespaceListResultIterator
type SBNamespaceListResultPage = original.SBNamespaceListResultPage
type SBNamespaceMigrate = original.SBNamespaceMigrate
type SBNamespaceProperties = original.SBNamespaceProperties
type SBNamespaceUpdateParameters = original.SBNamespaceUpdateParameters
type SBQueue = original.SBQueue
type SBQueueListResult = original.SBQueueListResult
type SBQueueListResultIterator = original.SBQueueListResultIterator
type SBQueueListResultPage = original.SBQueueListResultPage
type SBQueueProperties = original.SBQueueProperties
type SBSku = original.SBSku
type SBSubscription = original.SBSubscription
type SBSubscriptionListResult = original.SBSubscriptionListResult
type SBSubscriptionListResultIterator = original.SBSubscriptionListResultIterator
type SBSubscriptionListResultPage = original.SBSubscriptionListResultPage
type SBSubscriptionProperties = original.SBSubscriptionProperties
type SBTopic = original.SBTopic
type SBTopicListResult = original.SBTopicListResult
type SBTopicListResultIterator = original.SBTopicListResultIterator
type SBTopicListResultPage = original.SBTopicListResultPage
type SBTopicProperties = original.SBTopicProperties
type SQLFilter = original.SQLFilter
type SQLRuleAction = original.SQLRuleAction
type Subnet = original.Subnet
type SubscriptionsClient = original.SubscriptionsClient
type TopicsClient = original.TopicsClient
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewArmDisasterRecoveryListResultIterator(page ArmDisasterRecoveryListResultPage) ArmDisasterRecoveryListResultIterator {
	return original.NewArmDisasterRecoveryListResultIterator(page)
}
func NewArmDisasterRecoveryListResultPage(getNextPage func(context.Context, ArmDisasterRecoveryListResult) (ArmDisasterRecoveryListResult, error)) ArmDisasterRecoveryListResultPage {
	return original.NewArmDisasterRecoveryListResultPage(getNextPage)
}
func NewDisasterRecoveryConfigsClient(subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClient(subscriptionID)
}
func NewDisasterRecoveryConfigsClientWithBaseURI(baseURI string, subscriptionID string) DisasterRecoveryConfigsClient {
	return original.NewDisasterRecoveryConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventHubListResultIterator(page EventHubListResultPage) EventHubListResultIterator {
	return original.NewEventHubListResultIterator(page)
}
func NewEventHubListResultPage(getNextPage func(context.Context, EventHubListResult) (EventHubListResult, error)) EventHubListResultPage {
	return original.NewEventHubListResultPage(getNextPage)
}
func NewEventHubsClient(subscriptionID string) EventHubsClient {
	return original.NewEventHubsClient(subscriptionID)
}
func NewEventHubsClientWithBaseURI(baseURI string, subscriptionID string) EventHubsClient {
	return original.NewEventHubsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMigrationConfigListResultIterator(page MigrationConfigListResultPage) MigrationConfigListResultIterator {
	return original.NewMigrationConfigListResultIterator(page)
}
func NewMigrationConfigListResultPage(getNextPage func(context.Context, MigrationConfigListResult) (MigrationConfigListResult, error)) MigrationConfigListResultPage {
	return original.NewMigrationConfigListResultPage(getNextPage)
}
func NewMigrationConfigsClient(subscriptionID string) MigrationConfigsClient {
	return original.NewMigrationConfigsClient(subscriptionID)
}
func NewMigrationConfigsClientWithBaseURI(baseURI string, subscriptionID string) MigrationConfigsClient {
	return original.NewMigrationConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
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
func NewPremiumMessagingRegionsClient(subscriptionID string) PremiumMessagingRegionsClient {
	return original.NewPremiumMessagingRegionsClient(subscriptionID)
}
func NewPremiumMessagingRegionsClientWithBaseURI(baseURI string, subscriptionID string) PremiumMessagingRegionsClient {
	return original.NewPremiumMessagingRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPremiumMessagingRegionsListResultIterator(page PremiumMessagingRegionsListResultPage) PremiumMessagingRegionsListResultIterator {
	return original.NewPremiumMessagingRegionsListResultIterator(page)
}
func NewPremiumMessagingRegionsListResultPage(getNextPage func(context.Context, PremiumMessagingRegionsListResult) (PremiumMessagingRegionsListResult, error)) PremiumMessagingRegionsListResultPage {
	return original.NewPremiumMessagingRegionsListResultPage(getNextPage)
}
func NewQueuesClient(subscriptionID string) QueuesClient {
	return original.NewQueuesClient(subscriptionID)
}
func NewQueuesClientWithBaseURI(baseURI string, subscriptionID string) QueuesClient {
	return original.NewQueuesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegionsClient(subscriptionID string) RegionsClient {
	return original.NewRegionsClient(subscriptionID)
}
func NewRegionsClientWithBaseURI(baseURI string, subscriptionID string) RegionsClient {
	return original.NewRegionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRuleListResultIterator(page RuleListResultPage) RuleListResultIterator {
	return original.NewRuleListResultIterator(page)
}
func NewRuleListResultPage(getNextPage func(context.Context, RuleListResult) (RuleListResult, error)) RuleListResultPage {
	return original.NewRuleListResultPage(getNextPage)
}
func NewRulesClient(subscriptionID string) RulesClient {
	return original.NewRulesClient(subscriptionID)
}
func NewRulesClientWithBaseURI(baseURI string, subscriptionID string) RulesClient {
	return original.NewRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSBAuthorizationRuleListResultIterator(page SBAuthorizationRuleListResultPage) SBAuthorizationRuleListResultIterator {
	return original.NewSBAuthorizationRuleListResultIterator(page)
}
func NewSBAuthorizationRuleListResultPage(getNextPage func(context.Context, SBAuthorizationRuleListResult) (SBAuthorizationRuleListResult, error)) SBAuthorizationRuleListResultPage {
	return original.NewSBAuthorizationRuleListResultPage(getNextPage)
}
func NewSBNamespaceListResultIterator(page SBNamespaceListResultPage) SBNamespaceListResultIterator {
	return original.NewSBNamespaceListResultIterator(page)
}
func NewSBNamespaceListResultPage(getNextPage func(context.Context, SBNamespaceListResult) (SBNamespaceListResult, error)) SBNamespaceListResultPage {
	return original.NewSBNamespaceListResultPage(getNextPage)
}
func NewSBQueueListResultIterator(page SBQueueListResultPage) SBQueueListResultIterator {
	return original.NewSBQueueListResultIterator(page)
}
func NewSBQueueListResultPage(getNextPage func(context.Context, SBQueueListResult) (SBQueueListResult, error)) SBQueueListResultPage {
	return original.NewSBQueueListResultPage(getNextPage)
}
func NewSBSubscriptionListResultIterator(page SBSubscriptionListResultPage) SBSubscriptionListResultIterator {
	return original.NewSBSubscriptionListResultIterator(page)
}
func NewSBSubscriptionListResultPage(getNextPage func(context.Context, SBSubscriptionListResult) (SBSubscriptionListResult, error)) SBSubscriptionListResultPage {
	return original.NewSBSubscriptionListResultPage(getNextPage)
}
func NewSBTopicListResultIterator(page SBTopicListResultPage) SBTopicListResultIterator {
	return original.NewSBTopicListResultIterator(page)
}
func NewSBTopicListResultPage(getNextPage func(context.Context, SBTopicListResult) (SBTopicListResult, error)) SBTopicListResultPage {
	return original.NewSBTopicListResultPage(getNextPage)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopicsClient(subscriptionID string) TopicsClient {
	return original.NewTopicsClient(subscriptionID)
}
func NewTopicsClientWithBaseURI(baseURI string, subscriptionID string) TopicsClient {
	return original.NewTopicsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return original.PossibleEncodingCaptureDescriptionValues()
}
func PossibleEntityStatusValues() []EntityStatus {
	return original.PossibleEntityStatusValues()
}
func PossibleFilterTypeValues() []FilterType {
	return original.PossibleFilterTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleNameSpaceTypeValues() []NameSpaceType {
	return original.PossibleNameSpaceTypeValues()
}
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return original.PossibleNetworkRuleIPActionValues()
}
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return original.PossibleProvisioningStateDRValues()
}
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return original.PossibleRoleDisasterRecoveryValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleUnavailableReasonValues() []UnavailableReason {
	return original.PossibleUnavailableReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
