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

package batchai

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/batchai/mgmt/2018-05-01/batchai"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AllocationState = original.AllocationState

const (
	Resizing AllocationState = original.Resizing
	Steady   AllocationState = original.Steady
)

type CachingType = original.CachingType

const (
	None      CachingType = original.None
	Readonly  CachingType = original.Readonly
	Readwrite CachingType = original.Readwrite
)

type DeallocationOption = original.DeallocationOption

const (
	Requeue              DeallocationOption = original.Requeue
	Terminate            DeallocationOption = original.Terminate
	Waitforjobcompletion DeallocationOption = original.Waitforjobcompletion
)

type ExecutionState = original.ExecutionState

const (
	Failed      ExecutionState = original.Failed
	Queued      ExecutionState = original.Queued
	Running     ExecutionState = original.Running
	Succeeded   ExecutionState = original.Succeeded
	Terminating ExecutionState = original.Terminating
)

type FileServerProvisioningState = original.FileServerProvisioningState

const (
	FileServerProvisioningStateCreating  FileServerProvisioningState = original.FileServerProvisioningStateCreating
	FileServerProvisioningStateDeleting  FileServerProvisioningState = original.FileServerProvisioningStateDeleting
	FileServerProvisioningStateFailed    FileServerProvisioningState = original.FileServerProvisioningStateFailed
	FileServerProvisioningStateSucceeded FileServerProvisioningState = original.FileServerProvisioningStateSucceeded
	FileServerProvisioningStateUpdating  FileServerProvisioningState = original.FileServerProvisioningStateUpdating
)

type FileType = original.FileType

const (
	FileTypeDirectory FileType = original.FileTypeDirectory
	FileTypeFile      FileType = original.FileTypeFile
)

type JobPriority = original.JobPriority

