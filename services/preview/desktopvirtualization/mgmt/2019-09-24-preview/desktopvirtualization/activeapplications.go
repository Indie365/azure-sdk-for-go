package desktopvirtualization

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

// ActiveApplicationsClient is the client for the ActiveApplications methods of the Desktopvirtualization service.
type ActiveApplicationsClient struct {
	BaseClient
}

// NewActiveApplicationsClient creates an instance of the ActiveApplicationsClient client.
func NewActiveApplicationsClient(subscriptionID string) ActiveApplicationsClient {
	return NewActiveApplicationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActiveApplicationsClientWithBaseURI creates an instance of the ActiveApplicationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewActiveApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ActiveApplicationsClient {
	return ActiveApplicationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListBySessionHost list applications for the given session host.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// filter - oData filter expression. Valid properties for filtering are userprincipalname and sessionstate.
func (client ActiveApplicationsClient) ListBySessionHost(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, filter string) (result ApplicationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActiveApplicationsClient.ListBySessionHost")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.ActiveApplicationsClient", "ListBySessionHost", err.Error())
	}

	result.fn = client.listBySessionHostNextResults
	req, err := client.ListBySessionHostPreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ActiveApplicationsClient", "ListBySessionHost", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySessionHostSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ActiveApplicationsClient", "ListBySessionHost", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListBySessionHostResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ActiveApplicationsClient", "ListBySessionHost", resp, "Failure responding to request")
	}

	return
}

// ListBySessionHostPreparer prepares the ListBySessionHost request.
func (client ActiveApplicationsClient) ListBySessionHostPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}/activeApplications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySessionHostSender sends the ListBySessionHost request. The method will close the
// http.Response Body if it receives an error.
func (client ActiveApplicationsClient) ListBySessionHostSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySessionHostResponder handles the response to the ListBySessionHost request. The method always
// closes the http.Response Body.
func (client ActiveApplicationsClient) ListBySessionHostResponder(resp *http.Response) (result ApplicationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySessionHostNextResults retrieves the next set of results, if any.
func (client ActiveApplicationsClient) listBySessionHostNextResults(ctx context.Context, lastResults ApplicationList) (result ApplicationList, err error) {
	req, err := lastResults.applicationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.ActiveApplicationsClient", "listBySessionHostNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySessionHostSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.ActiveApplicationsClient", "listBySessionHostNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySessionHostResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.ActiveApplicationsClient", "listBySessionHostNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySessionHostComplete enumerates all values, automatically crossing page boundaries as required.
func (client ActiveApplicationsClient) ListBySessionHostComplete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, filter string) (result ApplicationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActiveApplicationsClient.ListBySessionHost")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySessionHost(ctx, resourceGroupName, hostPoolName, sessionHostName, filter)
	return
}
