//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiembedded

type AzureSKU struct {
	// REQUIRED; SKU name
	Name *AzureSKUName `json:"name,omitempty"`

	// REQUIRED; SKU tier
	Tier *AzureSKUTier `json:"tier,omitempty"`
}

type CheckNameRequest struct {
	// Workspace collection name
	Name *string `json:"name,omitempty"`

	// Resource type
	Type *string `json:"type,omitempty"`
}

type CheckNameResponse struct {
	// Message indicating an unavailable name due to a conflict, or a description of the naming rules that are violated.
	Message *string `json:"message,omitempty"`

	// Specifies a Boolean value that indicates whether the specified Power BI Workspace Collection name is available to use.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// Reason why the workspace collection name cannot be used.
	Reason *CheckNameReason `json:"reason,omitempty"`
}

type CreateWorkspaceCollectionRequest struct {
	// Azure location
	Location *string   `json:"location,omitempty"`
	SKU      *AzureSKU `json:"sku,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`
}

type Display struct {
	// The localized friendly description for the operation as shown to the user. This description should be thorough, yet concise.
	// It will be used in tool-tips and detailed views.
	Description *string `json:"description,omitempty"`

	// The localized friendly name for the operation as shown to the user. This name should be concise (to fit in drop downs),
	// but clear (self-documenting). Use Title Casing and include the entity/resource
	// to which it applies.
	Operation *string `json:"operation,omitempty"`

	// The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX. Default
	// value is 'user,system'
	Origin *string `json:"origin,omitempty"`

	// The localized friendly form of the resource provider name. This form is also expected to include the publisher/company
	// responsible. Use Title Casing. Begin with "Microsoft" for 1st party services.
	Provider *string `json:"provider,omitempty"`

	// The localized friendly form of the resource type related to this action/operation. This form should match the public documentation
	// for the resource provider. Use Title Casing. For examples, refer to
	// the "name" section.
	Resource *string `json:"resource,omitempty"`
}

type Error struct {
	Code    *string        `json:"code,omitempty"`
	Details []*ErrorDetail `json:"details,omitempty"`
	Message *string        `json:"message,omitempty"`
	Target  *string        `json:"target,omitempty"`
}

type ErrorDetail struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// ManagementClientGetAvailableOperationsOptions contains the optional parameters for the ManagementClient.GetAvailableOperations
// method.
type ManagementClientGetAvailableOperationsOptions struct {
	// placeholder for future optional parameters
}

type MigrateWorkspaceCollectionRequest struct {
	Resources []*string `json:"resources,omitempty"`

	// Name of the resource group the Power BI workspace collections will be migrated to.
	TargetResourceGroup *string `json:"targetResourceGroup,omitempty"`
}

type Operation struct {
	Display *Display `json:"display,omitempty"`

	// The name of the operation being performed on this particular object. This name should match the action name that appears
	// in RBAC / the event service.
	Name *string `json:"name,omitempty"`
}

type OperationList struct {
	Value []*Operation `json:"value,omitempty"`
}

type UpdateWorkspaceCollectionRequest struct {
	SKU *AzureSKU `json:"sku,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`
}

type Workspace struct {
	// Workspace id
	ID *string `json:"id,omitempty"`

	// Workspace name
	Name *string `json:"name,omitempty"`

	// Property bag
	Properties interface{} `json:"properties,omitempty"`

	// Resource type
	Type *string `json:"type,omitempty"`
}

type WorkspaceCollection struct {
	// Resource id
	ID *string `json:"id,omitempty"`

	// Azure location
	Location *string `json:"location,omitempty"`

	// Workspace collection name
	Name *string `json:"name,omitempty"`

	// Properties
	Properties interface{} `json:"properties,omitempty"`
	SKU        *AzureSKU   `json:"sku,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`

	// Resource type
	Type *string `json:"type,omitempty"`
}

type WorkspaceCollectionAccessKey struct {
	// Key name
	KeyName *AccessKeyName `json:"keyName,omitempty"`
}

type WorkspaceCollectionAccessKeys struct {
	// Access key 1
	Key1 *string `json:"key1,omitempty"`

	// Access key 2
	Key2 *string `json:"key2,omitempty"`
}

type WorkspaceCollectionList struct {
	Value []*WorkspaceCollection `json:"value,omitempty"`
}

// WorkspaceCollectionsClientBeginDeleteOptions contains the optional parameters for the WorkspaceCollectionsClient.BeginDelete
// method.
type WorkspaceCollectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspaceCollectionsClientCheckNameAvailabilityOptions contains the optional parameters for the WorkspaceCollectionsClient.CheckNameAvailability
// method.
type WorkspaceCollectionsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientCreateOptions contains the optional parameters for the WorkspaceCollectionsClient.Create method.
type WorkspaceCollectionsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientGetAccessKeysOptions contains the optional parameters for the WorkspaceCollectionsClient.GetAccessKeys
// method.
type WorkspaceCollectionsClientGetAccessKeysOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientGetByNameOptions contains the optional parameters for the WorkspaceCollectionsClient.GetByName
// method.
type WorkspaceCollectionsClientGetByNameOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientListByResourceGroupOptions contains the optional parameters for the WorkspaceCollectionsClient.ListByResourceGroup
// method.
type WorkspaceCollectionsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientListBySubscriptionOptions contains the optional parameters for the WorkspaceCollectionsClient.ListBySubscription
// method.
type WorkspaceCollectionsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientMigrateOptions contains the optional parameters for the WorkspaceCollectionsClient.Migrate method.
type WorkspaceCollectionsClientMigrateOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientRegenerateKeyOptions contains the optional parameters for the WorkspaceCollectionsClient.RegenerateKey
// method.
type WorkspaceCollectionsClientRegenerateKeyOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceCollectionsClientUpdateOptions contains the optional parameters for the WorkspaceCollectionsClient.Update method.
type WorkspaceCollectionsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

type WorkspaceList struct {
	Value []*Workspace `json:"value,omitempty"`
}

// WorkspacesClientListOptions contains the optional parameters for the WorkspacesClient.List method.
type WorkspacesClientListOptions struct {
	// placeholder for future optional parameters
}