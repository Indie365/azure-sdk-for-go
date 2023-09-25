//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// ScopeAccessReviewHistoryDefinitionInstanceClient contains the methods for the ScopeAccessReviewHistoryDefinitionInstance group.
// Don't use this type directly, use NewScopeAccessReviewHistoryDefinitionInstanceClient() instead.
type ScopeAccessReviewHistoryDefinitionInstanceClient struct {
	internal *arm.Client
}

// NewScopeAccessReviewHistoryDefinitionInstanceClient creates a new instance of ScopeAccessReviewHistoryDefinitionInstanceClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopeAccessReviewHistoryDefinitionInstanceClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeAccessReviewHistoryDefinitionInstanceClient, error) {
	cl, err := arm.NewClient(moduleName+".ScopeAccessReviewHistoryDefinitionInstanceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopeAccessReviewHistoryDefinitionInstanceClient{
	internal: cl,
	}
	return client, nil
}

// GenerateDownloadURI - Generates a uri which can be used to retrieve review history data. This URI has a TTL of 1 day and
// can be retrieved by fetching the accessReviewHistoryDefinition object.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - historyDefinitionID - The id of the access review history definition.
//   - instanceID - The id of the access review history definition instance to generate a URI for.
//   - options - ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIOptions contains the optional parameters for
//     the ScopeAccessReviewHistoryDefinitionInstanceClient.GenerateDownloadURI method.
func (client *ScopeAccessReviewHistoryDefinitionInstanceClient) GenerateDownloadURI(ctx context.Context, scope string, historyDefinitionID string, instanceID string, options *ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIOptions) (ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse, error) {
	var err error
	req, err := client.generateDownloadURICreateRequest(ctx, scope, historyDefinitionID, instanceID, options)
	if err != nil {
		return ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	resp, err := client.generateDownloadURIHandleResponse(httpResp)
	return resp, err
}

// generateDownloadURICreateRequest creates the GenerateDownloadURI request.
func (client *ScopeAccessReviewHistoryDefinitionInstanceClient) generateDownloadURICreateRequest(ctx context.Context, scope string, historyDefinitionID string, instanceID string, options *ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}/instances/{instanceId}/generateDownloadUri"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateDownloadURIHandleResponse handles the GenerateDownloadURI response.
func (client *ScopeAccessReviewHistoryDefinitionInstanceClient) generateDownloadURIHandleResponse(resp *http.Response) (ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse, error) {
	result := ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryInstance); err != nil {
		return ScopeAccessReviewHistoryDefinitionInstanceClientGenerateDownloadURIResponse{}, err
	}
	return result, nil
}

