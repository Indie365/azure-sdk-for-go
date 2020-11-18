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

package compute

import original "github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AgentVMSizeTypes = original.AgentVMSizeTypes

const (
	StandardA0    AgentVMSizeTypes = original.StandardA0
	StandardA1    AgentVMSizeTypes = original.StandardA1
	StandardA10   AgentVMSizeTypes = original.StandardA10
	StandardA11   AgentVMSizeTypes = original.StandardA11
	StandardA2    AgentVMSizeTypes = original.StandardA2
	StandardA3    AgentVMSizeTypes = original.StandardA3
	StandardA4    AgentVMSizeTypes = original.StandardA4
	StandardA5    AgentVMSizeTypes = original.StandardA5
	StandardA6    AgentVMSizeTypes = original.StandardA6
	StandardA7    AgentVMSizeTypes = original.StandardA7
	StandardA8    AgentVMSizeTypes = original.StandardA8
	StandardA9    AgentVMSizeTypes = original.StandardA9
	StandardD1    AgentVMSizeTypes = original.StandardD1
	StandardD11   AgentVMSizeTypes = original.StandardD11
	StandardD11V2 AgentVMSizeTypes = original.StandardD11V2
	StandardD12   AgentVMSizeTypes = original.StandardD12
	StandardD12V2 AgentVMSizeTypes = original.StandardD12V2
	StandardD13   AgentVMSizeTypes = original.StandardD13
	StandardD13V2 AgentVMSizeTypes = original.StandardD13V2
	StandardD14   AgentVMSizeTypes = original.StandardD14
	StandardD14V2 AgentVMSizeTypes = original.StandardD14V2
	StandardD1V2  AgentVMSizeTypes = original.StandardD1V2
	StandardD2    AgentVMSizeTypes = original.StandardD2
	StandardD2V2  AgentVMSizeTypes = original.StandardD2V2
	StandardD3    AgentVMSizeTypes = original.StandardD3
	StandardD3V2  AgentVMSizeTypes = original.StandardD3V2
	StandardD4    AgentVMSizeTypes = original.StandardD4
	StandardD4V2  AgentVMSizeTypes = original.StandardD4V2
	StandardD5V2  AgentVMSizeTypes = original.StandardD5V2
	StandardDS1   AgentVMSizeTypes = original.StandardDS1
	StandardDS11  AgentVMSizeTypes = original.StandardDS11
	StandardDS12  AgentVMSizeTypes = original.StandardDS12
	StandardDS13  AgentVMSizeTypes = original.StandardDS13
	StandardDS14  AgentVMSizeTypes = original.StandardDS14
	StandardDS2   AgentVMSizeTypes = original.StandardDS2
	StandardDS3   AgentVMSizeTypes = original.StandardDS3
	StandardDS4   AgentVMSizeTypes = original.StandardDS4
	StandardG1    AgentVMSizeTypes = original.StandardG1
	StandardG2    AgentVMSizeTypes = original.StandardG2
	StandardG3    AgentVMSizeTypes = original.StandardG3
	StandardG4    AgentVMSizeTypes = original.StandardG4
	StandardG5    AgentVMSizeTypes = original.StandardG5
	StandardGS1   AgentVMSizeTypes = original.StandardGS1
	StandardGS2   AgentVMSizeTypes = original.StandardGS2
	StandardGS3   AgentVMSizeTypes = original.StandardGS3
	StandardGS4   AgentVMSizeTypes = original.StandardGS4
	StandardGS5   AgentVMSizeTypes = original.StandardGS5
)

type ClusterType = original.ClusterType

const (
	ACS   ClusterType = original.ACS
	Local ClusterType = original.Local
)

type OperationStatus = original.OperationStatus

const (
	Canceled  OperationStatus = original.Canceled
	Creating  OperationStatus = original.Creating
	Deleting  OperationStatus = original.Deleting
	Failed    OperationStatus = original.Failed
	Succeeded OperationStatus = original.Succeeded
	Unknown   OperationStatus = original.Unknown
	Updating  OperationStatus = original.Updating
)

type OrchestratorType = original.OrchestratorType

const (
	Kubernetes OrchestratorType = original.Kubernetes
	None       OrchestratorType = original.None
)

type Status = original.Status

const (
	Disabled Status = original.Disabled
	Enabled  Status = original.Enabled
)

type SystemServiceType = original.SystemServiceType

