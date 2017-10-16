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

package automation

import original "github.com/Azure/azure-sdk-for-go/services/automation/mgmt/2015-10-31/automation"

type RunbookDraftClient = original.RunbookDraftClient
type ScheduleClient = original.ScheduleClient
type VariableClient = original.VariableClient
type ActivityClient = original.ActivityClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type ConnectionClient = original.ConnectionClient
type AgentRegistrationInformationClient = original.AgentRegistrationInformationClient
type DscNodeClient = original.DscNodeClient
type DscNodeConfigurationClient = original.DscNodeConfigurationClient
type JobClient = original.JobClient
type NodeReportsClient = original.NodeReportsClient
type AccountClient = original.AccountClient
type AccountState = original.AccountState

const (
	Ok		AccountState	= original.Ok
	Suspended	AccountState	= original.Suspended
	Unavailable	AccountState	= original.Unavailable
)

type AgentRegistrationKeyName = original.AgentRegistrationKeyName

const (
	Primary		AgentRegistrationKeyName	= original.Primary
	Secondary	AgentRegistrationKeyName	= original.Secondary
)

type ContentSourceType = original.ContentSourceType

const (
	EmbeddedContent	ContentSourceType	= original.EmbeddedContent
	URI		ContentSourceType	= original.URI
)

type DscConfigurationProvisioningState = original.DscConfigurationProvisioningState

const (
	Succeeded DscConfigurationProvisioningState = original.Succeeded
)

type DscConfigurationState = original.DscConfigurationState

const (
	DscConfigurationStateEdit	DscConfigurationState	= original.DscConfigurationStateEdit
	DscConfigurationStateNew	DscConfigurationState	= original.DscConfigurationStateNew
	DscConfigurationStatePublished	DscConfigurationState	= original.DscConfigurationStatePublished
)

type HTTPStatusCode = original.HTTPStatusCode

const (
	Accepted			HTTPStatusCode	= original.Accepted
	Ambiguous			HTTPStatusCode	= original.Ambiguous
	BadGateway			HTTPStatusCode	= original.BadGateway
	BadRequest			HTTPStatusCode	= original.BadRequest
	Conflict			HTTPStatusCode	= original.Conflict
	Continue			HTTPStatusCode	= original.Continue
	Created				HTTPStatusCode	= original.Created
	ExpectationFailed		HTTPStatusCode	= original.ExpectationFailed
	Forbidden			HTTPStatusCode	= original.Forbidden
	Found				HTTPStatusCode	= original.Found
	GatewayTimeout			HTTPStatusCode	= original.GatewayTimeout
	Gone				HTTPStatusCode	= original.Gone
	HTTPVersionNotSupported		HTTPStatusCode	= original.HTTPVersionNotSupported
	InternalServerError		HTTPStatusCode	= original.InternalServerError
	LengthRequired			HTTPStatusCode	= original.LengthRequired
	MethodNotAllowed		HTTPStatusCode	= original.MethodNotAllowed
	Moved				HTTPStatusCode	= original.Moved
	MovedPermanently		HTTPStatusCode	= original.MovedPermanently
	MultipleChoices			HTTPStatusCode	= original.MultipleChoices
	NoContent			HTTPStatusCode	= original.NoContent
	NonAuthoritativeInformation	HTTPStatusCode	= original.NonAuthoritativeInformation
	NotAcceptable			HTTPStatusCode	= original.NotAcceptable
	NotFound			HTTPStatusCode	= original.NotFound
	NotImplemented			HTTPStatusCode	= original.NotImplemented
	NotModified			HTTPStatusCode	= original.NotModified
	OK				HTTPStatusCode	= original.OK
	PartialContent			HTTPStatusCode	= original.PartialContent
	PaymentRequired			HTTPStatusCode	= original.PaymentRequired
	PreconditionFailed		HTTPStatusCode	= original.PreconditionFailed
	ProxyAuthenticationRequired	HTTPStatusCode	= original.ProxyAuthenticationRequired
	Redirect			HTTPStatusCode	= original.Redirect
	RedirectKeepVerb		HTTPStatusCode	= original.RedirectKeepVerb
	RedirectMethod			HTTPStatusCode	= original.RedirectMethod
	RequestedRangeNotSatisfiable	HTTPStatusCode	= original.RequestedRangeNotSatisfiable
	RequestEntityTooLarge		HTTPStatusCode	= original.RequestEntityTooLarge
	RequestTimeout			HTTPStatusCode	= original.RequestTimeout
	RequestURITooLong		HTTPStatusCode	= original.RequestURITooLong
	ResetContent			HTTPStatusCode	= original.ResetContent
	SeeOther			HTTPStatusCode	= original.SeeOther
	ServiceUnavailable		HTTPStatusCode	= original.ServiceUnavailable
	SwitchingProtocols		HTTPStatusCode	= original.SwitchingProtocols
	TemporaryRedirect		HTTPStatusCode	= original.TemporaryRedirect
	Unauthorized			HTTPStatusCode	= original.Unauthorized
	UnsupportedMediaType		HTTPStatusCode	= original.UnsupportedMediaType
	Unused				HTTPStatusCode	= original.Unused
	UpgradeRequired			HTTPStatusCode	= original.UpgradeRequired
	UseProxy			HTTPStatusCode	= original.UseProxy
)

