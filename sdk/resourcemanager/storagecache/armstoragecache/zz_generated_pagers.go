//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// CachesListByResourceGroupPager provides operations for iterating over paged responses.
type CachesListByResourceGroupPager struct {
	client    *CachesClient
	current   CachesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CachesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CachesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CachesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CachesListResult.NextLink == nil || len(*p.current.CachesListResult.NextLink) == 0 {
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

// PageResponse returns the current CachesListByResourceGroupResponse page.
func (p *CachesListByResourceGroupPager) PageResponse() CachesListByResourceGroupResponse {
	return p.current
}

// CachesListPager provides operations for iterating over paged responses.
type CachesListPager struct {
	client    *CachesClient
	current   CachesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CachesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CachesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CachesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CachesListResult.NextLink == nil || len(*p.current.CachesListResult.NextLink) == 0 {
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

// PageResponse returns the current CachesListResponse page.
func (p *CachesListPager) PageResponse() CachesListResponse {
	return p.current
}

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
		if p.current.APIOperationListResult.NextLink == nil || len(*p.current.APIOperationListResult.NextLink) == 0 {
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

// SKUsListPager provides operations for iterating over paged responses.
type SKUsListPager struct {
	client    *SKUsClient
	current   SKUsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SKUsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SKUsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SKUsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceSKUsResult.NextLink == nil || len(*p.current.ResourceSKUsResult.NextLink) == 0 {
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

// PageResponse returns the current SKUsListResponse page.
func (p *SKUsListPager) PageResponse() SKUsListResponse {
	return p.current
}

// StorageTargetsListByCachePager provides operations for iterating over paged responses.
type StorageTargetsListByCachePager struct {
	client    *StorageTargetsClient
	current   StorageTargetsListByCacheResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, StorageTargetsListByCacheResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *StorageTargetsListByCachePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *StorageTargetsListByCachePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageTargetsResult.NextLink == nil || len(*p.current.StorageTargetsResult.NextLink) == 0 {
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
		p.err = p.client.listByCacheHandleError(resp)
		return false
	}
	result, err := p.client.listByCacheHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current StorageTargetsListByCacheResponse page.
func (p *StorageTargetsListByCachePager) PageResponse() StorageTargetsListByCacheResponse {
	return p.current
}

// UsageModelsListPager provides operations for iterating over paged responses.
type UsageModelsListPager struct {
	client    *UsageModelsClient
	current   UsageModelsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, UsageModelsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *UsageModelsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *UsageModelsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.UsageModelsResult.NextLink == nil || len(*p.current.UsageModelsResult.NextLink) == 0 {
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

// PageResponse returns the current UsageModelsListResponse page.
func (p *UsageModelsListPager) PageResponse() UsageModelsListResponse {
	return p.current
}
