// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package securityinsight

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2017-08-01-preview/securityinsight"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationsKind = original.AggregationsKind

const (
	CasesAggregation AggregationsKind = original.CasesAggregation
)

type AlertRuleKind = original.AlertRuleKind

const (
	Scheduled AlertRuleKind = original.Scheduled
)

type AlertSeverity = original.AlertSeverity

const (
	High          AlertSeverity = original.High
	Informational AlertSeverity = original.Informational
	Low           AlertSeverity = original.Low
	Medium        AlertSeverity = original.Medium
)

type CaseSeverity = original.CaseSeverity

const (
	CaseSeverityCritical      CaseSeverity = original.CaseSeverityCritical
	CaseSeverityHigh          CaseSeverity = original.CaseSeverityHigh
	CaseSeverityInformational CaseSeverity = original.CaseSeverityInformational
	CaseSeverityLow           CaseSeverity = original.CaseSeverityLow
	CaseSeverityMedium        CaseSeverity = original.CaseSeverityMedium
)

type CaseStatus = original.CaseStatus

const (
	Closed     CaseStatus = original.Closed
	Draft      CaseStatus = original.Draft
	InProgress CaseStatus = original.InProgress
	Open       CaseStatus = original.Open
)

type CloseReason = original.CloseReason

const (
	Dismissed CloseReason = original.Dismissed
	Other     CloseReason = original.Other
	Resolved  CloseReason = original.Resolved
)

type DataConnectorKind = original.DataConnectorKind

const (
	AzureActiveDirectory      DataConnectorKind = original.AzureActiveDirectory
	AzureSecurityCenter       DataConnectorKind = original.AzureSecurityCenter
	MicrosoftCloudAppSecurity DataConnectorKind = original.MicrosoftCloudAppSecurity
	Office365                 DataConnectorKind = original.Office365
	ThreatIntelligence        DataConnectorKind = original.ThreatIntelligence
)

type DataTypeState = original.DataTypeState

const (
	Disabled DataTypeState = original.Disabled
	Enabled  DataTypeState = original.Enabled
)

type EntityKind = original.EntityKind

const (
	Account EntityKind = original.Account
	File    EntityKind = original.File
	Host    EntityKind = original.Host
)

type Kind = original.Kind

const (
	KindAlertRule Kind = original.KindAlertRule
	KindScheduled Kind = original.KindScheduled
)

type KindBasicAggregations = original.KindBasicAggregations

const (
	KindAggregations KindBasicAggregations = original.KindAggregations
)

type KindBasicDataConnector = original.KindBasicDataConnector

const (
	KindAzureActiveDirectory      KindBasicDataConnector = original.KindAzureActiveDirectory
	KindAzureSecurityCenter       KindBasicDataConnector = original.KindAzureSecurityCenter
	KindDataConnector             KindBasicDataConnector = original.KindDataConnector
	KindMicrosoftCloudAppSecurity KindBasicDataConnector = original.KindMicrosoftCloudAppSecurity
	KindOffice365                 KindBasicDataConnector = original.KindOffice365
	KindThreatIntelligence        KindBasicDataConnector = original.KindThreatIntelligence
)

type KindBasicEntity = original.KindBasicEntity

const (
	KindAccount KindBasicEntity = original.KindAccount
	KindEntity  KindBasicEntity = original.KindEntity
	KindFile    KindBasicEntity = original.KindFile
	KindHost    KindBasicEntity = original.KindHost
)

type KindBasicSettings = original.KindBasicSettings

const (
	KindSettings       KindBasicSettings = original.KindSettings
	KindToggleSettings KindBasicSettings = original.KindToggleSettings
	KindUebaSettings   KindBasicSettings = original.KindUebaSettings
)

type OSFamily = original.OSFamily

const (
	Android OSFamily = original.Android
	IOS     OSFamily = original.IOS
	Linux   OSFamily = original.Linux
	Windows OSFamily = original.Windows
)

type SettingKind = original.SettingKind

const (
	SettingKindToggleSettings SettingKind = original.SettingKindToggleSettings
	SettingKindUebaSettings   SettingKind = original.SettingKindUebaSettings
)

type StatusInMcas = original.StatusInMcas

