// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
	"net/http"
	"net/url"
	"regexp"
)

// SKUServer is a fake server for instances of the armdevopsinfrastructure.SKUClient type.
type SKUServer struct {
	// NewListByLocationPager is the fake for method SKUClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(locationName string, options *armdevopsinfrastructure.SKUClientListByLocationOptions) (resp azfake.PagerResponder[armdevopsinfrastructure.SKUClientListByLocationResponse])
}

// NewSKUServerTransport creates a new instance of SKUServerTransport with the provided implementation.
// The returned SKUServerTransport instance is connected to an instance of armdevopsinfrastructure.SKUClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSKUServerTransport(srv *SKUServer) *SKUServerTransport {
	return &SKUServerTransport{
		srv:                    srv,
		newListByLocationPager: newTracker[azfake.PagerResponder[armdevopsinfrastructure.SKUClientListByLocationResponse]](),
	}
}

// SKUServerTransport connects instances of armdevopsinfrastructure.SKUClient to instances of SKUServer.
// Don't use this type directly, use NewSKUServerTransport instead.
type SKUServerTransport struct {
	srv                    *SKUServer
	newListByLocationPager *tracker[azfake.PagerResponder[armdevopsinfrastructure.SKUClientListByLocationResponse]]
}

// Do implements the policy.Transporter interface for SKUServerTransport.
func (s *SKUServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SKUServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "SKUClient.NewListByLocationPager":
		resp, err = s.dispatchNewListByLocationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *SKUServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	newListByLocationPager := s.newListByLocationPager.get(req)
	if newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevOpsInfrastructure/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByLocationPager(locationNameParam, nil)
		newListByLocationPager = &resp
		s.newListByLocationPager.add(req, newListByLocationPager)
		server.PagerResponderInjectNextLinks(newListByLocationPager, req, func(page *armdevopsinfrastructure.SKUClientListByLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocationPager) {
		s.newListByLocationPager.remove(req)
	}
	return resp, nil
}