const (
	SystemServiceTypeBatchFrontEnd   SystemServiceType = original.SystemServiceTypeBatchFrontEnd
	SystemServiceTypeNone            SystemServiceType = original.SystemServiceTypeNone
	SystemServiceTypeScoringFrontEnd SystemServiceType = original.SystemServiceTypeScoringFrontEnd
)

type UpdatesAvailable = original.UpdatesAvailable

const (
	No  UpdatesAvailable = original.No
	Yes UpdatesAvailable = original.Yes
)

type AcsClusterProperties = original.AcsClusterProperties
type AppInsightsCredentials = original.AppInsightsCredentials
type AppInsightsProperties = original.AppInsightsProperties
type AutoScaleConfiguration = original.AutoScaleConfiguration
type AvailableOperations = original.AvailableOperations
type BaseClient = original.BaseClient
type CheckSystemServicesUpdatesAvailableResponse = original.CheckSystemServicesUpdatesAvailableResponse
type ContainerRegistryCredentials = original.ContainerRegistryCredentials
type ContainerRegistryProperties = original.ContainerRegistryProperties
type ContainerServiceCredentials = original.ContainerServiceCredentials
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type ErrorResponseWrapper = original.ErrorResponseWrapper
type GlobalServiceConfiguration = original.GlobalServiceConfiguration
type KubernetesClusterProperties = original.KubernetesClusterProperties
type MachineLearningComputeClient = original.MachineLearningComputeClient
type OperationalizationCluster = original.OperationalizationCluster
type OperationalizationClusterCredentials = original.OperationalizationClusterCredentials
type OperationalizationClusterProperties = original.OperationalizationClusterProperties
type OperationalizationClusterUpdateParameters = original.OperationalizationClusterUpdateParameters
type OperationalizationClustersClient = original.OperationalizationClustersClient
type OperationalizationClustersCreateOrUpdateFuture = original.OperationalizationClustersCreateOrUpdateFuture
type OperationalizationClustersDeleteFuture = original.OperationalizationClustersDeleteFuture
type OperationalizationClustersUpdateSystemServicesFuture = original.OperationalizationClustersUpdateSystemServicesFuture
type PaginatedOperationalizationClustersList = original.PaginatedOperationalizationClustersList
type PaginatedOperationalizationClustersListIterator = original.PaginatedOperationalizationClustersListIterator
type PaginatedOperationalizationClustersListPage = original.PaginatedOperationalizationClustersListPage
type Resource = original.Resource
type ResourceOperation = original.ResourceOperation
type ResourceOperationDisplay = original.ResourceOperationDisplay
type ServiceAuthConfiguration = original.ServiceAuthConfiguration
type ServicePrincipalProperties = original.ServicePrincipalProperties
type SslConfiguration = original.SslConfiguration
type StorageAccountCredentials = original.StorageAccountCredentials
type StorageAccountProperties = original.StorageAccountProperties
type SystemService = original.SystemService
type UpdateSystemServicesResponse = original.UpdateSystemServicesResponse

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewMachineLearningComputeClient(subscriptionID string) MachineLearningComputeClient {
	return original.NewMachineLearningComputeClient(subscriptionID)
}
func NewMachineLearningComputeClientWithBaseURI(baseURI string, subscriptionID string) MachineLearningComputeClient {
	return original.NewMachineLearningComputeClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationalizationClustersClient(subscriptionID string) OperationalizationClustersClient {
	return original.NewOperationalizationClustersClient(subscriptionID)
}
func NewOperationalizationClustersClientWithBaseURI(baseURI string, subscriptionID string) OperationalizationClustersClient {
	return original.NewOperationalizationClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewPaginatedOperationalizationClustersListIterator(page PaginatedOperationalizationClustersListPage) PaginatedOperationalizationClustersListIterator {
	return original.NewPaginatedOperationalizationClustersListIterator(page)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAgentVMSizeTypesValues() []AgentVMSizeTypes {
	return original.PossibleAgentVMSizeTypesValues()
}
func PossibleClusterTypeValues() []ClusterType {
	return original.PossibleClusterTypeValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func PossibleOrchestratorTypeValues() []OrchestratorType {
	return original.PossibleOrchestratorTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleSystemServiceTypeValues() []SystemServiceType {
	return original.PossibleSystemServiceTypeValues()
}
func PossibleUpdatesAvailableValues() []UpdatesAvailable {
	return original.PossibleUpdatesAvailableValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
