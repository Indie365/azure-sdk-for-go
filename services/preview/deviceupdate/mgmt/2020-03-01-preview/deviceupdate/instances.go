package deviceupdate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// InstancesClient is the microsoft Device Update resource provider.
type InstancesClient struct {
	BaseClient
}

// NewInstancesClient creates an instance of the InstancesClient client.
func NewInstancesClient(subscriptionID string) InstancesClient {
	return NewInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInstancesClientWithBaseURI creates an instance of the InstancesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInstancesClientWithBaseURI(baseURI string, subscriptionID string) InstancesClient {
	return InstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates or updates instance.
// Parameters:
// resourceGroupName - the resource group name.
// accountName - account name.
// instanceName - instance name.
// instance - instance details.
func (client InstancesClient) Create(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance) (result InstancesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}},
		{TargetValue: instanceName,
			Constraints: []validation.Constraint{{Target: "instanceName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "instanceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "instanceName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}},
		{TargetValue: instance,
			Constraints: []validation.Constraint{{Target: "instance.InstanceProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deviceupdate.InstancesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, instanceName, instance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client InstancesClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"instanceName":      autorest.Encode("path", instanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}", pathParameters),
		autorest.WithJSON(instance),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client InstancesClient) CreateSender(req *http.Request) (future InstancesCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client InstancesClient) (i Instance, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "deviceupdate.InstancesCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("deviceupdate.InstancesCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		i.Response.Response, err = future.GetResult(sender)
		if i.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "deviceupdate.InstancesCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && i.Response.Response.StatusCode != http.StatusNoContent {
			i, err = client.CreateResponder(i.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deviceupdate.InstancesCreateFuture", "Result", i.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client InstancesClient) CreateResponder(resp *http.Response) (result Instance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes instance.
// Parameters:
// resourceGroupName - the resource group name.
// accountName - account name.
// instanceName - instance name.
func (client InstancesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, instanceName string) (result InstancesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}},
		{TargetValue: instanceName,
			Constraints: []validation.Constraint{{Target: "instanceName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "instanceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "instanceName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deviceupdate.InstancesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, instanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InstancesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, instanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"instanceName":      autorest.Encode("path", instanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InstancesClient) DeleteSender(req *http.Request) (future InstancesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client InstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "deviceupdate.InstancesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("deviceupdate.InstancesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns instances for the given account and instance name.
// Parameters:
// resourceGroupName - the resource group name.
// accountName - account name.
// instanceName - instance name.
func (client InstancesClient) Get(ctx context.Context, resourceGroupName string, accountName string, instanceName string) (result Instance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}},
		{TargetValue: instanceName,
			Constraints: []validation.Constraint{{Target: "instanceName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "instanceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "instanceName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deviceupdate.InstancesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, instanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client InstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, instanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"instanceName":      autorest.Encode("path", instanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InstancesClient) GetResponder(resp *http.Response) (result Instance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccount returns instances for the given account name.
// Parameters:
// resourceGroupName - the resource group name.
// accountName - account name.
func (client InstancesClient) ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result InstanceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.ListByAccount")
		defer func() {
			sc := -1
			if result.il.Response.Response != nil {
				sc = result.il.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deviceupdate.InstancesClient", "ListByAccount", err.Error())
	}

	result.fn = client.listByAccountNextResults
	req, err := client.ListByAccountPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "ListByAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.il.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "ListByAccount", resp, "Failure sending request")
		return
	}

	result.il, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "ListByAccount", resp, "Failure responding to request")
		return
	}
	if result.il.hasNextLink() && result.il.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByAccountPreparer prepares the ListByAccount request.
func (client InstancesClient) ListByAccountPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAccountSender sends the ListByAccount request. The method will close the
// http.Response Body if it receives an error.
func (client InstancesClient) ListByAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAccountResponder handles the response to the ListByAccount request. The method always
// closes the http.Response Body.
func (client InstancesClient) ListByAccountResponder(resp *http.Response) (result InstanceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAccountNextResults retrieves the next set of results, if any.
func (client InstancesClient) listByAccountNextResults(ctx context.Context, lastResults InstanceList) (result InstanceList, err error) {
	req, err := lastResults.instanceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "listByAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "listByAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "listByAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client InstancesClient) ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result InstanceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.ListByAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAccount(ctx, resourceGroupName, accountName)
	return
}

// ListBySubscription returns instances for the given account name.
// Parameters:
// accountName - account name.
func (client InstancesClient) ListBySubscription(ctx context.Context, accountName string) (result InstanceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.il.Response.Response != nil {
				sc = result.il.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deviceupdate.InstancesClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.il.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.il, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.il.hasNextLink() && result.il.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client InstancesClient) ListBySubscriptionPreparer(ctx context.Context, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":    autorest.Encode("path", accountName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client InstancesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client InstancesClient) ListBySubscriptionResponder(resp *http.Response) (result InstanceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client InstancesClient) listBySubscriptionNextResults(ctx context.Context, lastResults InstanceList) (result InstanceList, err error) {
	req, err := lastResults.instanceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client InstancesClient) ListBySubscriptionComplete(ctx context.Context, accountName string) (result InstanceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, accountName)
	return
}

// Update updates instance's tags.
// Parameters:
// resourceGroupName - the resource group name.
// accountName - account name.
// instanceName - instance name.
// tagUpdatePayload - updated tags.
func (client InstancesClient) Update(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload TagUpdate) (result Instance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InstancesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}},
		{TargetValue: instanceName,
			Constraints: []validation.Constraint{{Target: "instanceName", Name: validation.MaxLength, Rule: 36, Chain: nil},
				{Target: "instanceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "instanceName", Name: validation.Pattern, Rule: `^[A-Za-z][A-Za-z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("deviceupdate.InstancesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, instanceName, tagUpdatePayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deviceupdate.InstancesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client InstancesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload TagUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"instanceName":      autorest.Encode("path", instanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}", pathParameters),
		autorest.WithJSON(tagUpdatePayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client InstancesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client InstancesClient) UpdateResponder(resp *http.Response) (result Instance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
