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

// TriggerClient is the client for the Trigger methods of the Artifacts service.
type TriggerClient struct {
	BaseClient
}

// NewTriggerClient creates an instance of the TriggerClient client.
func NewTriggerClient(endpoint string) TriggerClient {
	return TriggerClient{New(endpoint)}
}

// CreateOrUpdateTrigger creates or updates a trigger.
// Parameters:
// triggerName - the trigger name.
// trigger - trigger resource definition.
// ifMatch - eTag of the trigger entity.  Should only be specified for update, for which it should match
// existing entity or can be * for unconditional update.
func (client TriggerClient) CreateOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, ifMatch string) (result TriggerResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.CreateOrUpdateTrigger")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: trigger,
			Constraints: []validation.Constraint{{Target: "trigger.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "CreateOrUpdateTrigger", err.Error())
	}

	req, err := client.CreateOrUpdateTriggerPreparer(ctx, triggerName, trigger, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "CreateOrUpdateTrigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateTriggerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "CreateOrUpdateTrigger", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateTriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "CreateOrUpdateTrigger", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateTriggerPreparer prepares the CreateOrUpdateTrigger request.
func (client TriggerClient) CreateOrUpdateTriggerPreparer(ctx context.Context, triggerName string, trigger TriggerResource, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}", pathParameters),
		autorest.WithJSON(trigger),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateTriggerSender sends the CreateOrUpdateTrigger request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) CreateOrUpdateTriggerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateTriggerResponder handles the response to the CreateOrUpdateTrigger request. The method always
// closes the http.Response Body.
func (client TriggerClient) CreateOrUpdateTriggerResponder(resp *http.Response) (result TriggerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteTrigger deletes a trigger.
// Parameters:
// triggerName - the trigger name.
func (client TriggerClient) DeleteTrigger(ctx context.Context, triggerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.DeleteTrigger")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "DeleteTrigger", err.Error())
	}

	req, err := client.DeleteTriggerPreparer(ctx, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "DeleteTrigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteTriggerSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "DeleteTrigger", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteTriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "DeleteTrigger", resp, "Failure responding to request")
	}

	return
}

// DeleteTriggerPreparer prepares the DeleteTrigger request.
func (client TriggerClient) DeleteTriggerPreparer(ctx context.Context, triggerName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteTriggerSender sends the DeleteTrigger request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) DeleteTriggerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteTriggerResponder handles the response to the DeleteTrigger request. The method always
// closes the http.Response Body.
func (client TriggerClient) DeleteTriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetEventSubscriptionStatus get a trigger's event subscription status.
// Parameters:
// triggerName - the trigger name.
func (client TriggerClient) GetEventSubscriptionStatus(ctx context.Context, triggerName string) (result TriggerSubscriptionOperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.GetEventSubscriptionStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "GetEventSubscriptionStatus", err.Error())
	}

	req, err := client.GetEventSubscriptionStatusPreparer(ctx, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetEventSubscriptionStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEventSubscriptionStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetEventSubscriptionStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetEventSubscriptionStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetEventSubscriptionStatus", resp, "Failure responding to request")
	}

	return
}

// GetEventSubscriptionStatusPreparer prepares the GetEventSubscriptionStatus request.
func (client TriggerClient) GetEventSubscriptionStatusPreparer(ctx context.Context, triggerName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/getEventSubscriptionStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEventSubscriptionStatusSender sends the GetEventSubscriptionStatus request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) GetEventSubscriptionStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEventSubscriptionStatusResponder handles the response to the GetEventSubscriptionStatus request. The method always
// closes the http.Response Body.
func (client TriggerClient) GetEventSubscriptionStatusResponder(resp *http.Response) (result TriggerSubscriptionOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTrigger gets a trigger.
// Parameters:
// triggerName - the trigger name.
// ifNoneMatch - eTag of the trigger entity. Should only be specified for get. If the ETag matches the existing
// entity tag, or if * was provided, then no content will be returned.
func (client TriggerClient) GetTrigger(ctx context.Context, triggerName string, ifNoneMatch string) (result TriggerResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.GetTrigger")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "GetTrigger", err.Error())
	}

	req, err := client.GetTriggerPreparer(ctx, triggerName, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetTrigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTriggerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetTrigger", resp, "Failure sending request")
		return
	}

	result, err = client.GetTriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetTrigger", resp, "Failure responding to request")
	}

	return
}

// GetTriggerPreparer prepares the GetTrigger request.
func (client TriggerClient) GetTriggerPreparer(ctx context.Context, triggerName string, ifNoneMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTriggerSender sends the GetTrigger request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) GetTriggerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTriggerResponder handles the response to the GetTrigger request. The method always
// closes the http.Response Body.
func (client TriggerClient) GetTriggerResponder(resp *http.Response) (result TriggerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotModified),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTriggersByWorkspace lists triggers.
func (client TriggerClient) GetTriggersByWorkspace(ctx context.Context) (result TriggerListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.GetTriggersByWorkspace")
		defer func() {
			sc := -1
			if result.tlr.Response.Response != nil {
				sc = result.tlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getTriggersByWorkspaceNextResults
	req, err := client.GetTriggersByWorkspacePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetTriggersByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTriggersByWorkspaceSender(req)
	if err != nil {
		result.tlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetTriggersByWorkspace", resp, "Failure sending request")
		return
	}

	result.tlr, err = client.GetTriggersByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "GetTriggersByWorkspace", resp, "Failure responding to request")
	}

	return
}

