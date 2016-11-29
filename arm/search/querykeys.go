package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/satori/uuid"
	"net/http"
)

// QueryKeysClient is the client that can be used to manage Azure Search
// services and API keys.
type QueryKeysClient struct {
	ManagementClient
}

// NewQueryKeysClient creates an instance of the QueryKeysClient client.
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return NewQueryKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryKeysClientWithBaseURI creates an instance of the QueryKeysClient
// client.
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return QueryKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create generates a new query key for the specified Search service. You can
// create up to 50 query keys per service.
//
// resourceGroupName is the name of the resource group within the current
// subscription. You can obtain this value from the Azure Resource Manager
// API or the portal. searchServiceName is the name of the Azure Search
// service associated with the specified resource group. name is the name of
// the new query API key. xMsClientRequestID is a client-generated GUID value
// that identifies this request. If specified, this will be included in
// response information as a way to track the request.
func (client QueryKeysClient) Create(resourceGroupName string, searchServiceName string, name string, xMsClientRequestID *uuid.UUID) (result QueryKey, err error) {
	req, err := client.CreatePreparer(resourceGroupName, searchServiceName, name, xMsClientRequestID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client QueryKeysClient) CreatePreparer(resourceGroupName string, searchServiceName string, name string, xMsClientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/createQueryKey/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if xMsClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(xMsClientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) CreateResponder(resp *http.Response) (result QueryKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified query key. Unlike admin keys, query keys are
// not regenerated. The process for regenerating a query key is to delete and
// then recreate it.
//
// resourceGroupName is the name of the resource group within the current
// subscription. You can obtain this value from the Azure Resource Manager
// API or the portal. searchServiceName is the name of the Azure Search
// service associated with the specified resource group. key is the query key
// to be deleted. Query keys are identified by value, not by name.
// xMsClientRequestID is a client-generated GUID value that identifies this
// request. If specified, this will be included in response information as a
// way to track the request.
func (client QueryKeysClient) Delete(resourceGroupName string, searchServiceName string, key string, xMsClientRequestID *uuid.UUID) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, searchServiceName, key, xMsClientRequestID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client QueryKeysClient) DeletePreparer(resourceGroupName string, searchServiceName string, key string, xMsClientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"key":               autorest.Encode("path", key),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/deleteQueryKey/{key}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if xMsClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(xMsClientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListBySearchService returns the list of query API keys for the given Azure
// Search service.
//
// resourceGroupName is the name of the resource group within the current
// subscription. You can obtain this value from the Azure Resource Manager
// API or the portal. searchServiceName is the name of the Azure Search
// service associated with the specified resource group. xMsClientRequestID
// is a client-generated GUID value that identifies this request. If
// specified, this will be included in response information as a way to track
// the request.
func (client QueryKeysClient) ListBySearchService(resourceGroupName string, searchServiceName string, xMsClientRequestID *uuid.UUID) (result ListQueryKeysResult, err error) {
	req, err := client.ListBySearchServicePreparer(resourceGroupName, searchServiceName, xMsClientRequestID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "ListBySearchService", nil, "Failure preparing request")
	}

	resp, err := client.ListBySearchServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "ListBySearchService", resp, "Failure sending request")
	}

	result, err = client.ListBySearchServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "ListBySearchService", resp, "Failure responding to request")
	}

	return
}

// ListBySearchServicePreparer prepares the ListBySearchService request.
func (client QueryKeysClient) ListBySearchServicePreparer(resourceGroupName string, searchServiceName string, xMsClientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listQueryKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if xMsClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(xMsClientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// ListBySearchServiceSender sends the ListBySearchService request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) ListBySearchServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySearchServiceResponder handles the response to the ListBySearchService request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) ListBySearchServiceResponder(resp *http.Response) (result ListQueryKeysResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
