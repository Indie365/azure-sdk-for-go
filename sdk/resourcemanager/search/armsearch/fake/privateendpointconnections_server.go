//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionsServer is a fake server for instances of the armsearch.PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsServer struct {
	// Delete is the fake for method PrivateEndpointConnectionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	Delete func(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.PrivateEndpointConnectionsClientDeleteOptions) (resp azfake.Responder[armsearch.PrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.PrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armsearch.PrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method PrivateEndpointConnectionsClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.PrivateEndpointConnectionsClientListByServiceOptions) (resp azfake.PagerResponder[armsearch.PrivateEndpointConnectionsClientListByServiceResponse])

	// Update is the fake for method PrivateEndpointConnectionsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, searchServiceName string, privateEndpointConnectionName string, privateEndpointConnection armsearch.PrivateEndpointConnection, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.PrivateEndpointConnectionsClientUpdateOptions) (resp azfake.Responder[armsearch.PrivateEndpointConnectionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPrivateEndpointConnectionsServerTransport creates a new instance of PrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsServerTransport instance is connected to an instance of armsearch.PrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsServerTransport(srv *PrivateEndpointConnectionsServer) *PrivateEndpointConnectionsServerTransport {
	return &PrivateEndpointConnectionsServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armsearch.PrivateEndpointConnectionsClientListByServiceResponse]](),
	}
}

// PrivateEndpointConnectionsServerTransport connects instances of armsearch.PrivateEndpointConnectionsClient to instances of PrivateEndpointConnectionsServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsServerTransport instead.
type PrivateEndpointConnectionsServerTransport struct {
	srv                   *PrivateEndpointConnectionsServer
	newListByServicePager *tracker[azfake.PagerResponder[armsearch.PrivateEndpointConnectionsClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsServerTransport.
func (p *PrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PrivateEndpointConnectionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionsClient.NewListByServicePager":
		resp, err = p.dispatchNewListByServicePager(req)
	case "PrivateEndpointConnectionsClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
	if clientRequestIDParam != nil {
		searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, searchServiceNameParam, privateEndpointConnectionNameParam, searchManagementRequestOptions, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
	if clientRequestIDParam != nil {
		searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, searchServiceNameParam, privateEndpointConnectionNameParam, searchManagementRequestOptions, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := p.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
		if clientRequestIDParam != nil {
			searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
				ClientRequestID: clientRequestIDParam,
			}
		}
		resp := p.srv.NewListByServicePager(resourceGroupNameParam, searchServiceNameParam, searchManagementRequestOptions, nil)
		newListByServicePager = &resp
		p.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armsearch.PrivateEndpointConnectionsClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		p.newListByServicePager.remove(req)
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsearch.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
	if clientRequestIDParam != nil {
		searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameParam, searchServiceNameParam, privateEndpointConnectionNameParam, body, searchManagementRequestOptions, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
