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

package kusto

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2019-01-21/kusto"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AzureScaleType = original.AzureScaleType

const (
	Automatic AzureScaleType = original.Automatic
	Manual    AzureScaleType = original.Manual
	None      AzureScaleType = original.None
)

type AzureSkuName = original.AzureSkuName

const (
	DevNoSLAStandardD11V2 AzureSkuName = original.DevNoSLAStandardD11V2
	StandardD11V2         AzureSkuName = original.StandardD11V2
	StandardD12V2         AzureSkuName = original.StandardD12V2
	StandardD13V2         AzureSkuName = original.StandardD13V2
	StandardD14V2         AzureSkuName = original.StandardD14V2
	StandardDS13V21TBPS   AzureSkuName = original.StandardDS13V21TBPS
	StandardDS13V22TBPS   AzureSkuName = original.StandardDS13V22TBPS
	StandardDS14V23TBPS   AzureSkuName = original.StandardDS14V23TBPS
	StandardDS14V24TBPS   AzureSkuName = original.StandardDS14V24TBPS
	StandardL16s          AzureSkuName = original.StandardL16s
	StandardL4s           AzureSkuName = original.StandardL4s
	StandardL8s           AzureSkuName = original.StandardL8s
)

type AzureSkuTier = original.AzureSkuTier

const (
	Basic    AzureSkuTier = original.Basic
	Standard AzureSkuTier = original.Standard
)

type DataFormat = original.DataFormat

const (
	AVRO       DataFormat = original.AVRO
	CSV        DataFormat = original.CSV
	JSON       DataFormat = original.JSON
	MULTIJSON  DataFormat = original.MULTIJSON
	PSV        DataFormat = original.PSV
	RAW        DataFormat = original.RAW
	SCSV       DataFormat = original.SCSV
	SINGLEJSON DataFormat = original.SINGLEJSON
	SOHSV      DataFormat = original.SOHSV
	TSV        DataFormat = original.TSV
	TXT        DataFormat = original.TXT
)

type DatabasePrincipalRole = original.DatabasePrincipalRole

const (
	Admin               DatabasePrincipalRole = original.Admin
	Ingestor            DatabasePrincipalRole = original.Ingestor
	Monitor             DatabasePrincipalRole = original.Monitor
	UnrestrictedViewers DatabasePrincipalRole = original.UnrestrictedViewers
	User                DatabasePrincipalRole = original.User
	Viewer              DatabasePrincipalRole = original.Viewer
)

type DatabasePrincipalType = original.DatabasePrincipalType

const (
	DatabasePrincipalTypeApp   DatabasePrincipalType = original.DatabasePrincipalTypeApp
	DatabasePrincipalTypeGroup DatabasePrincipalType = original.DatabasePrincipalTypeGroup
	DatabasePrincipalTypeUser  DatabasePrincipalType = original.DatabasePrincipalTypeUser
)

type Kind = original.Kind

const (
	KindDatabase         Kind = original.KindDatabase
	KindReadOnlyAttached Kind = original.KindReadOnlyAttached
	KindReadWrite        Kind = original.KindReadWrite
)

type KindBasicDataConnection = original.KindBasicDataConnection

const (
	KindDataConnection KindBasicDataConnection = original.KindDataConnection
	KindEventGrid      KindBasicDataConnection = original.KindEventGrid
	KindEventHub       KindBasicDataConnection = original.KindEventHub
)

type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type State = original.State

const (
	StateCreating    State = original.StateCreating
	StateDeleted     State = original.StateDeleted
	StateDeleting    State = original.StateDeleting
	StateRunning     State = original.StateRunning
	StateStarting    State = original.StateStarting
	StateStopped     State = original.StateStopped
	StateStopping    State = original.StateStopping
	StateUnavailable State = original.StateUnavailable
	StateUpdating    State = original.StateUpdating
)

type Type = original.Type

const (
	MicrosoftKustoclustersattachedDatabaseConfigurations Type = original.MicrosoftKustoclustersattachedDatabaseConfigurations
	MicrosoftKustoclustersdatabases                      Type = original.MicrosoftKustoclustersdatabases
)

