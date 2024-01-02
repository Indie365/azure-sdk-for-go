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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
	"net/http"
	"net/url"
	"regexp"
)

// GitHubConnectorServer is a fake server for instances of the armsecuritydevops.GitHubConnectorClient type.
type GitHubConnectorServer struct {
	// BeginCreateOrUpdate is the fake for method GitHubConnectorClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector armsecuritydevops.GitHubConnector, options *armsecuritydevops.GitHubConnectorClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GitHubConnectorClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *armsecuritydevops.GitHubConnectorClientBeginDeleteOptions) (resp azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GitHubConnectorClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *armsecuritydevops.GitHubConnectorClientGetOptions) (resp azfake.Responder[armsecuritydevops.GitHubConnectorClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method GitHubConnectorClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsecuritydevops.GitHubConnectorClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsecuritydevops.GitHubConnectorClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method GitHubConnectorClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armsecuritydevops.GitHubConnectorClientListBySubscriptionOptions) (resp azfake.PagerResponder[armsecuritydevops.GitHubConnectorClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method GitHubConnectorClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector armsecuritydevops.GitHubConnector, options *armsecuritydevops.GitHubConnectorClientBeginUpdateOptions) (resp azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGitHubConnectorServerTransport creates a new instance of GitHubConnectorServerTransport with the provided implementation.
// The returned GitHubConnectorServerTransport instance is connected to an instance of armsecuritydevops.GitHubConnectorClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGitHubConnectorServerTransport(srv *GitHubConnectorServer) *GitHubConnectorServerTransport {
	return &GitHubConnectorServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armsecuritydevops.GitHubConnectorClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armsecuritydevops.GitHubConnectorClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientUpdateResponse]](),
	}
}

// GitHubConnectorServerTransport connects instances of armsecuritydevops.GitHubConnectorClient to instances of GitHubConnectorServer.
// Don't use this type directly, use NewGitHubConnectorServerTransport instead.
type GitHubConnectorServerTransport struct {
	srv                         *GitHubConnectorServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armsecuritydevops.GitHubConnectorClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armsecuritydevops.GitHubConnectorClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armsecuritydevops.GitHubConnectorClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for GitHubConnectorServerTransport.
func (g *GitHubConnectorServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GitHubConnectorClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GitHubConnectorClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GitHubConnectorClient.Get":
		resp, err = g.dispatchGet(req)
	case "GitHubConnectorClient.NewListByResourceGroupPager":
		resp, err = g.dispatchNewListByResourceGroupPager(req)
	case "GitHubConnectorClient.NewListBySubscriptionPager":
		resp, err = g.dispatchNewListBySubscriptionPager(req)
	case "GitHubConnectorClient.BeginUpdate":
		resp, err = g.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GitHubConnectorServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityDevOps/gitHubConnectors/(?P<gitHubConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsecuritydevops.GitHubConnector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gitHubConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gitHubConnectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, gitHubConnectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GitHubConnectorServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := g.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityDevOps/gitHubConnectors/(?P<gitHubConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gitHubConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gitHubConnectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), resourceGroupNameParam, gitHubConnectorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		g.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		g.beginDelete.remove(req)
	}

	return resp, nil
}

func (g *GitHubConnectorServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityDevOps/gitHubConnectors/(?P<gitHubConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gitHubConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gitHubConnectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, gitHubConnectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GitHubConnector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GitHubConnectorServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := g.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityDevOps/gitHubConnectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		g.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsecuritydevops.GitHubConnectorClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		g.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (g *GitHubConnectorServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := g.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityDevOps/gitHubConnectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		g.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armsecuritydevops.GitHubConnectorClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		g.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (g *GitHubConnectorServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := g.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityDevOps/gitHubConnectors/(?P<gitHubConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsecuritydevops.GitHubConnector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gitHubConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gitHubConnectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), resourceGroupNameParam, gitHubConnectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		g.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		g.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		g.beginUpdate.remove(req)
	}

	return resp, nil
}
