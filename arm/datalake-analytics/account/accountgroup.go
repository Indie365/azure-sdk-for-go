package account

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// GroupClient is the creates an Azure Data Lake Analytics account management
// client.
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient(subscriptionID string) GroupClient {
	return NewGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupClientWithBaseURI creates an instance of the GroupClient client.
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return GroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates the specified Data Lake Analytics account. This supplies the
// user with computation services for Data Lake Analytics workloads This method
// may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account.the account will be associated with. accountName
// is the name of the Data Lake Analytics account to create. parameters is
// parameters supplied to the create Data Lake Analytics account operation.
func (client GroupClient) Create(resourceGroupName string, accountName string, parameters DataLakeAnalyticsAccount, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.DefaultDataLakeStoreAccount", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxDegreeOfParallelism", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxDegreeOfParallelism", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.QueryStoreRetention", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.QueryStoreRetention", Name: validation.InclusiveMaximum, Rule: 180, Chain: nil},
							{Target: "parameters.DataLakeAnalyticsAccountProperties.QueryStoreRetention", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
						}},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxJobCount", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.DataLakeAnalyticsAccountProperties.MaxJobCount", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
					{Target: "parameters.DataLakeAnalyticsAccountProperties.DataLakeStoreAccounts", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.GroupClient", "Create")
	}

	req, err := client.CreatePreparer(resourceGroupName, accountName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client GroupClient) CreatePreparer(resourceGroupName string, accountName string, parameters DataLakeAnalyticsAccount, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete begins the delete delete process for the Data Lake Analytics account
// object specified by the account name. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to delete
func (client GroupClient) Delete(resourceGroupName string, accountName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, accountName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GroupClient) DeletePreparer(resourceGroupName string, accountName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details of the specified Data Lake Analytics account.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to retrieve.
func (client GroupClient) Get(resourceGroupName string, accountName string) (result DataLakeAnalyticsAccount, err error) {
	req, err := client.GetPreparer(resourceGroupName, accountName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupClient) GetPreparer(resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupClient) GetResponder(resp *http.Response) (result DataLakeAnalyticsAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the first page of Data Lake Analytics accounts, if any, within the
// current subscription. This includes a link to the next page, if any.
//
// filter is oData filter. Optional. top is the number of items to return.
// Optional. skip is the number of items to skip over before returning
// elements. Optional. selectParameter is oData Select statement. Limits the
// properties on each entry to just those requested, e.g.
// Categories?$select=CategoryName,Description. Optional. orderby is orderBy
// clause. One or more comma-separated expressions with an optional "asc" (the
// default) or "desc" depending on the order you'd like the values sorted, e.g.
// Categories?$orderby=CategoryName desc. Optional. count is the Boolean value
// of true or false to request a count of the matching resources included with
// the resources in the response, e.g. Categories?$count=true. Optional.
func (client GroupClient) List(filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result DataLakeAnalyticsAccountListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.GroupClient", "List")
	}

	req, err := client.ListPreparer(filter, top, skip, selectParameter, orderby, count)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupClient) ListPreparer(filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupClient) ListResponder(resp *http.Response) (result DataLakeAnalyticsAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client GroupClient) ListNextResults(lastResults DataLakeAnalyticsAccountListResult) (result DataLakeAnalyticsAccountListResult, err error) {
	req, err := lastResults.DataLakeAnalyticsAccountListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroup gets the first page of Data Lake Analytics accounts, if
// any, within a specific resource group. This includes a link to the next
// page, if any.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. filter is oData filter. Optional. top is the
// number of items to return. Optional. skip is the number of items to skip
// over before returning elements. Optional. selectParameter is oData Select
// statement. Limits the properties on each entry to just those requested, e.g.
// Categories?$select=CategoryName,Description. Optional. orderby is orderBy
// clause. One or more comma-separated expressions with an optional "asc" (the
// default) or "desc" depending on the order you'd like the values sorted, e.g.
// Categories?$orderby=CategoryName desc. Optional. count is the Boolean value
// of true or false to request a count of the matching resources included with
// the resources in the response, e.g. Categories?$count=true. Optional.
func (client GroupClient) ListByResourceGroup(resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result DataLakeAnalyticsAccountListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.GroupClient", "ListByResourceGroup")
	}

	req, err := client.ListByResourceGroupPreparer(resourceGroupName, filter, top, skip, selectParameter, orderby, count)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client GroupClient) ListByResourceGroupPreparer(resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client GroupClient) ListByResourceGroupResponder(resp *http.Response) (result DataLakeAnalyticsAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client GroupClient) ListByResourceGroupNextResults(lastResults DataLakeAnalyticsAccountListResult) (result DataLakeAnalyticsAccountListResult, err error) {
	req, err := lastResults.DataLakeAnalyticsAccountListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// Update updates the Data Lake Analytics account object specified by the
// accountName with the contents of the account object. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to update. parameters is parameters supplied to the update
// Data Lake Analytics account operation.
func (client GroupClient) Update(resourceGroupName string, accountName string, parameters *DataLakeAnalyticsAccountUpdateParameters, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, accountName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.GroupClient", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.GroupClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client GroupClient) UpdatePreparer(resourceGroupName string, accountName string, parameters *DataLakeAnalyticsAccountUpdateParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client GroupClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}
