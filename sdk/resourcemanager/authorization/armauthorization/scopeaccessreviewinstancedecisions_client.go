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

// ScopeAccessReviewInstanceDecisionsClient contains the methods for the ScopeAccessReviewInstanceDecisions group.
// Don't use this type directly, use NewScopeAccessReviewInstanceDecisionsClient() instead.
type ScopeAccessReviewInstanceDecisionsClient struct {
	internal *arm.Client
}

// NewScopeAccessReviewInstanceDecisionsClient creates a new instance of ScopeAccessReviewInstanceDecisionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopeAccessReviewInstanceDecisionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeAccessReviewInstanceDecisionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ScopeAccessReviewInstanceDecisionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopeAccessReviewInstanceDecisionsClient{
	internal: cl,
	}
	return client, nil
}

// NewListPager - Get access review instance decisions
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstanceDecisionsClientListOptions contains the optional parameters for the ScopeAccessReviewInstanceDecisionsClient.NewListPager
//     method.
func (client *ScopeAccessReviewInstanceDecisionsClient) NewListPager(scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceDecisionsClientListOptions) (*runtime.Pager[ScopeAccessReviewInstanceDecisionsClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ScopeAccessReviewInstanceDecisionsClientListResponse]{
		More: func(page ScopeAccessReviewInstanceDecisionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScopeAccessReviewInstanceDecisionsClientListResponse) (ScopeAccessReviewInstanceDecisionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ScopeAccessReviewInstanceDecisionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ScopeAccessReviewInstanceDecisionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ScopeAccessReviewInstanceDecisionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ScopeAccessReviewInstanceDecisionsClient) listCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceDecisionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScopeAccessReviewInstanceDecisionsClient) listHandleResponse(resp *http.Response) (ScopeAccessReviewInstanceDecisionsClientListResponse, error) {
	result := ScopeAccessReviewInstanceDecisionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecisionListResult); err != nil {
		return ScopeAccessReviewInstanceDecisionsClientListResponse{}, err
	}
	return result, nil
}

