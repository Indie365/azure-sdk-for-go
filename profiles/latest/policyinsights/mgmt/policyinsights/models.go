//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package policyinsights

import original "github.com/Azure/azure-sdk-for-go/services/policyinsights/mgmt/2018-04-04/policyinsights"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type PolicyStatesResource = original.PolicyStatesResource

const (
	Default PolicyStatesResource = original.Default
	Latest  PolicyStatesResource = original.Latest
)

type BaseClient = original.BaseClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsListResults = original.OperationsListResults
type PolicyAssignmentSummary = original.PolicyAssignmentSummary
type PolicyDefinitionSummary = original.PolicyDefinitionSummary
type PolicyEvent = original.PolicyEvent
type PolicyEventsClient = original.PolicyEventsClient
type PolicyEventsQueryResults = original.PolicyEventsQueryResults
type PolicyState = original.PolicyState
type PolicyStatesClient = original.PolicyStatesClient
type PolicyStatesQueryResults = original.PolicyStatesQueryResults
type QueryFailure = original.QueryFailure
type QueryFailureError = original.QueryFailureError
type String = original.String
type SummarizeResults = original.SummarizeResults
type Summary = original.Summary
type SummaryResults = original.SummaryResults

func New() BaseClient {
	return original.New()
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewPolicyEventsClient() PolicyEventsClient {
	return original.NewPolicyEventsClient()
}
func NewPolicyEventsClientWithBaseURI(baseURI string) PolicyEventsClient {
	return original.NewPolicyEventsClientWithBaseURI(baseURI)
}
func NewPolicyStatesClient() PolicyStatesClient {
	return original.NewPolicyStatesClient()
}
func NewPolicyStatesClientWithBaseURI(baseURI string) PolicyStatesClient {
	return original.NewPolicyStatesClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return original.PossiblePolicyStatesResourceValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
