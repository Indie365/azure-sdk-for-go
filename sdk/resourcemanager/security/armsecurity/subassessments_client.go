//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// SubAssessmentsClient contains the methods for the SubAssessments group.
// Don't use this type directly, use NewSubAssessmentsClient() instead.
type SubAssessmentsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewSubAssessmentsClient creates a new instance of SubAssessmentsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubAssessmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SubAssessmentsClient, error) {
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
	client := &SubAssessmentsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get a security sub-assessment on your scanned resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - The Assessment Key - Unique key for the assessment type
// subAssessmentName - The Sub-Assessment Key - Unique key for the sub-assessment type
// options - SubAssessmentsClientGetOptions contains the optional parameters for the SubAssessmentsClient.Get method.
func (client *SubAssessmentsClient) Get(ctx context.Context, scope string, assessmentName string, subAssessmentName string, options *SubAssessmentsClientGetOptions) (SubAssessmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, assessmentName, subAssessmentName, options)
	if err != nil {
		return SubAssessmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubAssessmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubAssessmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubAssessmentsClient) getCreateRequest(ctx context.Context, scope string, assessmentName string, subAssessmentName string, options *SubAssessmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments/{subAssessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if subAssessmentName == "" {
		return nil, errors.New("parameter subAssessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subAssessmentName}", url.PathEscape(subAssessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubAssessmentsClient) getHandleResponse(resp *http.Response) (SubAssessmentsClientGetResponse, error) {
	result := SubAssessmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubAssessment); err != nil {
		return SubAssessmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get security sub-assessments on all your scanned resources inside a scope
// Generated from API version 2019-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// assessmentName - The Assessment Key - Unique key for the assessment type
// options - SubAssessmentsClientListOptions contains the optional parameters for the SubAssessmentsClient.List method.
func (client *SubAssessmentsClient) NewListPager(scope string, assessmentName string, options *SubAssessmentsClientListOptions) *runtime.Pager[SubAssessmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubAssessmentsClientListResponse]{
		More: func(page SubAssessmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubAssessmentsClientListResponse) (SubAssessmentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, assessmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubAssessmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubAssessmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubAssessmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SubAssessmentsClient) listCreateRequest(ctx context.Context, scope string, assessmentName string, options *SubAssessmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubAssessmentsClient) listHandleResponse(resp *http.Response) (SubAssessmentsClientListResponse, error) {
	result := SubAssessmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubAssessmentList); err != nil {
		return SubAssessmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Get security sub-assessments on all your scanned resources inside a subscription scope
// Generated from API version 2019-01-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// options - SubAssessmentsClientListAllOptions contains the optional parameters for the SubAssessmentsClient.ListAll method.
func (client *SubAssessmentsClient) NewListAllPager(scope string, options *SubAssessmentsClientListAllOptions) *runtime.Pager[SubAssessmentsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubAssessmentsClientListAllResponse]{
		More: func(page SubAssessmentsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubAssessmentsClientListAllResponse) (SubAssessmentsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubAssessmentsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubAssessmentsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubAssessmentsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *SubAssessmentsClient) listAllCreateRequest(ctx context.Context, scope string, options *SubAssessmentsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/subAssessments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *SubAssessmentsClient) listAllHandleResponse(resp *http.Response) (SubAssessmentsClientListAllResponse, error) {
	result := SubAssessmentsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubAssessmentList); err != nil {
		return SubAssessmentsClientListAllResponse{}, err
	}
	return result, nil
}