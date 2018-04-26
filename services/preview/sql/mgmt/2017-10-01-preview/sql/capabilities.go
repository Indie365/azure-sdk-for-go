package sql

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
	"net/http"
)

// CapabilitiesClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type CapabilitiesClient struct {
	BaseClient
}

// NewCapabilitiesClient creates an instance of the CapabilitiesClient client.
func NewCapabilitiesClient(subscriptionID string) CapabilitiesClient {
	return NewCapabilitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCapabilitiesClientWithBaseURI creates an instance of the CapabilitiesClient client.
func NewCapabilitiesClientWithBaseURI(baseURI string, subscriptionID string) CapabilitiesClient {
	return CapabilitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByLocation gets the subscription capabilities available for the specified location.
//
// locationName is the location name whose capabilities are retrieved. include is if specified, restricts the
// response to only include the selected item.
func (client CapabilitiesClient) ListByLocation(ctx context.Context, locationName string, include CapabilityGroup) (result LocationCapabilities, err error) {
	req, err := client.ListByLocationPreparer(ctx, locationName, include)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.CapabilitiesClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.CapabilitiesClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.CapabilitiesClient", "ListByLocation", resp, "Failure responding to request")
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client CapabilitiesClient) ListByLocationPreparer(ctx context.Context, locationName string, include CapabilityGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(include)) > 0 {
		queryParameters["include"] = autorest.Encode("query", include)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client CapabilitiesClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client CapabilitiesClient) ListByLocationResponder(resp *http.Response) (result LocationCapabilities, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
