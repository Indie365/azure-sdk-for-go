//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// KeyVaultClientGetCertificateIssuersPager provides operations for iterating over paged responses.
type KeyVaultClientGetCertificateIssuersPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetCertificateIssuersResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetCertificateIssuersResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetCertificateIssuersPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetCertificateIssuersPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateIssuerListResult.NextLink == nil || len(*p.current.CertificateIssuerListResult.NextLink) == 0 {
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
	result, err := p.client.getCertificateIssuersHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetCertificateIssuersResponse page.
func (p *KeyVaultClientGetCertificateIssuersPager) PageResponse() KeyVaultClientGetCertificateIssuersResponse {
	return p.current
}

// KeyVaultClientGetCertificateVersionsPager provides operations for iterating over paged responses.
type KeyVaultClientGetCertificateVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetCertificateVersionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetCertificateVersionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetCertificateVersionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetCertificateVersionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
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
	result, err := p.client.getCertificateVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetCertificateVersionsResponse page.
func (p *KeyVaultClientGetCertificateVersionsPager) PageResponse() KeyVaultClientGetCertificateVersionsResponse {
	return p.current
}

// KeyVaultClientGetCertificatesPager provides operations for iterating over paged responses.
type KeyVaultClientGetCertificatesPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetCertificatesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetCertificatesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetCertificatesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetCertificatesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
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
	result, err := p.client.getCertificatesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetCertificatesResponse page.
func (p *KeyVaultClientGetCertificatesPager) PageResponse() KeyVaultClientGetCertificatesResponse {
	return p.current
}

// KeyVaultClientGetDeletedCertificatesPager provides operations for iterating over paged responses.
type KeyVaultClientGetDeletedCertificatesPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedCertificatesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedCertificatesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *KeyVaultClientGetDeletedCertificatesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *KeyVaultClientGetDeletedCertificatesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedCertificateListResult.NextLink == nil || len(*p.current.DeletedCertificateListResult.NextLink) == 0 {
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
	result, err := p.client.getDeletedCertificatesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current KeyVaultClientGetDeletedCertificatesResponse page.
func (p *KeyVaultClientGetDeletedCertificatesPager) PageResponse() KeyVaultClientGetDeletedCertificatesResponse {
	return p.current
}

// RoleAssignmentsClientListForScopePager provides operations for iterating over paged responses.
type RoleAssignmentsClientListForScopePager struct {
	client    *RoleAssignmentsClient
	current   RoleAssignmentsClientListForScopeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleAssignmentsClientListForScopeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RoleAssignmentsClientListForScopePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RoleAssignmentsClientListForScopePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentListResult.NextLink == nil || len(*p.current.RoleAssignmentListResult.NextLink) == 0 {
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
	result, err := p.client.listForScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RoleAssignmentsClientListForScopeResponse page.
func (p *RoleAssignmentsClientListForScopePager) PageResponse() RoleAssignmentsClientListForScopeResponse {
	return p.current
}

// RoleDefinitionsClientListPager provides operations for iterating over paged responses.
type RoleDefinitionsClientListPager struct {
	client    *RoleDefinitionsClient
	current   RoleDefinitionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleDefinitionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RoleDefinitionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RoleDefinitionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleDefinitionListResult.NextLink == nil || len(*p.current.RoleDefinitionListResult.NextLink) == 0 {
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

// PageResponse returns the current RoleDefinitionsClientListResponse page.
func (p *RoleDefinitionsClientListPager) PageResponse() RoleDefinitionsClientListResponse {
	return p.current
}
