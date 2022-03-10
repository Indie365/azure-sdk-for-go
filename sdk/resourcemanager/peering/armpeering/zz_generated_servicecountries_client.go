//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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
	"strings"
)

// ServiceCountriesClient contains the methods for the PeeringServiceCountries group.
// Don't use this type directly, use NewServiceCountriesClient() instead.
type ServiceCountriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceCountriesClient creates a new instance of ServiceCountriesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceCountriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServiceCountriesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServiceCountriesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Lists all of the available countries for peering service.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ServiceCountriesClientListOptions contains the optional parameters for the ServiceCountriesClient.List method.
func (client *ServiceCountriesClient) List(options *ServiceCountriesClientListOptions) *ServiceCountriesClientListPager {
	return &ServiceCountriesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ServiceCountriesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServiceCountryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServiceCountriesClient) listCreateRequest(ctx context.Context, options *ServiceCountriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceCountries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceCountriesClient) listHandleResponse(resp *http.Response) (ServiceCountriesClientListResponse, error) {
	result := ServiceCountriesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceCountryListResult); err != nil {
		return ServiceCountriesClientListResponse{}, err
	}
	return result, nil
}
