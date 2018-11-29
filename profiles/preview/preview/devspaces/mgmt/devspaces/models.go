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

package devspaces

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/devspaces/mgmt/2018-06-01-preview/devspaces"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type InstanceType = original.InstanceType

const (
	InstanceTypeKubernetes                            InstanceType = original.InstanceTypeKubernetes
	InstanceTypeOrchestratorSpecificConnectionDetails InstanceType = original.InstanceTypeOrchestratorSpecificConnectionDetails
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleted   ProvisioningState = original.Deleted
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type SkuTier = original.SkuTier

const (
	Standard SkuTier = original.Standard
)

type BaseClient = original.BaseClient
type BasicOrchestratorSpecificConnectionDetails = original.BasicOrchestratorSpecificConnectionDetails
type ContainerHostMapping = original.ContainerHostMapping
type ContainerHostMappingsClient = original.ContainerHostMappingsClient
type Controller = original.Controller
type ControllerConnectionDetails = original.ControllerConnectionDetails
type ControllerConnectionDetailsList = original.ControllerConnectionDetailsList
type ControllerList = original.ControllerList
type ControllerListIterator = original.ControllerListIterator
type ControllerListPage = original.ControllerListPage
type ControllerProperties = original.ControllerProperties
type ControllerUpdateParameters = original.ControllerUpdateParameters
type ControllersClient = original.ControllersClient
type ControllersCreateFuture = original.ControllersCreateFuture
type ControllersDeleteFuture = original.ControllersDeleteFuture
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type KubernetesConnectionDetails = original.KubernetesConnectionDetails
type OperationsClient = original.OperationsClient
type OrchestratorSpecificConnectionDetails = original.OrchestratorSpecificConnectionDetails
type Resource = original.Resource
type ResourceProviderOperationDefinition = original.ResourceProviderOperationDefinition
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type SetObject = original.SetObject
type Sku = original.Sku
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewContainerHostMappingsClient(subscriptionID string) ContainerHostMappingsClient {
	return original.NewContainerHostMappingsClient(subscriptionID)
}
func NewContainerHostMappingsClientWithBaseURI(baseURI string, subscriptionID string) ContainerHostMappingsClient {
	return original.NewContainerHostMappingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewControllerListIterator(page ControllerListPage) ControllerListIterator {
	return original.NewControllerListIterator(page)
}
func NewControllerListPage(getNextPage func(context.Context, ControllerList) (ControllerList, error)) ControllerListPage {
	return original.NewControllerListPage(getNextPage)
}
func NewControllersClient(subscriptionID string) ControllersClient {
	return original.NewControllersClient(subscriptionID)
}
func NewControllersClientWithBaseURI(baseURI string, subscriptionID string) ControllersClient {
	return original.NewControllersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return original.NewResourceProviderOperationListIterator(page)
}
func NewResourceProviderOperationListPage(getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return original.NewResourceProviderOperationListPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleInstanceTypeValues() []InstanceType {
	return original.PossibleInstanceTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
