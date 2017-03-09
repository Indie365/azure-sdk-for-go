package sql

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
	"net/http"
)

// ServersClient is the provides create, read, update and delete functionality
// for Azure SQL Database resources including servers, databases, elastic
// pools, recommendations, operations, and usage metrics.
type ServersClient struct {
	ManagementClient
}

// NewServersClient creates an instance of the ServersClient client.
func NewServersClient(subscriptionID string) ServersClient {
	return NewServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServersClientWithBaseURI creates an instance of the ServersClient client.
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return ServersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a new server.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. parameters is the required
// parameters for creating or updating a server.
func (client ServersClient) CreateOrUpdate(resourceGroupName string, serverName string, parameters Server) (result Server, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServersClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, parameters Server) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServersClient) CreateOrUpdateResponder(resp *http.Response) (result Server, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a SQL server.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client ServersClient) Delete(resourceGroupName string, serverName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, serverName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServersClient) DeletePreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetByResourceGroup gets a server.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client ServersClient) GetByResourceGroup(resourceGroupName string, serverName string) (result Server, err error) {
	req, err := client.GetByResourceGroupPreparer(resourceGroupName, serverName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "GetByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.GetByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "GetByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.GetByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "GetByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// GetByResourceGroupPreparer prepares the GetByResourceGroup request.
func (client ServersClient) GetByResourceGroupPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByResourceGroupSender sends the GetByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) GetByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetByResourceGroupResponder handles the response to the GetByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServersClient) GetByResourceGroupResponder(resp *http.Response) (result Server, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetServiceObjective gets a database service objective.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server. serviceObjectiveName is
// the name of the service objective to retrieve.
func (client ServersClient) GetServiceObjective(resourceGroupName string, serverName string, serviceObjectiveName string) (result ServiceObjective, err error) {
	req, err := client.GetServiceObjectivePreparer(resourceGroupName, serverName, serviceObjectiveName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "GetServiceObjective", nil, "Failure preparing request")
	}

	resp, err := client.GetServiceObjectiveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "GetServiceObjective", resp, "Failure sending request")
	}

	result, err = client.GetServiceObjectiveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "GetServiceObjective", resp, "Failure responding to request")
	}

	return
}

// GetServiceObjectivePreparer prepares the GetServiceObjective request.
func (client ServersClient) GetServiceObjectivePreparer(resourceGroupName string, serverName string, serviceObjectiveName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"serverName":           autorest.Encode("path", serverName),
		"serviceObjectiveName": autorest.Encode("path", serviceObjectiveName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/serviceObjectives/{serviceObjectiveName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetServiceObjectiveSender sends the GetServiceObjective request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) GetServiceObjectiveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetServiceObjectiveResponder handles the response to the GetServiceObjective request. The method always
// closes the http.Response Body.
func (client ServersClient) GetServiceObjectiveResponder(resp *http.Response) (result ServiceObjective, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns a list of servers.
func (client ServersClient) List() (result ServerListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServersClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/servers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServersClient) ListResponder(resp *http.Response) (result ServerListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns a server.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal.
func (client ServersClient) ListByResourceGroup(resourceGroupName string) (result ServerListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ServersClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServersClient) ListByResourceGroupResponder(resp *http.Response) (result ServerListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListServiceObjectives returns database service objectives.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client ServersClient) ListServiceObjectives(resourceGroupName string, serverName string) (result ServiceObjectiveListResult, err error) {
	req, err := client.ListServiceObjectivesPreparer(resourceGroupName, serverName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "ListServiceObjectives", nil, "Failure preparing request")
	}

	resp, err := client.ListServiceObjectivesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "ListServiceObjectives", resp, "Failure sending request")
	}

	result, err = client.ListServiceObjectivesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "ListServiceObjectives", resp, "Failure responding to request")
	}

	return
}

// ListServiceObjectivesPreparer prepares the ListServiceObjectives request.
func (client ServersClient) ListServiceObjectivesPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/serviceObjectives", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListServiceObjectivesSender sends the ListServiceObjectives request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) ListServiceObjectivesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListServiceObjectivesResponder handles the response to the ListServiceObjectives request. The method always
// closes the http.Response Body.
func (client ServersClient) ListServiceObjectivesResponder(resp *http.Response) (result ServiceObjectiveListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUsages returns server usages.
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client ServersClient) ListUsages(resourceGroupName string, serverName string) (result ServerMetricListResult, err error) {
	req, err := client.ListUsagesPreparer(resourceGroupName, serverName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "ListUsages", nil, "Failure preparing request")
	}

	resp, err := client.ListUsagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServersClient", "ListUsages", resp, "Failure sending request")
	}

	result, err = client.ListUsagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServersClient", "ListUsages", resp, "Failure responding to request")
	}

	return
}

// ListUsagesPreparer prepares the ListUsages request.
func (client ServersClient) ListUsagesPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListUsagesSender sends the ListUsages request. The method will close the
// http.Response Body if it receives an error.
func (client ServersClient) ListUsagesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListUsagesResponder handles the response to the ListUsages request. The method always
// closes the http.Response Body.
func (client ServersClient) ListUsagesResponder(resp *http.Response) (result ServerMetricListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
