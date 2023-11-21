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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// GroupUserServer is a fake server for instances of the armapimanagement.GroupUserClient type.
type GroupUserServer struct {
	// CheckEntityExists is the fake for method GroupUserClient.CheckEntityExists
	// HTTP status codes to indicate success: http.StatusNoContent, http.StatusNotFound
	CheckEntityExists func(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string, options *armapimanagement.GroupUserClientCheckEntityExistsOptions) (resp azfake.Responder[armapimanagement.GroupUserClientCheckEntityExistsResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method GroupUserClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string, options *armapimanagement.GroupUserClientCreateOptions) (resp azfake.Responder[armapimanagement.GroupUserClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method GroupUserClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string, options *armapimanagement.GroupUserClientDeleteOptions) (resp azfake.Responder[armapimanagement.GroupUserClientDeleteResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GroupUserClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, groupID string, options *armapimanagement.GroupUserClientListOptions) (resp azfake.PagerResponder[armapimanagement.GroupUserClientListResponse])
}

// NewGroupUserServerTransport creates a new instance of GroupUserServerTransport with the provided implementation.
// The returned GroupUserServerTransport instance is connected to an instance of armapimanagement.GroupUserClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGroupUserServerTransport(srv *GroupUserServer) *GroupUserServerTransport {
	return &GroupUserServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armapimanagement.GroupUserClientListResponse]](),
	}
}

// GroupUserServerTransport connects instances of armapimanagement.GroupUserClient to instances of GroupUserServer.
// Don't use this type directly, use NewGroupUserServerTransport instead.
type GroupUserServerTransport struct {
	srv          *GroupUserServer
	newListPager *tracker[azfake.PagerResponder[armapimanagement.GroupUserClientListResponse]]
}

// Do implements the policy.Transporter interface for GroupUserServerTransport.
func (g *GroupUserServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GroupUserClient.CheckEntityExists":
		resp, err = g.dispatchCheckEntityExists(req)
	case "GroupUserClient.Create":
		resp, err = g.dispatchCreate(req)
	case "GroupUserClient.Delete":
		resp, err = g.dispatchDelete(req)
	case "GroupUserClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GroupUserServerTransport) dispatchCheckEntityExists(req *http.Request) (*http.Response, error) {
	if g.srv.CheckEntityExists == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckEntityExists not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	userIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.CheckEntityExists(req.Context(), resourceGroupNameParam, serviceNameParam, groupIDParam, userIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GroupUserServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if g.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	userIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Create(req.Context(), resourceGroupNameParam, serviceNameParam, groupIDParam, userIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UserContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GroupUserServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if g.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	userIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("userId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, groupIDParam, userIDParam, nil)
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

func (g *GroupUserServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.GroupUserClientListOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.GroupUserClientListOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := g.srv.NewListPager(resourceGroupNameParam, serviceNameParam, groupIDParam, options)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapimanagement.GroupUserClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		g.newListPager.remove(req)
	}
	return resp, nil
}
