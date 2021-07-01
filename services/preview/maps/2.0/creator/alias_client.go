// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package creator

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AliasClient contains the methods for the Alias group.
// Don't use this type directly, use NewAliasClient() instead.
type AliasClient struct {
	con *Connection
}

// NewAliasClient creates a new instance of AliasClient with the specified values.
func NewAliasClient(con *Connection) *AliasClient {
	return &AliasClient{con: con}
}

// Assign - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to assign an alias to reference a resource.
// SUBMIT ASSIGN REQUEST To assign your alias to a resource, you will use a PUT request with the aliasId in the path and the creatorDataItemId passed as
// a query parameter.
// ASSIGN ALIAS RESPONSE The Assign API returns a HTTP 200 OK response with the updated alias resource in the body, if the alias was assigned successfully.
// A sample of the assign response is
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId": "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14",
// "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns the *ErrorResponse error type.
func (client *AliasClient) Assign(ctx context.Context, aliasID string, creatorDataItemID string, options *AliasAssignOptions) (AliasListItemResponse, error) {
	req, err := client.assignCreateRequest(ctx, aliasID, creatorDataItemID, options)
	if err != nil {
		return AliasListItemResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AliasListItemResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AliasListItemResponse{}, client.assignHandleError(resp)
	}
	return client.assignHandleResponse(resp)
}

