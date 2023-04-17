//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpurview

// AccountsClientAddRootCollectionAdminResponse contains the response from method AccountsClient.AddRootCollectionAdmin.
type AccountsClientAddRootCollectionAdminResponse struct {
	// placeholder for future response values
}

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// AccountsClientCreateOrUpdateResponse contains the response from method AccountsClient.BeginCreateOrUpdate.
type AccountsClientCreateOrUpdateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	AccountList
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	AccountList
}

// AccountsClientListKeysResponse contains the response from method AccountsClient.ListKeys.
type AccountsClientListKeysResponse struct {
	AccessKeys
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	Account
}

// DefaultAccountsClientGetResponse contains the response from method DefaultAccountsClient.Get.
type DefaultAccountsClientGetResponse struct {
	DefaultAccountPayload
}

// DefaultAccountsClientRemoveResponse contains the response from method DefaultAccountsClient.Remove.
type DefaultAccountsClientRemoveResponse struct {
	// placeholder for future response values
}

// DefaultAccountsClientSetResponse contains the response from method DefaultAccountsClient.Set.
type DefaultAccountsClientSetResponse struct {
	DefaultAccountPayload
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationList
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByAccountResponse contains the response from method PrivateEndpointConnectionsClient.NewListByAccountPager.
type PrivateEndpointConnectionsClientListByAccountResponse struct {
	PrivateEndpointConnectionList
}

// PrivateLinkResourcesClientGetByGroupIDResponse contains the response from method PrivateLinkResourcesClient.GetByGroupID.
type PrivateLinkResourcesClientGetByGroupIDResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByAccountResponse contains the response from method PrivateLinkResourcesClient.NewListByAccountPager.
type PrivateLinkResourcesClientListByAccountResponse struct {
	PrivateLinkResourceList
}
