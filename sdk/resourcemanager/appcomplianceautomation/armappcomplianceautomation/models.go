//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

import "time"

// Assessment - A class represent the assessment.
type Assessment struct {
	// READ-ONLY; The description of the assessment.
	Description *string

	// READ-ONLY; Indicates whether all the resource(s) are compliant.
	IsPass *IsPass

	// READ-ONLY; The name of the assessment.
	Name *string

	// READ-ONLY; The policy id mapping to this assessment.
	PolicyID *string

	// READ-ONLY; The remediation of the assessment.
	Remediation *string

	// READ-ONLY; List of resource assessments.
	ResourceList []*AssessmentResource

	// READ-ONLY; The severity level of this assessment.
	Severity *AssessmentSeverity
}

// AssessmentResource - A class represent the assessment resource.
type AssessmentResource struct {
	// READ-ONLY; The reason for the N/A resource.
	Reason *string

	// READ-ONLY; The Id of the resource.
	ResourceID *string

	// READ-ONLY; Resource status.
	ResourceStatus *ResourceStatus

	// READ-ONLY; The status change date for the resource. For unavailable date, set it as N/A.
	StatusChangeDate *string
}

// Category - A class represent the compliance category.
type Category struct {
	// READ-ONLY; The name of the compliance category. e.g. "Operational Security"
	CategoryName *string

	// READ-ONLY; Category status.
	CategoryStatus *CategoryStatus

	// READ-ONLY; The category type
	CategoryType *CategoryType

	// READ-ONLY; List of control families.
	ControlFamilies []*ControlFamily
}

// ComplianceReportItem - Object that includes all the content for single compliance result.
type ComplianceReportItem struct {
	// READ-ONLY; The category name.
	CategoryName *string

	// READ-ONLY; The compliance result's status.
	ComplianceState *ComplianceState

	// READ-ONLY; The control Id - e.g. "1".
	ControlID *string

	// READ-ONLY; The control name.
	ControlName *string

	// READ-ONLY; The control type.
	ControlType *ControlType

	// READ-ONLY; The policy's detail description.
	PolicyDescription *string

	// READ-ONLY; The policy's display name.
	PolicyDisplayName *string

	// READ-ONLY; The compliance result mapped policy Id.
	PolicyID *string

	// READ-ONLY; The compliance result mapped resource group.
	ResourceGroup *string

	// READ-ONLY; The compliance result mapped resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
	ResourceID *string

	// READ-ONLY; The compliance result mapped resource type.
	ResourceType *string

	// READ-ONLY; The compliance result last changed date - e.g. "2022-10-24T02:55:16.3274379Z". For unavailable date, set it
// as "N/A".
	StatusChangeDate *string

	// READ-ONLY; The compliance result mapped subscription Id.
	SubscriptionID *string
}

// ComplianceResult - A class represent the compliance result.
type ComplianceResult struct {
	// READ-ONLY; List of categories.
	Categories []*Category

	// READ-ONLY; The name of the compliance. e.g. "M365"
	ComplianceName *string
}

// Control - A class represent the control.
type Control struct {
	// READ-ONLY; List of assessments.
	Assessments []*Assessment

	// READ-ONLY; The control's description
	ControlDescription *string

	// READ-ONLY; The hyper link to the control's description'.
	ControlDescriptionHyperLink *string

	// READ-ONLY; The full name of the control. e.g. "Validate that unsupported operating systems and software components are
// not in use."
	ControlFullName *string

	// READ-ONLY; The Id of the control. e.g. "Operational Security#10"
	ControlID *string

	// READ-ONLY; The short name of the control. e.g. "Unsupported OS and Software."
	ControlShortName *string

	// READ-ONLY; Control status.
	ControlStatus *ControlStatus

	// READ-ONLY; The control type
	ControlType *ControlType
}

// ControlFamily - A class represent the control family.
type ControlFamily struct {
	// READ-ONLY; List of controls.
	Controls []*Control

	// READ-ONLY; The name of the control family. e.g. "Malware Protection - Anti-Virus"
	FamilyName *string

	// READ-ONLY; Control family status.
	FamilyStatus *ControlFamilyStatus

	// READ-ONLY; The control family type
	FamilyType *ControlFamilyType
}

// DownloadResponse - Object that includes all the possible response for the download operation.
type DownloadResponse struct {
	// READ-ONLY; compliance detailed pdf report
	ComplianceDetailedPDFReport *DownloadResponseComplianceDetailedPDFReport

	// READ-ONLY; compliance pdf report
	CompliancePDFReport *DownloadResponseCompliancePDFReport

	// READ-ONLY; List of the compliance result
	ComplianceReport []*ComplianceReportItem

	// READ-ONLY; List of the reports
	ResourceList []*ResourceItem
}

