//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsignalr

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// SignalRListByResourceGroupPager provides operations for iterating over paged responses.
type SignalRListByResourceGroupPager struct {
	client    *SignalRClient
	current   SignalRListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SignalRListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SignalRListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SignalRListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SignalRResourceList.NextLink == nil || len(*p.current.SignalRResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SignalRListByResourceGroupResponse page.
func (p *SignalRListByResourceGroupPager) PageResponse() SignalRListByResourceGroupResponse {
	return p.current
}

// SignalRListBySubscriptionPager provides operations for iterating over paged responses.
type SignalRListBySubscriptionPager struct {
	client    *SignalRClient
	current   SignalRListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SignalRListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SignalRListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SignalRListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SignalRResourceList.NextLink == nil || len(*p.current.SignalRResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SignalRListBySubscriptionResponse page.
func (p *SignalRListBySubscriptionPager) PageResponse() SignalRListBySubscriptionResponse {
	return p.current
}

// SignalRPrivateEndpointConnectionsListPager provides operations for iterating over paged responses.
type SignalRPrivateEndpointConnectionsListPager struct {
	client    *SignalRPrivateEndpointConnectionsClient
	current   SignalRPrivateEndpointConnectionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SignalRPrivateEndpointConnectionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SignalRPrivateEndpointConnectionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SignalRPrivateEndpointConnectionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionList.NextLink == nil || len(*p.current.PrivateEndpointConnectionList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SignalRPrivateEndpointConnectionsListResponse page.
func (p *SignalRPrivateEndpointConnectionsListPager) PageResponse() SignalRPrivateEndpointConnectionsListResponse {
	return p.current
}

// SignalRPrivateLinkResourcesListPager provides operations for iterating over paged responses.
type SignalRPrivateLinkResourcesListPager struct {
	client    *SignalRPrivateLinkResourcesClient
	current   SignalRPrivateLinkResourcesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SignalRPrivateLinkResourcesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SignalRPrivateLinkResourcesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SignalRPrivateLinkResourcesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkResourceList.NextLink == nil || len(*p.current.PrivateLinkResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SignalRPrivateLinkResourcesListResponse page.
func (p *SignalRPrivateLinkResourcesListPager) PageResponse() SignalRPrivateLinkResourcesListResponse {
	return p.current
}

// SignalRSharedPrivateLinkResourcesListPager provides operations for iterating over paged responses.
type SignalRSharedPrivateLinkResourcesListPager struct {
	client    *SignalRSharedPrivateLinkResourcesClient
	current   SignalRSharedPrivateLinkResourcesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SignalRSharedPrivateLinkResourcesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SignalRSharedPrivateLinkResourcesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SignalRSharedPrivateLinkResourcesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SharedPrivateLinkResourceList.NextLink == nil || len(*p.current.SharedPrivateLinkResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SignalRSharedPrivateLinkResourcesListResponse page.
func (p *SignalRSharedPrivateLinkResourcesListPager) PageResponse() SignalRSharedPrivateLinkResourcesListResponse {
	return p.current
}

// UsagesListPager provides operations for iterating over paged responses.
type UsagesListPager struct {
	client    *UsagesClient
	current   UsagesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, UsagesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *UsagesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *UsagesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SignalRUsageList.NextLink == nil || len(*p.current.SignalRUsageList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current UsagesListResponse page.
func (p *UsagesListPager) PageResponse() UsagesListResponse {
	return p.current
}
