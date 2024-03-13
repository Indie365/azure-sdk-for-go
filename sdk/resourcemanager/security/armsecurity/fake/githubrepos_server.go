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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// GitHubReposServer is a fake server for instances of the armsecurity.GitHubReposClient type.
type GitHubReposServer struct {
	// Get is the fake for method GitHubReposClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, securityConnectorName string, ownerName string, repoName string, options *armsecurity.GitHubReposClientGetOptions) (resp azfake.Responder[armsecurity.GitHubReposClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GitHubReposClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, securityConnectorName string, ownerName string, options *armsecurity.GitHubReposClientListOptions) (resp azfake.PagerResponder[armsecurity.GitHubReposClientListResponse])
}

// NewGitHubReposServerTransport creates a new instance of GitHubReposServerTransport with the provided implementation.
// The returned GitHubReposServerTransport instance is connected to an instance of armsecurity.GitHubReposClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGitHubReposServerTransport(srv *GitHubReposServer) *GitHubReposServerTransport {
	return &GitHubReposServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurity.GitHubReposClientListResponse]](),
	}
}

// GitHubReposServerTransport connects instances of armsecurity.GitHubReposClient to instances of GitHubReposServer.
// Don't use this type directly, use NewGitHubReposServerTransport instead.
type GitHubReposServerTransport struct {
	srv          *GitHubReposServer
	newListPager *tracker[azfake.PagerResponder[armsecurity.GitHubReposClientListResponse]]
}

// Do implements the policy.Transporter interface for GitHubReposServerTransport.
func (g *GitHubReposServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GitHubReposClient.Get":
		resp, err = g.dispatchGet(req)
	case "GitHubReposClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GitHubReposServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/securityConnectors/(?P<securityConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devops/default/gitHubOwners/(?P<ownerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repos/(?P<repoName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	securityConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("securityConnectorName")])
	if err != nil {
		return nil, err
	}
	ownerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ownerName")])
	if err != nil {
		return nil, err
	}
	repoNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("repoName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, securityConnectorNameParam, ownerNameParam, repoNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GitHubRepository, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GitHubReposServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/securityConnectors/(?P<securityConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devops/default/gitHubOwners/(?P<ownerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repos`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		securityConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("securityConnectorName")])
		if err != nil {
			return nil, err
		}
		ownerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ownerName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListPager(resourceGroupNameParam, securityConnectorNameParam, ownerNameParam, nil)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.GitHubReposClientListResponse, createLink func() string) {
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
