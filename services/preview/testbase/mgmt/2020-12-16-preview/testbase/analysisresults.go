package testbase

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

// AnalysisResultsClient is the REST API for Test Base
type AnalysisResultsClient struct {
	BaseClient
}

// NewAnalysisResultsClient creates an instance of the AnalysisResultsClient client.
func NewAnalysisResultsClient(subscriptionID string) AnalysisResultsClient {
	return NewAnalysisResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAnalysisResultsClientWithBaseURI creates an instance of the AnalysisResultsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAnalysisResultsClientWithBaseURI(baseURI string, subscriptionID string) AnalysisResultsClient {
	return AnalysisResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the Analysis Results of a Test Result. The result collection will only contain one element as all the
// data will be nested in a singleton object.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource.
// testBaseAccountName - the resource name of the Test Base Account.
// packageName - the resource name of the Test Base Package.
// testResultName - the Test Result Name. It equals to {osName}-{TestResultId} string.
// analysisResultType - the type of the Analysis Result of a Test Result.
func (client AnalysisResultsClient) List(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string, analysisResultType AnalysisResultType) (result AnalysisResultListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AnalysisResultsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, testBaseAccountName, packageName, testResultName, analysisResultType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.AnalysisResultsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "testbase.AnalysisResultsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testbase.AnalysisResultsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AnalysisResultsClient) ListPreparer(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, testResultName string, analysisResultType AnalysisResultType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"packageName":         autorest.Encode("path", packageName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
		"testBaseAccountName": autorest.Encode("path", testBaseAccountName),
		"testResultName":      autorest.Encode("path", testResultName),
	}

	const APIVersion = "2020-12-16-preview"
	queryParameters := map[string]interface{}{
		"analysisResultType": autorest.Encode("query", analysisResultType),
		"api-version":        APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/testResults/{testResultName}/analysisResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AnalysisResultsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AnalysisResultsClient) ListResponder(resp *http.Response) (result AnalysisResultListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
