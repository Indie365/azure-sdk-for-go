//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

// AccountsClientBeginCreateOptions contains the optional parameters for the AccountsClient.BeginCreate method.
type AccountsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
type AccountsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
type AccountsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
// method.
type AccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListKeysOptions contains the optional parameters for the AccountsClient.ListKeys method.
type AccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListModelsOptions contains the optional parameters for the AccountsClient.NewListModelsPager method.
type AccountsClientListModelsOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListOptions contains the optional parameters for the AccountsClient.NewListPager method.
type AccountsClientListOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListSKUsOptions contains the optional parameters for the AccountsClient.ListSKUs method.
type AccountsClientListSKUsOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListUsagesOptions contains the optional parameters for the AccountsClient.ListUsages method.
type AccountsClientListUsagesOptions struct {
	// An OData filter expression that describes a subset of usages to return. The supported parameter is name.value (name of
// the metric, can have an or of multiple names).
	Filter *string
}

// AccountsClientRegenerateKeyOptions contains the optional parameters for the AccountsClient.RegenerateKey method.
type AccountsClientRegenerateKeyOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientBeginCreateOrUpdateAssociationOptions contains the optional parameters for the CommitmentPlansClient.BeginCreateOrUpdateAssociation
// method.
type CommitmentPlansClientBeginCreateOrUpdateAssociationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommitmentPlansClientBeginCreateOrUpdatePlanOptions contains the optional parameters for the CommitmentPlansClient.BeginCreateOrUpdatePlan
// method.
type CommitmentPlansClientBeginCreateOrUpdatePlanOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommitmentPlansClientBeginDeleteAssociationOptions contains the optional parameters for the CommitmentPlansClient.BeginDeleteAssociation
// method.
type CommitmentPlansClientBeginDeleteAssociationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommitmentPlansClientBeginDeleteOptions contains the optional parameters for the CommitmentPlansClient.BeginDelete method.
type CommitmentPlansClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommitmentPlansClientBeginDeletePlanOptions contains the optional parameters for the CommitmentPlansClient.BeginDeletePlan
// method.
type CommitmentPlansClientBeginDeletePlanOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommitmentPlansClientBeginUpdatePlanOptions contains the optional parameters for the CommitmentPlansClient.BeginUpdatePlan
// method.
type CommitmentPlansClientBeginUpdatePlanOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommitmentPlansClientCreateOrUpdateOptions contains the optional parameters for the CommitmentPlansClient.CreateOrUpdate
// method.
type CommitmentPlansClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientGetAssociationOptions contains the optional parameters for the CommitmentPlansClient.GetAssociation
// method.
type CommitmentPlansClientGetAssociationOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientGetOptions contains the optional parameters for the CommitmentPlansClient.Get method.
type CommitmentPlansClientGetOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientGetPlanOptions contains the optional parameters for the CommitmentPlansClient.GetPlan method.
type CommitmentPlansClientGetPlanOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientListAssociationsOptions contains the optional parameters for the CommitmentPlansClient.NewListAssociationsPager
// method.
type CommitmentPlansClientListAssociationsOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientListOptions contains the optional parameters for the CommitmentPlansClient.NewListPager method.
type CommitmentPlansClientListOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientListPlansByResourceGroupOptions contains the optional parameters for the CommitmentPlansClient.NewListPlansByResourceGroupPager
// method.
type CommitmentPlansClientListPlansByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CommitmentPlansClientListPlansBySubscriptionOptions contains the optional parameters for the CommitmentPlansClient.NewListPlansBySubscriptionPager
// method.
type CommitmentPlansClientListPlansBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// CommitmentTiersClientListOptions contains the optional parameters for the CommitmentTiersClient.NewListPager method.
type CommitmentTiersClientListOptions struct {
	// placeholder for future optional parameters
}

// DeletedAccountsClientBeginPurgeOptions contains the optional parameters for the DeletedAccountsClient.BeginPurge method.
type DeletedAccountsClientBeginPurgeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeletedAccountsClientGetOptions contains the optional parameters for the DeletedAccountsClient.Get method.
type DeletedAccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DeletedAccountsClientListOptions contains the optional parameters for the DeletedAccountsClient.NewListPager method.
type DeletedAccountsClientListOptions struct {
	// placeholder for future optional parameters
}

// DeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the DeploymentsClient.BeginCreateOrUpdate
// method.
type DeploymentsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientBeginDeleteOptions contains the optional parameters for the DeploymentsClient.BeginDelete method.
type DeploymentsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientGetOptions contains the optional parameters for the DeploymentsClient.Get method.
type DeploymentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DeploymentsClientListOptions contains the optional parameters for the DeploymentsClient.NewListPager method.
type DeploymentsClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientCheckDomainAvailabilityOptions contains the optional parameters for the ManagementClient.CheckDomainAvailability
// method.
type ManagementClientCheckDomainAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientCheckSKUAvailabilityOptions contains the optional parameters for the ManagementClient.CheckSKUAvailability
// method.
type ManagementClientCheckSKUAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ModelsClientListOptions contains the optional parameters for the ModelsClient.NewListPager method.
type ModelsClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
type PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.List
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkResourcesClient.List method.
type PrivateLinkResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// ResourceSKUsClientListOptions contains the optional parameters for the ResourceSKUsClient.NewListPager method.
type ResourceSKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// UsagesClientListOptions contains the optional parameters for the UsagesClient.NewListPager method.
type UsagesClientListOptions struct {
	// An OData filter expression that describes a subset of usages to return. The supported parameter is name.value (name of
// the metric, can have an or of multiple names).
	Filter *string
}

