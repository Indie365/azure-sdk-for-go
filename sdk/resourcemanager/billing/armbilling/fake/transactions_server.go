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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"net/url"
	"regexp"
)

// TransactionsServer is a fake server for instances of the armbilling.TransactionsClient type.
type TransactionsServer struct {
	// NewListByInvoicePager is the fake for method TransactionsClient.NewListByInvoicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInvoicePager func(billingAccountName string, invoiceName string, options *armbilling.TransactionsClientListByInvoiceOptions) (resp azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceResponse])
}

// NewTransactionsServerTransport creates a new instance of TransactionsServerTransport with the provided implementation.
// The returned TransactionsServerTransport instance is connected to an instance of armbilling.TransactionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTransactionsServerTransport(srv *TransactionsServer) *TransactionsServerTransport {
	return &TransactionsServerTransport{
		srv:                   srv,
		newListByInvoicePager: newTracker[azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceResponse]](),
	}
}

// TransactionsServerTransport connects instances of armbilling.TransactionsClient to instances of TransactionsServer.
// Don't use this type directly, use NewTransactionsServerTransport instead.
type TransactionsServerTransport struct {
	srv                   *TransactionsServer
	newListByInvoicePager *tracker[azfake.PagerResponder[armbilling.TransactionsClientListByInvoiceResponse]]
}

// Do implements the policy.Transporter interface for TransactionsServerTransport.
func (t *TransactionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TransactionsClient.NewListByInvoicePager":
		resp, err = t.dispatchNewListByInvoicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionsServerTransport) dispatchNewListByInvoicePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByInvoicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInvoicePager not implemented")}
	}
	newListByInvoicePager := t.newListByInvoicePager.get(req)
	if newListByInvoicePager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoices/(?P<invoiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transactions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		invoiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByInvoicePager(billingAccountNameParam, invoiceNameParam, nil)
		newListByInvoicePager = &resp
		t.newListByInvoicePager.add(req, newListByInvoicePager)
		server.PagerResponderInjectNextLinks(newListByInvoicePager, req, func(page *armbilling.TransactionsClientListByInvoiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInvoicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByInvoicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInvoicePager) {
		t.newListByInvoicePager.remove(req)
	}
	return resp, nil
}
