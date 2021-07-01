// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package traffic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TrafficClient contains the methods for the Traffic group.
// Don't use this type directly, use NewTrafficClient() instead.
type TrafficClient struct {
	con *Connection
	xmsClientID *string
}

// NewTrafficClient creates a new instance of TrafficClient with the specified values.
func NewTrafficClient(con *Connection, xmsClientID *string) *TrafficClient {
	return &TrafficClient{con: con, xmsClientID: xmsClientID}
}

// GetTrafficFlowSegment - Traffic Flow Segment
// Applies to: S0 and S1 pricing tiers.
// This service provides information about the speeds and travel times of the road fragment closest to the given coordinates. It is designed to work alongside
// the Flow layer of the Render Service to
// support clickable flow data visualizations. With this API, the client side can connect any place in the map with flow data on the closest road and present
// it to the user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TrafficClient) GetTrafficFlowSegment(ctx context.Context, formatParam TextFormat, style TrafficFlowSegmentStyle, zoom int32, query string, options *TrafficGetTrafficFlowSegmentOptions) (TrafficFlowSegmentResultResponse, error) {
	req, err := client.getTrafficFlowSegmentCreateRequest(ctx, formatParam, style, zoom, query, options)
	if err != nil {
		return TrafficFlowSegmentResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TrafficFlowSegmentResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TrafficFlowSegmentResultResponse{}, client.getTrafficFlowSegmentHandleError(resp)
	}
	return client.getTrafficFlowSegmentHandleResponse(resp)
}

// getTrafficFlowSegmentCreateRequest creates the GetTrafficFlowSegment request.
func (client *TrafficClient) getTrafficFlowSegmentCreateRequest(ctx context.Context, formatParam TextFormat, style TrafficFlowSegmentStyle, zoom int32, query string, options *TrafficGetTrafficFlowSegmentOptions) (*azcore.Request, error) {
	urlPath := "/traffic/flow/segment/{format}"
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
	reqQP.Set("style", string(style))
	reqQP.Set("zoom", strconv.FormatInt(int64(zoom), 10))
	reqQP.Set("query", query)
	if options != nil && options.Unit != nil {
		reqQP.Set("unit", string(*options.Unit))
	}
	if options != nil && options.Thickness != nil {
		reqQP.Set("thickness", strconv.FormatInt(int64(*options.Thickness), 10))
	}
	if options != nil && options.OpenLr != nil {
		reqQP.Set("openLr", strconv.FormatBool(*options.OpenLr))
	}
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTrafficFlowSegmentHandleResponse handles the GetTrafficFlowSegment response.
func (client *TrafficClient) getTrafficFlowSegmentHandleResponse(resp *azcore.Response) (TrafficFlowSegmentResultResponse, error) {
	var val *TrafficFlowSegmentResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TrafficFlowSegmentResultResponse{}, err
	}
return TrafficFlowSegmentResultResponse{RawResponse: resp.Response, TrafficFlowSegmentResult: val}, nil
}

// getTrafficFlowSegmentHandleError handles the GetTrafficFlowSegment error response.
func (client *TrafficClient) getTrafficFlowSegmentHandleError(resp *azcore.Response) error {
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

// GetTrafficFlowTile - Traffic Flow Tile
// Applies to: S0 and S1 pricing tiers.
// The Azure Flow Tile API serves 256 x 256 pixel tiles showing traffic flow. All tiles use the same grid system. Because the traffic tiles use transparent
// images, they can be layered on top of map tiles
// to create a compound display. The Flow tiles use colors to indicate either the speed of traffic on different road segments, or the difference between
// that speed and the free-flow speed on the road
// segment in question.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TrafficClient) GetTrafficFlowTile(ctx context.Context, formatParam TileFormat, style TrafficFlowTileStyle, zoom int32, xTileIndex int32, yTileIndex int32, options *TrafficGetTrafficFlowTileOptions) (TrafficGetTrafficFlowTileResponse, error) {
	req, err := client.getTrafficFlowTileCreateRequest(ctx, formatParam, style, zoom, xTileIndex, yTileIndex, options)
	if err != nil {
		return TrafficGetTrafficFlowTileResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TrafficGetTrafficFlowTileResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TrafficGetTrafficFlowTileResponse{}, client.getTrafficFlowTileHandleError(resp)
	}
	return client.getTrafficFlowTileHandleResponse(resp)
}

