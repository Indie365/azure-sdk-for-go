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

package commerce

import original "github.com/Azure/azure-sdk-for-go/services/commerce/mgmt/2015-06-01-preview/commerce"

type UsageAggregatesClient = original.UsageAggregatesClient
type AggregationGranularity = original.AggregationGranularity

const (
	Daily  AggregationGranularity = original.Daily
	Hourly AggregationGranularity = original.Hourly
)

type Name = original.Name

const (
	NameMonetaryCommitment Name = original.NameMonetaryCommitment
	NameMonetaryCredit     Name = original.NameMonetaryCredit
	NameOfferTermInfo      Name = original.NameOfferTermInfo
	NameRecurringCharge    Name = original.NameRecurringCharge
)

type ErrorResponse = original.ErrorResponse
type InfoField = original.InfoField
type MeterInfo = original.MeterInfo
type MonetaryCommitment = original.MonetaryCommitment
type MonetaryCredit = original.MonetaryCredit
type BasicOfferTermInfo = original.BasicOfferTermInfo
type OfferTermInfo = original.OfferTermInfo
type RateCardQueryParameters = original.RateCardQueryParameters
type RecurringCharge = original.RecurringCharge
type ResourceRateCardInfo = original.ResourceRateCardInfo
type UsageAggregation = original.UsageAggregation
type UsageAggregationListResult = original.UsageAggregationListResult
type UsageAggregationListResultIterator = original.UsageAggregationListResultIterator
type UsageAggregationListResultPage = original.UsageAggregationListResultPage
type UsageSample = original.UsageSample
type RateCardClient = original.RateCardClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func NewUsageAggregatesClient(subscriptionID string) UsageAggregatesClient {
	return original.NewUsageAggregatesClient(subscriptionID)
}
func NewUsageAggregatesClientWithBaseURI(baseURI string, subscriptionID string) UsageAggregatesClient {
	return original.NewUsageAggregatesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAggregationGranularityValues() []AggregationGranularity {
	return original.PossibleAggregationGranularityValues()
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewRateCardClient(subscriptionID string) RateCardClient {
	return original.NewRateCardClient(subscriptionID)
}
func NewRateCardClientWithBaseURI(baseURI string, subscriptionID string) RateCardClient {
	return original.NewRateCardClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
