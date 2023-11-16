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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedPrivateEndpointsServer is a fake server for instances of the armdatafactory.ManagedPrivateEndpointsClient type.
type ManagedPrivateEndpointsServer struct {
	// CreateOrUpdate is the fake for method ManagedPrivateEndpointsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint armdatafactory.ManagedPrivateEndpointResource, options *armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateOptions) (resp azfake.Responder[armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagedPrivateEndpointsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *armdatafactory.ManagedPrivateEndpointsClientDeleteOptions) (resp azfake.Responder[armdatafactory.ManagedPrivateEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedPrivateEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *armdatafactory.ManagedPrivateEndpointsClientGetOptions) (resp azfake.Responder[armdatafactory.ManagedPrivateEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFactoryPager is the fake for method ManagedPrivateEndpointsClient.NewListByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFactoryPager func(resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *armdatafactory.ManagedPrivateEndpointsClientListByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.ManagedPrivateEndpointsClientListByFactoryResponse])
}

// NewManagedPrivateEndpointsServerTransport creates a new instance of ManagedPrivateEndpointsServerTransport with the provided implementation.
// The returned ManagedPrivateEndpointsServerTransport instance is connected to an instance of armdatafactory.ManagedPrivateEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedPrivateEndpointsServerTransport(srv *ManagedPrivateEndpointsServer) *ManagedPrivateEndpointsServerTransport {
	return &ManagedPrivateEndpointsServerTransport{
		srv:                   srv,
		newListByFactoryPager: newTracker[azfake.PagerResponder[armdatafactory.ManagedPrivateEndpointsClientListByFactoryResponse]](),
	}
}

// ManagedPrivateEndpointsServerTransport connects instances of armdatafactory.ManagedPrivateEndpointsClient to instances of ManagedPrivateEndpointsServer.
// Don't use this type directly, use NewManagedPrivateEndpointsServerTransport instead.
type ManagedPrivateEndpointsServerTransport struct {
	srv                   *ManagedPrivateEndpointsServer
	newListByFactoryPager *tracker[azfake.PagerResponder[armdatafactory.ManagedPrivateEndpointsClientListByFactoryResponse]]
}

// Do implements the policy.Transporter interface for ManagedPrivateEndpointsServerTransport.
func (m *ManagedPrivateEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedPrivateEndpointsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagedPrivateEndpointsClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ManagedPrivateEndpointsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedPrivateEndpointsClient.NewListByFactoryPager":
		resp, err = m.dispatchNewListByFactoryPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedVirtualNetworks/(?P<managedVirtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints/(?P<managedPrivateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.ManagedPrivateEndpointResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	managedVirtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedVirtualNetworkName")])
	if err != nil {
		return nil, err
	}
	managedPrivateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedPrivateEndpointName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, factoryNameParam, managedVirtualNetworkNameParam, managedPrivateEndpointNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedPrivateEndpointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedVirtualNetworks/(?P<managedVirtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints/(?P<managedPrivateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	managedVirtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedVirtualNetworkName")])
	if err != nil {
		return nil, err
	}
	managedPrivateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedPrivateEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, managedVirtualNetworkNameParam, managedPrivateEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedVirtualNetworks/(?P<managedVirtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints/(?P<managedPrivateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	managedVirtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedVirtualNetworkName")])
	if err != nil {
		return nil, err
	}
	managedPrivateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedPrivateEndpointName")])
	if err != nil {
		return nil, err
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armdatafactory.ManagedPrivateEndpointsClientGetOptions
	if ifNoneMatchParam != nil {
		options = &armdatafactory.ManagedPrivateEndpointsClientGetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, managedVirtualNetworkNameParam, managedPrivateEndpointNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedPrivateEndpointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedPrivateEndpointsServerTransport) dispatchNewListByFactoryPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFactoryPager not implemented")}
	}
	newListByFactoryPager := m.newListByFactoryPager.get(req)
	if newListByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedVirtualNetworks/(?P<managedVirtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedPrivateEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		managedVirtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedVirtualNetworkName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByFactoryPager(resourceGroupNameParam, factoryNameParam, managedVirtualNetworkNameParam, nil)
		newListByFactoryPager = &resp
		m.newListByFactoryPager.add(req, newListByFactoryPager)
		server.PagerResponderInjectNextLinks(newListByFactoryPager, req, func(page *armdatafactory.ManagedPrivateEndpointsClientListByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFactoryPager) {
		m.newListByFactoryPager.remove(req)
	}
	return resp, nil
}
