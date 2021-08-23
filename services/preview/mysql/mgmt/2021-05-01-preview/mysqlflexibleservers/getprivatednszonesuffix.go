package mysqlflexibleservers

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

// GetPrivateDNSZoneSuffixClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure MySQL resources including servers, databases, firewall rules, VNET rules, log files and
// configurations with new business model.
type GetPrivateDNSZoneSuffixClient struct {
	BaseClient
}

// NewGetPrivateDNSZoneSuffixClient creates an instance of the GetPrivateDNSZoneSuffixClient client.
func NewGetPrivateDNSZoneSuffixClient(subscriptionID string) GetPrivateDNSZoneSuffixClient {
	return NewGetPrivateDNSZoneSuffixClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGetPrivateDNSZoneSuffixClientWithBaseURI creates an instance of the GetPrivateDNSZoneSuffixClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewGetPrivateDNSZoneSuffixClientWithBaseURI(baseURI string, subscriptionID string) GetPrivateDNSZoneSuffixClient {
	return GetPrivateDNSZoneSuffixClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Execute get private DNS zone suffix in the cloud.
func (client GetPrivateDNSZoneSuffixClient) Execute(ctx context.Context) (result GetPrivateDNSZoneSuffixResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GetPrivateDNSZoneSuffixClient.Execute")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExecutePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysqlflexibleservers.GetPrivateDNSZoneSuffixClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysqlflexibleservers.GetPrivateDNSZoneSuffixClient", "Execute", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysqlflexibleservers.GetPrivateDNSZoneSuffixClient", "Execute", resp, "Failure responding to request")
		return
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client GetPrivateDNSZoneSuffixClient) ExecutePreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2021-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.DBforMySQL/getPrivateDnsZoneSuffix"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client GetPrivateDNSZoneSuffixClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client GetPrivateDNSZoneSuffixClient) ExecuteResponder(resp *http.Response) (result GetPrivateDNSZoneSuffixResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
