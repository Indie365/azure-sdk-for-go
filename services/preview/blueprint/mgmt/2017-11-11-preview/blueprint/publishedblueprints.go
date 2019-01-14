package blueprint

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

// PublishedBlueprintsClient is the blueprint Client
type PublishedBlueprintsClient struct {
	BaseClient
}

// NewPublishedBlueprintsClient creates an instance of the PublishedBlueprintsClient client.
func NewPublishedBlueprintsClient() PublishedBlueprintsClient {
	return NewPublishedBlueprintsClientWithBaseURI(DefaultBaseURI)
}

// NewPublishedBlueprintsClientWithBaseURI creates an instance of the PublishedBlueprintsClient client.
func NewPublishedBlueprintsClientWithBaseURI(baseURI string) PublishedBlueprintsClient {
	return PublishedBlueprintsClient{NewWithBaseURI(baseURI)}
}

// Create publish a new version of the Blueprint with the latest artifacts. Published Blueprints are immutable.
// Parameters:
// scope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// blueprintName - name of the blueprint.
// versionID - version of the published blueprint.
func (client PublishedBlueprintsClient) Create(ctx context.Context, scope string, blueprintName string, versionID string) (result PublishedBlueprint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublishedBlueprintsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, scope, blueprintName, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PublishedBlueprintsClient) CreatePreparer(ctx context.Context, scope string, blueprintName string, versionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName": autorest.Encode("path", blueprintName),
		"scope":         scope,
		"versionId":     autorest.Encode("path", versionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PublishedBlueprintsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PublishedBlueprintsClient) CreateResponder(resp *http.Response) (result PublishedBlueprint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a published Blueprint.
// Parameters:
// scope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// blueprintName - name of the blueprint.
// versionID - version of the published blueprint.
func (client PublishedBlueprintsClient) Delete(ctx context.Context, scope string, blueprintName string, versionID string) (result PublishedBlueprint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublishedBlueprintsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, blueprintName, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PublishedBlueprintsClient) DeletePreparer(ctx context.Context, scope string, blueprintName string, versionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName": autorest.Encode("path", blueprintName),
		"scope":         scope,
		"versionId":     autorest.Encode("path", versionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublishedBlueprintsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PublishedBlueprintsClient) DeleteResponder(resp *http.Response) (result PublishedBlueprint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a published Blueprint.
// Parameters:
// scope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// blueprintName - name of the blueprint.
// versionID - version of the published blueprint.
func (client PublishedBlueprintsClient) Get(ctx context.Context, scope string, blueprintName string, versionID string) (result PublishedBlueprint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublishedBlueprintsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, blueprintName, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PublishedBlueprintsClient) GetPreparer(ctx context.Context, scope string, blueprintName string, versionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName": autorest.Encode("path", blueprintName),
		"scope":         scope,
		"versionId":     autorest.Encode("path", versionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PublishedBlueprintsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PublishedBlueprintsClient) GetResponder(resp *http.Response) (result PublishedBlueprint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list published versions of given Blueprint.
// Parameters:
// scope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// blueprintName - name of the blueprint.
func (client PublishedBlueprintsClient) List(ctx context.Context, scope string, blueprintName string) (result PublishedBlueprintListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublishedBlueprintsClient.List")
		defer func() {
			sc := -1
			if result.pbl.Response.Response != nil {
				sc = result.pbl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, blueprintName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pbl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "List", resp, "Failure sending request")
		return
	}

	result.pbl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PublishedBlueprintsClient) ListPreparer(ctx context.Context, scope string, blueprintName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blueprintName": autorest.Encode("path", blueprintName),
		"scope":         scope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PublishedBlueprintsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PublishedBlueprintsClient) ListResponder(resp *http.Response) (result PublishedBlueprintList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PublishedBlueprintsClient) listNextResults(ctx context.Context, lastResults PublishedBlueprintList) (result PublishedBlueprintList, err error) {
	req, err := lastResults.publishedBlueprintListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.PublishedBlueprintsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublishedBlueprintsClient) ListComplete(ctx context.Context, scope string, blueprintName string) (result PublishedBlueprintListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublishedBlueprintsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope, blueprintName)
	return
}
