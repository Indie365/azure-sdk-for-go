//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeducation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// StudentsClient contains the methods for the Students group.
// Don't use this type directly, use NewStudentsClient() instead.
type StudentsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewStudentsClient creates a new instance of StudentsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStudentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*StudentsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &StudentsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CreateOrUpdate - Create and add a new student to the specified lab or update the details of an existing student in a lab.
// Note the student must have a valid tenant to accept the lab after they have been added to lab.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// studentAlias - Student alias.
// parameters - Request parameters that are provided to update student properties.
// options - StudentsClientCreateOrUpdateOptions contains the optional parameters for the StudentsClient.CreateOrUpdate method.
func (client *StudentsClient) CreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, studentAlias string, parameters StudentDetails, options *StudentsClientCreateOrUpdateOptions) (StudentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, studentAlias, parameters, options)
	if err != nil {
		return StudentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StudentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return StudentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StudentsClient) createOrUpdateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, studentAlias string, parameters StudentDetails, options *StudentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/students/{studentAlias}"
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
	if studentAlias == "" {
		return nil, errors.New("parameter studentAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{studentAlias}", url.PathEscape(studentAlias))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StudentsClient) createOrUpdateHandleResponse(resp *http.Response) (StudentsClientCreateOrUpdateResponse, error) {
	result := StudentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StudentDetails); err != nil {
		return StudentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the specified student based on the student alias.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// studentAlias - Student alias.
// options - StudentsClientDeleteOptions contains the optional parameters for the StudentsClient.Delete method.
func (client *StudentsClient) Delete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, studentAlias string, options *StudentsClientDeleteOptions) (StudentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, studentAlias, options)
	if err != nil {
		return StudentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StudentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return StudentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return StudentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StudentsClient) deleteCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, studentAlias string, options *StudentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/students/{studentAlias}"
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
	if studentAlias == "" {
		return nil, errors.New("parameter studentAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{studentAlias}", url.PathEscape(studentAlias))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the details for a specific student in the specified lab by student alias
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// studentAlias - Student alias.
// options - StudentsClientGetOptions contains the optional parameters for the StudentsClient.Get method.
func (client *StudentsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, studentAlias string, options *StudentsClientGetOptions) (StudentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, studentAlias, options)
	if err != nil {
		return StudentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StudentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StudentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StudentsClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, studentAlias string, options *StudentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/students/{studentAlias}"
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
	if studentAlias == "" {
		return nil, errors.New("parameter studentAlias cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{studentAlias}", url.PathEscape(studentAlias))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StudentsClient) getHandleResponse(resp *http.Response) (StudentsClientGetResponse, error) {
	result := StudentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StudentDetails); err != nil {
		return StudentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of details about students that are associated with the specified lab.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// options - StudentsClientListOptions contains the optional parameters for the StudentsClient.List method.
func (client *StudentsClient) NewListPager(billingAccountName string, billingProfileName string, invoiceSectionName string, options *StudentsClientListOptions) *runtime.Pager[StudentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[StudentsClientListResponse]{
		More: func(page StudentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StudentsClientListResponse) (StudentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StudentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StudentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StudentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *StudentsClient) listCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *StudentsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/students"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.IncludeDeleted != nil {
		reqQP.Set("includeDeleted", strconv.FormatBool(*options.IncludeDeleted))
	}
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StudentsClient) listHandleResponse(resp *http.Response) (StudentsClientListResponse, error) {
	result := StudentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StudentListResult); err != nil {
		return StudentsClientListResponse{}, err
	}
	return result, nil
}
