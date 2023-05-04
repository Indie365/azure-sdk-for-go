//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos

import "time"

// ActionClassification provides polymorphic access to related types.
// Call the interface's GetAction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Action, *ContinuousAction, *DelayAction, *DiscreteAction
type ActionClassification interface {
	// GetAction returns the Action content of the underlying type.
	GetAction() *Action
}

// Action - Model that represents the base action model.
type Action struct {
	// REQUIRED; String that represents a Capability URN.
	Name *string

	// REQUIRED; Enum that discriminates between action models.
	Type *string
}

// GetAction implements the ActionClassification interface for type Action.
func (a *Action) GetAction() *Action { return a }

// ActionStatus - Model that represents the an action and its status.
type ActionStatus struct {
	// READ-ONLY; The id of the action status.
	ActionID *string

	// READ-ONLY; The name of the action status.
	ActionName *string

	// READ-ONLY; String that represents the end time of the action.
	EndTime *time.Time

	// READ-ONLY; String that represents the start time of the action.
	StartTime *time.Time

	// READ-ONLY; The status of the action.
	Status *string

	// READ-ONLY; The array of targets.
	Targets []*ExperimentExecutionActionTargetDetailsProperties
}

// Branch - Model that represents a branch in the step.
type Branch struct {
	// REQUIRED; List of actions.
	Actions []ActionClassification

	// REQUIRED; String of the branch name.
	Name *string
}

// BranchStatus - Model that represents the a list of actions and action statuses.
type BranchStatus struct {
	// READ-ONLY; The array of actions.
	Actions []*ActionStatus

	// READ-ONLY; The id of the branch status.
	BranchID *string

	// READ-ONLY; The name of the branch status.
	BranchName *string

	// READ-ONLY; The status of the branch.
	Status *string
}

// CapabilitiesClientCreateOrUpdateOptions contains the optional parameters for the CapabilitiesClient.CreateOrUpdate method.
type CapabilitiesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesClientDeleteOptions contains the optional parameters for the CapabilitiesClient.Delete method.
type CapabilitiesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesClientGetOptions contains the optional parameters for the CapabilitiesClient.Get method.
type CapabilitiesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CapabilitiesClientListOptions contains the optional parameters for the CapabilitiesClient.NewListPager method.
type CapabilitiesClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// Capability - Model that represents a Capability resource.
type Capability struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The properties of a capability resource.
	Properties *CapabilityProperties

	// READ-ONLY; The standard system metadata of a resource type.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CapabilityListResult - Model that represents a list of Capability resources and a link for pagination.
type CapabilityListResult struct {
	// READ-ONLY; URL to retrieve the next page of Capability resources.
	NextLink *string

	// READ-ONLY; List of Capability resources.
	Value []*Capability
}

// CapabilityProperties - Model that represents the Capability properties model.
type CapabilityProperties struct {
	// READ-ONLY; Localized string of the description.
	Description *string

	// READ-ONLY; URL to retrieve JSON schema of the Capability parameters.
	ParametersSchema *string

	// READ-ONLY; String of the Publisher that this Capability extends.
	Publisher *string

	// READ-ONLY; String of the Target Type that this Capability extends.
	TargetType *string

	// READ-ONLY; String of the URN for this Capability Type.
	Urn *string
}

