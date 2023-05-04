//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armportal

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

// TenantConfigurationsClient contains the methods for the TenantConfigurations group.
// Don't use this type directly, use NewTenantConfigurationsClient() instead.
type TenantConfigurationsClient struct {
	internal *arm.Client
}

// NewTenantConfigurationsClient creates a new instance of TenantConfigurationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTenantConfigurationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".TenantConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TenantConfigurationsClient{
		internal: cl,
	}
	return client, nil
}

// Create - Create the tenant configuration. If configuration already exists - update it. User has to be a Tenant Admin for
// this operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01-preview
//   - configurationName - The configuration name. Value must be 'default'
//   - tenantConfiguration - The parameters required to create or update tenant configuration.
//   - options - TenantConfigurationsClientCreateOptions contains the optional parameters for the TenantConfigurationsClient.Create
//     method.
func (client *TenantConfigurationsClient) Create(ctx context.Context, configurationName ConfigurationName, tenantConfiguration Configuration, options *TenantConfigurationsClientCreateOptions) (TenantConfigurationsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, configurationName, tenantConfiguration, options)
	if err != nil {
		return TenantConfigurationsClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantConfigurationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return TenantConfigurationsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *TenantConfigurationsClient) createCreateRequest(ctx context.Context, configurationName ConfigurationName, tenantConfiguration Configuration, options *TenantConfigurationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations/{configurationName}"
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, tenantConfiguration)
}

// createHandleResponse handles the Create response.
func (client *TenantConfigurationsClient) createHandleResponse(resp *http.Response) (TenantConfigurationsClientCreateResponse, error) {
	result := TenantConfigurationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return TenantConfigurationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the tenant configuration. User has to be a Tenant Admin for this operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01-preview
//   - configurationName - The configuration name. Value must be 'default'
//   - options - TenantConfigurationsClientDeleteOptions contains the optional parameters for the TenantConfigurationsClient.Delete
//     method.
func (client *TenantConfigurationsClient) Delete(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsClientDeleteOptions) (TenantConfigurationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, configurationName, options)
	if err != nil {
		return TenantConfigurationsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TenantConfigurationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return TenantConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TenantConfigurationsClient) deleteCreateRequest(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations/{configurationName}"
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the tenant configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01-preview
//   - configurationName - The configuration name. Value must be 'default'
//   - options - TenantConfigurationsClientGetOptions contains the optional parameters for the TenantConfigurationsClient.Get
//     method.
func (client *TenantConfigurationsClient) Get(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsClientGetOptions) (TenantConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, configurationName, options)
	if err != nil {
		return TenantConfigurationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return TenantConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TenantConfigurationsClient) getCreateRequest(ctx context.Context, configurationName ConfigurationName, options *TenantConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations/{configurationName}"
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(string(configurationName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TenantConfigurationsClient) getHandleResponse(resp *http.Response) (TenantConfigurationsClientGetResponse, error) {
	result := TenantConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configuration); err != nil {
		return TenantConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets list of the tenant configurations.
//
// Generated from API version 2020-09-01-preview
//   - options - TenantConfigurationsClientListOptions contains the optional parameters for the TenantConfigurationsClient.NewListPager
//     method.
func (client *TenantConfigurationsClient) NewListPager(options *TenantConfigurationsClientListOptions) *runtime.Pager[TenantConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TenantConfigurationsClientListResponse]{
		More: func(page TenantConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantConfigurationsClientListResponse) (TenantConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TenantConfigurationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TenantConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TenantConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TenantConfigurationsClient) listCreateRequest(ctx context.Context, options *TenantConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/tenantConfigurations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TenantConfigurationsClient) listHandleResponse(resp *http.Response) (TenantConfigurationsClientListResponse, error) {
	result := TenantConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationList); err != nil {
		return TenantConfigurationsClientListResponse{}, err
	}
	return result, nil
}