// DownloadResponseComplianceDetailedPDFReport - compliance detailed pdf report
type DownloadResponseComplianceDetailedPDFReport struct {
	// READ-ONLY; uri of compliance detailed pdf report
	SasURI *string
}

// DownloadResponseCompliancePDFReport - compliance pdf report
type DownloadResponseCompliancePDFReport struct {
	// READ-ONLY; uri of compliance pdf report
	SasURI *string
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

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OverviewStatus - The overview of the compliance result for one report.
type OverviewStatus struct {
	// The count of all failed full automation control.
	FailedCount *int32

	// The count of all manual control.
	ManualCount *int32

	// The count of all passed full automation control.
	PassedCount *int32
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ReportComplianceStatus - A list which includes all the compliance result for one report.
type ReportComplianceStatus struct {
	// The Microsoft 365 certification name.
	M365 *OverviewStatus
}

// ReportProperties - Report's properties.
type ReportProperties struct {
	// REQUIRED; List of resource data.
	Resources []*ResourceMetadata

	// REQUIRED; Report collection trigger time's time zone, the available list can be obtained by executing "Get-TimeZone -ListAvailable"
// in PowerShell. An example of valid timezone id is "Pacific Standard Time".
	TimeZone *string

	// REQUIRED; Report collection trigger time.
	TriggerTime *time.Time

	// Report offer Guid.
	OfferGUID *string

	// READ-ONLY; Report compliance status.
	ComplianceStatus *ReportComplianceStatus

	// READ-ONLY; Report id in database.
	ID *string

	// READ-ONLY; Report last collection trigger time.
	LastTriggerTime *time.Time

	// READ-ONLY; Report next collection trigger time.
	NextTriggerTime *time.Time

	// READ-ONLY; Azure lifecycle management
	ProvisioningState *ProvisioningState

	// READ-ONLY; Report name.
	ReportName *string

	// READ-ONLY; Report status.
	Status *ReportStatus

	// READ-ONLY; List of subscription Ids.
	Subscriptions []*string

	// READ-ONLY; Report's tenant id.
	TenantID *string
}

// ReportResource - A class represent an AppComplianceAutomation report resource.
type ReportResource struct {
	// REQUIRED; Report property.
	Properties *ReportProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ReportResourceList - Object that includes an array of resources and a possible link for next set.
type ReportResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; List of the reports
	Value []*ReportResource
}

// ReportResourcePatch - A class represent a AppComplianceAutomation report resource update properties.
type ReportResourcePatch struct {
	// Report property.
	Properties *ReportProperties
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceItem - Resource Id.
type ResourceItem struct {
	// READ-ONLY; The resource group name of this resource.
	ResourceGroup *string

	// READ-ONLY; The resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
	ResourceID *string

	// READ-ONLY; The resource type of this resource.
	ResourceType *string

	// READ-ONLY; The subscription Id of this resource.
	SubscriptionID *string
}

// ResourceMetadata - Single resource Id's metadata.
type ResourceMetadata struct {
	// REQUIRED; Resource Id - e.g. "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1".
	ResourceID *string

	// Resource kind.
	ResourceKind *string

	// Resource name.
	ResourceName *string

	// Resource type.
	ResourceType *string

	// Resource's tag type.
	Tags map[string]*string
}

// SnapshotDownloadRequest - Snapshot's download request.
type SnapshotDownloadRequest struct {
	// REQUIRED; Indicates the download type.
	DownloadType *DownloadType

	// The offerGuid which mapping to the reports.
	OfferGUID *string

	// Tenant id.
	ReportCreatorTenantID *string
}

// SnapshotProperties - Snapshot's properties.
type SnapshotProperties struct {
	// READ-ONLY; List of compliance results.
	ComplianceResults []*ComplianceResult

	// READ-ONLY; The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// READ-ONLY; Snapshot id in the database.
	ID *string

	// READ-ONLY; Azure lifecycle management
	ProvisioningState *ProvisioningState

	// READ-ONLY; The report essential info.
	ReportProperties *ReportProperties

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	ReportSystemData *SystemData

	// READ-ONLY; Snapshot name.
	SnapshotName *string
}

// SnapshotResource - A class represent a AppComplianceAutomation snapshot resource.
type SnapshotResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Snapshot's property'.
	Properties *SnapshotProperties

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SnapshotResourceList - Object that includes an array of resources and a possible link for next set.
type SnapshotResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; List of the snapshots
	Value []*SnapshotResource
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

