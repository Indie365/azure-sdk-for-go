//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// APIClient contains the methods for the ManagementGroupsAPI group.
// Don't use this type directly, use NewAPIClient() instead.
type APIClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAPIClient creates a new instance of APIClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*APIClient, error) {
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
	client := &APIClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CheckNameAvailability - Checks if the specified management group name is valid and unique
// If the operation fails it returns an *azcore.ResponseError type.
// checkNameAvailabilityRequest - Management group name availability check parameters.
// options - APIClientCheckNameAvailabilityOptions contains the optional parameters for the APIClient.CheckNameAvailability
// method.
func (client *APIClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *APIClientCheckNameAvailabilityOptions) (APIClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityRequest, options)
	if err != nil {
		return APIClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *APIClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityRequest CheckNameAvailabilityRequest, options *APIClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/checkNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *APIClient) checkNameAvailabilityHandleResponse(resp *http.Response) (APIClientCheckNameAvailabilityResponse, error) {
	result := APIClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return APIClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// StartTenantBackfill - Starts backfilling subscriptions for the Tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - APIClientStartTenantBackfillOptions contains the optional parameters for the APIClient.StartTenantBackfill method.
func (client *APIClient) StartTenantBackfill(ctx context.Context, options *APIClientStartTenantBackfillOptions) (APIClientStartTenantBackfillResponse, error) {
	req, err := client.startTenantBackfillCreateRequest(ctx, options)
	if err != nil {
		return APIClientStartTenantBackfillResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIClientStartTenantBackfillResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIClientStartTenantBackfillResponse{}, runtime.NewResponseError(resp)
	}
	return client.startTenantBackfillHandleResponse(resp)
}

// startTenantBackfillCreateRequest creates the StartTenantBackfill request.
func (client *APIClient) startTenantBackfillCreateRequest(ctx context.Context, options *APIClientStartTenantBackfillOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/startTenantBackfill"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startTenantBackfillHandleResponse handles the StartTenantBackfill response.
func (client *APIClient) startTenantBackfillHandleResponse(resp *http.Response) (APIClientStartTenantBackfillResponse, error) {
	result := APIClientStartTenantBackfillResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantBackfillStatusResult); err != nil {
		return APIClientStartTenantBackfillResponse{}, err
	}
	return result, nil
}

// TenantBackfillStatus - Gets tenant backfill status
// If the operation fails it returns an *azcore.ResponseError type.
// options - APIClientTenantBackfillStatusOptions contains the optional parameters for the APIClient.TenantBackfillStatus
// method.
func (client *APIClient) TenantBackfillStatus(ctx context.Context, options *APIClientTenantBackfillStatusOptions) (APIClientTenantBackfillStatusResponse, error) {
	req, err := client.tenantBackfillStatusCreateRequest(ctx, options)
	if err != nil {
		return APIClientTenantBackfillStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIClientTenantBackfillStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIClientTenantBackfillStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.tenantBackfillStatusHandleResponse(resp)
}

// tenantBackfillStatusCreateRequest creates the TenantBackfillStatus request.
func (client *APIClient) tenantBackfillStatusCreateRequest(ctx context.Context, options *APIClientTenantBackfillStatusOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/tenantBackfillStatus"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// tenantBackfillStatusHandleResponse handles the TenantBackfillStatus response.
func (client *APIClient) tenantBackfillStatusHandleResponse(resp *http.Response) (APIClientTenantBackfillStatusResponse, error) {
	result := APIClientTenantBackfillStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantBackfillStatusResult); err != nil {
		return APIClientTenantBackfillStatusResponse{}, err
	}
	return result, nil
}
