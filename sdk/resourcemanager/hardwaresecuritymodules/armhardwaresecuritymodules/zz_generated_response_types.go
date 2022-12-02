//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

// DedicatedHsmClientCreateOrUpdateResponse contains the response from method DedicatedHsmClient.CreateOrUpdate.
type DedicatedHsmClientCreateOrUpdateResponse struct {
	DedicatedHsm
}

// DedicatedHsmClientDeleteResponse contains the response from method DedicatedHsmClient.Delete.
type DedicatedHsmClientDeleteResponse struct {
	// placeholder for future response values
}

// DedicatedHsmClientGetResponse contains the response from method DedicatedHsmClient.Get.
type DedicatedHsmClientGetResponse struct {
	DedicatedHsm
}

// DedicatedHsmClientListByResourceGroupResponse contains the response from method DedicatedHsmClient.ListByResourceGroup.
type DedicatedHsmClientListByResourceGroupResponse struct {
	DedicatedHsmListResult
}

// DedicatedHsmClientListBySubscriptionResponse contains the response from method DedicatedHsmClient.ListBySubscription.
type DedicatedHsmClientListBySubscriptionResponse struct {
	DedicatedHsmListResult
}

// DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method DedicatedHsmClient.ListOutboundNetworkDependenciesEndpoints.
type DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse struct {
	OutboundEnvironmentEndpointCollection
}

// DedicatedHsmClientUpdateResponse contains the response from method DedicatedHsmClient.Update.
type DedicatedHsmClientUpdateResponse struct {
	DedicatedHsm
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	DedicatedHsmOperationListResult
}