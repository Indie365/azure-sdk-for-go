//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

import "time"

type APIKey struct {
	// REQUIRED; The value of the API key.
	Key *string

	// The time of creation of the API key.
	Created *string

	// The user that created the API key.
	CreatedBy *string

	// The name of the API key.
	Name *string
}

// APIKeyListResponse - Response of a list operation.
type APIKeyListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*APIKey
}

// AgreementProperties - Terms properties.
type AgreementProperties struct {
	// If any version of the terms have been accepted, otherwise false.
	Accepted *bool

	// Link to HTML with Microsoft and Publisher terms.
	LicenseTextLink *string

	// Plan identifier string.
	Plan *string

	// Link to the privacy policy of the publisher.
	PrivacyPolicyLink *string

	// Product identifier string.
	Product *string

	// Publisher identifier string.
	Publisher *string

	// Date and time in UTC of when the terms were accepted. This is empty if Accepted is false.
	RetrieveDatetime *time.Time

	// Terms signature.
	Signature *string
}

type AgreementResource struct {
	// Represents the properties of the resource.
	Properties *AgreementProperties

	// READ-ONLY; ARM id of the resource.
	ID *string

	// READ-ONLY; Name of the agreement.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource.
	Type *string
}

// AgreementResourceListResponse - Response of a list operation.
type AgreementResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*AgreementResource
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them
// from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag. Exclusion takes priority over inclusion.
	Action *TagAction

	// The name (also known as the key) of the tag.
	Name *string

	// The value of the tag.
	Value *string
}

type Host struct {
	// The aliases for the host.
	Aliases []*string

	// The Datadog integrations reporting metrics for the host.
	Apps []*string
	Meta *HostMetadata

	// The name of the host.
	Name *string
}

// HostListResponse - Response of a list operation.
type HostListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*Host
}

type HostMetadata struct {
	// The agent version.
	AgentVersion *string
	InstallMethod *InstallMethod
	LogsAgent *LogsAgent
}

type IdentityProperties struct {
	// Identity type
	Type *ManagedIdentityTypes

	// READ-ONLY; The identity ID.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

type InstallMethod struct {
	// The installer version.
	InstallerVersion *string

	// The tool.
	Tool *string

	// The tool version.
	ToolVersion *string
}

// LinkedResource - The definition of a linked resource.
type LinkedResource struct {
	// The ARM id of the linked resource.
	ID *string
}

// LinkedResourceListResponse - Response of a list operation.
type LinkedResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*LinkedResource
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendResourceLogs flag is enabled. If empty,
// all resources will be captured. If only Exclude action is specified, the
// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
// resources with the associated tags.
	FilteringTags []*FilteringTag

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *bool

	// Flag specifying if Azure resource logs should be sent for the Monitor resource.
	SendResourceLogs *bool

	// Flag specifying if Azure subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *bool
}

type LogsAgent struct {
	// The transport.
	Transport *string
}

// MetricRules - Set of rules for sending metrics for the Monitor resource.
type MetricRules struct {
	// List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action
// is specified, the rules will apply to the list of all available resources. If
// Include actions are specified, the rules will only include resources with the associated tags.
	FilteringTags []*FilteringTag
}

// MonitorProperties - Properties specific to the monitor resource.
type MonitorProperties struct {
	// Datadog organization properties
	DatadogOrganizationProperties *OrganizationProperties

	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus

	// User info
	UserInfo *UserInfo

	// READ-ONLY
	LiftrResourceCategory *LiftrResourceCategories

	// READ-ONLY; The priority of the resource.
	LiftrResourcePreference *int32

	// READ-ONLY; Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in time, the resource
// will go in Suspended state.
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

type MonitorResource struct {
	// REQUIRED
	Location *string
	Identity *IdentityProperties

	// Properties specific to the monitor resource.
	Properties *MonitorProperties
	SKU *ResourceSKU

	// Dictionary of
	Tags map[string]*string

	// READ-ONLY; ARM id of the monitor resource.
	ID *string

	// READ-ONLY; Name of the monitor resource.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the monitor resource.
	Type *string
}

// MonitorResourceListResponse - Response of a list operation.
type MonitorResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*MonitorResource
}

