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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
	"net/http"
	"net/url"
	"regexp"
)

// IscsiServersServer is a fake server for instances of the armstorsimple1200series.IscsiServersClient type.
type IscsiServersServer struct {
	// BeginBackupNow is the fake for method IscsiServersClient.BeginBackupNow
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginBackupNow func(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientBeginBackupNowOptions) (resp azfake.PollerResponder[armstorsimple1200series.IscsiServersClientBackupNowResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method IscsiServersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, iscsiServer armstorsimple1200series.ISCSIServer, options *armstorsimple1200series.IscsiServersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armstorsimple1200series.IscsiServersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method IscsiServersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientBeginDeleteOptions) (resp azfake.PollerResponder[armstorsimple1200series.IscsiServersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IscsiServersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientGetOptions) (resp azfake.Responder[armstorsimple1200series.IscsiServersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDevicePager is the fake for method IscsiServersClient.NewListByDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDevicePager func(deviceName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientListByDeviceOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListByDeviceResponse])

	// NewListByManagerPager is the fake for method IscsiServersClient.NewListByManagerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagerPager func(resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientListByManagerOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListByManagerResponse])

	// NewListMetricDefinitionPager is the fake for method IscsiServersClient.NewListMetricDefinitionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricDefinitionPager func(deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientListMetricDefinitionOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListMetricDefinitionResponse])

	// NewListMetricsPager is the fake for method IscsiServersClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(deviceName string, iscsiServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.IscsiServersClientListMetricsOptions) (resp azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListMetricsResponse])
}

// NewIscsiServersServerTransport creates a new instance of IscsiServersServerTransport with the provided implementation.
// The returned IscsiServersServerTransport instance is connected to an instance of armstorsimple1200series.IscsiServersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIscsiServersServerTransport(srv *IscsiServersServer) *IscsiServersServerTransport {
	return &IscsiServersServerTransport{
		srv:                          srv,
		beginBackupNow:               newTracker[azfake.PollerResponder[armstorsimple1200series.IscsiServersClientBackupNowResponse]](),
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armstorsimple1200series.IscsiServersClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armstorsimple1200series.IscsiServersClientDeleteResponse]](),
		newListByDevicePager:         newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListByDeviceResponse]](),
		newListByManagerPager:        newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListByManagerResponse]](),
		newListMetricDefinitionPager: newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListMetricDefinitionResponse]](),
		newListMetricsPager:          newTracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListMetricsResponse]](),
	}
}

// IscsiServersServerTransport connects instances of armstorsimple1200series.IscsiServersClient to instances of IscsiServersServer.
// Don't use this type directly, use NewIscsiServersServerTransport instead.
type IscsiServersServerTransport struct {
	srv                          *IscsiServersServer
	beginBackupNow               *tracker[azfake.PollerResponder[armstorsimple1200series.IscsiServersClientBackupNowResponse]]
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armstorsimple1200series.IscsiServersClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armstorsimple1200series.IscsiServersClientDeleteResponse]]
	newListByDevicePager         *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListByDeviceResponse]]
	newListByManagerPager        *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListByManagerResponse]]
	newListMetricDefinitionPager *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListMetricDefinitionResponse]]
	newListMetricsPager          *tracker[azfake.PagerResponder[armstorsimple1200series.IscsiServersClientListMetricsResponse]]
}

// Do implements the policy.Transporter interface for IscsiServersServerTransport.
func (i *IscsiServersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IscsiServersClient.BeginBackupNow":
		resp, err = i.dispatchBeginBackupNow(req)
	case "IscsiServersClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "IscsiServersClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "IscsiServersClient.Get":
		resp, err = i.dispatchGet(req)
	case "IscsiServersClient.NewListByDevicePager":
		resp, err = i.dispatchNewListByDevicePager(req)
	case "IscsiServersClient.NewListByManagerPager":
		resp, err = i.dispatchNewListByManagerPager(req)
	case "IscsiServersClient.NewListMetricDefinitionPager":
		resp, err = i.dispatchNewListMetricDefinitionPager(req)
	case "IscsiServersClient.NewListMetricsPager":
		resp, err = i.dispatchNewListMetricsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchBeginBackupNow(req *http.Request) (*http.Response, error) {
	if i.srv.BeginBackupNow == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginBackupNow not implemented")}
	}
	beginBackupNow := i.beginBackupNow.get(req)
	if beginBackupNow == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginBackupNow(req.Context(), deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginBackupNow = &respr
		i.beginBackupNow.add(req, beginBackupNow)
	}

	resp, err := server.PollerResponderNext(beginBackupNow, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginBackupNow.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginBackupNow) {
		i.beginBackupNow.remove(req)
	}

	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple1200series.ISCSIServer](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ISCSIServer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchNewListByDevicePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDevicePager not implemented")}
	}
	newListByDevicePager := i.newListByDevicePager.get(req)
	if newListByDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByDevicePager(deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByDevicePager = &resp
		i.newListByDevicePager.add(req, newListByDevicePager)
	}
	resp, err := server.PagerResponderNext(newListByDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDevicePager) {
		i.newListByDevicePager.remove(req)
	}
	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchNewListByManagerPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByManagerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagerPager not implemented")}
	}
	newListByManagerPager := i.newListByManagerPager.get(req)
	if newListByManagerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByManagerPager(resourceGroupNameParam, managerNameParam, nil)
		newListByManagerPager = &resp
		i.newListByManagerPager.add(req, newListByManagerPager)
	}
	resp, err := server.PagerResponderNext(newListByManagerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByManagerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByManagerPager) {
		i.newListByManagerPager.remove(req)
	}
	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchNewListMetricDefinitionPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListMetricDefinitionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricDefinitionPager not implemented")}
	}
	newListMetricDefinitionPager := i.newListMetricDefinitionPager.get(req)
	if newListMetricDefinitionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metricsDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListMetricDefinitionPager(deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListMetricDefinitionPager = &resp
		i.newListMetricDefinitionPager.add(req, newListMetricDefinitionPager)
	}
	resp, err := server.PagerResponderNext(newListMetricDefinitionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListMetricDefinitionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricDefinitionPager) {
		i.newListMetricDefinitionPager.remove(req)
	}
	return resp, nil
}

func (i *IscsiServersServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := i.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iscsiservers/(?P<iscsiServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		iscsiServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iscsiServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armstorsimple1200series.IscsiServersClientListMetricsOptions
		if filterParam != nil {
			options = &armstorsimple1200series.IscsiServersClientListMetricsOptions{
				Filter: filterParam,
			}
		}
		resp := i.srv.NewListMetricsPager(deviceNameParam, iscsiServerNameParam, resourceGroupNameParam, managerNameParam, options)
		newListMetricsPager = &resp
		i.newListMetricsPager.add(req, newListMetricsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListMetricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricsPager) {
		i.newListMetricsPager.remove(req)
	}
	return resp, nil
}
