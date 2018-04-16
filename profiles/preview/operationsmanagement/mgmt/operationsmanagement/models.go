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

package operationsmanagement

import original "github.com/Azure/azure-sdk-for-go/services/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"

type ArmTemplateParameter = original.ArmTemplateParameter
type CodeMessageError = original.CodeMessageError
type CodeMessageErrorError = original.CodeMessageErrorError
type ManagementAssociation = original.ManagementAssociation
type ManagementAssociationProperties = original.ManagementAssociationProperties
type ManagementAssociationPropertiesList = original.ManagementAssociationPropertiesList
type ManagementConfiguration = original.ManagementConfiguration
type ManagementConfigurationProperties = original.ManagementConfigurationProperties
type ManagementConfigurationPropertiesList = original.ManagementConfigurationPropertiesList
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Solution = original.Solution
type SolutionPlan = original.SolutionPlan
type SolutionProperties = original.SolutionProperties
type SolutionPropertiesList = original.SolutionPropertiesList
type ManagementConfigurationsClient = original.ManagementConfigurationsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ManagementAssociationsClient = original.ManagementAssociationsClient
type OperationsClient = original.OperationsClient
type SolutionsClient = original.SolutionsClient

func NewSolutionsClient(subscriptionID string, providerName string, resourceType string, resourceName string) SolutionsClient {
	return original.NewSolutionsClient(subscriptionID, providerName, resourceType, resourceName)
}
func NewSolutionsClientWithBaseURI(baseURI string, subscriptionID string, providerName string, resourceType string, resourceName string) SolutionsClient {
	return original.NewSolutionsClientWithBaseURI(baseURI, subscriptionID, providerName, resourceType, resourceName)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewManagementConfigurationsClient(subscriptionID string, providerName string, resourceType string, resourceName string) ManagementConfigurationsClient {
	return original.NewManagementConfigurationsClient(subscriptionID, providerName, resourceType, resourceName)
}
func NewManagementConfigurationsClientWithBaseURI(baseURI string, subscriptionID string, providerName string, resourceType string, resourceName string) ManagementConfigurationsClient {
	return original.NewManagementConfigurationsClientWithBaseURI(baseURI, subscriptionID, providerName, resourceType, resourceName)
}
func New(subscriptionID string, providerName string, resourceType string, resourceName string) BaseClient {
	return original.New(subscriptionID, providerName, resourceType, resourceName)
}
func NewWithBaseURI(baseURI string, subscriptionID string, providerName string, resourceType string, resourceName string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, providerName, resourceType, resourceName)
}
func NewManagementAssociationsClient(subscriptionID string, providerName string, resourceType string, resourceName string) ManagementAssociationsClient {
	return original.NewManagementAssociationsClient(subscriptionID, providerName, resourceType, resourceName)
}
func NewManagementAssociationsClientWithBaseURI(baseURI string, subscriptionID string, providerName string, resourceType string, resourceName string) ManagementAssociationsClient {
	return original.NewManagementAssociationsClientWithBaseURI(baseURI, subscriptionID, providerName, resourceType, resourceName)
}
func NewOperationsClient(subscriptionID string, providerName string, resourceType string, resourceName string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, providerName, resourceType, resourceName)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, providerName string, resourceType string, resourceName string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, providerName, resourceType, resourceName)
}
