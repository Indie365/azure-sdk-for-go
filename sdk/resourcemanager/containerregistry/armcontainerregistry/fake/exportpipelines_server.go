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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"net/http"
	"net/url"
	"regexp"
)

// ExportPipelinesServer is a fake server for instances of the armcontainerregistry.ExportPipelinesClient type.
type ExportPipelinesServer struct {
	// BeginCreate is the fake for method ExportPipelinesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters armcontainerregistry.ExportPipeline, options *armcontainerregistry.ExportPipelinesClientBeginCreateOptions) (resp azfake.PollerResponder[armcontainerregistry.ExportPipelinesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExportPipelinesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *armcontainerregistry.ExportPipelinesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerregistry.ExportPipelinesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExportPipelinesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *armcontainerregistry.ExportPipelinesClientGetOptions) (resp azfake.Responder[armcontainerregistry.ExportPipelinesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExportPipelinesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, options *armcontainerregistry.ExportPipelinesClientListOptions) (resp azfake.PagerResponder[armcontainerregistry.ExportPipelinesClientListResponse])
}

// NewExportPipelinesServerTransport creates a new instance of ExportPipelinesServerTransport with the provided implementation.
// The returned ExportPipelinesServerTransport instance is connected to an instance of armcontainerregistry.ExportPipelinesClient by way of the
// undefined.Transporter field.
func NewExportPipelinesServerTransport(srv *ExportPipelinesServer) *ExportPipelinesServerTransport {
	return &ExportPipelinesServerTransport{srv: srv}
}

// ExportPipelinesServerTransport connects instances of armcontainerregistry.ExportPipelinesClient to instances of ExportPipelinesServer.
// Don't use this type directly, use NewExportPipelinesServerTransport instead.
type ExportPipelinesServerTransport struct {
	srv          *ExportPipelinesServer
	beginCreate  *azfake.PollerResponder[armcontainerregistry.ExportPipelinesClientCreateResponse]
	beginDelete  *azfake.PollerResponder[armcontainerregistry.ExportPipelinesClientDeleteResponse]
	newListPager *azfake.PagerResponder[armcontainerregistry.ExportPipelinesClientListResponse]
}

// Do implements the policy.Transporter interface for ExportPipelinesServerTransport.
func (e *ExportPipelinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExportPipelinesClient.BeginCreate":
		resp, err = e.dispatchBeginCreate(req)
	case "ExportPipelinesClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExportPipelinesClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExportPipelinesClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExportPipelinesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	if e.beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportPipelines/(?P<exportPipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.ExportPipeline](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		exportPipelineNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("exportPipelineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreate(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, exportPipelineNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginCreate = &respr
	}

	resp, err := server.PollerResponderNext(e.beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginCreate) {
		e.beginCreate = nil
	}

	return resp, nil
}

func (e *ExportPipelinesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if e.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportPipelines/(?P<exportPipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		exportPipelineNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("exportPipelineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, exportPipelineNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(e.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginDelete) {
		e.beginDelete = nil
	}

	return resp, nil
}

func (e *ExportPipelinesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportPipelines/(?P<exportPipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	exportPipelineNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("exportPipelineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, exportPipelineNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExportPipeline, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExportPipelinesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if e.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportPipelines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameUnescaped, registryNameUnescaped, nil)
		e.newListPager = &resp
		server.PagerResponderInjectNextLinks(e.newListPager, req, func(page *armcontainerregistry.ExportPipelinesClientListResponse, createLink func() string) {
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
