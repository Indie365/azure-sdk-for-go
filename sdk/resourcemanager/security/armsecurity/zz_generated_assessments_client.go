//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// AssessmentsClient contains the methods for the Assessments group.
// Don't use this type directly, use NewAssessmentsClient() instead.
type AssessmentsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAssessmentsClient creates a new instance of AssessmentsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAssessmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessmentsClient, error) {
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
	client := &AssessmentsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CreateOrUpdate - Create a security assessment on your resource. An assessment metadata that describes this assessment must
// be predefined with the same name before inserting the assessment result
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The identifier of the resource.
// assessmentName - The Assessment Key - Unique key for the assessment type
// assessment - Calculated assessment on a pre-defined assessment metadata
// options - AssessmentsClientCreateOrUpdateOptions contains the optional parameters for the AssessmentsClient.CreateOrUpdate
// method.
func (client *AssessmentsClient) CreateOrUpdate(ctx context.Context, resourceID string, assessmentName string, assessment Assessment, options *AssessmentsClientCreateOrUpdateOptions) (AssessmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceID, assessmentName, assessment, options)
	if err != nil {
		return AssessmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AssessmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssessmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceID string, assessmentName string, assessment Assessment, options *AssessmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, assessment)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AssessmentsClient) createOrUpdateHandleResponse(resp *http.Response) (AssessmentsClientCreateOrUpdateResponse, error) {
	result := AssessmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentResponse); err != nil {
		return AssessmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a security assessment on your resource. An assessment metadata that describes this assessment must be predefined
// with the same name before inserting the assessment result
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The identifier of the resource.
// assessmentName - The Assessment Key - Unique key for the assessment type
// options - AssessmentsClientDeleteOptions contains the optional parameters for the AssessmentsClient.Delete method.
func (client *AssessmentsClient) Delete(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsClientDeleteOptions) (AssessmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceID, assessmentName, options)
	if err != nil {
		return AssessmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AssessmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AssessmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssessmentsClient) deleteCreateRequest(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a security assessment on your scanned resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceID - The identifier of the resource.
// assessmentName - The Assessment Key - Unique key for the assessment type
// options - AssessmentsClientGetOptions contains the optional parameters for the AssessmentsClient.Get method.
func (client *AssessmentsClient) Get(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsClientGetOptions) (AssessmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceID, assessmentName, options)
	if err != nil {
		return AssessmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssessmentsClient) getCreateRequest(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssessmentsClient) getHandleResponse(resp *http.Response) (AssessmentsClientGetResponse, error) {
	result := AssessmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentResponse); err != nil {
		return AssessmentsClientGetResponse{}, err
	}
	return result, nil
}

// List - Get security assessments on all your scanned resources inside a scope
// If the operation fails it returns an *azcore.ResponseError type.
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// options - AssessmentsClientListOptions contains the optional parameters for the AssessmentsClient.List method.
func (client *AssessmentsClient) List(scope string, options *AssessmentsClientListOptions) *runtime.Pager[AssessmentsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AssessmentsClientListResponse]{
		More: func(page AssessmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessmentsClientListResponse) (AssessmentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AssessmentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AssessmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssessmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AssessmentsClient) listCreateRequest(ctx context.Context, scope string, options *AssessmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssessmentsClient) listHandleResponse(resp *http.Response) (AssessmentsClientListResponse, error) {
	result := AssessmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentList); err != nil {
		return AssessmentsClientListResponse{}, err
	}
	return result, nil
}
