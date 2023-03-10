//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RegistriesClient contains the methods for the Registries group.
// Don't use this type directly, use NewRegistriesClient() instead.
type RegistriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistriesClient creates a new instance of RegistriesClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistriesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+".RegistriesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks whether the container registry name is available for use. The name must contain only alphanumeric
// characters, be globally unique, and between 5 and 50 characters in length.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - registryNameCheckRequest - The object containing information for the availability request.
//   - options - RegistriesClientCheckNameAvailabilityOptions contains the optional parameters for the RegistriesClient.CheckNameAvailability
//     method.
func (client *RegistriesClient) CheckNameAvailability(ctx context.Context, registryNameCheckRequest RegistryNameCheckRequest, options *RegistriesClientCheckNameAvailabilityOptions) (RegistriesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, registryNameCheckRequest, options)
	if err != nil {
		return RegistriesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistriesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistriesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *RegistriesClient) checkNameAvailabilityCreateRequest(ctx context.Context, registryNameCheckRequest RegistryNameCheckRequest, options *RegistriesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, registryNameCheckRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *RegistriesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (RegistriesClientCheckNameAvailabilityResponse, error) {
	result := RegistriesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistryNameStatus); err != nil {
		return RegistriesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Creates a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - registry - The parameters for creating a container registry.
//   - options - RegistriesClientBeginCreateOptions contains the optional parameters for the RegistriesClient.BeginCreate method.
func (client *RegistriesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, registry Registry, options *RegistriesClientBeginCreateOptions) (*runtime.Poller[RegistriesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, registry, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RegistriesClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RegistriesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *RegistriesClient) create(ctx context.Context, resourceGroupName string, registryName string, registry Registry, options *RegistriesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, registry, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *RegistriesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, registry Registry, options *RegistriesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, registry)
}

// BeginDelete - Deletes a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - RegistriesClientBeginDeleteOptions contains the optional parameters for the RegistriesClient.BeginDelete method.
func (client *RegistriesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientBeginDeleteOptions) (*runtime.Poller[RegistriesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RegistriesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RegistriesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *RegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the properties of the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - RegistriesClientGetOptions contains the optional parameters for the RegistriesClient.Get method.
func (client *RegistriesClient) Get(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientGetOptions) (RegistriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, options)
	if err != nil {
		return RegistriesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistriesClient) getHandleResponse(resp *http.Response) (RegistriesClientGetResponse, error) {
	result := RegistriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Registry); err != nil {
		return RegistriesClientGetResponse{}, err
	}
	return result, nil
}

// BeginImportImage - Copies an image to this container registry from the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - parameters - The parameters specifying the image to copy and the source container registry.
//   - options - RegistriesClientBeginImportImageOptions contains the optional parameters for the RegistriesClient.BeginImportImage
//     method.
func (client *RegistriesClient) BeginImportImage(ctx context.Context, resourceGroupName string, registryName string, parameters ImportImageParameters, options *RegistriesClientBeginImportImageOptions) (*runtime.Poller[RegistriesClientImportImageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.importImage(ctx, resourceGroupName, registryName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RegistriesClientImportImageResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RegistriesClientImportImageResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ImportImage - Copies an image to this container registry from the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *RegistriesClient) importImage(ctx context.Context, resourceGroupName string, registryName string, parameters ImportImageParameters, options *RegistriesClientBeginImportImageOptions) (*http.Response, error) {
	req, err := client.importImageCreateRequest(ctx, resourceGroupName, registryName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// importImageCreateRequest creates the ImportImage request.
func (client *RegistriesClient) importImageCreateRequest(ctx context.Context, resourceGroupName string, registryName string, parameters ImportImageParameters, options *RegistriesClientBeginImportImageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importImage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// NewListPager - Lists all the container registries under the specified subscription.
//
// Generated from API version 2019-05-01
//   - options - RegistriesClientListOptions contains the optional parameters for the RegistriesClient.NewListPager method.
func (client *RegistriesClient) NewListPager(options *RegistriesClientListOptions) *runtime.Pager[RegistriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistriesClientListResponse]{
		More: func(page RegistriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistriesClientListResponse) (RegistriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistriesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RegistriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistriesClient) listCreateRequest(ctx context.Context, options *RegistriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistriesClient) listHandleResponse(resp *http.Response) (RegistriesClientListResponse, error) {
	result := RegistriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistryListResult); err != nil {
		return RegistriesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the container registries under the specified resource group.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - options - RegistriesClientListByResourceGroupOptions contains the optional parameters for the RegistriesClient.NewListByResourceGroupPager
//     method.
func (client *RegistriesClient) NewListByResourceGroupPager(resourceGroupName string, options *RegistriesClientListByResourceGroupOptions) *runtime.Pager[RegistriesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistriesClientListByResourceGroupResponse]{
		More: func(page RegistriesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistriesClientListByResourceGroupResponse) (RegistriesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistriesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RegistriesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistriesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RegistriesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RegistriesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RegistriesClient) listByResourceGroupHandleResponse(resp *http.Response) (RegistriesClientListByResourceGroupResponse, error) {
	result := RegistriesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistryListResult); err != nil {
		return RegistriesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListCredentials - Lists the login credentials for the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - RegistriesClientListCredentialsOptions contains the optional parameters for the RegistriesClient.ListCredentials
//     method.
func (client *RegistriesClient) ListCredentials(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientListCredentialsOptions) (RegistriesClientListCredentialsResponse, error) {
	req, err := client.listCredentialsCreateRequest(ctx, resourceGroupName, registryName, options)
	if err != nil {
		return RegistriesClientListCredentialsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistriesClientListCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistriesClientListCredentialsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listCredentialsHandleResponse(resp)
}

// listCredentialsCreateRequest creates the ListCredentials request.
func (client *RegistriesClient) listCredentialsCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientListCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/listCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCredentialsHandleResponse handles the ListCredentials response.
func (client *RegistriesClient) listCredentialsHandleResponse(resp *http.Response) (RegistriesClientListCredentialsResponse, error) {
	result := RegistriesClientListCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistryListCredentialsResult); err != nil {
		return RegistriesClientListCredentialsResponse{}, err
	}
	return result, nil
}

// ListUsages - Gets the quota usages for the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - RegistriesClientListUsagesOptions contains the optional parameters for the RegistriesClient.ListUsages method.
func (client *RegistriesClient) ListUsages(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientListUsagesOptions) (RegistriesClientListUsagesResponse, error) {
	req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, registryName, options)
	if err != nil {
		return RegistriesClientListUsagesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistriesClientListUsagesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistriesClientListUsagesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listUsagesHandleResponse(resp)
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *RegistriesClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistriesClientListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/listUsages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *RegistriesClient) listUsagesHandleResponse(resp *http.Response) (RegistriesClientListUsagesResponse, error) {
	result := RegistriesClientListUsagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistryUsageListResult); err != nil {
		return RegistriesClientListUsagesResponse{}, err
	}
	return result, nil
}

// RegenerateCredential - Regenerates one of the login credentials for the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - regenerateCredentialParameters - Specifies name of the password which should be regenerated -- password or password2.
//   - options - RegistriesClientRegenerateCredentialOptions contains the optional parameters for the RegistriesClient.RegenerateCredential
//     method.
func (client *RegistriesClient) RegenerateCredential(ctx context.Context, resourceGroupName string, registryName string, regenerateCredentialParameters RegenerateCredentialParameters, options *RegistriesClientRegenerateCredentialOptions) (RegistriesClientRegenerateCredentialResponse, error) {
	req, err := client.regenerateCredentialCreateRequest(ctx, resourceGroupName, registryName, regenerateCredentialParameters, options)
	if err != nil {
		return RegistriesClientRegenerateCredentialResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistriesClientRegenerateCredentialResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistriesClientRegenerateCredentialResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateCredentialHandleResponse(resp)
}

// regenerateCredentialCreateRequest creates the RegenerateCredential request.
func (client *RegistriesClient) regenerateCredentialCreateRequest(ctx context.Context, resourceGroupName string, registryName string, regenerateCredentialParameters RegenerateCredentialParameters, options *RegistriesClientRegenerateCredentialOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/regenerateCredential"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, regenerateCredentialParameters)
}

// regenerateCredentialHandleResponse handles the RegenerateCredential response.
func (client *RegistriesClient) regenerateCredentialHandleResponse(resp *http.Response) (RegistriesClientRegenerateCredentialResponse, error) {
	result := RegistriesClientRegenerateCredentialResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistryListCredentialsResult); err != nil {
		return RegistriesClientRegenerateCredentialResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - registryUpdateParameters - The parameters for updating a container registry.
//   - options - RegistriesClientBeginUpdateOptions contains the optional parameters for the RegistriesClient.BeginUpdate method.
func (client *RegistriesClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters, options *RegistriesClientBeginUpdateOptions) (*runtime.Poller[RegistriesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, registryUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RegistriesClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RegistriesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *RegistriesClient) update(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters, options *RegistriesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, registryUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *RegistriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters, options *RegistriesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, registryUpdateParameters)
}
