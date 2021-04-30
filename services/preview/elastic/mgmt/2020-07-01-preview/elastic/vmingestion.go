package elastic

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

// VMIngestionClient is the client for the VMIngestion methods of the Elastic service.
type VMIngestionClient struct {
	BaseClient
}

// NewVMIngestionClient creates an instance of the VMIngestionClient client.
func NewVMIngestionClient(subscriptionID string) VMIngestionClient {
	return NewVMIngestionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVMIngestionClientWithBaseURI creates an instance of the VMIngestionClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVMIngestionClientWithBaseURI(baseURI string, subscriptionID string) VMIngestionClient {
	return VMIngestionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Details sends the details request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Elastic resource belongs.
// monitorName - monitor resource name
func (client VMIngestionClient) Details(ctx context.Context, resourceGroupName string, monitorName string) (result VMIngestionDetailsResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VMIngestionClient.Details")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetailsPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMIngestionClient", "Details", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "elastic.VMIngestionClient", "Details", resp, "Failure sending request")
		return
	}

	result, err = client.DetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elastic.VMIngestionClient", "Details", resp, "Failure responding to request")
		return
	}

	return
}

// DetailsPreparer prepares the Details request.
func (client VMIngestionClient) DetailsPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/vmIngestionDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetailsSender sends the Details request. The method will close the
// http.Response Body if it receives an error.
func (client VMIngestionClient) DetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DetailsResponder handles the response to the Details request. The method always
// closes the http.Response Body.
func (client VMIngestionClient) DetailsResponder(resp *http.Response) (result VMIngestionDetailsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
