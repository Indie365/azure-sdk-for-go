//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// RoleDefinitionsClient contains the methods for the BillingRoleDefinitions group.
// Don't use this type directly, use NewRoleDefinitionsClient() instead.
type RoleDefinitionsClient struct {
	internal *arm.Client
}

// NewRoleDefinitionsClient creates a new instance of RoleDefinitionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".RoleDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleDefinitionsClient{
	internal: cl,
	}
	return client, nil
}

// GetByBillingAccount - Gets the definition for a role on a billing account. The operation is supported for billing accounts
// with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingRoleDefinitionName - The ID that uniquely identifies a role definition.
//   - options - RoleDefinitionsClientGetByBillingAccountOptions contains the optional parameters for the RoleDefinitionsClient.GetByBillingAccount
//     method.
func (client *RoleDefinitionsClient) GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleDefinitionName string, options *RoleDefinitionsClientGetByBillingAccountOptions) (RoleDefinitionsClientGetByBillingAccountResponse, error) {
	var err error
	req, err := client.getByBillingAccountCreateRequest(ctx, billingAccountName, billingRoleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsClientGetByBillingAccountResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleDefinitionsClientGetByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleDefinitionsClientGetByBillingAccountResponse{}, err
	}
	resp, err := client.getByBillingAccountHandleResponse(httpResp)
	return resp, err
}

