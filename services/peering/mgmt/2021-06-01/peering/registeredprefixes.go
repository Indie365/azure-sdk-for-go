package peering

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RegisteredPrefixesClient is the peering Client
type RegisteredPrefixesClient struct {
	BaseClient
}

// NewRegisteredPrefixesClient creates an instance of the RegisteredPrefixesClient client.
func NewRegisteredPrefixesClient(subscriptionID string) RegisteredPrefixesClient {
	return NewRegisteredPrefixesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegisteredPrefixesClientWithBaseURI creates an instance of the RegisteredPrefixesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRegisteredPrefixesClientWithBaseURI(baseURI string, subscriptionID string) RegisteredPrefixesClient {
	return RegisteredPrefixesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new registered prefix with the specified name under the given subscription, resource group
// and peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// registeredPrefixName - the name of the registered prefix.
// registeredPrefix - the properties needed to create a registered prefix.
func (client RegisteredPrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix RegisteredPrefix) (result RegisteredPrefix, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredPrefixesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, peeringName, registeredPrefixName, registeredPrefix)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegisteredPrefixesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix RegisteredPrefix) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":          autorest.Encode("path", peeringName),
		"registeredPrefixName": autorest.Encode("path", registeredPrefixName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}", pathParameters),
		autorest.WithJSON(registeredPrefix),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredPrefixesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RegisteredPrefixesClient) CreateOrUpdateResponder(resp *http.Response) (result RegisteredPrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing registered prefix with the specified name under the given subscription, resource group
// and peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// registeredPrefixName - the name of the registered prefix.
func (client RegisteredPrefixesClient) Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredPrefixesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, peeringName, registeredPrefixName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegisteredPrefixesClient) DeletePreparer(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":          autorest.Encode("path", peeringName),
		"registeredPrefixName": autorest.Encode("path", registeredPrefixName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredPrefixesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegisteredPrefixesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an existing registered prefix with the specified name under the given subscription, resource group and
// peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
// registeredPrefixName - the name of the registered prefix.
func (client RegisteredPrefixesClient) Get(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string) (result RegisteredPrefix, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredPrefixesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, peeringName, registeredPrefixName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegisteredPrefixesClient) GetPreparer(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":          autorest.Encode("path", peeringName),
		"registeredPrefixName": autorest.Encode("path", registeredPrefixName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredPrefixesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegisteredPrefixesClient) GetResponder(resp *http.Response) (result RegisteredPrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPeering lists all registered prefixes under the given subscription, resource group and peering.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringName - the name of the peering.
func (client RegisteredPrefixesClient) ListByPeering(ctx context.Context, resourceGroupName string, peeringName string) (result RegisteredPrefixListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredPrefixesClient.ListByPeering")
		defer func() {
			sc := -1
			if result.rplr.Response.Response != nil {
				sc = result.rplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPeeringNextResults
	req, err := client.ListByPeeringPreparer(ctx, resourceGroupName, peeringName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "ListByPeering", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPeeringSender(req)
	if err != nil {
		result.rplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "ListByPeering", resp, "Failure sending request")
		return
	}

	result.rplr, err = client.ListByPeeringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "ListByPeering", resp, "Failure responding to request")
		return
	}
	if result.rplr.hasNextLink() && result.rplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByPeeringPreparer prepares the ListByPeering request.
func (client RegisteredPrefixesClient) ListByPeeringPreparer(ctx context.Context, resourceGroupName string, peeringName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringName":       autorest.Encode("path", peeringName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPeeringSender sends the ListByPeering request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredPrefixesClient) ListByPeeringSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPeeringResponder handles the response to the ListByPeering request. The method always
// closes the http.Response Body.
func (client RegisteredPrefixesClient) ListByPeeringResponder(resp *http.Response) (result RegisteredPrefixListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPeeringNextResults retrieves the next set of results, if any.
func (client RegisteredPrefixesClient) listByPeeringNextResults(ctx context.Context, lastResults RegisteredPrefixListResult) (result RegisteredPrefixListResult, err error) {
	req, err := lastResults.registeredPrefixListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "listByPeeringNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPeeringSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "listByPeeringNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPeeringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.RegisteredPrefixesClient", "listByPeeringNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPeeringComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegisteredPrefixesClient) ListByPeeringComplete(ctx context.Context, resourceGroupName string, peeringName string) (result RegisteredPrefixListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredPrefixesClient.ListByPeering")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPeering(ctx, resourceGroupName, peeringName)
	return
}
