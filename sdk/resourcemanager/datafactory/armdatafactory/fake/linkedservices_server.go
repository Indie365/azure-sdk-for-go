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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
	"net/http"
	"net/url"
	"regexp"
)

// LinkedServicesServer is a fake server for instances of the armdatafactory.LinkedServicesClient type.
type LinkedServicesServer struct {
	// CreateOrUpdate is the fake for method LinkedServicesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, linkedService armdatafactory.LinkedServiceResource, options *armdatafactory.LinkedServicesClientCreateOrUpdateOptions) (resp azfake.Responder[armdatafactory.LinkedServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method LinkedServicesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, options *armdatafactory.LinkedServicesClientDeleteOptions) (resp azfake.Responder[armdatafactory.LinkedServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LinkedServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	Get func(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, options *armdatafactory.LinkedServicesClientGetOptions) (resp azfake.Responder[armdatafactory.LinkedServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFactoryPager is the fake for method LinkedServicesClient.NewListByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFactoryPager func(resourceGroupName string, factoryName string, options *armdatafactory.LinkedServicesClientListByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.LinkedServicesClientListByFactoryResponse])
}

// NewLinkedServicesServerTransport creates a new instance of LinkedServicesServerTransport with the provided implementation.
// The returned LinkedServicesServerTransport instance is connected to an instance of armdatafactory.LinkedServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLinkedServicesServerTransport(srv *LinkedServicesServer) *LinkedServicesServerTransport {
	return &LinkedServicesServerTransport{
		srv:                   srv,
		newListByFactoryPager: newTracker[azfake.PagerResponder[armdatafactory.LinkedServicesClientListByFactoryResponse]](),
	}
}

// LinkedServicesServerTransport connects instances of armdatafactory.LinkedServicesClient to instances of LinkedServicesServer.
// Don't use this type directly, use NewLinkedServicesServerTransport instead.
type LinkedServicesServerTransport struct {
	srv                   *LinkedServicesServer
	newListByFactoryPager *tracker[azfake.PagerResponder[armdatafactory.LinkedServicesClientListByFactoryResponse]]
}

// Do implements the policy.Transporter interface for LinkedServicesServerTransport.
func (l *LinkedServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LinkedServicesClient.CreateOrUpdate":
		resp, err = l.dispatchCreateOrUpdate(req)
	case "LinkedServicesClient.Delete":
		resp, err = l.dispatchDelete(req)
	case "LinkedServicesClient.Get":
		resp, err = l.dispatchGet(req)
	case "LinkedServicesClient.NewListByFactoryPager":
		resp, err = l.dispatchNewListByFactoryPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LinkedServicesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkedservices/(?P<linkedServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.LinkedServiceResource](req)
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
	linkedServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkedServiceName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armdatafactory.LinkedServicesClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armdatafactory.LinkedServicesClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := l.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, factoryNameParam, linkedServiceNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkedServiceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkedServicesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if l.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkedservices/(?P<linkedServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	linkedServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkedServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, linkedServiceNameParam, nil)
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

func (l *LinkedServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkedservices/(?P<linkedServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	linkedServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkedServiceName")])
	if err != nil {
		return nil, err
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armdatafactory.LinkedServicesClientGetOptions
	if ifNoneMatchParam != nil {
		options = &armdatafactory.LinkedServicesClientGetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, linkedServiceNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkedServiceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinkedServicesServerTransport) dispatchNewListByFactoryPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFactoryPager not implemented")}
	}
	newListByFactoryPager := l.newListByFactoryPager.get(req)
	if newListByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkedservices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := l.srv.NewListByFactoryPager(resourceGroupNameParam, factoryNameParam, nil)
		newListByFactoryPager = &resp
		l.newListByFactoryPager.add(req, newListByFactoryPager)
		server.PagerResponderInjectNextLinks(newListByFactoryPager, req, func(page *armdatafactory.LinkedServicesClientListByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFactoryPager) {
		l.newListByFactoryPager.remove(req)
	}
	return resp, nil
}
