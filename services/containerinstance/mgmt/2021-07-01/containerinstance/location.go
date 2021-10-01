package containerinstance

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

// LocationClient is the client for the Location methods of the Containerinstance service.
type LocationClient struct {
	BaseClient
}

// NewLocationClient creates an instance of the LocationClient client.
func NewLocationClient(subscriptionID string) LocationClient {
	return NewLocationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationClientWithBaseURI creates an instance of the LocationClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string) LocationClient {
	return LocationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListCachedImages get the list of cached images on specific OS type for a subscription in a region.
// Parameters:
// location - the identifier for the physical azure location.
func (client LocationClient) ListCachedImages(ctx context.Context, location string) (result CachedImagesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.ListCachedImages")
		defer func() {
			sc := -1
			if result.cilr.Response.Response != nil {
				sc = result.cilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listCachedImagesNextResults
	req, err := client.ListCachedImagesPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListCachedImages", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCachedImagesSender(req)
	if err != nil {
		result.cilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListCachedImages", resp, "Failure sending request")
		return
	}

	result.cilr, err = client.ListCachedImagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListCachedImages", resp, "Failure responding to request")
		return
	}
	if result.cilr.hasNextLink() && result.cilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListCachedImagesPreparer prepares the ListCachedImages request.
func (client LocationClient) ListCachedImagesPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/cachedImages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCachedImagesSender sends the ListCachedImages request. The method will close the
// http.Response Body if it receives an error.
func (client LocationClient) ListCachedImagesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCachedImagesResponder handles the response to the ListCachedImages request. The method always
// closes the http.Response Body.
func (client LocationClient) ListCachedImagesResponder(resp *http.Response) (result CachedImagesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listCachedImagesNextResults retrieves the next set of results, if any.
func (client LocationClient) listCachedImagesNextResults(ctx context.Context, lastResults CachedImagesListResult) (result CachedImagesListResult, err error) {
	req, err := lastResults.cachedImagesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerinstance.LocationClient", "listCachedImagesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListCachedImagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerinstance.LocationClient", "listCachedImagesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListCachedImagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "listCachedImagesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListCachedImagesComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocationClient) ListCachedImagesComplete(ctx context.Context, location string) (result CachedImagesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.ListCachedImages")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListCachedImages(ctx, location)
	return
}

// ListCapabilities get the list of CPU/memory/GPU capabilities of a region.
// Parameters:
// location - the identifier for the physical azure location.
func (client LocationClient) ListCapabilities(ctx context.Context, location string) (result CapabilitiesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.ListCapabilities")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listCapabilitiesNextResults
	req, err := client.ListCapabilitiesPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListCapabilities", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCapabilitiesSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListCapabilities", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListCapabilities", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListCapabilitiesPreparer prepares the ListCapabilities request.
func (client LocationClient) ListCapabilitiesPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCapabilitiesSender sends the ListCapabilities request. The method will close the
// http.Response Body if it receives an error.
func (client LocationClient) ListCapabilitiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCapabilitiesResponder handles the response to the ListCapabilities request. The method always
// closes the http.Response Body.
func (client LocationClient) ListCapabilitiesResponder(resp *http.Response) (result CapabilitiesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listCapabilitiesNextResults retrieves the next set of results, if any.
func (client LocationClient) listCapabilitiesNextResults(ctx context.Context, lastResults CapabilitiesListResult) (result CapabilitiesListResult, err error) {
	req, err := lastResults.capabilitiesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerinstance.LocationClient", "listCapabilitiesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListCapabilitiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerinstance.LocationClient", "listCapabilitiesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "listCapabilitiesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListCapabilitiesComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocationClient) ListCapabilitiesComplete(ctx context.Context, location string) (result CapabilitiesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.ListCapabilities")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListCapabilities(ctx, location)
	return
}

// ListUsage get the usage for a subscription
// Parameters:
// location - the identifier for the physical azure location.
func (client LocationClient) ListUsage(ctx context.Context, location string) (result UsageListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.ListUsage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListUsagePreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListUsage", resp, "Failure sending request")
		return
	}

	result, err = client.ListUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.LocationClient", "ListUsage", resp, "Failure responding to request")
		return
	}

	return
}

// ListUsagePreparer prepares the ListUsage request.
func (client LocationClient) ListUsagePreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUsageSender sends the ListUsage request. The method will close the
// http.Response Body if it receives an error.
func (client LocationClient) ListUsageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListUsageResponder handles the response to the ListUsage request. The method always
// closes the http.Response Body.
func (client LocationClient) ListUsageResponder(resp *http.Response) (result UsageListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
