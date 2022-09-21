//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package managedapplications

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/solutions/mgmt/2021-07-01/managedapplications"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	Internal ActionType = original.Internal
)

type ApplicationArtifactName = original.ApplicationArtifactName

const (
	Authorizations       ApplicationArtifactName = original.Authorizations
	CustomRoleDefinition ApplicationArtifactName = original.CustomRoleDefinition
	NotSpecified         ApplicationArtifactName = original.NotSpecified
	ViewDefinition       ApplicationArtifactName = original.ViewDefinition
)

type ApplicationArtifactType = original.ApplicationArtifactType

const (
	ApplicationArtifactTypeCustom       ApplicationArtifactType = original.ApplicationArtifactTypeCustom
	ApplicationArtifactTypeNotSpecified ApplicationArtifactType = original.ApplicationArtifactTypeNotSpecified
	ApplicationArtifactTypeTemplate     ApplicationArtifactType = original.ApplicationArtifactTypeTemplate
)

type ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactName

const (
	ApplicationDefinitionArtifactNameApplicationResourceTemplate ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameApplicationResourceTemplate
	ApplicationDefinitionArtifactNameCreateUIDefinition          ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameCreateUIDefinition
	ApplicationDefinitionArtifactNameMainTemplateParameters      ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameMainTemplateParameters
	ApplicationDefinitionArtifactNameNotSpecified                ApplicationDefinitionArtifactName = original.ApplicationDefinitionArtifactNameNotSpecified
)

type ApplicationLockLevel = original.ApplicationLockLevel

const (
	CanNotDelete ApplicationLockLevel = original.CanNotDelete
	None         ApplicationLockLevel = original.None
	ReadOnly     ApplicationLockLevel = original.ReadOnly
)

type ApplicationManagementMode = original.ApplicationManagementMode

