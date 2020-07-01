package reservations

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

// QuotaClient is the client for the Quota methods of the Reservations service.
type QuotaClient struct {
	BaseClient
}

// NewQuotaClient creates an instance of the QuotaClient client.
func NewQuotaClient() QuotaClient {
	return NewQuotaClientWithBaseURI(DefaultBaseURI)
}

// NewQuotaClientWithBaseURI creates an instance of the QuotaClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQuotaClientWithBaseURI(baseURI string) QuotaClient {
	return QuotaClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate create or update the service limits (quota) of a resource to requested value.
// Steps:
//
// 1. Make the Get request to get the quota information for specific resource.
//
// 2. To increase the quota, update the limit field in the response from Get request to new value.
//
// 3. Submit the JSON to the quota request API to update the quota.
// The Create quota request may be constructed as follows. The PUT operation can be used to update the quota.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource provider id.
// location - azure region.
// resourceName - the resource name for a resource provider, such as SKU name for Microsoft.Compute, Sku or
// TotalLowPriorityCores for Microsoft.MachineLearningServices
// createQuotaRequest - quota requests payload.
func (client QuotaClient) CreateOrUpdate(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase) (result QuotaCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, subscriptionID, providerID, location, resourceName, createQuotaRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client QuotaClient) CreateOrUpdatePreparer(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"resourceName":   autorest.Encode("path", resourceName),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", pathParameters),
		autorest.WithJSON(createQuotaRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) CreateOrUpdateSender(req *http.Request) (future QuotaCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client QuotaClient) CreateOrUpdateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the current service limits (quotas) and usage of a resource. The response from Get API can be leveraged to
// submit quota update requests.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource provider id.
// location - azure region.
// resourceName - the resource name for a resource provider, such as SKU name for Microsoft.Compute, Sku or
// TotalLowPriorityCores for Microsoft.MachineLearningServices
func (client QuotaClient) Get(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string) (result CurrentQuotaLimitBase, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, subscriptionID, providerID, location, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client QuotaClient) GetPreparer(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"resourceName":   autorest.Encode("path", resourceName),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client QuotaClient) GetResponder(resp *http.Response) (result CurrentQuotaLimitBase, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of current service limits (quota) and usages of all the resources. The response from List API can be
// leveraged to submit quota update requests.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource provider id.
// location - azure region.
func (client QuotaClient) List(ctx context.Context, subscriptionID string, providerID string, location string) (result QuotaLimitsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.List")
		defer func() {
			sc := -1
			if result.ql.Response.Response != nil {
				sc = result.ql.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, subscriptionID, providerID, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ql.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "List", resp, "Failure sending request")
		return
	}

	result.ql, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client QuotaClient) ListPreparer(ctx context.Context, subscriptionID string, providerID string, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client QuotaClient) ListResponder(resp *http.Response) (result QuotaLimits, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client QuotaClient) listNextResults(ctx context.Context, lastResults QuotaLimits) (result QuotaLimits, err error) {
	req, err := lastResults.quotaLimitsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "reservations.QuotaClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "reservations.QuotaClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client QuotaClient) ListComplete(ctx context.Context, subscriptionID string, providerID string, location string) (result QuotaLimitsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, subscriptionID, providerID, location)
	return
}

// Update update the service limits (quota) of a resource to requested value.
// Steps:
//
// 1. Make the Get request to get the quota information for specific resource.
//
// 2. To increase the quota, update the limit field in the response from Get request to new value.
//
// 3. Submit the JSON to the quota request API to update the quota.
// The Update quota request may be constructed as follows. The PATCH operation can be used to update the quota.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource provider id.
// location - azure region.
// resourceName - the resource name for a resource provider, such as SKU name for Microsoft.Compute, Sku or
// TotalLowPriorityCores for Microsoft.MachineLearningServices
// createQuotaRequest - quota requests payload.
func (client QuotaClient) Update(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase) (result QuotaUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, subscriptionID, providerID, location, resourceName, createQuotaRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client QuotaClient) UpdatePreparer(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"resourceName":   autorest.Encode("path", resourceName),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", pathParameters),
		autorest.WithJSON(createQuotaRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) UpdateSender(req *http.Request) (future QuotaUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client QuotaClient) UpdateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
