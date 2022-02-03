//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridconnectivity

import "net/http"

// EndpointsClientCreateOrUpdateResponse contains the response from method EndpointsClient.CreateOrUpdate.
type EndpointsClientCreateOrUpdateResponse struct {
	EndpointsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointsClientCreateOrUpdateResult contains the result from method EndpointsClient.CreateOrUpdate.
type EndpointsClientCreateOrUpdateResult struct {
	EndpointResource
}

// EndpointsClientDeleteResponse contains the response from method EndpointsClient.Delete.
type EndpointsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointsClientGetResponse contains the response from method EndpointsClient.Get.
type EndpointsClientGetResponse struct {
	EndpointsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointsClientGetResult contains the result from method EndpointsClient.Get.
type EndpointsClientGetResult struct {
	EndpointResource
}

// EndpointsClientListCredentialsResponse contains the response from method EndpointsClient.ListCredentials.
type EndpointsClientListCredentialsResponse struct {
	EndpointsClientListCredentialsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointsClientListCredentialsResult contains the result from method EndpointsClient.ListCredentials.
type EndpointsClientListCredentialsResult struct {
	EndpointAccessResource
}

// EndpointsClientListResponse contains the response from method EndpointsClient.List.
type EndpointsClientListResponse struct {
	EndpointsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointsClientListResult contains the result from method EndpointsClient.List.
type EndpointsClientListResult struct {
	EndpointsList
}

// EndpointsClientUpdateResponse contains the response from method EndpointsClient.Update.
type EndpointsClientUpdateResponse struct {
	EndpointsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EndpointsClientUpdateResult contains the result from method EndpointsClient.Update.
type EndpointsClientUpdateResult struct {
	EndpointResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}
