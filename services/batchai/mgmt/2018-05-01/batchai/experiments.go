package batchai

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

// ExperimentsClient is the the Azure BatchAI Management API.
type ExperimentsClient struct {
	BaseClient
}

// NewExperimentsClient creates an instance of the ExperimentsClient client.
func NewExperimentsClient(subscriptionID string) ExperimentsClient {
	return NewExperimentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExperimentsClientWithBaseURI creates an instance of the ExperimentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExperimentsClientWithBaseURI(baseURI string, subscriptionID string) ExperimentsClient {
	return ExperimentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates an Experiment.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// experimentName - the name of the experiment. Experiment names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client ExperimentsClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (result ExperimentsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExperimentsClient.Create")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: experimentName,
			Constraints: []validation.Constraint{{Target: "experimentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "experimentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "experimentName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ExperimentsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceName, experimentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ExperimentsClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"experimentName":    autorest.Encode("path", experimentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ExperimentsClient) CreateSender(req *http.Request) (future ExperimentsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ExperimentsClient) CreateResponder(resp *http.Response) (result Experiment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an Experiment.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// experimentName - the name of the experiment. Experiment names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client ExperimentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (result ExperimentsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExperimentsClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: experimentName,
			Constraints: []validation.Constraint{{Target: "experimentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "experimentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "experimentName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ExperimentsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, experimentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExperimentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"experimentName":    autorest.Encode("path", experimentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExperimentsClient) DeleteSender(req *http.Request) (future ExperimentsDeleteFuture, err error) {
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
func (client ExperimentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about an Experiment.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// experimentName - the name of the experiment. Experiment names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client ExperimentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (result Experiment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExperimentsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: experimentName,
			Constraints: []validation.Constraint{{Target: "experimentName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "experimentName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "experimentName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ExperimentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, experimentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExperimentsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"experimentName":    autorest.Encode("path", experimentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExperimentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExperimentsClient) GetResponder(resp *http.Response) (result Experiment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace gets a list of Experiments within the specified Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client ExperimentsClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result ExperimentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExperimentsClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.elr.Response.Response != nil {
				sc = result.elr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ExperimentsClient", "ListByWorkspace", err.Error())
	}

	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.elr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.elr, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "ListByWorkspace", resp, "Failure responding to request")
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client ExperimentsClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client ExperimentsClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client ExperimentsClient) ListByWorkspaceResponder(resp *http.Response) (result ExperimentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client ExperimentsClient) listByWorkspaceNextResults(ctx context.Context, lastResults ExperimentListResult) (result ExperimentListResult, err error) {
	req, err := lastResults.experimentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ExperimentsClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExperimentsClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result ExperimentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExperimentsClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName, maxResults)
	return
}
