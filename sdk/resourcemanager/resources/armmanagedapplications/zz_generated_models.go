//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// Application - Information about managed application.
type Application struct {
	GenericResource
	// REQUIRED; The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
	Kind *string `json:"kind,omitempty"`

	// REQUIRED; The managed application properties.
	Properties *ApplicationProperties `json:"properties,omitempty"`

	// The plan information.
	Plan *Plan `json:"plan,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Application.
func (a Application) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	a.GenericResource.marshalInternal(objectMap)
	populate(objectMap, "kind", a.Kind)
	populate(objectMap, "plan", a.Plan)
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// ApplicationArtifact - Managed application artifact.
type ApplicationArtifact struct {
	// The managed application artifact name.
	Name *string `json:"name,omitempty"`

	// The managed application artifact type.
	Type *ApplicationArtifactType `json:"type,omitempty"`

	// The managed application artifact blob uri.
	URI *string `json:"uri,omitempty"`
}

// ApplicationClientListOperationsOptions contains the optional parameters for the ApplicationClient.ListOperations method.
type ApplicationClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinition - Information about managed application definition.
type ApplicationDefinition struct {
	GenericResource
	// REQUIRED; The managed application definition properties.
	Properties *ApplicationDefinitionProperties `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationDefinition.
func (a ApplicationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	a.GenericResource.marshalInternal(objectMap)
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// ApplicationDefinitionListResult - List of managed application definitions.
type ApplicationDefinitionListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of managed application definitions.
	Value []*ApplicationDefinition `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationDefinitionListResult.
func (a ApplicationDefinitionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// ApplicationDefinitionProperties - The managed application definition properties.
type ApplicationDefinitionProperties struct {
	// REQUIRED; The managed application provider authorizations.
	Authorizations []*ApplicationProviderAuthorization `json:"authorizations,omitempty"`

	// REQUIRED; The managed application lock level.
	LockLevel *ApplicationLockLevel `json:"lockLevel,omitempty"`

	// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a
	// managed application from a managed application
	// definition.
	Artifacts []*ApplicationArtifact `json:"artifacts,omitempty"`

	// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
	CreateUIDefinition map[string]interface{} `json:"createUiDefinition,omitempty"`

	// The managed application definition description.
	Description *string `json:"description,omitempty"`

	// The managed application definition display name.
	DisplayName *string `json:"displayName,omitempty"`

	// A value indicating whether the package is enabled or not.
	IsEnabled *string `json:"isEnabled,omitempty"`

	// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
	MainTemplate map[string]interface{} `json:"mainTemplate,omitempty"`

	// The managed application definition package file Uri. Use this element
	PackageFileURI *string `json:"packageFileUri,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationDefinitionProperties.
func (a ApplicationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifacts", a.Artifacts)
	populate(objectMap, "authorizations", a.Authorizations)
	populate(objectMap, "createUiDefinition", a.CreateUIDefinition)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "displayName", a.DisplayName)
	populate(objectMap, "isEnabled", a.IsEnabled)
	populate(objectMap, "lockLevel", a.LockLevel)
	populate(objectMap, "mainTemplate", a.MainTemplate)
	populate(objectMap, "packageFileUri", a.PackageFileURI)
	return json.Marshal(objectMap)
}

// ApplicationDefinitionsBeginCreateOrUpdateByIDOptions contains the optional parameters for the ApplicationDefinitions.BeginCreateOrUpdateByID method.
type ApplicationDefinitionsBeginCreateOrUpdateByIDOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinitionsBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationDefinitions.BeginCreateOrUpdate method.
type ApplicationDefinitionsBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinitionsBeginDeleteByIDOptions contains the optional parameters for the ApplicationDefinitions.BeginDeleteByID method.
type ApplicationDefinitionsBeginDeleteByIDOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinitionsBeginDeleteOptions contains the optional parameters for the ApplicationDefinitions.BeginDelete method.
type ApplicationDefinitionsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinitionsGetByIDOptions contains the optional parameters for the ApplicationDefinitions.GetByID method.
type ApplicationDefinitionsGetByIDOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinitionsGetOptions contains the optional parameters for the ApplicationDefinitions.Get method.
type ApplicationDefinitionsGetOptions struct {
	// placeholder for future optional parameters
}

// ApplicationDefinitionsListByResourceGroupOptions contains the optional parameters for the ApplicationDefinitions.ListByResourceGroup method.
type ApplicationDefinitionsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ApplicationListResult - List of managed applications.
type ApplicationListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of managed applications.
	Value []*Application `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationListResult.
func (a ApplicationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// ApplicationPatchable - Information about managed application.
type ApplicationPatchable struct {
	GenericResource
	// The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
	Kind *string `json:"kind,omitempty"`

	// The plan information.
	Plan *PlanPatchable `json:"plan,omitempty"`

	// The managed application properties.
	Properties *ApplicationPropertiesPatchable `json:"properties,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationPatchable.
func (a ApplicationPatchable) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	a.GenericResource.marshalInternal(objectMap)
	populate(objectMap, "kind", a.Kind)
	populate(objectMap, "plan", a.Plan)
	populate(objectMap, "properties", a.Properties)
	return json.Marshal(objectMap)
}

