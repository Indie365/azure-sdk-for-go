package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// AppsClient is the REST API for Azure Spring Cloud
type AppsClient struct {
	BaseClient
}

// NewAppsClient creates an instance of the AppsClient client.
func NewAppsClient(subscriptionID string) AppsClient {
	return NewAppsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAppsClientWithBaseURI creates an instance of the AppsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return AppsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new App or update an exiting App.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// appResource - parameters for the create or update operation
func (client AppsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource) (result AppsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: appResource,
			Constraints: []validation.Constraint{{Target: "appResource.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "appResource.Properties.TemporaryDisk", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "appResource.Properties.TemporaryDisk.SizeInGB", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "appResource.Properties.TemporaryDisk.SizeInGB", Name: validation.InclusiveMaximum, Rule: int64(5), Chain: nil},
							{Target: "appResource.Properties.TemporaryDisk.SizeInGB", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
						}},
					}},
					{Target: "appResource.Properties.PersistentDisk", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "appResource.Properties.PersistentDisk.SizeInGB", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "appResource.Properties.PersistentDisk.SizeInGB", Name: validation.InclusiveMaximum, Rule: int64(50), Chain: nil},
								{Target: "appResource.Properties.PersistentDisk.SizeInGB", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
							}},
							{Target: "appResource.Properties.PersistentDisk.UsedInGB", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "appResource.Properties.PersistentDisk.UsedInGB", Name: validation.InclusiveMaximum, Rule: int64(50), Chain: nil},
									{Target: "appResource.Properties.PersistentDisk.UsedInGB", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("appplatform.AppsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, appName, appResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AppsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", pathParameters),
		autorest.WithJSON(appResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) CreateOrUpdateSender(req *http.Request) (future AppsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AppsClient) CreateOrUpdateResponder(resp *http.Response) (result AppResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete operation to delete an App.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
func (client AppsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result AppsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, appName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AppsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) DeleteSender(req *http.Request) (future AppsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AppsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an App and its properties.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// syncStatus - indicates whether sync status
func (client AppsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, syncStatus string) (result AppResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, appName, syncStatus)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AppsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, syncStatus string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(syncStatus) > 0 {
		queryParameters["syncStatus"] = autorest.Encode("query", syncStatus)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AppsClient) GetResponder(resp *http.Response) (result AppResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetResourceUploadURL get an resource upload URL for an App, which may be artifacts or source archive.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
func (client AppsClient) GetResourceUploadURL(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result ResourceUploadDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.GetResourceUploadURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetResourceUploadURLPreparer(ctx, resourceGroupName, serviceName, appName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "GetResourceUploadURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceUploadURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "GetResourceUploadURL", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceUploadURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "GetResourceUploadURL", resp, "Failure responding to request")
		return
	}

	return
}

// GetResourceUploadURLPreparer prepares the GetResourceUploadURL request.
func (client AppsClient) GetResourceUploadURLPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/getResourceUploadUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceUploadURLSender sends the GetResourceUploadURL request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) GetResourceUploadURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResourceUploadURLResponder handles the response to the GetResourceUploadURL request. The method always
// closes the http.Response Body.
func (client AppsClient) GetResourceUploadURLResponder(resp *http.Response) (result ResourceUploadDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List handles requests to list all resources in a Service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
func (client AppsClient) List(ctx context.Context, resourceGroupName string, serviceName string) (result AppResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.List")
		defer func() {
			sc := -1
			if result.arc.Response.Response != nil {
				sc = result.arc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.arc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "List", resp, "Failure sending request")
		return
	}

	result.arc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.arc.hasNextLink() && result.arc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AppsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AppsClient) ListResponder(resp *http.Response) (result AppResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AppsClient) listNextResults(ctx context.Context, lastResults AppResourceCollection) (result AppResourceCollection, err error) {
	req, err := lastResults.appResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.AppsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.AppsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AppsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result AppResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, serviceName)
	return
}

// Update operation to update an exiting App.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// appResource - parameters for the update operation
func (client AppsClient) Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource) (result AppsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, appName, appResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AppsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", pathParameters),
		autorest.WithJSON(appResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) UpdateSender(req *http.Request) (future AppsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AppsClient) UpdateResponder(resp *http.Response) (result AppResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateDomain check the resource name is valid as well as not in use.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// appName - the name of the App resource.
// validatePayload - custom domain payload to be validated
func (client AppsClient) ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, appName string, validatePayload CustomDomainValidatePayload) (result CustomDomainValidateResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AppsClient.ValidateDomain")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: validatePayload,
			Constraints: []validation.Constraint{{Target: "validatePayload.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appplatform.AppsClient", "ValidateDomain", err.Error())
	}

	req, err := client.ValidateDomainPreparer(ctx, resourceGroupName, serviceName, appName, validatePayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "ValidateDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "ValidateDomain", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppsClient", "ValidateDomain", resp, "Failure responding to request")
		return
	}

	return
}

// ValidateDomainPreparer prepares the ValidateDomain request.
func (client AppsClient) ValidateDomainPreparer(ctx context.Context, resourceGroupName string, serviceName string, appName string, validatePayload CustomDomainValidatePayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/validateDomain", pathParameters),
		autorest.WithJSON(validatePayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateDomainSender sends the ValidateDomain request. The method will close the
// http.Response Body if it receives an error.
func (client AppsClient) ValidateDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateDomainResponder handles the response to the ValidateDomain request. The method always
// closes the http.Response Body.
func (client AppsClient) ValidateDomainResponder(resp *http.Response) (result CustomDomainValidateResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
