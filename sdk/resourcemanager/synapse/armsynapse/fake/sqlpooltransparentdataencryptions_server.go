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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// SQLPoolTransparentDataEncryptionsServer is a fake server for instances of the armsynapse.SQLPoolTransparentDataEncryptionsClient type.
type SQLPoolTransparentDataEncryptionsServer struct {
	// CreateOrUpdate is the fake for method SQLPoolTransparentDataEncryptionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, transparentDataEncryptionName armsynapse.TransparentDataEncryptionName, parameters armsynapse.TransparentDataEncryption, options *armsynapse.SQLPoolTransparentDataEncryptionsClientCreateOrUpdateOptions) (resp azfake.Responder[armsynapse.SQLPoolTransparentDataEncryptionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SQLPoolTransparentDataEncryptionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, transparentDataEncryptionName armsynapse.TransparentDataEncryptionName, options *armsynapse.SQLPoolTransparentDataEncryptionsClientGetOptions) (resp azfake.Responder[armsynapse.SQLPoolTransparentDataEncryptionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SQLPoolTransparentDataEncryptionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, sqlPoolName string, options *armsynapse.SQLPoolTransparentDataEncryptionsClientListOptions) (resp azfake.PagerResponder[armsynapse.SQLPoolTransparentDataEncryptionsClientListResponse])
}

// NewSQLPoolTransparentDataEncryptionsServerTransport creates a new instance of SQLPoolTransparentDataEncryptionsServerTransport with the provided implementation.
// The returned SQLPoolTransparentDataEncryptionsServerTransport instance is connected to an instance of armsynapse.SQLPoolTransparentDataEncryptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLPoolTransparentDataEncryptionsServerTransport(srv *SQLPoolTransparentDataEncryptionsServer) *SQLPoolTransparentDataEncryptionsServerTransport {
	return &SQLPoolTransparentDataEncryptionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsynapse.SQLPoolTransparentDataEncryptionsClientListResponse]](),
	}
}

// SQLPoolTransparentDataEncryptionsServerTransport connects instances of armsynapse.SQLPoolTransparentDataEncryptionsClient to instances of SQLPoolTransparentDataEncryptionsServer.
// Don't use this type directly, use NewSQLPoolTransparentDataEncryptionsServerTransport instead.
type SQLPoolTransparentDataEncryptionsServerTransport struct {
	srv          *SQLPoolTransparentDataEncryptionsServer
	newListPager *tracker[azfake.PagerResponder[armsynapse.SQLPoolTransparentDataEncryptionsClientListResponse]]
}

// Do implements the policy.Transporter interface for SQLPoolTransparentDataEncryptionsServerTransport.
func (s *SQLPoolTransparentDataEncryptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLPoolTransparentDataEncryptionsClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SQLPoolTransparentDataEncryptionsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SQLPoolTransparentDataEncryptionsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLPoolTransparentDataEncryptionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transparentDataEncryption/(?P<transparentDataEncryptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsynapse.TransparentDataEncryption](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
	if err != nil {
		return nil, err
	}
	transparentDataEncryptionNameParam, err := parseWithCast(matches[regex.SubexpIndex("transparentDataEncryptionName")], func(v string) (armsynapse.TransparentDataEncryptionName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsynapse.TransparentDataEncryptionName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, transparentDataEncryptionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TransparentDataEncryption, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLPoolTransparentDataEncryptionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transparentDataEncryption/(?P<transparentDataEncryptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
	if err != nil {
		return nil, err
	}
	transparentDataEncryptionNameParam, err := parseWithCast(matches[regex.SubexpIndex("transparentDataEncryptionName")], func(v string) (armsynapse.TransparentDataEncryptionName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsynapse.TransparentDataEncryptionName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, transparentDataEncryptionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TransparentDataEncryption, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLPoolTransparentDataEncryptionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transparentDataEncryption`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsynapse.SQLPoolTransparentDataEncryptionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}
