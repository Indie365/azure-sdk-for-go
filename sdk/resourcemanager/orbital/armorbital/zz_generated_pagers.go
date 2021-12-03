//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AvailableGroundStationsListByCapabilityPager provides operations for iterating over paged responses.
type AvailableGroundStationsListByCapabilityPager struct {
	client    *AvailableGroundStationsClient
	current   AvailableGroundStationsListByCapabilityResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AvailableGroundStationsListByCapabilityResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AvailableGroundStationsListByCapabilityPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AvailableGroundStationsListByCapabilityPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailableGroundStationListResult.NextLink == nil || len(*p.current.AvailableGroundStationListResult.NextLink) == 0 {
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
		p.err = p.client.listByCapabilityHandleError(resp)
		return false
	}
	result, err := p.client.listByCapabilityHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AvailableGroundStationsListByCapabilityResponse page.
func (p *AvailableGroundStationsListByCapabilityPager) PageResponse() AvailableGroundStationsListByCapabilityResponse {
	return p.current
}
