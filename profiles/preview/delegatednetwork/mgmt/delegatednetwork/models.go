//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package delegatednetwork

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/delegatednetwork/mgmt/2021-03-15/delegatednetwork"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	ActionTypeInternal ActionType = original.ActionTypeInternal
)

type ControllerState = original.ControllerState

const (
	ControllerStateDeleting     ControllerState = original.ControllerStateDeleting
	ControllerStateFailed       ControllerState = original.ControllerStateFailed
	ControllerStateProvisioning ControllerState = original.ControllerStateProvisioning
	ControllerStateSucceeded    ControllerState = original.ControllerStateSucceeded
)

type DelegatedSubnetState = original.DelegatedSubnetState

const (
	DelegatedSubnetStateDeleting     DelegatedSubnetState = original.DelegatedSubnetStateDeleting
	DelegatedSubnetStateFailed       DelegatedSubnetState = original.DelegatedSubnetStateFailed
	DelegatedSubnetStateProvisioning DelegatedSubnetState = original.DelegatedSubnetStateProvisioning
	DelegatedSubnetStateSucceeded    DelegatedSubnetState = original.DelegatedSubnetStateSucceeded
)

type OrchestratorInstanceState = original.OrchestratorInstanceState

const (
	OrchestratorInstanceStateDeleting     OrchestratorInstanceState = original.OrchestratorInstanceStateDeleting
	OrchestratorInstanceStateFailed       OrchestratorInstanceState = original.OrchestratorInstanceStateFailed
	OrchestratorInstanceStateProvisioning OrchestratorInstanceState = original.OrchestratorInstanceStateProvisioning
	OrchestratorInstanceStateSucceeded    OrchestratorInstanceState = original.OrchestratorInstanceStateSucceeded
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type BaseClient = original.BaseClient
type Client = original.Client
type ControllerClient = original.ControllerClient
type ControllerCreateFuture = original.ControllerCreateFuture
type ControllerDeleteFuture = original.ControllerDeleteFuture
type ControllerDetails = original.ControllerDetails
type ControllerResource = original.ControllerResource
type ControllerResourceUpdateParameters = original.ControllerResourceUpdateParameters
type DelegatedController = original.DelegatedController
type DelegatedControllerProperties = original.DelegatedControllerProperties
type DelegatedControllers = original.DelegatedControllers
type DelegatedControllersIterator = original.DelegatedControllersIterator
type DelegatedControllersPage = original.DelegatedControllersPage
type DelegatedSubnet = original.DelegatedSubnet
type DelegatedSubnetProperties = original.DelegatedSubnetProperties
type DelegatedSubnetResource = original.DelegatedSubnetResource
type DelegatedSubnetServiceClient = original.DelegatedSubnetServiceClient
type DelegatedSubnetServiceDeleteDetailsFuture = original.DelegatedSubnetServiceDeleteDetailsFuture
type DelegatedSubnetServicePatchDetailsFuture = original.DelegatedSubnetServicePatchDetailsFuture
type DelegatedSubnetServicePutDetailsFuture = original.DelegatedSubnetServicePutDetailsFuture
type DelegatedSubnets = original.DelegatedSubnets
type DelegatedSubnetsIterator = original.DelegatedSubnetsIterator
type DelegatedSubnetsPage = original.DelegatedSubnetsPage
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Orchestrator = original.Orchestrator
type OrchestratorIdentity = original.OrchestratorIdentity
type OrchestratorInstanceServiceClient = original.OrchestratorInstanceServiceClient
type OrchestratorInstanceServiceCreateFuture = original.OrchestratorInstanceServiceCreateFuture
type OrchestratorInstanceServiceDeleteFuture = original.OrchestratorInstanceServiceDeleteFuture
type OrchestratorResource = original.OrchestratorResource
type OrchestratorResourceProperties = original.OrchestratorResourceProperties
type OrchestratorResourceUpdateParameters = original.OrchestratorResourceUpdateParameters
type Orchestrators = original.Orchestrators
type OrchestratorsIterator = original.OrchestratorsIterator
type OrchestratorsPage = original.OrchestratorsPage
type ResourceUpdateParameters = original.ResourceUpdateParameters
type SubnetDetails = original.SubnetDetails

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewControllerClient(subscriptionID string) ControllerClient {
	return original.NewControllerClient(subscriptionID)
}
func NewControllerClientWithBaseURI(baseURI string, subscriptionID string) ControllerClient {
	return original.NewControllerClientWithBaseURI(baseURI, subscriptionID)
}
func NewDelegatedControllersIterator(page DelegatedControllersPage) DelegatedControllersIterator {
	return original.NewDelegatedControllersIterator(page)
}
func NewDelegatedControllersPage(cur DelegatedControllers, getNextPage func(context.Context, DelegatedControllers) (DelegatedControllers, error)) DelegatedControllersPage {
	return original.NewDelegatedControllersPage(cur, getNextPage)
}
func NewDelegatedSubnetServiceClient(subscriptionID string) DelegatedSubnetServiceClient {
	return original.NewDelegatedSubnetServiceClient(subscriptionID)
}
func NewDelegatedSubnetServiceClientWithBaseURI(baseURI string, subscriptionID string) DelegatedSubnetServiceClient {
	return original.NewDelegatedSubnetServiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewDelegatedSubnetsIterator(page DelegatedSubnetsPage) DelegatedSubnetsIterator {
	return original.NewDelegatedSubnetsIterator(page)
}
func NewDelegatedSubnetsPage(cur DelegatedSubnets, getNextPage func(context.Context, DelegatedSubnets) (DelegatedSubnets, error)) DelegatedSubnetsPage {
	return original.NewDelegatedSubnetsPage(cur, getNextPage)
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
func NewOrchestratorInstanceServiceClient(subscriptionID string) OrchestratorInstanceServiceClient {
	return original.NewOrchestratorInstanceServiceClient(subscriptionID)
}
func NewOrchestratorInstanceServiceClientWithBaseURI(baseURI string, subscriptionID string) OrchestratorInstanceServiceClient {
	return original.NewOrchestratorInstanceServiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewOrchestratorsIterator(page OrchestratorsPage) OrchestratorsIterator {
	return original.NewOrchestratorsIterator(page)
}
func NewOrchestratorsPage(cur Orchestrators, getNextPage func(context.Context, Orchestrators) (Orchestrators, error)) OrchestratorsPage {
	return original.NewOrchestratorsPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleControllerStateValues() []ControllerState {
	return original.PossibleControllerStateValues()
}
func PossibleDelegatedSubnetStateValues() []DelegatedSubnetState {
	return original.PossibleDelegatedSubnetStateValues()
}
func PossibleOrchestratorInstanceStateValues() []OrchestratorInstanceState {
	return original.PossibleOrchestratorInstanceStateValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
