//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

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

// RegistrationDefinitionsClient contains the methods for the RegistrationDefinitions group.
// Don't use this type directly, use NewRegistrationDefinitionsClient() instead.
type RegistrationDefinitionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRegistrationDefinitionsClient creates a new instance of RegistrationDefinitionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegistrationDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistrationDefinitionsClient, error) {
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
	client := &RegistrationDefinitionsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a registration definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// registrationDefinitionID - The GUID of the registration definition.
// scope - The scope of the resource.
// requestBody - The parameters required to create a new registration definition.
// options - RegistrationDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistrationDefinitionsClient.BeginCreateOrUpdate
// method.
func (client *RegistrationDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistrationDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, registrationDefinitionID, scope, requestBody, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RegistrationDefinitionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[RegistrationDefinitionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a registration definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *RegistrationDefinitionsClient) createOrUpdate(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, registrationDefinitionID, scope, requestBody, options)
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
func (client *RegistrationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, requestBody)
}

// Delete - Deletes the registration definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// registrationDefinitionID - The GUID of the registration definition.
// scope - The scope of the resource.
// options - RegistrationDefinitionsClientDeleteOptions contains the optional parameters for the RegistrationDefinitionsClient.Delete
// method.
func (client *RegistrationDefinitionsClient) Delete(ctx context.Context, registrationDefinitionID string, scope string, options *RegistrationDefinitionsClientDeleteOptions) (RegistrationDefinitionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, registrationDefinitionID, scope, options)
	if err != nil {
		return RegistrationDefinitionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RegistrationDefinitionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RegistrationDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistrationDefinitionsClient) deleteCreateRequest(ctx context.Context, registrationDefinitionID string, scope string, options *RegistrationDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the registration definition details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// scope - The scope of the resource.
// registrationDefinitionID - The GUID of the registration definition.
// options - RegistrationDefinitionsClientGetOptions contains the optional parameters for the RegistrationDefinitionsClient.Get
// method.
func (client *RegistrationDefinitionsClient) Get(ctx context.Context, scope string, registrationDefinitionID string, options *RegistrationDefinitionsClientGetOptions) (RegistrationDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, registrationDefinitionID, options)
	if err != nil {
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistrationDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistrationDefinitionsClient) getCreateRequest(ctx context.Context, scope string, registrationDefinitionID string, options *RegistrationDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistrationDefinitionsClient) getHandleResponse(resp *http.Response) (RegistrationDefinitionsClientGetResponse, error) {
	result := RegistrationDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationDefinition); err != nil {
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of the registration definitions.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// scope - The scope of the resource.
// options - RegistrationDefinitionsClientListOptions contains the optional parameters for the RegistrationDefinitionsClient.List
// method.
func (client *RegistrationDefinitionsClient) NewListPager(scope string, options *RegistrationDefinitionsClientListOptions) *runtime.Pager[RegistrationDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistrationDefinitionsClientListResponse]{
		More: func(page RegistrationDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistrationDefinitionsClientListResponse) (RegistrationDefinitionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistrationDefinitionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RegistrationDefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistrationDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistrationDefinitionsClient) listCreateRequest(ctx context.Context, scope string, options *RegistrationDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistrationDefinitionsClient) listHandleResponse(resp *http.Response) (RegistrationDefinitionsClientListResponse, error) {
	result := RegistrationDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationDefinitionList); err != nil {
		return RegistrationDefinitionsClientListResponse{}, err
	}
	return result, nil
}