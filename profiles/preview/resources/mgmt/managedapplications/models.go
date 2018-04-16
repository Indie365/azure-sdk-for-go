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

package managedapplications

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01-preview/managedapplications"

type AppliancesClient = original.AppliancesClient
type ApplianceDefinitionsClient = original.ApplianceDefinitionsClient
type ApplianceArtifactType = original.ApplianceArtifactType

const (
	Custom   ApplianceArtifactType = original.Custom
	Template ApplianceArtifactType = original.Template
)

type ApplianceLockLevel = original.ApplianceLockLevel

const (
	CanNotDelete ApplianceLockLevel = original.CanNotDelete
	None         ApplianceLockLevel = original.None
	ReadOnly     ApplianceLockLevel = original.ReadOnly
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Canceled  ProvisioningState = original.Canceled
	Created   ProvisioningState = original.Created
	Creating  ProvisioningState = original.Creating
	Deleted   ProvisioningState = original.Deleted
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Ready     ProvisioningState = original.Ready
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type Appliance = original.Appliance
type ApplianceArtifact = original.ApplianceArtifact
type ApplianceDefinition = original.ApplianceDefinition
type ApplianceDefinitionListResult = original.ApplianceDefinitionListResult
type ApplianceDefinitionListResultIterator = original.ApplianceDefinitionListResultIterator
type ApplianceDefinitionListResultPage = original.ApplianceDefinitionListResultPage
type ApplianceDefinitionProperties = original.ApplianceDefinitionProperties
type ApplianceDefinitionsCreateOrUpdateByIDFuture = original.ApplianceDefinitionsCreateOrUpdateByIDFuture
type ApplianceDefinitionsCreateOrUpdateFuture = original.ApplianceDefinitionsCreateOrUpdateFuture
type ApplianceDefinitionsDeleteByIDFuture = original.ApplianceDefinitionsDeleteByIDFuture
type ApplianceDefinitionsDeleteFuture = original.ApplianceDefinitionsDeleteFuture
type ApplianceListResult = original.ApplianceListResult
type ApplianceListResultIterator = original.ApplianceListResultIterator
type ApplianceListResultPage = original.ApplianceListResultPage
type AppliancePatchable = original.AppliancePatchable
type ApplianceProperties = original.ApplianceProperties
type AppliancePropertiesPatchable = original.AppliancePropertiesPatchable
type ApplianceProviderAuthorization = original.ApplianceProviderAuthorization
type AppliancesCreateOrUpdateByIDFuture = original.AppliancesCreateOrUpdateByIDFuture
type AppliancesCreateOrUpdateFuture = original.AppliancesCreateOrUpdateFuture
type AppliancesDeleteByIDFuture = original.AppliancesDeleteByIDFuture
type AppliancesDeleteFuture = original.AppliancesDeleteFuture
type ErrorResponse = original.ErrorResponse
type GenericResource = original.GenericResource
type Identity = original.Identity
type Plan = original.Plan
type PlanPatchable = original.PlanPatchable
type Resource = original.Resource
type Sku = original.Sku

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewAppliancesClient(subscriptionID string) AppliancesClient {
	return original.NewAppliancesClient(subscriptionID)
}
func NewAppliancesClientWithBaseURI(baseURI string, subscriptionID string) AppliancesClient {
	return original.NewAppliancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplianceDefinitionsClient(subscriptionID string) ApplianceDefinitionsClient {
	return original.NewApplianceDefinitionsClient(subscriptionID)
}
func NewApplianceDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ApplianceDefinitionsClient {
	return original.NewApplianceDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplianceArtifactTypeValues() []ApplianceArtifactType {
	return original.PossibleApplianceArtifactTypeValues()
}
func PossibleApplianceLockLevelValues() []ApplianceLockLevel {
	return original.PossibleApplianceLockLevelValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
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
