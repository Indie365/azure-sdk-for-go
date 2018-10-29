package graphrbac

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
	"net/http"
)

// OAuth2Client is the the Graph RBAC Management Client
type OAuth2Client struct {
	BaseClient
}

// NewOAuth2Client creates an instance of the OAuth2Client client.
func NewOAuth2Client(tenantID string) OAuth2Client {
	return NewOAuth2ClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewOAuth2ClientWithBaseURI creates an instance of the OAuth2Client client.
func NewOAuth2ClientWithBaseURI(baseURI string, tenantID string) OAuth2Client {
	return OAuth2Client{NewWithBaseURI(baseURI, tenantID)}
}

// Get queries OAuth2 permissions for the relevant SP ObjectId of an app.
// Parameters:
// filter - this is the Service Principal ObjectId associated with the app
func (client OAuth2Client) Get(ctx context.Context, filter string) (result Permissions, err error) {
	req, err := client.GetPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2Client", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OAuth2Client) GetPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/oauth2PermissionGrants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OAuth2Client) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OAuth2Client) GetResponder(resp *http.Response) (result Permissions, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Grant grants OAuth2 permissions for the relevant resource Ids of an app.
// Parameters:
// body - the relevant app Service Principal Object Id and the Service Principal Objecit Id you want to grant.
func (client OAuth2Client) Grant(ctx context.Context, body *Permissions) (result Permissions, err error) {
	req, err := client.GrantPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2Client", "Grant", nil, "Failure preparing request")
		return
	}

	resp, err := client.GrantSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2Client", "Grant", resp, "Failure sending request")
		return
	}

	result, err = client.GrantResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.OAuth2Client", "Grant", resp, "Failure responding to request")
	}

	return
}

// GrantPreparer prepares the Grant request.
func (client OAuth2Client) GrantPreparer(ctx context.Context, body *Permissions) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/oauth2PermissionGrants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GrantSender sends the Grant request. The method will close the
// http.Response Body if it receives an error.
func (client OAuth2Client) GrantSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GrantResponder handles the response to the Grant request. The method always
// closes the http.Response Body.
func (client OAuth2Client) GrantResponder(resp *http.Response) (result Permissions, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
