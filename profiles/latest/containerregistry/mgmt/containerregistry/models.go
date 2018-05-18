// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package containerregistry

import original "github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2017-10-01/containerregistry"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ImportMode = original.ImportMode

const (
	Force   ImportMode = original.Force
	NoForce ImportMode = original.NoForce
)

type PasswordName = original.PasswordName

const (
	Password  PasswordName = original.Password
	Password2 PasswordName = original.Password2
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type RegistryUsageUnit = original.RegistryUsageUnit

const (
	Bytes RegistryUsageUnit = original.Bytes
	Count RegistryUsageUnit = original.Count
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Classic  SkuName = original.Classic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierClassic  SkuTier = original.SkuTierClassic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type WebhookAction = original.WebhookAction

const (
	Delete     WebhookAction = original.Delete
	Push       WebhookAction = original.Push
	Quarantine WebhookAction = original.Quarantine
)

type WebhookStatus = original.WebhookStatus

const (
	Disabled WebhookStatus = original.Disabled
	Enabled  WebhookStatus = original.Enabled
)

type Actor = original.Actor
type CallbackConfig = original.CallbackConfig
type Event = original.Event
type EventContent = original.EventContent
type EventInfo = original.EventInfo
type EventListResult = original.EventListResult
type EventListResultIterator = original.EventListResultIterator
type EventListResultPage = original.EventListResultPage
type EventRequestMessage = original.EventRequestMessage
type EventResponseMessage = original.EventResponseMessage
type ImportImageParameters = original.ImportImageParameters
type ImportSource = original.ImportSource
type OperationDefinition = original.OperationDefinition
type OperationDisplayDefinition = original.OperationDisplayDefinition
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type RegenerateCredentialParameters = original.RegenerateCredentialParameters
type RegistriesCreateFuture = original.RegistriesCreateFuture
type RegistriesDeleteFuture = original.RegistriesDeleteFuture
type RegistriesImportImageFuture = original.RegistriesImportImageFuture
type RegistriesUpdateFuture = original.RegistriesUpdateFuture
type Registry = original.Registry
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type RegistryListResult = original.RegistryListResult
type RegistryListResultIterator = original.RegistryListResultIterator
type RegistryListResultPage = original.RegistryListResultPage
type RegistryNameCheckRequest = original.RegistryNameCheckRequest
type RegistryNameStatus = original.RegistryNameStatus
type RegistryPassword = original.RegistryPassword
type RegistryProperties = original.RegistryProperties
type RegistryPropertiesUpdateParameters = original.RegistryPropertiesUpdateParameters
type RegistryUpdateParameters = original.RegistryUpdateParameters
type RegistryUsage = original.RegistryUsage
type RegistryUsageListResult = original.RegistryUsageListResult
type Replication = original.Replication
type ReplicationListResult = original.ReplicationListResult
type ReplicationListResultIterator = original.ReplicationListResultIterator
type ReplicationListResultPage = original.ReplicationListResultPage
type ReplicationProperties = original.ReplicationProperties
type ReplicationsCreateFuture = original.ReplicationsCreateFuture
type ReplicationsDeleteFuture = original.ReplicationsDeleteFuture
type ReplicationsUpdateFuture = original.ReplicationsUpdateFuture
type ReplicationUpdateParameters = original.ReplicationUpdateParameters
type Request = original.Request
type Resource = original.Resource
type Sku = original.Sku
type Source = original.Source
type Status = original.Status
type StorageAccountProperties = original.StorageAccountProperties
type Target = original.Target
type Webhook = original.Webhook
type WebhookCreateParameters = original.WebhookCreateParameters
type WebhookListResult = original.WebhookListResult
type WebhookListResultIterator = original.WebhookListResultIterator
type WebhookListResultPage = original.WebhookListResultPage
type WebhookProperties = original.WebhookProperties
type WebhookPropertiesCreateParameters = original.WebhookPropertiesCreateParameters
type WebhookPropertiesUpdateParameters = original.WebhookPropertiesUpdateParameters
type WebhooksCreateFuture = original.WebhooksCreateFuture
type WebhooksDeleteFuture = original.WebhooksDeleteFuture
type WebhooksUpdateFuture = original.WebhooksUpdateFuture
type WebhookUpdateParameters = original.WebhookUpdateParameters
type OperationsClient = original.OperationsClient
type RegistriesClient = original.RegistriesClient
type ReplicationsClient = original.ReplicationsClient
type WebhooksClient = original.WebhooksClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleImportModeValues() []ImportMode {
	return original.PossibleImportModeValues()
}
func PossiblePasswordNameValues() []PasswordName {
	return original.PossiblePasswordNameValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return original.PossibleRegistryUsageUnitValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleWebhookActionValues() []WebhookAction {
	return original.PossibleWebhookActionValues()
}
func PossibleWebhookStatusValues() []WebhookStatus {
	return original.PossibleWebhookStatusValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return original.NewRegistriesClient(subscriptionID)
}
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return original.NewRegistriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationsClient(subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClient(subscriptionID)
}
func NewReplicationsClientWithBaseURI(baseURI string, subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func NewWebhooksClient(subscriptionID string) WebhooksClient {
	return original.NewWebhooksClient(subscriptionID)
}
func NewWebhooksClientWithBaseURI(baseURI string, subscriptionID string) WebhooksClient {
	return original.NewWebhooksClientWithBaseURI(baseURI, subscriptionID)
}
