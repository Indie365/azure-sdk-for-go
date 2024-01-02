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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
	"net/http"
	"net/url"
	"regexp"
)

// EndpointsServer is a fake server for instances of the armtrafficmanager.EndpointsClient type.
type EndpointsServer struct {
	// CreateOrUpdate is the fake for method EndpointsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, profileName string, endpointType armtrafficmanager.EndpointType, endpointName string, parameters armtrafficmanager.Endpoint, options *armtrafficmanager.EndpointsClientCreateOrUpdateOptions) (resp azfake.Responder[armtrafficmanager.EndpointsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method EndpointsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, profileName string, endpointType armtrafficmanager.EndpointType, endpointName string, options *armtrafficmanager.EndpointsClientDeleteOptions) (resp azfake.Responder[armtrafficmanager.EndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, endpointType armtrafficmanager.EndpointType, endpointName string, options *armtrafficmanager.EndpointsClientGetOptions) (resp azfake.Responder[armtrafficmanager.EndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method EndpointsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, profileName string, endpointType armtrafficmanager.EndpointType, endpointName string, parameters armtrafficmanager.Endpoint, options *armtrafficmanager.EndpointsClientUpdateOptions) (resp azfake.Responder[armtrafficmanager.EndpointsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewEndpointsServerTransport creates a new instance of EndpointsServerTransport with the provided implementation.
// The returned EndpointsServerTransport instance is connected to an instance of armtrafficmanager.EndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEndpointsServerTransport(srv *EndpointsServer) *EndpointsServerTransport {
	return &EndpointsServerTransport{srv: srv}
}

// EndpointsServerTransport connects instances of armtrafficmanager.EndpointsClient to instances of EndpointsServer.
// Don't use this type directly, use NewEndpointsServerTransport instead.
type EndpointsServerTransport struct {
	srv *EndpointsServer
}

// Do implements the policy.Transporter interface for EndpointsServerTransport.
func (e *EndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EndpointsClient.CreateOrUpdate":
		resp, err = e.dispatchCreateOrUpdate(req)
	case "EndpointsClient.Delete":
		resp, err = e.dispatchDelete(req)
	case "EndpointsClient.Get":
		resp, err = e.dispatchGet(req)
	case "EndpointsClient.Update":
		resp, err = e.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/trafficmanagerprofiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armtrafficmanager.Endpoint](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointTypeParam, err := parseWithCast(matches[regex.SubexpIndex("endpointType")], func(v string) (armtrafficmanager.EndpointType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armtrafficmanager.EndpointType(p), nil
	})
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, profileNameParam, endpointTypeParam, endpointNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Endpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/trafficmanagerprofiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointTypeParam, err := parseWithCast(matches[regex.SubexpIndex("endpointType")], func(v string) (armtrafficmanager.EndpointType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armtrafficmanager.EndpointType(p), nil
	})
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), resourceGroupNameParam, profileNameParam, endpointTypeParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeleteOperationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/trafficmanagerprofiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointTypeParam, err := parseWithCast(matches[regex.SubexpIndex("endpointType")], func(v string) (armtrafficmanager.EndpointType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armtrafficmanager.EndpointType(p), nil
	})
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, endpointTypeParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Endpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/trafficmanagerprofiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armtrafficmanager.Endpoint](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointTypeParam, err := parseWithCast(matches[regex.SubexpIndex("endpointType")], func(v string) (armtrafficmanager.EndpointType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armtrafficmanager.EndpointType(p), nil
	})
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Update(req.Context(), resourceGroupNameParam, profileNameParam, endpointTypeParam, endpointNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Endpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
