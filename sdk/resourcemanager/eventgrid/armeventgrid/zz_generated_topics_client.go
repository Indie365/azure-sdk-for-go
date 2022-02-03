//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TopicsClient contains the methods for the Topics group.
// Don't use this type directly, use NewTopicsClient() instead.
type TopicsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTopicsClient creates a new instance of TopicsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTopicsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TopicsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &TopicsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Asynchronously creates a new topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// topicName - Name of the topic.
// topicInfo - Topic information.
// options - TopicsClientBeginCreateOrUpdateOptions contains the optional parameters for the TopicsClient.BeginCreateOrUpdate
// method.
func (client *TopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo Topic, options *TopicsClientBeginCreateOrUpdateOptions) (TopicsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, topicName, topicInfo, options)
	if err != nil {
		return TopicsClientCreateOrUpdatePollerResponse{}, err
	}
	result := TopicsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TopicsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return TopicsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &TopicsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates a new topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo Topic, options *TopicsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, topicName, topicInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, topicName string, topicInfo Topic, options *TopicsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, topicInfo)
}

// BeginDelete - Delete existing topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// topicName - Name of the topic.
// options - TopicsClientBeginDeleteOptions contains the optional parameters for the TopicsClient.BeginDelete method.
func (client *TopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientBeginDeleteOptions) (TopicsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, topicName, options)
	if err != nil {
		return TopicsClientDeletePollerResponse{}, err
	}
	result := TopicsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TopicsClient.Delete", "", resp, client.pl)
	if err != nil {
		return TopicsClientDeletePollerResponse{}, err
	}
	result.Poller = &TopicsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete existing topic.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, topicName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// topicName - Name of the topic.
