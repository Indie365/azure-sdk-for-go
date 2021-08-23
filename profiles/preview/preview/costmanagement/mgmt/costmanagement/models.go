//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package costmanagement

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/costmanagement/mgmt/2019-03-01/costmanagement"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ConnectorBillingModel = original.ConnectorBillingModel

const (
	AutoUpgrade ConnectorBillingModel = original.AutoUpgrade
	Expired     ConnectorBillingModel = original.Expired
	Premium     ConnectorBillingModel = original.Premium
	Trial       ConnectorBillingModel = original.Trial
)

type ConnectorStatus = original.ConnectorStatus

const (
	ConnectorStatusActive  ConnectorStatus = original.ConnectorStatusActive
	ConnectorStatusError   ConnectorStatus = original.ConnectorStatusError
	ConnectorStatusExpired ConnectorStatus = original.ConnectorStatusExpired
	ConnectorStatusWarning ConnectorStatus = original.ConnectorStatusWarning
)

type CostAllocationPolicy = original.CostAllocationPolicy

const (
	Evenly       CostAllocationPolicy = original.Evenly
	Fixed        CostAllocationPolicy = original.Fixed
	Proportional CostAllocationPolicy = original.Proportional
)

type Direction = original.Direction

const (
	Ascending  Direction = original.Ascending
	Descending Direction = original.Descending
)

type FunctionType = original.FunctionType

const (
	AHUB         FunctionType = original.AHUB
	All          FunctionType = original.All
	None         FunctionType = original.None
	Reservations FunctionType = original.Reservations
)

type GranularityType = original.GranularityType

const (
	Daily   GranularityType = original.Daily
	Monthly GranularityType = original.Monthly
)

type OperatorType = original.OperatorType

const (
	Contains OperatorType = original.Contains
	In       OperatorType = original.In
)

type ReportConfigColumnType = original.ReportConfigColumnType

const (
	ReportConfigColumnTypeDimension ReportConfigColumnType = original.ReportConfigColumnTypeDimension
	ReportConfigColumnTypeTag       ReportConfigColumnType = original.ReportConfigColumnTypeTag
)

type RuleType = original.RuleType

const (
	RuleTypeCostAllocation         RuleType = original.RuleTypeCostAllocation
	RuleTypeCustomPrice            RuleType = original.RuleTypeCustomPrice
	RuleTypeShowbackRuleProperties RuleType = original.RuleTypeShowbackRuleProperties
)

type ShowbackRuleStatus = original.ShowbackRuleStatus

const (
	Active    ShowbackRuleStatus = original.Active
	NotActive ShowbackRuleStatus = original.NotActive
)

type TimeframeType = original.TimeframeType

const (
	Custom      TimeframeType = original.Custom
	MonthToDate TimeframeType = original.MonthToDate
	WeekToDate  TimeframeType = original.WeekToDate
	YearToDate  TimeframeType = original.YearToDate
)

