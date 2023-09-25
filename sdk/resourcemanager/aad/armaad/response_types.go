//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreate.
type PrivateEndpointConnectionsClientCreateResponse struct {
	// Private endpoint connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// Private endpoint connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPolicyNameResponse contains the response from method PrivateEndpointConnectionsClient.NewListByPolicyNamePager.
type PrivateEndpointConnectionsClientListByPolicyNameResponse struct {
	// A list of private link resources
	PrivateEndpointConnectionListResult
}

// PrivateLinkForAzureAdClientCreateResponse contains the response from method PrivateLinkForAzureAdClient.BeginCreate.
type PrivateLinkForAzureAdClientCreateResponse struct {
	// PrivateLink Policy configuration object.
	PrivateLinkPolicy
}

// PrivateLinkForAzureAdClientDeleteResponse contains the response from method PrivateLinkForAzureAdClient.Delete.
type PrivateLinkForAzureAdClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkForAzureAdClientGetResponse contains the response from method PrivateLinkForAzureAdClient.Get.
type PrivateLinkForAzureAdClientGetResponse struct {
	// PrivateLink Policy configuration object.
	PrivateLinkPolicy
}

// PrivateLinkForAzureAdClientListBySubscriptionResponse contains the response from method PrivateLinkForAzureAdClient.NewListBySubscriptionPager.
type PrivateLinkForAzureAdClientListBySubscriptionResponse struct {
	// A list of private link policies
	PrivateLinkPolicyListResult
}

// PrivateLinkForAzureAdClientListResponse contains the response from method PrivateLinkForAzureAdClient.NewListPager.
type PrivateLinkForAzureAdClientListResponse struct {
	// A list of private link policies
	PrivateLinkPolicyListResult
}

// PrivateLinkForAzureAdClientUpdateResponse contains the response from method PrivateLinkForAzureAdClient.Update.
type PrivateLinkForAzureAdClientUpdateResponse struct {
	// PrivateLink Policy configuration object.
	PrivateLinkPolicy
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkPolicyResponse contains the response from method PrivateLinkResourcesClient.NewListByPrivateLinkPolicyPager.
type PrivateLinkResourcesClientListByPrivateLinkPolicyResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