// options - TopicsClientGetOptions contains the optional parameters for the TopicsClient.Get method.
func (client *TopicsClient) Get(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientGetOptions) (TopicsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, topicName, options)
	if err != nil {
		return TopicsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TopicsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TopicsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicsClient) getHandleResponse(resp *http.Response) (TopicsClientGetResponse, error) {
	result := TopicsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Topic); err != nil {
		return TopicsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - List all the topics under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// options - TopicsClientListByResourceGroupOptions contains the optional parameters for the TopicsClient.ListByResourceGroup
// method.
func (client *TopicsClient) ListByResourceGroup(resourceGroupName string, options *TopicsClientListByResourceGroupOptions) *TopicsClientListByResourceGroupPager {
	return &TopicsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp TopicsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TopicsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TopicsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *TopicsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *TopicsClient) listByResourceGroupHandleResponse(resp *http.Response) (TopicsClientListByResourceGroupResponse, error) {
	result := TopicsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicsListResult); err != nil {
		return TopicsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - List all the topics under an Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - TopicsClientListBySubscriptionOptions contains the optional parameters for the TopicsClient.ListBySubscription
// method.
func (client *TopicsClient) ListBySubscription(options *TopicsClientListBySubscriptionOptions) *TopicsClientListBySubscriptionPager {
	return &TopicsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp TopicsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TopicsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *TopicsClient) listBySubscriptionCreateRequest(ctx context.Context, options *TopicsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *TopicsClient) listBySubscriptionHandleResponse(resp *http.Response) (TopicsClientListBySubscriptionResponse, error) {
	result := TopicsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicsListResult); err != nil {
		return TopicsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListEventTypes - List event types for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// providerNamespace - Namespace of the provider of the topic.
// resourceTypeName - Name of the topic type.
// resourceName - Name of the topic.
// options - TopicsClientListEventTypesOptions contains the optional parameters for the TopicsClient.ListEventTypes method.
func (client *TopicsClient) ListEventTypes(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, options *TopicsClientListEventTypesOptions) (TopicsClientListEventTypesResponse, error) {
	req, err := client.listEventTypesCreateRequest(ctx, resourceGroupName, providerNamespace, resourceTypeName, resourceName, options)
	if err != nil {
		return TopicsClientListEventTypesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TopicsClientListEventTypesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TopicsClientListEventTypesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listEventTypesHandleResponse(resp)
}

// listEventTypesCreateRequest creates the ListEventTypes request.
func (client *TopicsClient) listEventTypesCreateRequest(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, options *TopicsClientListEventTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if resourceTypeName == "" {
		return nil, errors.New("parameter resourceTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceTypeName}", url.PathEscape(resourceTypeName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listEventTypesHandleResponse handles the ListEventTypes response.
func (client *TopicsClient) listEventTypesHandleResponse(resp *http.Response) (TopicsClientListEventTypesResponse, error) {
	result := TopicsClientListEventTypesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventTypesListResult); err != nil {
		return TopicsClientListEventTypesResponse{}, err
	}
	return result, nil
}

// ListSharedAccessKeys - List the two keys used to publish to a topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// topicName - Name of the topic.
// options - TopicsClientListSharedAccessKeysOptions contains the optional parameters for the TopicsClient.ListSharedAccessKeys
// method.
func (client *TopicsClient) ListSharedAccessKeys(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientListSharedAccessKeysOptions) (TopicsClientListSharedAccessKeysResponse, error) {
	req, err := client.listSharedAccessKeysCreateRequest(ctx, resourceGroupName, topicName, options)
	if err != nil {
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TopicsClientListSharedAccessKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSharedAccessKeysHandleResponse(resp)
}

// listSharedAccessKeysCreateRequest creates the ListSharedAccessKeys request.
func (client *TopicsClient) listSharedAccessKeysCreateRequest(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientListSharedAccessKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSharedAccessKeysHandleResponse handles the ListSharedAccessKeys response.
func (client *TopicsClient) listSharedAccessKeysHandleResponse(resp *http.Response) (TopicsClientListSharedAccessKeysResponse, error) {
	result := TopicsClientListSharedAccessKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicSharedAccessKeys); err != nil {
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	return result, nil
}

// BeginRegenerateKey - Regenerate a shared access key for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// topicName - Name of the topic.
// regenerateKeyRequest - Request body to regenerate key.
// options - TopicsClientBeginRegenerateKeyOptions contains the optional parameters for the TopicsClient.BeginRegenerateKey
// method.
func (client *TopicsClient) BeginRegenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest TopicRegenerateKeyRequest, options *TopicsClientBeginRegenerateKeyOptions) (TopicsClientRegenerateKeyPollerResponse, error) {
	resp, err := client.regenerateKey(ctx, resourceGroupName, topicName, regenerateKeyRequest, options)
	if err != nil {
		return TopicsClientRegenerateKeyPollerResponse{}, err
	}
	result := TopicsClientRegenerateKeyPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TopicsClient.RegenerateKey", "", resp, client.pl)
	if err != nil {
		return TopicsClientRegenerateKeyPollerResponse{}, err
	}
	result.Poller = &TopicsClientRegenerateKeyPoller{
		pt: pt,
	}
	return result, nil
}

// RegenerateKey - Regenerate a shared access key for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TopicsClient) regenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest TopicRegenerateKeyRequest, options *TopicsClientBeginRegenerateKeyOptions) (*http.Response, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, topicName, regenerateKeyRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *TopicsClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest TopicRegenerateKeyRequest, options *TopicsClientBeginRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, regenerateKeyRequest)
}

// BeginUpdate - Asynchronously updates a topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// topicName - Name of the topic.
// topicUpdateParameters - Topic update information.
// options - TopicsClientBeginUpdateOptions contains the optional parameters for the TopicsClient.BeginUpdate method.
func (client *TopicsClient) BeginUpdate(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters TopicUpdateParameters, options *TopicsClientBeginUpdateOptions) (TopicsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, topicName, topicUpdateParameters, options)
	if err != nil {
		return TopicsClientUpdatePollerResponse{}, err
	}
	result := TopicsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TopicsClient.Update", "", resp, client.pl)
	if err != nil {
		return TopicsClientUpdatePollerResponse{}, err
	}
	result.Poller = &TopicsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Asynchronously updates a topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TopicsClient) update(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters TopicUpdateParameters, options *TopicsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, topicName, topicUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *TopicsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters TopicUpdateParameters, options *TopicsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, topicUpdateParameters)
}
