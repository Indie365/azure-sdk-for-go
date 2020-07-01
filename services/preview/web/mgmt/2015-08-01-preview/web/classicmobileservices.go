package web

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ClassicMobileServicesClient is the webSite Management Client
type ClassicMobileServicesClient struct {
	BaseClient
}

// NewClassicMobileServicesClient creates an instance of the ClassicMobileServicesClient client.
func NewClassicMobileServicesClient(subscriptionID string) ClassicMobileServicesClient {
	return NewClassicMobileServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClassicMobileServicesClientWithBaseURI creates an instance of the ClassicMobileServicesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewClassicMobileServicesClientWithBaseURI(baseURI string, subscriptionID string) ClassicMobileServicesClient {
	return ClassicMobileServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeleteClassicMobileService sends the delete classic mobile service request.
// Parameters:
// resourceGroupName - name of resource group
// name - name of mobile service
func (client ClassicMobileServicesClient) DeleteClassicMobileService(ctx context.Context, resourceGroupName string, name string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClassicMobileServicesClient.DeleteClassicMobileService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteClassicMobileServicePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "DeleteClassicMobileService", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteClassicMobileServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "DeleteClassicMobileService", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteClassicMobileServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "DeleteClassicMobileService", resp, "Failure responding to request")
	}

	return
}

// DeleteClassicMobileServicePreparer prepares the DeleteClassicMobileService request.
func (client ClassicMobileServicesClient) DeleteClassicMobileServicePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/classicMobileServices/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteClassicMobileServiceSender sends the DeleteClassicMobileService request. The method will close the
// http.Response Body if it receives an error.
func (client ClassicMobileServicesClient) DeleteClassicMobileServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteClassicMobileServiceResponder handles the response to the DeleteClassicMobileService request. The method always
// closes the http.Response Body.
func (client ClassicMobileServicesClient) DeleteClassicMobileServiceResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetClassicMobileService sends the get classic mobile service request.
// Parameters:
// resourceGroupName - name of resource group
// name - name of mobile service
func (client ClassicMobileServicesClient) GetClassicMobileService(ctx context.Context, resourceGroupName string, name string) (result ClassicMobileService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClassicMobileServicesClient.GetClassicMobileService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetClassicMobileServicePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileService", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetClassicMobileServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileService", resp, "Failure sending request")
		return
	}

	result, err = client.GetClassicMobileServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileService", resp, "Failure responding to request")
	}

	return
}

// GetClassicMobileServicePreparer prepares the GetClassicMobileService request.
func (client ClassicMobileServicesClient) GetClassicMobileServicePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/classicMobileServices/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetClassicMobileServiceSender sends the GetClassicMobileService request. The method will close the
// http.Response Body if it receives an error.
func (client ClassicMobileServicesClient) GetClassicMobileServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetClassicMobileServiceResponder handles the response to the GetClassicMobileService request. The method always
// closes the http.Response Body.
func (client ClassicMobileServicesClient) GetClassicMobileServiceResponder(resp *http.Response) (result ClassicMobileService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetClassicMobileServices sends the get classic mobile services request.
// Parameters:
// resourceGroupName - name of resource group
func (client ClassicMobileServicesClient) GetClassicMobileServices(ctx context.Context, resourceGroupName string) (result ClassicMobileServiceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClassicMobileServicesClient.GetClassicMobileServices")
		defer func() {
			sc := -1
			if result.cmsc.Response.Response != nil {
				sc = result.cmsc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getClassicMobileServicesNextResults
	req, err := client.GetClassicMobileServicesPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileServices", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetClassicMobileServicesSender(req)
	if err != nil {
		result.cmsc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileServices", resp, "Failure sending request")
		return
	}

	result.cmsc, err = client.GetClassicMobileServicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "GetClassicMobileServices", resp, "Failure responding to request")
	}

	return
}

// GetClassicMobileServicesPreparer prepares the GetClassicMobileServices request.
func (client ClassicMobileServicesClient) GetClassicMobileServicesPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/classicMobileServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetClassicMobileServicesSender sends the GetClassicMobileServices request. The method will close the
// http.Response Body if it receives an error.
func (client ClassicMobileServicesClient) GetClassicMobileServicesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetClassicMobileServicesResponder handles the response to the GetClassicMobileServices request. The method always
// closes the http.Response Body.
func (client ClassicMobileServicesClient) GetClassicMobileServicesResponder(resp *http.Response) (result ClassicMobileServiceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getClassicMobileServicesNextResults retrieves the next set of results, if any.
func (client ClassicMobileServicesClient) getClassicMobileServicesNextResults(ctx context.Context, lastResults ClassicMobileServiceCollection) (result ClassicMobileServiceCollection, err error) {
	req, err := lastResults.classicMobileServiceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "getClassicMobileServicesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetClassicMobileServicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "getClassicMobileServicesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetClassicMobileServicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ClassicMobileServicesClient", "getClassicMobileServicesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetClassicMobileServicesComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClassicMobileServicesClient) GetClassicMobileServicesComplete(ctx context.Context, resourceGroupName string) (result ClassicMobileServiceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClassicMobileServicesClient.GetClassicMobileServices")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetClassicMobileServices(ctx, resourceGroupName)
	return
}
