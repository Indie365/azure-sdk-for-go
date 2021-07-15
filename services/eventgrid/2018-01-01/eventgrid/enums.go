package eventgrid

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AppAction enumerates the values for app action.
type AppAction string

const (
	// ChangedAppSettings There was an operation to change app setting on the web app.
	ChangedAppSettings AppAction = "ChangedAppSettings"
	// Completed The job has completed.
	Completed AppAction = "Completed"
	// Failed The job has failed to complete.
	Failed AppAction = "Failed"
	// Restarted Web app was restarted.
	Restarted AppAction = "Restarted"
	// Started The job has started.
	Started AppAction = "Started"
	// Stopped Web app was stopped.
	Stopped AppAction = "Stopped"
)

// PossibleAppActionValues returns an array of possible values for the AppAction const type.
func PossibleAppActionValues() []AppAction {
	return []AppAction{ChangedAppSettings, Completed, Failed, Restarted, Started, Stopped}
}

// AppServicePlanAction enumerates the values for app service plan action.
type AppServicePlanAction string

const (
	// Updated App Service plan is being updated.
	Updated AppServicePlanAction = "Updated"
)

// PossibleAppServicePlanActionValues returns an array of possible values for the AppServicePlanAction const type.
func PossibleAppServicePlanActionValues() []AppServicePlanAction {
	return []AppServicePlanAction{Updated}
}

// AsyncStatus enumerates the values for async status.
type AsyncStatus string

const (
	// AsyncStatusCompleted Async operation has completed.
	AsyncStatusCompleted AsyncStatus = "Completed"
	// AsyncStatusFailed Async operation failed to complete.
	AsyncStatusFailed AsyncStatus = "Failed"
	// AsyncStatusStarted Async operation has started.
	AsyncStatusStarted AsyncStatus = "Started"
)

// PossibleAsyncStatusValues returns an array of possible values for the AsyncStatus const type.
func PossibleAsyncStatusValues() []AsyncStatus {
	return []AsyncStatus{AsyncStatusCompleted, AsyncStatusFailed, AsyncStatusStarted}
}

// CommunicationCloudEnvironmentModel enumerates the values for communication cloud environment model.
type CommunicationCloudEnvironmentModel string

const (
	// Dod ...
	Dod CommunicationCloudEnvironmentModel = "dod"
	// Gcch ...
	Gcch CommunicationCloudEnvironmentModel = "gcch"
	// Public ...
	Public CommunicationCloudEnvironmentModel = "public"
)

// PossibleCommunicationCloudEnvironmentModelValues returns an array of possible values for the CommunicationCloudEnvironmentModel const type.
func PossibleCommunicationCloudEnvironmentModelValues() []CommunicationCloudEnvironmentModel {
	return []CommunicationCloudEnvironmentModel{Dod, Gcch, Public}
}

// MediaJobErrorCategory enumerates the values for media job error category.
type MediaJobErrorCategory string

const (
	// Configuration The error is configuration related.
	Configuration MediaJobErrorCategory = "Configuration"
	// Content The error is related to data in the input files.
	Content MediaJobErrorCategory = "Content"
	// Download The error is download related.
	Download MediaJobErrorCategory = "Download"
	// Service The error is service related.
	Service MediaJobErrorCategory = "Service"
	// Upload The error is upload related.
	Upload MediaJobErrorCategory = "Upload"
)

// PossibleMediaJobErrorCategoryValues returns an array of possible values for the MediaJobErrorCategory const type.
func PossibleMediaJobErrorCategoryValues() []MediaJobErrorCategory {
	return []MediaJobErrorCategory{Configuration, Content, Download, Service, Upload}
}

// MediaJobErrorCode enumerates the values for media job error code.
type MediaJobErrorCode string

