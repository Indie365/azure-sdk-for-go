package botservice

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

// BotConnectionClient is the azure Bot Service is a platform for creating smart conversational agents.
type BotConnectionClient struct {
	BaseClient
}

// NewBotConnectionClient creates an instance of the BotConnectionClient client.
func NewBotConnectionClient(subscriptionID string) BotConnectionClient {
	return NewBotConnectionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBotConnectionClientWithBaseURI creates an instance of the BotConnectionClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBotConnectionClientWithBaseURI(baseURI string, subscriptionID string) BotConnectionClient {
	return BotConnectionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create register a new Auth Connection for a Bot Service
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// connectionName - the name of the Bot Service Connection Setting resource
// parameters - the parameters to provide for creating the Connection Setting.
func (client BotConnectionClient) Create(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting) (result ConnectionSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.Create")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: connectionName,
			Constraints: []validation.Constraint{{Target: "connectionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "connectionName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "connectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotConnectionClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, connectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client BotConnectionClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) CreateResponder(resp *http.Response) (result ConnectionSetting, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Connection Setting registration for a Bot Service
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// connectionName - the name of the Bot Service Connection Setting resource
func (client BotConnectionClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: connectionName,
			Constraints: []validation.Constraint{{Target: "connectionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "connectionName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "connectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotConnectionClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BotConnectionClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a Connection Setting registration for a Bot Service
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// connectionName - the name of the Bot Service Connection Setting resource
func (client BotConnectionClient) Get(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (result ConnectionSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: connectionName,
			Constraints: []validation.Constraint{{Target: "connectionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "connectionName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "connectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotConnectionClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BotConnectionClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) GetResponder(resp *http.Response) (result ConnectionSetting, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBotService returns all the Connection Settings registered to a particular BotService resource
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
func (client BotConnectionClient) ListByBotService(ctx context.Context, resourceGroupName string, resourceName string) (result ConnectionSettingResponseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.ListByBotService")
		defer func() {
			sc := -1
			if result.csrl.Response.Response != nil {
				sc = result.csrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotConnectionClient", "ListByBotService", err.Error())
	}

	result.fn = client.listByBotServiceNextResults
	req, err := client.ListByBotServicePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListByBotService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBotServiceSender(req)
	if err != nil {
		result.csrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListByBotService", resp, "Failure sending request")
		return
	}

	result.csrl, err = client.ListByBotServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListByBotService", resp, "Failure responding to request")
	}

	return
}

// ListByBotServicePreparer prepares the ListByBotService request.
func (client BotConnectionClient) ListByBotServicePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBotServiceSender sends the ListByBotService request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) ListByBotServiceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByBotServiceResponder handles the response to the ListByBotService request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) ListByBotServiceResponder(resp *http.Response) (result ConnectionSettingResponseList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBotServiceNextResults retrieves the next set of results, if any.
func (client BotConnectionClient) listByBotServiceNextResults(ctx context.Context, lastResults ConnectionSettingResponseList) (result ConnectionSettingResponseList, err error) {
	req, err := lastResults.connectionSettingResponseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "listByBotServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBotServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "listByBotServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBotServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "listByBotServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBotServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client BotConnectionClient) ListByBotServiceComplete(ctx context.Context, resourceGroupName string, resourceName string) (result ConnectionSettingResponseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.ListByBotService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBotService(ctx, resourceGroupName, resourceName)
	return
}

// ListServiceProviders lists the available Service Providers for creating Connection Settings
func (client BotConnectionClient) ListServiceProviders(ctx context.Context) (result ServiceProviderResponseList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.ListServiceProviders")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListServiceProvidersPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListServiceProviders", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListServiceProvidersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListServiceProviders", resp, "Failure sending request")
		return
	}

	result, err = client.ListServiceProvidersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListServiceProviders", resp, "Failure responding to request")
	}

	return
}

// ListServiceProvidersPreparer prepares the ListServiceProviders request.
func (client BotConnectionClient) ListServiceProvidersPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BotService/listAuthServiceProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListServiceProvidersSender sends the ListServiceProviders request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) ListServiceProvidersSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListServiceProvidersResponder handles the response to the ListServiceProviders request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) ListServiceProvidersResponder(resp *http.Response) (result ServiceProviderResponseList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListWithSecrets get a Connection Setting registration for a Bot Service
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// connectionName - the name of the Bot Service Connection Setting resource
func (client BotConnectionClient) ListWithSecrets(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (result ConnectionSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.ListWithSecrets")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: connectionName,
			Constraints: []validation.Constraint{{Target: "connectionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "connectionName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "connectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotConnectionClient", "ListWithSecrets", err.Error())
	}

	req, err := client.ListWithSecretsPreparer(ctx, resourceGroupName, resourceName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListWithSecrets", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListWithSecretsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListWithSecrets", resp, "Failure sending request")
		return
	}

	result, err = client.ListWithSecretsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "ListWithSecrets", resp, "Failure responding to request")
	}

	return
}

// ListWithSecretsPreparer prepares the ListWithSecrets request.
func (client BotConnectionClient) ListWithSecretsPreparer(ctx context.Context, resourceGroupName string, resourceName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}/listWithSecrets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListWithSecretsSender sends the ListWithSecrets request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) ListWithSecretsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListWithSecretsResponder handles the response to the ListWithSecrets request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) ListWithSecretsResponder(resp *http.Response) (result ConnectionSetting, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a Connection Setting registration for a Bot Service
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// connectionName - the name of the Bot Service Connection Setting resource
// parameters - the parameters to provide for updating the Connection Setting.
func (client BotConnectionClient) Update(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting) (result ConnectionSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BotConnectionClient.Update")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: connectionName,
			Constraints: []validation.Constraint{{Target: "connectionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "connectionName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "connectionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.BotConnectionClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, connectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.BotConnectionClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BotConnectionClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BotConnectionClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BotConnectionClient) UpdateResponder(resp *http.Response) (result ConnectionSetting, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