type BaseClient = original.BaseClient
type BasicShowbackRuleProperties = original.BasicShowbackRuleProperties
type CheckEligibilityDefinition = original.CheckEligibilityDefinition
type CloudConnectorClient = original.CloudConnectorClient
type ConnectorClient = original.ConnectorClient
type ConnectorCollectionErrorInfo = original.ConnectorCollectionErrorInfo
type ConnectorCollectionInfo = original.ConnectorCollectionInfo
type ConnectorDefinition = original.ConnectorDefinition
type ConnectorDefinitionListResult = original.ConnectorDefinitionListResult
type ConnectorProperties = original.ConnectorProperties
type CostAllocationDetails = original.CostAllocationDetails
type CostAllocationDetailsKind = original.CostAllocationDetailsKind
type CustomPriceDetails = original.CustomPriceDetails
type CustomPriceDetailsKind = original.CustomPriceDetailsKind
type Dimension = original.Dimension
type DimensionProperties = original.DimensionProperties
type DimensionsClient = original.DimensionsClient
type DimensionsListResult = original.DimensionsListResult
type ErrorBase = original.ErrorBase
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type ExternalBillingAccountClient = original.ExternalBillingAccountClient
type ExternalBillingAccountDefinition = original.ExternalBillingAccountDefinition
type ExternalBillingAccountDefinitionListResult = original.ExternalBillingAccountDefinitionListResult
type ExternalBillingAccountProperties = original.ExternalBillingAccountProperties
type ExternalSubscriptionClient = original.ExternalSubscriptionClient
type ExternalSubscriptionDefinition = original.ExternalSubscriptionDefinition
type ExternalSubscriptionDefinitionListResult = original.ExternalSubscriptionDefinitionListResult
type ExternalSubscriptionIDListRequest = original.ExternalSubscriptionIDListRequest
type ExternalSubscriptionProperties = original.ExternalSubscriptionProperties
type ForecastClient = original.ForecastClient
type Markup = original.Markup
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Query = original.Query
type QueryClient = original.QueryClient
type QueryColumn = original.QueryColumn
type QueryProperties = original.QueryProperties
type QueryResult = original.QueryResult
type ReportConfigAggregation = original.ReportConfigAggregation
type ReportConfigComparisonExpression = original.ReportConfigComparisonExpression
type ReportConfigDataset = original.ReportConfigDataset
type ReportConfigDatasetConfiguration = original.ReportConfigDatasetConfiguration
type ReportConfigDefinition = original.ReportConfigDefinition
type ReportConfigFilter = original.ReportConfigFilter
type ReportConfigGrouping = original.ReportConfigGrouping
type ReportConfigSorting = original.ReportConfigSorting
type ReportConfigTimePeriod = original.ReportConfigTimePeriod
type Resource = original.Resource
type Scope = original.Scope
type ShowbackRule = original.ShowbackRule
type ShowbackRuleClient = original.ShowbackRuleClient
type ShowbackRuleListResult = original.ShowbackRuleListResult
type ShowbackRuleProperties = original.ShowbackRuleProperties
type ShowbackRulesClient = original.ShowbackRulesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCloudConnectorClient(subscriptionID string) CloudConnectorClient {
	return original.NewCloudConnectorClient(subscriptionID)
}
func NewCloudConnectorClientWithBaseURI(baseURI string, subscriptionID string) CloudConnectorClient {
	return original.NewCloudConnectorClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectorClient(subscriptionID string) ConnectorClient {
	return original.NewConnectorClient(subscriptionID)
}
func NewConnectorClientWithBaseURI(baseURI string, subscriptionID string) ConnectorClient {
	return original.NewConnectorClientWithBaseURI(baseURI, subscriptionID)
}
func NewDimensionsClient(subscriptionID string) DimensionsClient {
	return original.NewDimensionsClient(subscriptionID)
}
func NewDimensionsClientWithBaseURI(baseURI string, subscriptionID string) DimensionsClient {
	return original.NewDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExternalBillingAccountClient(subscriptionID string) ExternalBillingAccountClient {
	return original.NewExternalBillingAccountClient(subscriptionID)
}
func NewExternalBillingAccountClientWithBaseURI(baseURI string, subscriptionID string) ExternalBillingAccountClient {
	return original.NewExternalBillingAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewExternalSubscriptionClient(subscriptionID string) ExternalSubscriptionClient {
	return original.NewExternalSubscriptionClient(subscriptionID)
}
func NewExternalSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) ExternalSubscriptionClient {
	return original.NewExternalSubscriptionClientWithBaseURI(baseURI, subscriptionID)
}
func NewForecastClient(subscriptionID string) ForecastClient {
	return original.NewForecastClient(subscriptionID)
}
func NewForecastClientWithBaseURI(baseURI string, subscriptionID string) ForecastClient {
	return original.NewForecastClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewQueryClient(subscriptionID string) QueryClient {
	return original.NewQueryClient(subscriptionID)
}
func NewQueryClientWithBaseURI(baseURI string, subscriptionID string) QueryClient {
	return original.NewQueryClientWithBaseURI(baseURI, subscriptionID)
}
func NewShowbackRuleClient(subscriptionID string) ShowbackRuleClient {
	return original.NewShowbackRuleClient(subscriptionID)
}
func NewShowbackRuleClientWithBaseURI(baseURI string, subscriptionID string) ShowbackRuleClient {
	return original.NewShowbackRuleClientWithBaseURI(baseURI, subscriptionID)
}
func NewShowbackRulesClient(subscriptionID string) ShowbackRulesClient {
	return original.NewShowbackRulesClient(subscriptionID)
}
func NewShowbackRulesClientWithBaseURI(baseURI string, subscriptionID string) ShowbackRulesClient {
	return original.NewShowbackRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConnectorBillingModelValues() []ConnectorBillingModel {
	return original.PossibleConnectorBillingModelValues()
}
func PossibleConnectorStatusValues() []ConnectorStatus {
	return original.PossibleConnectorStatusValues()
}
func PossibleCostAllocationPolicyValues() []CostAllocationPolicy {
	return original.PossibleCostAllocationPolicyValues()
}
func PossibleDirectionValues() []Direction {
	return original.PossibleDirectionValues()
}
func PossibleFunctionTypeValues() []FunctionType {
	return original.PossibleFunctionTypeValues()
}
func PossibleGranularityTypeValues() []GranularityType {
	return original.PossibleGranularityTypeValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossibleReportConfigColumnTypeValues() []ReportConfigColumnType {
	return original.PossibleReportConfigColumnTypeValues()
}
func PossibleRuleTypeValues() []RuleType {
	return original.PossibleRuleTypeValues()
}
func PossibleShowbackRuleStatusValues() []ShowbackRuleStatus {
	return original.PossibleShowbackRuleStatusValues()
}
func PossibleTimeframeTypeValues() []TimeframeType {
	return original.PossibleTimeframeTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
