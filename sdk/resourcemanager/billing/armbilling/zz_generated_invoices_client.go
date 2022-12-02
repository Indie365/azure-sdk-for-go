//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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
	"strings"
)

// InvoicesClient contains the methods for the Invoices group.
// Don't use this type directly, use NewInvoicesClient() instead.
type InvoicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInvoicesClient creates a new instance of InvoicesClient with the specified values.
// subscriptionID - The ID that uniquely identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInvoicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InvoicesClient, error) {
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
	client := &InvoicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDownloadBillingSubscriptionInvoice - Gets a URL to download an invoice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// invoiceName - The ID that uniquely identifies an invoice.
// downloadToken - Download token with document source and document ID.
// options - InvoicesClientBeginDownloadBillingSubscriptionInvoiceOptions contains the optional parameters for the InvoicesClient.BeginDownloadBillingSubscriptionInvoice
// method.
func (client *InvoicesClient) BeginDownloadBillingSubscriptionInvoice(ctx context.Context, invoiceName string, downloadToken string, options *InvoicesClientBeginDownloadBillingSubscriptionInvoiceOptions) (*runtime.Poller[InvoicesClientDownloadBillingSubscriptionInvoiceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.downloadBillingSubscriptionInvoice(ctx, invoiceName, downloadToken, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[InvoicesClientDownloadBillingSubscriptionInvoiceResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InvoicesClientDownloadBillingSubscriptionInvoiceResponse](options.ResumeToken, client.pl, nil)
	}
}

// DownloadBillingSubscriptionInvoice - Gets a URL to download an invoice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
func (client *InvoicesClient) downloadBillingSubscriptionInvoice(ctx context.Context, invoiceName string, downloadToken string, options *InvoicesClientBeginDownloadBillingSubscriptionInvoiceOptions) (*http.Response, error) {
	req, err := client.downloadBillingSubscriptionInvoiceCreateRequest(ctx, invoiceName, downloadToken, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// downloadBillingSubscriptionInvoiceCreateRequest creates the DownloadBillingSubscriptionInvoice request.
func (client *InvoicesClient) downloadBillingSubscriptionInvoiceCreateRequest(ctx context.Context, invoiceName string, downloadToken string, options *InvoicesClientBeginDownloadBillingSubscriptionInvoiceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	reqQP.Set("downloadToken", downloadToken)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDownloadInvoice - Gets a URL to download an invoice. The operation is supported for billing accounts with agreement
// type Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// invoiceName - The ID that uniquely identifies an invoice.
// downloadToken - Download token with document source and document ID.
// options - InvoicesClientBeginDownloadInvoiceOptions contains the optional parameters for the InvoicesClient.BeginDownloadInvoice
// method.
func (client *InvoicesClient) BeginDownloadInvoice(ctx context.Context, billingAccountName string, invoiceName string, downloadToken string, options *InvoicesClientBeginDownloadInvoiceOptions) (*runtime.Poller[InvoicesClientDownloadInvoiceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.downloadInvoice(ctx, billingAccountName, invoiceName, downloadToken, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[InvoicesClientDownloadInvoiceResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InvoicesClientDownloadInvoiceResponse](options.ResumeToken, client.pl, nil)
	}
}

// DownloadInvoice - Gets a URL to download an invoice. The operation is supported for billing accounts with agreement type
// Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
func (client *InvoicesClient) downloadInvoice(ctx context.Context, billingAccountName string, invoiceName string, downloadToken string, options *InvoicesClientBeginDownloadInvoiceOptions) (*http.Response, error) {
	req, err := client.downloadInvoiceCreateRequest(ctx, billingAccountName, invoiceName, downloadToken, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// downloadInvoiceCreateRequest creates the DownloadInvoice request.
func (client *InvoicesClient) downloadInvoiceCreateRequest(ctx context.Context, billingAccountName string, invoiceName string, downloadToken string, options *InvoicesClientBeginDownloadInvoiceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/download"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	reqQP.Set("downloadToken", downloadToken)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDownloadMultipleBillingProfileInvoices - Gets a URL to download multiple invoice documents (invoice pdf, tax receipts,
// credit notes) as a zip file. The operation is supported for billing accounts with agreement type Microsoft Partner
// Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// downloadUrls - An array of download urls for individual documents
// options - InvoicesClientBeginDownloadMultipleBillingProfileInvoicesOptions contains the optional parameters for the InvoicesClient.BeginDownloadMultipleBillingProfileInvoices
// method.
func (client *InvoicesClient) BeginDownloadMultipleBillingProfileInvoices(ctx context.Context, billingAccountName string, downloadUrls []*string, options *InvoicesClientBeginDownloadMultipleBillingProfileInvoicesOptions) (*runtime.Poller[InvoicesClientDownloadMultipleBillingProfileInvoicesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.downloadMultipleBillingProfileInvoices(ctx, billingAccountName, downloadUrls, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[InvoicesClientDownloadMultipleBillingProfileInvoicesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InvoicesClientDownloadMultipleBillingProfileInvoicesResponse](options.ResumeToken, client.pl, nil)
	}
}

// DownloadMultipleBillingProfileInvoices - Gets a URL to download multiple invoice documents (invoice pdf, tax receipts,
// credit notes) as a zip file. The operation is supported for billing accounts with agreement type Microsoft Partner
// Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
func (client *InvoicesClient) downloadMultipleBillingProfileInvoices(ctx context.Context, billingAccountName string, downloadUrls []*string, options *InvoicesClientBeginDownloadMultipleBillingProfileInvoicesOptions) (*http.Response, error) {
	req, err := client.downloadMultipleBillingProfileInvoicesCreateRequest(ctx, billingAccountName, downloadUrls, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// downloadMultipleBillingProfileInvoicesCreateRequest creates the DownloadMultipleBillingProfileInvoices request.
func (client *InvoicesClient) downloadMultipleBillingProfileInvoicesCreateRequest(ctx context.Context, billingAccountName string, downloadUrls []*string, options *InvoicesClientBeginDownloadMultipleBillingProfileInvoicesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/downloadDocuments"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, downloadUrls)
}

// BeginDownloadMultipleBillingSubscriptionInvoices - Gets a URL to download multiple invoice documents (invoice pdf, tax
// receipts, credit notes) as a zip file.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// downloadUrls - An array of download urls for individual documents
// options - InvoicesClientBeginDownloadMultipleBillingSubscriptionInvoicesOptions contains the optional parameters for the
// InvoicesClient.BeginDownloadMultipleBillingSubscriptionInvoices method.
func (client *InvoicesClient) BeginDownloadMultipleBillingSubscriptionInvoices(ctx context.Context, downloadUrls []*string, options *InvoicesClientBeginDownloadMultipleBillingSubscriptionInvoicesOptions) (*runtime.Poller[InvoicesClientDownloadMultipleBillingSubscriptionInvoicesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.downloadMultipleBillingSubscriptionInvoices(ctx, downloadUrls, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[InvoicesClientDownloadMultipleBillingSubscriptionInvoicesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InvoicesClientDownloadMultipleBillingSubscriptionInvoicesResponse](options.ResumeToken, client.pl, nil)
	}
}

// DownloadMultipleBillingSubscriptionInvoices - Gets a URL to download multiple invoice documents (invoice pdf, tax receipts,
// credit notes) as a zip file.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
func (client *InvoicesClient) downloadMultipleBillingSubscriptionInvoices(ctx context.Context, downloadUrls []*string, options *InvoicesClientBeginDownloadMultipleBillingSubscriptionInvoicesOptions) (*http.Response, error) {
	req, err := client.downloadMultipleBillingSubscriptionInvoicesCreateRequest(ctx, downloadUrls, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// downloadMultipleBillingSubscriptionInvoicesCreateRequest creates the DownloadMultipleBillingSubscriptionInvoices request.
func (client *InvoicesClient) downloadMultipleBillingSubscriptionInvoicesCreateRequest(ctx context.Context, downloadUrls []*string, options *InvoicesClientBeginDownloadMultipleBillingSubscriptionInvoicesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/downloadDocuments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, downloadUrls)
}

// Get - Gets an invoice by billing account name and ID. The operation is supported for billing accounts with agreement type
// Microsoft Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// invoiceName - The ID that uniquely identifies an invoice.
// options - InvoicesClientGetOptions contains the optional parameters for the InvoicesClient.Get method.
func (client *InvoicesClient) Get(ctx context.Context, billingAccountName string, invoiceName string, options *InvoicesClientGetOptions) (InvoicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, invoiceName, options)
	if err != nil {
		return InvoicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InvoicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InvoicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InvoicesClient) getCreateRequest(ctx context.Context, billingAccountName string, invoiceName string, options *InvoicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InvoicesClient) getHandleResponse(resp *http.Response) (InvoicesClientGetResponse, error) {
	result := InvoicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Invoice); err != nil {
		return InvoicesClientGetResponse{}, err
	}
	return result, nil
}

// GetByID - Gets an invoice by ID. The operation is supported for billing accounts with agreement type Microsoft Partner
// Agreement or Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// invoiceName - The ID that uniquely identifies an invoice.
// options - InvoicesClientGetByIDOptions contains the optional parameters for the InvoicesClient.GetByID method.
func (client *InvoicesClient) GetByID(ctx context.Context, invoiceName string, options *InvoicesClientGetByIDOptions) (InvoicesClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, invoiceName, options)
	if err != nil {
		return InvoicesClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InvoicesClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InvoicesClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *InvoicesClient) getByIDCreateRequest(ctx context.Context, invoiceName string, options *InvoicesClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/default/invoices/{invoiceName}"
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *InvoicesClient) getByIDHandleResponse(resp *http.Response) (InvoicesClientGetByIDResponse, error) {
	result := InvoicesClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Invoice); err != nil {
		return InvoicesClientGetByIDResponse{}, err
	}
	return result, nil
}

// GetBySubscriptionAndInvoiceID - Gets an invoice by subscription ID and invoice ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// invoiceName - The ID that uniquely identifies an invoice.
// options - InvoicesClientGetBySubscriptionAndInvoiceIDOptions contains the optional parameters for the InvoicesClient.GetBySubscriptionAndInvoiceID
// method.
func (client *InvoicesClient) GetBySubscriptionAndInvoiceID(ctx context.Context, invoiceName string, options *InvoicesClientGetBySubscriptionAndInvoiceIDOptions) (InvoicesClientGetBySubscriptionAndInvoiceIDResponse, error) {
	req, err := client.getBySubscriptionAndInvoiceIDCreateRequest(ctx, invoiceName, options)
	if err != nil {
		return InvoicesClientGetBySubscriptionAndInvoiceIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InvoicesClientGetBySubscriptionAndInvoiceIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InvoicesClientGetBySubscriptionAndInvoiceIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getBySubscriptionAndInvoiceIDHandleResponse(resp)
}

// getBySubscriptionAndInvoiceIDCreateRequest creates the GetBySubscriptionAndInvoiceID request.
func (client *InvoicesClient) getBySubscriptionAndInvoiceIDCreateRequest(ctx context.Context, invoiceName string, options *InvoicesClientGetBySubscriptionAndInvoiceIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBySubscriptionAndInvoiceIDHandleResponse handles the GetBySubscriptionAndInvoiceID response.
func (client *InvoicesClient) getBySubscriptionAndInvoiceIDHandleResponse(resp *http.Response) (InvoicesClientGetBySubscriptionAndInvoiceIDResponse, error) {
	result := InvoicesClientGetBySubscriptionAndInvoiceIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Invoice); err != nil {
		return InvoicesClientGetBySubscriptionAndInvoiceIDResponse{}, err
	}
	return result, nil
}

// NewListByBillingAccountPager - Lists the invoices for a billing account for a given start date and end date. The operation
// is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer
// Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// periodStartDate - The start date to fetch the invoices. The date should be specified in MM-DD-YYYY format.
// periodEndDate - The end date to fetch the invoices. The date should be specified in MM-DD-YYYY format.
// options - InvoicesClientListByBillingAccountOptions contains the optional parameters for the InvoicesClient.ListByBillingAccount
// method.
func (client *InvoicesClient) NewListByBillingAccountPager(billingAccountName string, periodStartDate string, periodEndDate string, options *InvoicesClientListByBillingAccountOptions) *runtime.Pager[InvoicesClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[InvoicesClientListByBillingAccountResponse]{
		More: func(page InvoicesClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InvoicesClientListByBillingAccountResponse) (InvoicesClientListByBillingAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingAccountCreateRequest(ctx, billingAccountName, periodStartDate, periodEndDate, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InvoicesClientListByBillingAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InvoicesClientListByBillingAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InvoicesClientListByBillingAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *InvoicesClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string, options *InvoicesClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	reqQP.Set("periodStartDate", periodStartDate)
	reqQP.Set("periodEndDate", periodEndDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *InvoicesClient) listByBillingAccountHandleResponse(resp *http.Response) (InvoicesClientListByBillingAccountResponse, error) {
	result := InvoicesClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceListResult); err != nil {
		return InvoicesClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the invoices for a billing profile for a given start date and end date. The operation
// is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer
// Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// billingAccountName - The ID that uniquely identifies a billing account.
// billingProfileName - The ID that uniquely identifies a billing profile.
// periodStartDate - The start date to fetch the invoices. The date should be specified in MM-DD-YYYY format.
// periodEndDate - The end date to fetch the invoices. The date should be specified in MM-DD-YYYY format.
// options - InvoicesClientListByBillingProfileOptions contains the optional parameters for the InvoicesClient.ListByBillingProfile
// method.
func (client *InvoicesClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string, options *InvoicesClientListByBillingProfileOptions) *runtime.Pager[InvoicesClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[InvoicesClientListByBillingProfileResponse]{
		More: func(page InvoicesClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InvoicesClientListByBillingProfileResponse) (InvoicesClientListByBillingProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, periodStartDate, periodEndDate, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InvoicesClientListByBillingProfileResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InvoicesClientListByBillingProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InvoicesClientListByBillingProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *InvoicesClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string, options *InvoicesClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	reqQP.Set("periodStartDate", periodStartDate)
	reqQP.Set("periodEndDate", periodEndDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *InvoicesClient) listByBillingProfileHandleResponse(resp *http.Response) (InvoicesClientListByBillingProfileResponse, error) {
	result := InvoicesClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceListResult); err != nil {
		return InvoicesClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// NewListByBillingSubscriptionPager - Lists the invoices for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// periodStartDate - Invoice period start date.
// periodEndDate - Invoice period end date.
// options - InvoicesClientListByBillingSubscriptionOptions contains the optional parameters for the InvoicesClient.ListByBillingSubscription
// method.
func (client *InvoicesClient) NewListByBillingSubscriptionPager(periodStartDate string, periodEndDate string, options *InvoicesClientListByBillingSubscriptionOptions) *runtime.Pager[InvoicesClientListByBillingSubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[InvoicesClientListByBillingSubscriptionResponse]{
		More: func(page InvoicesClientListByBillingSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InvoicesClientListByBillingSubscriptionResponse) (InvoicesClientListByBillingSubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingSubscriptionCreateRequest(ctx, periodStartDate, periodEndDate, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InvoicesClientListByBillingSubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InvoicesClientListByBillingSubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InvoicesClientListByBillingSubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingSubscriptionHandleResponse(resp)
		},
	})
}

// listByBillingSubscriptionCreateRequest creates the ListByBillingSubscription request.
func (client *InvoicesClient) listByBillingSubscriptionCreateRequest(ctx context.Context, periodStartDate string, periodEndDate string, options *InvoicesClientListByBillingSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("periodStartDate", periodStartDate)
	reqQP.Set("periodEndDate", periodEndDate)
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingSubscriptionHandleResponse handles the ListByBillingSubscription response.
func (client *InvoicesClient) listByBillingSubscriptionHandleResponse(resp *http.Response) (InvoicesClientListByBillingSubscriptionResponse, error) {
	result := InvoicesClientListByBillingSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvoiceListResult); err != nil {
		return InvoicesClientListByBillingSubscriptionResponse{}, err
	}
	return result, nil
}