package managednetwork

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ScopeAssignmentsClient is the the Microsoft Azure Managed Network management API provides a RESTful set of web
// services that interact with Microsoft Azure Networks service to programmatically view, control, change, and monitor
// your entire Azure network centrally and with ease.
type ScopeAssignmentsClient struct {
	BaseClient
}

// NewScopeAssignmentsClient creates an instance of the ScopeAssignmentsClient client.
func NewScopeAssignmentsClient(subscriptionID string) ScopeAssignmentsClient {
	return NewScopeAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScopeAssignmentsClientWithBaseURI creates an instance of the ScopeAssignmentsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewScopeAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ScopeAssignmentsClient {
	return ScopeAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a scope assignment.
// Parameters:
// parameters - parameters supplied to the specify which Managed Network this scope is being assigned
// scope - the base resource of the scope assignment to create. The scope can be any REST resource instance.
// For example, use 'subscriptions/{subscription-id}' for a subscription,
// 'subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
// 'subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource.
// scopeAssignmentName - the name of the scope assignment to create.
func (client ScopeAssignmentsClient) CreateOrUpdate(ctx context.Context, parameters ScopeAssignment, scope string, scopeAssignmentName string) (result ScopeAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScopeAssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters, scope, scopeAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ScopeAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, parameters ScopeAssignment, scope string, scopeAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope":               scope,
		"scopeAssignmentName": autorest.Encode("path", scopeAssignmentName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments/{scopeAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ScopeAssignmentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ScopeAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result ScopeAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a scope assignment.
// Parameters:
// scope - the scope of the scope assignment to delete.
// scopeAssignmentName - the name of the scope assignment to delete.
func (client ScopeAssignmentsClient) Delete(ctx context.Context, scope string, scopeAssignmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScopeAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, scopeAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ScopeAssignmentsClient) DeletePreparer(ctx context.Context, scope string, scopeAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope":               scope,
		"scopeAssignmentName": autorest.Encode("path", scopeAssignmentName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments/{scopeAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ScopeAssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScopeAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified scope assignment.
// Parameters:
// scope - the base resource of the scope assignment.
// scopeAssignmentName - the name of the scope assignment to get.
func (client ScopeAssignmentsClient) Get(ctx context.Context, scope string, scopeAssignmentName string) (result ScopeAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScopeAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, scopeAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ScopeAssignmentsClient) GetPreparer(ctx context.Context, scope string, scopeAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope":               scope,
		"scopeAssignmentName": autorest.Encode("path", scopeAssignmentName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments/{scopeAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ScopeAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScopeAssignmentsClient) GetResponder(resp *http.Response) (result ScopeAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get the specified scope assignment.
// Parameters:
// scope - the base resource of the scope assignment.
func (client ScopeAssignmentsClient) List(ctx context.Context, scope string) (result ScopeAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScopeAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.salr.Response.Response != nil {
				sc = result.salr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.salr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.salr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ScopeAssignmentsClient) ListPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScopeAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScopeAssignmentsClient) ListResponder(resp *http.Response) (result ScopeAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ScopeAssignmentsClient) listNextResults(ctx context.Context, lastResults ScopeAssignmentListResult) (result ScopeAssignmentListResult, err error) {
	req, err := lastResults.scopeAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managednetwork.ScopeAssignmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ScopeAssignmentsClient) ListComplete(ctx context.Context, scope string) (result ScopeAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScopeAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope)
	return
}
