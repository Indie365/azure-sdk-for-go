// +build go1.9

// Copyright 2020 Microsoft Corporation
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

package subscription

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted  ProvisioningState = original.Accepted
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type SpendingLimit = original.SpendingLimit

const (
	CurrentPeriodOff SpendingLimit = original.CurrentPeriodOff
	Off              SpendingLimit = original.Off
	On               SpendingLimit = original.On
)

type State = original.State

const (
	Deleted  State = original.Deleted
	Disabled State = original.Disabled
	Enabled  State = original.Enabled
	PastDue  State = original.PastDue
	Warned   State = original.Warned
)

type Workload = original.Workload

const (
	DevTest    Workload = original.DevTest
	Production Workload = original.Production
)

type AliasClient = original.AliasClient
type AliasCreateFuture = original.AliasCreateFuture
type BaseClient = original.BaseClient
type CanceledSubscriptionID = original.CanceledSubscriptionID
type Client = original.Client
type EnabledSubscriptionID = original.EnabledSubscriptionID
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Location = original.Location
type LocationListResult = original.LocationListResult
type Model = original.Model
type Name = original.Name
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Policies = original.Policies
type PutAliasListResult = original.PutAliasListResult
type PutAliasRequest = original.PutAliasRequest
type PutAliasRequestProperties = original.PutAliasRequestProperties
type PutAliasResponse = original.PutAliasResponse
type PutAliasResponseProperties = original.PutAliasResponseProperties
type RenamedSubscriptionID = original.RenamedSubscriptionID
type SubscriptionsClient = original.SubscriptionsClient
type TenantIDDescription = original.TenantIDDescription
type TenantListResult = original.TenantListResult
type TenantListResultIterator = original.TenantListResultIterator
type TenantListResultPage = original.TenantListResultPage
type TenantsClient = original.TenantsClient

func New() BaseClient {
	return original.New()
}
func NewAliasClient() AliasClient {
	return original.NewAliasClient()
}
func NewAliasClientWithBaseURI(baseURI string) AliasClient {
	return original.NewAliasClientWithBaseURI(baseURI)
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewSubscriptionsClient() SubscriptionsClient {
	return original.NewSubscriptionsClient()
}
func NewSubscriptionsClientWithBaseURI(baseURI string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI)
}
func NewTenantListResultIterator(page TenantListResultPage) TenantListResultIterator {
	return original.NewTenantListResultIterator(page)
}
func NewTenantsClient() TenantsClient {
	return original.NewTenantsClient()
}
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return original.NewTenantsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSpendingLimitValues() []SpendingLimit {
	return original.PossibleSpendingLimitValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleWorkloadValues() []Workload {
	return original.PossibleWorkloadValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
