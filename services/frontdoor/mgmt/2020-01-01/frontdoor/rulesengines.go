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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RulesEnginesClient is the frontDoor Client
type RulesEnginesClient struct {
	BaseClient
}

// NewRulesEnginesClient creates an instance of the RulesEnginesClient client.
func NewRulesEnginesClient(subscriptionID string) RulesEnginesClient {
	return NewRulesEnginesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRulesEnginesClientWithBaseURI creates an instance of the RulesEnginesClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRulesEnginesClientWithBaseURI(baseURI string, subscriptionID string) RulesEnginesClient {
	return RulesEnginesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new Rules Engine Configuration with the specified name within the specified Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// rulesEngineName - name of the Rules Engine which is unique within the Front Door.
// rulesEngineParameters - rules Engine Configuration properties needed to create a new Rules Engine
// Configuration.
func (client RulesEnginesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine) (result RulesEnginesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesEnginesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: rulesEngineName,
			Constraints: []validation.Constraint{{Target: "rulesEngineName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "rulesEngineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rulesEngineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.RulesEnginesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, frontDoorName, rulesEngineName, rulesEngineParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RulesEnginesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters RulesEngine) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rulesEngineName":   autorest.Encode("path", rulesEngineName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	rulesEngineParameters.Name = nil
	rulesEngineParameters.Type = nil
	rulesEngineParameters.ID = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}", pathParameters),
		autorest.WithJSON(rulesEngineParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RulesEnginesClient) CreateOrUpdateSender(req *http.Request) (future RulesEnginesCreateOrUpdateFuture, err error) {
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
func (client RulesEnginesClient) CreateOrUpdateResponder(resp *http.Response) (result RulesEngine, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing Rules Engine Configuration with the specified parameters.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// rulesEngineName - name of the Rules Engine which is unique within the Front Door.
func (client RulesEnginesClient) Delete(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string) (result RulesEnginesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesEnginesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: rulesEngineName,
			Constraints: []validation.Constraint{{Target: "rulesEngineName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "rulesEngineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rulesEngineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.RulesEnginesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, frontDoorName, rulesEngineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RulesEnginesClient) DeletePreparer(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rulesEngineName":   autorest.Encode("path", rulesEngineName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RulesEnginesClient) DeleteSender(req *http.Request) (future RulesEnginesDeleteFuture, err error) {
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
func (client RulesEnginesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Rules Engine Configuration with the specified name within the specified Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// rulesEngineName - name of the Rules Engine which is unique within the Front Door.
func (client RulesEnginesClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string) (result RulesEngine, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesEnginesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: rulesEngineName,
			Constraints: []validation.Constraint{{Target: "rulesEngineName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "rulesEngineName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "rulesEngineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.RulesEnginesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, frontDoorName, rulesEngineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RulesEnginesClient) GetPreparer(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"rulesEngineName":   autorest.Encode("path", rulesEngineName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RulesEnginesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RulesEnginesClient) GetResponder(resp *http.Response) (result RulesEngine, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByFrontDoor lists all of the Rules Engine Configurations within a Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
func (client RulesEnginesClient) ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result RulesEngineListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesEnginesClient.ListByFrontDoor")
		defer func() {
			sc := -1
			if result.relr.Response.Response != nil {
				sc = result.relr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.RulesEnginesClient", "ListByFrontDoor", err.Error())
	}

	result.fn = client.listByFrontDoorNextResults
	req, err := client.ListByFrontDoorPreparer(ctx, resourceGroupName, frontDoorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "ListByFrontDoor", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFrontDoorSender(req)
	if err != nil {
		result.relr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "ListByFrontDoor", resp, "Failure sending request")
		return
	}

	result.relr, err = client.ListByFrontDoorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "ListByFrontDoor", resp, "Failure responding to request")
	}

	return
}

// ListByFrontDoorPreparer prepares the ListByFrontDoor request.
func (client RulesEnginesClient) ListByFrontDoorPreparer(ctx context.Context, resourceGroupName string, frontDoorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByFrontDoorSender sends the ListByFrontDoor request. The method will close the
// http.Response Body if it receives an error.
func (client RulesEnginesClient) ListByFrontDoorSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByFrontDoorResponder handles the response to the ListByFrontDoor request. The method always
// closes the http.Response Body.
func (client RulesEnginesClient) ListByFrontDoorResponder(resp *http.Response) (result RulesEngineListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFrontDoorNextResults retrieves the next set of results, if any.
func (client RulesEnginesClient) listByFrontDoorNextResults(ctx context.Context, lastResults RulesEngineListResult) (result RulesEngineListResult, err error) {
	req, err := lastResults.rulesEngineListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "listByFrontDoorNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFrontDoorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "listByFrontDoorNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFrontDoorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.RulesEnginesClient", "listByFrontDoorNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByFrontDoorComplete enumerates all values, automatically crossing page boundaries as required.
func (client RulesEnginesClient) ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result RulesEngineListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RulesEnginesClient.ListByFrontDoor")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByFrontDoor(ctx, resourceGroupName, frontDoorName)
	return
}