// CapabilityType - Model that represents a Capability Type resource.
type CapabilityType struct {
	// Location of the Capability Type resource.
	Location *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The properties of the capability type resource.
	Properties *CapabilityTypeProperties

	// READ-ONLY; The system metadata properties of the capability type resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CapabilityTypeListResult - Model that represents a list of Capability Type resources and a link for pagination.
type CapabilityTypeListResult struct {
	// READ-ONLY; URL to retrieve the next page of Capability Type resources.
	NextLink *string

	// READ-ONLY; List of Capability Type resources.
	Value []*CapabilityType
}

// CapabilityTypeProperties - Model that represents the Capability Type properties model.
type CapabilityTypeProperties struct {
	// READ-ONLY; Localized string of the description.
	Description *string

	// READ-ONLY; Localized string of the display name.
	DisplayName *string

	// READ-ONLY; URL to retrieve JSON schema of the Capability Type parameters.
	ParametersSchema *string

	// READ-ONLY; String of the Publisher that this Capability Type extends.
	Publisher *string

	// READ-ONLY; String of the Target Type that this Capability Type extends.
	TargetType *string

	// READ-ONLY; String of the URN for this Capability Type.
	Urn *string
}

// CapabilityTypesClientGetOptions contains the optional parameters for the CapabilityTypesClient.Get method.
type CapabilityTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CapabilityTypesClientListOptions contains the optional parameters for the CapabilityTypesClient.NewListPager method.
type CapabilityTypesClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// ContinuousAction - Model that represents a continuous action.
type ContinuousAction struct {
	// REQUIRED; ISO8601 formatted string that represents a duration.
	Duration *string

	// REQUIRED; String that represents a Capability URN.
	Name *string

	// REQUIRED; List of key value pairs.
	Parameters []*KeyValuePair

	// REQUIRED; String that represents a selector.
	SelectorID *string

	// REQUIRED; Enum that discriminates between action models.
	Type *string
}

// GetAction implements the ActionClassification interface for type ContinuousAction.
func (c *ContinuousAction) GetAction() *Action {
	return &Action{
		Type: c.Type,
		Name: c.Name,
	}
}

// DelayAction - Model that represents a delay action.
type DelayAction struct {
	// REQUIRED; ISO8601 formatted string that represents a duration.
	Duration *string

	// REQUIRED; String that represents a Capability URN.
	Name *string

	// REQUIRED; Enum that discriminates between action models.
	Type *string
}

// GetAction implements the ActionClassification interface for type DelayAction.
func (d *DelayAction) GetAction() *Action {
	return &Action{
		Type: d.Type,
		Name: d.Name,
	}
}

// DiscreteAction - Model that represents a discrete action.
type DiscreteAction struct {
	// REQUIRED; String that represents a Capability URN.
	Name *string

	// REQUIRED; List of key value pairs.
	Parameters []*KeyValuePair

	// REQUIRED; String that represents a selector.
	SelectorID *string

	// REQUIRED; Enum that discriminates between action models.
	Type *string
}

// GetAction implements the ActionClassification interface for type DiscreteAction.
func (d *DiscreteAction) GetAction() *Action {
	return &Action{
		Type: d.Type,
		Name: d.Name,
	}
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

// Experiment - Model that represents a Experiment resource.
type Experiment struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The properties of the experiment resource.
	Properties *ExperimentProperties

	// The identity of the experiment resource.
	Identity *ResourceIdentity

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata of the experiment resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ExperimentCancelOperationResult - Model that represents the result of a cancel Experiment operation.
type ExperimentCancelOperationResult struct {
	// READ-ONLY; String of the Experiment name.
	Name *string

	// READ-ONLY; URL to retrieve the Experiment status.
	StatusURL *string
}

// ExperimentExecutionActionTargetDetailsError - Model that represents the Experiment action target details error model.
type ExperimentExecutionActionTargetDetailsError struct {
	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error message
	Message *string
}

// ExperimentExecutionActionTargetDetailsProperties - Model that represents the Experiment action target details properties
// model.
type ExperimentExecutionActionTargetDetailsProperties struct {
	// READ-ONLY; The error of the action.
	Error *ExperimentExecutionActionTargetDetailsError

	// READ-ONLY; The status of the execution.
	Status *string

	// READ-ONLY; The target for the action.
	Target *string

	// READ-ONLY; String that represents the completed date time.
	TargetCompletedTime *time.Time

	// READ-ONLY; String that represents the failed date time.
	TargetFailedTime *time.Time
}

// ExperimentExecutionDetails - Model that represents the execution details of a Experiment.
type ExperimentExecutionDetails struct {
	// READ-ONLY; String of the fully qualified resource ID.
	ID *string

	// READ-ONLY; String of the resource name.
	Name *string

	// READ-ONLY; The properties of the experiment execution details.
	Properties *ExperimentExecutionDetailsProperties

	// READ-ONLY; String of the resource type.
	Type *string
}

// ExperimentExecutionDetailsListResult - Model that represents a list of Experiment execution details and a link for pagination.
type ExperimentExecutionDetailsListResult struct {
	// READ-ONLY; URL to retrieve the next page of Experiment execution details.
	NextLink *string

	// READ-ONLY; List of Experiment execution details.
	Value []*ExperimentExecutionDetails
}

// ExperimentExecutionDetailsProperties - Model that represents the Experiment execution details properties model.
type ExperimentExecutionDetailsProperties struct {
	// READ-ONLY; String that represents the created date time.
	CreatedDateTime *time.Time

	// READ-ONLY; The id of the experiment.
	ExperimentID *string

	// READ-ONLY; The reason why the execution failed.
	FailureReason *string

	// READ-ONLY; String that represents the last action date time.
	LastActionDateTime *time.Time

	// READ-ONLY; The information of the experiment run.
	RunInformation *ExperimentExecutionDetailsPropertiesRunInformation

	// READ-ONLY; String that represents the start date time.
	StartDateTime *time.Time

	// READ-ONLY; The value of the status of the experiment execution.
	Status *string

	// READ-ONLY; String that represents the stop date time.
	StopDateTime *time.Time
}

// ExperimentExecutionDetailsPropertiesRunInformation - The information of the experiment run.
type ExperimentExecutionDetailsPropertiesRunInformation struct {
	// READ-ONLY; The steps of the experiment run.
	Steps []*StepStatus
}

// ExperimentListResult - Model that represents a list of Experiment resources and a link for pagination.
type ExperimentListResult struct {
	// READ-ONLY; URL to retrieve the next page of Experiment resources.
	NextLink *string

	// READ-ONLY; List of Experiment resources.
	Value []*Experiment
}

// ExperimentProperties - Model that represents the Experiment properties model.
type ExperimentProperties struct {
	// REQUIRED; List of selectors.
	Selectors []*Selector

	// REQUIRED; List of steps.
	Steps []*Step

	// A boolean value that indicates if experiment should be started on creation or not.
	StartOnCreation *bool
}

// ExperimentStartOperationResult - Model that represents the result of a start Experiment operation.
type ExperimentStartOperationResult struct {
	// READ-ONLY; String of the Experiment name.
	Name *string

	// READ-ONLY; URL to retrieve the Experiment status.
	StatusURL *string
}

// ExperimentStatus - Model that represents the status of a Experiment.
type ExperimentStatus struct {
	// The properties of experiment execution status.
	Properties *ExperimentStatusProperties

	// READ-ONLY; String of the fully qualified resource ID.
	ID *string

	// READ-ONLY; String of the resource name.
	Name *string

	// READ-ONLY; String of the resource type.
	Type *string
}

// ExperimentStatusListResult - Model that represents a list of Experiment statuses and a link for pagination.
type ExperimentStatusListResult struct {
	// READ-ONLY; URL to retrieve the next page of Experiment statuses.
	NextLink *string

	// READ-ONLY; List of Experiment statuses.
	Value []*ExperimentStatus
}

// ExperimentStatusProperties - Model that represents the Experiment status properties model.
type ExperimentStatusProperties struct {
	// READ-ONLY; String that represents the created date time of a Experiment.
	CreatedDateUTC *time.Time

	// READ-ONLY; String that represents the end date time of a Experiment.
	EndDateUTC *time.Time

	// READ-ONLY; String that represents the status of a Experiment.
	Status *string
}

// ExperimentsClientBeginCancelOptions contains the optional parameters for the ExperimentsClient.BeginCancel method.
type ExperimentsClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExperimentsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExperimentsClient.BeginCreateOrUpdate
// method.
type ExperimentsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExperimentsClientDeleteOptions contains the optional parameters for the ExperimentsClient.Delete method.
type ExperimentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientGetExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.GetExecutionDetails
// method.
type ExperimentsClientGetExecutionDetailsOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientGetOptions contains the optional parameters for the ExperimentsClient.Get method.
type ExperimentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientGetStatusOptions contains the optional parameters for the ExperimentsClient.GetStatus method.
type ExperimentsClientGetStatusOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientListAllOptions contains the optional parameters for the ExperimentsClient.NewListAllPager method.
type ExperimentsClientListAllOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
	// Optional value that indicates whether to filter results based on if the Experiment is currently running. If null, then
	// the results will not be filtered.
	Running *bool
}

// ExperimentsClientListAllStatusesOptions contains the optional parameters for the ExperimentsClient.NewListAllStatusesPager
// method.
type ExperimentsClientListAllStatusesOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientListExecutionDetailsOptions contains the optional parameters for the ExperimentsClient.NewListExecutionDetailsPager
// method.
type ExperimentsClientListExecutionDetailsOptions struct {
	// placeholder for future optional parameters
}

// ExperimentsClientListOptions contains the optional parameters for the ExperimentsClient.NewListPager method.
type ExperimentsClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
	// Optional value that indicates whether to filter results based on if the Experiment is currently running. If null, then
	// the results will not be filtered.
	Running *bool
}

