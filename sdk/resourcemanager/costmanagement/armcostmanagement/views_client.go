//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement

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

// ViewsClient contains the methods for the Views group.
// Don't use this type directly, use NewViewsClient() instead.
type ViewsClient struct {
	internal *arm.Client
}

// NewViewsClient creates a new instance of ViewsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewViewsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ViewsClient, error) {
	cl, err := arm.NewClient(moduleName+".ViewsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ViewsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - The operation to create or update a view. Update operation requires latest eTag to be set in the request.
// You may obtain the latest eTag by performing a get operation. Create operation does not
// require eTag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - viewName - View name
//   - parameters - Parameters supplied to the CreateOrUpdate View operation.
//   - options - ViewsClientCreateOrUpdateOptions contains the optional parameters for the ViewsClient.CreateOrUpdate method.
func (client *ViewsClient) CreateOrUpdate(ctx context.Context, viewName string, parameters View, options *ViewsClientCreateOrUpdateOptions) (ViewsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, viewName, parameters, options)
	if err != nil {
		return ViewsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ViewsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ViewsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ViewsClient) createOrUpdateCreateRequest(ctx context.Context, viewName string, parameters View, options *ViewsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/views/{viewName}"
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ViewsClient) createOrUpdateHandleResponse(resp *http.Response) (ViewsClientCreateOrUpdateResponse, error) {
	result := ViewsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.View); err != nil {
		return ViewsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateByScope - The operation to create or update a view. Update operation requires latest eTag to be set in the
// request. You may obtain the latest eTag by performing a get operation. Create operation does not
// require eTag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - scope - The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope,
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
//     scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}'
//     for
//     Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
//     for EnrollmentAccount scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile
//     scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection
//     scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for
//     Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External
//     Billing Account scope and
//     'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
//   - viewName - View name
//   - parameters - Parameters supplied to the CreateOrUpdate View operation.
//   - options - ViewsClientCreateOrUpdateByScopeOptions contains the optional parameters for the ViewsClient.CreateOrUpdateByScope
//     method.
func (client *ViewsClient) CreateOrUpdateByScope(ctx context.Context, scope string, viewName string, parameters View, options *ViewsClientCreateOrUpdateByScopeOptions) (ViewsClientCreateOrUpdateByScopeResponse, error) {
	req, err := client.createOrUpdateByScopeCreateRequest(ctx, scope, viewName, parameters, options)
	if err != nil {
		return ViewsClientCreateOrUpdateByScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ViewsClientCreateOrUpdateByScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ViewsClientCreateOrUpdateByScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateByScopeHandleResponse(resp)
}

// createOrUpdateByScopeCreateRequest creates the CreateOrUpdateByScope request.
func (client *ViewsClient) createOrUpdateByScopeCreateRequest(ctx context.Context, scope string, viewName string, parameters View, options *ViewsClientCreateOrUpdateByScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/views/{viewName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateByScopeHandleResponse handles the CreateOrUpdateByScope response.
func (client *ViewsClient) createOrUpdateByScopeHandleResponse(resp *http.Response) (ViewsClientCreateOrUpdateByScopeResponse, error) {
	result := ViewsClientCreateOrUpdateByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.View); err != nil {
		return ViewsClientCreateOrUpdateByScopeResponse{}, err
	}
	return result, nil
}

// Delete - The operation to delete a view.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - viewName - View name
//   - options - ViewsClientDeleteOptions contains the optional parameters for the ViewsClient.Delete method.
func (client *ViewsClient) Delete(ctx context.Context, viewName string, options *ViewsClientDeleteOptions) (ViewsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, viewName, options)
	if err != nil {
		return ViewsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ViewsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ViewsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ViewsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ViewsClient) deleteCreateRequest(ctx context.Context, viewName string, options *ViewsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/views/{viewName}"
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteByScope - The operation to delete a view.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - scope - The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope,
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
//     scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}'
//     for
//     Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
//     for EnrollmentAccount scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile
//     scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection
//     scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for
//     Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External
//     Billing Account scope and
//     'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
//   - viewName - View name
//   - options - ViewsClientDeleteByScopeOptions contains the optional parameters for the ViewsClient.DeleteByScope method.
func (client *ViewsClient) DeleteByScope(ctx context.Context, scope string, viewName string, options *ViewsClientDeleteByScopeOptions) (ViewsClientDeleteByScopeResponse, error) {
	req, err := client.deleteByScopeCreateRequest(ctx, scope, viewName, options)
	if err != nil {
		return ViewsClientDeleteByScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ViewsClientDeleteByScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ViewsClientDeleteByScopeResponse{}, runtime.NewResponseError(resp)
	}
	return ViewsClientDeleteByScopeResponse{}, nil
}

// deleteByScopeCreateRequest creates the DeleteByScope request.
func (client *ViewsClient) deleteByScopeCreateRequest(ctx context.Context, scope string, viewName string, options *ViewsClientDeleteByScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/views/{viewName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the view by view name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - viewName - View name
//   - options - ViewsClientGetOptions contains the optional parameters for the ViewsClient.Get method.
func (client *ViewsClient) Get(ctx context.Context, viewName string, options *ViewsClientGetOptions) (ViewsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, viewName, options)
	if err != nil {
		return ViewsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ViewsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ViewsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ViewsClient) getCreateRequest(ctx context.Context, viewName string, options *ViewsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/views/{viewName}"
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ViewsClient) getHandleResponse(resp *http.Response) (ViewsClientGetResponse, error) {
	result := ViewsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.View); err != nil {
		return ViewsClientGetResponse{}, err
	}
	return result, nil
}

// GetByScope - Gets the view for the defined scope by view name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - scope - The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope,
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
//     scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}'
//     for
//     Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
//     for EnrollmentAccount scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile
//     scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection
//     scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for
//     Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External
//     Billing Account scope and
//     'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
//   - viewName - View name
//   - options - ViewsClientGetByScopeOptions contains the optional parameters for the ViewsClient.GetByScope method.
func (client *ViewsClient) GetByScope(ctx context.Context, scope string, viewName string, options *ViewsClientGetByScopeOptions) (ViewsClientGetByScopeResponse, error) {
	req, err := client.getByScopeCreateRequest(ctx, scope, viewName, options)
	if err != nil {
		return ViewsClientGetByScopeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ViewsClientGetByScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ViewsClientGetByScopeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByScopeHandleResponse(resp)
}

// getByScopeCreateRequest creates the GetByScope request.
func (client *ViewsClient) getByScopeCreateRequest(ctx context.Context, scope string, viewName string, options *ViewsClientGetByScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/views/{viewName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if viewName == "" {
		return nil, errors.New("parameter viewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{viewName}", url.PathEscape(viewName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByScopeHandleResponse handles the GetByScope response.
func (client *ViewsClient) getByScopeHandleResponse(resp *http.Response) (ViewsClientGetByScopeResponse, error) {
	result := ViewsClientGetByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.View); err != nil {
		return ViewsClientGetByScopeResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all views by tenant and object.
//
// Generated from API version 2021-10-01
//   - options - ViewsClientListOptions contains the optional parameters for the ViewsClient.NewListPager method.
func (client *ViewsClient) NewListPager(options *ViewsClientListOptions) *runtime.Pager[ViewsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ViewsClientListResponse]{
		More: func(page ViewsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ViewsClientListResponse) (ViewsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ViewsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ViewsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ViewsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ViewsClient) listCreateRequest(ctx context.Context, options *ViewsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/views"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ViewsClient) listHandleResponse(resp *http.Response) (ViewsClientListResponse, error) {
	result := ViewsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ViewListResult); err != nil {
		return ViewsClientListResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - Lists all views at the given scope.
//
// Generated from API version 2021-10-01
//   - scope - The scope associated with view operations. This includes 'subscriptions/{subscriptionId}' for subscription scope,
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
//     scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}'
//     for
//     Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
//     for EnrollmentAccount scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile
//     scope,
//     'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection
//     scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for
//     Management Group scope, 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for External
//     Billing Account scope and
//     'providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for External Subscription scope.
//   - options - ViewsClientListByScopeOptions contains the optional parameters for the ViewsClient.NewListByScopePager method.
func (client *ViewsClient) NewListByScopePager(scope string, options *ViewsClientListByScopeOptions) *runtime.Pager[ViewsClientListByScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[ViewsClientListByScopeResponse]{
		More: func(page ViewsClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ViewsClientListByScopeResponse) (ViewsClientListByScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ViewsClientListByScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ViewsClientListByScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ViewsClientListByScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByScopeHandleResponse(resp)
		},
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *ViewsClient) listByScopeCreateRequest(ctx context.Context, scope string, options *ViewsClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/views"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByScopeHandleResponse handles the ListByScope response.
func (client *ViewsClient) listByScopeHandleResponse(resp *http.Response) (ViewsClientListByScopeResponse, error) {
	result := ViewsClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ViewListResult); err != nil {
		return ViewsClientListByScopeResponse{}, err
	}
	return result, nil
}