const (
	StatusInMcasDisabled StatusInMcas = original.StatusInMcasDisabled
	StatusInMcasEnabled  StatusInMcas = original.StatusInMcasEnabled
)

type TriggerOperator = original.TriggerOperator

const (
	Equal       TriggerOperator = original.Equal
	GreaterThan TriggerOperator = original.GreaterThan
	LessThan    TriggerOperator = original.LessThan
	NotEqual    TriggerOperator = original.NotEqual
)

type AADDataConnector = original.AADDataConnector
type AADDataConnectorProperties = original.AADDataConnectorProperties
type ASCDataConnector = original.ASCDataConnector
type ASCDataConnectorProperties = original.ASCDataConnectorProperties
type AccountEntity = original.AccountEntity
type AccountEntityProperties = original.AccountEntityProperties
type Action = original.Action
type ActionProperties = original.ActionProperties
type ActionsClient = original.ActionsClient
type ActionsList = original.ActionsList
type ActionsListIterator = original.ActionsListIterator
type ActionsListPage = original.ActionsListPage
type Aggregations = original.Aggregations
type AggregationsKind1 = original.AggregationsKind1
type AggregationsModel = original.AggregationsModel
type AlertRule = original.AlertRule
type AlertRuleKind1 = original.AlertRuleKind1
type AlertRuleModel = original.AlertRuleModel
type AlertRulesClient = original.AlertRulesClient
type AlertRulesList = original.AlertRulesList
type AlertRulesListIterator = original.AlertRulesListIterator
type AlertRulesListPage = original.AlertRulesListPage
type AlertsDataTypeOfDataConnector = original.AlertsDataTypeOfDataConnector
type AlertsDataTypeOfDataConnectorAlerts = original.AlertsDataTypeOfDataConnectorAlerts
type BaseClient = original.BaseClient
type BasicAggregations = original.BasicAggregations
type BasicAlertRule = original.BasicAlertRule
type BasicDataConnector = original.BasicDataConnector
type BasicEntity = original.BasicEntity
type BasicSettings = original.BasicSettings
type Bookmark = original.Bookmark
type BookmarkList = original.BookmarkList
type BookmarkListIterator = original.BookmarkListIterator
type BookmarkListPage = original.BookmarkListPage
type BookmarkProperties = original.BookmarkProperties
type BookmarksClient = original.BookmarksClient
type Case = original.Case
type CaseList = original.CaseList
type CaseListIterator = original.CaseListIterator
type CaseListPage = original.CaseListPage
type CaseProperties = original.CaseProperties
type CasesAggregationsClient = original.CasesAggregationsClient
type CasesClient = original.CasesClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type DataConnector = original.DataConnector
type DataConnectorDataTypeCommon = original.DataConnectorDataTypeCommon
type DataConnectorKind1 = original.DataConnectorKind1
type DataConnectorList = original.DataConnectorList
type DataConnectorListIterator = original.DataConnectorListIterator
type DataConnectorListPage = original.DataConnectorListPage
type DataConnectorModel = original.DataConnectorModel
type DataConnectorTenantID = original.DataConnectorTenantID
type DataConnectorWithAlertsProperties = original.DataConnectorWithAlertsProperties
type DataConnectorsClient = original.DataConnectorsClient
type EntitiesClient = original.EntitiesClient
type Entity = original.Entity
type EntityKind1 = original.EntityKind1
type EntityList = original.EntityList
type EntityListIterator = original.EntityListIterator
type EntityListPage = original.EntityListPage
type EntityModel = original.EntityModel
type EntityQueriesClient = original.EntityQueriesClient
type EntityQuery = original.EntityQuery
type EntityQueryList = original.EntityQueryList
type EntityQueryListIterator = original.EntityQueryListIterator
type EntityQueryListPage = original.EntityQueryListPage
type FileEntity = original.FileEntity
type FileEntityProperties = original.FileEntityProperties
type HostEntity = original.HostEntity
type HostEntityProperties = original.HostEntityProperties
type MCASDataConnector = original.MCASDataConnector
type MCASDataConnectorProperties = original.MCASDataConnectorProperties
type OfficeConsent = original.OfficeConsent
type OfficeConsentList = original.OfficeConsentList
type OfficeConsentListIterator = original.OfficeConsentListIterator
type OfficeConsentListPage = original.OfficeConsentListPage
type OfficeConsentProperties = original.OfficeConsentProperties
type OfficeConsentsClient = original.OfficeConsentsClient
type OfficeDataConnector = original.OfficeDataConnector
type OfficeDataConnectorDataTypes = original.OfficeDataConnectorDataTypes
type OfficeDataConnectorDataTypesExchange = original.OfficeDataConnectorDataTypesExchange
type OfficeDataConnectorDataTypesSharePoint = original.OfficeDataConnectorDataTypesSharePoint
type OfficeDataConnectorProperties = original.OfficeDataConnectorProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type ProductSettingsClient = original.ProductSettingsClient
type Resource = original.Resource
type ScheduledAlertRule = original.ScheduledAlertRule
type ScheduledAlertRuleProperties = original.ScheduledAlertRuleProperties
type Settings = original.Settings
type SettingsKind = original.SettingsKind
type SettingsModel = original.SettingsModel
type TIDataConnector = original.TIDataConnector
type TIDataConnectorDataTypes = original.TIDataConnectorDataTypes
type TIDataConnectorDataTypesIndicators = original.TIDataConnectorDataTypesIndicators
type TIDataConnectorProperties = original.TIDataConnectorProperties
type ToggleSettings = original.ToggleSettings
type ToggleSettingsProperties = original.ToggleSettingsProperties
type UebaSettings = original.UebaSettings
type UebaSettingsProperties = original.UebaSettingsProperties
type UserInfo = original.UserInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActionsClient(subscriptionID string) ActionsClient {
	return original.NewActionsClient(subscriptionID)
}
func NewActionsClientWithBaseURI(baseURI string, subscriptionID string) ActionsClient {
	return original.NewActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewActionsListIterator(page ActionsListPage) ActionsListIterator {
	return original.NewActionsListIterator(page)
}
func NewActionsListPage(getNextPage func(context.Context, ActionsList) (ActionsList, error)) ActionsListPage {
	return original.NewActionsListPage(getNextPage)
}
func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClient(subscriptionID)
}
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRulesListIterator(page AlertRulesListPage) AlertRulesListIterator {
	return original.NewAlertRulesListIterator(page)
}
func NewAlertRulesListPage(getNextPage func(context.Context, AlertRulesList) (AlertRulesList, error)) AlertRulesListPage {
	return original.NewAlertRulesListPage(getNextPage)
}
func NewBookmarkListIterator(page BookmarkListPage) BookmarkListIterator {
	return original.NewBookmarkListIterator(page)
}
func NewBookmarkListPage(getNextPage func(context.Context, BookmarkList) (BookmarkList, error)) BookmarkListPage {
	return original.NewBookmarkListPage(getNextPage)
}
func NewBookmarksClient(subscriptionID string) BookmarksClient {
	return original.NewBookmarksClient(subscriptionID)
}
func NewBookmarksClientWithBaseURI(baseURI string, subscriptionID string) BookmarksClient {
	return original.NewBookmarksClientWithBaseURI(baseURI, subscriptionID)
}
func NewCaseListIterator(page CaseListPage) CaseListIterator {
	return original.NewCaseListIterator(page)
}
func NewCaseListPage(getNextPage func(context.Context, CaseList) (CaseList, error)) CaseListPage {
	return original.NewCaseListPage(getNextPage)
}
func NewCasesAggregationsClient(subscriptionID string) CasesAggregationsClient {
	return original.NewCasesAggregationsClient(subscriptionID)
}
func NewCasesAggregationsClientWithBaseURI(baseURI string, subscriptionID string) CasesAggregationsClient {
	return original.NewCasesAggregationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCasesClient(subscriptionID string) CasesClient {
	return original.NewCasesClient(subscriptionID)
}
func NewCasesClientWithBaseURI(baseURI string, subscriptionID string) CasesClient {
	return original.NewCasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataConnectorListIterator(page DataConnectorListPage) DataConnectorListIterator {
	return original.NewDataConnectorListIterator(page)
}
func NewDataConnectorListPage(getNextPage func(context.Context, DataConnectorList) (DataConnectorList, error)) DataConnectorListPage {
	return original.NewDataConnectorListPage(getNextPage)
}
func NewDataConnectorsClient(subscriptionID string) DataConnectorsClient {
	return original.NewDataConnectorsClient(subscriptionID)
}
func NewDataConnectorsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectorsClient {
	return original.NewDataConnectorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEntitiesClient(subscriptionID string) EntitiesClient {
	return original.NewEntitiesClient(subscriptionID)
}
func NewEntitiesClientWithBaseURI(baseURI string, subscriptionID string) EntitiesClient {
	return original.NewEntitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEntityListIterator(page EntityListPage) EntityListIterator {
	return original.NewEntityListIterator(page)
}
func NewEntityListPage(getNextPage func(context.Context, EntityList) (EntityList, error)) EntityListPage {
	return original.NewEntityListPage(getNextPage)
}
func NewEntityQueriesClient(subscriptionID string) EntityQueriesClient {
	return original.NewEntityQueriesClient(subscriptionID)
}
func NewEntityQueriesClientWithBaseURI(baseURI string, subscriptionID string) EntityQueriesClient {
	return original.NewEntityQueriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEntityQueryListIterator(page EntityQueryListPage) EntityQueryListIterator {
	return original.NewEntityQueryListIterator(page)
}
func NewEntityQueryListPage(getNextPage func(context.Context, EntityQueryList) (EntityQueryList, error)) EntityQueryListPage {
	return original.NewEntityQueryListPage(getNextPage)
}
func NewOfficeConsentListIterator(page OfficeConsentListPage) OfficeConsentListIterator {
	return original.NewOfficeConsentListIterator(page)
}
func NewOfficeConsentListPage(getNextPage func(context.Context, OfficeConsentList) (OfficeConsentList, error)) OfficeConsentListPage {
	return original.NewOfficeConsentListPage(getNextPage)
}
func NewOfficeConsentsClient(subscriptionID string) OfficeConsentsClient {
	return original.NewOfficeConsentsClient(subscriptionID)
}
func NewOfficeConsentsClientWithBaseURI(baseURI string, subscriptionID string) OfficeConsentsClient {
	return original.NewOfficeConsentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(getNextPage)
}
func NewProductSettingsClient(subscriptionID string) ProductSettingsClient {
	return original.NewProductSettingsClient(subscriptionID)
}
func NewProductSettingsClientWithBaseURI(baseURI string, subscriptionID string) ProductSettingsClient {
	return original.NewProductSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAggregationsKindValues() []AggregationsKind {
	return original.PossibleAggregationsKindValues()
}
func PossibleAlertRuleKindValues() []AlertRuleKind {
	return original.PossibleAlertRuleKindValues()
}
func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}
func PossibleCaseSeverityValues() []CaseSeverity {
	return original.PossibleCaseSeverityValues()
}
func PossibleCaseStatusValues() []CaseStatus {
	return original.PossibleCaseStatusValues()
}
func PossibleCloseReasonValues() []CloseReason {
	return original.PossibleCloseReasonValues()
}
func PossibleDataConnectorKindValues() []DataConnectorKind {
	return original.PossibleDataConnectorKindValues()
}
func PossibleDataTypeStateValues() []DataTypeState {
	return original.PossibleDataTypeStateValues()
}
func PossibleEntityKindValues() []EntityKind {
	return original.PossibleEntityKindValues()
}
func PossibleKindBasicAggregationsValues() []KindBasicAggregations {
	return original.PossibleKindBasicAggregationsValues()
}
func PossibleKindBasicDataConnectorValues() []KindBasicDataConnector {
	return original.PossibleKindBasicDataConnectorValues()
}
func PossibleKindBasicEntityValues() []KindBasicEntity {
	return original.PossibleKindBasicEntityValues()
}
func PossibleKindBasicSettingsValues() []KindBasicSettings {
	return original.PossibleKindBasicSettingsValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleOSFamilyValues() []OSFamily {
	return original.PossibleOSFamilyValues()
}
func PossibleSettingKindValues() []SettingKind {
	return original.PossibleSettingKindValues()
}
func PossibleStatusInMcasValues() []StatusInMcas {
	return original.PossibleStatusInMcasValues()
}
func PossibleTriggerOperatorValues() []TriggerOperator {
	return original.PossibleTriggerOperatorValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
