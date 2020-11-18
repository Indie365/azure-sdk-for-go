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

package datashare

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/datashare/mgmt/2019-11-01/datashare"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DataSetMappingStatus = original.DataSetMappingStatus

const (
	Broken DataSetMappingStatus = original.Broken
	Ok     DataSetMappingStatus = original.Ok
)

type DataSetType = original.DataSetType

const (
	AdlsGen1File       DataSetType = original.AdlsGen1File
	AdlsGen1Folder     DataSetType = original.AdlsGen1Folder
	AdlsGen2File       DataSetType = original.AdlsGen2File
	AdlsGen2FileSystem DataSetType = original.AdlsGen2FileSystem
	AdlsGen2Folder     DataSetType = original.AdlsGen2Folder
	Blob               DataSetType = original.Blob
	BlobFolder         DataSetType = original.BlobFolder
	Container          DataSetType = original.Container
	KustoCluster       DataSetType = original.KustoCluster
	KustoDatabase      DataSetType = original.KustoDatabase
	SQLDBTable         DataSetType = original.SQLDBTable
	SQLDWTable         DataSetType = original.SQLDWTable
)

type InvitationStatus = original.InvitationStatus

const (
	Accepted  InvitationStatus = original.Accepted
	Pending   InvitationStatus = original.Pending
	Rejected  InvitationStatus = original.Rejected
	Withdrawn InvitationStatus = original.Withdrawn
)

type Kind = original.Kind

const (
	KindAdlsGen1File       Kind = original.KindAdlsGen1File
	KindAdlsGen1Folder     Kind = original.KindAdlsGen1Folder
	KindAdlsGen2File       Kind = original.KindAdlsGen2File
	KindAdlsGen2FileSystem Kind = original.KindAdlsGen2FileSystem
	KindAdlsGen2Folder     Kind = original.KindAdlsGen2Folder
	KindBlob               Kind = original.KindBlob
	KindBlobFolder         Kind = original.KindBlobFolder
	KindContainer          Kind = original.KindContainer
	KindDataSet            Kind = original.KindDataSet
	KindKustoCluster       Kind = original.KindKustoCluster
	KindKustoDatabase      Kind = original.KindKustoDatabase
	KindSQLDBTable         Kind = original.KindSQLDBTable
	KindSQLDWTable         Kind = original.KindSQLDWTable
)

type KindBasicDataSetMapping = original.KindBasicDataSetMapping

const (
	KindBasicDataSetMappingKindAdlsGen2File       KindBasicDataSetMapping = original.KindBasicDataSetMappingKindAdlsGen2File
	KindBasicDataSetMappingKindAdlsGen2FileSystem KindBasicDataSetMapping = original.KindBasicDataSetMappingKindAdlsGen2FileSystem
	KindBasicDataSetMappingKindAdlsGen2Folder     KindBasicDataSetMapping = original.KindBasicDataSetMappingKindAdlsGen2Folder
	KindBasicDataSetMappingKindBlob               KindBasicDataSetMapping = original.KindBasicDataSetMappingKindBlob
	KindBasicDataSetMappingKindBlobFolder         KindBasicDataSetMapping = original.KindBasicDataSetMappingKindBlobFolder
	KindBasicDataSetMappingKindContainer          KindBasicDataSetMapping = original.KindBasicDataSetMappingKindContainer
	KindBasicDataSetMappingKindDataSetMapping     KindBasicDataSetMapping = original.KindBasicDataSetMappingKindDataSetMapping
	KindBasicDataSetMappingKindKustoCluster       KindBasicDataSetMapping = original.KindBasicDataSetMappingKindKustoCluster
	KindBasicDataSetMappingKindKustoDatabase      KindBasicDataSetMapping = original.KindBasicDataSetMappingKindKustoDatabase
	KindBasicDataSetMappingKindSQLDBTable         KindBasicDataSetMapping = original.KindBasicDataSetMappingKindSQLDBTable
	KindBasicDataSetMappingKindSQLDWTable         KindBasicDataSetMapping = original.KindBasicDataSetMappingKindSQLDWTable
)

