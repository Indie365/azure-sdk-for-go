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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// PolicySnippetsClient is the use these REST APIs for performing operations on
// entities like API, Product, and Subscription associated with your Azure API
// Management deployment.
type PolicySnippetsClient struct {
	ManagementClient
}

// NewPolicySnippetsClient creates an instance of the PolicySnippetsClient
// client.
func NewPolicySnippetsClient(subscriptionID string) PolicySnippetsClient {
	return NewPolicySnippetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPolicySnippetsClientWithBaseURI creates an instance of the
// PolicySnippetsClient client.
func NewPolicySnippetsClientWithBaseURI(baseURI string, subscriptionID string) PolicySnippetsClient {
	return PolicySnippetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByService lists all policy snippets.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. scope is policy scope.
func (client PolicySnippetsClient) ListByService(resourceGroupName string, serviceName string, scope PolicyScopeContract) (result ListPolicySnippetContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.PolicySnippetsClient", "ListByService")
	}

	req, err := client.ListByServicePreparer(resourceGroupName, serviceName, scope)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.PolicySnippetsClient", "ListByService", nil, "Failure preparing request")
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.PolicySnippetsClient", "ListByService", resp, "Failure sending request")
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.PolicySnippetsClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client PolicySnippetsClient) ListByServicePreparer(resourceGroupName string, serviceName string, scope PolicyScopeContract) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(string(scope)) > 0 {
		queryParameters["scope"] = autorest.Encode("query", scope)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policySnippets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client PolicySnippetsClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client PolicySnippetsClient) ListByServiceResponder(resp *http.Response) (result ListPolicySnippetContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
