package kusto

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

// DatabasePrincipalAssignmentsClient is the the Azure Kusto management API provides a RESTful set of web services that
// interact with Azure Kusto services to manage your clusters and databases. The API enables you to create, update, and
// delete clusters and databases.
type DatabasePrincipalAssignmentsClient struct {
	BaseClient
}

// NewDatabasePrincipalAssignmentsClient creates an instance of the DatabasePrincipalAssignmentsClient client.
func NewDatabasePrincipalAssignmentsClient(subscriptionID string) DatabasePrincipalAssignmentsClient {
	return NewDatabasePrincipalAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabasePrincipalAssignmentsClientWithBaseURI creates an instance of the DatabasePrincipalAssignmentsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewDatabasePrincipalAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) DatabasePrincipalAssignmentsClient {
	return DatabasePrincipalAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks that the database principal assignment is valid and is not already in use.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// principalAssignmentName - the name of the resource.
func (client DatabasePrincipalAssignmentsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName DatabasePrincipalAssignmentCheckNameRequest) (result CheckNameResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasePrincipalAssignmentsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: principalAssignmentName,
			Constraints: []validation.Constraint{{Target: "principalAssignmentName.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "principalAssignmentName.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kusto.DatabasePrincipalAssignmentsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, resourceGroupName, clusterName, databaseName, principalAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client DatabasePrincipalAssignmentsClient) CheckNameAvailabilityPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName DatabasePrincipalAssignmentCheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/checkPrincipalAssignmentNameAvailability", pathParameters),
		autorest.WithJSON(principalAssignmentName),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasePrincipalAssignmentsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client DatabasePrincipalAssignmentsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates a Kusto cluster database principalAssignment.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// principalAssignmentName - the name of the Kusto principalAssignment.
// parameters - the Kusto principalAssignments parameters supplied for the operation.
func (client DatabasePrincipalAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string, parameters DatabasePrincipalAssignment) (result DatabasePrincipalAssignmentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasePrincipalAssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DatabasePrincipalProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.DatabasePrincipalProperties.PrincipalID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("kusto.DatabasePrincipalAssignmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, principalAssignmentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DatabasePrincipalAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string, parameters DatabasePrincipalAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":             autorest.Encode("path", clusterName),
		"databaseName":            autorest.Encode("path", databaseName),
		"principalAssignmentName": autorest.Encode("path", principalAssignmentName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/principalAssignments/{principalAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasePrincipalAssignmentsClient) CreateOrUpdateSender(req *http.Request) (future DatabasePrincipalAssignmentsCreateOrUpdateFuture, err error) {
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
func (client DatabasePrincipalAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result DatabasePrincipalAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Kusto principalAssignment.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// principalAssignmentName - the name of the Kusto principalAssignment.
func (client DatabasePrincipalAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string) (result DatabasePrincipalAssignmentsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasePrincipalAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, databaseName, principalAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DatabasePrincipalAssignmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":             autorest.Encode("path", clusterName),
		"databaseName":            autorest.Encode("path", databaseName),
		"principalAssignmentName": autorest.Encode("path", principalAssignmentName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/principalAssignments/{principalAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasePrincipalAssignmentsClient) DeleteSender(req *http.Request) (future DatabasePrincipalAssignmentsDeleteFuture, err error) {
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
func (client DatabasePrincipalAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Kusto cluster database principalAssignment.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// principalAssignmentName - the name of the Kusto principalAssignment.
func (client DatabasePrincipalAssignmentsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string) (result DatabasePrincipalAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasePrincipalAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, databaseName, principalAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DatabasePrincipalAssignmentsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, principalAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":             autorest.Encode("path", clusterName),
		"databaseName":            autorest.Encode("path", databaseName),
		"principalAssignmentName": autorest.Encode("path", principalAssignmentName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/principalAssignments/{principalAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasePrincipalAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabasePrincipalAssignmentsClient) GetResponder(resp *http.Response) (result DatabasePrincipalAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all Kusto cluster database principalAssignments.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
func (client DatabasePrincipalAssignmentsClient) List(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DatabasePrincipalAssignmentListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasePrincipalAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DatabasePrincipalAssignmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DatabasePrincipalAssignmentsClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/principalAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasePrincipalAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DatabasePrincipalAssignmentsClient) ListResponder(resp *http.Response) (result DatabasePrincipalAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
