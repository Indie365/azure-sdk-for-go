//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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

// ConsumerInvitationsClient contains the methods for the ConsumerInvitations group.
// Don't use this type directly, use NewConsumerInvitationsClient() instead.
type ConsumerInvitationsClient struct {
	internal *arm.Client
}

// NewConsumerInvitationsClient creates a new instance of ConsumerInvitationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConsumerInvitationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ConsumerInvitationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConsumerInvitationsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get an invitation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - location - Location of the invitation
//   - invitationID - An invitation id
//   - options - ConsumerInvitationsClientGetOptions contains the optional parameters for the ConsumerInvitationsClient.Get method.
func (client *ConsumerInvitationsClient) Get(ctx context.Context, location string, invitationID string, options *ConsumerInvitationsClientGetOptions) (ConsumerInvitationsClientGetResponse, error) {
	var err error
	const operationName = "ConsumerInvitationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, invitationID, options)
	if err != nil {
		return ConsumerInvitationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConsumerInvitationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConsumerInvitationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConsumerInvitationsClient) getCreateRequest(ctx context.Context, location string, invitationID string, options *ConsumerInvitationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DataShare/locations/{location}/consumerInvitations/{invitationId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if invitationID == "" {
		return nil, errors.New("parameter invitationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invitationId}", url.PathEscape(invitationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConsumerInvitationsClient) getHandleResponse(resp *http.Response) (ConsumerInvitationsClientGetResponse, error) {
	result := ConsumerInvitationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerInvitation); err != nil {
		return ConsumerInvitationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListInvitationsPager - Lists invitations
//
// Generated from API version 2020-09-01
//   - options - ConsumerInvitationsClientListInvitationsOptions contains the optional parameters for the ConsumerInvitationsClient.NewListInvitationsPager
//     method.
func (client *ConsumerInvitationsClient) NewListInvitationsPager(options *ConsumerInvitationsClientListInvitationsOptions) *runtime.Pager[ConsumerInvitationsClientListInvitationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConsumerInvitationsClientListInvitationsResponse]{
		More: func(page ConsumerInvitationsClientListInvitationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConsumerInvitationsClientListInvitationsResponse) (ConsumerInvitationsClientListInvitationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConsumerInvitationsClient.NewListInvitationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listInvitationsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ConsumerInvitationsClientListInvitationsResponse{}, err
			}
			return client.listInvitationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listInvitationsCreateRequest creates the ListInvitations request.
func (client *ConsumerInvitationsClient) listInvitationsCreateRequest(ctx context.Context, options *ConsumerInvitationsClientListInvitationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DataShare/listInvitations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listInvitationsHandleResponse handles the ListInvitations response.
func (client *ConsumerInvitationsClient) listInvitationsHandleResponse(resp *http.Response) (ConsumerInvitationsClientListInvitationsResponse, error) {
	result := ConsumerInvitationsClientListInvitationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerInvitationList); err != nil {
		return ConsumerInvitationsClientListInvitationsResponse{}, err
	}
	return result, nil
}

// RejectInvitation - Reject an invitation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - location - Location of the invitation
//   - invitation - An invitation payload
//   - options - ConsumerInvitationsClientRejectInvitationOptions contains the optional parameters for the ConsumerInvitationsClient.RejectInvitation
//     method.
func (client *ConsumerInvitationsClient) RejectInvitation(ctx context.Context, location string, invitation ConsumerInvitation, options *ConsumerInvitationsClientRejectInvitationOptions) (ConsumerInvitationsClientRejectInvitationResponse, error) {
	var err error
	const operationName = "ConsumerInvitationsClient.RejectInvitation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rejectInvitationCreateRequest(ctx, location, invitation, options)
	if err != nil {
		return ConsumerInvitationsClientRejectInvitationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConsumerInvitationsClientRejectInvitationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConsumerInvitationsClientRejectInvitationResponse{}, err
	}
	resp, err := client.rejectInvitationHandleResponse(httpResp)
	return resp, err
}

// rejectInvitationCreateRequest creates the RejectInvitation request.
func (client *ConsumerInvitationsClient) rejectInvitationCreateRequest(ctx context.Context, location string, invitation ConsumerInvitation, options *ConsumerInvitationsClientRejectInvitationOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DataShare/locations/{location}/rejectInvitation"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, invitation); err != nil {
		return nil, err
	}
	return req, nil
}

// rejectInvitationHandleResponse handles the RejectInvitation response.
func (client *ConsumerInvitationsClient) rejectInvitationHandleResponse(resp *http.Response) (ConsumerInvitationsClientRejectInvitationResponse, error) {
	result := ConsumerInvitationsClientRejectInvitationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerInvitation); err != nil {
		return ConsumerInvitationsClientRejectInvitationResponse{}, err
	}
	return result, nil
}
