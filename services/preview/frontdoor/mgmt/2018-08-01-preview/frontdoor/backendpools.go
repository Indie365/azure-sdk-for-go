package frontdoor

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

// BackendPoolsClient is the frontdoor Client
type BackendPoolsClient struct {
	BaseClient
}

// NewBackendPoolsClient creates an instance of the BackendPoolsClient client.
func NewBackendPoolsClient(subscriptionID string) BackendPoolsClient {
	return NewBackendPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackendPoolsClientWithBaseURI creates an instance of the BackendPoolsClient client.
func NewBackendPoolsClientWithBaseURI(baseURI string, subscriptionID string) BackendPoolsClient {
	return BackendPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new Backend Pool with the specified Pool name within the specified Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// backendPoolName - name of the Backend Pool which is unique within the Front Door.
// backendPoolParameters - backend Pool properties needed to create a new Pool.
func (client BackendPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, backendPoolName string, backendPoolParameters BackendPool) (result BackendPoolsCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: backendPoolName,
			Constraints: []validation.Constraint{{Target: "backendPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "backendPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.BackendPoolsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, frontDoorName, backendPoolName, backendPoolParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackendPoolsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, frontDoorName string, backendPoolName string, backendPoolParameters BackendPool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendPoolName":   autorest.Encode("path", backendPoolName),
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools/{backendPoolName}", pathParameters),
		autorest.WithJSON(backendPoolParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackendPoolsClient) CreateOrUpdateSender(req *http.Request) (future BackendPoolsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackendPoolsClient) CreateOrUpdateResponder(resp *http.Response) (result BackendPool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing Backend Pool with the specified parameters.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// backendPoolName - name of the Backend Pool which is unique within the Front Door.
func (client BackendPoolsClient) Delete(ctx context.Context, resourceGroupName string, frontDoorName string, backendPoolName string) (result BackendPoolsDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: backendPoolName,
			Constraints: []validation.Constraint{{Target: "backendPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "backendPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.BackendPoolsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, frontDoorName, backendPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackendPoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, frontDoorName string, backendPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendPoolName":   autorest.Encode("path", backendPoolName),
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools/{backendPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackendPoolsClient) DeleteSender(req *http.Request) (future BackendPoolsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackendPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Backend Pool with the specified Pool name within the specified Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// backendPoolName - name of the Backend Pool which is unique within the Front Door.
func (client BackendPoolsClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, backendPoolName string) (result BackendPool, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: backendPoolName,
			Constraints: []validation.Constraint{{Target: "backendPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "backendPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.BackendPoolsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, frontDoorName, backendPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackendPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, frontDoorName string, backendPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendPoolName":   autorest.Encode("path", backendPoolName),
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools/{backendPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackendPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackendPoolsClient) GetResponder(resp *http.Response) (result BackendPool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByFrontDoor lists all of the Backend Pools within a Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
func (client BackendPoolsClient) ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result BackendPoolListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.BackendPoolsClient", "ListByFrontDoor", err.Error())
	}

	result.fn = client.listByFrontDoorNextResults
	req, err := client.ListByFrontDoorPreparer(ctx, resourceGroupName, frontDoorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "ListByFrontDoor", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFrontDoorSender(req)
	if err != nil {
		result.bplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "ListByFrontDoor", resp, "Failure sending request")
		return
	}

	result.bplr, err = client.ListByFrontDoorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "ListByFrontDoor", resp, "Failure responding to request")
	}

	return
}

// ListByFrontDoorPreparer prepares the ListByFrontDoor request.
func (client BackendPoolsClient) ListByFrontDoorPreparer(ctx context.Context, resourceGroupName string, frontDoorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/backendPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByFrontDoorSender sends the ListByFrontDoor request. The method will close the
// http.Response Body if it receives an error.
func (client BackendPoolsClient) ListByFrontDoorSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByFrontDoorResponder handles the response to the ListByFrontDoor request. The method always
// closes the http.Response Body.
func (client BackendPoolsClient) ListByFrontDoorResponder(resp *http.Response) (result BackendPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFrontDoorNextResults retrieves the next set of results, if any.
func (client BackendPoolsClient) listByFrontDoorNextResults(lastResults BackendPoolListResult) (result BackendPoolListResult, err error) {
	req, err := lastResults.backendPoolListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "listByFrontDoorNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFrontDoorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "listByFrontDoorNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFrontDoorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BackendPoolsClient", "listByFrontDoorNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByFrontDoorComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackendPoolsClient) ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result BackendPoolListResultIterator, err error) {
	result.page, err = client.ListByFrontDoor(ctx, resourceGroupName, frontDoorName)
	return
}
