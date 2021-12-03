//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

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

// AssignmentOperationsClient contains the methods for the AssignmentOperations group.
// Don't use this type directly, use NewAssignmentOperationsClient() instead.
type AssignmentOperationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAssignmentOperationsClient creates a new instance of AssignmentOperationsClient with the specified values.
func NewAssignmentOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *AssignmentOperationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AssignmentOperationsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get a blueprint assignment operation.
// If the operation fails it returns the *CloudError error type.
func (client *AssignmentOperationsClient) Get(ctx context.Context, resourceScope string, assignmentName string, assignmentOperationName string, options *AssignmentOperationsGetOptions) (AssignmentOperationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceScope, assignmentName, assignmentOperationName, options)
	if err != nil {
		return AssignmentOperationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssignmentOperationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssignmentOperationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssignmentOperationsClient) getCreateRequest(ctx context.Context, resourceScope string, assignmentName string, assignmentOperationName string, options *AssignmentOperationsGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations/{assignmentOperationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	if assignmentOperationName == "" {
		return nil, errors.New("parameter assignmentOperationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentOperationName}", url.PathEscape(assignmentOperationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssignmentOperationsClient) getHandleResponse(resp *http.Response) (AssignmentOperationsGetResponse, error) {
	result := AssignmentOperationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentOperation); err != nil {
		return AssignmentOperationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AssignmentOperationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List operations for given blueprint assignment within a subscription or a management group.
// If the operation fails it returns the *CloudError error type.
func (client *AssignmentOperationsClient) List(resourceScope string, assignmentName string, options *AssignmentOperationsListOptions) *AssignmentOperationsListPager {
	return &AssignmentOperationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceScope, assignmentName, options)
		},
		advancer: func(ctx context.Context, resp AssignmentOperationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AssignmentOperationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AssignmentOperationsClient) listCreateRequest(ctx context.Context, resourceScope string, assignmentName string, options *AssignmentOperationsListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssignmentOperationsClient) listHandleResponse(resp *http.Response) (AssignmentOperationsListResponse, error) {
	result := AssignmentOperationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentOperationList); err != nil {
		return AssignmentOperationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AssignmentOperationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
