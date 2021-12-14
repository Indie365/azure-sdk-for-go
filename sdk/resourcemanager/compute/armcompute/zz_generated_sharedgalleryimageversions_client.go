//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// SharedGalleryImageVersionsClient contains the methods for the SharedGalleryImageVersions group.
// Don't use this type directly, use NewSharedGalleryImageVersionsClient() instead.
type SharedGalleryImageVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSharedGalleryImageVersionsClient creates a new instance of SharedGalleryImageVersionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSharedGalleryImageVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SharedGalleryImageVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &SharedGalleryImageVersionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// Get - Get a shared gallery image version by subscription id or tenant id.
// If the operation fails it returns the *CloudError error type.
// location - Resource location.
// galleryUniqueName - The unique name of the Shared Gallery.
// galleryImageName - The name of the Shared Gallery Image Definition from which the Image Versions are to be listed.
// galleryImageVersionName - The name of the gallery image version to be created. Needs to follow semantic version name pattern:
// The allowed characters are digit and period. Digits must be within the range of a 32-bit integer.
// Format: ..
// options - SharedGalleryImageVersionsGetOptions contains the optional parameters for the SharedGalleryImageVersionsClient.Get
// method.
func (client *SharedGalleryImageVersionsClient) Get(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, galleryImageVersionName string, options *SharedGalleryImageVersionsGetOptions) (SharedGalleryImageVersionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, galleryUniqueName, galleryImageName, galleryImageVersionName, options)
	if err != nil {
		return SharedGalleryImageVersionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharedGalleryImageVersionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharedGalleryImageVersionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharedGalleryImageVersionsClient) getCreateRequest(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, galleryImageVersionName string, options *SharedGalleryImageVersionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if galleryUniqueName == "" {
		return nil, errors.New("parameter galleryUniqueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryUniqueName}", url.PathEscape(galleryUniqueName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedGalleryImageVersionsClient) getHandleResponse(resp *http.Response) (SharedGalleryImageVersionsGetResponse, error) {
	result := SharedGalleryImageVersionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryImageVersion); err != nil {
		return SharedGalleryImageVersionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SharedGalleryImageVersionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List shared gallery image versions by subscription id or tenant id.
// If the operation fails it returns the *CloudError error type.
// location - Resource location.
// galleryUniqueName - The unique name of the Shared Gallery.
// galleryImageName - The name of the Shared Gallery Image Definition from which the Image Versions are to be listed.
// options - SharedGalleryImageVersionsListOptions contains the optional parameters for the SharedGalleryImageVersionsClient.List
// method.
func (client *SharedGalleryImageVersionsClient) List(location string, galleryUniqueName string, galleryImageName string, options *SharedGalleryImageVersionsListOptions) *SharedGalleryImageVersionsListPager {
	return &SharedGalleryImageVersionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, galleryUniqueName, galleryImageName, options)
		},
		advancer: func(ctx context.Context, resp SharedGalleryImageVersionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SharedGalleryImageVersionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SharedGalleryImageVersionsClient) listCreateRequest(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, options *SharedGalleryImageVersionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if galleryUniqueName == "" {
		return nil, errors.New("parameter galleryUniqueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryUniqueName}", url.PathEscape(galleryUniqueName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.SharedTo != nil {
		reqQP.Set("sharedTo", string(*options.SharedTo))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SharedGalleryImageVersionsClient) listHandleResponse(resp *http.Response) (SharedGalleryImageVersionsListResponse, error) {
	result := SharedGalleryImageVersionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryImageVersionList); err != nil {
		return SharedGalleryImageVersionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SharedGalleryImageVersionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
