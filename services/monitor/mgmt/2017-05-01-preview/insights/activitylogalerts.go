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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ActivityLogAlertsClient is the monitor Management Client
type ActivityLogAlertsClient struct {
	ManagementClient
}

// NewActivityLogAlertsClient creates an instance of the ActivityLogAlertsClient client.
func NewActivityLogAlertsClient(subscriptionID string) ActivityLogAlertsClient {
	return NewActivityLogAlertsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActivityLogAlertsClientWithBaseURI creates an instance of the ActivityLogAlertsClient client.
func NewActivityLogAlertsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogAlertsClient {
	return ActivityLogAlertsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new activity log alert or update an existing one.
//
// resourceGroupName is the name of the resource group. activityLogAlertName is the name of the activity log alert.
// activityLogAlert is the activity log alert to create or use for the update.
func (client ActivityLogAlertsClient) CreateOrUpdate(resourceGroupName string, activityLogAlertName string, activityLogAlert ActivityLogAlertResource) (result ActivityLogAlertResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: activityLogAlert,
			Constraints: []validation.Constraint{{Target: "activityLogAlert.ActivityLogAlert", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "activityLogAlert.ActivityLogAlert.Scopes", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "activityLogAlert.ActivityLogAlert.Condition", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "activityLogAlert.ActivityLogAlert.Condition.AllOf", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "activityLogAlert.ActivityLogAlert.Actions", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "insights.ActivityLogAlertsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, activityLogAlertName, activityLogAlert)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ActivityLogAlertsClient) CreateOrUpdatePreparer(resourceGroupName string, activityLogAlertName string, activityLogAlert ActivityLogAlertResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"activityLogAlertName": autorest.Encode("path", activityLogAlertName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}", pathParameters),
		autorest.WithJSON(activityLogAlert),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityLogAlertsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ActivityLogAlertsClient) CreateOrUpdateResponder(resp *http.Response) (result ActivityLogAlertResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an activity log alert.
//
// resourceGroupName is the name of the resource group. activityLogAlertName is the name of the activity log alert.
func (client ActivityLogAlertsClient) Delete(resourceGroupName string, activityLogAlertName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, activityLogAlertName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ActivityLogAlertsClient) DeletePreparer(resourceGroupName string, activityLogAlertName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"activityLogAlertName": autorest.Encode("path", activityLogAlertName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityLogAlertsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ActivityLogAlertsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an activity log alert.
//
// resourceGroupName is the name of the resource group. activityLogAlertName is the name of the activity log alert.
func (client ActivityLogAlertsClient) Get(resourceGroupName string, activityLogAlertName string) (result ActivityLogAlertResource, err error) {
	req, err := client.GetPreparer(resourceGroupName, activityLogAlertName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ActivityLogAlertsClient) GetPreparer(resourceGroupName string, activityLogAlertName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"activityLogAlertName": autorest.Encode("path", activityLogAlertName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityLogAlertsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ActivityLogAlertsClient) GetResponder(resp *http.Response) (result ActivityLogAlertResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup get a list of all activity log alerts in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ActivityLogAlertsClient) ListByResourceGroup(resourceGroupName string) (result ActivityLogAlertList, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ActivityLogAlertsClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityLogAlertsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ActivityLogAlertsClient) ListByResourceGroupResponder(resp *http.Response) (result ActivityLogAlertList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionID get a list of all activity log alerts in a subscription.
func (client ActivityLogAlertsClient) ListBySubscriptionID() (result ActivityLogAlertList, err error) {
	req, err := client.ListBySubscriptionIDPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "ListBySubscriptionID", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "ListBySubscriptionID", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "ListBySubscriptionID", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionIDPreparer prepares the ListBySubscriptionID request.
func (client ActivityLogAlertsClient) ListBySubscriptionIDPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/activityLogAlerts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionIDSender sends the ListBySubscriptionID request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityLogAlertsClient) ListBySubscriptionIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionIDResponder handles the response to the ListBySubscriptionID request. The method always
// closes the http.Response Body.
func (client ActivityLogAlertsClient) ListBySubscriptionIDResponder(resp *http.Response) (result ActivityLogAlertList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing ActivityLogAlertResource's tags. To update other fields use the CreateOrUpdate method.
//
// resourceGroupName is the name of the resource group. activityLogAlertName is the name of the activity log alert.
// activityLogAlertPatch is parameters supplied to the operation.
func (client ActivityLogAlertsClient) Update(resourceGroupName string, activityLogAlertName string, activityLogAlertPatch ActivityLogAlertPatchBody) (result ActivityLogAlertResource, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, activityLogAlertName, activityLogAlertPatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.ActivityLogAlertsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ActivityLogAlertsClient) UpdatePreparer(resourceGroupName string, activityLogAlertName string, activityLogAlertPatch ActivityLogAlertPatchBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"activityLogAlertName": autorest.Encode("path", activityLogAlertName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/activityLogAlerts/{activityLogAlertName}", pathParameters),
		autorest.WithJSON(activityLogAlertPatch),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ActivityLogAlertsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ActivityLogAlertsClient) UpdateResponder(resp *http.Response) (result ActivityLogAlertResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
