package virtualmachineimagebuilder

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

// VirtualMachineImageTemplatesClient is the azure Virtual Machine Image Builder Client
type VirtualMachineImageTemplatesClient struct {
	BaseClient
}

// NewVirtualMachineImageTemplatesClient creates an instance of the VirtualMachineImageTemplatesClient client.
func NewVirtualMachineImageTemplatesClient(subscriptionID string) VirtualMachineImageTemplatesClient {
	return NewVirtualMachineImageTemplatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineImageTemplatesClientWithBaseURI creates an instance of the VirtualMachineImageTemplatesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewVirtualMachineImageTemplatesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineImageTemplatesClient {
	return VirtualMachineImageTemplatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a virtual machine image template
// Parameters:
// parameters - parameters supplied to the CreateImageTemplate operation
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
func (client VirtualMachineImageTemplatesClient) CreateOrUpdate(ctx context.Context, parameters ImageTemplate, resourceGroupName string, imageTemplateName string) (result VirtualMachineImageTemplatesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ImageTemplateProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ImageTemplateProperties.Source", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ImageTemplateProperties.Distribute", Name: validation.Null, Rule: true, Chain: nil},
				}}}},
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, imageTemplateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineImageTemplatesClient) CreateOrUpdatePreparer(ctx context.Context, parameters ImageTemplate, resourceGroupName string, imageTemplateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) CreateOrUpdateSender(req *http.Request) (future VirtualMachineImageTemplatesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) CreateOrUpdateResponder(resp *http.Response) (result ImageTemplate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a virtual machine image template
// Parameters:
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
func (client VirtualMachineImageTemplatesClient) Delete(ctx context.Context, resourceGroupName string, imageTemplateName string) (result VirtualMachineImageTemplatesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, imageTemplateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineImageTemplatesClient) DeletePreparer(ctx context.Context, resourceGroupName string, imageTemplateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) DeleteSender(req *http.Request) (future VirtualMachineImageTemplatesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get information about a virtual machine image template
// Parameters:
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
func (client VirtualMachineImageTemplatesClient) Get(ctx context.Context, resourceGroupName string, imageTemplateName string) (result ImageTemplate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, imageTemplateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineImageTemplatesClient) GetPreparer(ctx context.Context, resourceGroupName string, imageTemplateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) GetResponder(resp *http.Response) (result ImageTemplate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRunOutput get the specified run output for the specified image template resource
// Parameters:
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
// runOutputName - the name of the run output
func (client VirtualMachineImageTemplatesClient) GetRunOutput(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string) (result RunOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.GetRunOutput")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}},
		{TargetValue: runOutputName,
			Constraints: []validation.Constraint{{Target: "runOutputName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "GetRunOutput", err.Error())
	}

	req, err := client.GetRunOutputPreparer(ctx, resourceGroupName, imageTemplateName, runOutputName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "GetRunOutput", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRunOutputSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "GetRunOutput", resp, "Failure sending request")
		return
	}

	result, err = client.GetRunOutputResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "GetRunOutput", resp, "Failure responding to request")
	}

	return
}

// GetRunOutputPreparer prepares the GetRunOutput request.
func (client VirtualMachineImageTemplatesClient) GetRunOutputPreparer(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runOutputName":     autorest.Encode("path", runOutputName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/runOutputs/{runOutputName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRunOutputSender sends the GetRunOutput request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) GetRunOutputSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetRunOutputResponder handles the response to the GetRunOutput request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) GetRunOutputResponder(resp *http.Response) (result RunOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets information about the VM image templates associated with the subscription.
func (client VirtualMachineImageTemplatesClient) List(ctx context.Context) (result ImageTemplateListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.List")
		defer func() {
			sc := -1
			if result.itlr.Response.Response != nil {
				sc = result.itlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "List", resp, "Failure sending request")
		return
	}

	result.itlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineImageTemplatesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VirtualMachineImages/imageTemplates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) ListResponder(resp *http.Response) (result ImageTemplateListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualMachineImageTemplatesClient) listNextResults(ctx context.Context, lastResults ImageTemplateListResult) (result ImageTemplateListResult, err error) {
	req, err := lastResults.imageTemplateListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineImageTemplatesClient) ListComplete(ctx context.Context) (result ImageTemplateListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets information about the VM image templates associated with the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client VirtualMachineImageTemplatesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ImageTemplateListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.itlr.Response.Response != nil {
				sc = result.itlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.itlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.itlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client VirtualMachineImageTemplatesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) ListByResourceGroupResponder(resp *http.Response) (result ImageTemplateListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client VirtualMachineImageTemplatesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ImageTemplateListResult) (result ImageTemplateListResult, err error) {
	req, err := lastResults.imageTemplateListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineImageTemplatesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ImageTemplateListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListRunOutputs list all run outputs for the specified Image Template resource
// Parameters:
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
func (client VirtualMachineImageTemplatesClient) ListRunOutputs(ctx context.Context, resourceGroupName string, imageTemplateName string) (result RunOutputCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.ListRunOutputs")
		defer func() {
			sc := -1
			if result.roc.Response.Response != nil {
				sc = result.roc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListRunOutputs", err.Error())
	}

	result.fn = client.listRunOutputsNextResults
	req, err := client.ListRunOutputsPreparer(ctx, resourceGroupName, imageTemplateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListRunOutputs", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRunOutputsSender(req)
	if err != nil {
		result.roc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListRunOutputs", resp, "Failure sending request")
		return
	}

	result.roc, err = client.ListRunOutputsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "ListRunOutputs", resp, "Failure responding to request")
	}

	return
}

// ListRunOutputsPreparer prepares the ListRunOutputs request.
func (client VirtualMachineImageTemplatesClient) ListRunOutputsPreparer(ctx context.Context, resourceGroupName string, imageTemplateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/runOutputs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRunOutputsSender sends the ListRunOutputs request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) ListRunOutputsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListRunOutputsResponder handles the response to the ListRunOutputs request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) ListRunOutputsResponder(resp *http.Response) (result RunOutputCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRunOutputsNextResults retrieves the next set of results, if any.
func (client VirtualMachineImageTemplatesClient) listRunOutputsNextResults(ctx context.Context, lastResults RunOutputCollection) (result RunOutputCollection, err error) {
	req, err := lastResults.runOutputCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listRunOutputsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRunOutputsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listRunOutputsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRunOutputsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "listRunOutputsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRunOutputsComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineImageTemplatesClient) ListRunOutputsComplete(ctx context.Context, resourceGroupName string, imageTemplateName string) (result RunOutputCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.ListRunOutputs")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRunOutputs(ctx, resourceGroupName, imageTemplateName)
	return
}

// Run create artifacts from a existing image template
// Parameters:
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
func (client VirtualMachineImageTemplatesClient) Run(ctx context.Context, resourceGroupName string, imageTemplateName string) (result VirtualMachineImageTemplatesRunFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.Run")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Run", err.Error())
	}

	req, err := client.RunPreparer(ctx, resourceGroupName, imageTemplateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Run", nil, "Failure preparing request")
		return
	}

	result, err = client.RunSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Run", result.Response(), "Failure sending request")
		return
	}

	return
}

// RunPreparer prepares the Run request.
func (client VirtualMachineImageTemplatesClient) RunPreparer(ctx context.Context, resourceGroupName string, imageTemplateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) RunSender(req *http.Request) (future VirtualMachineImageTemplatesRunFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update the tags for this Virtual Machine Image Template
// Parameters:
// parameters - additional parameters for Image Template update.
// resourceGroupName - the name of the resource group.
// imageTemplateName - the name of the image Template
func (client VirtualMachineImageTemplatesClient) Update(ctx context.Context, parameters ImageTemplateUpdateParameters, resourceGroupName string, imageTemplateName string) (result ImageTemplate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImageTemplatesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageTemplateName,
			Constraints: []validation.Constraint{{Target: "imageTemplateName", Name: validation.Pattern, Rule: `^[A-Za-z0-9-_]{1,64}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, parameters, resourceGroupName, imageTemplateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimagebuilder.VirtualMachineImageTemplatesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualMachineImageTemplatesClient) UpdatePreparer(ctx context.Context, parameters ImageTemplateUpdateParameters, resourceGroupName string, imageTemplateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageTemplateName": autorest.Encode("path", imageTemplateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImageTemplatesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualMachineImageTemplatesClient) UpdateResponder(resp *http.Response) (result ImageTemplate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