// ExperimentsClientStartOptions contains the optional parameters for the ExperimentsClient.Start method.
type ExperimentsClientStartOptions struct {
	// placeholder for future optional parameters
}

// KeyValuePair - A map to describe the settings of an action.
type KeyValuePair struct {
	// REQUIRED; The name of the setting for the action.
	Key *string

	// REQUIRED; The value of the setting for the action.
	Value *string
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

// OperationsClientListAllOptions contains the optional parameters for the OperationsClient.NewListAllPager method.
type OperationsClientListAllOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceIdentity - The managed identity of a resource.
type ResourceIdentity struct {
	// REQUIRED; String of the resource identity type.
	Type *ResourceIdentityType

	// READ-ONLY; GUID that represents the principal ID of this resource identity.
	PrincipalID *string

	// READ-ONLY; GUID that represents the tenant ID of this resource identity.
	TenantID *string
}

// Selector - Model that represents a selector in the Experiment resource.
type Selector struct {
	// REQUIRED; String of the selector ID.
	ID *string

	// REQUIRED; List of Target references.
	Targets []*TargetReference

	// REQUIRED; Enum of the selector type.
	Type *SelectorType
}

// Step - Model that represents a step in the Experiment resource.
type Step struct {
	// REQUIRED; List of branches.
	Branches []*Branch

	// REQUIRED; String of the step name.
	Name *string
}

// StepStatus - Model that represents the a list of branches and branch statuses.
type StepStatus struct {
	// READ-ONLY; The array of branches.
	Branches []*BranchStatus

	// READ-ONLY; The value of the status of the step.
	Status *string

	// READ-ONLY; The id of the step.
	StepID *string

	// READ-ONLY; The name of the step.
	StepName *string
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

// Target - Model that represents a Target resource.
type Target struct {
	// REQUIRED; The properties of the target resource.
	Properties map[string]any

	// Location of the target resource.
	Location *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata of the target resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TargetListResult - Model that represents a list of Target resources and a link for pagination.
type TargetListResult struct {
	// READ-ONLY; URL to retrieve the next page of Target resources.
	NextLink *string

	// READ-ONLY; List of Target resources.
	Value []*Target
}

// TargetReference - Model that represents a reference to a Target in the selector.
type TargetReference struct {
	// REQUIRED; String of the resource ID of a Target resource.
	ID *string

	// CONSTANT; Enum of the Target reference type.
	// Field has constant value "ChaosTarget", any specified value is ignored.
	Type *string
}

// TargetType - Model that represents a Target Type resource.
type TargetType struct {
	// REQUIRED; The properties of the target type resource.
	Properties *TargetTypeProperties

	// Location of the Target Type resource.
	Location *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata properties of the target type resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TargetTypeListResult - Model that represents a list of Target Type resources and a link for pagination.
type TargetTypeListResult struct {
	// READ-ONLY; URL to retrieve the next page of Target Type resources.
	NextLink *string

	// READ-ONLY; List of Target Type resources.
	Value []*TargetType
}

// TargetTypeProperties - Model that represents the base Target Type properties model.
type TargetTypeProperties struct {
	// READ-ONLY; Localized string of the description.
	Description *string

	// READ-ONLY; Localized string of the display name.
	DisplayName *string

	// READ-ONLY; URL to retrieve JSON schema of the Target Type properties.
	PropertiesSchema *string

	// READ-ONLY; List of resource types this Target Type can extend.
	ResourceTypes []*string
}

// TargetTypesClientGetOptions contains the optional parameters for the TargetTypesClient.Get method.
type TargetTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TargetTypesClientListOptions contains the optional parameters for the TargetTypesClient.NewListPager method.
type TargetTypesClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// TargetsClientCreateOrUpdateOptions contains the optional parameters for the TargetsClient.CreateOrUpdate method.
type TargetsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TargetsClientDeleteOptions contains the optional parameters for the TargetsClient.Delete method.
type TargetsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TargetsClientGetOptions contains the optional parameters for the TargetsClient.Get method.
type TargetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TargetsClientListOptions contains the optional parameters for the TargetsClient.NewListPager method.
type TargetsClientListOptions struct {
	// String that sets the continuation token.
	ContinuationToken *string
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}
