//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// FieldsClient contains the methods for the Fields group.
// Don't use this type directly, use NewFieldsClient() instead.
type FieldsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewFieldsClient creates a new instance of FieldsClient with the specified values.
func NewFieldsClient(con *arm.Connection, subscriptionID string) *FieldsClient {
	return &FieldsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListByType - Retrieve a list of fields of a given type identified by module name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FieldsClient) ListByType(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *FieldsListByTypeOptions) (FieldsListByTypeResponse, error) {
	req, err := client.listByTypeCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, typeName, options)
	if err != nil {
		return FieldsListByTypeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FieldsListByTypeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FieldsListByTypeResponse{}, client.listByTypeHandleError(resp)
	}
	return client.listByTypeHandleResponse(resp)
}

// listByTypeCreateRequest creates the ListByType request.
func (client *FieldsClient) listByTypeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *FieldsListByTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/types/{typeName}/fields"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if typeName == "" {
		return nil, errors.New("parameter typeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{typeName}", url.PathEscape(typeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTypeHandleResponse handles the ListByType response.
func (client *FieldsClient) listByTypeHandleResponse(resp *http.Response) (FieldsListByTypeResponse, error) {
	result := FieldsListByTypeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TypeFieldListResult); err != nil {
		return FieldsListByTypeResponse{}, err
	}
	return result, nil
}

// listByTypeHandleError handles the ListByType error response.
func (client *FieldsClient) listByTypeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
