// Package backup implements the Azure ARM Backup service API version .
//
// Open API 2.0 Specs for Azure RecoveryServices Backup service
package backup

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Backup
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Backup.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// GetOperationStatus sends the get operation status request.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// privateEndpointConnectionName - the name of the private endpoint connection.
// operationID - operation id
func (client BaseClient) GetOperationStatus(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationStatusPreparer(ctx, vaultName, resourceGroupName, privateEndpointConnectionName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BaseClient", "GetOperationStatus", resp, "Failure responding to request")
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client BaseClient) GetOperationStatusPreparer(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":                   autorest.Encode("path", operationID),
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"vaultName":                     autorest.Encode("path", vaultName),
	}

	const APIVersion = "2020-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}/operationsStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client BaseClient) GetOperationStatusResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
