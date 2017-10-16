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

package mobileengagement

import original "github.com/Azure/azure-sdk-for-go/services/mobileengagement/mgmt/2014-12-01/mobileengagement"

type AppCollectionsClient = original.AppCollectionsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type ImportTasksClient = original.ImportTasksClient
type AudienceOperators = original.AudienceOperators

const (
	EQ	AudienceOperators	= original.EQ
	GE	AudienceOperators	= original.GE
	GT	AudienceOperators	= original.GT
	LE	AudienceOperators	= original.LE
	LT	AudienceOperators	= original.LT
)

type CampaignFeedbacks = original.CampaignFeedbacks

const (
	Actioned	CampaignFeedbacks	= original.Actioned
	Exited		CampaignFeedbacks	= original.Exited
	Pushed		CampaignFeedbacks	= original.Pushed
	Replied		CampaignFeedbacks	= original.Replied
)

type CampaignKinds = original.CampaignKinds

const (
	Announcements	CampaignKinds	= original.Announcements
	DataPushes	CampaignKinds	= original.DataPushes
	NativePushes	CampaignKinds	= original.NativePushes
	Polls		CampaignKinds	= original.Polls
)

type CampaignStates = original.CampaignStates

const (
	Draft		CampaignStates	= original.Draft
	Finished	CampaignStates	= original.Finished
	InProgress	CampaignStates	= original.InProgress
	Queued		CampaignStates	= original.Queued
	Scheduled	CampaignStates	= original.Scheduled
)

type CampaignType = original.CampaignType

const (
	Announcement	CampaignType	= original.Announcement
	DataPush	CampaignType	= original.DataPush
	NativePush	CampaignType	= original.NativePush
	Poll		CampaignType	= original.Poll
)

type CampaignTypes = original.CampaignTypes

const (
	OnlyNotif	CampaignTypes	= original.OnlyNotif
	Textbase64	CampaignTypes	= original.Textbase64
	Texthtml	CampaignTypes	= original.Texthtml
	Textplain	CampaignTypes	= original.Textplain
)

type DeliveryTimes = original.DeliveryTimes

const (
	Any		DeliveryTimes	= original.Any
	Background	DeliveryTimes	= original.Background
	Session		DeliveryTimes	= original.Session
)

type ExportFormat = original.ExportFormat

const (
	CsvBlob		ExportFormat	= original.CsvBlob
	JSONBlob	ExportFormat	= original.JSONBlob
)

type ExportState = original.ExportState

const (
	ExportStateFailed	ExportState	= original.ExportStateFailed
	ExportStateQueued	ExportState	= original.ExportStateQueued
	ExportStateStarted	ExportState	= original.ExportStateStarted
	ExportStateSucceeded	ExportState	= original.ExportStateSucceeded
)

type ExportType = original.ExportType

const (
	ExportTypeActivity	ExportType	= original.ExportTypeActivity
	ExportTypeCrash		ExportType	= original.ExportTypeCrash
	ExportTypeError		ExportType	= original.ExportTypeError
	ExportTypeEvent		ExportType	= original.ExportTypeEvent
	ExportTypeJob		ExportType	= original.ExportTypeJob
	ExportTypePush		ExportType	= original.ExportTypePush
	ExportTypeSession	ExportType	= original.ExportTypeSession
	ExportTypeTag		ExportType	= original.ExportTypeTag
	ExportTypeToken		ExportType	= original.ExportTypeToken
)

type JobStates = original.JobStates

const (
	JobStatesFailed		JobStates	= original.JobStatesFailed
	JobStatesQueued		JobStates	= original.JobStatesQueued
	JobStatesStarted	JobStates	= original.JobStatesStarted
	JobStatesSucceeded	JobStates	= original.JobStatesSucceeded
)

type NotificationTypes = original.NotificationTypes

const (
	Popup	NotificationTypes	= original.Popup
	System	NotificationTypes	= original.System
)

type ProvisioningStates = original.ProvisioningStates

const (
	Creating	ProvisioningStates	= original.Creating
	Succeeded	ProvisioningStates	= original.Succeeded
)

type PushModes = original.PushModes

const (
	Manual		PushModes	= original.Manual
	OneShot		PushModes	= original.OneShot
	RealTime	PushModes	= original.RealTime
)

type Type = original.Type

