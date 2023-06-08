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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"net/http"
	"reflect"
)

// SubscriptionServer is a fake server for instances of the armsubscriptions.SubscriptionClient type.
type SubscriptionServer struct {
	// CheckResourceName is the fake for method SubscriptionClient.CheckResourceName
	// HTTP status codes to indicate success: http.StatusOK
	CheckResourceName func(ctx context.Context, options *armsubscriptions.SubscriptionClientCheckResourceNameOptions) (resp azfake.Responder[armsubscriptions.SubscriptionClientCheckResourceNameResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionServerTransport creates a new instance of SubscriptionServerTransport with the provided implementation.
// The returned SubscriptionServerTransport instance is connected to an instance of armsubscriptions.SubscriptionClient by way of the
// undefined.Transporter field.
func NewSubscriptionServerTransport(srv *SubscriptionServer) *SubscriptionServerTransport {
	return &SubscriptionServerTransport{srv: srv}
}

// SubscriptionServerTransport connects instances of armsubscriptions.SubscriptionClient to instances of SubscriptionServer.
// Don't use this type directly, use NewSubscriptionServerTransport instead.
type SubscriptionServerTransport struct {
	srv *SubscriptionServer
}

// Do implements the policy.Transporter interface for SubscriptionServerTransport.
func (s *SubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubscriptionClient.CheckResourceName":
		resp, err = s.dispatchCheckResourceName(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubscriptionServerTransport) dispatchCheckResourceName(req *http.Request) (*http.Response, error) {
	if s.srv.CheckResourceName == nil {
		return nil, &nonRetriableError{errors.New("method CheckResourceName not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[armsubscriptions.ResourceName](req)
	if err != nil {
		return nil, err
	}
	var options *armsubscriptions.SubscriptionClientCheckResourceNameOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armsubscriptions.SubscriptionClientCheckResourceNameOptions{
			ResourceNameDefinition: &body,
		}
	}
	respr, errRespr := s.srv.CheckResourceName(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckResourceNameResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
