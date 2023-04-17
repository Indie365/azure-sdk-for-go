//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationsmanagement

// ArmTemplateParameter - Parameter to pass to ARM template
type ArmTemplateParameter struct {
	// name of the parameter.
	Name *string

	// value for the parameter. In Jtoken
	Value *string
}

// CodeMessageError - The error body contract.
type CodeMessageError struct {
	// The error details for a failed request.
	Error *CodeMessageErrorError
}

// CodeMessageErrorError - The error details for a failed request.
type CodeMessageErrorError struct {
	// The error type.
	Code *string

	// The error message.
	Message *string
}

// ManagementAssociation - The container for solution.
type ManagementAssociation struct {
	// Resource location
	Location *string

	// Properties for ManagementAssociation object supported by the OperationsManagement resource provider.
	Properties *ManagementAssociationProperties

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// ManagementAssociationProperties - ManagementAssociation properties supported by the OperationsManagement resource provider.
type ManagementAssociationProperties struct {
	// REQUIRED; The applicationId of the appliance for this association.
	ApplicationID *string
}

// ManagementAssociationPropertiesList - the list of ManagementAssociation response
type ManagementAssociationPropertiesList struct {
	// List of Management Association properties within the subscription.
	Value []*ManagementAssociation
}

// ManagementAssociationsClientCreateOrUpdateOptions contains the optional parameters for the ManagementAssociationsClient.CreateOrUpdate
// method.
type ManagementAssociationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagementAssociationsClientDeleteOptions contains the optional parameters for the ManagementAssociationsClient.Delete
// method.
type ManagementAssociationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ManagementAssociationsClientGetOptions contains the optional parameters for the ManagementAssociationsClient.Get method.
type ManagementAssociationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagementAssociationsClientListBySubscriptionOptions contains the optional parameters for the ManagementAssociationsClient.ListBySubscription
// method.
type ManagementAssociationsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ManagementConfiguration - The container for solution.
type ManagementConfiguration struct {
	// Resource location
	Location *string

	// Properties for ManagementConfiguration object supported by the OperationsManagement resource provider.
	Properties *ManagementConfigurationProperties

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// ManagementConfigurationProperties - ManagementConfiguration properties supported by the OperationsManagement resource provider.
type ManagementConfigurationProperties struct {
	// REQUIRED; Parameters to run the ARM template
	Parameters []*ArmTemplateParameter

	// REQUIRED; The type of the parent resource.
	ParentResourceType *string

	// REQUIRED; The Json object containing the ARM template to deploy
	Template any

	// The applicationId of the appliance for this Management.
	ApplicationID *string

	// READ-ONLY; The provisioning state for the ManagementConfiguration.
	ProvisioningState *string
}

// ManagementConfigurationPropertiesList - the list of ManagementConfiguration response
type ManagementConfigurationPropertiesList struct {
	// List of Management Configuration properties within the subscription.
	Value []*ManagementConfiguration
}

// ManagementConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the ManagementConfigurationsClient.CreateOrUpdate
// method.
type ManagementConfigurationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagementConfigurationsClientDeleteOptions contains the optional parameters for the ManagementConfigurationsClient.Delete
// method.
type ManagementConfigurationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ManagementConfigurationsClientGetOptions contains the optional parameters for the ManagementConfigurationsClient.Get method.
type ManagementConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagementConfigurationsClientListBySubscriptionOptions contains the optional parameters for the ManagementConfigurationsClient.ListBySubscription
// method.
type ManagementConfigurationsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// Operation - Supported operation of OperationsManagement resource provider.
type Operation struct {
	// Display metadata associated with the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Type of operation: get, read, delete, etc.
	Operation *string

	// Service provider: Microsoft OperationsManagement.
	Provider *string

	// Resource on which the operation is performed etc.
	Resource *string
}

// OperationListResult - Result of the request to list solution operations.
type OperationListResult struct {
	// List of solution operations supported by the OperationsManagement resource provider.
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Solution - The container for solution.
type Solution struct {
	// Resource location
	Location *string

	// Plan for solution object supported by the OperationsManagement resource provider.
	Plan *SolutionPlan

	// Properties for solution object supported by the OperationsManagement resource provider.
	Properties *SolutionProperties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// SolutionPatch - The properties of a Solution that can be patched.
type SolutionPatch struct {
	// Resource tags
	Tags map[string]*string
}

// SolutionPlan - Plan for solution object supported by the OperationsManagement resource provider.
type SolutionPlan struct {
	// name of the solution to be created. For Microsoft published solution it should be in the format of solutionType(workspaceName).
	// SolutionType part is case sensitive. For third party solution, it can be
	// anything.
	Name *string

	// name of the solution to enabled/add. For Microsoft published gallery solution it should be in the format of OMSGallery/.
	// This is case sensitive
	Product *string

	// promotionCode, Not really used now, can you left as empty
	PromotionCode *string

	// Publisher name. For gallery solution, it is Microsoft.
	Publisher *string
}

// SolutionProperties - Solution properties supported by the OperationsManagement resource provider.
type SolutionProperties struct {
	// REQUIRED; The azure resourceId for the workspace where the solution will be deployed/enabled.
	WorkspaceResourceID *string

	// The azure resources that will be contained within the solutions. They will be locked and gets deleted automatically when
	// the solution is deleted.
	ContainedResources []*string

	// The resources that will be referenced from this solution. Deleting any of those solution out of band will break the solution.
	ReferencedResources []*string

	// READ-ONLY; The provisioning state for the solution.
	ProvisioningState *string
}

// SolutionPropertiesList - the list of solution response
type SolutionPropertiesList struct {
	// List of solution properties within the subscription.
	Value []*Solution
}

// SolutionsClientBeginCreateOrUpdateOptions contains the optional parameters for the SolutionsClient.BeginCreateOrUpdate
// method.
type SolutionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SolutionsClientBeginDeleteOptions contains the optional parameters for the SolutionsClient.BeginDelete method.
type SolutionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SolutionsClientBeginUpdateOptions contains the optional parameters for the SolutionsClient.BeginUpdate method.
type SolutionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SolutionsClientGetOptions contains the optional parameters for the SolutionsClient.Get method.
type SolutionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SolutionsClientListByResourceGroupOptions contains the optional parameters for the SolutionsClient.ListByResourceGroup
// method.
type SolutionsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SolutionsClientListBySubscriptionOptions contains the optional parameters for the SolutionsClient.ListBySubscription method.
type SolutionsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}
