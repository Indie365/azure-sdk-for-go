package backup

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

// ValidateOperationStatusesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ValidateOperationStatusesClient struct {
	BaseClient
}

// NewValidateOperationStatusesClient creates an instance of the ValidateOperationStatusesClient client.
func NewValidateOperationStatusesClient(subscriptionID string) ValidateOperationStatusesClient {
	return NewValidateOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewValidateOperationStatusesClientWithBaseURI creates an instance of the ValidateOperationStatusesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewValidateOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) ValidateOperationStatusesClient {
	return ValidateOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get fetches the status of a triggered validate operation. The status can be in progress, completed
// or failed. You can refer to the OperationStatus enum for all the possible states of the operation.
// If operation has completed, this method returns the list of errors obtained while validating the operation.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// operationID - operationID represents the operation whose status needs to be fetched.
func (client ValidateOperationStatusesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ValidateOperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ValidateOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ValidateOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ValidateOperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ValidateOperationStatusesClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupValidateOperationsStatuses/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ValidateOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ValidateOperationStatusesClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
