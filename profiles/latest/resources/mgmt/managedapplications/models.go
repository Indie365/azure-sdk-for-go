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

package managedapplications

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ApplicationArtifactType = original.ApplicationArtifactType

const (
	Custom   ApplicationArtifactType = original.Custom
	Template ApplicationArtifactType = original.Template
)

type ApplicationLockLevel = original.ApplicationLockLevel

const (
	CanNotDelete ApplicationLockLevel = original.CanNotDelete
	None         ApplicationLockLevel = original.None
	ReadOnly     ApplicationLockLevel = original.ReadOnly
)

type JITStatusUpdate = original.JITStatusUpdate

const (
	JITStatusUpdateElevate      JITStatusUpdate = original.JITStatusUpdateElevate
	JITStatusUpdateNotSpecified JITStatusUpdate = original.JITStatusUpdateNotSpecified
	JITStatusUpdateRemove       JITStatusUpdate = original.JITStatusUpdateRemove
)

type JITSubStatus = original.JITSubStatus

const (
	JITSubStatusApproved     JITSubStatus = original.JITSubStatusApproved
	JITSubStatusDenied       JITSubStatus = original.JITSubStatusDenied
	JITSubStatusExpired      JITSubStatus = original.JITSubStatusExpired
	JITSubStatusFailed       JITSubStatus = original.JITSubStatusFailed
	JITSubStatusNotSpecified JITSubStatus = original.JITSubStatusNotSpecified
	JITSubStatusTimeout      JITSubStatus = original.JITSubStatusTimeout
)

type JitRequestState = original.JitRequestState

const (
	Approved     JitRequestState = original.Approved
	Canceled     JitRequestState = original.Canceled
	Denied       JitRequestState = original.Denied
	Expired      JitRequestState = original.Expired
	Failed       JitRequestState = original.Failed
	NotSpecified JitRequestState = original.NotSpecified
	Pending      JitRequestState = original.Pending
	Timeout      JitRequestState = original.Timeout
)

type JitSchedulingType = original.JitSchedulingType

