//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
)

// TenantActivityLogsServer is a fake server for instances of the armmonitor.TenantActivityLogsClient type.
type TenantActivityLogsServer struct {
	// NewListPager is the fake for method TenantActivityLogsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armmonitor.TenantActivityLogsClientListOptions) (resp azfake.PagerResponder[armmonitor.TenantActivityLogsClientListResponse])
}

// NewTenantActivityLogsServerTransport creates a new instance of TenantActivityLogsServerTransport with the provided implementation.
// The returned TenantActivityLogsServerTransport instance is connected to an instance of armmonitor.TenantActivityLogsClient by way of the
// undefined.Transporter field.
func NewTenantActivityLogsServerTransport(srv *TenantActivityLogsServer) *TenantActivityLogsServerTransport {
	return &TenantActivityLogsServerTransport{srv: srv}
}

// TenantActivityLogsServerTransport connects instances of armmonitor.TenantActivityLogsClient to instances of TenantActivityLogsServer.
// Don't use this type directly, use NewTenantActivityLogsServerTransport instead.
type TenantActivityLogsServerTransport struct {
	srv          *TenantActivityLogsServer
	newListPager *azfake.PagerResponder[armmonitor.TenantActivityLogsClientListResponse]
}

// Do implements the policy.Transporter interface for TenantActivityLogsServerTransport.
func (t *TenantActivityLogsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TenantActivityLogsClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TenantActivityLogsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if t.newListPager == nil {
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
		if err != nil {
			return nil, err
		}
		selectParam := getOptional(selectUnescaped)
		var options *armmonitor.TenantActivityLogsClientListOptions
		if filterParam != nil || selectParam != nil {
			options = &armmonitor.TenantActivityLogsClientListOptions{
				Filter: filterParam,
				Select: selectParam,
			}
		}
		resp := t.srv.NewListPager(options)
		t.newListPager = &resp
		server.PagerResponderInjectNextLinks(t.newListPager, req, func(page *armmonitor.TenantActivityLogsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(t.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(t.newListPager) {
		t.newListPager = nil
	}
	return resp, nil
}
