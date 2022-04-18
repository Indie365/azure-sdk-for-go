//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.Create.
type PrivateEndpointConnectionsClientCreateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPolicyNameResponse contains the response from method PrivateEndpointConnectionsClient.ListByPolicyName.
type PrivateEndpointConnectionsClientListByPolicyNameResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkForAzureAdClientCreateResponse contains the response from method PrivateLinkForAzureAdClient.Create.
type PrivateLinkForAzureAdClientCreateResponse struct {
	PrivateLinkPolicy
}

// PrivateLinkForAzureAdClientDeleteResponse contains the response from method PrivateLinkForAzureAdClient.Delete.
type PrivateLinkForAzureAdClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkForAzureAdClientGetResponse contains the response from method PrivateLinkForAzureAdClient.Get.
type PrivateLinkForAzureAdClientGetResponse struct {
	PrivateLinkPolicy
}

// PrivateLinkForAzureAdClientListBySubscriptionResponse contains the response from method PrivateLinkForAzureAdClient.ListBySubscription.
type PrivateLinkForAzureAdClientListBySubscriptionResponse struct {
	PrivateLinkPolicyListResult
}

// PrivateLinkForAzureAdClientListResponse contains the response from method PrivateLinkForAzureAdClient.List.
type PrivateLinkForAzureAdClientListResponse struct {
	PrivateLinkPolicyListResult
}

// PrivateLinkForAzureAdClientUpdateResponse contains the response from method PrivateLinkForAzureAdClient.Update.
type PrivateLinkForAzureAdClientUpdateResponse struct {
	PrivateLinkPolicy
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkPolicyResponse contains the response from method PrivateLinkResourcesClient.ListByPrivateLinkPolicy.
type PrivateLinkResourcesClientListByPrivateLinkPolicyResponse struct {
	PrivateLinkResourceListResult
}
