//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package policy

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-09-01/policy"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EnforcementMode = original.EnforcementMode

const (
	Default      EnforcementMode = original.Default
	DoNotEnforce EnforcementMode = original.DoNotEnforce
)

type ParameterType = original.ParameterType

const (
	Array    ParameterType = original.Array
	Boolean  ParameterType = original.Boolean
	DateTime ParameterType = original.DateTime
	Float    ParameterType = original.Float
	Integer  ParameterType = original.Integer
	Object   ParameterType = original.Object
	String   ParameterType = original.String
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None           ResourceIdentityType = original.None
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type Type = original.Type

const (
	BuiltIn      Type = original.BuiltIn
	Custom       Type = original.Custom
	NotSpecified Type = original.NotSpecified
	Static       Type = original.Static
)

type Assignment = original.Assignment
type AssignmentListResult = original.AssignmentListResult
type AssignmentListResultIterator = original.AssignmentListResultIterator
type AssignmentListResultPage = original.AssignmentListResultPage
type AssignmentProperties = original.AssignmentProperties
type AssignmentsClient = original.AssignmentsClient
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type Definition = original.Definition
type DefinitionGroup = original.DefinitionGroup
type DefinitionListResult = original.DefinitionListResult
type DefinitionListResultIterator = original.DefinitionListResultIterator
type DefinitionListResultPage = original.DefinitionListResultPage
type DefinitionProperties = original.DefinitionProperties
type DefinitionReference = original.DefinitionReference
type DefinitionsClient = original.DefinitionsClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type Identity = original.Identity
type ParameterDefinitionsValue = original.ParameterDefinitionsValue
type ParameterDefinitionsValueMetadata = original.ParameterDefinitionsValueMetadata
type ParameterValuesValue = original.ParameterValuesValue
type SetDefinition = original.SetDefinition
type SetDefinitionListResult = original.SetDefinitionListResult
type SetDefinitionListResultIterator = original.SetDefinitionListResultIterator
type SetDefinitionListResultPage = original.SetDefinitionListResultPage
type SetDefinitionProperties = original.SetDefinitionProperties
type SetDefinitionsClient = original.SetDefinitionsClient
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAssignmentListResultIterator(page AssignmentListResultPage) AssignmentListResultIterator {
	return original.NewAssignmentListResultIterator(page)
}
func NewAssignmentListResultPage(cur AssignmentListResult, getNextPage func(context.Context, AssignmentListResult) (AssignmentListResult, error)) AssignmentListResultPage {
	return original.NewAssignmentListResultPage(cur, getNextPage)
}
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClient(subscriptionID)
}
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDefinitionListResultIterator(page DefinitionListResultPage) DefinitionListResultIterator {
	return original.NewDefinitionListResultIterator(page)
}
func NewDefinitionListResultPage(cur DefinitionListResult, getNextPage func(context.Context, DefinitionListResult) (DefinitionListResult, error)) DefinitionListResultPage {
	return original.NewDefinitionListResultPage(cur, getNextPage)
}
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClient(subscriptionID)
}
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSetDefinitionListResultIterator(page SetDefinitionListResultPage) SetDefinitionListResultIterator {
	return original.NewSetDefinitionListResultIterator(page)
}
func NewSetDefinitionListResultPage(cur SetDefinitionListResult, getNextPage func(context.Context, SetDefinitionListResult) (SetDefinitionListResult, error)) SetDefinitionListResultPage {
	return original.NewSetDefinitionListResultPage(cur, getNextPage)
}
func NewSetDefinitionsClient(subscriptionID string) SetDefinitionsClient {
	return original.NewSetDefinitionsClient(subscriptionID)
}
func NewSetDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) SetDefinitionsClient {
	return original.NewSetDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleEnforcementModeValues() []EnforcementMode {
	return original.PossibleEnforcementModeValues()
}
func PossibleParameterTypeValues() []ParameterType {
	return original.PossibleParameterTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
