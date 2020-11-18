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

package containerinstance

import original "github.com/Azure/azure-sdk-for-go/services/preview/containerinstance/mgmt/2018-02-01-preview/containerinstance"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ContainerGroupNetworkProtocol = original.ContainerGroupNetworkProtocol

const (
	TCP ContainerGroupNetworkProtocol = original.TCP
	UDP ContainerGroupNetworkProtocol = original.UDP
)

type ContainerGroupRestartPolicy = original.ContainerGroupRestartPolicy

const (
	Always    ContainerGroupRestartPolicy = original.Always
	Never     ContainerGroupRestartPolicy = original.Never
	OnFailure ContainerGroupRestartPolicy = original.OnFailure
)

type ContainerNetworkProtocol = original.ContainerNetworkProtocol

const (
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = original.ContainerNetworkProtocolTCP
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = original.ContainerNetworkProtocolUDP
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	Linux   OperatingSystemTypes = original.Linux
	Windows OperatingSystemTypes = original.Windows
)

type OperationsOrigin = original.OperationsOrigin

const (
	System OperationsOrigin = original.System
	User   OperationsOrigin = original.User
)

type AzureFileVolume = original.AzureFileVolume
type BaseClient = original.BaseClient
type Container = original.Container
type ContainerExecRequest = original.ContainerExecRequest
type ContainerExecRequestTerminalSize = original.ContainerExecRequestTerminalSize
type ContainerExecResponse = original.ContainerExecResponse
type ContainerGroup = original.ContainerGroup
type ContainerGroupListResult = original.ContainerGroupListResult
type ContainerGroupListResultIterator = original.ContainerGroupListResultIterator
type ContainerGroupListResultPage = original.ContainerGroupListResultPage
type ContainerGroupProperties = original.ContainerGroupProperties
type ContainerGroupPropertiesInstanceView = original.ContainerGroupPropertiesInstanceView
type ContainerGroupUsageClient = original.ContainerGroupUsageClient
type ContainerGroupsClient = original.ContainerGroupsClient
type ContainerGroupsCreateOrUpdateFuture = original.ContainerGroupsCreateOrUpdateFuture
type ContainerLogsClient = original.ContainerLogsClient
type ContainerPort = original.ContainerPort
type ContainerProperties = original.ContainerProperties
type ContainerPropertiesInstanceView = original.ContainerPropertiesInstanceView
type ContainerState = original.ContainerState
type EnvironmentVariable = original.EnvironmentVariable
type Event = original.Event
type GitRepoVolume = original.GitRepoVolume
type IPAddress = original.IPAddress
type ImageRegistryCredential = original.ImageRegistryCredential
type Logs = original.Logs
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Port = original.Port
type Resource = original.Resource
type ResourceLimits = original.ResourceLimits
type ResourceRequests = original.ResourceRequests
type ResourceRequirements = original.ResourceRequirements
type StartContainerClient = original.StartContainerClient
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type Volume = original.Volume
type VolumeMount = original.VolumeMount

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewContainerGroupListResultIterator(page ContainerGroupListResultPage) ContainerGroupListResultIterator {
	return original.NewContainerGroupListResultIterator(page)
}
func NewContainerGroupUsageClient(subscriptionID string) ContainerGroupUsageClient {
	return original.NewContainerGroupUsageClient(subscriptionID)
}
func NewContainerGroupUsageClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupUsageClient {
	return original.NewContainerGroupUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewContainerGroupsClient(subscriptionID string) ContainerGroupsClient {
	return original.NewContainerGroupsClient(subscriptionID)
}
func NewContainerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ContainerGroupsClient {
	return original.NewContainerGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewContainerLogsClient(subscriptionID string) ContainerLogsClient {
	return original.NewContainerLogsClient(subscriptionID)
}
func NewContainerLogsClientWithBaseURI(baseURI string, subscriptionID string) ContainerLogsClient {
	return original.NewContainerLogsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStartContainerClient(subscriptionID string) StartContainerClient {
	return original.NewStartContainerClient(subscriptionID)
}
func NewStartContainerClientWithBaseURI(baseURI string, subscriptionID string) StartContainerClient {
	return original.NewStartContainerClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return original.PossibleContainerGroupNetworkProtocolValues()
}
func PossibleContainerGroupRestartPolicyValues() []ContainerGroupRestartPolicy {
	return original.PossibleContainerGroupRestartPolicyValues()
}
func PossibleContainerNetworkProtocolValues() []ContainerNetworkProtocol {
	return original.PossibleContainerNetworkProtocolValues()
}
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return original.PossibleOperatingSystemTypesValues()
}
func PossibleOperationsOriginValues() []OperationsOrigin {
	return original.PossibleOperationsOriginValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
