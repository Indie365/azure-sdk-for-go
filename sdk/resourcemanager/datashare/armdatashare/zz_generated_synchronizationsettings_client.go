//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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

// SynchronizationSettingsClient contains the methods for the SynchronizationSettings group.
// Don't use this type directly, use NewSynchronizationSettingsClient() instead.
type SynchronizationSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSynchronizationSettingsClient creates a new instance of SynchronizationSettingsClient with the specified values.
// subscriptionID - The subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSynchronizationSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SynchronizationSettingsClient, error) {
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
	client := &SynchronizationSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create a synchronizationSetting
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share to add the synchronization setting to.
// synchronizationSettingName - The name of the synchronizationSetting.
// synchronizationSetting - The new synchronization setting information.
// options - SynchronizationSettingsClientCreateOptions contains the optional parameters for the SynchronizationSettingsClient.Create
// method.
func (client *SynchronizationSettingsClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting SynchronizationSettingClassification, options *SynchronizationSettingsClientCreateOptions) (SynchronizationSettingsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, synchronizationSetting, options)
	if err != nil {
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SynchronizationSettingsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SynchronizationSettingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting SynchronizationSettingClassification, options *SynchronizationSettingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if synchronizationSettingName == "" {
		return nil, errors.New("parameter synchronizationSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{synchronizationSettingName}", url.PathEscape(synchronizationSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, synchronizationSetting)
}

// createHandleResponse handles the Create response.
func (client *SynchronizationSettingsClient) createHandleResponse(resp *http.Response) (SynchronizationSettingsClientCreateResponse, error) {
	result := SynchronizationSettingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a synchronizationSetting in a share
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// synchronizationSettingName - The name of the synchronizationSetting .
// options - SynchronizationSettingsClientBeginDeleteOptions contains the optional parameters for the SynchronizationSettingsClient.BeginDelete
// method.
func (client *SynchronizationSettingsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientBeginDeleteOptions) (*runtime.Poller[SynchronizationSettingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SynchronizationSettingsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SynchronizationSettingsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a synchronizationSetting in a share
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
func (client *SynchronizationSettingsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, options)
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
func (client *SynchronizationSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if synchronizationSettingName == "" {
		return nil, errors.New("parameter synchronizationSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{synchronizationSettingName}", url.PathEscape(synchronizationSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a synchronizationSetting in a share
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// synchronizationSettingName - The name of the synchronizationSetting.
// options - SynchronizationSettingsClientGetOptions contains the optional parameters for the SynchronizationSettingsClient.Get
// method.
func (client *SynchronizationSettingsClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientGetOptions) (SynchronizationSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, options)
	if err != nil {
		return SynchronizationSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SynchronizationSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SynchronizationSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SynchronizationSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	if synchronizationSettingName == "" {
		return nil, errors.New("parameter synchronizationSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{synchronizationSettingName}", url.PathEscape(synchronizationSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SynchronizationSettingsClient) getHandleResponse(resp *http.Response) (SynchronizationSettingsClientGetResponse, error) {
	result := SynchronizationSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SynchronizationSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySharePager - List synchronizationSettings in a share
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-09-01
// resourceGroupName - The resource group name.
// accountName - The name of the share account.
// shareName - The name of the share.
// options - SynchronizationSettingsClientListByShareOptions contains the optional parameters for the SynchronizationSettingsClient.ListByShare
// method.
func (client *SynchronizationSettingsClient) NewListBySharePager(resourceGroupName string, accountName string, shareName string, options *SynchronizationSettingsClientListByShareOptions) *runtime.Pager[SynchronizationSettingsClientListByShareResponse] {
	return runtime.NewPager(runtime.PagingHandler[SynchronizationSettingsClientListByShareResponse]{
		More: func(page SynchronizationSettingsClientListByShareResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SynchronizationSettingsClientListByShareResponse) (SynchronizationSettingsClientListByShareResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByShareCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SynchronizationSettingsClientListByShareResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SynchronizationSettingsClientListByShareResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SynchronizationSettingsClientListByShareResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByShareHandleResponse(resp)
		},
	})
}

// listByShareCreateRequest creates the ListByShare request.
func (client *SynchronizationSettingsClient) listByShareCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SynchronizationSettingsClientListByShareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByShareHandleResponse handles the ListByShare response.
func (client *SynchronizationSettingsClient) listByShareHandleResponse(resp *http.Response) (SynchronizationSettingsClientListByShareResponse, error) {
	result := SynchronizationSettingsClientListByShareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SynchronizationSettingList); err != nil {
		return SynchronizationSettingsClientListByShareResponse{}, err
	}
	return result, nil
}