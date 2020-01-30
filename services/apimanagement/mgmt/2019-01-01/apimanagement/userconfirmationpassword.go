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

// UserConfirmationPasswordClient is the apiManagement Client
type UserConfirmationPasswordClient struct {
	BaseClient
}

// NewUserConfirmationPasswordClient creates an instance of the UserConfirmationPasswordClient client.
func NewUserConfirmationPasswordClient(subscriptionID string) UserConfirmationPasswordClient {
	return NewUserConfirmationPasswordClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserConfirmationPasswordClientWithBaseURI creates an instance of the UserConfirmationPasswordClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewUserConfirmationPasswordClientWithBaseURI(baseURI string, subscriptionID string) UserConfirmationPasswordClient {
	return UserConfirmationPasswordClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Send sends confirmation
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client UserConfirmationPasswordClient) Send(ctx context.Context, resourceGroupName string, serviceName string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserConfirmationPasswordClient.Send")
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
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserConfirmationPasswordClient", "Send", err.Error())
	}

	req, err := client.SendPreparer(ctx, resourceGroupName, serviceName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserConfirmationPasswordClient", "Send", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.UserConfirmationPasswordClient", "Send", resp, "Failure sending request")
		return
	}

	result, err = client.SendResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserConfirmationPasswordClient", "Send", resp, "Failure responding to request")
	}

	return
}

// SendPreparer prepares the Send request.
func (client UserConfirmationPasswordClient) SendPreparer(ctx context.Context, resourceGroupName string, serviceName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/confirmations/password/send", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SendSender sends the Send request. The method will close the
// http.Response Body if it receives an error.
func (client UserConfirmationPasswordClient) SendSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// SendResponder handles the response to the Send request. The method always
// closes the http.Response Body.
func (client UserConfirmationPasswordClient) SendResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
