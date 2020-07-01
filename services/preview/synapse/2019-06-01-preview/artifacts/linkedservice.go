package artifacts

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

// LinkedServiceClient is the client for the LinkedService methods of the Artifacts service.
type LinkedServiceClient struct {
	BaseClient
}

// NewLinkedServiceClient creates an instance of the LinkedServiceClient client.
func NewLinkedServiceClient(endpoint string) LinkedServiceClient {
	return LinkedServiceClient{New(endpoint)}
}

// CreateOrUpdateLinkedService creates or updates a linked service.
// Parameters:
// linkedServiceName - the linked service name.
// linkedService - linked service resource definition.
// ifMatch - eTag of the linkedService entity.  Should only be specified for update, for which it should match
// existing entity or can be * for unconditional update.
func (client LinkedServiceClient) CreateOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, ifMatch string) (result LinkedServiceResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServiceClient.CreateOrUpdateLinkedService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: linkedServiceName,
			Constraints: []validation.Constraint{{Target: "linkedServiceName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "linkedServiceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "linkedServiceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: linkedService,
			Constraints: []validation.Constraint{{Target: "linkedService.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "linkedService.Properties.ConnectVia", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "linkedService.Properties.ConnectVia.Type", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "linkedService.Properties.ConnectVia.ReferenceName", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("artifacts.LinkedServiceClient", "CreateOrUpdateLinkedService", err.Error())
	}

	req, err := client.CreateOrUpdateLinkedServicePreparer(ctx, linkedServiceName, linkedService, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "CreateOrUpdateLinkedService", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateLinkedServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "CreateOrUpdateLinkedService", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateLinkedServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "CreateOrUpdateLinkedService", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateLinkedServicePreparer prepares the CreateOrUpdateLinkedService request.
func (client LinkedServiceClient) CreateOrUpdateLinkedServicePreparer(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"linkedServiceName": autorest.Encode("path", linkedServiceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/linkedservices/{linkedServiceName}", pathParameters),
		autorest.WithJSON(linkedService),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateLinkedServiceSender sends the CreateOrUpdateLinkedService request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServiceClient) CreateOrUpdateLinkedServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateLinkedServiceResponder handles the response to the CreateOrUpdateLinkedService request. The method always
// closes the http.Response Body.
func (client LinkedServiceClient) CreateOrUpdateLinkedServiceResponder(resp *http.Response) (result LinkedServiceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteLinkedService deletes a linked service.
// Parameters:
// linkedServiceName - the linked service name.
func (client LinkedServiceClient) DeleteLinkedService(ctx context.Context, linkedServiceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServiceClient.DeleteLinkedService")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: linkedServiceName,
			Constraints: []validation.Constraint{{Target: "linkedServiceName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "linkedServiceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "linkedServiceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.LinkedServiceClient", "DeleteLinkedService", err.Error())
	}

	req, err := client.DeleteLinkedServicePreparer(ctx, linkedServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "DeleteLinkedService", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteLinkedServiceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "DeleteLinkedService", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteLinkedServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "DeleteLinkedService", resp, "Failure responding to request")
	}

	return
}

// DeleteLinkedServicePreparer prepares the DeleteLinkedService request.
func (client LinkedServiceClient) DeleteLinkedServicePreparer(ctx context.Context, linkedServiceName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"linkedServiceName": autorest.Encode("path", linkedServiceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/linkedservices/{linkedServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteLinkedServiceSender sends the DeleteLinkedService request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServiceClient) DeleteLinkedServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteLinkedServiceResponder handles the response to the DeleteLinkedService request. The method always
// closes the http.Response Body.
func (client LinkedServiceClient) DeleteLinkedServiceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetLinkedService gets a linked service.
// Parameters:
// linkedServiceName - the linked service name.
// ifNoneMatch - eTag of the linked service entity. Should only be specified for get. If the ETag matches the
// existing entity tag, or if * was provided, then no content will be returned.
func (client LinkedServiceClient) GetLinkedService(ctx context.Context, linkedServiceName string, ifNoneMatch string) (result LinkedServiceResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServiceClient.GetLinkedService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: linkedServiceName,
			Constraints: []validation.Constraint{{Target: "linkedServiceName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "linkedServiceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "linkedServiceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.LinkedServiceClient", "GetLinkedService", err.Error())
	}

	req, err := client.GetLinkedServicePreparer(ctx, linkedServiceName, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "GetLinkedService", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLinkedServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "GetLinkedService", resp, "Failure sending request")
		return
	}

	result, err = client.GetLinkedServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "GetLinkedService", resp, "Failure responding to request")
	}

	return
}

// GetLinkedServicePreparer prepares the GetLinkedService request.
func (client LinkedServiceClient) GetLinkedServicePreparer(ctx context.Context, linkedServiceName string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"linkedServiceName": autorest.Encode("path", linkedServiceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/linkedservices/{linkedServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLinkedServiceSender sends the GetLinkedService request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServiceClient) GetLinkedServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetLinkedServiceResponder handles the response to the GetLinkedService request. The method always
// closes the http.Response Body.
func (client LinkedServiceClient) GetLinkedServiceResponder(resp *http.Response) (result LinkedServiceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotModified),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLinkedServicesByWorkspace lists linked services.
func (client LinkedServiceClient) GetLinkedServicesByWorkspace(ctx context.Context) (result LinkedServiceListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServiceClient.GetLinkedServicesByWorkspace")
		defer func() {
			sc := -1
			if result.lslr.Response.Response != nil {
				sc = result.lslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getLinkedServicesByWorkspaceNextResults
	req, err := client.GetLinkedServicesByWorkspacePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "GetLinkedServicesByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLinkedServicesByWorkspaceSender(req)
	if err != nil {
		result.lslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "GetLinkedServicesByWorkspace", resp, "Failure sending request")
		return
	}

	result.lslr, err = client.GetLinkedServicesByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "GetLinkedServicesByWorkspace", resp, "Failure responding to request")
	}

	return
}

// GetLinkedServicesByWorkspacePreparer prepares the GetLinkedServicesByWorkspace request.
func (client LinkedServiceClient) GetLinkedServicesByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/linkedservices"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLinkedServicesByWorkspaceSender sends the GetLinkedServicesByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServiceClient) GetLinkedServicesByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetLinkedServicesByWorkspaceResponder handles the response to the GetLinkedServicesByWorkspace request. The method always
// closes the http.Response Body.
func (client LinkedServiceClient) GetLinkedServicesByWorkspaceResponder(resp *http.Response) (result LinkedServiceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getLinkedServicesByWorkspaceNextResults retrieves the next set of results, if any.
func (client LinkedServiceClient) getLinkedServicesByWorkspaceNextResults(ctx context.Context, lastResults LinkedServiceListResponse) (result LinkedServiceListResponse, err error) {
	req, err := lastResults.linkedServiceListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "getLinkedServicesByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetLinkedServicesByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "getLinkedServicesByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetLinkedServicesByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.LinkedServiceClient", "getLinkedServicesByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetLinkedServicesByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client LinkedServiceClient) GetLinkedServicesByWorkspaceComplete(ctx context.Context) (result LinkedServiceListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServiceClient.GetLinkedServicesByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetLinkedServicesByWorkspace(ctx)
	return
}