// assignCreateRequest creates the Assign request.
func (client *AliasClient) assignCreateRequest(ctx context.Context, aliasID string, creatorDataItemID string, options *AliasAssignOptions) (*azcore.Request, error) {
	urlPath := "/aliases/{aliasId}"
	if aliasID == "" {
		return nil, errors.New("parameter aliasID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{aliasId}", url.PathEscape(aliasID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	reqQP.Set("creatorDataItemId", creatorDataItemID)
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// assignHandleResponse handles the Assign response.
func (client *AliasClient) assignHandleResponse(resp *azcore.Response) (AliasListItemResponse, error) {
	var val *AliasListItem
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AliasListItemResponse{}, err
	}
return AliasListItemResponse{RawResponse: resp.Response, AliasListItem: val}, nil
}

// assignHandleError handles the Assign error response.
func (client *AliasClient) assignHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Create - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create an alias. You can also assign the alias during the create request. An alias can reference an ID generated by a creator
// service, but cannot reference another alias
// ID.
// SUBMIT CREATE REQUEST To create your alias, you will use a POST request. If you would like to assign the alias during the creation, you will pass the
// resourceId query parameter.
// CREATE ALIAS RESPONSE The Create API returns a HTTP 201 Created response with the alias resource in the body.
// A sample response from creating an alias:
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId": "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14",
// "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns the *ErrorResponse error type.
func (client *AliasClient) Create(ctx context.Context, options *AliasCreateOptions) (AliasesCreateResponseResponse, error) {
	req, err := client.createCreateRequest(ctx, options)
	if err != nil {
		return AliasesCreateResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AliasesCreateResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AliasesCreateResponseResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *AliasClient) createCreateRequest(ctx context.Context, options *AliasCreateOptions) (*azcore.Request, error) {
	urlPath := "/aliases"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	if options != nil && options.CreatorDataItemID != nil {
		reqQP.Set("creatorDataItemId", *options.CreatorDataItemID)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *AliasClient) createHandleResponse(resp *azcore.Response) (AliasesCreateResponseResponse, error) {
	var val *AliasesCreateResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AliasesCreateResponseResponse{}, err
	}
	result := AliasesCreateResponseResponse{RawResponse: resp.Response, AliasesCreateResponse: val}
	if val := resp.Header.Get("Access-Control-Expose-Headers"); val != "" {
		result.AccessControlExposeHeaders = &val
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *AliasClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to delete a previously created alias. You can also use this API to delete old/unused aliases to create space for new content.This
// API does not delete the references
// resource, only the alias referencing the resource.
// SUBMIT DELETE REQUEST To delete your alias you will issue a DELETE request where the path will contain the aliasId of the alias to delete.
// DELETE ALIAS RESPONSE The Delete API returns a HTTP 204 No Content response with an empty body, if the alias was deleted successfully.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AliasClient) Delete(ctx context.Context, aliasID string, options *AliasDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, aliasID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AliasClient) deleteCreateRequest(ctx context.Context, aliasID string, options *AliasDeleteOptions) (*azcore.Request, error) {
	urlPath := "/aliases/{aliasId}"
	if aliasID == "" {
		return nil, errors.New("parameter aliasID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{aliasId}", url.PathEscape(aliasID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AliasClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetDetails - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch the details of a previously created alias.
// SUBMIT GET DETAILS REQUEST To get the details of your alias, you will issue a GET request with the aliasId in the path.
// GET DETAILS RESPONSE The Get Details API returns the previously created alias in json format. The response contains the following details for the alias
// resource:
// > createdTimestamp - The timestamp that the alias was created. aliasId - The id for the alias. creatorDataItemId - The id for the creator data item that
// this alias references (could be null if the
// alias has not been assigned). lastUpdatedTimestamp - The last time the alias was assigned to a resource.
// Here's a sample response:
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId": "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14",
// "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns the *ErrorResponse error type.
func (client *AliasClient) GetDetails(ctx context.Context, aliasID string, options *AliasGetDetailsOptions) (AliasListItemResponse, error) {
	req, err := client.getDetailsCreateRequest(ctx, aliasID, options)
	if err != nil {
		return AliasListItemResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AliasListItemResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AliasListItemResponse{}, client.getDetailsHandleError(resp)
	}
	return client.getDetailsHandleResponse(resp)
}

// getDetailsCreateRequest creates the GetDetails request.
func (client *AliasClient) getDetailsCreateRequest(ctx context.Context, aliasID string, options *AliasGetDetailsOptions) (*azcore.Request, error) {
	urlPath := "/aliases/{aliasId}"
	if aliasID == "" {
		return nil, errors.New("parameter aliasID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{aliasId}", url.PathEscape(aliasID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDetailsHandleResponse handles the GetDetails response.
func (client *AliasClient) getDetailsHandleResponse(resp *azcore.Response) (AliasListItemResponse, error) {
	var val *AliasListItem
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AliasListItemResponse{}, err
	}
return AliasListItemResponse{RawResponse: resp.Response, AliasListItem: val}, nil
}

// getDetailsHandleError handles the GetDetails error response.
func (client *AliasClient) getDetailsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch a list of all previously successfully created aliases.
// SUBMIT LIST REQUEST To list all your aliases, you will issue a GET request with no additional parameters.
// LIST DATA RESPONSE The List API returns the complete list of all aliases in json format. The response contains the following details for each alias resource:
// > createdTimestamp - The timestamp that the alias was created. Format yyyy-MM-ddTHH:mm:ss.sssZ aliasId - The id for the alias. creatorDataItemId - The
// id for the creator data item that this alias
// references (could be null if the alias has not been assigned). lastUpdatedTimestamp - The last time the alias was assigned to a resource. Format yyyy-MM-ddTHH:mm:ss.sssZ
// A sample response returning 2 alias resources:
// { "aliases": [ { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId": "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14",
// "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }, { "createdTimestamp": "2020-02-18T19:53:33.123Z", "aliasId": "1856dbfc-7a66-ee5a-bf8d-51dbfe1906f6", "creatorDataItemId":
// null, "lastUpdatedTimestamp":
// "2020-02-18T19:53:33.123Z" } ] }
// If the operation fails it returns the *ErrorResponse error type.
func (client *AliasClient) List(options *AliasListOptions) (AliasListResponsePager) {
	return &aliasListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp AliasListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AliasListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *AliasClient) listCreateRequest(ctx context.Context, options *AliasListOptions) (*azcore.Request, error) {
	urlPath := "/aliases"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AliasClient) listHandleResponse(resp *azcore.Response) (AliasListResponseResponse, error) {
	var val *AliasListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AliasListResponseResponse{}, err
	}
return AliasListResponseResponse{RawResponse: resp.Response, AliasListResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *AliasClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

