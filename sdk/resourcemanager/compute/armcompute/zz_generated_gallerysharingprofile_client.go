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

// GallerySharingProfileClient contains the methods for the GallerySharingProfile group.
// Don't use this type directly, use NewGallerySharingProfileClient() instead.
type GallerySharingProfileClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGallerySharingProfileClient creates a new instance of GallerySharingProfileClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGallerySharingProfileClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GallerySharingProfileClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &GallerySharingProfileClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginUpdate - Update sharing profile of a gallery.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery.
// sharingUpdate - Parameters supplied to the update gallery sharing profile.
// options - GallerySharingProfileBeginUpdateOptions contains the optional parameters for the GallerySharingProfileClient.BeginUpdate
// method.
func (client *GallerySharingProfileClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileBeginUpdateOptions) (GallerySharingProfileUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, galleryName, sharingUpdate, options)
	if err != nil {
		return GallerySharingProfileUpdatePollerResponse{}, err
	}
	result := GallerySharingProfileUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GallerySharingProfileClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return GallerySharingProfileUpdatePollerResponse{}, err
	}
	result.Poller = &GallerySharingProfileUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update sharing profile of a gallery.
// If the operation fails it returns the *CloudError error type.
func (client *GallerySharingProfileClient) update(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, sharingUpdate, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GallerySharingProfileClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/share"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sharingUpdate)
}

// updateHandleError handles the Update error response.
func (client *GallerySharingProfileClient) updateHandleError(resp *http.Response) error {
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
