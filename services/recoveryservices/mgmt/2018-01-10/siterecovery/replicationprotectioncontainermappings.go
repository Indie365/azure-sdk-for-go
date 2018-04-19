package siterecovery

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ReplicationProtectionContainerMappingsClient is the client for the ReplicationProtectionContainerMappings methods of
// the Siterecovery service.
type ReplicationProtectionContainerMappingsClient struct {
	BaseClient
}

// NewReplicationProtectionContainerMappingsClient creates an instance of the
// ReplicationProtectionContainerMappingsClient client.
func NewReplicationProtectionContainerMappingsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectionContainerMappingsClient {
	return NewReplicationProtectionContainerMappingsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationProtectionContainerMappingsClientWithBaseURI creates an instance of the
// ReplicationProtectionContainerMappingsClient client.
func NewReplicationProtectionContainerMappingsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectionContainerMappingsClient {
	return ReplicationProtectionContainerMappingsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to create a protection container mapping.
//
// fabricName is fabric name. protectionContainerName is protection container name. mappingName is protection
// container mapping name. creationInput is mapping creation input.
func (client ReplicationProtectionContainerMappingsClient) Create(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput) (result ReplicationProtectionContainerMappingsCreateFuture, err error) {
	req, err := client.CreatePreparer(ctx, fabricName, protectionContainerName, mappingName, creationInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationProtectionContainerMappingsClient) CreatePreparer(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"mappingName":             autorest.Encode("path", mappingName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", pathParameters),
		autorest.WithJSON(creationInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) CreateSender(req *http.Request) (future ReplicationProtectionContainerMappingsCreateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) CreateResponder(resp *http.Response) (result ProtectionContainerMapping, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete or remove a protection container mapping.
//
// fabricName is fabric name. protectionContainerName is protection container name. mappingName is protection
// container mapping name. removalInput is removal input.
func (client ReplicationProtectionContainerMappingsClient) Delete(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput) (result ReplicationProtectionContainerMappingsDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, fabricName, protectionContainerName, mappingName, removalInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationProtectionContainerMappingsClient) DeletePreparer(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"mappingName":             autorest.Encode("path", mappingName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}/remove", pathParameters),
		autorest.WithJSON(removalInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) DeleteSender(req *http.Request) (future ReplicationProtectionContainerMappingsDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of a protection container mapping.
//
// fabricName is fabric name. protectionContainerName is protection container name. mappingName is protection
// Container mapping name.
func (client ReplicationProtectionContainerMappingsClient) Get(ctx context.Context, fabricName string, protectionContainerName string, mappingName string) (result ProtectionContainerMapping, err error) {
	req, err := client.GetPreparer(ctx, fabricName, protectionContainerName, mappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationProtectionContainerMappingsClient) GetPreparer(ctx context.Context, fabricName string, protectionContainerName string, mappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"mappingName":             autorest.Encode("path", mappingName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) GetResponder(resp *http.Response) (result ProtectionContainerMapping, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the protection container mappings in the vault.
func (client ReplicationProtectionContainerMappingsClient) List(ctx context.Context) (result ProtectionContainerMappingCollectionPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pcmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "List", resp, "Failure sending request")
		return
	}

	result.pcmc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationProtectionContainerMappingsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) ListResponder(resp *http.Response) (result ProtectionContainerMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationProtectionContainerMappingsClient) listNextResults(lastResults ProtectionContainerMappingCollection) (result ProtectionContainerMappingCollection, err error) {
	req, err := lastResults.protectionContainerMappingCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationProtectionContainerMappingsClient) ListComplete(ctx context.Context) (result ProtectionContainerMappingCollectionIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByReplicationProtectionContainers lists the protection container mappings for a protection container.
//
// fabricName is fabric name. protectionContainerName is protection container name.
func (client ReplicationProtectionContainerMappingsClient) ListByReplicationProtectionContainers(ctx context.Context, fabricName string, protectionContainerName string) (result ProtectionContainerMappingCollectionPage, err error) {
	result.fn = client.listByReplicationProtectionContainersNextResults
	req, err := client.ListByReplicationProtectionContainersPreparer(ctx, fabricName, protectionContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "ListByReplicationProtectionContainers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationProtectionContainersSender(req)
	if err != nil {
		result.pcmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "ListByReplicationProtectionContainers", resp, "Failure sending request")
		return
	}

	result.pcmc, err = client.ListByReplicationProtectionContainersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "ListByReplicationProtectionContainers", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationProtectionContainersPreparer prepares the ListByReplicationProtectionContainers request.
func (client ReplicationProtectionContainerMappingsClient) ListByReplicationProtectionContainersPreparer(ctx context.Context, fabricName string, protectionContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationProtectionContainersSender sends the ListByReplicationProtectionContainers request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) ListByReplicationProtectionContainersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationProtectionContainersResponder handles the response to the ListByReplicationProtectionContainers request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) ListByReplicationProtectionContainersResponder(resp *http.Response) (result ProtectionContainerMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationProtectionContainersNextResults retrieves the next set of results, if any.
func (client ReplicationProtectionContainerMappingsClient) listByReplicationProtectionContainersNextResults(lastResults ProtectionContainerMappingCollection) (result ProtectionContainerMappingCollection, err error) {
	req, err := lastResults.protectionContainerMappingCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "listByReplicationProtectionContainersNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationProtectionContainersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "listByReplicationProtectionContainersNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationProtectionContainersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "listByReplicationProtectionContainersNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationProtectionContainersComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationProtectionContainerMappingsClient) ListByReplicationProtectionContainersComplete(ctx context.Context, fabricName string, protectionContainerName string) (result ProtectionContainerMappingCollectionIterator, err error) {
	result.page, err = client.ListByReplicationProtectionContainers(ctx, fabricName, protectionContainerName)
	return
}

// Purge the operation to purge(force delete) a protection container mapping
//
// fabricName is fabric name. protectionContainerName is protection container name. mappingName is protection
// container mapping name.
func (client ReplicationProtectionContainerMappingsClient) Purge(ctx context.Context, fabricName string, protectionContainerName string, mappingName string) (result ReplicationProtectionContainerMappingsPurgeFuture, err error) {
	req, err := client.PurgePreparer(ctx, fabricName, protectionContainerName, mappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Purge", nil, "Failure preparing request")
		return
	}

	result, err = client.PurgeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Purge", result.Response(), "Failure sending request")
		return
	}

	return
}

// PurgePreparer prepares the Purge request.
func (client ReplicationProtectionContainerMappingsClient) PurgePreparer(ctx context.Context, fabricName string, protectionContainerName string, mappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"mappingName":             autorest.Encode("path", mappingName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) PurgeSender(req *http.Request) (future ReplicationProtectionContainerMappingsPurgeFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) PurgeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update the operation to update protection container mapping.
//
// fabricName is fabric name. protectionContainerName is protection container name. mappingName is protection
// container mapping name. updateInput is mapping update input.
func (client ReplicationProtectionContainerMappingsClient) Update(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput) (result ReplicationProtectionContainerMappingsUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, fabricName, protectionContainerName, mappingName, updateInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationProtectionContainerMappingsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ReplicationProtectionContainerMappingsClient) UpdatePreparer(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"mappingName":             autorest.Encode("path", mappingName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", pathParameters),
		autorest.WithJSON(updateInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectionContainerMappingsClient) UpdateSender(req *http.Request) (future ReplicationProtectionContainerMappingsUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ReplicationProtectionContainerMappingsClient) UpdateResponder(resp *http.Response) (result ProtectionContainerMapping, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
