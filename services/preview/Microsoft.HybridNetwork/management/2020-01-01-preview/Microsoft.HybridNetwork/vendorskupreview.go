package microsofthybridnetwork

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

// VendorSkuPreviewClient is the client for the VendorSkuPreview methods of the Microsofthybridnetwork service.
type VendorSkuPreviewClient struct {
	BaseClient
}

// NewVendorSkuPreviewClient creates an instance of the VendorSkuPreviewClient client.
func NewVendorSkuPreviewClient(subscriptionID string) VendorSkuPreviewClient {
	return NewVendorSkuPreviewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVendorSkuPreviewClientWithBaseURI creates an instance of the VendorSkuPreviewClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewVendorSkuPreviewClientWithBaseURI(baseURI string, subscriptionID string) VendorSkuPreviewClient {
	return VendorSkuPreviewClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates preview information of a vendor sku
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the vendor sku.
// previewSubscription - preview subscription id.
// parameters - parameters supplied to the update vendor virtual network function.
func (client VendorSkuPreviewClient) CreateOrUpdate(ctx context.Context, vendorName string, skuName string, previewSubscription string, parameters PreviewSubscription) (result VendorSkuPreviewCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkuPreviewClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, vendorName, skuName, previewSubscription, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VendorSkuPreviewClient) CreateOrUpdatePreparer(ctx context.Context, vendorName string, skuName string, previewSubscription string, parameters PreviewSubscription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"previewSubscription": autorest.Encode("path", previewSubscription),
		"skuName":             autorest.Encode("path", skuName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"vendorName":          autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Name = nil
	parameters.ID = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions/{previewSubscription}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkuPreviewClient) CreateOrUpdateSender(req *http.Request) (future VendorSkuPreviewCreateOrUpdateFuture, err error) {
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
func (client VendorSkuPreviewClient) CreateOrUpdateResponder(resp *http.Response) (result PreviewSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes preview information of a vendor sku
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the vendor sku.
// previewSubscription - preview subscription id.
func (client VendorSkuPreviewClient) Delete(ctx context.Context, vendorName string, skuName string, previewSubscription string) (result VendorSkuPreviewDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkuPreviewClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vendorName, skuName, previewSubscription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VendorSkuPreviewClient) DeletePreparer(ctx context.Context, vendorName string, skuName string, previewSubscription string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"previewSubscription": autorest.Encode("path", previewSubscription),
		"skuName":             autorest.Encode("path", skuName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"vendorName":          autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions/{previewSubscription}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkuPreviewClient) DeleteSender(req *http.Request) (future VendorSkuPreviewDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VendorSkuPreviewClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets preview information of a vendor sku
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the vendor sku.
// previewSubscription - preview subscription id.
func (client VendorSkuPreviewClient) Get(ctx context.Context, vendorName string, skuName string, previewSubscription string) (result PreviewSubscription, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkuPreviewClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vendorName, skuName, previewSubscription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VendorSkuPreviewClient) GetPreparer(ctx context.Context, vendorName string, skuName string, previewSubscription string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"previewSubscription": autorest.Encode("path", previewSubscription),
		"skuName":             autorest.Encode("path", skuName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"vendorName":          autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions/{previewSubscription}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkuPreviewClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VendorSkuPreviewClient) GetResponder(resp *http.Response) (result PreviewSubscription, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets preview information of a vendor sku
// Parameters:
// vendorName - the name of the vendor.
// skuName - the name of the sku.
func (client VendorSkuPreviewClient) List(ctx context.Context, vendorName string, skuName string) (result PreviewSubscriptionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkuPreviewClient.List")
		defer func() {
			sc := -1
			if result.psl.Response.Response != nil {
				sc = result.psl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vendorName, skuName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.psl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "List", resp, "Failure sending request")
		return
	}

	result.psl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "List", resp, "Failure responding to request")
	}
	if result.psl.hasNextLink() && result.psl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client VendorSkuPreviewClient) ListPreparer(ctx context.Context, vendorName string, skuName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"skuName":        autorest.Encode("path", skuName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VendorSkuPreviewClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VendorSkuPreviewClient) ListResponder(resp *http.Response) (result PreviewSubscriptionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VendorSkuPreviewClient) listNextResults(ctx context.Context, lastResults PreviewSubscriptionsList) (result PreviewSubscriptionsList, err error) {
	req, err := lastResults.previewSubscriptionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "microsofthybridnetwork.VendorSkuPreviewClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VendorSkuPreviewClient) ListComplete(ctx context.Context, vendorName string, skuName string) (result PreviewSubscriptionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorSkuPreviewClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vendorName, skuName)
	return
}