type JobStatus = original.JobStatus

const (
	JobStatusActivating	JobStatus	= original.JobStatusActivating
	JobStatusBlocked	JobStatus	= original.JobStatusBlocked
	JobStatusCompleted	JobStatus	= original.JobStatusCompleted
	JobStatusDisconnected	JobStatus	= original.JobStatusDisconnected
	JobStatusFailed		JobStatus	= original.JobStatusFailed
	JobStatusNew		JobStatus	= original.JobStatusNew
	JobStatusRemoving	JobStatus	= original.JobStatusRemoving
	JobStatusResuming	JobStatus	= original.JobStatusResuming
	JobStatusRunning	JobStatus	= original.JobStatusRunning
	JobStatusStopped	JobStatus	= original.JobStatusStopped
	JobStatusStopping	JobStatus	= original.JobStatusStopping
	JobStatusSuspended	JobStatus	= original.JobStatusSuspended
	JobStatusSuspending	JobStatus	= original.JobStatusSuspending
)

type JobStreamType = original.JobStreamType

const (
	Any		JobStreamType	= original.Any
	Debug		JobStreamType	= original.Debug
	Error		JobStreamType	= original.Error
	Output		JobStreamType	= original.Output
	Progress	JobStreamType	= original.Progress
	Verbose		JobStreamType	= original.Verbose
	Warning		JobStreamType	= original.Warning
)

type ModuleProvisioningState = original.ModuleProvisioningState

const (
	ModuleProvisioningStateActivitiesStored			ModuleProvisioningState	= original.ModuleProvisioningStateActivitiesStored
	ModuleProvisioningStateCancelled			ModuleProvisioningState	= original.ModuleProvisioningStateCancelled
	ModuleProvisioningStateConnectionTypeImported		ModuleProvisioningState	= original.ModuleProvisioningStateConnectionTypeImported
	ModuleProvisioningStateContentDownloaded		ModuleProvisioningState	= original.ModuleProvisioningStateContentDownloaded
	ModuleProvisioningStateContentRetrieved			ModuleProvisioningState	= original.ModuleProvisioningStateContentRetrieved
	ModuleProvisioningStateContentStored			ModuleProvisioningState	= original.ModuleProvisioningStateContentStored
	ModuleProvisioningStateContentValidated			ModuleProvisioningState	= original.ModuleProvisioningStateContentValidated
	ModuleProvisioningStateCreated				ModuleProvisioningState	= original.ModuleProvisioningStateCreated
	ModuleProvisioningStateCreating				ModuleProvisioningState	= original.ModuleProvisioningStateCreating
	ModuleProvisioningStateFailed				ModuleProvisioningState	= original.ModuleProvisioningStateFailed
	ModuleProvisioningStateModuleDataStored			ModuleProvisioningState	= original.ModuleProvisioningStateModuleDataStored
	ModuleProvisioningStateModuleImportRunbookComplete	ModuleProvisioningState	= original.ModuleProvisioningStateModuleImportRunbookComplete
	ModuleProvisioningStateRunningImportModuleRunbook	ModuleProvisioningState	= original.ModuleProvisioningStateRunningImportModuleRunbook
	ModuleProvisioningStateStartingImportModuleRunbook	ModuleProvisioningState	= original.ModuleProvisioningStateStartingImportModuleRunbook
	ModuleProvisioningStateSucceeded			ModuleProvisioningState	= original.ModuleProvisioningStateSucceeded
	ModuleProvisioningStateUpdating				ModuleProvisioningState	= original.ModuleProvisioningStateUpdating
)