const (
	ApplicationManagementModeManaged      ApplicationManagementMode = original.ApplicationManagementModeManaged
	ApplicationManagementModeNotSpecified ApplicationManagementMode = original.ApplicationManagementModeNotSpecified
	ApplicationManagementModeUnmanaged    ApplicationManagementMode = original.ApplicationManagementModeUnmanaged
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type DeploymentMode = original.DeploymentMode

const (
	DeploymentModeComplete     DeploymentMode = original.DeploymentModeComplete
	DeploymentModeIncremental  DeploymentMode = original.DeploymentModeIncremental
	DeploymentModeNotSpecified DeploymentMode = original.DeploymentModeNotSpecified
)

type JitApprovalMode = original.JitApprovalMode

const (
	JitApprovalModeAutoApprove   JitApprovalMode = original.JitApprovalModeAutoApprove
	JitApprovalModeManualApprove JitApprovalMode = original.JitApprovalModeManualApprove
	JitApprovalModeNotSpecified  JitApprovalMode = original.JitApprovalModeNotSpecified
)

type JitApproverType = original.JitApproverType

const (
	Group JitApproverType = original.Group
	User  JitApproverType = original.User
)

type JitRequestState = original.JitRequestState

const (
	JitRequestStateApproved     JitRequestState = original.JitRequestStateApproved
	JitRequestStateCanceled     JitRequestState = original.JitRequestStateCanceled
	JitRequestStateDenied       JitRequestState = original.JitRequestStateDenied
	JitRequestStateExpired      JitRequestState = original.JitRequestStateExpired
	JitRequestStateFailed       JitRequestState = original.JitRequestStateFailed
	JitRequestStateNotSpecified JitRequestState = original.JitRequestStateNotSpecified
	JitRequestStatePending      JitRequestState = original.JitRequestStatePending
	JitRequestStateTimeout      JitRequestState = original.JitRequestStateTimeout
)

type JitSchedulingType = original.JitSchedulingType

const (
	JitSchedulingTypeNotSpecified JitSchedulingType = original.JitSchedulingTypeNotSpecified
	JitSchedulingTypeOnce         JitSchedulingType = original.JitSchedulingTypeOnce
	JitSchedulingTypeRecurring    JitSchedulingType = original.JitSchedulingTypeRecurring
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateNotSpecified ProvisioningState = original.ProvisioningStateNotSpecified
	ProvisioningStateRunning      ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type Status = original.Status

const (
	StatusElevate      Status = original.StatusElevate
	StatusNotSpecified Status = original.StatusNotSpecified
	StatusRemove       Status = original.StatusRemove
)

type Substatus = original.Substatus

const (
	SubstatusApproved     Substatus = original.SubstatusApproved
	SubstatusDenied       Substatus = original.SubstatusDenied
	SubstatusExpired      Substatus = original.SubstatusExpired
	SubstatusFailed       Substatus = original.SubstatusFailed
	SubstatusNotSpecified Substatus = original.SubstatusNotSpecified
	SubstatusTimeout      Substatus = original.SubstatusTimeout
)

type AllowedUpgradePlansResult = original.AllowedUpgradePlansResult
type Application = original.Application
type ApplicationArtifact = original.ApplicationArtifact
type ApplicationAuthorization = original.ApplicationAuthorization
type ApplicationBillingDetailsDefinition = original.ApplicationBillingDetailsDefinition
type ApplicationClientDetails = original.ApplicationClientDetails
type ApplicationDefinition = original.ApplicationDefinition
type ApplicationDefinitionArtifact = original.ApplicationDefinitionArtifact
type ApplicationDefinitionListResult = original.ApplicationDefinitionListResult
type ApplicationDefinitionListResultIterator = original.ApplicationDefinitionListResultIterator
type ApplicationDefinitionListResultPage = original.ApplicationDefinitionListResultPage
type ApplicationDefinitionPatchable = original.ApplicationDefinitionPatchable
type ApplicationDefinitionProperties = original.ApplicationDefinitionProperties
type ApplicationDefinitionsClient = original.ApplicationDefinitionsClient
type ApplicationDeploymentPolicy = original.ApplicationDeploymentPolicy
type ApplicationJitAccessPolicy = original.ApplicationJitAccessPolicy
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationManagementPolicy = original.ApplicationManagementPolicy
type ApplicationNotificationEndpoint = original.ApplicationNotificationEndpoint
type ApplicationNotificationPolicy = original.ApplicationNotificationPolicy
type ApplicationPackageContact = original.ApplicationPackageContact
type ApplicationPackageLockingPolicyDefinition = original.ApplicationPackageLockingPolicyDefinition
type ApplicationPackageSupportUrls = original.ApplicationPackageSupportUrls
type ApplicationPatchable = original.ApplicationPatchable
type ApplicationPolicy = original.ApplicationPolicy
type ApplicationProperties = original.ApplicationProperties
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateOrUpdateByIDFuture = original.ApplicationsCreateOrUpdateByIDFuture
type ApplicationsCreateOrUpdateFuture = original.ApplicationsCreateOrUpdateFuture
type ApplicationsDeleteByIDFuture = original.ApplicationsDeleteByIDFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type ApplicationsRefreshPermissionsFuture = original.ApplicationsRefreshPermissionsFuture
type ApplicationsUpdateAccessFuture = original.ApplicationsUpdateAccessFuture
type ApplicationsUpdateByIDFuture = original.ApplicationsUpdateByIDFuture
type ApplicationsUpdateFuture = original.ApplicationsUpdateFuture
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type GenericResource = original.GenericResource
type Identity = original.Identity
type JitApproverDefinition = original.JitApproverDefinition
type JitAuthorizationPolicies = original.JitAuthorizationPolicies
type JitRequestDefinition = original.JitRequestDefinition
type JitRequestDefinitionListResult = original.JitRequestDefinitionListResult
type JitRequestMetadata = original.JitRequestMetadata
type JitRequestPatchable = original.JitRequestPatchable
type JitRequestProperties = original.JitRequestProperties
type JitRequestsClient = original.JitRequestsClient
type JitRequestsCreateOrUpdateFuture = original.JitRequestsCreateOrUpdateFuture
type JitSchedulingPolicy = original.JitSchedulingPolicy
type ListTokenRequest = original.ListTokenRequest
type ManagedIdentityToken = original.ManagedIdentityToken
type ManagedIdentityTokenResult = original.ManagedIdentityTokenResult
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type Plan = original.Plan
type PlanPatchable = original.PlanPatchable
type Resource = original.Resource
type Sku = original.Sku
type SystemData = original.SystemData
type UpdateAccessDefinition = original.UpdateAccessDefinition
type UserAssignedResourceIdentity = original.UserAssignedResourceIdentity

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationDefinitionListResultIterator(page ApplicationDefinitionListResultPage) ApplicationDefinitionListResultIterator {
	return original.NewApplicationDefinitionListResultIterator(page)
}
func NewApplicationDefinitionListResultPage(cur ApplicationDefinitionListResult, getNextPage func(context.Context, ApplicationDefinitionListResult) (ApplicationDefinitionListResult, error)) ApplicationDefinitionListResultPage {
	return original.NewApplicationDefinitionListResultPage(cur, getNextPage)
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
func NewApplicationListResultPage(cur ApplicationListResult, getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(cur, getNextPage)
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
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleApplicationArtifactNameValues() []ApplicationArtifactName {
	return original.PossibleApplicationArtifactNameValues()
}
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return original.PossibleApplicationArtifactTypeValues()
}
func PossibleApplicationDefinitionArtifactNameValues() []ApplicationDefinitionArtifactName {
	return original.PossibleApplicationDefinitionArtifactNameValues()
}
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return original.PossibleApplicationLockLevelValues()
}
func PossibleApplicationManagementModeValues() []ApplicationManagementMode {
	return original.PossibleApplicationManagementModeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleJitApprovalModeValues() []JitApprovalMode {
	return original.PossibleJitApprovalModeValues()
}
func PossibleJitApproverTypeValues() []JitApproverType {
	return original.PossibleJitApproverTypeValues()
}
func PossibleJitRequestStateValues() []JitRequestState {
	return original.PossibleJitRequestStateValues()
}
func PossibleJitSchedulingTypeValues() []JitSchedulingType {
	return original.PossibleJitSchedulingTypeValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleSubstatusValues() []Substatus {
	return original.PossibleSubstatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
