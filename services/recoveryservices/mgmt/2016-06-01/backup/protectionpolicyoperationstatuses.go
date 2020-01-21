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

// ProtectionPolicyOperationStatusesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionPolicyOperationStatusesClient struct {
	BaseClient
}

// NewProtectionPolicyOperationStatusesClient creates an instance of the ProtectionPolicyOperationStatusesClient
// client.
func NewProtectionPolicyOperationStatusesClient(subscriptionID string) ProtectionPolicyOperationStatusesClient {
	return NewProtectionPolicyOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionPolicyOperationStatusesClientWithBaseURI creates an instance of the
// ProtectionPolicyOperationStatusesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProtectionPolicyOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) ProtectionPolicyOperationStatusesClient {
	return ProtectionPolicyOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get provides the status of the asynchronous operations like backup or restore. The status can be: in progress,
// completed, or failed. You can refer to the Operation Status enumeration for the possible states of an operation.
// Some operations create jobs. This method returns the list of jobs associated with the operation.
// Parameters:
// vaultName - the name of the Recovery Services vault.
// resourceGroupName - the name of the resource group associated with the Recovery Services vault.
// policyName - the backup policy name used in this GET operation.
// operationID - the ID associated with this GET operation.
func (client ProtectionPolicyOperationStatusesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectionPolicyOperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, policyName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionPolicyOperationStatusesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectionPolicyOperationStatusesClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionPolicyOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionPolicyOperationStatusesClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
