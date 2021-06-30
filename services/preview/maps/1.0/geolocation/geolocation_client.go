// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package geolocation

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type geolocationClient struct {
	con *connection
	xmsClientID *string
}

// GetIPToLocationPreview - Applies to: S0 and S1 pricing tiers.
// This service will return the ISO country code for the provided IP address. Developers can use this information to block or alter certain content based
// on geographical locations where the application
// is being viewed from.
// Note: This service returns results from IANA and does not necessarily reflect the views of Microsoft Corporation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *geolocationClient) GetIPToLocationPreview(ctx context.Context, formatParam ResponseFormat, ip string, options *GeolocationGetIPToLocationPreviewOptions) (IPAddressToLocationResultResponse, error) {
	req, err := client.getIPToLocationPreviewCreateRequest(ctx, formatParam, ip, options)
	if err != nil {
		return IPAddressToLocationResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IPAddressToLocationResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IPAddressToLocationResultResponse{}, client.getIPToLocationPreviewHandleError(resp)
	}
	return client.getIPToLocationPreviewHandleResponse(resp)
}

// getIPToLocationPreviewCreateRequest creates the GetIPToLocationPreview request.
func (client *geolocationClient) getIPToLocationPreviewCreateRequest(ctx context.Context, formatParam ResponseFormat, ip string, options *GeolocationGetIPToLocationPreviewOptions) (*azcore.Request, error) {
	urlPath := "/geolocation/ip/{format}"
	if formatParam == "" {
		return nil, errors.New("parameter formatParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{format}", url.PathEscape(string(formatParam)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "1.0")
	reqQP.Set("ip", ip)
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getIPToLocationPreviewHandleResponse handles the GetIPToLocationPreview response.
func (client *geolocationClient) getIPToLocationPreviewHandleResponse(resp *azcore.Response) (IPAddressToLocationResultResponse, error) {
	var val *IPAddressToLocationResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IPAddressToLocationResultResponse{}, err
	}
return IPAddressToLocationResultResponse{RawResponse: resp.Response, IPAddressToLocationResult: val}, nil
}

// getIPToLocationPreviewHandleError handles the GetIPToLocationPreview error response.
func (client *geolocationClient) getIPToLocationPreviewHandleError(resp *azcore.Response) error {
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

