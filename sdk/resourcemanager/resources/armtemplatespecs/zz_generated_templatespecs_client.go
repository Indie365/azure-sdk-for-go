//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TemplateSpecsClient contains the methods for the TemplateSpecs group.
// Don't use this type directly, use NewTemplateSpecsClient() instead.
type TemplateSpecsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTemplateSpecsClient creates a new instance of TemplateSpecsClient with the specified values.
func NewTemplateSpecsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TemplateSpecsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TemplateSpecsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates a Template Spec.
// If the operation fails it returns the *TemplateSpecsError error type.
func (client *TemplateSpecsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpec TemplateSpec, options *TemplateSpecsCreateOrUpdateOptions) (TemplateSpecsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, templateSpecName, templateSpec, options)
	if err != nil {
		return TemplateSpecsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TemplateSpecsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TemplateSpecsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TemplateSpecsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpec TemplateSpec, options *TemplateSpecsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, templateSpec)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TemplateSpecsClient) createOrUpdateHandleResponse(resp *http.Response) (TemplateSpecsCreateOrUpdateResponse, error) {
	result := TemplateSpecsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return TemplateSpecsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TemplateSpecsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := TemplateSpecsError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a Template Spec by name. When operation completes, status code 200 returned without content.
// If the operation fails it returns the *TemplateSpecsError error type.
func (client *TemplateSpecsClient) Delete(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecsDeleteOptions) (TemplateSpecsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, templateSpecName, options)
	if err != nil {
		return TemplateSpecsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TemplateSpecsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TemplateSpecsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return TemplateSpecsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TemplateSpecsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *TemplateSpecsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := TemplateSpecsError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a Template Spec with a given name.
// If the operation fails it returns the *TemplateSpecsError error type.
func (client *TemplateSpecsClient) Get(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecsGetOptions) (TemplateSpecsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, templateSpecName, options)
	if err != nil {
		return TemplateSpecsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TemplateSpecsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TemplateSpecsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TemplateSpecsClient) getCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TemplateSpecsClient) getHandleResponse(resp *http.Response) (TemplateSpecsGetResponse, error) {
	result := TemplateSpecsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return TemplateSpecsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TemplateSpecsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := TemplateSpecsError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists all the Template Specs within the specified resource group.
// If the operation fails it returns the *TemplateSpecsError error type.
func (client *TemplateSpecsClient) ListByResourceGroup(resourceGroupName string, options *TemplateSpecsListByResourceGroupOptions) *TemplateSpecsListByResourceGroupPager {
	return &TemplateSpecsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp TemplateSpecsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TemplateSpecsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TemplateSpecsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *TemplateSpecsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *TemplateSpecsClient) listByResourceGroupHandleResponse(resp *http.Response) (TemplateSpecsListByResourceGroupResponse, error) {
	result := TemplateSpecsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecsListResult); err != nil {
		return TemplateSpecsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *TemplateSpecsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := TemplateSpecsError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists all the Template Specs within the specified subscriptions.
// If the operation fails it returns the *TemplateSpecsError error type.
func (client *TemplateSpecsClient) ListBySubscription(options *TemplateSpecsListBySubscriptionOptions) *TemplateSpecsListBySubscriptionPager {
	return &TemplateSpecsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp TemplateSpecsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TemplateSpecsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *TemplateSpecsClient) listBySubscriptionCreateRequest(ctx context.Context, options *TemplateSpecsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/templateSpecs/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *TemplateSpecsClient) listBySubscriptionHandleResponse(resp *http.Response) (TemplateSpecsListBySubscriptionResponse, error) {
	result := TemplateSpecsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecsListResult); err != nil {
		return TemplateSpecsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *TemplateSpecsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := TemplateSpecsError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates Template Spec tags with specified values.
// If the operation fails it returns the *TemplateSpecsError error type.
func (client *TemplateSpecsClient) Update(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecsUpdateOptions) (TemplateSpecsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, templateSpecName, options)
	if err != nil {
		return TemplateSpecsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TemplateSpecsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TemplateSpecsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TemplateSpecsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.TemplateSpec != nil {
		return req, runtime.MarshalAsJSON(req, *options.TemplateSpec)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TemplateSpecsClient) updateHandleResponse(resp *http.Response) (TemplateSpecsUpdateResponse, error) {
	result := TemplateSpecsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpec); err != nil {
		return TemplateSpecsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *TemplateSpecsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := TemplateSpecsError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
