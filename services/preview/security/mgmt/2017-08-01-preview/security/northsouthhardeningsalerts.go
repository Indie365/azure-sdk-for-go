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
	"net/http"
)

// NorthSouthHardeningsAlertsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type NorthSouthHardeningsAlertsClient struct {
	BaseClient
}

// NewNorthSouthHardeningsAlertsClient creates an instance of the NorthSouthHardeningsAlertsClient client.
func NewNorthSouthHardeningsAlertsClient(subscriptionID string, ascLocation string) NorthSouthHardeningsAlertsClient {
	return NewNorthSouthHardeningsAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewNorthSouthHardeningsAlertsClientWithBaseURI creates an instance of the NorthSouthHardeningsAlertsClient client.
func NewNorthSouthHardeningsAlertsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) NorthSouthHardeningsAlertsClient {
	return NorthSouthHardeningsAlertsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get get a north-south hardening alert
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// northSouthResourceName - name of a north-south resource.
// northSouthAlertName - name of a north-south alert.
func (client NorthSouthHardeningsAlertsClient) Get(ctx context.Context, resourceGroupName string, northSouthResourceName string, northSouthAlertName string) (result NorthSouthHardeningsAlert, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.NorthSouthHardeningsAlertsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, northSouthResourceName, northSouthAlertName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NorthSouthHardeningsAlertsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.NorthSouthHardeningsAlertsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NorthSouthHardeningsAlertsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NorthSouthHardeningsAlertsClient) GetPreparer(ctx context.Context, resourceGroupName string, northSouthResourceName string, northSouthAlertName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"northSouthAlertName":    autorest.Encode("path", northSouthAlertName),
		"northSouthResourceName": autorest.Encode("path", northSouthResourceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/northSouthHardenings/{northSouthResourceName}/alerts/{northSouthAlertName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NorthSouthHardeningsAlertsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NorthSouthHardeningsAlertsClient) GetResponder(resp *http.Response) (result NorthSouthHardeningsAlert, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list north-south hardening alerts for a north-south hardening resource
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// northSouthResourceName - name of a north-south resource.
func (client NorthSouthHardeningsAlertsClient) List(ctx context.Context, resourceGroupName string, northSouthResourceName string) (result NorthSouthHardeningsAlertList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.NorthSouthHardeningsAlertsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, northSouthResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NorthSouthHardeningsAlertsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.NorthSouthHardeningsAlertsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.NorthSouthHardeningsAlertsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client NorthSouthHardeningsAlertsClient) ListPreparer(ctx context.Context, resourceGroupName string, northSouthResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"northSouthResourceName": autorest.Encode("path", northSouthResourceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/northSouthHardenings/{northSouthResourceName}/alerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NorthSouthHardeningsAlertsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NorthSouthHardeningsAlertsClient) ListResponder(resp *http.Response) (result NorthSouthHardeningsAlertList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
