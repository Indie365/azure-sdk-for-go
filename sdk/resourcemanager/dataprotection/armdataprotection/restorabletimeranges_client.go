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

// RestorableTimeRangesClient contains the methods for the RestorableTimeRanges group.
// Don't use this type directly, use NewRestorableTimeRangesClient() instead.
type RestorableTimeRangesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableTimeRangesClient creates a new instance of RestorableTimeRangesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableTimeRangesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableTimeRangesClient, error) {
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
	client := &RestorableTimeRangesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Find -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group where the backup vault is present.
//   - vaultName - The name of the backup vault.
//   - backupInstanceName - The name of the backup instance
//   - parameters - Request body for operation
//   - options - RestorableTimeRangesClientFindOptions contains the optional parameters for the RestorableTimeRangesClient.Find
//     method.
func (client *RestorableTimeRangesClient) Find(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, parameters AzureBackupFindRestorableTimeRangesRequest, options *RestorableTimeRangesClientFindOptions) (RestorableTimeRangesClientFindResponse, error) {
	req, err := client.findCreateRequest(ctx, resourceGroupName, vaultName, backupInstanceName, parameters, options)
	if err != nil {
		return RestorableTimeRangesClientFindResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableTimeRangesClientFindResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableTimeRangesClientFindResponse{}, runtime.NewResponseError(resp)
	}
	return client.findHandleResponse(resp)
}

// findCreateRequest creates the Find request.
func (client *RestorableTimeRangesClient) findCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, parameters AzureBackupFindRestorableTimeRangesRequest, options *RestorableTimeRangesClientFindOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/findRestorableTimeRanges"
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
	if backupInstanceName == "" {
		return nil, errors.New("parameter backupInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupInstanceName}", url.PathEscape(backupInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// findHandleResponse handles the Find response.
func (client *RestorableTimeRangesClient) findHandleResponse(resp *http.Response) (RestorableTimeRangesClientFindResponse, error) {
	result := RestorableTimeRangesClientFindResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBackupFindRestorableTimeRangesResponseResource); err != nil {
		return RestorableTimeRangesClientFindResponse{}, err
	}
	return result, nil
}
