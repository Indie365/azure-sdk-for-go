//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcomplianceautomation

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

// SnapshotClient contains the methods for the Snapshot group.
// Don't use this type directly, use NewSnapshotClient() instead.
type SnapshotClient struct {
	internal *arm.Client
}

// NewSnapshotClient creates a new instance of SnapshotClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSnapshotClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SnapshotClient, error) {
	cl, err := arm.NewClient(moduleName+".SnapshotClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SnapshotClient{
		internal: cl,
	}
	return client, nil
}

// BeginDownload - Download compliance needs from snapshot, like: Compliance Report, Resource List.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-16-preview
//   - reportName - Report Name.
//   - snapshotName - Snapshot Name.
//   - parameters - Parameters for the query operation
//   - options - SnapshotClientBeginDownloadOptions contains the optional parameters for the SnapshotClient.BeginDownload method.
func (client *SnapshotClient) BeginDownload(ctx context.Context, reportName string, snapshotName string, parameters SnapshotDownloadRequest, options *SnapshotClientBeginDownloadOptions) (*runtime.Poller[SnapshotClientDownloadResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.download(ctx, reportName, snapshotName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SnapshotClientDownloadResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SnapshotClientDownloadResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Download - Download compliance needs from snapshot, like: Compliance Report, Resource List.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-16-preview
func (client *SnapshotClient) download(ctx context.Context, reportName string, snapshotName string, parameters SnapshotDownloadRequest, options *SnapshotClientBeginDownloadOptions) (*http.Response, error) {
	req, err := client.downloadCreateRequest(ctx, reportName, snapshotName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// downloadCreateRequest creates the Download request.
func (client *SnapshotClient) downloadCreateRequest(ctx context.Context, reportName string, snapshotName string, parameters SnapshotDownloadRequest, options *SnapshotClientBeginDownloadOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/snapshots/{snapshotName}/download"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get the AppComplianceAutomation snapshot and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-16-preview
//   - reportName - Report Name.
//   - snapshotName - Snapshot Name.
//   - options - SnapshotClientGetOptions contains the optional parameters for the SnapshotClient.Get method.
func (client *SnapshotClient) Get(ctx context.Context, reportName string, snapshotName string, options *SnapshotClientGetOptions) (SnapshotClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, reportName, snapshotName, options)
	if err != nil {
		return SnapshotClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SnapshotClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SnapshotClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SnapshotClient) getCreateRequest(ctx context.Context, reportName string, snapshotName string, options *SnapshotClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/snapshots/{snapshotName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SnapshotClient) getHandleResponse(resp *http.Response) (SnapshotClientGetResponse, error) {
	result := SnapshotClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotResource); err != nil {
		return SnapshotClientGetResponse{}, err
	}
	return result, nil
}
