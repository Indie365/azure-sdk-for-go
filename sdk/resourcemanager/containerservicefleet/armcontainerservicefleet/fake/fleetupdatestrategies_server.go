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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
	"net/http"
	"net/url"
	"regexp"
)

// FleetUpdateStrategiesServer is a fake server for instances of the armcontainerservicefleet.FleetUpdateStrategiesClient type.
type FleetUpdateStrategiesServer struct {
	// BeginCreateOrUpdate is the fake for method FleetUpdateStrategiesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, resource armcontainerservicefleet.FleetUpdateStrategy, options *armcontainerservicefleet.FleetUpdateStrategiesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FleetUpdateStrategiesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *armcontainerservicefleet.FleetUpdateStrategiesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FleetUpdateStrategiesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *armcontainerservicefleet.FleetUpdateStrategiesClientGetOptions) (resp azfake.Responder[armcontainerservicefleet.FleetUpdateStrategiesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFleetPager is the fake for method FleetUpdateStrategiesClient.NewListByFleetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFleetPager func(resourceGroupName string, fleetName string, options *armcontainerservicefleet.FleetUpdateStrategiesClientListByFleetOptions) (resp azfake.PagerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientListByFleetResponse])
}

// NewFleetUpdateStrategiesServerTransport creates a new instance of FleetUpdateStrategiesServerTransport with the provided implementation.
// The returned FleetUpdateStrategiesServerTransport instance is connected to an instance of armcontainerservicefleet.FleetUpdateStrategiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFleetUpdateStrategiesServerTransport(srv *FleetUpdateStrategiesServer) *FleetUpdateStrategiesServerTransport {
	return &FleetUpdateStrategiesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientDeleteResponse]](),
		newListByFleetPager: newTracker[azfake.PagerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientListByFleetResponse]](),
	}
}

// FleetUpdateStrategiesServerTransport connects instances of armcontainerservicefleet.FleetUpdateStrategiesClient to instances of FleetUpdateStrategiesServer.
// Don't use this type directly, use NewFleetUpdateStrategiesServerTransport instead.
type FleetUpdateStrategiesServerTransport struct {
	srv                 *FleetUpdateStrategiesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientDeleteResponse]]
	newListByFleetPager *tracker[azfake.PagerResponder[armcontainerservicefleet.FleetUpdateStrategiesClientListByFleetResponse]]
}

// Do implements the policy.Transporter interface for FleetUpdateStrategiesServerTransport.
func (f *FleetUpdateStrategiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FleetUpdateStrategiesClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FleetUpdateStrategiesClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FleetUpdateStrategiesClient.Get":
		resp, err = f.dispatchGet(req)
	case "FleetUpdateStrategiesClient.NewListByFleetPager":
		resp, err = f.dispatchNewListByFleetPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FleetUpdateStrategiesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateStrategies/(?P<updateStrategyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerservicefleet.FleetUpdateStrategy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateStrategyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateStrategyName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.FleetUpdateStrategiesClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armcontainerservicefleet.FleetUpdateStrategiesClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, updateStrategyNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FleetUpdateStrategiesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateStrategies/(?P<updateStrategyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		updateStrategyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateStrategyName")])
		if err != nil {
			return nil, err
		}
		var options *armcontainerservicefleet.FleetUpdateStrategiesClientBeginDeleteOptions
		if ifMatchParam != nil {
			options = &armcontainerservicefleet.FleetUpdateStrategiesClientBeginDeleteOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, fleetNameParam, updateStrategyNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FleetUpdateStrategiesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateStrategies/(?P<updateStrategyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
	if err != nil {
		return nil, err
	}
	updateStrategyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("updateStrategyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, fleetNameParam, updateStrategyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FleetUpdateStrategy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FleetUpdateStrategiesServerTransport) dispatchNewListByFleetPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByFleetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFleetPager not implemented")}
	}
	newListByFleetPager := f.newListByFleetPager.get(req)
	if newListByFleetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateStrategies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByFleetPager(resourceGroupNameParam, fleetNameParam, nil)
		newListByFleetPager = &resp
		f.newListByFleetPager.add(req, newListByFleetPager)
		server.PagerResponderInjectNextLinks(newListByFleetPager, req, func(page *armcontainerservicefleet.FleetUpdateStrategiesClientListByFleetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFleetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByFleetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFleetPager) {
		f.newListByFleetPager.remove(req)
	}
	return resp, nil
}
