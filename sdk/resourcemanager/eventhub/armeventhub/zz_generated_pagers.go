//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClustersClientListByResourceGroupPager provides operations for iterating over paged responses.
type ClustersClientListByResourceGroupPager struct {
	client    *ClustersClient
	current   ClustersClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterListResult.NextLink == nil || len(*p.current.ClusterListResult.NextLink) == 0 {
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
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersClientListByResourceGroupResponse page.
func (p *ClustersClientListByResourceGroupPager) PageResponse() ClustersClientListByResourceGroupResponse {
	return p.current
}

// ClustersClientListBySubscriptionPager provides operations for iterating over paged responses.
type ClustersClientListBySubscriptionPager struct {
	client    *ClustersClient
	current   ClustersClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClustersClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClustersClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClustersClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterListResult.NextLink == nil || len(*p.current.ClusterListResult.NextLink) == 0 {
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
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClustersClientListBySubscriptionResponse page.
func (p *ClustersClientListBySubscriptionPager) PageResponse() ClustersClientListBySubscriptionResponse {
	return p.current
}

// ConsumerGroupsClientListByEventHubPager provides operations for iterating over paged responses.
type ConsumerGroupsClientListByEventHubPager struct {
	client    *ConsumerGroupsClient
	current   ConsumerGroupsClientListByEventHubResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConsumerGroupsClientListByEventHubResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConsumerGroupsClientListByEventHubPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConsumerGroupsClientListByEventHubPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ConsumerGroupListResult.NextLink == nil || len(*p.current.ConsumerGroupListResult.NextLink) == 0 {
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
	result, err := p.client.listByEventHubHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConsumerGroupsClientListByEventHubResponse page.
func (p *ConsumerGroupsClientListByEventHubPager) PageResponse() ConsumerGroupsClientListByEventHubResponse {
	return p.current
}

// DisasterRecoveryConfigsClientListAuthorizationRulesPager provides operations for iterating over paged responses.
type DisasterRecoveryConfigsClientListAuthorizationRulesPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsClientListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsClientListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DisasterRecoveryConfigsClientListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DisasterRecoveryConfigsClientListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
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
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DisasterRecoveryConfigsClientListAuthorizationRulesResponse page.
func (p *DisasterRecoveryConfigsClientListAuthorizationRulesPager) PageResponse() DisasterRecoveryConfigsClientListAuthorizationRulesResponse {
	return p.current
}

// DisasterRecoveryConfigsClientListPager provides operations for iterating over paged responses.
type DisasterRecoveryConfigsClientListPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DisasterRecoveryConfigsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DisasterRecoveryConfigsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArmDisasterRecoveryListResult.NextLink == nil || len(*p.current.ArmDisasterRecoveryListResult.NextLink) == 0 {
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
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DisasterRecoveryConfigsClientListResponse page.
func (p *DisasterRecoveryConfigsClientListPager) PageResponse() DisasterRecoveryConfigsClientListResponse {
	return p.current
}

// EventHubsClientListAuthorizationRulesPager provides operations for iterating over paged responses.
type EventHubsClientListAuthorizationRulesPager struct {
	client    *EventHubsClient
	current   EventHubsClientListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EventHubsClientListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EventHubsClientListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EventHubsClientListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
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
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current EventHubsClientListAuthorizationRulesResponse page.
func (p *EventHubsClientListAuthorizationRulesPager) PageResponse() EventHubsClientListAuthorizationRulesResponse {
	return p.current
}

// EventHubsClientListByNamespacePager provides operations for iterating over paged responses.
type EventHubsClientListByNamespacePager struct {
	client    *EventHubsClient
	current   EventHubsClientListByNamespaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EventHubsClientListByNamespaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EventHubsClientListByNamespacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EventHubsClientListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListResult.NextLink == nil || len(*p.current.ListResult.NextLink) == 0 {
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
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current EventHubsClientListByNamespaceResponse page.
func (p *EventHubsClientListByNamespacePager) PageResponse() EventHubsClientListByNamespaceResponse {
	return p.current
}

// NamespacesClientListAuthorizationRulesPager provides operations for iterating over paged responses.
type NamespacesClientListAuthorizationRulesPager struct {
	client    *NamespacesClient
	current   NamespacesClientListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesClientListAuthorizationRulesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesClientListAuthorizationRulesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesClientListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
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
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesClientListAuthorizationRulesResponse page.
func (p *NamespacesClientListAuthorizationRulesPager) PageResponse() NamespacesClientListAuthorizationRulesResponse {
	return p.current
}

// NamespacesClientListByResourceGroupPager provides operations for iterating over paged responses.
type NamespacesClientListByResourceGroupPager struct {
	client    *NamespacesClient
	current   NamespacesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EHNamespaceListResult.NextLink == nil || len(*p.current.EHNamespaceListResult.NextLink) == 0 {
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
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesClientListByResourceGroupResponse page.
func (p *NamespacesClientListByResourceGroupPager) PageResponse() NamespacesClientListByResourceGroupResponse {
	return p.current
}

// NamespacesClientListPager provides operations for iterating over paged responses.
type NamespacesClientListPager struct {
	client    *NamespacesClient
	current   NamespacesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, NamespacesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *NamespacesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *NamespacesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EHNamespaceListResult.NextLink == nil || len(*p.current.EHNamespaceListResult.NextLink) == 0 {
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
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current NamespacesClientListResponse page.
func (p *NamespacesClientListPager) PageResponse() NamespacesClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
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
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// PrivateEndpointConnectionsClientListPager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsClientListPager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
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
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsClientListResponse page.
func (p *PrivateEndpointConnectionsClientListPager) PageResponse() PrivateEndpointConnectionsClientListResponse {
	return p.current
}

// SchemaRegistryClientListByNamespacePager provides operations for iterating over paged responses.
type SchemaRegistryClientListByNamespacePager struct {
	client    *SchemaRegistryClient
	current   SchemaRegistryClientListByNamespaceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SchemaRegistryClientListByNamespaceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SchemaRegistryClientListByNamespacePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SchemaRegistryClientListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SchemaGroupListResult.NextLink == nil || len(*p.current.SchemaGroupListResult.NextLink) == 0 {
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
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SchemaRegistryClientListByNamespaceResponse page.
func (p *SchemaRegistryClientListByNamespacePager) PageResponse() SchemaRegistryClientListByNamespaceResponse {
	return p.current
}
