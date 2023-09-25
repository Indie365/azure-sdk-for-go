//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

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

// ValidationsClient contains the methods for the Validations group.
// Don't use this type directly, use NewValidationsClient() instead.
type ValidationsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewValidationsClient creates a new instance of ValidationsClient with the specified values.
//   - subscriptionID - Microsoft Azure subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewValidationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ValidationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ValidationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ValidationsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// ValidateOrganization - Organization Validate proxy resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - Resource group name
//   - organizationName - Organization resource name
//   - body - Organization resource model
//   - options - ValidationsClientValidateOrganizationOptions contains the optional parameters for the ValidationsClient.ValidateOrganization
//     method.
func (client *ValidationsClient) ValidateOrganization(ctx context.Context, resourceGroupName string, organizationName string, body OrganizationResource, options *ValidationsClientValidateOrganizationOptions) (ValidationsClientValidateOrganizationResponse, error) {
	var err error
	req, err := client.validateOrganizationCreateRequest(ctx, resourceGroupName, organizationName, body, options)
	if err != nil {
		return ValidationsClientValidateOrganizationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ValidationsClientValidateOrganizationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ValidationsClientValidateOrganizationResponse{}, err
	}
	resp, err := client.validateOrganizationHandleResponse(httpResp)
	return resp, err
}

// validateOrganizationCreateRequest creates the ValidateOrganization request.
func (client *ValidationsClient) validateOrganizationCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, body OrganizationResource, options *ValidationsClientValidateOrganizationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/validations/{organizationName}/orgvalidate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
	return req, nil
}

// validateOrganizationHandleResponse handles the ValidateOrganization response.
func (client *ValidationsClient) validateOrganizationHandleResponse(resp *http.Response) (ValidationsClientValidateOrganizationResponse, error) {
	result := ValidationsClientValidateOrganizationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResource); err != nil {
		return ValidationsClientValidateOrganizationResponse{}, err
	}
	return result, nil
}

