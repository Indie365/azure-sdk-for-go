//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// RestorableSQLResourcesClient contains the methods for the RestorableSQLResources group.
// Don't use this type directly, use NewRestorableSQLResourcesClient() instead.
type RestorableSQLResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableSQLResourcesClient creates a new instance of RestorableSQLResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableSQLResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableSQLResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableSQLResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Return a list of database and container combo that exist on the account at the given timestamp and location.
// This helps in scenarios to validate what resources exist at given timestamp and location.
// This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission.
//
// Generated from API version 2023-03-15-preview
//   - location - Cosmos DB region, with spaces between words and each word capitalized.
//   - instanceID - The instanceId GUID of a restorable database account.
//   - options - RestorableSQLResourcesClientListOptions contains the optional parameters for the RestorableSQLResourcesClient.NewListPager
//     method.
func (client *RestorableSQLResourcesClient) NewListPager(location string, instanceID string, options *RestorableSQLResourcesClientListOptions) *runtime.Pager[RestorableSQLResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableSQLResourcesClientListResponse]{
		More: func(page RestorableSQLResourcesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableSQLResourcesClientListResponse) (RestorableSQLResourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorableSQLResourcesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableSQLResourcesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorableSQLResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableSQLResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RestorableSQLResourcesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableSQLResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	if options != nil && options.RestoreLocation != nil {
		reqQP.Set("restoreLocation", *options.RestoreLocation)
	}
	if options != nil && options.RestoreTimestampInUTC != nil {
		reqQP.Set("restoreTimestampInUtc", *options.RestoreTimestampInUTC)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableSQLResourcesClient) listHandleResponse(resp *http.Response) (RestorableSQLResourcesClientListResponse, error) {
	result := RestorableSQLResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableSQLResourcesListResult); err != nil {
		return RestorableSQLResourcesClientListResponse{}, err
	}
	return result, nil
}
