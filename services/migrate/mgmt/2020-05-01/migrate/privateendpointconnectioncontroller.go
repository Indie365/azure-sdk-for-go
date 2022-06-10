package migrate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateEndpointConnectionControllerClient is the client for the PrivateEndpointConnectionController methods of the
// Migrate service.
type PrivateEndpointConnectionControllerClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionControllerClient creates an instance of the PrivateEndpointConnectionControllerClient
// client.
func NewPrivateEndpointConnectionControllerClient(subscriptionID string) PrivateEndpointConnectionControllerClient {
	return NewPrivateEndpointConnectionControllerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionControllerClientWithBaseURI creates an instance of the
// PrivateEndpointConnectionControllerClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrivateEndpointConnectionControllerClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionControllerClient {
	return PrivateEndpointConnectionControllerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeletePrivateEndpointConnection delete the private endpoint. Deleting non-existent private endpoint is a
// no-operation.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// peConnectionName - private endpoint connection name.
// APIVersion - the API version to use for this operation.
func (client PrivateEndpointConnectionControllerClient) DeletePrivateEndpointConnection(ctx context.Context, resourceGroupName string, migrateProjectName string, peConnectionName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionControllerClient.DeletePrivateEndpointConnection")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePrivateEndpointConnectionPreparer(ctx, resourceGroupName, migrateProjectName, peConnectionName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "DeletePrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeletePrivateEndpointConnectionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "DeletePrivateEndpointConnection", resp, "Failure sending request")
		return
	}

	result, err = client.DeletePrivateEndpointConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "DeletePrivateEndpointConnection", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePrivateEndpointConnectionPreparer prepares the DeletePrivateEndpointConnection request.
func (client PrivateEndpointConnectionControllerClient) DeletePrivateEndpointConnectionPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, peConnectionName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"peConnectionName":   autorest.Encode("path", peConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/privateEndpointConnections/{peConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeletePrivateEndpointConnectionSender sends the DeletePrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionControllerClient) DeletePrivateEndpointConnectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeletePrivateEndpointConnectionResponder handles the response to the DeletePrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionControllerClient) DeletePrivateEndpointConnectionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetPrivateEndpointConnection get the private endpoint with the specified name.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// peConnectionName - private endpoint connection name.
// APIVersion - the API version to use for this operation.
func (client PrivateEndpointConnectionControllerClient) GetPrivateEndpointConnection(ctx context.Context, resourceGroupName string, migrateProjectName string, peConnectionName string, APIVersion string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionControllerClient.GetPrivateEndpointConnection")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPrivateEndpointConnectionPreparer(ctx, resourceGroupName, migrateProjectName, peConnectionName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "GetPrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPrivateEndpointConnectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "GetPrivateEndpointConnection", resp, "Failure sending request")
		return
	}

	result, err = client.GetPrivateEndpointConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "GetPrivateEndpointConnection", resp, "Failure responding to request")
		return
	}

	return
}

// GetPrivateEndpointConnectionPreparer prepares the GetPrivateEndpointConnection request.
func (client PrivateEndpointConnectionControllerClient) GetPrivateEndpointConnectionPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, peConnectionName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"peConnectionName":   autorest.Encode("path", peConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/privateEndpointConnections/{peConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPrivateEndpointConnectionSender sends the GetPrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionControllerClient) GetPrivateEndpointConnectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetPrivateEndpointConnectionResponder handles the response to the GetPrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionControllerClient) GetPrivateEndpointConnectionResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutPrivateEndpointConnection create a private endpoint with specified name. If a private endpoint already exists,
// update it.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// peConnectionName - private endpoint connection name.
// APIVersion - the API version to use for this operation.
// body - modify connection state body.
func (client PrivateEndpointConnectionControllerClient) PutPrivateEndpointConnection(ctx context.Context, resourceGroupName string, migrateProjectName string, peConnectionName string, APIVersion string, body ModifyConnectionStateBody) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionControllerClient.PutPrivateEndpointConnection")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutPrivateEndpointConnectionPreparer(ctx, resourceGroupName, migrateProjectName, peConnectionName, APIVersion, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "PutPrivateEndpointConnection", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutPrivateEndpointConnectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "PutPrivateEndpointConnection", resp, "Failure sending request")
		return
	}

	result, err = client.PutPrivateEndpointConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.PrivateEndpointConnectionControllerClient", "PutPrivateEndpointConnection", resp, "Failure responding to request")
		return
	}

	return
}

// PutPrivateEndpointConnectionPreparer prepares the PutPrivateEndpointConnection request.
func (client PrivateEndpointConnectionControllerClient) PutPrivateEndpointConnectionPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, peConnectionName string, APIVersion string, body ModifyConnectionStateBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"peConnectionName":   autorest.Encode("path", peConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/privateEndpointConnections/{peConnectionName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutPrivateEndpointConnectionSender sends the PutPrivateEndpointConnection request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionControllerClient) PutPrivateEndpointConnectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutPrivateEndpointConnectionResponder handles the response to the PutPrivateEndpointConnection request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionControllerClient) PutPrivateEndpointConnectionResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
