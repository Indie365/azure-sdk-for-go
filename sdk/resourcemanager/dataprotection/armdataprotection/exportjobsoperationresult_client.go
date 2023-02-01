//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExportJobsOperationResultClient contains the methods for the ExportJobsOperationResult group.
// Don't use this type directly, use NewExportJobsOperationResultClient() instead.
type ExportJobsOperationResultClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExportJobsOperationResultClient creates a new instance of ExportJobsOperationResultClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExportJobsOperationResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExportJobsOperationResultClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ExportJobsOperationResultClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the operation result of operation triggered by Export Jobs API. If the operation is successful, then it also
// contains URL of a Blob and a SAS key to access the same. The blob contains exported
// jobs in JSON serialized format.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - vaultName - The name of the backup vault.
//   - operationID - OperationID which represents the export job.
//   - options - ExportJobsOperationResultClientGetOptions contains the optional parameters for the ExportJobsOperationResultClient.Get
//     method.
func (client *ExportJobsOperationResultClient) Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *ExportJobsOperationResultClientGetOptions) (ExportJobsOperationResultClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, operationID, options)
	if err != nil {
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return ExportJobsOperationResultClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExportJobsOperationResultClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *ExportJobsOperationResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupJobs/operations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExportJobsOperationResultClient) getHandleResponse(resp *http.Response) (ExportJobsOperationResultClientGetResponse, error) {
	result := ExportJobsOperationResultClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExportJobsResult); err != nil {
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	return result, nil
}
