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

// UserSubscriptionClient is the client for the UserSubscription methods of the Apimanagement service.
type UserSubscriptionClient struct {
	BaseClient
}

// NewUserSubscriptionClient creates an instance of the UserSubscriptionClient client.
func NewUserSubscriptionClient() UserSubscriptionClient {
	return UserSubscriptionClient{New()}
}

// List lists the collection of subscriptions of the specified user.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// UID - user identifier. Must be unique in the current API Management service instance.
// filter - | Field        | Supported operators    | Supported functions                         |
// |--------------|------------------------|---------------------------------------------|
// | id           | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | stateComment | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | userId       | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | productId    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | state        | eq                     |                                             |
// top - number of records to return.
// skip - number of records to skip.
func (client UserSubscriptionClient) List(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.sc.Response.Response != nil {
				sc = result.sc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserSubscriptionClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, apimBaseURL, UID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "List", resp, "Failure sending request")
		return
	}

	result.sc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "List", resp, "Failure responding to request")
	}
	if result.sc.hasNextLink() && result.sc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client UserSubscriptionClient) ListPreparer(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"uid": autorest.Encode("path", UID),
	}

	const APIVersion = "2017-03-01"
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
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/users/{uid}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UserSubscriptionClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UserSubscriptionClient) ListResponder(resp *http.Response) (result SubscriptionCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UserSubscriptionClient) listNextResults(ctx context.Context, lastResults SubscriptionCollection) (result SubscriptionCollection, err error) {
	req, err := lastResults.subscriptionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserSubscriptionClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UserSubscriptionClient) ListComplete(ctx context.Context, apimBaseURL string, UID string, filter string, top *int32, skip *int32) (result SubscriptionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSubscriptionClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, apimBaseURL, UID, filter, top, skip)
	return
}