// getTrafficFlowTileCreateRequest creates the GetTrafficFlowTile request.
func (client *TrafficClient) getTrafficFlowTileCreateRequest(ctx context.Context, formatParam TileFormat, style TrafficFlowTileStyle, zoom int32, xTileIndex int32, yTileIndex int32, options *TrafficGetTrafficFlowTileOptions) (*azcore.Request, error) {
	urlPath := "/traffic/flow/tile/{format}"
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
	reqQP.Set("style", string(style))
	reqQP.Set("zoom", strconv.FormatInt(int64(zoom), 10))
	reqQP.Set("x", strconv.FormatInt(int64(xTileIndex), 10))
	reqQP.Set("y", strconv.FormatInt(int64(yTileIndex), 10))
	if options != nil && options.Thickness != nil {
		reqQP.Set("thickness", strconv.FormatInt(int64(*options.Thickness), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.SkipBodyDownload()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json, image/jpeg, image/png, image/pbf, application/vnd.mapbox-vector-tile")
	return req, nil
}

// getTrafficFlowTileHandleResponse handles the GetTrafficFlowTile response.
func (client *TrafficClient) getTrafficFlowTileHandleResponse(resp *azcore.Response) (TrafficGetTrafficFlowTileResponse, error) {
	result := TrafficGetTrafficFlowTileResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	return result, nil
}

// getTrafficFlowTileHandleError handles the GetTrafficFlowTile error response.
func (client *TrafficClient) getTrafficFlowTileHandleError(resp *azcore.Response) error {
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

// GetTrafficIncidentDetail - Traffic Incident Detail
// Applies to: S0 and S1 pricing tiers.
// This API provides information on traffic incidents inside a given bounding box, based on the current Traffic Model ID. The Traffic Model ID is available
// to grant synchronization of data between calls
// and API's. The Traffic Model ID is a key value for determining the currency of traffic incidents. It is updated every minute, and is valid for two minutes
// before it times out. It is used in rendering
// incident tiles [https://docs.microsoft.com/en-us/rest/api/maps/traffic/gettrafficincidenttile]. It can be obtained from the Viewport API
// [https://docs.microsoft.com/en-us/rest/api/maps/traffic/gettrafficincidentviewport].
// If the operation fails it returns the *ErrorResponse error type.
func (client *TrafficClient) GetTrafficIncidentDetail(ctx context.Context, formatParam TextFormat, style TrafficIncidentDetailStyle, boundingbox string, boundingZoom int32, trafficmodelid string, options *TrafficGetTrafficIncidentDetailOptions) (TrafficIncidentDetailResultResponse, error) {
	req, err := client.getTrafficIncidentDetailCreateRequest(ctx, formatParam, style, boundingbox, boundingZoom, trafficmodelid, options)
	if err != nil {
		return TrafficIncidentDetailResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TrafficIncidentDetailResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TrafficIncidentDetailResultResponse{}, client.getTrafficIncidentDetailHandleError(resp)
	}
	return client.getTrafficIncidentDetailHandleResponse(resp)
}

// getTrafficIncidentDetailCreateRequest creates the GetTrafficIncidentDetail request.
func (client *TrafficClient) getTrafficIncidentDetailCreateRequest(ctx context.Context, formatParam TextFormat, style TrafficIncidentDetailStyle, boundingbox string, boundingZoom int32, trafficmodelid string, options *TrafficGetTrafficIncidentDetailOptions) (*azcore.Request, error) {
	urlPath := "/traffic/incident/detail/{format}"
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
	reqQP.Set("style", string(style))
	reqQP.Set("boundingbox", boundingbox)
	reqQP.Set("boundingZoom", strconv.FormatInt(int64(boundingZoom), 10))
	reqQP.Set("trafficmodelid", trafficmodelid)
	if options != nil && options.Language != nil {
		reqQP.Set("language", *options.Language)
	}
	if options != nil && options.Projection != nil {
		reqQP.Set("projection", string(*options.Projection))
	}
	if options != nil && options.Geometries != nil {
		reqQP.Set("geometries", string(*options.Geometries))
	}
	if options != nil && options.ExpandCluster != nil {
		reqQP.Set("expandCluster", strconv.FormatBool(*options.ExpandCluster))
	}
	if options != nil && options.OriginalPosition != nil {
		reqQP.Set("originalPosition", strconv.FormatBool(*options.OriginalPosition))
	}
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTrafficIncidentDetailHandleResponse handles the GetTrafficIncidentDetail response.
func (client *TrafficClient) getTrafficIncidentDetailHandleResponse(resp *azcore.Response) (TrafficIncidentDetailResultResponse, error) {
	var val *TrafficIncidentDetailResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TrafficIncidentDetailResultResponse{}, err
	}
return TrafficIncidentDetailResultResponse{RawResponse: resp.Response, TrafficIncidentDetailResult: val}, nil
}

// getTrafficIncidentDetailHandleError handles the GetTrafficIncidentDetail error response.
func (client *TrafficClient) getTrafficIncidentDetailHandleError(resp *azcore.Response) error {
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

// GetTrafficIncidentTile - Traffic Incident Tile
// Applies to: S0 and S1 pricing tiers.
// This service serves 256 x 256 pixel tiles showing traffic incidents. All tiles use the same grid system. Because the traffic tiles use transparent images,
// they can be layered on top of map tiles to
// create a compound display. Traffic tiles render graphics to indicate traffic on the roads in the specified area.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TrafficClient) GetTrafficIncidentTile(ctx context.Context, formatParam TileFormat, style TrafficIncidentTileStyle, zoom int32, xTileIndex int32, yTileIndex int32, options *TrafficGetTrafficIncidentTileOptions) (TrafficGetTrafficIncidentTileResponse, error) {
	req, err := client.getTrafficIncidentTileCreateRequest(ctx, formatParam, style, zoom, xTileIndex, yTileIndex, options)
	if err != nil {
		return TrafficGetTrafficIncidentTileResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TrafficGetTrafficIncidentTileResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TrafficGetTrafficIncidentTileResponse{}, client.getTrafficIncidentTileHandleError(resp)
	}
	return client.getTrafficIncidentTileHandleResponse(resp)
}

// getTrafficIncidentTileCreateRequest creates the GetTrafficIncidentTile request.
func (client *TrafficClient) getTrafficIncidentTileCreateRequest(ctx context.Context, formatParam TileFormat, style TrafficIncidentTileStyle, zoom int32, xTileIndex int32, yTileIndex int32, options *TrafficGetTrafficIncidentTileOptions) (*azcore.Request, error) {
	urlPath := "/traffic/incident/tile/{format}"
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
	reqQP.Set("style", string(style))
	reqQP.Set("zoom", strconv.FormatInt(int64(zoom), 10))
	reqQP.Set("x", strconv.FormatInt(int64(xTileIndex), 10))
	reqQP.Set("y", strconv.FormatInt(int64(yTileIndex), 10))
	if options != nil && options.TrafficState != nil {
		reqQP.Set("t", *options.TrafficState)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.SkipBodyDownload()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json, image/jpeg, image/png, image/pbf, application/vnd.mapbox-vector-tile")
	return req, nil
}

// getTrafficIncidentTileHandleResponse handles the GetTrafficIncidentTile response.
func (client *TrafficClient) getTrafficIncidentTileHandleResponse(resp *azcore.Response) (TrafficGetTrafficIncidentTileResponse, error) {
	result := TrafficGetTrafficIncidentTileResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	return result, nil
}

// getTrafficIncidentTileHandleError handles the GetTrafficIncidentTile error response.
func (client *TrafficClient) getTrafficIncidentTileHandleError(resp *azcore.Response) error {
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

// GetTrafficIncidentViewport - Traffic Incident Viewport
// Applies to: S0 and S1 pricing tiers.
// This API returns legal and technical information for the viewport described in the request. It should be called by client applications whenever the viewport
// changes (for instance, through zooming,
// panning, going to a location, or displaying a route). The request should contain the bounding box and zoom level of the viewport whose information is
// needed. The return will contain map version
// information, as well as the current Traffic Model ID and copyright IDs. The Traffic Model ID returned by the Viewport Description is used by other APIs
// to retrieve last traffic information for further
// processing.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TrafficClient) GetTrafficIncidentViewport(ctx context.Context, formatParam TextFormat, boundingbox string, boundingzoom int32, overviewbox string, overviewzoom int32, options *TrafficGetTrafficIncidentViewportOptions) (TrafficIncidentViewportResultResponse, error) {
	req, err := client.getTrafficIncidentViewportCreateRequest(ctx, formatParam, boundingbox, boundingzoom, overviewbox, overviewzoom, options)
	if err != nil {
		return TrafficIncidentViewportResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TrafficIncidentViewportResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TrafficIncidentViewportResultResponse{}, client.getTrafficIncidentViewportHandleError(resp)
	}
	return client.getTrafficIncidentViewportHandleResponse(resp)
}

// getTrafficIncidentViewportCreateRequest creates the GetTrafficIncidentViewport request.
func (client *TrafficClient) getTrafficIncidentViewportCreateRequest(ctx context.Context, formatParam TextFormat, boundingbox string, boundingzoom int32, overviewbox string, overviewzoom int32, options *TrafficGetTrafficIncidentViewportOptions) (*azcore.Request, error) {
	urlPath := "/traffic/incident/viewport/{format}"
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
	reqQP.Set("boundingbox", boundingbox)
	reqQP.Set("boundingzoom", strconv.FormatInt(int64(boundingzoom), 10))
	reqQP.Set("overviewbox", overviewbox)
	reqQP.Set("overviewzoom", strconv.FormatInt(int64(overviewzoom), 10))
	if options != nil && options.Copyright != nil {
		reqQP.Set("copyright", strconv.FormatBool(*options.Copyright))
	}
	req.URL.RawQuery = reqQP.Encode()
	if client.xmsClientID != nil {
		req.Header.Set("x-ms-client-id", *client.xmsClientID)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTrafficIncidentViewportHandleResponse handles the GetTrafficIncidentViewport response.
func (client *TrafficClient) getTrafficIncidentViewportHandleResponse(resp *azcore.Response) (TrafficIncidentViewportResultResponse, error) {
	var val *TrafficIncidentViewportResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TrafficIncidentViewportResultResponse{}, err
	}
return TrafficIncidentViewportResultResponse{RawResponse: resp.Response, TrafficIncidentViewportResult: val}, nil
}

// getTrafficIncidentViewportHandleError handles the GetTrafficIncidentViewport error response.
func (client *TrafficClient) getTrafficIncidentViewportHandleError(resp *azcore.Response) error {
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

