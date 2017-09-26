package recoveryservicessiterecovery

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

// ReplicationProtectableItemsClient is the client for the ReplicationProtectableItems methods of the
// Recoveryservicessiterecovery service.
type ReplicationProtectableItemsClient struct {
	ManagementClient
}

// NewReplicationProtectableItemsClient creates an instance of the ReplicationProtectableItemsClient client.
func NewReplicationProtectableItemsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectableItemsClient {
	return NewReplicationProtectableItemsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationProtectableItemsClientWithBaseURI creates an instance of the ReplicationProtectableItemsClient client.
func NewReplicationProtectableItemsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationProtectableItemsClient {
	return ReplicationProtectableItemsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get the operation to get the details of a protectable item.
//
// fabricName is fabric name. protectionContainerName is protection container name. protectableItemName is protectable
// item name.
func (client ReplicationProtectableItemsClient) Get(fabricName string, protectionContainerName string, protectableItemName string) (result ProtectableItem, err error) {
	req, err := client.GetPreparer(fabricName, protectionContainerName, protectableItemName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationProtectableItemsClient) GetPreparer(fabricName string, protectionContainerName string, protectableItemName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectableItemName":     autorest.Encode("path", protectableItemName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems/{protectableItemName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectableItemsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationProtectableItemsClient) GetResponder(resp *http.Response) (result ProtectableItem, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationProtectionContainers lists the protectable items in a protection container.
//
// fabricName is fabric name. protectionContainerName is protection container name.
func (client ReplicationProtectableItemsClient) ListByReplicationProtectionContainers(fabricName string, protectionContainerName string) (result ProtectableItemCollection, err error) {
	req, err := client.ListByReplicationProtectionContainersPreparer(fabricName, protectionContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "ListByReplicationProtectionContainers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationProtectionContainersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "ListByReplicationProtectionContainers", resp, "Failure sending request")
		return
	}

	result, err = client.ListByReplicationProtectionContainersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "ListByReplicationProtectionContainers", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationProtectionContainersPreparer prepares the ListByReplicationProtectionContainers request.
func (client ReplicationProtectableItemsClient) ListByReplicationProtectionContainersPreparer(fabricName string, protectionContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":              autorest.Encode("path", fabricName),
		"protectionContainerName": autorest.Encode("path", protectionContainerName),
		"resourceGroupName":       autorest.Encode("path", client.ResourceGroupName),
		"resourceName":            autorest.Encode("path", client.ResourceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByReplicationProtectionContainersSender sends the ListByReplicationProtectionContainers request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationProtectableItemsClient) ListByReplicationProtectionContainersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByReplicationProtectionContainersResponder handles the response to the ListByReplicationProtectionContainers request. The method always
// closes the http.Response Body.
func (client ReplicationProtectableItemsClient) ListByReplicationProtectionContainersResponder(resp *http.Response) (result ProtectableItemCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationProtectionContainersNextResults retrieves the next set of results, if any.
func (client ReplicationProtectableItemsClient) ListByReplicationProtectionContainersNextResults(lastResults ProtectableItemCollection) (result ProtectableItemCollection, err error) {
	req, err := lastResults.ProtectableItemCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "ListByReplicationProtectionContainers", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByReplicationProtectionContainersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "ListByReplicationProtectionContainers", resp, "Failure sending next results request")
	}

	result, err = client.ListByReplicationProtectionContainersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationProtectableItemsClient", "ListByReplicationProtectionContainers", resp, "Failure responding to next results request")
	}

	return
}

// ListByReplicationProtectionContainersComplete gets all elements from the list without paging.
func (client ReplicationProtectableItemsClient) ListByReplicationProtectionContainersComplete(fabricName string, protectionContainerName string, cancel <-chan struct{}) (<-chan ProtectableItem, <-chan error) {
	resultChan := make(chan ProtectableItem)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByReplicationProtectionContainers(fabricName, protectionContainerName)
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
			list, err = client.ListByReplicationProtectionContainersNextResults(list)
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