type RunbookProvisioningState = original.RunbookProvisioningState

const (
	RunbookProvisioningStateSucceeded RunbookProvisioningState = original.RunbookProvisioningStateSucceeded
)

type RunbookState = original.RunbookState

const (
	RunbookStateEdit	RunbookState	= original.RunbookStateEdit
	RunbookStateNew		RunbookState	= original.RunbookStateNew
	RunbookStatePublished	RunbookState	= original.RunbookStatePublished
)

type RunbookTypeEnum = original.RunbookTypeEnum

const (
	Graph			RunbookTypeEnum	= original.Graph
	GraphPowerShell		RunbookTypeEnum	= original.GraphPowerShell
	GraphPowerShellWorkflow	RunbookTypeEnum	= original.GraphPowerShellWorkflow
	PowerShell		RunbookTypeEnum	= original.PowerShell
	PowerShellWorkflow	RunbookTypeEnum	= original.PowerShellWorkflow
	Script			RunbookTypeEnum	= original.Script
)

type ScheduleDay = original.ScheduleDay

const (
	Friday		ScheduleDay	= original.Friday
	Monday		ScheduleDay	= original.Monday
	Saturday	ScheduleDay	= original.Saturday
	Sunday		ScheduleDay	= original.Sunday
	Thursday	ScheduleDay	= original.Thursday
	Tuesday		ScheduleDay	= original.Tuesday
	Wednesday	ScheduleDay	= original.Wednesday
)

type ScheduleFrequency = original.ScheduleFrequency

const (
	Day	ScheduleFrequency	= original.Day
	Hour	ScheduleFrequency	= original.Hour
	Month	ScheduleFrequency	= original.Month
	OneTime	ScheduleFrequency	= original.OneTime
	Week	ScheduleFrequency	= original.Week
)

type SkuNameEnum = original.SkuNameEnum

const (
	Basic	SkuNameEnum	= original.Basic
	Free	SkuNameEnum	= original.Free
)

