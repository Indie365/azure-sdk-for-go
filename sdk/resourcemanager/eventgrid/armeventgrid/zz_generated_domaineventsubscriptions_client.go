//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DomainEventSubscriptionsClient contains the methods for the DomainEventSubscriptions group.
// Don't use this type directly, use NewDomainEventSubscriptionsClient() instead.
type DomainEventSubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDomainEventSubscriptionsClient creates a new instance of DomainEventSubscriptionsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDomainEventSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DomainEventSubscriptionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DomainEventSubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates a new event subscription or updates an existing event subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain topic.
// eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// eventSubscriptionInfo - Event subscription properties containing the destination and filter information.
// options - DomainEventSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the DomainEventSubscriptionsClient.BeginCreateOrUpdate
// method.
func (client *DomainEventSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *DomainEventSubscriptionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[DomainEventSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, domainName, eventSubscriptionName, eventSubscriptionInfo, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DomainEventSubscriptionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DomainEventSubscriptionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Asynchronously creates a new event subscription or updates an existing event subscription.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DomainEventSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *DomainEventSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainName, eventSubscriptionName, eventSubscriptionInfo, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainEventSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *DomainEventSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventSubscriptionInfo)
}

// BeginDelete - Delete an existing event subscription for a domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// eventSubscriptionName - Name of the event subscription to be deleted. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// options - DomainEventSubscriptionsClientBeginDeleteOptions contains the optional parameters for the DomainEventSubscriptionsClient.BeginDelete
// method.
func (client *DomainEventSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientBeginDeleteOptions) (*armruntime.Poller[DomainEventSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, domainName, eventSubscriptionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DomainEventSubscriptionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DomainEventSubscriptionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an existing event subscription for a domain.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DomainEventSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainName, eventSubscriptionName, options)
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
func (client *DomainEventSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of an event subscription of a domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the partner topic.
// eventSubscriptionName - Name of the event subscription to be found. Event subscription names must be between 3 and 100
// characters in length and use alphanumeric letters only.
// options - DomainEventSubscriptionsClientGetOptions contains the optional parameters for the DomainEventSubscriptionsClient.Get
// method.
func (client *DomainEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientGetOptions) (DomainEventSubscriptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainName, eventSubscriptionName, options)
	if err != nil {
		return DomainEventSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainEventSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainEventSubscriptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DomainEventSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainEventSubscriptionsClient) getHandleResponse(resp *http.Response) (DomainEventSubscriptionsClientGetResponse, error) {
	result := DomainEventSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscription); err != nil {
		return DomainEventSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeliveryAttributes - Get all delivery attributes for an event subscription for domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain topic.
// eventSubscriptionName - Name of the event subscription.
// options - DomainEventSubscriptionsClientGetDeliveryAttributesOptions contains the optional parameters for the DomainEventSubscriptionsClient.GetDeliveryAttributes
// method.
func (client *DomainEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientGetDeliveryAttributesOptions) (DomainEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	req, err := client.getDeliveryAttributesCreateRequest(ctx, resourceGroupName, domainName, eventSubscriptionName, options)
	if err != nil {
		return DomainEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainEventSubscriptionsClientGetDeliveryAttributesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeliveryAttributesHandleResponse(resp)
}

// getDeliveryAttributesCreateRequest creates the GetDeliveryAttributes request.
func (client *DomainEventSubscriptionsClient) getDeliveryAttributesCreateRequest(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientGetDeliveryAttributesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDeliveryAttributesHandleResponse handles the GetDeliveryAttributes response.
func (client *DomainEventSubscriptionsClient) getDeliveryAttributesHandleResponse(resp *http.Response) (DomainEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	result := DomainEventSubscriptionsClientGetDeliveryAttributesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeliveryAttributeListResult); err != nil {
		return DomainEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	return result, nil
}

// GetFullURL - Get the full endpoint URL for an event subscription for domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain topic.
// eventSubscriptionName - Name of the event subscription.
// options - DomainEventSubscriptionsClientGetFullURLOptions contains the optional parameters for the DomainEventSubscriptionsClient.GetFullURL
// method.
func (client *DomainEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientGetFullURLOptions) (DomainEventSubscriptionsClientGetFullURLResponse, error) {
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, domainName, eventSubscriptionName, options)
	if err != nil {
		return DomainEventSubscriptionsClientGetFullURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainEventSubscriptionsClientGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainEventSubscriptionsClientGetFullURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFullURLHandleResponse(resp)
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *DomainEventSubscriptionsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, options *DomainEventSubscriptionsClientGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *DomainEventSubscriptionsClient) getFullURLHandleResponse(resp *http.Response) (DomainEventSubscriptionsClientGetFullURLResponse, error) {
	result := DomainEventSubscriptionsClientGetFullURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionFullURL); err != nil {
		return DomainEventSubscriptionsClientGetFullURLResponse{}, err
	}
	return result, nil
}

// NewListPager - List all event subscriptions that have been created for a specific topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// options - DomainEventSubscriptionsClientListOptions contains the optional parameters for the DomainEventSubscriptionsClient.List
// method.
func (client *DomainEventSubscriptionsClient) NewListPager(resourceGroupName string, domainName string, options *DomainEventSubscriptionsClientListOptions) *runtime.Pager[DomainEventSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[DomainEventSubscriptionsClientListResponse]{
		More: func(page DomainEventSubscriptionsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DomainEventSubscriptionsClientListResponse) (DomainEventSubscriptionsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, domainName, options)
			if err != nil {
				return DomainEventSubscriptionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DomainEventSubscriptionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DomainEventSubscriptionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DomainEventSubscriptionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainEventSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DomainEventSubscriptionsClient) listHandleResponse(resp *http.Response) (DomainEventSubscriptionsClientListResponse, error) {
	result := DomainEventSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionsListResult); err != nil {
		return DomainEventSubscriptionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an existing event subscription for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// eventSubscriptionName - Name of the event subscription to be updated.
// eventSubscriptionUpdateParameters - Updated event subscription information.
// options - DomainEventSubscriptionsClientBeginUpdateOptions contains the optional parameters for the DomainEventSubscriptionsClient.BeginUpdate
// method.
func (client *DomainEventSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *DomainEventSubscriptionsClientBeginUpdateOptions) (*armruntime.Poller[DomainEventSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, domainName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[DomainEventSubscriptionsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[DomainEventSubscriptionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update an existing event subscription for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DomainEventSubscriptionsClient) update(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *DomainEventSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *DomainEventSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *DomainEventSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventSubscriptionUpdateParameters)
}
