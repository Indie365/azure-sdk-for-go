// +build go1.9

// Copyright 2017 Microsoft Corporation
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
// commit ID: 714052a3359963ba46703fe033cf9dd4838bea11

package insights

import original "github.com/Azure/azure-sdk-for-go/services/monitor/2017-05-01/insights"

type UsageMetricsClient = original.UsageMetricsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type EventCategoriesClient = original.EventCategoriesClient
type MetricDefinitionsClient = original.MetricDefinitionsClient
type MetricsClient = original.MetricsClient
type AggregationType = original.AggregationType

const (
	Average	AggregationType	= original.Average
	Count	AggregationType	= original.Count
	Maximum	AggregationType	= original.Maximum
	Minimum	AggregationType	= original.Minimum
	None	AggregationType	= original.None
	Total	AggregationType	= original.Total
)

type EventLevel = original.EventLevel

const (
	Critical	EventLevel	= original.Critical
	Error		EventLevel	= original.Error
	Informational	EventLevel	= original.Informational
	Verbose		EventLevel	= original.Verbose
	Warning		EventLevel	= original.Warning
)

type ResultType = original.ResultType

const (
	Data		ResultType	= original.Data
	Metadata	ResultType	= original.Metadata
)

type Unit = original.Unit

const (
	UnitBytes		Unit	= original.UnitBytes
	UnitByteSeconds		Unit	= original.UnitByteSeconds
	UnitBytesPerSecond	Unit	= original.UnitBytesPerSecond
	UnitCount		Unit	= original.UnitCount
	UnitCountPerSecond	Unit	= original.UnitCountPerSecond
	UnitMilliSeconds	Unit	= original.UnitMilliSeconds
	UnitPercent		Unit	= original.UnitPercent
	UnitSeconds		Unit	= original.UnitSeconds
	UnitUnspecified		Unit	= original.UnitUnspecified
)

type ErrorResponse = original.ErrorResponse
type EventCategoryCollection = original.EventCategoryCollection
type EventData = original.EventData
type EventDataCollection = original.EventDataCollection
type HTTPRequestInfo = original.HTTPRequestInfo
type LocalizableString = original.LocalizableString
type MetadataValue = original.MetadataValue
type Metric = original.Metric
type MetricAvailability = original.MetricAvailability
type MetricDefinition = original.MetricDefinition
type MetricDefinitionCollection = original.MetricDefinitionCollection
type MetricValue = original.MetricValue
type Response = original.Response
type SenderAuthorization = original.SenderAuthorization
type TimeSeriesElement = original.TimeSeriesElement
type UsageMetric = original.UsageMetric
type UsageMetricCollection = original.UsageMetricCollection
type TenantActivityLogsClient = original.TenantActivityLogsClient
type ActivityLogsClient = original.ActivityLogsClient

func NewUsageMetricsClient(subscriptionID string) UsageMetricsClient {
	return original.NewUsageMetricsClient(subscriptionID)
}
func NewUsageMetricsClientWithBaseURI(baseURI string, subscriptionID string) UsageMetricsClient {
	return original.NewUsageMetricsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewEventCategoriesClient(subscriptionID string) EventCategoriesClient {
	return original.NewEventCategoriesClient(subscriptionID)
}
func NewEventCategoriesClientWithBaseURI(baseURI string, subscriptionID string) EventCategoriesClient {
	return original.NewEventCategoriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricDefinitionsClient(subscriptionID string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClient(subscriptionID)
}
func NewMetricDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricsClient(subscriptionID string) MetricsClient {
	return original.NewMetricsClient(subscriptionID)
}
func NewMetricsClientWithBaseURI(baseURI string, subscriptionID string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTenantActivityLogsClient(subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClient(subscriptionID)
}
func NewTenantActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}
func NewActivityLogsClient(subscriptionID string) ActivityLogsClient {
	return original.NewActivityLogsClient(subscriptionID)
}
func NewActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogsClient {
	return original.NewActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
