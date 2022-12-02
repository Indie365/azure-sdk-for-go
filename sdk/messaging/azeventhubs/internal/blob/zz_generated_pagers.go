//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package blob

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// containerClientListBlobFlatSegmentPager provides operations for iterating over paged responses.
type containerClientListBlobFlatSegmentPager struct {
	client    *containerClient
	current   containerClientListBlobFlatSegmentResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, containerClientListBlobFlatSegmentResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *containerClientListBlobFlatSegmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *containerClientListBlobFlatSegmentPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListBlobsFlatSegmentResponse.NextMarker == nil || len(*p.current.ListBlobsFlatSegmentResponse.NextMarker) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listBlobFlatSegmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current containerClientListBlobFlatSegmentResponse page.
func (p *containerClientListBlobFlatSegmentPager) PageResponse() containerClientListBlobFlatSegmentResponse {
	return p.current
}

// containerClientListBlobHierarchySegmentPager provides operations for iterating over paged responses.
type containerClientListBlobHierarchySegmentPager struct {
	client    *containerClient
	current   containerClientListBlobHierarchySegmentResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, containerClientListBlobHierarchySegmentResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *containerClientListBlobHierarchySegmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *containerClientListBlobHierarchySegmentPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListBlobsHierarchySegmentResponse.NextMarker == nil || len(*p.current.ListBlobsHierarchySegmentResponse.NextMarker) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listBlobHierarchySegmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current containerClientListBlobHierarchySegmentResponse page.
func (p *containerClientListBlobHierarchySegmentPager) PageResponse() containerClientListBlobHierarchySegmentResponse {
	return p.current
}

// pageBlobClientGetPageRangesDiffPager provides operations for iterating over paged responses.
type pageBlobClientGetPageRangesDiffPager struct {
	client    *pageBlobClient
	current   pageBlobClientGetPageRangesDiffResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, pageBlobClientGetPageRangesDiffResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *pageBlobClientGetPageRangesDiffPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *pageBlobClientGetPageRangesDiffPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PageList.NextMarker == nil || len(*p.current.PageList.NextMarker) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.getPageRangesDiffHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current pageBlobClientGetPageRangesDiffResponse page.
func (p *pageBlobClientGetPageRangesDiffPager) PageResponse() pageBlobClientGetPageRangesDiffResponse {
	return p.current
}

// pageBlobClientGetPageRangesPager provides operations for iterating over paged responses.
type pageBlobClientGetPageRangesPager struct {
	client    *pageBlobClient
	current   pageBlobClientGetPageRangesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, pageBlobClientGetPageRangesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *pageBlobClientGetPageRangesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *pageBlobClientGetPageRangesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PageList.NextMarker == nil || len(*p.current.PageList.NextMarker) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.getPageRangesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current pageBlobClientGetPageRangesResponse page.
func (p *pageBlobClientGetPageRangesPager) PageResponse() pageBlobClientGetPageRangesResponse {
	return p.current
}

// serviceClientListContainersSegmentPager provides operations for iterating over paged responses.
type serviceClientListContainersSegmentPager struct {
	client    *serviceClient
	current   serviceClientListContainersSegmentResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, serviceClientListContainersSegmentResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *serviceClientListContainersSegmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *serviceClientListContainersSegmentPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListContainersSegmentResponse.NextMarker == nil || len(*p.current.ListContainersSegmentResponse.NextMarker) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listContainersSegmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current serviceClientListContainersSegmentResponse page.
func (p *serviceClientListContainersSegmentPager) PageResponse() serviceClientListContainersSegmentResponse {
	return p.current
}