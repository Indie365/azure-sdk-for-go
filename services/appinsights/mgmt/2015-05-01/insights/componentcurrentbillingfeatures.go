package insights

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

// ComponentCurrentBillingFeaturesClient is the composite Swagger for Application Insights Management Client
type ComponentCurrentBillingFeaturesClient struct {
	BaseClient
}

// NewComponentCurrentBillingFeaturesClient creates an instance of the ComponentCurrentBillingFeaturesClient client.
func NewComponentCurrentBillingFeaturesClient(subscriptionID string, purgeID string) ComponentCurrentBillingFeaturesClient {
	return NewComponentCurrentBillingFeaturesClientWithBaseURI(DefaultBaseURI, subscriptionID, purgeID)
}

// NewComponentCurrentBillingFeaturesClientWithBaseURI creates an instance of the ComponentCurrentBillingFeaturesClient
// client.
func NewComponentCurrentBillingFeaturesClientWithBaseURI(baseURI string, subscriptionID string, purgeID string) ComponentCurrentBillingFeaturesClient {
	return ComponentCurrentBillingFeaturesClient{NewWithBaseURI(baseURI, subscriptionID, purgeID)}
}

// Get returns current billing features for an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource.
func (client ComponentCurrentBillingFeaturesClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result ApplicationInsightsComponentBillingFeatures, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ComponentCurrentBillingFeaturesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ComponentCurrentBillingFeaturesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ComponentCurrentBillingFeaturesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ComponentCurrentBillingFeaturesClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ComponentCurrentBillingFeaturesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ComponentCurrentBillingFeaturesClient) GetResponder(resp *http.Response) (result ApplicationInsightsComponentBillingFeatures, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update current billing features for an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. billingFeaturesProperties is properties that need to be specified to update billing features
// for an Application Insights component.
func (client ComponentCurrentBillingFeaturesClient) Update(ctx context.Context, resourceGroupName string, resourceName string, billingFeaturesProperties ApplicationInsightsComponentBillingFeatures) (result ApplicationInsightsComponentBillingFeatures, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, billingFeaturesProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ComponentCurrentBillingFeaturesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ComponentCurrentBillingFeaturesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ComponentCurrentBillingFeaturesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ComponentCurrentBillingFeaturesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, billingFeaturesProperties ApplicationInsightsComponentBillingFeatures) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", pathParameters),
		autorest.WithJSON(billingFeaturesProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ComponentCurrentBillingFeaturesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ComponentCurrentBillingFeaturesClient) UpdateResponder(resp *http.Response) (result ApplicationInsightsComponentBillingFeatures, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
