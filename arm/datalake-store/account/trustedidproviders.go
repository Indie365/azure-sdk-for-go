package account

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

// TrustedIDProvidersClient is the creates an Azure Data Lake Store account
// management client.
type TrustedIDProvidersClient struct {
	ManagementClient
}

// NewTrustedIDProvidersClient creates an instance of the
// TrustedIDProvidersClient client.
func NewTrustedIDProvidersClient(subscriptionID string) TrustedIDProvidersClient {
	return NewTrustedIDProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTrustedIDProvidersClientWithBaseURI creates an instance of the
// TrustedIDProvidersClient client.
func NewTrustedIDProvidersClientWithBaseURI(baseURI string, subscriptionID string) TrustedIDProvidersClient {
	return TrustedIDProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified trusted identity provider.
// During update, the trusted identity provider with the specified name will be
// replaced with this new provider
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Store account. accountName is the name of the Data Lake Store
// account to which to add the trusted identity provider. trustedIDProviderName
// is the name of the trusted identity provider. This is used for
// differentiation of providers in the account. parameters is parameters
// supplied to create the create the trusted identity provider.
func (client TrustedIDProvidersClient) CreateOrUpdate(resourceGroupName string, accountName string, trustedIDProviderName string, parameters TrustedIDProvider) (result TrustedIDProvider, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.TrustedIDProviderProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.TrustedIDProviderProperties.IDProvider", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.TrustedIDProvidersClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, accountName, trustedIDProviderName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client TrustedIDProvidersClient) CreateOrUpdatePreparer(resourceGroupName string, accountName string, trustedIDProviderName string, parameters TrustedIDProvider) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"trustedIdProviderName": autorest.Encode("path", trustedIDProviderName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedIDProvidersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TrustedIDProvidersClient) CreateOrUpdateResponder(resp *http.Response) (result TrustedIDProvider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified trusted identity provider from the specified
// Data Lake Store account
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Store account. accountName is the name of the Data Lake Store
// account from which to delete the trusted identity provider.
// trustedIDProviderName is the name of the trusted identity provider to
// delete.
func (client TrustedIDProvidersClient) Delete(resourceGroupName string, accountName string, trustedIDProviderName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, accountName, trustedIDProviderName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TrustedIDProvidersClient) DeletePreparer(resourceGroupName string, accountName string, trustedIDProviderName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"trustedIdProviderName": autorest.Encode("path", trustedIDProviderName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedIDProvidersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TrustedIDProvidersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Data Lake Store trusted identity provider.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Store account. accountName is the name of the Data Lake Store
// account from which to get the trusted identity provider.
// trustedIDProviderName is the name of the trusted identity provider to
// retrieve.
func (client TrustedIDProvidersClient) Get(resourceGroupName string, accountName string, trustedIDProviderName string) (result TrustedIDProvider, err error) {
	req, err := client.GetPreparer(resourceGroupName, accountName, trustedIDProviderName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TrustedIDProvidersClient) GetPreparer(resourceGroupName string, accountName string, trustedIDProviderName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"trustedIdProviderName": autorest.Encode("path", trustedIDProviderName),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedIDProvidersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TrustedIDProvidersClient) GetResponder(resp *http.Response) (result TrustedIDProvider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccount lists the Data Lake Store trusted identity providers within
// the specified Data Lake Store account.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Store account. accountName is the name of the Data Lake Store
// account from which to get the trusted identity providers.
func (client TrustedIDProvidersClient) ListByAccount(resourceGroupName string, accountName string) (result DataLakeStoreTrustedIDProviderListResult, err error) {
	req, err := client.ListByAccountPreparer(resourceGroupName, accountName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "ListByAccount", nil, "Failure preparing request")
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "ListByAccount", resp, "Failure sending request")
	}

	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "ListByAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAccountPreparer prepares the ListByAccount request.
func (client TrustedIDProvidersClient) ListByAccountPreparer(resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByAccountSender sends the ListByAccount request. The method will close the
// http.Response Body if it receives an error.
func (client TrustedIDProvidersClient) ListByAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByAccountResponder handles the response to the ListByAccount request. The method always
// closes the http.Response Body.
func (client TrustedIDProvidersClient) ListByAccountResponder(resp *http.Response) (result DataLakeStoreTrustedIDProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccountNextResults retrieves the next set of results, if any.
func (client TrustedIDProvidersClient) ListByAccountNextResults(lastResults DataLakeStoreTrustedIDProviderListResult) (result DataLakeStoreTrustedIDProviderListResult, err error) {
	req, err := lastResults.DataLakeStoreTrustedIDProviderListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "ListByAccount", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "ListByAccount", resp, "Failure sending next results request")
	}

	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.TrustedIDProvidersClient", "ListByAccount", resp, "Failure responding to next results request")
	}

	return
}
