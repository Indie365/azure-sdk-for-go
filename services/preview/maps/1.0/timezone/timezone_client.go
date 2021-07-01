// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package timezone

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TimezoneClient contains the methods for the Timezone group.
// Don't use this type directly, use NewTimezoneClient() instead.
type TimezoneClient struct {
	con *Connection
	xmsClientID *string
}

// NewTimezoneClient creates a new instance of TimezoneClient with the specified values.
func NewTimezoneClient(con *Connection, xmsClientID *string) *TimezoneClient {
	return &TimezoneClient{con: con, xmsClientID: xmsClientID}
}

// GetTimezoneByCoordinates - Time Zone by Coordinates
// Applies to: S0 and S1 pricing tiers.
// This API returns current, historical, and future time zone information for a specified latitude-longitude pair. In addition, the API provides sunset
// and sunrise times for a given location.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TimezoneClient) GetTimezoneByCoordinates(ctx context.Context, formatParam ResponseFormat, query string, options *TimezoneGetTimezoneByCoordinatesOptions) (TimezoneByCoordinatesResultResponse, error) {
	req, err := client.getTimezoneByCoordinatesCreateRequest(ctx, formatParam, query, options)
	if err != nil {
		return TimezoneByCoordinatesResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimezoneByCoordinatesResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimezoneByCoordinatesResultResponse{}, client.getTimezoneByCoordinatesHandleError(resp)
	}
	return client.getTimezoneByCoordinatesHandleResponse(resp)
}

