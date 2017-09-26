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

// BackupProtectionContainersClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type BackupProtectionContainersClient struct {
	ManagementClient
}

// NewBackupProtectionContainersClient creates an instance of the BackupProtectionContainersClient client.
func NewBackupProtectionContainersClient(subscriptionID string) BackupProtectionContainersClient {
	return NewBackupProtectionContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupProtectionContainersClientWithBaseURI creates an instance of the BackupProtectionContainersClient client.
func NewBackupProtectionContainersClientWithBaseURI(baseURI string, subscriptionID string) BackupProtectionContainersClient {
	return BackupProtectionContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the containers registered to Recovery Services Vault.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. filter is oData filter options.
func (client BackupProtectionContainersClient) List(vaultName string, resourceGroupName string, filter string) (result ProtectionContainerResourceList, err error) {
	req, err := client.ListPreparer(vaultName, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectionContainersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectionContainersClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectionContainersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BackupProtectionContainersClient) ListPreparer(vaultName string, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BackupProtectionContainersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BackupProtectionContainersClient) ListResponder(resp *http.Response) (result ProtectionContainerResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client BackupProtectionContainersClient) ListNextResults(lastResults ProtectionContainerResourceList) (result ProtectionContainerResourceList, err error) {
	req, err := lastResults.ProtectionContainerResourceListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectionContainersClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectionContainersClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.BackupProtectionContainersClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client BackupProtectionContainersClient) ListComplete(vaultName string, resourceGroupName string, filter string, cancel <-chan struct{}) (<-chan ProtectionContainerResource, <-chan error) {
	resultChan := make(chan ProtectionContainerResource)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(vaultName, resourceGroupName, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
