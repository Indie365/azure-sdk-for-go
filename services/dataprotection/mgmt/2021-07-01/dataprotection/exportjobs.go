package dataprotection

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

// ExportJobsClient is the open API 2.0 Specs for Azure Data Protection service
type ExportJobsClient struct {
	BaseClient
}

// NewExportJobsClient creates an instance of the ExportJobsClient client.
func NewExportJobsClient(subscriptionID string) ExportJobsClient {
	return NewExportJobsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExportJobsClientWithBaseURI creates an instance of the ExportJobsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExportJobsClientWithBaseURI(baseURI string, subscriptionID string) ExportJobsClient {
	return ExportJobsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Trigger triggers export of jobs and returns an OperationID to track.
// Parameters:
// resourceGroupName - the name of the resource group where the backup vault is present.
// vaultName - the name of the backup vault.
func (client ExportJobsClient) Trigger(ctx context.Context, resourceGroupName string, vaultName string) (result ExportJobsTriggerFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportJobsClient.Trigger")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TriggerPreparer(ctx, resourceGroupName, vaultName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.ExportJobsClient", "Trigger", nil, "Failure preparing request")
		return
	}

	result, err = client.TriggerSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.ExportJobsClient", "Trigger", nil, "Failure sending request")
		return
	}

	return
}

// TriggerPreparer prepares the Trigger request.
func (client ExportJobsClient) TriggerPreparer(ctx context.Context, resourceGroupName string, vaultName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/exportBackupJobs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerSender sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (client ExportJobsClient) TriggerSender(req *http.Request) (future ExportJobsTriggerFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// TriggerResponder handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (client ExportJobsClient) TriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
