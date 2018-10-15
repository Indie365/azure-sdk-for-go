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

package consumption

import original "github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2017-11-30/consumption"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type Datagrain = original.Datagrain

const (
	DailyGrain   Datagrain = original.DailyGrain
	MonthlyGrain Datagrain = original.MonthlyGrain
)

type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type MeterDetails = original.MeterDetails
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type ReservationDetails = original.ReservationDetails
type ReservationDetailsListResult = original.ReservationDetailsListResult
type ReservationDetailsProperties = original.ReservationDetailsProperties
type ReservationSummaries = original.ReservationSummaries
type ReservationSummariesListResult = original.ReservationSummariesListResult
type ReservationSummariesProperties = original.ReservationSummariesProperties
type Resource = original.Resource
type UsageDetail = original.UsageDetail
type UsageDetailProperties = original.UsageDetailProperties
type UsageDetailsListResult = original.UsageDetailsListResult
type UsageDetailsListResultIterator = original.UsageDetailsListResultIterator
type UsageDetailsListResultPage = original.UsageDetailsListResultPage
type OperationsClient = original.OperationsClient
type OperationsGroupClient = original.OperationsGroupClient
type ReservationsDetailsClient = original.ReservationsDetailsClient
type ReservationsDetailsGroupClient = original.ReservationsDetailsGroupClient
type ReservationsSummariesClient = original.ReservationsSummariesClient
type ReservationsSummariesGroupClient = original.ReservationsSummariesGroupClient
type UsageDetailsClient = original.UsageDetailsClient
type UsageDetailsGroupClient = original.UsageDetailsGroupClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDatagrainValues() []Datagrain {
	return original.PossibleDatagrainValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsGroupClient(subscriptionID string) OperationsGroupClient {
	return original.NewOperationsGroupClient(subscriptionID)
}
func NewOperationsGroupClientWithBaseURI(baseURI string, subscriptionID string) OperationsGroupClient {
	return original.NewOperationsGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsDetailsClient(subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClient(subscriptionID)
}
func NewReservationsDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsDetailsGroupClient(subscriptionID string) ReservationsDetailsGroupClient {
	return original.NewReservationsDetailsGroupClient(subscriptionID)
}
func NewReservationsDetailsGroupClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsGroupClient {
	return original.NewReservationsDetailsGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsSummariesClient(subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClient(subscriptionID)
}
func NewReservationsSummariesClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsSummariesGroupClient(subscriptionID string) ReservationsSummariesGroupClient {
	return original.NewReservationsSummariesGroupClient(subscriptionID)
}
func NewReservationsSummariesGroupClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesGroupClient {
	return original.NewReservationsSummariesGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClient(subscriptionID)
}
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageDetailsGroupClient(subscriptionID string) UsageDetailsGroupClient {
	return original.NewUsageDetailsGroupClient(subscriptionID)
}
func NewUsageDetailsGroupClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsGroupClient {
	return original.NewUsageDetailsGroupClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