const (
	TypeAnnouncementFeedback	Type	= original.TypeAnnouncementFeedback
	TypeApplicationVersion		Type	= original.TypeApplicationVersion
	TypeBooleanTag			Type	= original.TypeBooleanTag
	TypeCarrierCountry		Type	= original.TypeCarrierCountry
	TypeCarrierName			Type	= original.TypeCarrierName
	TypeDatapushFeedback		Type	= original.TypeDatapushFeedback
	TypeDateTag			Type	= original.TypeDateTag
	TypeDeviceManufacturer		Type	= original.TypeDeviceManufacturer
	TypeDeviceModel			Type	= original.TypeDeviceModel
	TypeFirmwareVersion		Type	= original.TypeFirmwareVersion
	TypeGeoFencing			Type	= original.TypeGeoFencing
	TypeIntegerTag			Type	= original.TypeIntegerTag
	TypeLanguage			Type	= original.TypeLanguage
	TypeLocation			Type	= original.TypeLocation
	TypeNetworkType			Type	= original.TypeNetworkType
	TypePollAnswerFeedback		Type	= original.TypePollAnswerFeedback
	TypePollFeedback		Type	= original.TypePollFeedback
	TypeScreenSize			Type	= original.TypeScreenSize
	TypeSegment			Type	= original.TypeSegment
	TypeStringTag			Type	= original.TypeStringTag
)

type TypeFilter = original.TypeFilter

const (
	TypeAppInfo		TypeFilter	= original.TypeAppInfo
	TypeEngageActiveUsers	TypeFilter	= original.TypeEngageActiveUsers
	TypeEngageIdleUsers	TypeFilter	= original.TypeEngageIdleUsers
	TypeEngageNewUsers	TypeFilter	= original.TypeEngageNewUsers
	TypeEngageOldUsers	TypeFilter	= original.TypeEngageOldUsers
	TypeEngageSubset	TypeFilter	= original.TypeEngageSubset
	TypeNativePushEnabled	TypeFilter	= original.TypeNativePushEnabled
	TypePushQuota		TypeFilter	= original.TypePushQuota
)