const (
	JitSchedulingTypeNotSpecified JitSchedulingType = original.JitSchedulingTypeNotSpecified
	JitSchedulingTypeOnce         JitSchedulingType = original.JitSchedulingTypeOnce
	JitSchedulingTypeRecurring    JitSchedulingType = original.JitSchedulingTypeRecurring
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted  ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreated   ProvisioningState = original.ProvisioningStateCreated
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted   ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateReady     ProvisioningState = original.ProvisioningStateReady
	ProvisioningStateRunning   ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type Application = original.Application
type ApplicationArtifact = original.ApplicationArtifact
type ApplicationClientDetails = original.ApplicationClientDetails
type ApplicationDefinition = original.ApplicationDefinition
type ApplicationDefinitionArtifactsClient = original.ApplicationDefinitionArtifactsClient
type ApplicationDefinitionListResult = original.ApplicationDefinitionListResult
type ApplicationDefinitionListResultIterator = original.ApplicationDefinitionListResultIterator
type ApplicationDefinitionListResultPage = original.ApplicationDefinitionListResultPage
type ApplicationDefinitionProperties = original.ApplicationDefinitionProperties
type ApplicationDefinitionsClient = original.ApplicationDefinitionsClient
type ApplicationDefinitionsCreateOrUpdateByIDFuture = original.ApplicationDefinitionsCreateOrUpdateByIDFuture
type ApplicationDefinitionsCreateOrUpdateFuture = original.ApplicationDefinitionsCreateOrUpdateFuture
type ApplicationDefinitionsDeleteByIDFuture = original.ApplicationDefinitionsDeleteByIDFuture
type ApplicationDefinitionsDeleteFuture = original.ApplicationDefinitionsDeleteFuture
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationPatchable = original.ApplicationPatchable
type ApplicationProperties = original.ApplicationProperties
type ApplicationPropertiesPatchable = original.ApplicationPropertiesPatchable
type ApplicationProviderAuthorization = original.ApplicationProviderAuthorization
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateOrUpdateByIDFuture = original.ApplicationsCreateOrUpdateByIDFuture
type ApplicationsCreateOrUpdateFuture = original.ApplicationsCreateOrUpdateFuture
type ApplicationsDeleteByIDFuture = original.ApplicationsDeleteByIDFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type BaseClient = original.BaseClient
type ErrorResponse = original.ErrorResponse
type GenericResource = original.GenericResource
type Identity = original.Identity
type JitAuthorizationPolicies = original.JitAuthorizationPolicies
type JitRequestDefinition = original.JitRequestDefinition
type JitRequestDefinitionListResult = original.JitRequestDefinitionListResult
type JitRequestMetadata = original.JitRequestMetadata
type JitRequestProperties = original.JitRequestProperties
type JitRequestsClient = original.JitRequestsClient
type JitRequestsCreateOrUpdateFuture = original.JitRequestsCreateOrUpdateFuture
type JitRequestsPatchFuture = original.JitRequestsPatchFuture
type JitSchedulingPolicy = original.JitSchedulingPolicy
type JitUpdateAccessDefinition = original.JitUpdateAccessDefinition
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type Plan = original.Plan
type PlanPatchable = original.PlanPatchable
type RefreshPermissionsClient = original.RefreshPermissionsClient
type RefreshPermissionsPostFuture = original.RefreshPermissionsPostFuture
type Resource = original.Resource
type Sku = original.Sku
type String = original.String
type UpdateAccessClient = original.UpdateAccessClient
type UpdateAccessPostFuture = original.UpdateAccessPostFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationDefinitionArtifactsClient(subscriptionID string) ApplicationDefinitionArtifactsClient {
	return original.NewApplicationDefinitionArtifactsClient(subscriptionID)
}
func NewApplicationDefinitionArtifactsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationDefinitionArtifactsClient {
	return original.NewApplicationDefinitionArtifactsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationDefinitionListResultIterator(page ApplicationDefinitionListResultPage) ApplicationDefinitionListResultIterator {
	return original.NewApplicationDefinitionListResultIterator(page)
}
func NewApplicationDefinitionListResultPage(getNextPage func(context.Context, ApplicationDefinitionListResult) (ApplicationDefinitionListResult, error)) ApplicationDefinitionListResultPage {
	return original.NewApplicationDefinitionListResultPage(getNextPage)
}
func NewApplicationDefinitionsClient(subscriptionID string) ApplicationDefinitionsClient {
	return original.NewApplicationDefinitionsClient(subscriptionID)
}
func NewApplicationDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationDefinitionsClient {
	return original.NewApplicationDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationListResultIterator(page ApplicationListResultPage) ApplicationListResultIterator {
	return original.NewApplicationListResultIterator(page)
}
func NewApplicationListResultPage(getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(getNextPage)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewJitRequestsClient(subscriptionID string) JitRequestsClient {
	return original.NewJitRequestsClient(subscriptionID)
}
func NewJitRequestsClientWithBaseURI(baseURI string, subscriptionID string) JitRequestsClient {
	return original.NewJitRequestsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewRefreshPermissionsClient(subscriptionID string) RefreshPermissionsClient {
	return original.NewRefreshPermissionsClient(subscriptionID)
}
func NewRefreshPermissionsClientWithBaseURI(baseURI string, subscriptionID string) RefreshPermissionsClient {
	return original.NewRefreshPermissionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUpdateAccessClient(subscriptionID string) UpdateAccessClient {
	return original.NewUpdateAccessClient(subscriptionID)
}
func NewUpdateAccessClientWithBaseURI(baseURI string, subscriptionID string) UpdateAccessClient {
	return original.NewUpdateAccessClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return original.PossibleApplicationArtifactTypeValues()
}
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return original.PossibleApplicationLockLevelValues()
}
func PossibleJITStatusUpdateValues() []JITStatusUpdate {
	return original.PossibleJITStatusUpdateValues()
}
func PossibleJITSubStatusValues() []JITSubStatus {
	return original.PossibleJITSubStatusValues()
}
func PossibleJitRequestStateValues() []JitRequestState {
	return original.PossibleJitRequestStateValues()
}
func PossibleJitSchedulingTypeValues() []JitSchedulingType {
	return original.PossibleJitSchedulingTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
