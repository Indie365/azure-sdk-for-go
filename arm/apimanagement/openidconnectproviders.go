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

// OpenIDConnectProvidersClient is the use these REST APIs for performing
// operations on entities like API, Product, and Subscription associated with
// your Azure API Management deployment.
type OpenIDConnectProvidersClient struct {
	ManagementClient
}

// NewOpenIDConnectProvidersClient creates an instance of the
// OpenIDConnectProvidersClient client.
func NewOpenIDConnectProvidersClient(subscriptionID string) OpenIDConnectProvidersClient {
	return NewOpenIDConnectProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOpenIDConnectProvidersClientWithBaseURI creates an instance of the
// OpenIDConnectProvidersClient client.
func NewOpenIDConnectProvidersClientWithBaseURI(baseURI string, subscriptionID string) OpenIDConnectProvidersClient {
	return OpenIDConnectProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the OpenID Connect Provider.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. opid is identifier of the OpenID Connect
// Provider. parameters is create parameters.
func (client OpenIDConnectProvidersClient) CreateOrUpdate(resourceGroupName string, serviceName string, opid string, parameters OpenidConnectProviderCreateContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: opid,
			Constraints: []validation.Constraint{{Target: "opid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "opid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Name", Name: validation.MaxLength, Rule: 50, Chain: nil}}},
				{Target: "parameters.MetadataEndpoint", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ClientID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.OpenIDConnectProvidersClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serviceName, opid, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client OpenIDConnectProvidersClient) CreateOrUpdatePreparer(resourceGroupName string, serviceName string, opid string, parameters OpenidConnectProviderCreateContract) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"opid":              autorest.Encode("path", opid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client OpenIDConnectProvidersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client OpenIDConnectProvidersClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes specific OpenID Connect Provider of the API Management
// service instance.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. opid is identifier of the OpenID Connect
// Provider. ifMatch is the entity state (Etag) version of the OpenID Connect
// Provider to delete. A value of "*" can be used for If-Match to
// unconditionally apply the operation.
func (client OpenIDConnectProvidersClient) Delete(resourceGroupName string, serviceName string, opid string, ifMatch string) (result ErrorBodyContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.OpenIDConnectProvidersClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, serviceName, opid, ifMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OpenIDConnectProvidersClient) DeletePreparer(resourceGroupName string, serviceName string, opid string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"opid":              autorest.Encode("path", opid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OpenIDConnectProvidersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client OpenIDConnectProvidersClient) DeleteResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets specific OpenID Connect Provider.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. opid is identifier of the OpenID Connect
// Provider.
func (client OpenIDConnectProvidersClient) Get(resourceGroupName string, serviceName string, opid string) (result OpenidConnectProviderContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.OpenIDConnectProvidersClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, serviceName, opid)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OpenIDConnectProvidersClient) GetPreparer(resourceGroupName string, serviceName string, opid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"opid":              autorest.Encode("path", opid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OpenIDConnectProvidersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OpenIDConnectProvidersClient) GetResponder(resp *http.Response) (result OpenidConnectProviderContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists all OpenID Connect Providers.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. filter is | Field | Supported operators    |
// Supported functions                         |
// |-------|------------------------|---------------------------------------------|
// | id    | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith | top is number of records to return. skip is number of records to
// skip.
func (client OpenIDConnectProvidersClient) ListByService(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result OpenIDConnectProviderCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService")
	}

	req, err := client.ListByServicePreparer(resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService", nil, "Failure preparing request")
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService", resp, "Failure sending request")
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client OpenIDConnectProvidersClient) ListByServicePreparer(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
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

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client OpenIDConnectProvidersClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client OpenIDConnectProvidersClient) ListByServiceResponder(resp *http.Response) (result OpenIDConnectProviderCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServiceNextResults retrieves the next set of results, if any.
func (client OpenIDConnectProvidersClient) ListByServiceNextResults(lastResults OpenIDConnectProviderCollection) (result OpenIDConnectProviderCollection, err error) {
	req, err := lastResults.OpenIDConnectProviderCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService", resp, "Failure sending next results request")
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "ListByService", resp, "Failure responding to next results request")
	}

	return
}

// Update updates the specific OpenID Connect Provider.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. opid is identifier of the OpenID Connect
// Provider. parameters is update parameters. ifMatch is the entity state
// (Etag) version of the OpenID Connect Provider to update. A value of "*" can
// be used for If-Match to unconditionally apply the operation.
func (client OpenIDConnectProvidersClient) Update(resourceGroupName string, serviceName string, opid string, parameters OpenidConnectProviderUpdateContract, ifMatch string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.OpenIDConnectProvidersClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, serviceName, opid, parameters, ifMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OpenIDConnectProvidersClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client OpenIDConnectProvidersClient) UpdatePreparer(resourceGroupName string, serviceName string, opid string, parameters OpenidConnectProviderUpdateContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"opid":              autorest.Encode("path", opid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OpenIDConnectProvidersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OpenIDConnectProvidersClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
