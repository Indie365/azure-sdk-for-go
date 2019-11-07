package security

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NetworkDataClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type NetworkDataClient struct {
	BaseClient
}

// NewNetworkDataClient creates an instance of the NetworkDataClient client.
func NewNetworkDataClient(subscriptionID string, ascLocation string) NetworkDataClient {
	return NewNetworkDataClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewNetworkDataClientWithBaseURI creates an instance of the NetworkDataClient client.
func NewNetworkDataClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) NetworkDataClient {
	return NetworkDataClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get get the network data on your scanned resource
// Parameters:
// resourceID - the identifier of the resource.
// expand - expand whether you want to get more information about the network data (ports and connections
// details)
func (client NetworkDataClient) Get(ctx context.Context, resourceID string, expand ExpandValues) (result NetworkData, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkDataClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NetworkDataClient) GetPreparer(ctx context.Context, resourceID string, expand ExpandValues) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": autorest.Encode("path", resourceID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/NetworkData/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkDataClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NetworkDataClient) GetResponder(resp *http.Response) (result NetworkData, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetResourceCollection get the network data collection on your scanned resource
// Parameters:
// resourceID - the identifier of the resource.
// expand - expand whether you want to get more information about the network data (ports and connections
// details)
func (client NetworkDataClient) GetResourceCollection(ctx context.Context, resourceID string, expand ExpandValues) (result NetworkDataList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkDataClient.GetResourceCollection")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetResourceCollectionPreparer(ctx, resourceID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "GetResourceCollection", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceCollectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "GetResourceCollection", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceCollectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "GetResourceCollection", resp, "Failure responding to request")
	}

	return
}

// GetResourceCollectionPreparer prepares the GetResourceCollection request.
func (client NetworkDataClient) GetResourceCollectionPreparer(ctx context.Context, resourceID string, expand ExpandValues) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": autorest.Encode("path", resourceID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/NetworkData", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceCollectionSender sends the GetResourceCollection request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkDataClient) GetResourceCollectionSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResourceCollectionResponder handles the response to the GetResourceCollection request. The method always
// closes the http.Response Body.
func (client NetworkDataClient) GetResourceCollectionResponder(resp *http.Response) (result NetworkDataList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get the network data on all your scanned resources inside a scope
// Parameters:
// expand - expand whether you want to get more information about the network data (ports and connections
// details)
func (client NetworkDataClient) List(ctx context.Context, expand ExpandValues) (result NetworkDataListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkDataClient.List")
		defer func() {
			sc := -1
			if result.ndl.Response.Response != nil {
				sc = result.ndl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.NetworkDataClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ndl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "List", resp, "Failure sending request")
		return
	}

	result.ndl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client NetworkDataClient) ListPreparer(ctx context.Context, expand ExpandValues) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/networkData", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkDataClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NetworkDataClient) ListResponder(resp *http.Response) (result NetworkDataList, err error) {
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
func (client NetworkDataClient) listNextResults(ctx context.Context, lastResults NetworkDataList) (result NetworkDataList, err error) {
	req, err := lastResults.networkDataListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.NetworkDataClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.NetworkDataClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NetworkDataClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client NetworkDataClient) ListComplete(ctx context.Context, expand ExpandValues) (result NetworkDataListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkDataClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, expand)
	return
}
