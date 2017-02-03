package graphrbac

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ApplicationsClient is the the Graph RBAC Management Client
type ApplicationsClient struct {
	ManagementClient
}

// NewApplicationsClient creates an instance of the ApplicationsClient client.
func NewApplicationsClient(tenantID string) ApplicationsClient {
	return NewApplicationsClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewApplicationsClientWithBaseURI creates an instance of the
// ApplicationsClient client.
func NewApplicationsClientWithBaseURI(baseURI string, tenantID string) ApplicationsClient {
	return ApplicationsClient{NewWithBaseURI(baseURI, tenantID)}
}

// Create create a new application.
//
// parameters is the parameters for creating an application.
func (client ApplicationsClient) Create(parameters ApplicationCreateParameters) (result Application, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AvailableToOtherTenants", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.IdentifierUris", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "graphrbac.ApplicationsClient", "Create")
	}

	req, err := client.CreatePreparer(parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationsClient) CreatePreparer(parameters ApplicationCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) CreateResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an application.
//
// applicationObjectID is application object ID.
func (client ApplicationsClient) Delete(applicationObjectID string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(applicationObjectID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationsClient) DeletePreparer(applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an application by object ID.
//
// applicationObjectID is application object ID.
func (client ApplicationsClient) Get(applicationObjectID string) (result Application, err error) {
	req, err := client.GetPreparer(applicationObjectID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationsClient) GetPreparer(applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) GetResponder(resp *http.Response) (result Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists applications by filter parameters.
//
// filter is the filters to apply to the operation.
func (client ApplicationsClient) List(filter string) (result ApplicationListResult, err error) {
	req, err := client.ListPreparer(filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationsClient) ListPreparer(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListResponder(resp *http.Response) (result ApplicationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListKeyCredentials get the keyCredentials associated with an application.
//
// applicationObjectID is application object ID.
func (client ApplicationsClient) ListKeyCredentials(applicationObjectID string) (result KeyCredentialListResult, err error) {
	req, err := client.ListKeyCredentialsPreparer(applicationObjectID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListKeyCredentials", nil, "Failure preparing request")
	}

	resp, err := client.ListKeyCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListKeyCredentials", resp, "Failure sending request")
	}

	result, err = client.ListKeyCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListKeyCredentials", resp, "Failure responding to request")
	}

	return
}

// ListKeyCredentialsPreparer prepares the ListKeyCredentials request.
func (client ApplicationsClient) ListKeyCredentialsPreparer(applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/keyCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListKeyCredentialsSender sends the ListKeyCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListKeyCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListKeyCredentialsResponder handles the response to the ListKeyCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListKeyCredentialsResponder(resp *http.Response) (result KeyCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNext gets a list of applications from the current tenant.
//
// nextLink is next link for the list operation.
func (client ApplicationsClient) ListNext(nextLink string) (result ApplicationListResult, err error) {
	req, err := client.ListNextPreparer(nextLink)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListNext", nil, "Failure preparing request")
	}

	resp, err := client.ListNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListNext", resp, "Failure sending request")
	}

	result, err = client.ListNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListNext", resp, "Failure responding to request")
	}

	return
}

// ListNextPreparer prepares the ListNext request.
func (client ApplicationsClient) ListNextPreparer(nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink": nextLink,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/{nextLink}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListNextSender sends the ListNext request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListNextSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListNextResponder handles the response to the ListNext request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListNextResponder(resp *http.Response) (result ApplicationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPasswordCredentials get the passwordCredentials associated with an
// application.
//
// applicationObjectID is application object ID.
func (client ApplicationsClient) ListPasswordCredentials(applicationObjectID string) (result PasswordCredentialListResult, err error) {
	req, err := client.ListPasswordCredentialsPreparer(applicationObjectID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", nil, "Failure preparing request")
	}

	resp, err := client.ListPasswordCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", resp, "Failure sending request")
	}

	result, err = client.ListPasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// ListPasswordCredentialsPreparer prepares the ListPasswordCredentials request.
func (client ApplicationsClient) ListPasswordCredentialsPreparer(applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/passwordCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListPasswordCredentialsSender sends the ListPasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) ListPasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListPasswordCredentialsResponder handles the response to the ListPasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) ListPasswordCredentialsResponder(resp *http.Response) (result PasswordCredentialListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch update an existing application.
//
// applicationObjectID is application object ID. parameters is parameters to
// update an existing application.
func (client ApplicationsClient) Patch(applicationObjectID string, parameters ApplicationUpdateParameters) (result autorest.Response, err error) {
	req, err := client.PatchPreparer(applicationObjectID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", nil, "Failure preparing request")
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", resp, "Failure sending request")
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client ApplicationsClient) PatchPreparer(applicationObjectID string, parameters ApplicationUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) PatchResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdateKeyCredentials update the keyCredentials associated with an
// application.
//
// applicationObjectID is application object ID. parameters is parameters to
// update the keyCredentials of an existing application.
func (client ApplicationsClient) UpdateKeyCredentials(applicationObjectID string, parameters KeyCredentialsUpdateParameters) (result autorest.Response, err error) {
	req, err := client.UpdateKeyCredentialsPreparer(applicationObjectID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdateKeyCredentials", nil, "Failure preparing request")
	}

	resp, err := client.UpdateKeyCredentialsSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdateKeyCredentials", resp, "Failure sending request")
	}

	result, err = client.UpdateKeyCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdateKeyCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdateKeyCredentialsPreparer prepares the UpdateKeyCredentials request.
func (client ApplicationsClient) UpdateKeyCredentialsPreparer(applicationObjectID string, parameters KeyCredentialsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/keyCredentials", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateKeyCredentialsSender sends the UpdateKeyCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) UpdateKeyCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateKeyCredentialsResponder handles the response to the UpdateKeyCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) UpdateKeyCredentialsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdatePasswordCredentials update passwordCredentials associated with an
// application.
//
// applicationObjectID is application object ID. parameters is parameters to
// update passwordCredentials of an existing application.
func (client ApplicationsClient) UpdatePasswordCredentials(applicationObjectID string, parameters PasswordCredentialsUpdateParameters) (result autorest.Response, err error) {
	req, err := client.UpdatePasswordCredentialsPreparer(applicationObjectID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", nil, "Failure preparing request")
	}

	resp, err := client.UpdatePasswordCredentialsSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure sending request")
	}

	result, err = client.UpdatePasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdatePasswordCredentialsPreparer prepares the UpdatePasswordCredentials request.
func (client ApplicationsClient) UpdatePasswordCredentialsPreparer(applicationObjectID string, parameters PasswordCredentialsUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": applicationObjectID,
		"tenantID":            autorest.Encode("path", client.TenantID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/applications/{applicationObjectId}/passwordCredentials", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdatePasswordCredentialsSender sends the UpdatePasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) UpdatePasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdatePasswordCredentialsResponder handles the response to the UpdatePasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) UpdatePasswordCredentialsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
