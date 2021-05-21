package aad

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

// OuContainerOperationsClient is the the AAD Domain Services API.
type OuContainerOperationsClient struct {
	BaseClient
}

// NewOuContainerOperationsClient creates an instance of the OuContainerOperationsClient client.
func NewOuContainerOperationsClient(subscriptionID string) OuContainerOperationsClient {
	return NewOuContainerOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOuContainerOperationsClientWithBaseURI creates an instance of the OuContainerOperationsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewOuContainerOperationsClientWithBaseURI(baseURI string, subscriptionID string) OuContainerOperationsClient {
	return OuContainerOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all the available OuContainer operations.
func (client OuContainerOperationsClient) List(ctx context.Context) (result OperationEntityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerOperationsClient.List")
		defer func() {
			sc := -1
			if result.oelr.Response.Response != nil {
				sc = result.oelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerOperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.oelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.OuContainerOperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.oelr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerOperationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.oelr.hasNextLink() && result.oelr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client OuContainerOperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Aad/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OuContainerOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OuContainerOperationsClient) ListResponder(resp *http.Response) (result OperationEntityListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client OuContainerOperationsClient) listNextResults(ctx context.Context, lastResults OperationEntityListResult) (result OperationEntityListResult, err error) {
	req, err := lastResults.operationEntityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "aad.OuContainerOperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "aad.OuContainerOperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerOperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OuContainerOperationsClient) ListComplete(ctx context.Context) (result OperationEntityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerOperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