type AnnouncementFeedbackCriterion = original.AnnouncementFeedbackCriterion
type APIError = original.APIError
type APIErrorError = original.APIErrorError
type App = original.App
type AppCollection = original.AppCollection
type AppCollectionListResult = original.AppCollectionListResult
type AppCollectionNameAvailability = original.AppCollectionNameAvailability
type AppCollectionProperties = original.AppCollectionProperties
type AppInfoFilter = original.AppInfoFilter
type ApplicationVersionCriterion = original.ApplicationVersionCriterion
type AppListResult = original.AppListResult
type AppProperties = original.AppProperties
type BooleanTagCriterion = original.BooleanTagCriterion
type Campaign = original.Campaign
type CampaignAudience = original.CampaignAudience
type CampaignListResult = original.CampaignListResult
type CampaignLocalization = original.CampaignLocalization
type CampaignPushParameters = original.CampaignPushParameters
type CampaignPushResult = original.CampaignPushResult
type CampaignResult = original.CampaignResult
type CampaignsListResult = original.CampaignsListResult
type CampaignState = original.CampaignState
type CampaignStateResult = original.CampaignStateResult
type CampaignStatisticsResult = original.CampaignStatisticsResult
type CampaignTestNewParameters = original.CampaignTestNewParameters
type CampaignTestSavedParameters = original.CampaignTestSavedParameters
type CarrierCountryCriterion = original.CarrierCountryCriterion
type CarrierNameCriterion = original.CarrierNameCriterion
type Criterion = original.Criterion
type DatapushFeedbackCriterion = original.DatapushFeedbackCriterion
type DateRangeExportTaskParameter = original.DateRangeExportTaskParameter
type DateTagCriterion = original.DateTagCriterion
type Device = original.Device
type DeviceInfo = original.DeviceInfo
type DeviceLocation = original.DeviceLocation
type DeviceManufacturerCriterion = original.DeviceManufacturerCriterion
type DeviceMeta = original.DeviceMeta
type DeviceModelCriterion = original.DeviceModelCriterion
type DeviceQueryResult = original.DeviceQueryResult
type DevicesQueryResult = original.DevicesQueryResult
type DeviceTagsParameters = original.DeviceTagsParameters
type DeviceTagsResult = original.DeviceTagsResult
type EngageActiveUsersFilter = original.EngageActiveUsersFilter
type EngageIdleUsersFilter = original.EngageIdleUsersFilter
type EngageNewUsersFilter = original.EngageNewUsersFilter
type EngageOldUsersFilter = original.EngageOldUsersFilter
type EngageSubsetFilter = original.EngageSubsetFilter
type ExportOptions = original.ExportOptions
type ExportTaskListResult = original.ExportTaskListResult
type ExportTaskParameter = original.ExportTaskParameter
type ExportTaskResult = original.ExportTaskResult
type FeedbackByCampaignParameter = original.FeedbackByCampaignParameter
type FeedbackByDateRangeParameter = original.FeedbackByDateRangeParameter
type Filter = original.Filter
type FirmwareVersionCriterion = original.FirmwareVersionCriterion
type GeoFencingCriterion = original.GeoFencingCriterion
type ImportTask = original.ImportTask
type ImportTaskListResult = original.ImportTaskListResult
type ImportTaskResult = original.ImportTaskResult
type IntegerTagCriterion = original.IntegerTagCriterion
type LanguageCriterion = original.LanguageCriterion
type LocationCriterion = original.LocationCriterion
type NativePushEnabledFilter = original.NativePushEnabledFilter
type NetworkTypeCriterion = original.NetworkTypeCriterion
type NotificationOptions = original.NotificationOptions
type PollAnswerFeedbackCriterion = original.PollAnswerFeedbackCriterion
type PollFeedbackCriterion = original.PollFeedbackCriterion
type PollQuestion = original.PollQuestion
type PollQuestionChoice = original.PollQuestionChoice
type PollQuestionChoiceLocalization = original.PollQuestionChoiceLocalization
type PollQuestionLocalization = original.PollQuestionLocalization
type PushQuotaFilter = original.PushQuotaFilter
type Resource = original.Resource
type ScreenSizeCriterion = original.ScreenSizeCriterion
type SegmentCriterion = original.SegmentCriterion
type StringTagCriterion = original.StringTagCriterion
type SupportedPlatformsListResult = original.SupportedPlatformsListResult
type SupportedPlatformsClient = original.SupportedPlatformsClient
type AppsClient = original.AppsClient
type CampaignsClient = original.CampaignsClient
type DevicesClient = original.DevicesClient
type ExportTasksClient = original.ExportTasksClient

func NewAppsClient(subscriptionID string) AppsClient {
	return original.NewAppsClient(subscriptionID)
}
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return original.NewAppsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCampaignsClient(subscriptionID string) CampaignsClient {
	return original.NewCampaignsClient(subscriptionID)
}
func NewCampaignsClientWithBaseURI(baseURI string, subscriptionID string) CampaignsClient {
	return original.NewCampaignsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDevicesClient(subscriptionID string) DevicesClient {
	return original.NewDevicesClient(subscriptionID)
}
func NewDevicesClientWithBaseURI(baseURI string, subscriptionID string) DevicesClient {
	return original.NewDevicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewExportTasksClient(subscriptionID string) ExportTasksClient {
	return original.NewExportTasksClient(subscriptionID)
}
func NewExportTasksClientWithBaseURI(baseURI string, subscriptionID string) ExportTasksClient {
	return original.NewExportTasksClientWithBaseURI(baseURI, subscriptionID)
}
func NewSupportedPlatformsClient(subscriptionID string) SupportedPlatformsClient {
	return original.NewSupportedPlatformsClient(subscriptionID)
}
func NewSupportedPlatformsClientWithBaseURI(baseURI string, subscriptionID string) SupportedPlatformsClient {
	return original.NewSupportedPlatformsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewAppCollectionsClient(subscriptionID string) AppCollectionsClient {
	return original.NewAppCollectionsClient(subscriptionID)
}
func NewAppCollectionsClientWithBaseURI(baseURI string, subscriptionID string) AppCollectionsClient {
	return original.NewAppCollectionsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewImportTasksClient(subscriptionID string) ImportTasksClient {
	return original.NewImportTasksClient(subscriptionID)
}
func NewImportTasksClientWithBaseURI(baseURI string, subscriptionID string) ImportTasksClient {
	return original.NewImportTasksClientWithBaseURI(baseURI, subscriptionID)
}
