package logic

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

// IntegrationAccountAssembliesClient is the REST API for Azure Logic Apps.
type IntegrationAccountAssembliesClient struct {
	BaseClient
}

// NewIntegrationAccountAssembliesClient creates an instance of the IntegrationAccountAssembliesClient client.
func NewIntegrationAccountAssembliesClient(subscriptionID string) IntegrationAccountAssembliesClient {
	return NewIntegrationAccountAssembliesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationAccountAssembliesClientWithBaseURI creates an instance of the IntegrationAccountAssembliesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewIntegrationAccountAssembliesClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountAssembliesClient {
	return IntegrationAccountAssembliesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update an assembly for an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// assemblyArtifactName - the assembly artifact name.
// assemblyArtifact - the assembly artifact.
func (client IntegrationAccountAssembliesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact AssemblyDefinition) (result AssemblyDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountAssembliesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: assemblyArtifact,
			Constraints: []validation.Constraint{{Target: "assemblyArtifact.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "assemblyArtifact.Properties.AssemblyName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("logic.IntegrationAccountAssembliesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, assemblyArtifact)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountAssembliesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact AssemblyDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assemblyArtifactName":   autorest.Encode("path", assemblyArtifactName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", pathParameters),
		autorest.WithJSON(assemblyArtifact),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAssembliesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAssembliesClient) CreateOrUpdateResponder(resp *http.Response) (result AssemblyDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an assembly for an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// assemblyArtifactName - the assembly artifact name.
func (client IntegrationAccountAssembliesClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountAssembliesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationAccountAssembliesClient) DeletePreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assemblyArtifactName":   autorest.Encode("path", assemblyArtifactName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAssembliesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAssembliesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an assembly for an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// assemblyArtifactName - the assembly artifact name.
func (client IntegrationAccountAssembliesClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (result AssemblyDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountAssembliesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationAccountAssembliesClient) GetPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assemblyArtifactName":   autorest.Encode("path", assemblyArtifactName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAssembliesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAssembliesClient) GetResponder(resp *http.Response) (result AssemblyDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list the assemblies for an integration account.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
func (client IntegrationAccountAssembliesClient) List(ctx context.Context, resourceGroupName string, integrationAccountName string) (result AssemblyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountAssembliesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, integrationAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationAccountAssembliesClient) ListPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAssembliesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAssembliesClient) ListResponder(resp *http.Response) (result AssemblyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListContentCallbackURL get the content callback url for an integration account assembly.
// Parameters:
// resourceGroupName - the resource group name.
// integrationAccountName - the integration account name.
// assemblyArtifactName - the assembly artifact name.
func (client IntegrationAccountAssembliesClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (result WorkflowTriggerCallbackURL, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationAccountAssembliesClient.ListContentCallbackURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListContentCallbackURLPreparer(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "ListContentCallbackURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListContentCallbackURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "ListContentCallbackURL", resp, "Failure sending request")
		return
	}

	result, err = client.ListContentCallbackURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAssembliesClient", "ListContentCallbackURL", resp, "Failure responding to request")
	}

	return
}

// ListContentCallbackURLPreparer prepares the ListContentCallbackURL request.
func (client IntegrationAccountAssembliesClient) ListContentCallbackURLPreparer(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assemblyArtifactName":   autorest.Encode("path", assemblyArtifactName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}/listContentCallbackUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListContentCallbackURLSender sends the ListContentCallbackURL request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAssembliesClient) ListContentCallbackURLSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListContentCallbackURLResponder handles the response to the ListContentCallbackURL request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAssembliesClient) ListContentCallbackURLResponder(resp *http.Response) (result WorkflowTriggerCallbackURL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
