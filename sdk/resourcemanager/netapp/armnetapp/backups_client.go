//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp

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

// BackupsClient contains the methods for the Backups group.
// Don't use this type directly, use NewBackupsClient() instead.
type BackupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupsClient creates a new instance of BackupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsClient, error) {
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
	client := &BackupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create a backup for the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - backupName - The name of the backup
//   - body - Backup object supplied in the body of the operation.
//   - options - BackupsClientBeginCreateOptions contains the optional parameters for the BackupsClient.BeginCreate method.
func (client *BackupsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup, options *BackupsClientBeginCreateOptions) (*runtime.Poller[BackupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[BackupsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[BackupsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a backup for the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *BackupsClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup, options *BackupsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *BackupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body Backup, options *BackupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete a backup of the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - backupName - The name of the backup
//   - options - BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
func (client *BackupsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsClientBeginDeleteOptions) (*runtime.Poller[BackupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[BackupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[BackupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a backup of the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *BackupsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BackupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified backup of the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - backupName - The name of the backup
//   - options - BackupsClientGetOptions contains the optional parameters for the BackupsClient.Get method.
func (client *BackupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsClientGetOptions) (BackupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, options)
	if err != nil {
		return BackupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, options *BackupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupsClient) getHandleResponse(resp *http.Response) (BackupsClientGetResponse, error) {
	result := BackupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Backup); err != nil {
		return BackupsClientGetResponse{}, err
	}
	return result, nil
}

// GetStatus - Get the status of the backup for a volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - BackupsClientGetStatusOptions contains the optional parameters for the BackupsClient.GetStatus method.
func (client *BackupsClient) GetStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetStatusOptions) (BackupsClientGetStatusResponse, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsClientGetStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsClientGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsClientGetStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStatusHandleResponse(resp)
}

// getStatusCreateRequest creates the GetStatus request.
func (client *BackupsClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backupStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *BackupsClient) getStatusHandleResponse(resp *http.Response) (BackupsClientGetStatusResponse, error) {
	result := BackupsClientGetStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupStatus); err != nil {
		return BackupsClientGetStatusResponse{}, err
	}
	return result, nil
}

// GetVolumeRestoreStatus - Get the status of the restore for a volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - BackupsClientGetVolumeRestoreStatusOptions contains the optional parameters for the BackupsClient.GetVolumeRestoreStatus
//     method.
func (client *BackupsClient) GetVolumeRestoreStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetVolumeRestoreStatusOptions) (BackupsClientGetVolumeRestoreStatusResponse, error) {
	req, err := client.getVolumeRestoreStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupsClientGetVolumeRestoreStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getVolumeRestoreStatusHandleResponse(resp)
}

// getVolumeRestoreStatusCreateRequest creates the GetVolumeRestoreStatus request.
func (client *BackupsClient) getVolumeRestoreStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetVolumeRestoreStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/restoreStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getVolumeRestoreStatusHandleResponse handles the GetVolumeRestoreStatus response.
func (client *BackupsClient) getVolumeRestoreStatusHandleResponse(resp *http.Response) (BackupsClientGetVolumeRestoreStatusResponse, error) {
	result := BackupsClientGetVolumeRestoreStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestoreStatus); err != nil {
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - List all backups for a volume
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - BackupsClientListOptions contains the optional parameters for the BackupsClient.NewListPager method.
func (client *BackupsClient) NewListPager(resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientListOptions) *runtime.Pager[BackupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupsClientListResponse]{
		More: func(page BackupsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BackupsClientListResponse) (BackupsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
			if err != nil {
				return BackupsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupsClient) listHandleResponse(resp *http.Response) (BackupsClientListResponse, error) {
	result := BackupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupsList); err != nil {
		return BackupsClientListResponse{}, err
	}
	return result, nil
}

// BeginRestoreFiles - Restore the specified files from the specified backup to the active filesystem
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - backupName - The name of the backup
//   - body - Restore payload supplied in the body of the operation.
//   - options - BackupsClientBeginRestoreFilesOptions contains the optional parameters for the BackupsClient.BeginRestoreFiles
//     method.
func (client *BackupsClient) BeginRestoreFiles(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body BackupRestoreFiles, options *BackupsClientBeginRestoreFilesOptions) (*runtime.Poller[BackupsClientRestoreFilesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restoreFiles(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[BackupsClientRestoreFilesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[BackupsClientRestoreFilesResponse](options.ResumeToken, client.pl, nil)
	}
}

// RestoreFiles - Restore the specified files from the specified backup to the active filesystem
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *BackupsClient) restoreFiles(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body BackupRestoreFiles, options *BackupsClientBeginRestoreFilesOptions) (*http.Response, error) {
	req, err := client.restoreFilesCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restoreFilesCreateRequest creates the RestoreFiles request.
func (client *BackupsClient) restoreFilesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body BackupRestoreFiles, options *BackupsClientBeginRestoreFilesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}/restoreFiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginUpdate - Patch a backup for the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - backupName - The name of the backup
//   - body - Backup object supplied in the body of the operation.
//   - options - BackupsClientBeginUpdateOptions contains the optional parameters for the BackupsClient.BeginUpdate method.
func (client *BackupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body BackupPatch, options *BackupsClientBeginUpdateOptions) (*runtime.Poller[BackupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[BackupsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[BackupsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch a backup for the volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *BackupsClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body BackupPatch, options *BackupsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, backupName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *BackupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, backupName string, body BackupPatch, options *BackupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
