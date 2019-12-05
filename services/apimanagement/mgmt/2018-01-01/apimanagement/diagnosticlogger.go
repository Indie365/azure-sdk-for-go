package apimanagement

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

// DiagnosticLoggerClient is the apiManagement Client
type DiagnosticLoggerClient struct {
	BaseClient
}

// NewDiagnosticLoggerClient creates an instance of the DiagnosticLoggerClient client.
func NewDiagnosticLoggerClient(subscriptionID string) DiagnosticLoggerClient {
	return NewDiagnosticLoggerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDiagnosticLoggerClientWithBaseURI creates an instance of the DiagnosticLoggerClient client.
func NewDiagnosticLoggerClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticLoggerClient {
	return DiagnosticLoggerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckEntityExists checks that logger entity specified by identifier is associated with the diagnostics entity.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// diagnosticID - diagnostic identifier. Must be unique in the current API Management service instance.
// loggerid - logger identifier. Must be unique in the API Management service instance.
func (client DiagnosticLoggerClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiagnosticLoggerClient.CheckEntityExists")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: diagnosticID,
			Constraints: []validation.Constraint{{Target: "diagnosticID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "diagnosticID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diagnosticID", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: loggerid,
			Constraints: []validation.Constraint{{Target: "loggerid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "loggerid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.DiagnosticLoggerClient", "CheckEntityExists", err.Error())
	}

	req, err := client.CheckEntityExistsPreparer(ctx, resourceGroupName, serviceName, diagnosticID, loggerid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "CheckEntityExists", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckEntityExistsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "CheckEntityExists", resp, "Failure sending request")
		return
	}

	result, err = client.CheckEntityExistsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "CheckEntityExists", resp, "Failure responding to request")
	}

	return
}

// CheckEntityExistsPreparer prepares the CheckEntityExists request.
func (client DiagnosticLoggerClient) CheckEntityExistsPreparer(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diagnosticId":      autorest.Encode("path", diagnosticID),
		"loggerid":          autorest.Encode("path", loggerid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers/{loggerid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckEntityExistsSender sends the CheckEntityExists request. The method will close the
// http.Response Body if it receives an error.
func (client DiagnosticLoggerClient) CheckEntityExistsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckEntityExistsResponder handles the response to the CheckEntityExists request. The method always
// closes the http.Response Body.
func (client DiagnosticLoggerClient) CheckEntityExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate attaches a logger to a diagnostic.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// diagnosticID - diagnostic identifier. Must be unique in the current API Management service instance.
// loggerid - logger identifier. Must be unique in the API Management service instance.
func (client DiagnosticLoggerClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string) (result LoggerContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiagnosticLoggerClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: diagnosticID,
			Constraints: []validation.Constraint{{Target: "diagnosticID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "diagnosticID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diagnosticID", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: loggerid,
			Constraints: []validation.Constraint{{Target: "loggerid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "loggerid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.DiagnosticLoggerClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, diagnosticID, loggerid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DiagnosticLoggerClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diagnosticId":      autorest.Encode("path", diagnosticID),
		"loggerid":          autorest.Encode("path", loggerid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers/{loggerid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DiagnosticLoggerClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DiagnosticLoggerClient) CreateOrUpdateResponder(resp *http.Response) (result LoggerContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Logger from Diagnostic.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// diagnosticID - diagnostic identifier. Must be unique in the current API Management service instance.
// loggerid - logger identifier. Must be unique in the API Management service instance.
func (client DiagnosticLoggerClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiagnosticLoggerClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: diagnosticID,
			Constraints: []validation.Constraint{{Target: "diagnosticID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "diagnosticID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diagnosticID", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: loggerid,
			Constraints: []validation.Constraint{{Target: "loggerid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "loggerid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.DiagnosticLoggerClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, diagnosticID, loggerid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DiagnosticLoggerClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, loggerid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diagnosticId":      autorest.Encode("path", diagnosticID),
		"loggerid":          autorest.Encode("path", loggerid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers/{loggerid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DiagnosticLoggerClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DiagnosticLoggerClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService lists all loggers associated with the specified Diagnostic of the API Management service instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// diagnosticID - diagnostic identifier. Must be unique in the current API Management service instance.
// filter - | Field       | Supported operators    | Supported functions               |
// |-------------|------------------------|-----------------------------------|
// | id          | ge, le, eq, ne, gt, lt | substringof, startswith, endswith |
// | type        | eq                     |                                   |
// top - number of records to return.
// skip - number of records to skip.
func (client DiagnosticLoggerClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, filter string, top *int32, skip *int32) (result LoggerCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiagnosticLoggerClient.ListByService")
		defer func() {
			sc := -1
			if result.lc.Response.Response != nil {
				sc = result.lc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: diagnosticID,
			Constraints: []validation.Constraint{{Target: "diagnosticID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "diagnosticID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diagnosticID", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.DiagnosticLoggerClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, diagnosticID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.lc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.lc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client DiagnosticLoggerClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diagnosticId":      autorest.Encode("path", diagnosticID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}/loggers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client DiagnosticLoggerClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client DiagnosticLoggerClient) ListByServiceResponder(resp *http.Response) (result LoggerCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client DiagnosticLoggerClient) listByServiceNextResults(ctx context.Context, lastResults LoggerCollection) (result LoggerCollection, err error) {
	req, err := lastResults.loggerCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.DiagnosticLoggerClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiagnosticLoggerClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, diagnosticID string, filter string, top *int32, skip *int32) (result LoggerCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiagnosticLoggerClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, diagnosticID, filter, top, skip)
	return
}
