//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ApplicationsListPager provides operations for iterating over paged responses.
type ApplicationsListPager struct {
	client    *ApplicationsClient
	current   ApplicationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ApplicationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ApplicationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ApplicationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SaasAppResponseWithContinuation.NextLink == nil || len(*p.current.SaasAppResponseWithContinuation.NextLink) == 0 {
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

// PageResponse returns the current ApplicationsListResponse page.
func (p *ApplicationsListPager) PageResponse() ApplicationsListResponse {
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
		if p.current.SaasAppOperationsResponseWithContinuation.NextLink == nil || len(*p.current.SaasAppOperationsResponseWithContinuation.NextLink) == 0 {
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

// SaasResourcesListPager provides operations for iterating over paged responses.
type SaasResourcesListPager struct {
	client    *SaasResourcesClient
	current   SaasResourcesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SaasResourcesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SaasResourcesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SaasResourcesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SaasResourceResponseWithContinuation.NextLink == nil || len(*p.current.SaasResourceResponseWithContinuation.NextLink) == 0 {
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

// PageResponse returns the current SaasResourcesListResponse page.
func (p *SaasResourcesListPager) PageResponse() SaasResourcesListResponse {
	return p.current
}

// SaasSubscriptionLevelListByAzureSubscriptionPager provides operations for iterating over paged responses.
type SaasSubscriptionLevelListByAzureSubscriptionPager struct {
	client    *SaasSubscriptionLevelClient
	current   SaasSubscriptionLevelListByAzureSubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SaasSubscriptionLevelListByAzureSubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SaasSubscriptionLevelListByAzureSubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SaasSubscriptionLevelListByAzureSubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SaasResourceResponseWithContinuation.NextLink == nil || len(*p.current.SaasResourceResponseWithContinuation.NextLink) == 0 {
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
		p.err = p.client.listByAzureSubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listByAzureSubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SaasSubscriptionLevelListByAzureSubscriptionResponse page.
func (p *SaasSubscriptionLevelListByAzureSubscriptionPager) PageResponse() SaasSubscriptionLevelListByAzureSubscriptionResponse {
	return p.current
}

// SaasSubscriptionLevelListByResourceGroupPager provides operations for iterating over paged responses.
type SaasSubscriptionLevelListByResourceGroupPager struct {
	client    *SaasSubscriptionLevelClient
	current   SaasSubscriptionLevelListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SaasSubscriptionLevelListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SaasSubscriptionLevelListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SaasSubscriptionLevelListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SaasResourceResponseWithContinuation.NextLink == nil || len(*p.current.SaasResourceResponseWithContinuation.NextLink) == 0 {
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

// PageResponse returns the current SaasSubscriptionLevelListByResourceGroupResponse page.
func (p *SaasSubscriptionLevelListByResourceGroupPager) PageResponse() SaasSubscriptionLevelListByResourceGroupResponse {
	return p.current
}