const (
	// ConfigurationUnsupported There was a problem with the combination of input files and the configuration
	// settings applied, fix the configuration settings and retry with the same input, or change input to match
	// the configuration.
	ConfigurationUnsupported MediaJobErrorCode = "ConfigurationUnsupported"
	// ContentMalformed There was a problem with the input content (for example: zero byte files, or
	// corrupt/non-decodable files), check the input files.
	ContentMalformed MediaJobErrorCode = "ContentMalformed"
	// ContentUnsupported There was a problem with the format of the input (not valid media file, or an
	// unsupported file/codec), check the validity of the input files.
	ContentUnsupported MediaJobErrorCode = "ContentUnsupported"
	// DownloadNotAccessible While trying to download the input files, the files were not accessible, please
	// check the availability of the source.
	DownloadNotAccessible MediaJobErrorCode = "DownloadNotAccessible"
	// DownloadTransientError While trying to download the input files, there was an issue during transfer
	// (storage service, network errors), see details and check your source.
	DownloadTransientError MediaJobErrorCode = "DownloadTransientError"
	// ServiceError Fatal service error, please contact support.
	ServiceError MediaJobErrorCode = "ServiceError"
	// ServiceTransientError Transient error, please retry, if retry is unsuccessful, please contact support.
	ServiceTransientError MediaJobErrorCode = "ServiceTransientError"
	// UploadNotAccessible While trying to upload the output files, the destination was not reachable, please
	// check the availability of the destination.
	UploadNotAccessible MediaJobErrorCode = "UploadNotAccessible"
	// UploadTransientError While trying to upload the output files, there was an issue during transfer
	// (storage service, network errors), see details and check your destination.
	UploadTransientError MediaJobErrorCode = "UploadTransientError"
)

// PossibleMediaJobErrorCodeValues returns an array of possible values for the MediaJobErrorCode const type.
func PossibleMediaJobErrorCodeValues() []MediaJobErrorCode {
	return []MediaJobErrorCode{ConfigurationUnsupported, ContentMalformed, ContentUnsupported, DownloadNotAccessible, DownloadTransientError, ServiceError, ServiceTransientError, UploadNotAccessible, UploadTransientError}
}

// MediaJobRetry enumerates the values for media job retry.
type MediaJobRetry string

const (
	// DoNotRetry Issue needs to be investigated and then the job resubmitted with corrections or retried once
	// the underlying issue has been corrected.
	DoNotRetry MediaJobRetry = "DoNotRetry"
	// MayRetry Issue may be resolved after waiting for a period of time and resubmitting the same Job.
	MayRetry MediaJobRetry = "MayRetry"
)

// PossibleMediaJobRetryValues returns an array of possible values for the MediaJobRetry const type.
func PossibleMediaJobRetryValues() []MediaJobRetry {
	return []MediaJobRetry{DoNotRetry, MayRetry}
}

// MediaJobState enumerates the values for media job state.
type MediaJobState string

const (
	// Canceled The job was canceled. This is a final state for the job.
	Canceled MediaJobState = "Canceled"
	// Canceling The job is in the process of being canceled. This is a transient state for the job.
	Canceling MediaJobState = "Canceling"
	// Error The job has encountered an error. This is a final state for the job.
	Error MediaJobState = "Error"
	// Finished The job is finished. This is a final state for the job.
	Finished MediaJobState = "Finished"
	// Processing The job is processing. This is a transient state for the job.
	Processing MediaJobState = "Processing"
	// Queued The job is in a queued state, waiting for resources to become available. This is a transient
	// state.
	Queued MediaJobState = "Queued"
	// Scheduled The job is being scheduled to run on an available resource. This is a transient state, between
	// queued and processing states.
	Scheduled MediaJobState = "Scheduled"
)

// PossibleMediaJobStateValues returns an array of possible values for the MediaJobState const type.
func PossibleMediaJobStateValues() []MediaJobState {
	return []MediaJobState{Canceled, Canceling, Error, Finished, Processing, Queued, Scheduled}
}

// OdataType enumerates the values for odata type.
type OdataType string

const (
	// OdataTypeMediaJobOutput ...
	OdataTypeMediaJobOutput OdataType = "MediaJobOutput"
	// OdataTypeMicrosoftMediaJobOutputAsset ...
	OdataTypeMicrosoftMediaJobOutputAsset OdataType = "#Microsoft.Media.JobOutputAsset"
)

// PossibleOdataTypeValues returns an array of possible values for the OdataType const type.
func PossibleOdataTypeValues() []OdataType {
	return []OdataType{OdataTypeMediaJobOutput, OdataTypeMicrosoftMediaJobOutputAsset}
}

// StampKind enumerates the values for stamp kind.
type StampKind string

const (
	// StampKindAseV1 App Service Plan is running on an App Service Environment V1.
	StampKindAseV1 StampKind = "AseV1"
	// StampKindAseV2 App Service Plan is running on an App Service Environment V2.
	StampKindAseV2 StampKind = "AseV2"
	// StampKindPublic App Service Plan is running on a public stamp.
	StampKindPublic StampKind = "Public"
)

// PossibleStampKindValues returns an array of possible values for the StampKind const type.
func PossibleStampKindValues() []StampKind {
	return []StampKind{StampKindAseV1, StampKindAseV2, StampKindPublic}
}
