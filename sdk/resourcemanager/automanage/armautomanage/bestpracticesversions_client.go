//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BestPracticesVersionsClient contains the methods for the BestPracticesVersions group.
// Don't use this type directly, use NewBestPracticesVersionsClient() instead.
type BestPracticesVersionsClient struct {
	internal *arm.Client
}

// NewBestPracticesVersionsClient creates a new instance of BestPracticesVersionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBestPracticesVersionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BestPracticesVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".BestPracticesVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BestPracticesVersionsClient{
	internal: cl,
	}
	return client, nil
}

// Get - Get information about a Automanage best practice version
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - bestPracticeName - The Automanage best practice name.
//   - versionName - The Automanage best practice version name.
//   - options - BestPracticesVersionsClientGetOptions contains the optional parameters for the BestPracticesVersionsClient.Get
//     method.
func (client *BestPracticesVersionsClient) Get(ctx context.Context, bestPracticeName string, versionName string, options *BestPracticesVersionsClientGetOptions) (BestPracticesVersionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, bestPracticeName, versionName, options)
	if err != nil {
		return BestPracticesVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BestPracticesVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BestPracticesVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BestPracticesVersionsClient) getCreateRequest(ctx context.Context, bestPracticeName string, versionName string, options *BestPracticesVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Automanage/bestPractices/{bestPracticeName}/versions/{versionName}"
	if bestPracticeName == "" {
		return nil, errors.New("parameter bestPracticeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bestPracticeName}", url.PathEscape(bestPracticeName))
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BestPracticesVersionsClient) getHandleResponse(resp *http.Response) (BestPracticesVersionsClientGetResponse, error) {
	result := BestPracticesVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BestPractice); err != nil {
		return BestPracticesVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTenantPager - Retrieve a list of Automanage best practices versions
//
// Generated from API version 2022-05-04
//   - bestPracticeName - The Automanage best practice name.
//   - options - BestPracticesVersionsClientListByTenantOptions contains the optional parameters for the BestPracticesVersionsClient.NewListByTenantPager
//     method.
func (client *BestPracticesVersionsClient) NewListByTenantPager(bestPracticeName string, options *BestPracticesVersionsClientListByTenantOptions) (*runtime.Pager[BestPracticesVersionsClientListByTenantResponse]) {
	return runtime.NewPager(runtime.PagingHandler[BestPracticesVersionsClientListByTenantResponse]{
		More: func(page BestPracticesVersionsClientListByTenantResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BestPracticesVersionsClientListByTenantResponse) (BestPracticesVersionsClientListByTenantResponse, error) {
			req, err := client.listByTenantCreateRequest(ctx, bestPracticeName, options)
			if err != nil {
				return BestPracticesVersionsClientListByTenantResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BestPracticesVersionsClientListByTenantResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BestPracticesVersionsClientListByTenantResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTenantHandleResponse(resp)
		},
	})
}

// listByTenantCreateRequest creates the ListByTenant request.
func (client *BestPracticesVersionsClient) listByTenantCreateRequest(ctx context.Context, bestPracticeName string, options *BestPracticesVersionsClientListByTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Automanage/bestPractices/{bestPracticeName}/versions"
	if bestPracticeName == "" {
		return nil, errors.New("parameter bestPracticeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bestPracticeName}", url.PathEscape(bestPracticeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTenantHandleResponse handles the ListByTenant response.
func (client *BestPracticesVersionsClient) listByTenantHandleResponse(resp *http.Response) (BestPracticesVersionsClientListByTenantResponse, error) {
	result := BestPracticesVersionsClientListByTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BestPracticeList); err != nil {
		return BestPracticesVersionsClientListByTenantResponse{}, err
	}
	return result, nil
}

