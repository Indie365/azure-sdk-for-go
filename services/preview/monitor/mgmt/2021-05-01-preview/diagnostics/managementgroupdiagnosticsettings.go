package diagnostics

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

// ManagementGroupDiagnosticSettingsClient is the monitor Management Client
type ManagementGroupDiagnosticSettingsClient struct {
	BaseClient
}

// NewManagementGroupDiagnosticSettingsClient creates an instance of the ManagementGroupDiagnosticSettingsClient
// client.
func NewManagementGroupDiagnosticSettingsClient(subscriptionID string) ManagementGroupDiagnosticSettingsClient {
	return NewManagementGroupDiagnosticSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagementGroupDiagnosticSettingsClientWithBaseURI creates an instance of the
// ManagementGroupDiagnosticSettingsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewManagementGroupDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) ManagementGroupDiagnosticSettingsClient {
	return ManagementGroupDiagnosticSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates management group diagnostic settings for the specified resource.
// Parameters:
// managementGroupID - the management group id.
// parameters - parameters supplied to the operation.
// name - the name of the diagnostic setting.
func (client ManagementGroupDiagnosticSettingsClient) CreateOrUpdate(ctx context.Context, managementGroupID string, parameters ManagementGroupDiagnosticSettingsResource, name string) (result ManagementGroupDiagnosticSettingsResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupDiagnosticSettingsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, managementGroupID, parameters, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagementGroupDiagnosticSettingsClient) CreateOrUpdatePreparer(ctx context.Context, managementGroupID string, parameters ManagementGroupDiagnosticSettingsResource, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": managementGroupID,
		"name":              autorest.Encode("path", name),
	}

	const APIVersion = "2021-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/diagnosticSettings/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupDiagnosticSettingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagementGroupDiagnosticSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagementGroupDiagnosticSettingsResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes existing management group diagnostic settings for the specified resource.
// Parameters:
// managementGroupID - the management group id.
// name - the name of the diagnostic setting.
func (client ManagementGroupDiagnosticSettingsClient) Delete(ctx context.Context, managementGroupID string, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupDiagnosticSettingsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, managementGroupID, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagementGroupDiagnosticSettingsClient) DeletePreparer(ctx context.Context, managementGroupID string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": managementGroupID,
		"name":              autorest.Encode("path", name),
	}

	const APIVersion = "2021-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/diagnosticSettings/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupDiagnosticSettingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagementGroupDiagnosticSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the active management group diagnostic settings for the specified resource.
// Parameters:
// managementGroupID - the management group id.
// name - the name of the diagnostic setting.
func (client ManagementGroupDiagnosticSettingsClient) Get(ctx context.Context, managementGroupID string, name string) (result ManagementGroupDiagnosticSettingsResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupDiagnosticSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, managementGroupID, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagementGroupDiagnosticSettingsClient) GetPreparer(ctx context.Context, managementGroupID string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": managementGroupID,
		"name":              autorest.Encode("path", name),
	}

	const APIVersion = "2021-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/diagnosticSettings/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupDiagnosticSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementGroupDiagnosticSettingsClient) GetResponder(resp *http.Response) (result ManagementGroupDiagnosticSettingsResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the active management group diagnostic settings list for the specified management group.
// Parameters:
// managementGroupID - the management group id.
func (client ManagementGroupDiagnosticSettingsClient) List(ctx context.Context, managementGroupID string) (result ManagementGroupDiagnosticSettingsResourceCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupDiagnosticSettingsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ManagementGroupDiagnosticSettingsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagementGroupDiagnosticSettingsClient) ListPreparer(ctx context.Context, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": managementGroupID,
	}

	const APIVersion = "2021-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/diagnosticSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupDiagnosticSettingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagementGroupDiagnosticSettingsClient) ListResponder(resp *http.Response) (result ManagementGroupDiagnosticSettingsResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