// GetTriggersByWorkspacePreparer prepares the GetTriggersByWorkspace request.
func (client TriggerClient) GetTriggersByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
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
		autorest.WithPath("/triggers"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTriggersByWorkspaceSender sends the GetTriggersByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) GetTriggersByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTriggersByWorkspaceResponder handles the response to the GetTriggersByWorkspace request. The method always
// closes the http.Response Body.
func (client TriggerClient) GetTriggersByWorkspaceResponder(resp *http.Response) (result TriggerListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getTriggersByWorkspaceNextResults retrieves the next set of results, if any.
func (client TriggerClient) getTriggersByWorkspaceNextResults(ctx context.Context, lastResults TriggerListResponse) (result TriggerListResponse, err error) {
	req, err := lastResults.triggerListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "artifacts.TriggerClient", "getTriggersByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetTriggersByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "artifacts.TriggerClient", "getTriggersByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetTriggersByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "getTriggersByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetTriggersByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client TriggerClient) GetTriggersByWorkspaceComplete(ctx context.Context) (result TriggerListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.GetTriggersByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetTriggersByWorkspace(ctx)
	return
}

// StartTrigger starts a trigger.
// Parameters:
// triggerName - the trigger name.
func (client TriggerClient) StartTrigger(ctx context.Context, triggerName string) (result TriggerStartTriggerFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.StartTrigger")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "StartTrigger", err.Error())
	}

	req, err := client.StartTriggerPreparer(ctx, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "StartTrigger", nil, "Failure preparing request")
		return
	}

	result, err = client.StartTriggerSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "StartTrigger", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartTriggerPreparer prepares the StartTrigger request.
func (client TriggerClient) StartTriggerPreparer(ctx context.Context, triggerName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartTriggerSender sends the StartTrigger request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) StartTriggerSender(req *http.Request) (future TriggerStartTriggerFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StartTriggerResponder handles the response to the StartTrigger request. The method always
// closes the http.Response Body.
func (client TriggerClient) StartTriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// StopTrigger stops a trigger.
// Parameters:
// triggerName - the trigger name.
func (client TriggerClient) StopTrigger(ctx context.Context, triggerName string) (result TriggerStopTriggerFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.StopTrigger")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "StopTrigger", err.Error())
	}

	req, err := client.StopTriggerPreparer(ctx, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "StopTrigger", nil, "Failure preparing request")
		return
	}

	result, err = client.StopTriggerSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "StopTrigger", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopTriggerPreparer prepares the StopTrigger request.
func (client TriggerClient) StopTriggerPreparer(ctx context.Context, triggerName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopTriggerSender sends the StopTrigger request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) StopTriggerSender(req *http.Request) (future TriggerStopTriggerFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopTriggerResponder handles the response to the StopTrigger request. The method always
// closes the http.Response Body.
func (client TriggerClient) StopTriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// SubscribeTriggerToEvents subscribe event trigger to events.
// Parameters:
// triggerName - the trigger name.
func (client TriggerClient) SubscribeTriggerToEvents(ctx context.Context, triggerName string) (result TriggerSubscribeTriggerToEventsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.SubscribeTriggerToEvents")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "SubscribeTriggerToEvents", err.Error())
	}

	req, err := client.SubscribeTriggerToEventsPreparer(ctx, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "SubscribeTriggerToEvents", nil, "Failure preparing request")
		return
	}

	result, err = client.SubscribeTriggerToEventsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "SubscribeTriggerToEvents", result.Response(), "Failure sending request")
		return
	}

	return
}

// SubscribeTriggerToEventsPreparer prepares the SubscribeTriggerToEvents request.
func (client TriggerClient) SubscribeTriggerToEventsPreparer(ctx context.Context, triggerName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/subscribeToEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SubscribeTriggerToEventsSender sends the SubscribeTriggerToEvents request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) SubscribeTriggerToEventsSender(req *http.Request) (future TriggerSubscribeTriggerToEventsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// SubscribeTriggerToEventsResponder handles the response to the SubscribeTriggerToEvents request. The method always
// closes the http.Response Body.
func (client TriggerClient) SubscribeTriggerToEventsResponder(resp *http.Response) (result TriggerSubscriptionOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UnsubscribeTriggerFromEvents unsubscribe event trigger from events.
// Parameters:
// triggerName - the trigger name.
func (client TriggerClient) UnsubscribeTriggerFromEvents(ctx context.Context, triggerName string) (result TriggerUnsubscribeTriggerFromEventsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerClient.UnsubscribeTriggerFromEvents")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("artifacts.TriggerClient", "UnsubscribeTriggerFromEvents", err.Error())
	}

	req, err := client.UnsubscribeTriggerFromEventsPreparer(ctx, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "UnsubscribeTriggerFromEvents", nil, "Failure preparing request")
		return
	}

	result, err = client.UnsubscribeTriggerFromEventsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "artifacts.TriggerClient", "UnsubscribeTriggerFromEvents", result.Response(), "Failure sending request")
		return
	}

	return
}

// UnsubscribeTriggerFromEventsPreparer prepares the UnsubscribeTriggerFromEvents request.
func (client TriggerClient) UnsubscribeTriggerFromEventsPreparer(ctx context.Context, triggerName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/unsubscribeFromEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UnsubscribeTriggerFromEventsSender sends the UnsubscribeTriggerFromEvents request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerClient) UnsubscribeTriggerFromEventsSender(req *http.Request) (future TriggerUnsubscribeTriggerFromEventsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UnsubscribeTriggerFromEventsResponder handles the response to the UnsubscribeTriggerFromEvents request. The method always
// closes the http.Response Body.
func (client TriggerClient) UnsubscribeTriggerFromEventsResponder(resp *http.Response) (result TriggerSubscriptionOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