type AttachedDatabaseConfiguration = original.AttachedDatabaseConfiguration
type AttachedDatabaseConfigurationListResult = original.AttachedDatabaseConfigurationListResult
type AttachedDatabaseConfigurationProperties = original.AttachedDatabaseConfigurationProperties
type AttachedDatabaseConfigurationsClient = original.AttachedDatabaseConfigurationsClient
type AttachedDatabaseConfigurationsCreateOrUpdateFuture = original.AttachedDatabaseConfigurationsCreateOrUpdateFuture
type AttachedDatabaseConfigurationsDeleteFuture = original.AttachedDatabaseConfigurationsDeleteFuture
type AzureCapacity = original.AzureCapacity
type AzureEntityResource = original.AzureEntityResource
type AzureResourceSku = original.AzureResourceSku
type AzureSku = original.AzureSku
type BaseClient = original.BaseClient
type BasicDataConnection = original.BasicDataConnection
type BasicDatabase = original.BasicDatabase
type CheckNameRequest = original.CheckNameRequest
type CheckNameResult = original.CheckNameResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Cluster = original.Cluster
type ClusterCheckNameRequest = original.ClusterCheckNameRequest
type ClusterListResult = original.ClusterListResult
type ClusterProperties = original.ClusterProperties
type ClusterUpdate = original.ClusterUpdate
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersStartFuture = original.ClustersStartFuture
type ClustersStopFuture = original.ClustersStopFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type DataConnection = original.DataConnection
type DataConnectionCheckNameRequest = original.DataConnectionCheckNameRequest
type DataConnectionListResult = original.DataConnectionListResult
type DataConnectionModel = original.DataConnectionModel
type DataConnectionValidation = original.DataConnectionValidation
type DataConnectionValidationListResult = original.DataConnectionValidationListResult
type DataConnectionValidationResult = original.DataConnectionValidationResult
type DataConnectionsClient = original.DataConnectionsClient
type DataConnectionsCreateOrUpdateFuture = original.DataConnectionsCreateOrUpdateFuture
type DataConnectionsDeleteFuture = original.DataConnectionsDeleteFuture
type DataConnectionsUpdateFuture = original.DataConnectionsUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseModel = original.DatabaseModel
type DatabasePrincipal = original.DatabasePrincipal
type DatabasePrincipalListRequest = original.DatabasePrincipalListRequest
type DatabasePrincipalListResult = original.DatabasePrincipalListResult
type DatabaseStatistics = original.DatabaseStatistics
type DatabasesClient = original.DatabasesClient
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type DatabasesUpdateFuture = original.DatabasesUpdateFuture
type EventGridConnectionProperties = original.EventGridConnectionProperties
type EventGridDataConnection = original.EventGridDataConnection
type EventHubConnectionProperties = original.EventHubConnectionProperties
type EventHubDataConnection = original.EventHubDataConnection
type FollowerDatabaseListResult = original.FollowerDatabaseListResult
type FollowerDatabaseRequest = original.FollowerDatabaseRequest
type FollowerDatabaseResult = original.FollowerDatabaseResult
type IntelligentAutoscale = original.IntelligentAutoscale
type ListResourceSkusResult = original.ListResourceSkusResult
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type ReadOnlyAttachedDatabase = original.ReadOnlyAttachedDatabase
type ReadOnlyAttachedDatabaseProperties = original.ReadOnlyAttachedDatabaseProperties
type ReadWriteDatabase = original.ReadWriteDatabase
type ReadWriteDatabaseProperties = original.ReadWriteDatabaseProperties
type Resource = original.Resource
type SkuDescription = original.SkuDescription
type SkuDescriptionList = original.SkuDescriptionList
type SkuLocationInfoItem = original.SkuLocationInfoItem
type TrackedResource = original.TrackedResource
type TrustedExternalTenant = original.TrustedExternalTenant

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAttachedDatabaseConfigurationsClient(subscriptionID string) AttachedDatabaseConfigurationsClient {
	return original.NewAttachedDatabaseConfigurationsClient(subscriptionID)
}
func NewAttachedDatabaseConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) AttachedDatabaseConfigurationsClient {
	return original.NewAttachedDatabaseConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataConnectionsClient(subscriptionID string) DataConnectionsClient {
	return original.NewDataConnectionsClient(subscriptionID)
}
func NewDataConnectionsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectionsClient {
	return original.NewDataConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
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
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAzureScaleTypeValues() []AzureScaleType {
	return original.PossibleAzureScaleTypeValues()
}
func PossibleAzureSkuNameValues() []AzureSkuName {
	return original.PossibleAzureSkuNameValues()
}
func PossibleAzureSkuTierValues() []AzureSkuTier {
	return original.PossibleAzureSkuTierValues()
}
func PossibleDataFormatValues() []DataFormat {
	return original.PossibleDataFormatValues()
}
func PossibleDatabasePrincipalRoleValues() []DatabasePrincipalRole {
	return original.PossibleDatabasePrincipalRoleValues()
}
func PossibleDatabasePrincipalTypeValues() []DatabasePrincipalType {
	return original.PossibleDatabasePrincipalTypeValues()
}
func PossibleKindBasicDataConnectionValues() []KindBasicDataConnection {
	return original.PossibleKindBasicDataConnectionValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
