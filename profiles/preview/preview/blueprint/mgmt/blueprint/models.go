// +build go1.9

// Copyright 2021 Microsoft Corporation
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

package blueprint

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/blueprint/mgmt/2018-11-01-preview/blueprint"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AssignmentDeleteBehavior = original.AssignmentDeleteBehavior

const (
	All  AssignmentDeleteBehavior = original.All
	None AssignmentDeleteBehavior = original.None
)

type AssignmentLockMode = original.AssignmentLockMode

const (
	AssignmentLockModeAllResourcesDoNotDelete AssignmentLockMode = original.AssignmentLockModeAllResourcesDoNotDelete
	AssignmentLockModeAllResourcesReadOnly    AssignmentLockMode = original.AssignmentLockModeAllResourcesReadOnly
	AssignmentLockModeNone                    AssignmentLockMode = original.AssignmentLockModeNone
)

type AssignmentProvisioningState = original.AssignmentProvisioningState

const (
	Canceled   AssignmentProvisioningState = original.Canceled
	Cancelling AssignmentProvisioningState = original.Cancelling
	Creating   AssignmentProvisioningState = original.Creating
	Deleting   AssignmentProvisioningState = original.Deleting
	Deploying  AssignmentProvisioningState = original.Deploying
	Failed     AssignmentProvisioningState = original.Failed
	Locking    AssignmentProvisioningState = original.Locking
	Succeeded  AssignmentProvisioningState = original.Succeeded
	Validating AssignmentProvisioningState = original.Validating
	Waiting    AssignmentProvisioningState = original.Waiting
)

type Kind = original.Kind

const (
	KindArtifact         Kind = original.KindArtifact
	KindPolicyAssignment Kind = original.KindPolicyAssignment
	KindRoleAssignment   Kind = original.KindRoleAssignment
	KindTemplate         Kind = original.KindTemplate
)

type ManagedServiceIdentityType = original.ManagedServiceIdentityType

const (
	ManagedServiceIdentityTypeNone           ManagedServiceIdentityType = original.ManagedServiceIdentityTypeNone
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = original.ManagedServiceIdentityTypeSystemAssigned
	ManagedServiceIdentityTypeUserAssigned   ManagedServiceIdentityType = original.ManagedServiceIdentityTypeUserAssigned
)

type TargetScope = original.TargetScope

const (
	ManagementGroup TargetScope = original.ManagementGroup
	Subscription    TargetScope = original.Subscription
)

type TemplateParameterType = original.TemplateParameterType

const (
	Array        TemplateParameterType = original.Array
	Bool         TemplateParameterType = original.Bool
	Int          TemplateParameterType = original.Int
	Object       TemplateParameterType = original.Object
	SecureObject TemplateParameterType = original.SecureObject
	SecureString TemplateParameterType = original.SecureString
	String       TemplateParameterType = original.String
)

type Artifact = original.Artifact
type ArtifactList = original.ArtifactList
type ArtifactListIterator = original.ArtifactListIterator
type ArtifactListPage = original.ArtifactListPage
type ArtifactModel = original.ArtifactModel
type ArtifactPropertiesBase = original.ArtifactPropertiesBase
type ArtifactsClient = original.ArtifactsClient
type Assignment = original.Assignment
type AssignmentDeploymentJob = original.AssignmentDeploymentJob
type AssignmentDeploymentJobResult = original.AssignmentDeploymentJobResult
type AssignmentJobCreatedResource = original.AssignmentJobCreatedResource
type AssignmentList = original.AssignmentList
type AssignmentListIterator = original.AssignmentListIterator
type AssignmentListPage = original.AssignmentListPage
type AssignmentLockSettings = original.AssignmentLockSettings
type AssignmentOperation = original.AssignmentOperation
type AssignmentOperationList = original.AssignmentOperationList
type AssignmentOperationListIterator = original.AssignmentOperationListIterator
type AssignmentOperationListPage = original.AssignmentOperationListPage
type AssignmentOperationProperties = original.AssignmentOperationProperties
type AssignmentOperationsClient = original.AssignmentOperationsClient
type AssignmentProperties = original.AssignmentProperties
type AssignmentStatus = original.AssignmentStatus
type AssignmentsClient = original.AssignmentsClient
type AzureResourceBase = original.AzureResourceBase
type AzureResourceManagerError = original.AzureResourceManagerError
type BaseClient = original.BaseClient
type BasicArtifact = original.BasicArtifact
type BlueprintsClient = original.BlueprintsClient
type KeyVaultReference = original.KeyVaultReference
type List = original.List
type ListIterator = original.ListIterator
type ListPage = original.ListPage
type ManagedServiceIdentity = original.ManagedServiceIdentity
type Model = original.Model
type ParameterDefinition = original.ParameterDefinition
type ParameterDefinitionMetadata = original.ParameterDefinitionMetadata
type ParameterValue = original.ParameterValue
type PolicyAssignmentArtifact = original.PolicyAssignmentArtifact
type PolicyAssignmentArtifactProperties = original.PolicyAssignmentArtifactProperties
type Properties = original.Properties
type PublishedArtifactsClient = original.PublishedArtifactsClient
type PublishedBlueprint = original.PublishedBlueprint
type PublishedBlueprintList = original.PublishedBlueprintList
type PublishedBlueprintListIterator = original.PublishedBlueprintListIterator
type PublishedBlueprintListPage = original.PublishedBlueprintListPage
type PublishedBlueprintProperties = original.PublishedBlueprintProperties
type PublishedBlueprintsClient = original.PublishedBlueprintsClient
type ResourceGroupDefinition = original.ResourceGroupDefinition
type ResourceGroupValue = original.ResourceGroupValue
type ResourcePropertiesBase = original.ResourcePropertiesBase
type ResourceProviderOperation = original.ResourceProviderOperation
type ResourceProviderOperationDisplay = original.ResourceProviderOperationDisplay
type ResourceProviderOperationList = original.ResourceProviderOperationList
type ResourceStatusBase = original.ResourceStatusBase
type RoleAssignmentArtifact = original.RoleAssignmentArtifact
type RoleAssignmentArtifactProperties = original.RoleAssignmentArtifactProperties
type SecretValueReference = original.SecretValueReference
type SharedBlueprintProperties = original.SharedBlueprintProperties
type Status = original.Status
type TemplateArtifact = original.TemplateArtifact
type TemplateArtifactProperties = original.TemplateArtifactProperties
type TrackedResource = original.TrackedResource
type UserAssignedIdentity = original.UserAssignedIdentity
type WhoIsBlueprintContract = original.WhoIsBlueprintContract

