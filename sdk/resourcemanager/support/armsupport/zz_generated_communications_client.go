//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CommunicationsClient contains the methods for the Communications group.
// Don't use this type directly, use NewCommunicationsClient() instead.
type CommunicationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCommunicationsClient creates a new instance of CommunicationsClient with the specified values.
// subscriptionID - Azure subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCommunicationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CommunicationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &CommunicationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CheckNameAvailability - Check the availability of a resource name. This API should be used to check the uniqueness of the
// name for adding a new communication to the support ticket.
// If the operation fails it returns an *azcore.ResponseError type.
// supportTicketName - Support ticket name.
// checkNameAvailabilityInput - Input to check.
// options - CommunicationsClientCheckNameAvailabilityOptions contains the optional parameters for the CommunicationsClient.CheckNameAvailability
// method.
func (client *CommunicationsClient) CheckNameAvailability(ctx context.Context, supportTicketName string, checkNameAvailabilityInput CheckNameAvailabilityInput, options *CommunicationsClientCheckNameAvailabilityOptions) (CommunicationsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, supportTicketName, checkNameAvailabilityInput, options)
	if err != nil {
		return CommunicationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommunicationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunicationsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *CommunicationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, supportTicketName string, checkNameAvailabilityInput CheckNameAvailabilityInput, options *CommunicationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/checkNameAvailability"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityInput)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *CommunicationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (CommunicationsClientCheckNameAvailabilityResponse, error) {
	result := CommunicationsClientCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return CommunicationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Adds a new customer communication to an Azure support ticket.
// If the operation fails it returns an *azcore.ResponseError type.
// supportTicketName - Support ticket name.
// communicationName - Communication name.
// createCommunicationParameters - Communication object.
// options - CommunicationsClientBeginCreateOptions contains the optional parameters for the CommunicationsClient.BeginCreate
// method.
func (client *CommunicationsClient) BeginCreate(ctx context.Context, supportTicketName string, communicationName string, createCommunicationParameters CommunicationDetails, options *CommunicationsClientBeginCreateOptions) (CommunicationsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, supportTicketName, communicationName, createCommunicationParameters, options)
	if err != nil {
		return CommunicationsClientCreatePollerResponse{}, err
	}
	result := CommunicationsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CommunicationsClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return CommunicationsClientCreatePollerResponse{}, err
	}
	result.Poller = &CommunicationsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Adds a new customer communication to an Azure support ticket.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *CommunicationsClient) create(ctx context.Context, supportTicketName string, communicationName string, createCommunicationParameters CommunicationDetails, options *CommunicationsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, supportTicketName, communicationName, createCommunicationParameters, options)
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

// createCreateRequest creates the Create request.
func (client *CommunicationsClient) createCreateRequest(ctx context.Context, supportTicketName string, communicationName string, createCommunicationParameters CommunicationDetails, options *CommunicationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications/{communicationName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if communicationName == "" {
		return nil, errors.New("parameter communicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationName}", url.PathEscape(communicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, createCommunicationParameters)
}

// Get - Returns communication details for a support ticket.
// If the operation fails it returns an *azcore.ResponseError type.
// supportTicketName - Support ticket name.
// communicationName - Communication name.
// options - CommunicationsClientGetOptions contains the optional parameters for the CommunicationsClient.Get method.
func (client *CommunicationsClient) Get(ctx context.Context, supportTicketName string, communicationName string, options *CommunicationsClientGetOptions) (CommunicationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, supportTicketName, communicationName, options)
	if err != nil {
		return CommunicationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CommunicationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CommunicationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CommunicationsClient) getCreateRequest(ctx context.Context, supportTicketName string, communicationName string, options *CommunicationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications/{communicationName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if communicationName == "" {
		return nil, errors.New("parameter communicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationName}", url.PathEscape(communicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CommunicationsClient) getHandleResponse(resp *http.Response) (CommunicationsClientGetResponse, error) {
	result := CommunicationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunicationDetails); err != nil {
		return CommunicationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all communications (attachments not included) for a support ticket.
// You can also filter support ticket communications by CreatedDate or CommunicationType using the $filter parameter. The
// only type of communication supported today is Web. Output will be a paged result
// with nextLink, using which you can retrieve the next set of Communication results.
// Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago,
// a request for data might cause an error.
// If the operation fails it returns an *azcore.ResponseError type.
// supportTicketName - Support ticket name.
// options - CommunicationsClientListOptions contains the optional parameters for the CommunicationsClient.List method.
func (client *CommunicationsClient) List(supportTicketName string, options *CommunicationsClientListOptions) *CommunicationsClientListPager {
	return &CommunicationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, supportTicketName, options)
		},
		advancer: func(ctx context.Context, resp CommunicationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CommunicationsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *CommunicationsClient) listCreateRequest(ctx context.Context, supportTicketName string, options *CommunicationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CommunicationsClient) listHandleResponse(resp *http.Response) (CommunicationsClientListResponse, error) {
	result := CommunicationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunicationsListResult); err != nil {
		return CommunicationsClientListResponse{}, err
	}
	return result, nil
}
