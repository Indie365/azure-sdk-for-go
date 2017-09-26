package recoveryservicesbackup

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ProtectionContainersClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionContainersClient struct {
	ManagementClient
}

// NewProtectionContainersClient creates an instance of the ProtectionContainersClient client.
func NewProtectionContainersClient(subscriptionID string) ProtectionContainersClient {
	return NewProtectionContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionContainersClientWithBaseURI creates an instance of the ProtectionContainersClient client.
func NewProtectionContainersClientWithBaseURI(baseURI string, subscriptionID string) ProtectionContainersClient {
	return ProtectionContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets details of the specific container registered to your Recovery Services Vault.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. fabricName is name of the fabric where the container belongs. containerName is
// name of the container whose details need to be fetched.
func (client ProtectionContainersClient) Get(vaultName string, resourceGroupName string, fabricName string, containerName string) (result ProtectionContainerResource, err error) {
	req, err := client.GetPreparer(vaultName, resourceGroupName, fabricName, containerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectionContainersClient) GetPreparer(vaultName string, resourceGroupName string, fabricName string, containerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionContainersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionContainersClient) GetResponder(resp *http.Response) (result ProtectionContainerResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Refresh discovers all the containers in the subscription that can be backed up to Recovery Services Vault. This is
// an asynchronous operation. To know the status of the operation, call GetRefreshOperationResult API.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. fabricName is fabric name associated the container.
func (client ProtectionContainersClient) Refresh(vaultName string, resourceGroupName string, fabricName string) (result autorest.Response, err error) {
	req, err := client.RefreshPreparer(vaultName, resourceGroupName, fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainersClient", "Refresh", nil, "Failure preparing request")
		return
	}

	resp, err := client.RefreshSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainersClient", "Refresh", resp, "Failure sending request")
		return
	}

	result, err = client.RefreshResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectionContainersClient", "Refresh", resp, "Failure responding to request")
	}

	return
}

// RefreshPreparer prepares the Refresh request.
func (client ProtectionContainersClient) RefreshPreparer(vaultName string, resourceGroupName string, fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/refreshContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RefreshSender sends the Refresh request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionContainersClient) RefreshSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RefreshResponder handles the response to the Refresh request. The method always
// closes the http.Response Body.
func (client ProtectionContainersClient) RefreshResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
