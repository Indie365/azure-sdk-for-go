package containerregistry

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

// ConnectedRegistriesClient is the client for the ConnectedRegistries methods of the Containerregistry service.
type ConnectedRegistriesClient struct {
	BaseClient
}

// NewConnectedRegistriesClient creates an instance of the ConnectedRegistriesClient client.
func NewConnectedRegistriesClient(subscriptionID string) ConnectedRegistriesClient {
	return NewConnectedRegistriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConnectedRegistriesClientWithBaseURI creates an instance of the ConnectedRegistriesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewConnectedRegistriesClientWithBaseURI(baseURI string, subscriptionID string) ConnectedRegistriesClient {
	return ConnectedRegistriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a connected registry for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// connectedRegistryName - the name of the connected registry.
// connectedRegistryCreateParameters - the parameters for creating a connectedRegistry.
func (client ConnectedRegistriesClient) Create(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry) (result ConnectedRegistriesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: connectedRegistryName,
			Constraints: []validation.Constraint{{Target: "connectedRegistryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: connectedRegistryCreateParameters,
			Constraints: []validation.Constraint{{Target: "connectedRegistryCreateParameters.ConnectedRegistryProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "connectedRegistryCreateParameters.ConnectedRegistryProperties.Parent", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "connectedRegistryCreateParameters.ConnectedRegistryProperties.Parent.SyncProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "connectedRegistryCreateParameters.ConnectedRegistryProperties.Parent.SyncProperties.TokenID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "connectedRegistryCreateParameters.ConnectedRegistryProperties.Parent.SyncProperties.MessageTTL", Name: validation.Null, Rule: true, Chain: nil},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("containerregistry.ConnectedRegistriesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ConnectedRegistriesClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectedRegistryName": autorest.Encode("path", connectedRegistryName),
		"registryName":          autorest.Encode("path", registryName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}", pathParameters),
		autorest.WithJSON(connectedRegistryCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectedRegistriesClient) CreateSender(req *http.Request) (future ConnectedRegistriesCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ConnectedRegistriesClient) CreateResponder(resp *http.Response) (result ConnectedRegistry, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deactivate deactivates the connected registry instance.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// connectedRegistryName - the name of the connected registry.
func (client ConnectedRegistriesClient) Deactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string) (result ConnectedRegistriesDeactivateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.Deactivate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: connectedRegistryName,
			Constraints: []validation.Constraint{{Target: "connectedRegistryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ConnectedRegistriesClient", "Deactivate", err.Error())
	}

	req, err := client.DeactivatePreparer(ctx, resourceGroupName, registryName, connectedRegistryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Deactivate", nil, "Failure preparing request")
		return
	}

	result, err = client.DeactivateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Deactivate", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeactivatePreparer prepares the Deactivate request.
func (client ConnectedRegistriesClient) DeactivatePreparer(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectedRegistryName": autorest.Encode("path", connectedRegistryName),
		"registryName":          autorest.Encode("path", registryName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}/deactivate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeactivateSender sends the Deactivate request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectedRegistriesClient) DeactivateSender(req *http.Request) (future ConnectedRegistriesDeactivateFuture, err error) {
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

// DeactivateResponder handles the response to the Deactivate request. The method always
// closes the http.Response Body.
func (client ConnectedRegistriesClient) DeactivateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes a connected registry from a container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// connectedRegistryName - the name of the connected registry.
func (client ConnectedRegistriesClient) Delete(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string) (result ConnectedRegistriesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: connectedRegistryName,
			Constraints: []validation.Constraint{{Target: "connectedRegistryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ConnectedRegistriesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, connectedRegistryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConnectedRegistriesClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectedRegistryName": autorest.Encode("path", connectedRegistryName),
		"registryName":          autorest.Encode("path", registryName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectedRegistriesClient) DeleteSender(req *http.Request) (future ConnectedRegistriesDeleteFuture, err error) {
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
func (client ConnectedRegistriesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the connected registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// connectedRegistryName - the name of the connected registry.
func (client ConnectedRegistriesClient) Get(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string) (result ConnectedRegistry, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: connectedRegistryName,
			Constraints: []validation.Constraint{{Target: "connectedRegistryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ConnectedRegistriesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, connectedRegistryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConnectedRegistriesClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectedRegistryName": autorest.Encode("path", connectedRegistryName),
		"registryName":          autorest.Encode("path", registryName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectedRegistriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConnectedRegistriesClient) GetResponder(resp *http.Response) (result ConnectedRegistry, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all connected registries for the specified container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// filter - an OData filter expression that describes a subset of connectedRegistries to return. The parameters
// that can be filtered are parent.id (the resource id of the connectedRegistry parent), mode, and
// connectionState. The supported operator is eq.
func (client ConnectedRegistriesClient) List(ctx context.Context, resourceGroupName string, registryName string, filter string) (result ConnectedRegistryListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.List")
		defer func() {
			sc := -1
			if result.crlr.Response.Response != nil {
				sc = result.crlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ConnectedRegistriesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.crlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "List", resp, "Failure sending request")
		return
	}

	result.crlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.crlr.hasNextLink() && result.crlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ConnectedRegistriesClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectedRegistriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConnectedRegistriesClient) ListResponder(resp *http.Response) (result ConnectedRegistryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ConnectedRegistriesClient) listNextResults(ctx context.Context, lastResults ConnectedRegistryListResult) (result ConnectedRegistryListResult, err error) {
	req, err := lastResults.connectedRegistryListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConnectedRegistriesClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string, filter string) (result ConnectedRegistryListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, registryName, filter)
	return
}

// Update updates a connected registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// connectedRegistryName - the name of the connected registry.
// connectedRegistryUpdateParameters - the parameters for updating a connectedRegistry.
func (client ConnectedRegistriesClient) Update(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters) (result ConnectedRegistriesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedRegistriesClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: connectedRegistryName,
			Constraints: []validation.Constraint{{Target: "connectedRegistryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "connectedRegistryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ConnectedRegistriesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ConnectedRegistriesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ConnectedRegistriesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectedRegistryName": autorest.Encode("path", connectedRegistryName),
		"registryName":          autorest.Encode("path", registryName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}", pathParameters),
		autorest.WithJSON(connectedRegistryUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectedRegistriesClient) UpdateSender(req *http.Request) (future ConnectedRegistriesUpdateFuture, err error) {
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
func (client ConnectedRegistriesClient) UpdateResponder(resp *http.Response) (result ConnectedRegistry, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