// getByBillingAccountCreateRequest creates the GetByBillingAccount request.
func (client *RoleDefinitionsClient) getByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, billingRoleDefinitionName string, options *RoleDefinitionsClientGetByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions/{billingRoleDefinitionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingRoleDefinitionName == "" {
		return nil, errors.New("parameter billingRoleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleDefinitionName}", url.PathEscape(billingRoleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingAccountHandleResponse handles the GetByBillingAccount response.
func (client *RoleDefinitionsClient) getByBillingAccountHandleResponse(resp *http.Response) (RoleDefinitionsClientGetByBillingAccountResponse, error) {
	result := RoleDefinitionsClientGetByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsClientGetByBillingAccountResponse{}, err
	}
	return result, nil
}

// GetByBillingProfile - Gets the definition for a role on a billing profile. The operation is supported for billing accounts
// with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - billingRoleDefinitionName - The ID that uniquely identifies a role definition.
//   - options - RoleDefinitionsClientGetByBillingProfileOptions contains the optional parameters for the RoleDefinitionsClient.GetByBillingProfile
//     method.
func (client *RoleDefinitionsClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string, options *RoleDefinitionsClientGetByBillingProfileOptions) (RoleDefinitionsClientGetByBillingProfileResponse, error) {
	var err error
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, billingRoleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsClientGetByBillingProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleDefinitionsClientGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleDefinitionsClientGetByBillingProfileResponse{}, err
	}
	resp, err := client.getByBillingProfileHandleResponse(httpResp)
	return resp, err
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *RoleDefinitionsClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string, options *RoleDefinitionsClientGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions/{billingRoleDefinitionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if billingRoleDefinitionName == "" {
		return nil, errors.New("parameter billingRoleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleDefinitionName}", url.PathEscape(billingRoleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *RoleDefinitionsClient) getByBillingProfileHandleResponse(resp *http.Response) (RoleDefinitionsClientGetByBillingProfileResponse, error) {
	result := RoleDefinitionsClientGetByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsClientGetByBillingProfileResponse{}, err
	}
	return result, nil
}

// GetByInvoiceSection - Gets the definition for a role on an invoice section. The operation is supported only for billing
// accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - billingRoleDefinitionName - The ID that uniquely identifies a role definition.
//   - options - RoleDefinitionsClientGetByInvoiceSectionOptions contains the optional parameters for the RoleDefinitionsClient.GetByInvoiceSection
//     method.
func (client *RoleDefinitionsClient) GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string, options *RoleDefinitionsClientGetByInvoiceSectionOptions) (RoleDefinitionsClientGetByInvoiceSectionResponse, error) {
	var err error
	req, err := client.getByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleDefinitionName, options)
	if err != nil {
		return RoleDefinitionsClientGetByInvoiceSectionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleDefinitionsClientGetByInvoiceSectionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleDefinitionsClientGetByInvoiceSectionResponse{}, err
	}
	resp, err := client.getByInvoiceSectionHandleResponse(httpResp)
	return resp, err
}

// getByInvoiceSectionCreateRequest creates the GetByInvoiceSection request.
func (client *RoleDefinitionsClient) getByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string, options *RoleDefinitionsClientGetByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions/{billingRoleDefinitionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	if billingRoleDefinitionName == "" {
		return nil, errors.New("parameter billingRoleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleDefinitionName}", url.PathEscape(billingRoleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByInvoiceSectionHandleResponse handles the GetByInvoiceSection response.
func (client *RoleDefinitionsClient) getByInvoiceSectionHandleResponse(resp *http.Response) (RoleDefinitionsClientGetByInvoiceSectionResponse, error) {
	result := RoleDefinitionsClientGetByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return RoleDefinitionsClientGetByInvoiceSectionResponse{}, err
	}
	return result, nil
}

// NewListByBillingAccountPager - Lists the role definitions for a billing account. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - RoleDefinitionsClientListByBillingAccountOptions contains the optional parameters for the RoleDefinitionsClient.NewListByBillingAccountPager
//     method.
func (client *RoleDefinitionsClient) NewListByBillingAccountPager(billingAccountName string, options *RoleDefinitionsClientListByBillingAccountOptions) (*runtime.Pager[RoleDefinitionsClientListByBillingAccountResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RoleDefinitionsClientListByBillingAccountResponse]{
		More: func(page RoleDefinitionsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleDefinitionsClientListByBillingAccountResponse) (RoleDefinitionsClientListByBillingAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleDefinitionsClientListByBillingAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleDefinitionsClientListByBillingAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleDefinitionsClientListByBillingAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *RoleDefinitionsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *RoleDefinitionsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *RoleDefinitionsClient) listByBillingAccountHandleResponse(resp *http.Response) (RoleDefinitionsClientListByBillingAccountResponse, error) {
	result := RoleDefinitionsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinitionListResult); err != nil {
		return RoleDefinitionsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the role definitions for a billing profile. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - RoleDefinitionsClientListByBillingProfileOptions contains the optional parameters for the RoleDefinitionsClient.NewListByBillingProfilePager
//     method.
func (client *RoleDefinitionsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *RoleDefinitionsClientListByBillingProfileOptions) (*runtime.Pager[RoleDefinitionsClientListByBillingProfileResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RoleDefinitionsClientListByBillingProfileResponse]{
		More: func(page RoleDefinitionsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleDefinitionsClientListByBillingProfileResponse) (RoleDefinitionsClientListByBillingProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleDefinitionsClientListByBillingProfileResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleDefinitionsClientListByBillingProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleDefinitionsClientListByBillingProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *RoleDefinitionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *RoleDefinitionsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *RoleDefinitionsClient) listByBillingProfileHandleResponse(resp *http.Response) (RoleDefinitionsClientListByBillingProfileResponse, error) {
	result := RoleDefinitionsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinitionListResult); err != nil {
		return RoleDefinitionsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// NewListByInvoiceSectionPager - Lists the role definitions for an invoice section. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - options - RoleDefinitionsClientListByInvoiceSectionOptions contains the optional parameters for the RoleDefinitionsClient.NewListByInvoiceSectionPager
//     method.
func (client *RoleDefinitionsClient) NewListByInvoiceSectionPager(billingAccountName string, billingProfileName string, invoiceSectionName string, options *RoleDefinitionsClientListByInvoiceSectionOptions) (*runtime.Pager[RoleDefinitionsClientListByInvoiceSectionResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RoleDefinitionsClientListByInvoiceSectionResponse]{
		More: func(page RoleDefinitionsClientListByInvoiceSectionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleDefinitionsClientListByInvoiceSectionResponse) (RoleDefinitionsClientListByInvoiceSectionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleDefinitionsClientListByInvoiceSectionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleDefinitionsClientListByInvoiceSectionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleDefinitionsClientListByInvoiceSectionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInvoiceSectionHandleResponse(resp)
		},
	})
}

// listByInvoiceSectionCreateRequest creates the ListByInvoiceSection request.
func (client *RoleDefinitionsClient) listByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *RoleDefinitionsClientListByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInvoiceSectionHandleResponse handles the ListByInvoiceSection response.
func (client *RoleDefinitionsClient) listByInvoiceSectionHandleResponse(resp *http.Response) (RoleDefinitionsClientListByInvoiceSectionResponse, error) {
	result := RoleDefinitionsClientListByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinitionListResult); err != nil {
		return RoleDefinitionsClientListByInvoiceSectionResponse{}, err
	}
	return result, nil
}