const (
	High   JobPriority = original.High
	Low    JobPriority = original.Low
	Normal JobPriority = original.Normal
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type StorageAccountType = original.StorageAccountType

const (
	PremiumLRS  StorageAccountType = original.PremiumLRS
	StandardLRS StorageAccountType = original.StandardLRS
)

type ToolType = original.ToolType

const (
	Caffe      ToolType = original.Caffe
	Caffe2     ToolType = original.Caffe2
	Chainer    ToolType = original.Chainer
	Cntk       ToolType = original.Cntk
	Custom     ToolType = original.Custom
	Custommpi  ToolType = original.Custommpi
	Horovod    ToolType = original.Horovod
	Tensorflow ToolType = original.Tensorflow
)

type UsageUnit = original.UsageUnit

const (
	Count UsageUnit = original.Count
)

type VMPriority = original.VMPriority

const (
	Dedicated   VMPriority = original.Dedicated
	Lowpriority VMPriority = original.Lowpriority
)

type AppInsightsReference = original.AppInsightsReference
type AutoScaleSettings = original.AutoScaleSettings
type AzureBlobFileSystemReference = original.AzureBlobFileSystemReference
type AzureFileShareReference = original.AzureFileShareReference
type AzureStorageCredentialsInfo = original.AzureStorageCredentialsInfo
type BaseClient = original.BaseClient
type CNTKsettings = original.CNTKsettings
type Caffe2Settings = original.Caffe2Settings
type CaffeSettings = original.CaffeSettings
type ChainerSettings = original.ChainerSettings
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Cluster = original.Cluster
type ClusterBaseProperties = original.ClusterBaseProperties
type ClusterCreateParameters = original.ClusterCreateParameters
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterProperties = original.ClusterProperties
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpdateProperties = original.ClusterUpdateProperties
type ClustersClient = original.ClustersClient
type ClustersCreateFuture = original.ClustersCreateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ContainerSettings = original.ContainerSettings
type CustomMpiSettings = original.CustomMpiSettings
type CustomToolkitSettings = original.CustomToolkitSettings
type DataDisks = original.DataDisks
type EnvironmentVariable = original.EnvironmentVariable
type EnvironmentVariableWithSecretValue = original.EnvironmentVariableWithSecretValue
type Error = original.Error
type Experiment = original.Experiment
type ExperimentListResult = original.ExperimentListResult
type ExperimentListResultIterator = original.ExperimentListResultIterator
type ExperimentListResultPage = original.ExperimentListResultPage
type ExperimentProperties = original.ExperimentProperties
type ExperimentsClient = original.ExperimentsClient
type ExperimentsCreateFuture = original.ExperimentsCreateFuture
type ExperimentsDeleteFuture = original.ExperimentsDeleteFuture
type File = original.File
type FileListResult = original.FileListResult
type FileListResultIterator = original.FileListResultIterator
type FileListResultPage = original.FileListResultPage
type FileProperties = original.FileProperties
type FileServer = original.FileServer
type FileServerBaseProperties = original.FileServerBaseProperties
type FileServerCreateParameters = original.FileServerCreateParameters
type FileServerListResult = original.FileServerListResult
type FileServerListResultIterator = original.FileServerListResultIterator
type FileServerListResultPage = original.FileServerListResultPage
type FileServerProperties = original.FileServerProperties
type FileServerReference = original.FileServerReference
type FileServersClient = original.FileServersClient
type FileServersCreateFuture = original.FileServersCreateFuture
type FileServersDeleteFuture = original.FileServersDeleteFuture
type HorovodSettings = original.HorovodSettings
type ImageReference = original.ImageReference
type ImageSourceRegistry = original.ImageSourceRegistry
type InputDirectory = original.InputDirectory
type Job = original.Job
type JobBaseProperties = original.JobBaseProperties
type JobBasePropertiesConstraints = original.JobBasePropertiesConstraints
type JobCreateParameters = original.JobCreateParameters
type JobListResult = original.JobListResult
type JobListResultIterator = original.JobListResultIterator
type JobListResultPage = original.JobListResultPage
type JobPreparation = original.JobPreparation
type JobProperties = original.JobProperties
type JobPropertiesConstraints = original.JobPropertiesConstraints
type JobPropertiesExecutionInfo = original.JobPropertiesExecutionInfo
type JobsClient = original.JobsClient
type JobsCreateFuture = original.JobsCreateFuture
type JobsDeleteFuture = original.JobsDeleteFuture
type JobsTerminateFuture = original.JobsTerminateFuture
type KeyVaultSecretReference = original.KeyVaultSecretReference
type ListUsagesResult = original.ListUsagesResult
type ListUsagesResultIterator = original.ListUsagesResultIterator
type ListUsagesResultPage = original.ListUsagesResultPage
type ManualScaleSettings = original.ManualScaleSettings
type MountSettings = original.MountSettings
type MountVolumes = original.MountVolumes
type NameValuePair = original.NameValuePair
type NodeSetup = original.NodeSetup
type NodeStateCounts = original.NodeStateCounts
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type OutputDirectory = original.OutputDirectory
type PerformanceCountersSettings = original.PerformanceCountersSettings
type PrivateRegistryCredentials = original.PrivateRegistryCredentials
type ProxyResource = original.ProxyResource
type PyTorchSettings = original.PyTorchSettings
type RemoteLoginInformation = original.RemoteLoginInformation
type RemoteLoginInformationListResult = original.RemoteLoginInformationListResult
type RemoteLoginInformationListResultIterator = original.RemoteLoginInformationListResultIterator
type RemoteLoginInformationListResultPage = original.RemoteLoginInformationListResultPage
type Resource = original.Resource
type ResourceID = original.ResourceID
type SSHConfiguration = original.SSHConfiguration
type ScaleSettings = original.ScaleSettings
type SetupTask = original.SetupTask
type TensorFlowSettings = original.TensorFlowSettings
type UnmanagedFileSystemReference = original.UnmanagedFileSystemReference
type Usage = original.Usage
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type UserAccountSettings = original.UserAccountSettings
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type Workspace = original.Workspace
type WorkspaceCreateParameters = original.WorkspaceCreateParameters
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListResultIterator = original.WorkspaceListResultIterator
type WorkspaceListResultPage = original.WorkspaceListResultPage
type WorkspaceProperties = original.WorkspaceProperties
type WorkspaceUpdateParameters = original.WorkspaceUpdateParameters
type WorkspacesClient = original.WorkspacesClient
type WorkspacesCreateFuture = original.WorkspacesCreateFuture
type WorkspacesDeleteFuture = original.WorkspacesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClusterListResultIterator(page ClusterListResultPage) ClusterListResultIterator {
	return original.NewClusterListResultIterator(page)
}
func NewClusterListResultPage(cur ClusterListResult, getNextPage func(context.Context, ClusterListResult) (ClusterListResult, error)) ClusterListResultPage {
	return original.NewClusterListResultPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewExperimentListResultIterator(page ExperimentListResultPage) ExperimentListResultIterator {
	return original.NewExperimentListResultIterator(page)
}
func NewExperimentListResultPage(cur ExperimentListResult, getNextPage func(context.Context, ExperimentListResult) (ExperimentListResult, error)) ExperimentListResultPage {
	return original.NewExperimentListResultPage(cur, getNextPage)
}
func NewExperimentsClient(subscriptionID string) ExperimentsClient {
	return original.NewExperimentsClient(subscriptionID)
}
func NewExperimentsClientWithBaseURI(baseURI string, subscriptionID string) ExperimentsClient {
	return original.NewExperimentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFileListResultIterator(page FileListResultPage) FileListResultIterator {
	return original.NewFileListResultIterator(page)
}
func NewFileListResultPage(cur FileListResult, getNextPage func(context.Context, FileListResult) (FileListResult, error)) FileListResultPage {
	return original.NewFileListResultPage(cur, getNextPage)
}
func NewFileServerListResultIterator(page FileServerListResultPage) FileServerListResultIterator {
	return original.NewFileServerListResultIterator(page)
}
func NewFileServerListResultPage(cur FileServerListResult, getNextPage func(context.Context, FileServerListResult) (FileServerListResult, error)) FileServerListResultPage {
	return original.NewFileServerListResultPage(cur, getNextPage)
}
func NewFileServersClient(subscriptionID string) FileServersClient {
	return original.NewFileServersClient(subscriptionID)
}
func NewFileServersClientWithBaseURI(baseURI string, subscriptionID string) FileServersClient {
	return original.NewFileServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobListResultIterator(page JobListResultPage) JobListResultIterator {
	return original.NewJobListResultIterator(page)
}
func NewJobListResultPage(cur JobListResult, getNextPage func(context.Context, JobListResult) (JobListResult, error)) JobListResultPage {
	return original.NewJobListResultPage(cur, getNextPage)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewListUsagesResultIterator(page ListUsagesResultPage) ListUsagesResultIterator {
	return original.NewListUsagesResultIterator(page)
}
func NewListUsagesResultPage(cur ListUsagesResult, getNextPage func(context.Context, ListUsagesResult) (ListUsagesResult, error)) ListUsagesResultPage {
	return original.NewListUsagesResultPage(cur, getNextPage)
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
func NewRemoteLoginInformationListResultIterator(page RemoteLoginInformationListResultPage) RemoteLoginInformationListResultIterator {
	return original.NewRemoteLoginInformationListResultIterator(page)
}
func NewRemoteLoginInformationListResultPage(cur RemoteLoginInformationListResult, getNextPage func(context.Context, RemoteLoginInformationListResult) (RemoteLoginInformationListResult, error)) RemoteLoginInformationListResultPage {
	return original.NewRemoteLoginInformationListResultPage(cur, getNextPage)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceListResultIterator(page WorkspaceListResultPage) WorkspaceListResultIterator {
	return original.NewWorkspaceListResultIterator(page)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleCachingTypeValues() []CachingType {
	return original.PossibleCachingTypeValues()
}
func PossibleDeallocationOptionValues() []DeallocationOption {
	return original.PossibleDeallocationOptionValues()
}
func PossibleExecutionStateValues() []ExecutionState {
	return original.PossibleExecutionStateValues()
}
func PossibleFileServerProvisioningStateValues() []FileServerProvisioningState {
	return original.PossibleFileServerProvisioningStateValues()
}
func PossibleFileTypeValues() []FileType {
	return original.PossibleFileTypeValues()
}
func PossibleJobPriorityValues() []JobPriority {
	return original.PossibleJobPriorityValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return original.PossibleStorageAccountTypeValues()
}
func PossibleToolTypeValues() []ToolType {
	return original.PossibleToolTypeValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func PossibleVMPriorityValues() []VMPriority {
	return original.PossibleVMPriorityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
