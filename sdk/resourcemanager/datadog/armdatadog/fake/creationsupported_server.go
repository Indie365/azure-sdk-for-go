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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
	"net/http"
	"net/url"
	"regexp"
)

// CreationSupportedServer is a fake server for instances of the armdatadog.CreationSupportedClient type.
type CreationSupportedServer struct {
	// Get is the fake for method CreationSupportedClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, datadogOrganizationID string, options *armdatadog.CreationSupportedClientGetOptions) (resp azfake.Responder[armdatadog.CreationSupportedClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CreationSupportedClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(datadogOrganizationID string, options *armdatadog.CreationSupportedClientListOptions) (resp azfake.PagerResponder[armdatadog.CreationSupportedClientListResponse])
}

// NewCreationSupportedServerTransport creates a new instance of CreationSupportedServerTransport with the provided implementation.
// The returned CreationSupportedServerTransport instance is connected to an instance of armdatadog.CreationSupportedClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCreationSupportedServerTransport(srv *CreationSupportedServer) *CreationSupportedServerTransport {
	return &CreationSupportedServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdatadog.CreationSupportedClientListResponse]](),
	}
}

// CreationSupportedServerTransport connects instances of armdatadog.CreationSupportedClient to instances of CreationSupportedServer.
// Don't use this type directly, use NewCreationSupportedServerTransport instead.
type CreationSupportedServerTransport struct {
	srv          *CreationSupportedServer
	newListPager *tracker[azfake.PagerResponder[armdatadog.CreationSupportedClientListResponse]]
}

// Do implements the policy.Transporter interface for CreationSupportedServerTransport.
func (c *CreationSupportedServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CreationSupportedClient.Get":
		resp, err = c.dispatchGet(req)
	case "CreationSupportedClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CreationSupportedServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Datadog/subscriptionStatuses/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	datadogOrganizationIDParam, err := url.QueryUnescape(qp.Get("datadogOrganizationId"))
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), datadogOrganizationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CreateResourceSupportedResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CreationSupportedServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Datadog/subscriptionStatuses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		datadogOrganizationIDParam, err := url.QueryUnescape(qp.Get("datadogOrganizationId"))
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(datadogOrganizationIDParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}
