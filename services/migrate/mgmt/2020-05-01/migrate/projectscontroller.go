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

// ProjectsControllerClient is the client for the ProjectsController methods of the Migrate service.
type ProjectsControllerClient struct {
	BaseClient
}

// NewProjectsControllerClient creates an instance of the ProjectsControllerClient client.
func NewProjectsControllerClient(subscriptionID string) ProjectsControllerClient {
	return NewProjectsControllerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProjectsControllerClientWithBaseURI creates an instance of the ProjectsControllerClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewProjectsControllerClientWithBaseURI(baseURI string, subscriptionID string) ProjectsControllerClient {
	return ProjectsControllerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeleteMigrateProject delete the migrate project. It deletes summary of the project.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// APIVersion - the API version to use for this operation.
func (client ProjectsControllerClient) DeleteMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsControllerClient.DeleteMigrateProject")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "DeleteMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteMigrateProjectSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "DeleteMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "DeleteMigrateProject", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteMigrateProjectPreparer prepares the DeleteMigrateProject request.
func (client ProjectsControllerClient) DeleteMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteMigrateProjectSender sends the DeleteMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsControllerClient) DeleteMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteMigrateProjectResponder handles the response to the DeleteMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsControllerClient) DeleteMigrateProjectResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetMigrateProject get information related to a specific migrate project. Returns a json object of type
// 'migrateProject' as specified in the models section.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// APIVersion - the API version to use for this operation.
func (client ProjectsControllerClient) GetMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, APIVersion string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsControllerClient.GetMigrateProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "GetMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMigrateProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "GetMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.GetMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "GetMigrateProject", resp, "Failure responding to request")
		return
	}

	return
}

// GetMigrateProjectPreparer prepares the GetMigrateProject request.
func (client ProjectsControllerClient) GetMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMigrateProjectSender sends the GetMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsControllerClient) GetMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMigrateProjectResponder handles the response to the GetMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsControllerClient) GetMigrateProjectResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PatchMigrateProject update a project with specified name. Supports partial updates, for example only tags can be
// provided.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// body - migrate project body.
// APIVersion - the API version to use for this operation.
func (client ProjectsControllerClient) PatchMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project, APIVersion string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsControllerClient.PatchMigrateProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PatchMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "PatchMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchMigrateProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "PatchMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.PatchMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "PatchMigrateProject", resp, "Failure responding to request")
		return
	}

	return
}

// PatchMigrateProjectPreparer prepares the PatchMigrateProject request.
func (client ProjectsControllerClient) PatchMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	body.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchMigrateProjectSender sends the PatchMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsControllerClient) PatchMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchMigrateProjectResponder handles the response to the PatchMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsControllerClient) PatchMigrateProjectResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutMigrateProject create a new project by sending a json object of type 'migrateproject' as given in Models section
// as part of the Request Body. The project name is unique.
//
// This operation is Idempotent.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that project is part of.
// migrateProjectName - migrate project name.
// body - migrate project body.
// APIVersion - the API version to use for this operation.
func (client ProjectsControllerClient) PutMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project, APIVersion string) (result Project, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProjectsControllerClient.PutMigrateProject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutMigrateProjectPreparer(ctx, resourceGroupName, migrateProjectName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "PutMigrateProject", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMigrateProjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "PutMigrateProject", resp, "Failure sending request")
		return
	}

	result, err = client.PutMigrateProjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.ProjectsControllerClient", "PutMigrateProject", resp, "Failure responding to request")
		return
	}

	return
}

// PutMigrateProjectPreparer prepares the PutMigrateProject request.
func (client ProjectsControllerClient) PutMigrateProjectPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, body Project, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	body.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMigrateProjectSender sends the PutMigrateProject request. The method will close the
// http.Response Body if it receives an error.
func (client ProjectsControllerClient) PutMigrateProjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutMigrateProjectResponder handles the response to the PutMigrateProject request. The method always
// closes the http.Response Body.
func (client ProjectsControllerClient) PutMigrateProjectResponder(resp *http.Response) (result Project, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