type Account = original.Account
type AccountCreateOrUpdateParameters = original.AccountCreateOrUpdateParameters
type AccountCreateOrUpdateProperties = original.AccountCreateOrUpdateProperties
type AccountListResult = original.AccountListResult
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountUpdateProperties = original.AccountUpdateProperties
type Activity = original.Activity
type ActivityListResult = original.ActivityListResult
type ActivityOutputType = original.ActivityOutputType
type ActivityParameter = original.ActivityParameter
type ActivityParameterSet = original.ActivityParameterSet
type ActivityProperties = original.ActivityProperties
type AdvancedSchedule = original.AdvancedSchedule
type AdvancedScheduleMonthlyOccurrence = original.AdvancedScheduleMonthlyOccurrence
type AgentRegistration = original.AgentRegistration
type AgentRegistrationKeys = original.AgentRegistrationKeys
type AgentRegistrationRegenerateKeyParameter = original.AgentRegistrationRegenerateKeyParameter
type Certificate = original.Certificate
type CertificateCreateOrUpdateParameters = original.CertificateCreateOrUpdateParameters
type CertificateCreateOrUpdateProperties = original.CertificateCreateOrUpdateProperties
type CertificateListResult = original.CertificateListResult
type CertificateProperties = original.CertificateProperties
type CertificateUpdateParameters = original.CertificateUpdateParameters
type CertificateUpdateProperties = original.CertificateUpdateProperties
type Connection = original.Connection
type ConnectionCreateOrUpdateParameters = original.ConnectionCreateOrUpdateParameters
type ConnectionCreateOrUpdateProperties = original.ConnectionCreateOrUpdateProperties
type ConnectionListResult = original.ConnectionListResult
type ConnectionProperties = original.ConnectionProperties
type ConnectionType = original.ConnectionType
type ConnectionTypeAssociationProperty = original.ConnectionTypeAssociationProperty
type ConnectionTypeCreateOrUpdateParameters = original.ConnectionTypeCreateOrUpdateParameters
type ConnectionTypeCreateOrUpdateProperties = original.ConnectionTypeCreateOrUpdateProperties
type ConnectionTypeListResult = original.ConnectionTypeListResult
type ConnectionTypeProperties = original.ConnectionTypeProperties
type ConnectionUpdateParameters = original.ConnectionUpdateParameters
type ConnectionUpdateProperties = original.ConnectionUpdateProperties
type ContentHash = original.ContentHash
type ContentLink = original.ContentLink
type ContentSource = original.ContentSource
type Credential = original.Credential
type CredentialCreateOrUpdateParameters = original.CredentialCreateOrUpdateParameters
type CredentialCreateOrUpdateProperties = original.CredentialCreateOrUpdateProperties
type CredentialListResult = original.CredentialListResult
type CredentialProperties = original.CredentialProperties
type CredentialUpdateParameters = original.CredentialUpdateParameters
type CredentialUpdateProperties = original.CredentialUpdateProperties
type DscCompilationJob = original.DscCompilationJob
type DscCompilationJobCreateParameters = original.DscCompilationJobCreateParameters
type DscCompilationJobCreateProperties = original.DscCompilationJobCreateProperties
type DscCompilationJobListResult = original.DscCompilationJobListResult
type DscCompilationJobProperties = original.DscCompilationJobProperties
type DscConfiguration = original.DscConfiguration
type DscConfigurationAssociationProperty = original.DscConfigurationAssociationProperty
type DscConfigurationCreateOrUpdateParameters = original.DscConfigurationCreateOrUpdateParameters
type DscConfigurationCreateOrUpdateProperties = original.DscConfigurationCreateOrUpdateProperties
type DscConfigurationListResult = original.DscConfigurationListResult
type DscConfigurationParameter = original.DscConfigurationParameter
type DscConfigurationProperties = original.DscConfigurationProperties
type DscMetaConfiguration = original.DscMetaConfiguration
type DscNode = original.DscNode
type DscNodeConfiguration = original.DscNodeConfiguration
type DscNodeConfigurationAssociationProperty = original.DscNodeConfigurationAssociationProperty
type DscNodeConfigurationCreateOrUpdateParameters = original.DscNodeConfigurationCreateOrUpdateParameters
type DscNodeConfigurationListResult = original.DscNodeConfigurationListResult
type DscNodeListResult = original.DscNodeListResult
type DscNodeReport = original.DscNodeReport
type DscNodeReportListResult = original.DscNodeReportListResult
type DscNodeUpdateParameters = original.DscNodeUpdateParameters
type DscReportError = original.DscReportError
type DscReportResource = original.DscReportResource
type DscReportResourceNavigation = original.DscReportResourceNavigation
type ErrorResponse = original.ErrorResponse
type FieldDefinition = original.FieldDefinition
type HybridRunbookWorker = original.HybridRunbookWorker
type HybridRunbookWorkerGroup = original.HybridRunbookWorkerGroup
type HybridRunbookWorkerGroupsListResult = original.HybridRunbookWorkerGroupsListResult
type HybridRunbookWorkerGroupUpdateParameters = original.HybridRunbookWorkerGroupUpdateParameters
type Job = original.Job
type JobCreateParameters = original.JobCreateParameters
type JobCreateProperties = original.JobCreateProperties
type JobListResult = original.JobListResult
type JobProperties = original.JobProperties
type JobSchedule = original.JobSchedule
type JobScheduleCreateParameters = original.JobScheduleCreateParameters
type JobScheduleCreateProperties = original.JobScheduleCreateProperties
type JobScheduleListResult = original.JobScheduleListResult
type JobScheduleProperties = original.JobScheduleProperties
type JobStream = original.JobStream
type JobStreamListResult = original.JobStreamListResult
type JobStreamProperties = original.JobStreamProperties
type Module = original.Module
type ModuleCreateOrUpdateParameters = original.ModuleCreateOrUpdateParameters
type ModuleCreateOrUpdateProperties = original.ModuleCreateOrUpdateProperties
type ModuleErrorInfo = original.ModuleErrorInfo
type ModuleListResult = original.ModuleListResult
type ModuleProperties = original.ModuleProperties
type ModuleUpdateParameters = original.ModuleUpdateParameters
type ModuleUpdateProperties = original.ModuleUpdateProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type ReadCloser = original.ReadCloser
type Resource = original.Resource
type RunAsCredentialAssociationProperty = original.RunAsCredentialAssociationProperty
type Runbook = original.Runbook
type RunbookAssociationProperty = original.RunbookAssociationProperty
type RunbookCreateOrUpdateDraftParameters = original.RunbookCreateOrUpdateDraftParameters
type RunbookCreateOrUpdateDraftProperties = original.RunbookCreateOrUpdateDraftProperties
type RunbookCreateOrUpdateParameters = original.RunbookCreateOrUpdateParameters
type RunbookCreateOrUpdateProperties = original.RunbookCreateOrUpdateProperties
type RunbookDraft = original.RunbookDraft
type RunbookDraftUndoEditResult = original.RunbookDraftUndoEditResult
type RunbookListResult = original.RunbookListResult
type RunbookParameter = original.RunbookParameter
type RunbookProperties = original.RunbookProperties
type RunbookUpdateParameters = original.RunbookUpdateParameters
type RunbookUpdateProperties = original.RunbookUpdateProperties
type Schedule = original.Schedule
type ScheduleAssociationProperty = original.ScheduleAssociationProperty
type ScheduleCreateOrUpdateParameters = original.ScheduleCreateOrUpdateParameters
type ScheduleCreateOrUpdateProperties = original.ScheduleCreateOrUpdateProperties
type ScheduleListResult = original.ScheduleListResult
type ScheduleProperties = original.ScheduleProperties
type ScheduleUpdateParameters = original.ScheduleUpdateParameters
type ScheduleUpdateProperties = original.ScheduleUpdateProperties
type Sku = original.Sku
type Statistics = original.Statistics
type StatisticsListResult = original.StatisticsListResult
type String = original.String
type SubResource = original.SubResource
type TestJob = original.TestJob
type TestJobCreateParameters = original.TestJobCreateParameters
type TypeField = original.TypeField
type TypeFieldListResult = original.TypeFieldListResult
type Usage = original.Usage
type UsageCounterName = original.UsageCounterName
type UsageListResult = original.UsageListResult
type Variable = original.Variable
type VariableCreateOrUpdateParameters = original.VariableCreateOrUpdateParameters
type VariableCreateOrUpdateProperties = original.VariableCreateOrUpdateProperties
type VariableListResult = original.VariableListResult
type VariableProperties = original.VariableProperties
type VariableUpdateParameters = original.VariableUpdateParameters
type VariableUpdateProperties = original.VariableUpdateProperties
type Webhook = original.Webhook
type WebhookCreateOrUpdateParameters = original.WebhookCreateOrUpdateParameters
type WebhookCreateOrUpdateProperties = original.WebhookCreateOrUpdateProperties
type WebhookListResult = original.WebhookListResult
type WebhookProperties = original.WebhookProperties
type WebhookUpdateParameters = original.WebhookUpdateParameters
type WebhookUpdateProperties = original.WebhookUpdateProperties
type ModuleClient = original.ModuleClient
type RunbookClient = original.RunbookClient
type StatisticsClient = original.StatisticsClient
type JobStreamClient = original.JobStreamClient
type CredentialClient = original.CredentialClient
type DscConfigurationClient = original.DscConfigurationClient
type FieldsClient = original.FieldsClient
type ObjectDataTypesClient = original.ObjectDataTypesClient
type TestJobStreamsClient = original.TestJobStreamsClient
type UsagesClient = original.UsagesClient
type ConnectionTypeClient = original.ConnectionTypeClient
type WebhookClient = original.WebhookClient
type JobScheduleClient = original.JobScheduleClient
type DscCompilationJobClient = original.DscCompilationJobClient
type HybridRunbookWorkerGroupClient = original.HybridRunbookWorkerGroupClient
type OperationsClient = original.OperationsClient
type TestJobsClient = original.TestJobsClient
type CertificateClient = original.CertificateClient