func New() BaseClient {
	return original.New()
}
func NewArtifactListIterator(page ArtifactListPage) ArtifactListIterator {
	return original.NewArtifactListIterator(page)
}
func NewArtifactListPage(cur ArtifactList, getNextPage func(context.Context, ArtifactList) (ArtifactList, error)) ArtifactListPage {
	return original.NewArtifactListPage(cur, getNextPage)
}
func NewArtifactsClient() ArtifactsClient {
	return original.NewArtifactsClient()
}
func NewArtifactsClientWithBaseURI(baseURI string) ArtifactsClient {
	return original.NewArtifactsClientWithBaseURI(baseURI)
}
func NewAssignmentListIterator(page AssignmentListPage) AssignmentListIterator {
	return original.NewAssignmentListIterator(page)
}
func NewAssignmentListPage(cur AssignmentList, getNextPage func(context.Context, AssignmentList) (AssignmentList, error)) AssignmentListPage {
	return original.NewAssignmentListPage(cur, getNextPage)
}
func NewAssignmentOperationListIterator(page AssignmentOperationListPage) AssignmentOperationListIterator {
	return original.NewAssignmentOperationListIterator(page)
}
func NewAssignmentOperationListPage(cur AssignmentOperationList, getNextPage func(context.Context, AssignmentOperationList) (AssignmentOperationList, error)) AssignmentOperationListPage {
	return original.NewAssignmentOperationListPage(cur, getNextPage)
}
func NewAssignmentOperationsClient() AssignmentOperationsClient {
	return original.NewAssignmentOperationsClient()
}
func NewAssignmentOperationsClientWithBaseURI(baseURI string) AssignmentOperationsClient {
	return original.NewAssignmentOperationsClientWithBaseURI(baseURI)
}
func NewAssignmentsClient() AssignmentsClient {
	return original.NewAssignmentsClient()
}
func NewAssignmentsClientWithBaseURI(baseURI string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI)
}
func NewBlueprintsClient() BlueprintsClient {
	return original.NewBlueprintsClient()
}
func NewBlueprintsClientWithBaseURI(baseURI string) BlueprintsClient {
	return original.NewBlueprintsClientWithBaseURI(baseURI)
}
func NewListIterator(page ListPage) ListIterator {
	return original.NewListIterator(page)
}
func NewListPage(cur List, getNextPage func(context.Context, List) (List, error)) ListPage {
	return original.NewListPage(cur, getNextPage)
}
func NewPublishedArtifactsClient() PublishedArtifactsClient {
	return original.NewPublishedArtifactsClient()
}
func NewPublishedArtifactsClientWithBaseURI(baseURI string) PublishedArtifactsClient {
	return original.NewPublishedArtifactsClientWithBaseURI(baseURI)
}
func NewPublishedBlueprintListIterator(page PublishedBlueprintListPage) PublishedBlueprintListIterator {
	return original.NewPublishedBlueprintListIterator(page)
}
func NewPublishedBlueprintListPage(cur PublishedBlueprintList, getNextPage func(context.Context, PublishedBlueprintList) (PublishedBlueprintList, error)) PublishedBlueprintListPage {
	return original.NewPublishedBlueprintListPage(cur, getNextPage)
}
func NewPublishedBlueprintsClient() PublishedBlueprintsClient {
	return original.NewPublishedBlueprintsClient()
}
func NewPublishedBlueprintsClientWithBaseURI(baseURI string) PublishedBlueprintsClient {
	return original.NewPublishedBlueprintsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleAssignmentDeleteBehaviorValues() []AssignmentDeleteBehavior {
	return original.PossibleAssignmentDeleteBehaviorValues()
}
func PossibleAssignmentLockModeValues() []AssignmentLockMode {
	return original.PossibleAssignmentLockModeValues()
}
func PossibleAssignmentProvisioningStateValues() []AssignmentProvisioningState {
	return original.PossibleAssignmentProvisioningStateValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return original.PossibleManagedServiceIdentityTypeValues()
}
func PossibleTargetScopeValues() []TargetScope {
	return original.PossibleTargetScopeValues()
}
func PossibleTemplateParameterTypeValues() []TemplateParameterType {
	return original.PossibleTemplateParameterTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
