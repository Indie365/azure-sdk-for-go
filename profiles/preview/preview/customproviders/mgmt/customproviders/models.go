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

package customproviders

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/customproviders/mgmt/2018-09-01-preview/customproviders"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionRouting = original.ActionRouting

const (
	Proxy ActionRouting = original.Proxy
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
)

type ResourceTypeRouting = original.ResourceTypeRouting

const (
	ResourceTypeRoutingProxy      ResourceTypeRouting = original.ResourceTypeRoutingProxy
	ResourceTypeRoutingProxyCache ResourceTypeRouting = original.ResourceTypeRoutingProxyCache
)

type ValidationType = original.ValidationType

const (
	Swagger ValidationType = original.Swagger
)

type Association = original.Association
type AssociationClient = original.AssociationClient
type AssociationCreateOrUpdateFuture = original.AssociationCreateOrUpdateFuture
type AssociationDeleteFuture = original.AssociationDeleteFuture
type AssociationProperties = original.AssociationProperties
type AssociationsList = original.AssociationsList
type AssociationsListIterator = original.AssociationsListIterator
type AssociationsListPage = original.AssociationsListPage
type BaseClient = original.BaseClient
type CustomRPActionRouteDefinition = original.CustomRPActionRouteDefinition
type CustomRPManifest = original.CustomRPManifest
type CustomRPManifestProperties = original.CustomRPManifestProperties
type CustomRPResourceTypeRouteDefinition = original.CustomRPResourceTypeRouteDefinition
type CustomRPRouteDefinition = original.CustomRPRouteDefinition
type CustomRPValidations = original.CustomRPValidations
type CustomResourceProviderClient = original.CustomResourceProviderClient
type CustomResourceProviderCreateOrUpdateFuture = original.CustomResourceProviderCreateOrUpdateFuture
type CustomResourceProviderDeleteFuture = original.CustomResourceProviderDeleteFuture
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type ListByCustomRPManifest = original.ListByCustomRPManifest
type ListByCustomRPManifestIterator = original.ListByCustomRPManifestIterator
type ListByCustomRPManifestPage = original.ListByCustomRPManifestPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type ResourceProviderOperation = original.ResourceProviderOperation
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceProviderOperationListIterator = original.ResourceProviderOperationListIterator
type ResourceProviderOperationListPage = original.ResourceProviderOperationListPage
type ResourceProvidersUpdate = original.ResourceProvidersUpdate

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAssociationClient(subscriptionID string) AssociationClient {
	return original.NewAssociationClient(subscriptionID)
}
func NewAssociationClientWithBaseURI(baseURI string, subscriptionID string) AssociationClient {
	return original.NewAssociationClientWithBaseURI(baseURI, subscriptionID)
}
func NewAssociationsListIterator(page AssociationsListPage) AssociationsListIterator {
	return original.NewAssociationsListIterator(page)
}
func NewAssociationsListPage(getNextPage func(context.Context, AssociationsList) (AssociationsList, error)) AssociationsListPage {
	return original.NewAssociationsListPage(getNextPage)
}
func NewCustomResourceProviderClient(subscriptionID string) CustomResourceProviderClient {
	return original.NewCustomResourceProviderClient(subscriptionID)
}
func NewCustomResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) CustomResourceProviderClient {
	return original.NewCustomResourceProviderClientWithBaseURI(baseURI, subscriptionID)
}
func NewListByCustomRPManifestIterator(page ListByCustomRPManifestPage) ListByCustomRPManifestIterator {
	return original.NewListByCustomRPManifestIterator(page)
}
func NewListByCustomRPManifestPage(getNextPage func(context.Context, ListByCustomRPManifest) (ListByCustomRPManifest, error)) ListByCustomRPManifestPage {
	return original.NewListByCustomRPManifestPage(getNextPage)
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
func PossibleActionRoutingValues() []ActionRouting {
	return original.PossibleActionRoutingValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceTypeRoutingValues() []ResourceTypeRouting {
	return original.PossibleResourceTypeRoutingValues()
}
func PossibleValidationTypeValues() []ValidationType {
	return original.PossibleValidationTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
