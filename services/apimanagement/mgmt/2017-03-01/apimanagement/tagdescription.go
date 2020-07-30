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

// TagDescriptionClient is the apiManagement Client
type TagDescriptionClient struct {
	BaseClient
}

// NewTagDescriptionClient creates an instance of the TagDescriptionClient client.
func NewTagDescriptionClient(subscriptionID string) TagDescriptionClient {
	return NewTagDescriptionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTagDescriptionClientWithBaseURI creates an instance of the TagDescriptionClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTagDescriptionClientWithBaseURI(baseURI string, subscriptionID string) TagDescriptionClient {
	return TagDescriptionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create/Update tag description in scope of the Api.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// tagID - tag identifier. Must be unique in the current API Management service instance.
// parameters - create parameters.
// ifMatch - the entity state (Etag) version of the Tag to update. A value of "*" can be used for If-Match to
// unconditionally apply the operation.
func (client TagDescriptionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string, parameters TagDescriptionCreateParameters, ifMatch string) (result TagDescriptionContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagDescriptionClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: tagID,
			Constraints: []validation.Constraint{{Target: "tagID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "tagID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "tagID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.TagDescriptionBaseProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.TagDescriptionBaseProperties.ExternalDocsURL", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.TagDescriptionBaseProperties.ExternalDocsURL", Name: validation.MaxLength, Rule: 2000, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("apimanagement.TagDescriptionClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, apiid, tagID, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client TagDescriptionClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string, parameters TagDescriptionCreateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tagId":             autorest.Encode("path", tagID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client TagDescriptionClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TagDescriptionClient) CreateOrUpdateResponder(resp *http.Response) (result TagDescriptionContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete tag description for the Api.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// tagID - tag identifier. Must be unique in the current API Management service instance.
// ifMatch - the entity state (Etag) version of the Api schema to update. A value of "*" can be used for
// If-Match to unconditionally apply the operation.
func (client TagDescriptionClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagDescriptionClient.Delete")
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
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: tagID,
			Constraints: []validation.Constraint{{Target: "tagID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "tagID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "tagID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TagDescriptionClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, apiid, tagID, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TagDescriptionClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tagId":             autorest.Encode("path", tagID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TagDescriptionClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TagDescriptionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get tag associated with the API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// tagID - tag identifier. Must be unique in the current API Management service instance.
func (client TagDescriptionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string) (result TagDescriptionContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagDescriptionClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: tagID,
			Constraints: []validation.Constraint{{Target: "tagID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "tagID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "tagID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TagDescriptionClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, apiid, tagID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TagDescriptionClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tagId":             autorest.Encode("path", tagID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TagDescriptionClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TagDescriptionClient) GetResponder(resp *http.Response) (result TagDescriptionContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityState gets the entity state version of the tag specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// tagID - tag identifier. Must be unique in the current API Management service instance.
func (client TagDescriptionClient) GetEntityState(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagDescriptionClient.GetEntityState")
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
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: tagID,
			Constraints: []validation.Constraint{{Target: "tagID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "tagID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "tagID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TagDescriptionClient", "GetEntityState", err.Error())
	}

	req, err := client.GetEntityStatePreparer(ctx, resourceGroupName, serviceName, apiid, tagID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "GetEntityState", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityStateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "GetEntityState", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "GetEntityState", resp, "Failure responding to request")
	}

	return
}

// GetEntityStatePreparer prepares the GetEntityState request.
func (client TagDescriptionClient) GetEntityStatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, tagID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tagId":             autorest.Encode("path", tagID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityStateSender sends the GetEntityState request. The method will close the
// http.Response Body if it receives an error.
func (client TagDescriptionClient) GetEntityStateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEntityStateResponder handles the response to the GetEntityState request. The method always
// closes the http.Response Body.
func (client TagDescriptionClient) GetEntityStateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByAPI lists all Tags descriptions in scope of API. Model similar to swagger - tagDescription is defined on API
// level but tag may be assigned to the Operations
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// filter - | Field       | Supported operators    | Supported functions                         |
// |-------------|------------------------|---------------------------------------------|
// | id          | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name        | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client TagDescriptionClient) ListByAPI(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result TagDescriptionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagDescriptionClient.ListByAPI")
		defer func() {
			sc := -1
			if result.tdc.Response.Response != nil {
				sc = result.tdc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.TagDescriptionClient", "ListByAPI", err.Error())
	}

	result.fn = client.listByAPINextResults
	req, err := client.ListByAPIPreparer(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "ListByAPI", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAPISender(req)
	if err != nil {
		result.tdc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "ListByAPI", resp, "Failure sending request")
		return
	}

	result.tdc, err = client.ListByAPIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "ListByAPI", resp, "Failure responding to request")
	}
	if result.tdc.hasNextLink() && result.tdc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByAPIPreparer prepares the ListByAPI request.
func (client TagDescriptionClient) ListByAPIPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
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
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAPISender sends the ListByAPI request. The method will close the
// http.Response Body if it receives an error.
func (client TagDescriptionClient) ListByAPISender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAPIResponder handles the response to the ListByAPI request. The method always
// closes the http.Response Body.
func (client TagDescriptionClient) ListByAPIResponder(resp *http.Response) (result TagDescriptionCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAPINextResults retrieves the next set of results, if any.
func (client TagDescriptionClient) listByAPINextResults(ctx context.Context, lastResults TagDescriptionCollection) (result TagDescriptionCollection, err error) {
	req, err := lastResults.tagDescriptionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "listByAPINextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAPISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "listByAPINextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAPIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TagDescriptionClient", "listByAPINextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAPIComplete enumerates all values, automatically crossing page boundaries as required.
func (client TagDescriptionClient) ListByAPIComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result TagDescriptionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagDescriptionClient.ListByAPI")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAPI(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
	return
}
