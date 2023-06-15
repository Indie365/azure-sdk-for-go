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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ExpressRouteCrossConnectionsServer is a fake server for instances of the armnetwork.ExpressRouteCrossConnectionsClient type.
type ExpressRouteCrossConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRouteCrossConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters armnetwork.ExpressRouteCrossConnection, options *armnetwork.ExpressRouteCrossConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRouteCrossConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, crossConnectionName string, options *armnetwork.ExpressRouteCrossConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteCrossConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRouteCrossConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.ExpressRouteCrossConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteCrossConnectionsClientListResponse])

	// BeginListArpTable is the fake for method ExpressRouteCrossConnectionsClient.BeginListArpTable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListArpTable func(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *armnetwork.ExpressRouteCrossConnectionsClientBeginListArpTableOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientListArpTableResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ExpressRouteCrossConnectionsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.ExpressRouteCrossConnectionsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteCrossConnectionsClientListByResourceGroupResponse])

	// BeginListRoutesTable is the fake for method ExpressRouteCrossConnectionsClient.BeginListRoutesTable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListRoutesTable func(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *armnetwork.ExpressRouteCrossConnectionsClientBeginListRoutesTableOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientListRoutesTableResponse], errResp azfake.ErrorResponder)

	// BeginListRoutesTableSummary is the fake for method ExpressRouteCrossConnectionsClient.BeginListRoutesTableSummary
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListRoutesTableSummary func(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *armnetwork.ExpressRouteCrossConnectionsClientBeginListRoutesTableSummaryOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method ExpressRouteCrossConnectionsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters armnetwork.TagsObject, options *armnetwork.ExpressRouteCrossConnectionsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ExpressRouteCrossConnectionsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewExpressRouteCrossConnectionsServerTransport creates a new instance of ExpressRouteCrossConnectionsServerTransport with the provided implementation.
// The returned ExpressRouteCrossConnectionsServerTransport instance is connected to an instance of armnetwork.ExpressRouteCrossConnectionsClient by way of the
// undefined.Transporter field.
func NewExpressRouteCrossConnectionsServerTransport(srv *ExpressRouteCrossConnectionsServer) *ExpressRouteCrossConnectionsServerTransport {
	return &ExpressRouteCrossConnectionsServerTransport{srv: srv}
}

// ExpressRouteCrossConnectionsServerTransport connects instances of armnetwork.ExpressRouteCrossConnectionsClient to instances of ExpressRouteCrossConnectionsServer.
// Don't use this type directly, use NewExpressRouteCrossConnectionsServerTransport instead.
type ExpressRouteCrossConnectionsServerTransport struct {
	srv                         *ExpressRouteCrossConnectionsServer
	beginCreateOrUpdate         *azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientCreateOrUpdateResponse]
	newListPager                *azfake.PagerResponder[armnetwork.ExpressRouteCrossConnectionsClientListResponse]
	beginListArpTable           *azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientListArpTableResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armnetwork.ExpressRouteCrossConnectionsClientListByResourceGroupResponse]
	beginListRoutesTable        *azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientListRoutesTableResponse]
	beginListRoutesTableSummary *azfake.PollerResponder[armnetwork.ExpressRouteCrossConnectionsClientListRoutesTableSummaryResponse]
}

// Do implements the policy.Transporter interface for ExpressRouteCrossConnectionsServerTransport.
func (e *ExpressRouteCrossConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteCrossConnectionsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRouteCrossConnectionsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRouteCrossConnectionsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	case "ExpressRouteCrossConnectionsClient.BeginListArpTable":
		resp, err = e.dispatchBeginListArpTable(req)
	case "ExpressRouteCrossConnectionsClient.NewListByResourceGroupPager":
		resp, err = e.dispatchNewListByResourceGroupPager(req)
	case "ExpressRouteCrossConnectionsClient.BeginListRoutesTable":
		resp, err = e.dispatchBeginListRoutesTable(req)
	case "ExpressRouteCrossConnectionsClient.BeginListRoutesTableSummary":
		resp, err = e.dispatchBeginListRoutesTableSummary(req)
	case "ExpressRouteCrossConnectionsClient.UpdateTags":
		resp, err = e.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if e.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections/(?P<crossConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRouteCrossConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		crossConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("crossConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, crossConnectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(e.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginCreateOrUpdate) {
		e.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections/(?P<crossConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	crossConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("crossConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, crossConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCrossConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if e.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListPager(nil)
		e.newListPager = &resp
		server.PagerResponderInjectNextLinks(e.newListPager, req, func(page *armnetwork.ExpressRouteCrossConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(e.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(e.newListPager) {
		e.newListPager = nil
	}
	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchBeginListArpTable(req *http.Request) (*http.Response, error) {
	if e.srv.BeginListArpTable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListArpTable not implemented")}
	}
	if e.beginListArpTable == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections/(?P<crossConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arpTables/(?P<devicePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		crossConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("crossConnectionName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		devicePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("devicePath")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginListArpTable(req.Context(), resourceGroupNameUnescaped, crossConnectionNameUnescaped, peeringNameUnescaped, devicePathUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginListArpTable = &respr
	}

	resp, err := server.PollerResponderNext(e.beginListArpTable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginListArpTable) {
		e.beginListArpTable = nil
	}

	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if e.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		e.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(e.newListByResourceGroupPager, req, func(page *armnetwork.ExpressRouteCrossConnectionsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(e.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(e.newListByResourceGroupPager) {
		e.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchBeginListRoutesTable(req *http.Request) (*http.Response, error) {
	if e.srv.BeginListRoutesTable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListRoutesTable not implemented")}
	}
	if e.beginListRoutesTable == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections/(?P<crossConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeTables/(?P<devicePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		crossConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("crossConnectionName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		devicePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("devicePath")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginListRoutesTable(req.Context(), resourceGroupNameUnescaped, crossConnectionNameUnescaped, peeringNameUnescaped, devicePathUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginListRoutesTable = &respr
	}

	resp, err := server.PollerResponderNext(e.beginListRoutesTable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginListRoutesTable) {
		e.beginListRoutesTable = nil
	}

	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchBeginListRoutesTableSummary(req *http.Request) (*http.Response, error) {
	if e.srv.BeginListRoutesTableSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListRoutesTableSummary not implemented")}
	}
	if e.beginListRoutesTableSummary == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections/(?P<crossConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeTablesSummary/(?P<devicePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		crossConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("crossConnectionName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		devicePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("devicePath")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginListRoutesTableSummary(req.Context(), resourceGroupNameUnescaped, crossConnectionNameUnescaped, peeringNameUnescaped, devicePathUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginListRoutesTableSummary = &respr
	}

	resp, err := server.PollerResponderNext(e.beginListRoutesTableSummary, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginListRoutesTableSummary) {
		e.beginListRoutesTableSummary = nil
	}

	return resp, nil
}

func (e *ExpressRouteCrossConnectionsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if e.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCrossConnections/(?P<crossConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	crossConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("crossConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, crossConnectionNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCrossConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
