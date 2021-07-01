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
	"time"
)

// DatasetClient contains the methods for the Dataset group.
// Don't use this type directly, use NewDatasetClient() instead.
type DatasetClient struct {
	con *Connection
}

// NewDatasetClient creates a new instance of DatasetClient with the specified values.
func NewDatasetClient(con *Connection) *DatasetClient {
	return &DatasetClient{con: con}
}

// BeginCreate - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create a dataset from data that was uploaded to the Azure Maps Data Service and converted using the Azure Maps Conversion
// Service.
// You can use this API in a scenario like uploading a DWG zip package for a building, converting the zip package using the Azure Maps Conversion Service,
// and creating a dataset from the converted zip
// package. The created dataset can be used to create tilesets using the Azure Maps Tileset Service and can be queried via the Azure Maps WFS Service.
// SUBMIT CREATE REQUEST To create your dataset, you will use a POST request where the conversionId query parameter is an ID that represents the converted
// DWG zip package, the datasetId parameter will be
// the ID of a previously created dataset to append with the current dataset and, optionally, the description query parameter will contain a description
// (if description is not provided a default
// description will be given).
// The Create API is a long-running request [https://aka.ms/am-creator-lrt-v2].
// If the operation fails it returns the *ErrorResponse error type.
func (client *DatasetClient) BeginCreate(ctx context.Context, conversionID string, options *DatasetBeginCreateOptions) (LongRunningOperationResultPollerResponse, error) {
	resp, err := client.create(ctx, conversionID, options)
	if err != nil {
		return LongRunningOperationResultPollerResponse{}, err
	}
	result := LongRunningOperationResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("DatasetClient.Create",resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return LongRunningOperationResultPollerResponse{}, err
	}
	poller := &longRunningOperationResultPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LongRunningOperationResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new LongRunningOperationResultPoller from the specified resume token.
// token - The value must come from a previous call to LongRunningOperationResultPoller.ResumeToken().
func (client *DatasetClient) ResumeCreate(ctx context.Context, token string) (LongRunningOperationResultPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("DatasetClient.Create",token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return LongRunningOperationResultPollerResponse{}, err
	}
	poller := &longRunningOperationResultPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LongRunningOperationResultPollerResponse{}, err
	}
	result := LongRunningOperationResultPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LongRunningOperationResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create a dataset from data that was uploaded to the Azure Maps Data Service and converted using the Azure Maps Conversion
// Service.
// You can use this API in a scenario like uploading a DWG zip package for a building, converting the zip package using the Azure Maps Conversion Service,
// and creating a dataset from the converted zip
// package. The created dataset can be used to create tilesets using the Azure Maps Tileset Service and can be queried via the Azure Maps WFS Service.
// SUBMIT CREATE REQUEST To create your dataset, you will use a POST request where the conversionId query parameter is an ID that represents the converted
// DWG zip package, the datasetId parameter will be
// the ID of a previously created dataset to append with the current dataset and, optionally, the description query parameter will contain a description
// (if description is not provided a default
// description will be given).
// The Create API is a long-running request [https://aka.ms/am-creator-lrt-v2].
// If the operation fails it returns the *ErrorResponse error type.
func (client *DatasetClient) create(ctx context.Context, conversionID string, options *DatasetBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, conversionID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	 return resp, nil
}

// createCreateRequest creates the Create request.
func (client *DatasetClient) createCreateRequest(ctx context.Context, conversionID string, options *DatasetBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/datasets"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	reqQP.Set("conversionId", conversionID)
	if options != nil && options.DatasetID != nil {
		reqQP.Set("datasetId", *options.DatasetID)
	}
	if options != nil && options.DescriptionDataset != nil {
		reqQP.Set("description", *options.DescriptionDataset)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleError handles the Create error response.
func (client *DatasetClient) createHandleError(resp *azcore.Response) error {
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
// You can also use this API to delete old/unused datasets to create space for new Creator content.
// SUBMIT DELETE REQUEST To delete your content you will issue a DELETE request where the path will contain the datasetId of the dataset to delete.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DatasetClient) Delete(ctx context.Context, datasetID string, options *DatasetDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, datasetID, options)
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
func (client *DatasetClient) deleteCreateRequest(ctx context.Context, datasetID string, options *DatasetDeleteOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetId}"
	if datasetID == "" {
		return nil, errors.New("parameter datasetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetId}", url.PathEscape(datasetID))
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
func (client *DatasetClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps]
// article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch a previously successfully created dataset.
// SUBMIT GET DETAILS REQUEST To get the details for a previously created dataset, you will issue a GET request with the datasetId in the path.
// GET DETAILS RESPONSE The Get Details API returns the details for a dataset in json format. The response contains the following fields (if they are not
// null or empty):
// > created - The timestamp the dataset was created. datasetId - The id for the dataset. description - The description for the dataset. datasetSources
// - The source data that was used when the create
// request was issued. ontology - The source ontology [https://docs.microsoft.com/en-us/azure/azure-maps/creator-facility-ontology] that was used in the
// conversion service for the input data.
// The datasetSources describes the source data that was used when the create request was issued and contains the following elements (if they are not null
// or empty):
// > conversionIds - The list of conversionId (null if none were provided). appendDatasetId - The datasetId that was used for an append operation (null
// if none was used). featureCounts - The counts for
// each feature type in the dataset.
// Here's a sample response returning the timestamp, datasetId, description, datasetSources, and ontology of a dataset resource:
// { "timestamp": "2020-01-01T22:50:48.123Z", "datasetId": "f6495f62-94f8-0ec2-c252-45626f82fcb2", "description": "Some description or comment for the dataset.",
// "datasetSources": { "conversionIds": [
// "15d21452-c9bb-27b6-5e79-743ca5c3205d" ], }, "ontology": "facility-2.0", "featureCounts": { "directoryInfo": 2, "category": 10, "facility": 1, "level":
// 3, "unit": 183, "zone": 3,
// "verticalPenetration": 6, "opening": 48, "areaElement": 108 } }
// If the operation fails it returns the *ErrorResponse error type.
func (client *DatasetClient) Get(ctx context.Context, datasetID string, options *DatasetGetOptions) (DatasetDetailInfoResponse, error) {
	req, err := client.getCreateRequest(ctx, datasetID, options)
	if err != nil {
		return DatasetDetailInfoResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DatasetDetailInfoResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DatasetDetailInfoResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatasetClient) getCreateRequest(ctx context.Context, datasetID string, options *DatasetGetOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetId}"
	if datasetID == "" {
		return nil, errors.New("parameter datasetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetId}", url.PathEscape(datasetID))
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

// getHandleResponse handles the Get response.
func (client *DatasetClient) getHandleResponse(resp *azcore.Response) (DatasetDetailInfoResponse, error) {
	var val *DatasetDetailInfo
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DatasetDetailInfoResponse{}, err
	}
return DatasetDetailInfoResponse{RawResponse: resp.Response, DatasetDetailInfo: val}, nil
}

// getHandleError handles the Get error response.
func (client *DatasetClient) getHandleError(resp *azcore.Response) error {
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

// GetOperation - This API allows the caller to view the current progress of a dataset operation and the path is obtained from a call to the Create API.
// SUBMIT OPERATIONS REQUEST To view the current progress of a dataset operation, you will use a GET request where the operationId given the path is the
// ID that represents the operation.
// OPERATION RESPONSE While in progress, a 200-OK http status code will be returned with no extra headers. If the operation succeeds, a 200-OK http status
// code with Resource-Location header will be
// returned.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DatasetClient) GetOperation(ctx context.Context, operationID string, options *DatasetGetOperationOptions) (LongRunningOperationResultResponse, error) {
	req, err := client.getOperationCreateRequest(ctx, operationID, options)
	if err != nil {
		return LongRunningOperationResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LongRunningOperationResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LongRunningOperationResultResponse{}, client.getOperationHandleError(resp)
	}
	return client.getOperationHandleResponse(resp)
}

// getOperationCreateRequest creates the GetOperation request.
func (client *DatasetClient) getOperationCreateRequest(ctx context.Context, operationID string, options *DatasetGetOperationOptions) (*azcore.Request, error) {
	urlPath := "/datasets/operations/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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

// getOperationHandleResponse handles the GetOperation response.
func (client *DatasetClient) getOperationHandleResponse(resp *azcore.Response) (LongRunningOperationResultResponse, error) {
	var val *LongRunningOperationResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LongRunningOperationResultResponse{}, err
	}
	result := LongRunningOperationResultResponse{RawResponse: resp.Response, LongRunningOperationResult: val}
	if val := resp.Header.Get("Resource-Location"); val != "" {
		result.ResourceLocation = &val
	}
	return result, nil
}

// getOperationHandleError handles the GetOperation error response.
func (client *DatasetClient) getOperationHandleError(resp *azcore.Response) error {
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
// This API allows the caller to fetch a list of all previously successfully created datasets.
// SUBMIT LIST REQUEST To list all your datasets, you will issue a GET request with no additional parameters.
// LIST DATA RESPONSE The List API returns the complete list of all datasets in json format. The response contains the following fields (if they are not
// null or empty):
// > created - The timestamp the dataset was created. datasetId - The id for the dataset. description - The description for the dataset. datasetSources
// - The source data that was used when the create
// request was issued. ontology - The source ontology [https://docs.microsoft.com/en-us/azure/azure-maps/creator-facility-ontology] that was used in the
// conversion service for the input data.
// The datasetSources describes the source data that was used when the create request was issued and contains the following elements (if they are not null
// or empty):
// > conversionIds - The list of conversionId (null if none were provided). appendDatasetId - The datasetId that was used for an append operation (null
// if none was used). featureCounts - The counts for
// each feature type in the dataset.
// Here's a sample response returning the timestamp, datasetId, description, datasetSources, and ontology of 3 dataset resources:
// { "datasets": [ { "timestamp": "2020-01-01T22:50:48.123Z", "datasetId": "f6495f62-94f8-0ec2-c252-45626f82fcb2", "description": "Some description or comment
// for the dataset.", "datasetSources": {
// "conversionIds": [ "15d21452-c9bb-27b6-5e79-743ca5c3205d" ], }, "ontology": "facility-2.0", "featureCounts": { "directoryInfo": 2, "category": 10, "facility":
// 1, "level": 3, "unit": 183, "zone": 3,
// "verticalPenetration": 6, "opening": 48, "areaElement": 108 } }, { "timestamp": "2020-01-01T22:57:53.123Z", "datasetId": "8b1288fa-1958-4a2b-b68e-13a7i5af7d7c",
// "description": "Create from upload
// '0c1288fa-2058-4a1b-b68d-13a5f5af7d7c'.", "datasetSources": { "conversionIds": [ "0c1288fa-2058-4a1b-b68d-13a5f5af7d7c" ], "appendDatasetId": "46d1edb6-d29e-4786-9589-dbd4efd7a977"
// }, "ontology":
// "facility-2.0", "featureCounts": { "directoryInfo": 2, "category": 10, "facility": 1, "level": 3, "unit": 183, "zone": 3, "verticalPenetration": 6, "opening":
// 48, "areaElement": 108 } } ] }
// If the operation fails it returns the *ErrorResponse error type.
func (client *DatasetClient) List(options *DatasetListOptions) (DatasetListResponsePager) {
	return &datasetListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp DatasetListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DatasetListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *DatasetClient) listCreateRequest(ctx context.Context, options *DatasetListOptions) (*azcore.Request, error) {
	urlPath := "/datasets"
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
func (client *DatasetClient) listHandleResponse(resp *azcore.Response) (DatasetListResponseResponse, error) {
	var val *DatasetListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DatasetListResponseResponse{}, err
	}
return DatasetListResponseResponse{RawResponse: resp.Response, DatasetListResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *DatasetClient) listHandleError(resp *azcore.Response) error {
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