// getTimezoneByCoordinatesCreateRequest creates the GetTimezoneByCoordinates request.
func (client *TimezoneClient) getTimezoneByCoordinatesCreateRequest(ctx context.Context, formatParam ResponseFormat, query string, options *TimezoneGetTimezoneByCoordinatesOptions) (*azcore.Request, error) {
	urlPath := "/timezone/byCoordinates/{format}"
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
	if options != nil && options.Options != nil {
		reqQP.Set("options", string(*options.Options))
	}
	if options != nil && options.TimeStamp != nil {
		reqQP.Set("timeStamp", options.TimeStamp.Format(time.RFC3339Nano))
	}
	if options != nil && options.TransitionsFrom != nil {
		reqQP.Set("transitionsFrom", options.TransitionsFrom.Format(time.RFC3339Nano))
	}
	if options != nil && options.TransitionsYears != nil {
		reqQP.Set("transitionsYears", strconv.FormatInt(int64(*options.TransitionsYears), 10))
	}
	reqQP.Set("query", query)
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	if options != nil && options.AcceptLanguage != nil {
		req.Header.Set("Accept-Language", *options.AcceptLanguage)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTimezoneByCoordinatesHandleResponse handles the GetTimezoneByCoordinates response.
func (client *TimezoneClient) getTimezoneByCoordinatesHandleResponse(resp *azcore.Response) (TimezoneByCoordinatesResultResponse, error) {
	var val *TimezoneByCoordinatesResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TimezoneByCoordinatesResultResponse{}, err
	}
return TimezoneByCoordinatesResultResponse{RawResponse: resp.Response, TimezoneByCoordinatesResult: val}, nil
}

// getTimezoneByCoordinatesHandleError handles the GetTimezoneByCoordinates error response.
func (client *TimezoneClient) getTimezoneByCoordinatesHandleError(resp *azcore.Response) error {
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

// GetTimezoneByID - Time Zone by Id
// Applies to: S0 and S1 pricing tiers.
// This API returns current, historical, and future time zone information for the specified IANA time zone ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TimezoneClient) GetTimezoneByID(ctx context.Context, formatParam ResponseFormat, query string, options *TimezoneGetTimezoneByIDOptions) (TimezoneByIDResultResponse, error) {
	req, err := client.getTimezoneByIDCreateRequest(ctx, formatParam, query, options)
	if err != nil {
		return TimezoneByIDResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimezoneByIDResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimezoneByIDResultResponse{}, client.getTimezoneByIDHandleError(resp)
	}
	return client.getTimezoneByIDHandleResponse(resp)
}

// getTimezoneByIDCreateRequest creates the GetTimezoneByID request.
func (client *TimezoneClient) getTimezoneByIDCreateRequest(ctx context.Context, formatParam ResponseFormat, query string, options *TimezoneGetTimezoneByIDOptions) (*azcore.Request, error) {
	urlPath := "/timezone/byId/{format}"
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
	if options != nil && options.Options != nil {
		reqQP.Set("options", string(*options.Options))
	}
	if options != nil && options.TimeStamp != nil {
		reqQP.Set("timeStamp", options.TimeStamp.Format(time.RFC3339Nano))
	}
	if options != nil && options.TransitionsFrom != nil {
		reqQP.Set("transitionsFrom", options.TransitionsFrom.Format(time.RFC3339Nano))
	}
	if options != nil && options.TransitionsYears != nil {
		reqQP.Set("transitionsYears", strconv.FormatInt(int64(*options.TransitionsYears), 10))
	}
	reqQP.Set("query", query)
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	if options != nil && options.AcceptLanguage != nil {
		req.Header.Set("Accept-Language", *options.AcceptLanguage)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTimezoneByIDHandleResponse handles the GetTimezoneByID response.
func (client *TimezoneClient) getTimezoneByIDHandleResponse(resp *azcore.Response) (TimezoneByIDResultResponse, error) {
	var val *TimezoneByIDResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TimezoneByIDResultResponse{}, err
	}
return TimezoneByIDResultResponse{RawResponse: resp.Response, TimezoneByIDResult: val}, nil
}

// getTimezoneByIDHandleError handles the GetTimezoneByID error response.
func (client *TimezoneClient) getTimezoneByIDHandleError(resp *azcore.Response) error {
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

// GetTimezoneEnumIANA - Enumerate IANA Time Zones
// Applies to: S0 and S1 pricing tiers.
// This API returns a full list of IANA time zone IDs. Updates to the IANA service will be reflected in the system within one day.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TimezoneClient) GetTimezoneEnumIANA(ctx context.Context, formatParam ResponseFormat, options *TimezoneGetTimezoneEnumIANAOptions) (IanaIDArrayResponse, error) {
	req, err := client.getTimezoneEnumIANACreateRequest(ctx, formatParam, options)
	if err != nil {
		return IanaIDArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IanaIDArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IanaIDArrayResponse{}, client.getTimezoneEnumIANAHandleError(resp)
	}
	return client.getTimezoneEnumIANAHandleResponse(resp)
}

// getTimezoneEnumIANACreateRequest creates the GetTimezoneEnumIANA request.
func (client *TimezoneClient) getTimezoneEnumIANACreateRequest(ctx context.Context, formatParam ResponseFormat, options *TimezoneGetTimezoneEnumIANAOptions) (*azcore.Request, error) {
	urlPath := "/timezone/enumIana/{format}"
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
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTimezoneEnumIANAHandleResponse handles the GetTimezoneEnumIANA response.
func (client *TimezoneClient) getTimezoneEnumIANAHandleResponse(resp *azcore.Response) (IanaIDArrayResponse, error) {
	var val []*IanaID
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IanaIDArrayResponse{}, err
	}
return IanaIDArrayResponse{RawResponse: resp.Response, IanaIDArray: val}, nil
}

// getTimezoneEnumIANAHandleError handles the GetTimezoneEnumIANA error response.
func (client *TimezoneClient) getTimezoneEnumIANAHandleError(resp *azcore.Response) error {
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

// GetTimezoneEnumWindows - Enumerate Windows Time Zones
// Applies to: S0 and S1 pricing tiers.
// This API returns a full list of Windows Time Zone IDs.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TimezoneClient) GetTimezoneEnumWindows(ctx context.Context, formatParam ResponseFormat, options *TimezoneGetTimezoneEnumWindowsOptions) (TimezoneEnumWindowArrayResponse, error) {
	req, err := client.getTimezoneEnumWindowsCreateRequest(ctx, formatParam, options)
	if err != nil {
		return TimezoneEnumWindowArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimezoneEnumWindowArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimezoneEnumWindowArrayResponse{}, client.getTimezoneEnumWindowsHandleError(resp)
	}
	return client.getTimezoneEnumWindowsHandleResponse(resp)
}

// getTimezoneEnumWindowsCreateRequest creates the GetTimezoneEnumWindows request.
func (client *TimezoneClient) getTimezoneEnumWindowsCreateRequest(ctx context.Context, formatParam ResponseFormat, options *TimezoneGetTimezoneEnumWindowsOptions) (*azcore.Request, error) {
	urlPath := "/timezone/enumWindows/{format}"
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
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTimezoneEnumWindowsHandleResponse handles the GetTimezoneEnumWindows response.
func (client *TimezoneClient) getTimezoneEnumWindowsHandleResponse(resp *azcore.Response) (TimezoneEnumWindowArrayResponse, error) {
	var val []*TimezoneEnumWindow
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TimezoneEnumWindowArrayResponse{}, err
	}
return TimezoneEnumWindowArrayResponse{RawResponse: resp.Response, TimezoneEnumWindowArray: val}, nil
}

// getTimezoneEnumWindowsHandleError handles the GetTimezoneEnumWindows error response.
func (client *TimezoneClient) getTimezoneEnumWindowsHandleError(resp *azcore.Response) error {
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

// GetTimezoneIANAVersion - Time Zone IANA Version
// Applies to: S0 and S1 pricing tiers.
// This API returns the current IANA version number.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TimezoneClient) GetTimezoneIANAVersion(ctx context.Context, formatParam ResponseFormat, options *TimezoneGetTimezoneIANAVersionOptions) (TimezoneIanaVersionResultResponse, error) {
	req, err := client.getTimezoneIANAVersionCreateRequest(ctx, formatParam, options)
	if err != nil {
		return TimezoneIanaVersionResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimezoneIanaVersionResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimezoneIanaVersionResultResponse{}, client.getTimezoneIANAVersionHandleError(resp)
	}
	return client.getTimezoneIANAVersionHandleResponse(resp)
}

// getTimezoneIANAVersionCreateRequest creates the GetTimezoneIANAVersion request.
func (client *TimezoneClient) getTimezoneIANAVersionCreateRequest(ctx context.Context, formatParam ResponseFormat, options *TimezoneGetTimezoneIANAVersionOptions) (*azcore.Request, error) {
	urlPath := "/timezone/ianaVersion/{format}"
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
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTimezoneIANAVersionHandleResponse handles the GetTimezoneIANAVersion response.
func (client *TimezoneClient) getTimezoneIANAVersionHandleResponse(resp *azcore.Response) (TimezoneIanaVersionResultResponse, error) {
	var val *TimezoneIanaVersionResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TimezoneIanaVersionResultResponse{}, err
	}
return TimezoneIanaVersionResultResponse{RawResponse: resp.Response, TimezoneIanaVersionResult: val}, nil
}

// getTimezoneIANAVersionHandleError handles the GetTimezoneIANAVersion error response.
func (client *TimezoneClient) getTimezoneIANAVersionHandleError(resp *azcore.Response) error {
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

// GetTimezoneWindowsToIANA - Windows to IANA Time Zone
// Applies to: S0 and S1 pricing tiers.
// This API returns a corresponding IANA ID, given a valid Windows Time Zone ID. Multiple IANA IDs may be returned for a single Windows ID. It is possible
// to narrow these results by adding an optional
// territory parameter.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TimezoneClient) GetTimezoneWindowsToIANA(ctx context.Context, formatParam ResponseFormat, query string, options *TimezoneGetTimezoneWindowsToIANAOptions) (IanaIDArrayResponse, error) {
	req, err := client.getTimezoneWindowsToIANACreateRequest(ctx, formatParam, query, options)
	if err != nil {
		return IanaIDArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IanaIDArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IanaIDArrayResponse{}, client.getTimezoneWindowsToIANAHandleError(resp)
	}
	return client.getTimezoneWindowsToIANAHandleResponse(resp)
}

// getTimezoneWindowsToIANACreateRequest creates the GetTimezoneWindowsToIANA request.
func (client *TimezoneClient) getTimezoneWindowsToIANACreateRequest(ctx context.Context, formatParam ResponseFormat, query string, options *TimezoneGetTimezoneWindowsToIANAOptions) (*azcore.Request, error) {
	urlPath := "/timezone/windowsToIana/{format}"
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
	reqQP.Set("query", query)
	if options != nil && options.Territory != nil {
		reqQP.Set("territory", *options.Territory)
	}
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTimezoneWindowsToIANAHandleResponse handles the GetTimezoneWindowsToIANA response.
func (client *TimezoneClient) getTimezoneWindowsToIANAHandleResponse(resp *azcore.Response) (IanaIDArrayResponse, error) {
	var val []*IanaID
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IanaIDArrayResponse{}, err
	}
return IanaIDArrayResponse{RawResponse: resp.Response, IanaIDArray: val}, nil
}

// getTimezoneWindowsToIANAHandleError handles the GetTimezoneWindowsToIANA error response.
func (client *TimezoneClient) getTimezoneWindowsToIANAHandleError(resp *azcore.Response) error {
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