type KindBasicSourceShareSynchronizationSetting = original.KindBasicSourceShareSynchronizationSetting

const (
	KindScheduleBased                     KindBasicSourceShareSynchronizationSetting = original.KindScheduleBased
	KindSourceShareSynchronizationSetting KindBasicSourceShareSynchronizationSetting = original.KindSourceShareSynchronizationSetting
)

type KindBasicSynchronizationSetting = original.KindBasicSynchronizationSetting

const (
	KindBasicSynchronizationSettingKindScheduleBased          KindBasicSynchronizationSetting = original.KindBasicSynchronizationSettingKindScheduleBased
	KindBasicSynchronizationSettingKindSynchronizationSetting KindBasicSynchronizationSetting = original.KindBasicSynchronizationSettingKindSynchronizationSetting
)

type KindBasicTrigger = original.KindBasicTrigger

const (
	KindBasicTriggerKindScheduleBased KindBasicTrigger = original.KindBasicTriggerKindScheduleBased
	KindBasicTriggerKindTrigger       KindBasicTrigger = original.KindBasicTriggerKindTrigger
)

type OutputType = original.OutputType

const (
	Csv     OutputType = original.Csv
	Parquet OutputType = original.Parquet
)

type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Succeeded ProvisioningState = original.Succeeded
)

type RecurrenceInterval = original.RecurrenceInterval

const (
	Day  RecurrenceInterval = original.Day
	Hour RecurrenceInterval = original.Hour
)

type ShareKind = original.ShareKind

const (
	CopyBased ShareKind = original.CopyBased
	InPlace   ShareKind = original.InPlace
)

type ShareSubscriptionStatus = original.ShareSubscriptionStatus

const (
	Active        ShareSubscriptionStatus = original.Active
	Revoked       ShareSubscriptionStatus = original.Revoked
	Revoking      ShareSubscriptionStatus = original.Revoking
	SourceDeleted ShareSubscriptionStatus = original.SourceDeleted
)

type Status = original.Status

const (
	StatusAccepted         Status = original.StatusAccepted
	StatusCanceled         Status = original.StatusCanceled
	StatusFailed           Status = original.StatusFailed
	StatusInProgress       Status = original.StatusInProgress
	StatusSucceeded        Status = original.StatusSucceeded
	StatusTransientFailure Status = original.StatusTransientFailure
)

type SynchronizationMode = original.SynchronizationMode

const (
	FullSync    SynchronizationMode = original.FullSync
	Incremental SynchronizationMode = original.Incremental
)

type TriggerStatus = original.TriggerStatus

const (
	TriggerStatusActive                              TriggerStatus = original.TriggerStatusActive
	TriggerStatusInactive                            TriggerStatus = original.TriggerStatusInactive
	TriggerStatusSourceSynchronizationSettingDeleted TriggerStatus = original.TriggerStatusSourceSynchronizationSettingDeleted
)

type Type = original.Type

const (
	SystemAssigned Type = original.SystemAssigned
)