// MonitorResourceUpdateParameters - The parameters for a PATCH request to a monitor resource.
type MonitorResourceUpdateParameters struct {
	// The set of properties that can be update in a PATCH request to a monitor resource.
	Properties *MonitorUpdateProperties
	SKU *ResourceSKU

	// The new tags of the monitor resource.
	Tags map[string]*string
}

// MonitorUpdateProperties - The set of properties that can be update in a PATCH request to a monitor resource.
type MonitorUpdateProperties struct {
	// Flag specifying if the resource monitoring is enabled or disabled.
	MonitoringStatus *MonitoringStatus
}

// MonitoredResource - The properties of a resource currently being monitored by the Datadog monitor resource.
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string

	// Flag indicating if resource is sending logs to Datadog.
	SendingLogs *bool

	// Flag indicating if resource is sending metrics to Datadog.
	SendingMetrics *bool
}

// MonitoredResourceListResponse - Response of a list operation.
type MonitoredResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*MonitoredResource
}

// MonitoringTagRules - Capture logs and metrics of Azure resources based on ARM tags.
type MonitoringTagRules struct {
	// Definition of the properties for a TagRules resource.
	Properties *MonitoringTagRulesProperties

	// READ-ONLY; The id of the rule set.
	ID *string

	// READ-ONLY; Name of the rule set.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the rule set.
	Type *string
}

// MonitoringTagRulesListResponse - Response of a list operation.
type MonitoringTagRulesListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*MonitoringTagRules
}

// MonitoringTagRulesProperties - Definition of the properties for a TagRules resource.
type MonitoringTagRulesProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write monitors'.
	Description *string

	// Operation type, e.g., read, write, delete, etc.
	Operation *string

	// Service provider, i.e., Microsoft.Datadog.
	Provider *string

	// Type on which the operation is performed, e.g., 'monitors'.
	Resource *string
}

// OperationListResult - Result of GET request to list the Microsoft.Datadog operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of operations supported by the Microsoft.Datadog provider.
	Value []*OperationResult
}

// OperationResult - A Microsoft.Datadog REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name, i.e., {provider}/{resource}/{operation}.
	Name *string
}

// OrganizationProperties - Datadog organization properties
type OrganizationProperties struct {
	// Api key associated to the Datadog organization.
	APIKey *string

	// Application key associated to the Datadog organization.
	ApplicationKey *string

	// The Id of the Enterprise App used for Single sign on.
	EnterpriseAppID *string

	// The auth code used to linking to an existing datadog organization.
	LinkingAuthCode *string

	// The client_id from an existing in exchange for an auth token to link organization.
	LinkingClientID *string

	// The redirect uri for linking.
	RedirectURI *string

	// READ-ONLY; Id of the Datadog organization.
	ID *string

	// READ-ONLY; Name of the Datadog organization.
	Name *string
}

type ResourceSKU struct {
	// REQUIRED; Name of the SKU.
	Name *string
}

type SetPasswordLink struct {
	SetPasswordLink *string
}

type SingleSignOnProperties struct {
	// The Id of the Enterprise App used for Single sign-on.
	EnterpriseAppID *string

	// Various states of the SSO resource
	SingleSignOnState *SingleSignOnStates

	// READ-ONLY
	ProvisioningState *ProvisioningState

	// READ-ONLY; The login URL specific to this Datadog Organization.
	SingleSignOnURL *string
}

type SingleSignOnResource struct {
	Properties *SingleSignOnProperties

	// READ-ONLY; ARM id of the resource.
	ID *string

	// READ-ONLY; Name of the configuration.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource.
	Type *string
}

// SingleSignOnResourceListResponse - Response of a list operation.
type SingleSignOnResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*SingleSignOnResource
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserInfo - User info
type UserInfo struct {
	// Email of the user used by Datadog for contacting them if needed
	EmailAddress *string

	// Name of the user
	Name *string

	// Phone number of the user used by Datadog for contacting them if needed
	PhoneNumber *string
}

