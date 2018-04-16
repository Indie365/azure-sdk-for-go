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

package subscription

import original "github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2018-03-01-preview/subscription"

type OperationsClient = original.OperationsClient
type OfferType = original.OfferType

const (
	MSAZR0017P OfferType = original.MSAZR0017P
	MSAZR0148P OfferType = original.MSAZR0148P
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

type AdPrincipal = original.AdPrincipal
type CreationParameters = original.CreationParameters
type CreationResult = original.CreationResult
type ErrorResponse = original.ErrorResponse
type FactoryCreateSubscriptionInEnrollmentAccountFuture = original.FactoryCreateSubscriptionInEnrollmentAccountFuture
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Location = original.Location
type LocationListResult = original.LocationListResult
type Model = original.Model
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultType = original.OperationListResultType
type OperationType = original.OperationType
type Policies = original.Policies
type TenantIDDescription = original.TenantIDDescription
type TenantListResult = original.TenantListResult
type TenantListResultIterator = original.TenantListResultIterator
type TenantListResultPage = original.TenantListResultPage
type OperationsGroupClient = original.OperationsGroupClient
type SubscriptionsClient = original.SubscriptionsClient
type FactoryClient = original.FactoryClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type TenantsClient = original.TenantsClient

func NewOperationsGroupClient() OperationsGroupClient {
	return original.NewOperationsGroupClient()
}
func NewOperationsGroupClientWithBaseURI(baseURI string) OperationsGroupClient {
	return original.NewOperationsGroupClientWithBaseURI(baseURI)
}
func NewSubscriptionsClient() SubscriptionsClient {
	return original.NewSubscriptionsClient()
}
func NewSubscriptionsClientWithBaseURI(baseURI string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewFactoryClient() FactoryClient {
	return original.NewFactoryClient()
}
func NewFactoryClientWithBaseURI(baseURI string) FactoryClient {
	return original.NewFactoryClientWithBaseURI(baseURI)
}
func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewTenantsClient() TenantsClient {
	return original.NewTenantsClient()
}
func NewTenantsClientWithBaseURI(baseURI string) TenantsClient {
	return original.NewTenantsClientWithBaseURI(baseURI)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func PossibleOfferTypeValues() []OfferType {
	return original.PossibleOfferTypeValues()
}
func PossibleSpendingLimitValues() []SpendingLimit {
	return original.PossibleSpendingLimitValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