type ADLSGen1FileDataSet = original.ADLSGen1FileDataSet
type ADLSGen1FileProperties = original.ADLSGen1FileProperties
type ADLSGen1FolderDataSet = original.ADLSGen1FolderDataSet
type ADLSGen1FolderProperties = original.ADLSGen1FolderProperties
type ADLSGen2FileDataSet = original.ADLSGen2FileDataSet
type ADLSGen2FileDataSetMapping = original.ADLSGen2FileDataSetMapping
type ADLSGen2FileDataSetMappingProperties = original.ADLSGen2FileDataSetMappingProperties
type ADLSGen2FileProperties = original.ADLSGen2FileProperties
type ADLSGen2FileSystemDataSet = original.ADLSGen2FileSystemDataSet
type ADLSGen2FileSystemDataSetMapping = original.ADLSGen2FileSystemDataSetMapping
type ADLSGen2FileSystemDataSetMappingProperties = original.ADLSGen2FileSystemDataSetMappingProperties
type ADLSGen2FileSystemProperties = original.ADLSGen2FileSystemProperties
type ADLSGen2FolderDataSet = original.ADLSGen2FolderDataSet
type ADLSGen2FolderDataSetMapping = original.ADLSGen2FolderDataSetMapping
type ADLSGen2FolderDataSetMappingProperties = original.ADLSGen2FolderDataSetMappingProperties
type ADLSGen2FolderProperties = original.ADLSGen2FolderProperties
type Account = original.Account
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type BaseClient = original.BaseClient
type BasicDataSet = original.BasicDataSet
type BasicDataSetMapping = original.BasicDataSetMapping
type BasicSourceShareSynchronizationSetting = original.BasicSourceShareSynchronizationSetting
type BasicSynchronizationSetting = original.BasicSynchronizationSetting
type BasicTrigger = original.BasicTrigger
type BlobContainerDataSet = original.BlobContainerDataSet
type BlobContainerDataSetMapping = original.BlobContainerDataSetMapping
type BlobContainerMappingProperties = original.BlobContainerMappingProperties
type BlobContainerProperties = original.BlobContainerProperties
type BlobDataSet = original.BlobDataSet
type BlobDataSetMapping = original.BlobDataSetMapping
type BlobFolderDataSet = original.BlobFolderDataSet
type BlobFolderDataSetMapping = original.BlobFolderDataSetMapping
type BlobFolderMappingProperties = original.BlobFolderMappingProperties
type BlobFolderProperties = original.BlobFolderProperties
type BlobMappingProperties = original.BlobMappingProperties
type BlobProperties = original.BlobProperties
type ConsumerInvitation = original.ConsumerInvitation
type ConsumerInvitationList = original.ConsumerInvitationList
type ConsumerInvitationListIterator = original.ConsumerInvitationListIterator
type ConsumerInvitationListPage = original.ConsumerInvitationListPage
type ConsumerInvitationProperties = original.ConsumerInvitationProperties
type ConsumerInvitationsClient = original.ConsumerInvitationsClient
type ConsumerSourceDataSet = original.ConsumerSourceDataSet
type ConsumerSourceDataSetList = original.ConsumerSourceDataSetList
type ConsumerSourceDataSetListIterator = original.ConsumerSourceDataSetListIterator
type ConsumerSourceDataSetListPage = original.ConsumerSourceDataSetListPage
type ConsumerSourceDataSetProperties = original.ConsumerSourceDataSetProperties
type ConsumerSourceDataSetsClient = original.ConsumerSourceDataSetsClient
type DataSet = original.DataSet
type DataSetList = original.DataSetList
type DataSetListIterator = original.DataSetListIterator
type DataSetListPage = original.DataSetListPage
type DataSetMapping = original.DataSetMapping
type DataSetMappingList = original.DataSetMappingList
type DataSetMappingListIterator = original.DataSetMappingListIterator
type DataSetMappingListPage = original.DataSetMappingListPage
type DataSetMappingModel = original.DataSetMappingModel
type DataSetMappingsClient = original.DataSetMappingsClient
type DataSetModel = original.DataSetModel
type DataSetsClient = original.DataSetsClient
type DataSetsDeleteFuture = original.DataSetsDeleteFuture
type DefaultDto = original.DefaultDto
type DimensionProperties = original.DimensionProperties
type Error = original.Error
type ErrorInfo = original.ErrorInfo
type Identity = original.Identity
type Invitation = original.Invitation
type InvitationList = original.InvitationList
type InvitationListIterator = original.InvitationListIterator
type InvitationListPage = original.InvitationListPage
type InvitationProperties = original.InvitationProperties
type InvitationsClient = original.InvitationsClient
type KustoClusterDataSet = original.KustoClusterDataSet
type KustoClusterDataSetMapping = original.KustoClusterDataSetMapping
type KustoClusterDataSetMappingProperties = original.KustoClusterDataSetMappingProperties
type KustoClusterDataSetProperties = original.KustoClusterDataSetProperties
type KustoDatabaseDataSet = original.KustoDatabaseDataSet
type KustoDatabaseDataSetMapping = original.KustoDatabaseDataSetMapping
type KustoDatabaseDataSetMappingProperties = original.KustoDatabaseDataSetMappingProperties
type KustoDatabaseDataSetProperties = original.KustoDatabaseDataSetProperties
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationMetaLogSpecification = original.OperationMetaLogSpecification
type OperationMetaMetricSpecification = original.OperationMetaMetricSpecification
type OperationMetaPropertyInfo = original.OperationMetaPropertyInfo
type OperationMetaServiceSpecification = original.OperationMetaServiceSpecification
type OperationModel = original.OperationModel
type OperationModelProperties = original.OperationModelProperties
type OperationResponse = original.OperationResponse
type OperationsClient = original.OperationsClient
type ProviderShareSubscription = original.ProviderShareSubscription
type ProviderShareSubscriptionList = original.ProviderShareSubscriptionList
type ProviderShareSubscriptionListIterator = original.ProviderShareSubscriptionListIterator
type ProviderShareSubscriptionListPage = original.ProviderShareSubscriptionListPage
type ProviderShareSubscriptionProperties = original.ProviderShareSubscriptionProperties
type ProviderShareSubscriptionsClient = original.ProviderShareSubscriptionsClient
type ProviderShareSubscriptionsRevokeFuture = original.ProviderShareSubscriptionsRevokeFuture
type ProxyDto = original.ProxyDto
type SQLDBTableDataSet = original.SQLDBTableDataSet
type SQLDBTableDataSetMapping = original.SQLDBTableDataSetMapping
type SQLDBTableDataSetMappingProperties = original.SQLDBTableDataSetMappingProperties
type SQLDBTableProperties = original.SQLDBTableProperties
type SQLDWTableDataSet = original.SQLDWTableDataSet
type SQLDWTableDataSetMapping = original.SQLDWTableDataSetMapping
type SQLDWTableDataSetMappingProperties = original.SQLDWTableDataSetMappingProperties
type SQLDWTableProperties = original.SQLDWTableProperties
type ScheduledSourceShareSynchronizationSettingProperties = original.ScheduledSourceShareSynchronizationSettingProperties
type ScheduledSourceSynchronizationSetting = original.ScheduledSourceSynchronizationSetting
type ScheduledSynchronizationSetting = original.ScheduledSynchronizationSetting
type ScheduledSynchronizationSettingProperties = original.ScheduledSynchronizationSettingProperties
type ScheduledTrigger = original.ScheduledTrigger
type ScheduledTriggerProperties = original.ScheduledTriggerProperties
type Share = original.Share
type ShareList = original.ShareList
type ShareListIterator = original.ShareListIterator
type ShareListPage = original.ShareListPage
type ShareProperties = original.ShareProperties
type ShareSubscription = original.ShareSubscription
type ShareSubscriptionList = original.ShareSubscriptionList
type ShareSubscriptionListIterator = original.ShareSubscriptionListIterator
type ShareSubscriptionListPage = original.ShareSubscriptionListPage
type ShareSubscriptionProperties = original.ShareSubscriptionProperties
type ShareSubscriptionSynchronization = original.ShareSubscriptionSynchronization
type ShareSubscriptionSynchronizationList = original.ShareSubscriptionSynchronizationList
type ShareSubscriptionSynchronizationListIterator = original.ShareSubscriptionSynchronizationListIterator
type ShareSubscriptionSynchronizationListPage = original.ShareSubscriptionSynchronizationListPage
type ShareSubscriptionsCancelSynchronizationFuture = original.ShareSubscriptionsCancelSynchronizationFuture
type ShareSubscriptionsClient = original.ShareSubscriptionsClient
type ShareSubscriptionsDeleteFuture = original.ShareSubscriptionsDeleteFuture
type ShareSubscriptionsSynchronizeMethodFuture = original.ShareSubscriptionsSynchronizeMethodFuture
type ShareSynchronization = original.ShareSynchronization
type ShareSynchronizationList = original.ShareSynchronizationList
type ShareSynchronizationListIterator = original.ShareSynchronizationListIterator
type ShareSynchronizationListPage = original.ShareSynchronizationListPage
type SharesClient = original.SharesClient
type SharesDeleteFuture = original.SharesDeleteFuture
type SourceShareSynchronizationSetting = original.SourceShareSynchronizationSetting
type SourceShareSynchronizationSettingList = original.SourceShareSynchronizationSettingList
type SourceShareSynchronizationSettingListIterator = original.SourceShareSynchronizationSettingListIterator
type SourceShareSynchronizationSettingListPage = original.SourceShareSynchronizationSettingListPage
type SynchronizationDetails = original.SynchronizationDetails
type SynchronizationDetailsList = original.SynchronizationDetailsList
type SynchronizationDetailsListIterator = original.SynchronizationDetailsListIterator
type SynchronizationDetailsListPage = original.SynchronizationDetailsListPage
type SynchronizationSetting = original.SynchronizationSetting
type SynchronizationSettingList = original.SynchronizationSettingList
type SynchronizationSettingListIterator = original.SynchronizationSettingListIterator
type SynchronizationSettingListPage = original.SynchronizationSettingListPage
type SynchronizationSettingModel = original.SynchronizationSettingModel
type SynchronizationSettingsClient = original.SynchronizationSettingsClient
type SynchronizationSettingsDeleteFuture = original.SynchronizationSettingsDeleteFuture
type Synchronize = original.Synchronize
type Trigger = original.Trigger
type TriggerList = original.TriggerList
type TriggerListIterator = original.TriggerListIterator
type TriggerListPage = original.TriggerListPage
type TriggerModel = original.TriggerModel
type TriggersClient = original.TriggersClient
type TriggersCreateFuture = original.TriggersCreateFuture
type TriggersDeleteFuture = original.TriggersDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return original.NewAccountListIterator(page)
}
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return original.NewAccountListPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConsumerInvitationListIterator(page ConsumerInvitationListPage) ConsumerInvitationListIterator {
	return original.NewConsumerInvitationListIterator(page)
}
func NewConsumerInvitationListPage(cur ConsumerInvitationList, getNextPage func(context.Context, ConsumerInvitationList) (ConsumerInvitationList, error)) ConsumerInvitationListPage {
	return original.NewConsumerInvitationListPage(cur, getNextPage)
}
func NewConsumerInvitationsClient(subscriptionID string) ConsumerInvitationsClient {
	return original.NewConsumerInvitationsClient(subscriptionID)
}
func NewConsumerInvitationsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerInvitationsClient {
	return original.NewConsumerInvitationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConsumerSourceDataSetListIterator(page ConsumerSourceDataSetListPage) ConsumerSourceDataSetListIterator {
	return original.NewConsumerSourceDataSetListIterator(page)
}
func NewConsumerSourceDataSetListPage(cur ConsumerSourceDataSetList, getNextPage func(context.Context, ConsumerSourceDataSetList) (ConsumerSourceDataSetList, error)) ConsumerSourceDataSetListPage {
	return original.NewConsumerSourceDataSetListPage(cur, getNextPage)
}
func NewConsumerSourceDataSetsClient(subscriptionID string) ConsumerSourceDataSetsClient {
	return original.NewConsumerSourceDataSetsClient(subscriptionID)
}
func NewConsumerSourceDataSetsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerSourceDataSetsClient {
	return original.NewConsumerSourceDataSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataSetListIterator(page DataSetListPage) DataSetListIterator {
	return original.NewDataSetListIterator(page)
}
func NewDataSetListPage(cur DataSetList, getNextPage func(context.Context, DataSetList) (DataSetList, error)) DataSetListPage {
	return original.NewDataSetListPage(cur, getNextPage)
}
func NewDataSetMappingListIterator(page DataSetMappingListPage) DataSetMappingListIterator {
	return original.NewDataSetMappingListIterator(page)
}
func NewDataSetMappingListPage(cur DataSetMappingList, getNextPage func(context.Context, DataSetMappingList) (DataSetMappingList, error)) DataSetMappingListPage {
	return original.NewDataSetMappingListPage(cur, getNextPage)
}
func NewDataSetMappingsClient(subscriptionID string) DataSetMappingsClient {
	return original.NewDataSetMappingsClient(subscriptionID)
}
func NewDataSetMappingsClientWithBaseURI(baseURI string, subscriptionID string) DataSetMappingsClient {
	return original.NewDataSetMappingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataSetsClient(subscriptionID string) DataSetsClient {
	return original.NewDataSetsClient(subscriptionID)
}
func NewDataSetsClientWithBaseURI(baseURI string, subscriptionID string) DataSetsClient {
	return original.NewDataSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInvitationListIterator(page InvitationListPage) InvitationListIterator {
	return original.NewInvitationListIterator(page)
}
func NewInvitationListPage(cur InvitationList, getNextPage func(context.Context, InvitationList) (InvitationList, error)) InvitationListPage {
	return original.NewInvitationListPage(cur, getNextPage)
}
func NewInvitationsClient(subscriptionID string) InvitationsClient {
	return original.NewInvitationsClient(subscriptionID)
}
func NewInvitationsClientWithBaseURI(baseURI string, subscriptionID string) InvitationsClient {
	return original.NewInvitationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProviderShareSubscriptionListIterator(page ProviderShareSubscriptionListPage) ProviderShareSubscriptionListIterator {
	return original.NewProviderShareSubscriptionListIterator(page)
}
func NewProviderShareSubscriptionListPage(cur ProviderShareSubscriptionList, getNextPage func(context.Context, ProviderShareSubscriptionList) (ProviderShareSubscriptionList, error)) ProviderShareSubscriptionListPage {
	return original.NewProviderShareSubscriptionListPage(cur, getNextPage)
}
func NewProviderShareSubscriptionsClient(subscriptionID string) ProviderShareSubscriptionsClient {
	return original.NewProviderShareSubscriptionsClient(subscriptionID)
}
func NewProviderShareSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) ProviderShareSubscriptionsClient {
	return original.NewProviderShareSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewShareListIterator(page ShareListPage) ShareListIterator {
	return original.NewShareListIterator(page)
}
func NewShareListPage(cur ShareList, getNextPage func(context.Context, ShareList) (ShareList, error)) ShareListPage {
	return original.NewShareListPage(cur, getNextPage)
}
func NewShareSubscriptionListIterator(page ShareSubscriptionListPage) ShareSubscriptionListIterator {
	return original.NewShareSubscriptionListIterator(page)
}
func NewShareSubscriptionListPage(cur ShareSubscriptionList, getNextPage func(context.Context, ShareSubscriptionList) (ShareSubscriptionList, error)) ShareSubscriptionListPage {
	return original.NewShareSubscriptionListPage(cur, getNextPage)
}
func NewShareSubscriptionSynchronizationListIterator(page ShareSubscriptionSynchronizationListPage) ShareSubscriptionSynchronizationListIterator {
	return original.NewShareSubscriptionSynchronizationListIterator(page)
}
func NewShareSubscriptionSynchronizationListPage(cur ShareSubscriptionSynchronizationList, getNextPage func(context.Context, ShareSubscriptionSynchronizationList) (ShareSubscriptionSynchronizationList, error)) ShareSubscriptionSynchronizationListPage {
	return original.NewShareSubscriptionSynchronizationListPage(cur, getNextPage)
}
func NewShareSubscriptionsClient(subscriptionID string) ShareSubscriptionsClient {
	return original.NewShareSubscriptionsClient(subscriptionID)
}
func NewShareSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) ShareSubscriptionsClient {
	return original.NewShareSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewShareSynchronizationListIterator(page ShareSynchronizationListPage) ShareSynchronizationListIterator {
	return original.NewShareSynchronizationListIterator(page)
}
func NewShareSynchronizationListPage(cur ShareSynchronizationList, getNextPage func(context.Context, ShareSynchronizationList) (ShareSynchronizationList, error)) ShareSynchronizationListPage {
	return original.NewShareSynchronizationListPage(cur, getNextPage)
}
func NewSharesClient(subscriptionID string) SharesClient {
	return original.NewSharesClient(subscriptionID)
}
func NewSharesClientWithBaseURI(baseURI string, subscriptionID string) SharesClient {
	return original.NewSharesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSourceShareSynchronizationSettingListIterator(page SourceShareSynchronizationSettingListPage) SourceShareSynchronizationSettingListIterator {
	return original.NewSourceShareSynchronizationSettingListIterator(page)
}
func NewSourceShareSynchronizationSettingListPage(cur SourceShareSynchronizationSettingList, getNextPage func(context.Context, SourceShareSynchronizationSettingList) (SourceShareSynchronizationSettingList, error)) SourceShareSynchronizationSettingListPage {
	return original.NewSourceShareSynchronizationSettingListPage(cur, getNextPage)
}
func NewSynchronizationDetailsListIterator(page SynchronizationDetailsListPage) SynchronizationDetailsListIterator {
	return original.NewSynchronizationDetailsListIterator(page)
}
func NewSynchronizationDetailsListPage(cur SynchronizationDetailsList, getNextPage func(context.Context, SynchronizationDetailsList) (SynchronizationDetailsList, error)) SynchronizationDetailsListPage {
	return original.NewSynchronizationDetailsListPage(cur, getNextPage)
}
func NewSynchronizationSettingListIterator(page SynchronizationSettingListPage) SynchronizationSettingListIterator {
	return original.NewSynchronizationSettingListIterator(page)
}
func NewSynchronizationSettingListPage(cur SynchronizationSettingList, getNextPage func(context.Context, SynchronizationSettingList) (SynchronizationSettingList, error)) SynchronizationSettingListPage {
	return original.NewSynchronizationSettingListPage(cur, getNextPage)
}
func NewSynchronizationSettingsClient(subscriptionID string) SynchronizationSettingsClient {
	return original.NewSynchronizationSettingsClient(subscriptionID)
}
func NewSynchronizationSettingsClientWithBaseURI(baseURI string, subscriptionID string) SynchronizationSettingsClient {
	return original.NewSynchronizationSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTriggerListIterator(page TriggerListPage) TriggerListIterator {
	return original.NewTriggerListIterator(page)
}
func NewTriggersClient(subscriptionID string) TriggersClient {
	return original.NewTriggersClient(subscriptionID)
}
func NewTriggersClientWithBaseURI(baseURI string, subscriptionID string) TriggersClient {
	return original.NewTriggersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDataSetMappingStatusValues() []DataSetMappingStatus {
	return original.PossibleDataSetMappingStatusValues()
}
func PossibleDataSetTypeValues() []DataSetType {
	return original.PossibleDataSetTypeValues()
}
func PossibleInvitationStatusValues() []InvitationStatus {
	return original.PossibleInvitationStatusValues()
}
func PossibleKindBasicDataSetMappingValues() []KindBasicDataSetMapping {
	return original.PossibleKindBasicDataSetMappingValues()
}
func PossibleKindBasicSourceShareSynchronizationSettingValues() []KindBasicSourceShareSynchronizationSetting {
	return original.PossibleKindBasicSourceShareSynchronizationSettingValues()
}
func PossibleKindBasicSynchronizationSettingValues() []KindBasicSynchronizationSetting {
	return original.PossibleKindBasicSynchronizationSettingValues()
}
func PossibleKindBasicTriggerValues() []KindBasicTrigger {
	return original.PossibleKindBasicTriggerValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleOutputTypeValues() []OutputType {
	return original.PossibleOutputTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRecurrenceIntervalValues() []RecurrenceInterval {
	return original.PossibleRecurrenceIntervalValues()
}
func PossibleShareKindValues() []ShareKind {
	return original.PossibleShareKindValues()
}
func PossibleShareSubscriptionStatusValues() []ShareSubscriptionStatus {
	return original.PossibleShareSubscriptionStatusValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleSynchronizationModeValues() []SynchronizationMode {
	return original.PossibleSynchronizationModeValues()
}
func PossibleTriggerStatusValues() []TriggerStatus {
	return original.PossibleTriggerStatusValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
