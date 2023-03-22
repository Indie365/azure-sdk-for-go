//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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
	"strconv"
	"strings"
)

// VirtualMachineExtensionImagesClient contains the methods for the VirtualMachineExtensionImages group.
// Don't use this type directly, use NewVirtualMachineExtensionImagesClient() instead.
type VirtualMachineExtensionImagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineExtensionImagesClient creates a new instance of VirtualMachineExtensionImagesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineExtensionImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineExtensionImagesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armcompute.VirtualMachineExtensionImagesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineExtensionImagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a virtual machine extension image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - location - The name of a supported Azure region.
//   - options - VirtualMachineExtensionImagesClientGetOptions contains the optional parameters for the VirtualMachineExtensionImagesClient.Get
//     method.
func (client *VirtualMachineExtensionImagesClient) Get(ctx context.Context, location string, publisherName string, typeParam string, version string, options *VirtualMachineExtensionImagesClientGetOptions) (VirtualMachineExtensionImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publisherName, typeParam, version, options)
	if err != nil {
		return VirtualMachineExtensionImagesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineExtensionImagesClient) getCreateRequest(ctx context.Context, location string, publisherName string, typeParam string, version string, options *VirtualMachineExtensionImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if typeParam == "" {
		return nil, errors.New("parameter typeParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{type}", url.PathEscape(typeParam))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineExtensionImagesClient) getHandleResponse(resp *http.Response) (VirtualMachineExtensionImagesClientGetResponse, error) {
	result := VirtualMachineExtensionImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionImage); err != nil {
		return VirtualMachineExtensionImagesClientGetResponse{}, err
	}
	return result, nil
}

// ListTypes - Gets a list of virtual machine extension image types.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - location - The name of a supported Azure region.
//   - options - VirtualMachineExtensionImagesClientListTypesOptions contains the optional parameters for the VirtualMachineExtensionImagesClient.ListTypes
//     method.
func (client *VirtualMachineExtensionImagesClient) ListTypes(ctx context.Context, location string, publisherName string, options *VirtualMachineExtensionImagesClientListTypesOptions) (VirtualMachineExtensionImagesClientListTypesResponse, error) {
	req, err := client.listTypesCreateRequest(ctx, location, publisherName, options)
	if err != nil {
		return VirtualMachineExtensionImagesClientListTypesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionImagesClientListTypesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionImagesClientListTypesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listTypesHandleResponse(resp)
}

// listTypesCreateRequest creates the ListTypes request.
func (client *VirtualMachineExtensionImagesClient) listTypesCreateRequest(ctx context.Context, location string, publisherName string, options *VirtualMachineExtensionImagesClientListTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listTypesHandleResponse handles the ListTypes response.
func (client *VirtualMachineExtensionImagesClient) listTypesHandleResponse(resp *http.Response) (VirtualMachineExtensionImagesClientListTypesResponse, error) {
	result := VirtualMachineExtensionImagesClientListTypesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionImageArray); err != nil {
		return VirtualMachineExtensionImagesClientListTypesResponse{}, err
	}
	return result, nil
}

// ListVersions - Gets a list of virtual machine extension image versions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - location - The name of a supported Azure region.
//   - options - VirtualMachineExtensionImagesClientListVersionsOptions contains the optional parameters for the VirtualMachineExtensionImagesClient.ListVersions
//     method.
func (client *VirtualMachineExtensionImagesClient) ListVersions(ctx context.Context, location string, publisherName string, typeParam string, options *VirtualMachineExtensionImagesClientListVersionsOptions) (VirtualMachineExtensionImagesClientListVersionsResponse, error) {
	req, err := client.listVersionsCreateRequest(ctx, location, publisherName, typeParam, options)
	if err != nil {
		return VirtualMachineExtensionImagesClientListVersionsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionImagesClientListVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionImagesClientListVersionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listVersionsHandleResponse(resp)
}

// listVersionsCreateRequest creates the ListVersions request.
func (client *VirtualMachineExtensionImagesClient) listVersionsCreateRequest(ctx context.Context, location string, publisherName string, typeParam string, options *VirtualMachineExtensionImagesClientListVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if typeParam == "" {
		return nil, errors.New("parameter typeParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{type}", url.PathEscape(typeParam))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listVersionsHandleResponse handles the ListVersions response.
func (client *VirtualMachineExtensionImagesClient) listVersionsHandleResponse(resp *http.Response) (VirtualMachineExtensionImagesClientListVersionsResponse, error) {
	result := VirtualMachineExtensionImagesClientListVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionImageArray); err != nil {
		return VirtualMachineExtensionImagesClientListVersionsResponse{}, err
	}
	return result, nil
}
