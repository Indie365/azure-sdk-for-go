//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// AssessmentsClient contains the methods for the Assessments group.
// Don't use this type directly, use NewAssessmentsClient() instead.
type AssessmentsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAssessmentsClient creates a new instance of AssessmentsClient with the specified values.
func NewAssessmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AssessmentsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AssessmentsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create a security assessment on your resource. An assessment metadata that describes this assessment must be predefined with the same
// name before inserting the assessment result
// If the operation fails it returns the *CloudError error type.
func (client *AssessmentsClient) CreateOrUpdate(ctx context.Context, resourceID string, assessmentName string, assessment SecurityAssessment, options *AssessmentsCreateOrUpdateOptions) (AssessmentsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceID, assessmentName, assessment, options)
	if err != nil {
		return AssessmentsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AssessmentsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssessmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceID string, assessmentName string, assessment SecurityAssessment, options *AssessmentsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
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
func (client *AssessmentsClient) createOrUpdateHandleResponse(resp *http.Response) (AssessmentsCreateOrUpdateResponse, error) {
	result := AssessmentsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityAssessmentResponse); err != nil {
		return AssessmentsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AssessmentsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete a security assessment on your resource. An assessment metadata that describes this assessment must be predefined with the same name before
// inserting the assessment result
// If the operation fails it returns the *CloudError error type.
func (client *AssessmentsClient) Delete(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsDeleteOptions) (AssessmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceID, assessmentName, options)
	if err != nil {
		return AssessmentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AssessmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return AssessmentsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssessmentsClient) deleteCreateRequest(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AssessmentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a security assessment on your scanned resource
// If the operation fails it returns the *CloudError error type.
func (client *AssessmentsClient) Get(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsGetOptions) (AssessmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceID, assessmentName, options)
	if err != nil {
		return AssessmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssessmentsClient) getCreateRequest(ctx context.Context, resourceID string, assessmentName string, options *AssessmentsGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *AssessmentsClient) getHandleResponse(resp *http.Response) (AssessmentsGetResponse, error) {
	result := AssessmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityAssessmentResponse); err != nil {
		return AssessmentsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AssessmentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get security assessments on all your scanned resources inside a scope
// If the operation fails it returns the *CloudError error type.
func (client *AssessmentsClient) List(scope string, options *AssessmentsListOptions) *AssessmentsListPager {
	return &AssessmentsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp AssessmentsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityAssessmentList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AssessmentsClient) listCreateRequest(ctx context.Context, scope string, options *AssessmentsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *AssessmentsClient) listHandleResponse(resp *http.Response) (AssessmentsListResponse, error) {
	result := AssessmentsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityAssessmentList); err != nil {
		return AssessmentsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AssessmentsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