func NewActivityClient(subscriptionID string) ActivityClient {
	return original.NewActivityClient(subscriptionID)
}
func NewActivityClientWithBaseURI(baseURI string, subscriptionID string) ActivityClient {
	return original.NewActivityClientWithBaseURI(baseURI, subscriptionID)
}
func NewRunbookDraftClient(subscriptionID string) RunbookDraftClient {
	return original.NewRunbookDraftClient(subscriptionID)
}
func NewRunbookDraftClientWithBaseURI(baseURI string, subscriptionID string) RunbookDraftClient {
	return original.NewRunbookDraftClientWithBaseURI(baseURI, subscriptionID)
}
func NewScheduleClient(subscriptionID string) ScheduleClient {
	return original.NewScheduleClient(subscriptionID)
}
func NewScheduleClientWithBaseURI(baseURI string, subscriptionID string) ScheduleClient {
	return original.NewScheduleClientWithBaseURI(baseURI, subscriptionID)
}
func NewVariableClient(subscriptionID string) VariableClient {
	return original.NewVariableClient(subscriptionID)
}
func NewVariableClientWithBaseURI(baseURI string, subscriptionID string) VariableClient {
	return original.NewVariableClientWithBaseURI(baseURI, subscriptionID)
}
func NewAgentRegistrationInformationClient(subscriptionID string) AgentRegistrationInformationClient {
	return original.NewAgentRegistrationInformationClient(subscriptionID)
}
func NewAgentRegistrationInformationClientWithBaseURI(baseURI string, subscriptionID string) AgentRegistrationInformationClient {
	return original.NewAgentRegistrationInformationClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewConnectionClient(subscriptionID string) ConnectionClient {
	return original.NewConnectionClient(subscriptionID)
}
func NewConnectionClientWithBaseURI(baseURI string, subscriptionID string) ConnectionClient {
	return original.NewConnectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccountClient(subscriptionID string) AccountClient {
	return original.NewAccountClient(subscriptionID)
}
func NewAccountClientWithBaseURI(baseURI string, subscriptionID string) AccountClient {
	return original.NewAccountClientWithBaseURI(baseURI, subscriptionID)
}
func NewDscNodeClient(subscriptionID string) DscNodeClient {
	return original.NewDscNodeClient(subscriptionID)
}
func NewDscNodeClientWithBaseURI(baseURI string, subscriptionID string) DscNodeClient {
	return original.NewDscNodeClientWithBaseURI(baseURI, subscriptionID)
}
func NewDscNodeConfigurationClient(subscriptionID string) DscNodeConfigurationClient {
	return original.NewDscNodeConfigurationClient(subscriptionID)
}
func NewDscNodeConfigurationClientWithBaseURI(baseURI string, subscriptionID string) DscNodeConfigurationClient {
	return original.NewDscNodeConfigurationClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobClient(subscriptionID string) JobClient {
	return original.NewJobClient(subscriptionID)
}
func NewJobClientWithBaseURI(baseURI string, subscriptionID string) JobClient {
	return original.NewJobClientWithBaseURI(baseURI, subscriptionID)
}
func NewNodeReportsClient(subscriptionID string) NodeReportsClient {
	return original.NewNodeReportsClient(subscriptionID)
}
func NewNodeReportsClientWithBaseURI(baseURI string, subscriptionID string) NodeReportsClient {
	return original.NewNodeReportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobStreamClient(subscriptionID string) JobStreamClient {
	return original.NewJobStreamClient(subscriptionID)
}
func NewJobStreamClientWithBaseURI(baseURI string, subscriptionID string) JobStreamClient {
	return original.NewJobStreamClientWithBaseURI(baseURI, subscriptionID)
}
func NewModuleClient(subscriptionID string) ModuleClient {
	return original.NewModuleClient(subscriptionID)
}
func NewModuleClientWithBaseURI(baseURI string, subscriptionID string) ModuleClient {
	return original.NewModuleClientWithBaseURI(baseURI, subscriptionID)
}
func NewRunbookClient(subscriptionID string) RunbookClient {
	return original.NewRunbookClient(subscriptionID)
}
func NewRunbookClientWithBaseURI(baseURI string, subscriptionID string) RunbookClient {
	return original.NewRunbookClientWithBaseURI(baseURI, subscriptionID)
}
func NewStatisticsClient(subscriptionID string) StatisticsClient {
	return original.NewStatisticsClient(subscriptionID)
}
func NewStatisticsClientWithBaseURI(baseURI string, subscriptionID string) StatisticsClient {
	return original.NewStatisticsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectionTypeClient(subscriptionID string) ConnectionTypeClient {
	return original.NewConnectionTypeClient(subscriptionID)
}
func NewConnectionTypeClientWithBaseURI(baseURI string, subscriptionID string) ConnectionTypeClient {
	return original.NewConnectionTypeClientWithBaseURI(baseURI, subscriptionID)
}
func NewCredentialClient(subscriptionID string) CredentialClient {
	return original.NewCredentialClient(subscriptionID)
}
func NewCredentialClientWithBaseURI(baseURI string, subscriptionID string) CredentialClient {
	return original.NewCredentialClientWithBaseURI(baseURI, subscriptionID)
}
func NewDscConfigurationClient(subscriptionID string) DscConfigurationClient {
	return original.NewDscConfigurationClient(subscriptionID)
}
func NewDscConfigurationClientWithBaseURI(baseURI string, subscriptionID string) DscConfigurationClient {
	return original.NewDscConfigurationClientWithBaseURI(baseURI, subscriptionID)
}
func NewFieldsClient(subscriptionID string) FieldsClient {
	return original.NewFieldsClient(subscriptionID)
}
func NewFieldsClientWithBaseURI(baseURI string, subscriptionID string) FieldsClient {
	return original.NewFieldsClientWithBaseURI(baseURI, subscriptionID)
}
func NewObjectDataTypesClient(subscriptionID string) ObjectDataTypesClient {
	return original.NewObjectDataTypesClient(subscriptionID)
}
func NewObjectDataTypesClientWithBaseURI(baseURI string, subscriptionID string) ObjectDataTypesClient {
	return original.NewObjectDataTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTestJobStreamsClient(subscriptionID string) TestJobStreamsClient {
	return original.NewTestJobStreamsClient(subscriptionID)
}
func NewTestJobStreamsClientWithBaseURI(baseURI string, subscriptionID string) TestJobStreamsClient {
	return original.NewTestJobStreamsClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobScheduleClient(subscriptionID string) JobScheduleClient {
	return original.NewJobScheduleClient(subscriptionID)
}
func NewJobScheduleClientWithBaseURI(baseURI string, subscriptionID string) JobScheduleClient {
	return original.NewJobScheduleClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewWebhookClient(subscriptionID string) WebhookClient {
	return original.NewWebhookClient(subscriptionID)
}
func NewWebhookClientWithBaseURI(baseURI string, subscriptionID string) WebhookClient {
	return original.NewWebhookClientWithBaseURI(baseURI, subscriptionID)
}
func NewCertificateClient(subscriptionID string) CertificateClient {
	return original.NewCertificateClient(subscriptionID)
}
func NewCertificateClientWithBaseURI(baseURI string, subscriptionID string) CertificateClient {
	return original.NewCertificateClientWithBaseURI(baseURI, subscriptionID)
}
func NewDscCompilationJobClient(subscriptionID string) DscCompilationJobClient {
	return original.NewDscCompilationJobClient(subscriptionID)
}
func NewDscCompilationJobClientWithBaseURI(baseURI string, subscriptionID string) DscCompilationJobClient {
	return original.NewDscCompilationJobClientWithBaseURI(baseURI, subscriptionID)
}
func NewHybridRunbookWorkerGroupClient(subscriptionID string) HybridRunbookWorkerGroupClient {
	return original.NewHybridRunbookWorkerGroupClient(subscriptionID)
}
func NewHybridRunbookWorkerGroupClientWithBaseURI(baseURI string, subscriptionID string) HybridRunbookWorkerGroupClient {
	return original.NewHybridRunbookWorkerGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTestJobsClient(subscriptionID string) TestJobsClient {
	return original.NewTestJobsClient(subscriptionID)
}
func NewTestJobsClientWithBaseURI(baseURI string, subscriptionID string) TestJobsClient {
	return original.NewTestJobsClientWithBaseURI(baseURI, subscriptionID)
}