// ApplicationProperties - The managed application properties.
type ApplicationProperties struct {
	// REQUIRED; The managed resource group Id.
	ManagedResourceGroupID *string `json:"managedResourceGroupId,omitempty"`

	// The fully qualified path of managed application definition Id.
	ApplicationDefinitionID *string `json:"applicationDefinitionId,omitempty"`

	// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// READ-ONLY; Name and value pairs that define the managed application outputs.
	Outputs map[string]interface{} `json:"outputs,omitempty" azure:"ro"`

	// READ-ONLY; The managed application provisioning state.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ApplicationPropertiesPatchable - The managed application properties.
type ApplicationPropertiesPatchable struct {
	// The fully qualified path of managed application definition Id.
	ApplicationDefinitionID *string `json:"applicationDefinitionId,omitempty"`

	// The managed resource group Id.
	ManagedResourceGroupID *string `json:"managedResourceGroupId,omitempty"`

	// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// READ-ONLY; Name and value pairs that define the managed application outputs.
	Outputs map[string]interface{} `json:"outputs,omitempty" azure:"ro"`

	// READ-ONLY; The managed application provisioning state.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ApplicationProviderAuthorization - The managed application provider authorization.
type ApplicationProviderAuthorization struct {
	// REQUIRED; The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the managed application resources.
	PrincipalID *string `json:"principalId,omitempty"`

	// REQUIRED; The provider's role definition identifier. This role will define all the permissions that the provider must have on the managed application's
	// container resource group. This role definition cannot have
	// permission to delete the resource group.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
}

// ApplicationsBeginCreateOrUpdateByIDOptions contains the optional parameters for the Applications.BeginCreateOrUpdateByID method.
type ApplicationsBeginCreateOrUpdateByIDOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsBeginCreateOrUpdateOptions contains the optional parameters for the Applications.BeginCreateOrUpdate method.
type ApplicationsBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsBeginDeleteByIDOptions contains the optional parameters for the Applications.BeginDeleteByID method.
type ApplicationsBeginDeleteByIDOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsBeginDeleteOptions contains the optional parameters for the Applications.BeginDelete method.
type ApplicationsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsGetByIDOptions contains the optional parameters for the Applications.GetByID method.
type ApplicationsGetByIDOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsGetOptions contains the optional parameters for the Applications.Get method.
type ApplicationsGetOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsListByResourceGroupOptions contains the optional parameters for the Applications.ListByResourceGroup method.
type ApplicationsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsListBySubscriptionOptions contains the optional parameters for the Applications.ListBySubscription method.
type ApplicationsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsUpdateByIDOptions contains the optional parameters for the Applications.UpdateByID method.
type ApplicationsUpdateByIDOptions struct {
	// Parameters supplied to update an existing managed application.
	Parameters *Application
}

// ApplicationsUpdateOptions contains the optional parameters for the Applications.Update method.
type ApplicationsUpdateOptions struct {
	// Parameters supplied to update an existing managed application.
	Parameters *ApplicationPatchable
}

// ErrorResponse - Error response indicates managed application is not able to process the incoming request. The reason is provided in the error message.
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// Error code.
	ErrorCode *string `json:"errorCode,omitempty"`

	// Error message indicating why the operation failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Http status code.
	HTTPStatus *string `json:"httpStatus,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// GenericResource - Resource information.
type GenericResource struct {
	Resource
	// The identity of the resource.
	Identity *Identity `json:"identity,omitempty"`

	// ID of the resource that manages this resource.
	ManagedBy *string `json:"managedBy,omitempty"`

	// The SKU of the resource.
	SKU *SKU `json:"sku,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type GenericResource.
func (g GenericResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	g.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (g GenericResource) marshalInternal(objectMap map[string]interface{}) {
	g.Resource.marshalInternal(objectMap)
	populate(objectMap, "identity", g.Identity)
	populate(objectMap, "managedBy", g.ManagedBy)
	populate(objectMap, "sku", g.SKU)
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// Operation - Microsoft.Solutions operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.Solutions
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Application, JitRequest, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Microsoft.Solutions operations. It contains a list of operations and a URL link to get the next set
// of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Microsoft.Solutions operations.
	Value []*Operation `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// Plan for the managed application.
type Plan struct {
	// REQUIRED; The plan name.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The product code.
	Product *string `json:"product,omitempty"`

	// REQUIRED; The publisher ID.
	Publisher *string `json:"publisher,omitempty"`

	// REQUIRED; The plan's version.
	Version *string `json:"version,omitempty"`

	// The promotion code.
	PromotionCode *string `json:"promotionCode,omitempty"`
}

// PlanPatchable - Plan for the managed application.
type PlanPatchable struct {
	// The plan name.
	Name *string `json:"name,omitempty"`

	// The product code.
	Product *string `json:"product,omitempty"`

	// The promotion code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// The publisher ID.
	Publisher *string `json:"publisher,omitempty"`

	// The plan's version.
	Version *string `json:"version,omitempty"`
}

// Resource information.
type Resource struct {
	// Resource location
	Location *string `json:"location,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Resource ID
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
}

// SKU for the resource.
type SKU struct {
	// REQUIRED; The SKU name.
	Name *string `json:"name,omitempty"`

	// The SKU capacity.
	Capacity *int32 `json:"capacity,omitempty"`

	// The SKU family.
	Family *string `json:"family,omitempty"`

	// The SKU model.
	Model *string `json:"model,omitempty"`

	// The SKU size.
	Size *string `json:"size,omitempty"`

	// The SKU tier.
	Tier *string `json:"tier,omitempty"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
