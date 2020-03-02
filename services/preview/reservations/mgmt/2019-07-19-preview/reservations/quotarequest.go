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

// QuotaRequestClient is the client for the QuotaRequest methods of the Reservations service.
type QuotaRequestClient struct {
	BaseClient
}

// NewQuotaRequestClient creates an instance of the QuotaRequestClient client.
func NewQuotaRequestClient() QuotaRequestClient {
	return NewQuotaRequestClientWithBaseURI(DefaultBaseURI)
}

// NewQuotaRequestClientWithBaseURI creates an instance of the QuotaRequestClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQuotaRequestClientWithBaseURI(baseURI string) QuotaRequestClient {
	return QuotaRequestClient{NewWithBaseURI(baseURI)}
}

// Create submits Quota change request for a resource provider for the specified location for the specific resource in
// the parameter. To use, first make a Get request to get quota information for the specific resource. This information
// consists of information regarding that specific resources. For the specific resource, if it requires an update to
// the quota, update the limit field in the response from the Get request to the new value of quota. Then, submit this
// updated JSON object to this quota request API. This will update the quota to the value specified. The location
// header in the response will be used to track the status of the quota request. Please check the provisioningState
// field in the response.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource Provider id.
// location - azure region.
// resourceName - the Resource name for the specific resource provider, such as SKU name for Microsoft.Compute,
// pool for Microsoft.Batch.
// createQuotaRequest - quota requests payload.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client QuotaRequestClient) Create(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase, ifMatch string) (result QuotaRequestCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaRequestClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, subscriptionID, providerID, location, resourceName, createQuotaRequest, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client QuotaRequestClient) CreatePreparer(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase, ifMatch string) (*http.Request, error) {
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
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaRequestClient) CreateSender(req *http.Request) (future QuotaRequestCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client QuotaRequestClient) CreateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update submits Quota change request for a resource provider for the specified location for the specific resource in
// the parameter. To use, first make a Get request to get quota information for the specific resource. This information
// consists of information regarding that specific resources. For the specific resource, if it requires an update to
// the quota, update the limit field in the response from the Get request to the new value of quota. Then, submit this
// updated JSON object to this quota request API. This will update the quota to the value specified. The location
// header in the response will be used to track the status of the quota request. Please check the provisioningState
// field in the response.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource Provider id.
// location - azure region.
// resourceName - the Resource name for the specific resource provider, such as SKU name for Microsoft.Compute,
// pool for Microsoft.Batch.
// createQuotaRequest - quota requests payload.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client QuotaRequestClient) Update(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase, ifMatch string) (result QuotaRequestUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaRequestClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, subscriptionID, providerID, location, resourceName, createQuotaRequest, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client QuotaRequestClient) UpdatePreparer(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest CurrentQuotaLimitBase, ifMatch string) (*http.Request, error) {
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
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaRequestClient) UpdateSender(req *http.Request) (future QuotaRequestUpdateFuture, err error) {
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
func (client QuotaRequestClient) UpdateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
