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

// TasksServer is a fake server for instances of the armcontainerregistry.TasksClient type.
type TasksServer struct {
	// BeginCreate is the fake for method TasksClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskCreateParameters armcontainerregistry.Task, options *armcontainerregistry.TasksClientBeginCreateOptions) (resp azfake.PollerResponder[armcontainerregistry.TasksClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method TasksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *armcontainerregistry.TasksClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerregistry.TasksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TasksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *armcontainerregistry.TasksClientGetOptions) (resp azfake.Responder[armcontainerregistry.TasksClientGetResponse], errResp azfake.ErrorResponder)

	// GetDetails is the fake for method TasksClient.GetDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetDetails func(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *armcontainerregistry.TasksClientGetDetailsOptions) (resp azfake.Responder[armcontainerregistry.TasksClientGetDetailsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TasksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, options *armcontainerregistry.TasksClientListOptions) (resp azfake.PagerResponder[armcontainerregistry.TasksClientListResponse])

	// BeginUpdate is the fake for method TasksClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskUpdateParameters armcontainerregistry.TaskUpdateParameters, options *armcontainerregistry.TasksClientBeginUpdateOptions) (resp azfake.PollerResponder[armcontainerregistry.TasksClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTasksServerTransport creates a new instance of TasksServerTransport with the provided implementation.
// The returned TasksServerTransport instance is connected to an instance of armcontainerregistry.TasksClient by way of the
// undefined.Transporter field.
func NewTasksServerTransport(srv *TasksServer) *TasksServerTransport {
	return &TasksServerTransport{srv: srv}
}

// TasksServerTransport connects instances of armcontainerregistry.TasksClient to instances of TasksServer.
// Don't use this type directly, use NewTasksServerTransport instead.
type TasksServerTransport struct {
	srv          *TasksServer
	beginCreate  *azfake.PollerResponder[armcontainerregistry.TasksClientCreateResponse]
	beginDelete  *azfake.PollerResponder[armcontainerregistry.TasksClientDeleteResponse]
	newListPager *azfake.PagerResponder[armcontainerregistry.TasksClientListResponse]
	beginUpdate  *azfake.PollerResponder[armcontainerregistry.TasksClientUpdateResponse]
}

// Do implements the policy.Transporter interface for TasksServerTransport.
func (t *TasksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TasksClient.BeginCreate":
		resp, err = t.dispatchBeginCreate(req)
	case "TasksClient.BeginDelete":
		resp, err = t.dispatchBeginDelete(req)
	case "TasksClient.Get":
		resp, err = t.dispatchGet(req)
	case "TasksClient.GetDetails":
		resp, err = t.dispatchGetDetails(req)
	case "TasksClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	case "TasksClient.BeginUpdate":
		resp, err = t.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TasksServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	if t.beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.Task](req)
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
		taskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginCreate(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, taskNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginCreate = &respr
	}

	resp, err := server.PollerResponderNext(t.beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginCreate) {
		t.beginCreate = nil
	}

	return resp, nil
}

func (t *TasksServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if t.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		taskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, taskNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(t.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginDelete) {
		t.beginDelete = nil
	}

	return resp, nil
}

func (t *TasksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	taskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, taskNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Task, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TasksServerTransport) dispatchGetDetails(req *http.Request) (*http.Response, error) {
	if t.srv.GetDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listDetails`
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
	taskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.GetDetails(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, taskNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Task, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TasksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if t.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tasks`
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
		resp := t.srv.NewListPager(resourceGroupNameUnescaped, registryNameUnescaped, nil)
		t.newListPager = &resp
		server.PagerResponderInjectNextLinks(t.newListPager, req, func(page *armcontainerregistry.TasksClientListResponse, createLink func() string) {
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

func (t *TasksServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	if t.beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerRegistry/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerregistry.TaskUpdateParameters](req)
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
		taskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, registryNameUnescaped, taskNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		t.beginUpdate = &respr
	}

	resp, err := server.PollerResponderNext(t.beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(t.beginUpdate) {
		t.beginUpdate = nil
	}

	return resp, nil
}
