package web

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TopLevelDomainsClient is the webSite Management Client
type TopLevelDomainsClient struct {
	BaseClient
}

// NewTopLevelDomainsClient creates an instance of the TopLevelDomainsClient client.
func NewTopLevelDomainsClient(subscriptionID string) TopLevelDomainsClient {
	return NewTopLevelDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTopLevelDomainsClientWithBaseURI creates an instance of the TopLevelDomainsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTopLevelDomainsClientWithBaseURI(baseURI string, subscriptionID string) TopLevelDomainsClient {
	return TopLevelDomainsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetGetTopLevelDomains sends the get get top level domains request.
func (client TopLevelDomainsClient) GetGetTopLevelDomains(ctx context.Context) (result TopLevelDomainCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.GetGetTopLevelDomains")
		defer func() {
			sc := -1
			if result.tldc.Response.Response != nil {
				sc = result.tldc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getGetTopLevelDomainsNextResults
	req, err := client.GetGetTopLevelDomainsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "GetGetTopLevelDomains", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGetTopLevelDomainsSender(req)
	if err != nil {
		result.tldc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "GetGetTopLevelDomains", resp, "Failure sending request")
		return
	}

	result.tldc, err = client.GetGetTopLevelDomainsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "GetGetTopLevelDomains", resp, "Failure responding to request")
	}

	return
}

// GetGetTopLevelDomainsPreparer prepares the GetGetTopLevelDomains request.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetGetTopLevelDomainsSender sends the GetGetTopLevelDomains request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetGetTopLevelDomainsResponder handles the response to the GetGetTopLevelDomains request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsResponder(resp *http.Response) (result TopLevelDomainCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getGetTopLevelDomainsNextResults retrieves the next set of results, if any.
func (client TopLevelDomainsClient) getGetTopLevelDomainsNextResults(ctx context.Context, lastResults TopLevelDomainCollection) (result TopLevelDomainCollection, err error) {
	req, err := lastResults.topLevelDomainCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "getGetTopLevelDomainsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetGetTopLevelDomainsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "getGetTopLevelDomainsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetGetTopLevelDomainsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "getGetTopLevelDomainsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetGetTopLevelDomainsComplete enumerates all values, automatically crossing page boundaries as required.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsComplete(ctx context.Context) (result TopLevelDomainCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.GetGetTopLevelDomains")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetGetTopLevelDomains(ctx)
	return
}

// GetTopLevelDomain sends the get top level domain request.
// Parameters:
// name - name of the top level domain
func (client TopLevelDomainsClient) GetTopLevelDomain(ctx context.Context, name string) (result TopLevelDomain, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.GetTopLevelDomain")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetTopLevelDomainPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "GetTopLevelDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTopLevelDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "GetTopLevelDomain", resp, "Failure sending request")
		return
	}

	result, err = client.GetTopLevelDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "GetTopLevelDomain", resp, "Failure responding to request")
	}

	return
}

// GetTopLevelDomainPreparer prepares the GetTopLevelDomain request.
func (client TopLevelDomainsClient) GetTopLevelDomainPreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTopLevelDomainSender sends the GetTopLevelDomain request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) GetTopLevelDomainSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetTopLevelDomainResponder handles the response to the GetTopLevelDomain request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) GetTopLevelDomainResponder(resp *http.Response) (result TopLevelDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListTopLevelDomainAgreements sends the list top level domain agreements request.
// Parameters:
// name - name of the top level domain
// agreementOption - domain agreement options
func (client TopLevelDomainsClient) ListTopLevelDomainAgreements(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption) (result TldLegalAgreementCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.ListTopLevelDomainAgreements")
		defer func() {
			sc := -1
			if result.tlac.Response.Response != nil {
				sc = result.tlac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listTopLevelDomainAgreementsNextResults
	req, err := client.ListTopLevelDomainAgreementsPreparer(ctx, name, agreementOption)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "ListTopLevelDomainAgreements", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListTopLevelDomainAgreementsSender(req)
	if err != nil {
		result.tlac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "ListTopLevelDomainAgreements", resp, "Failure sending request")
		return
	}

	result.tlac, err = client.ListTopLevelDomainAgreementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "ListTopLevelDomainAgreements", resp, "Failure responding to request")
	}

	return
}

// ListTopLevelDomainAgreementsPreparer prepares the ListTopLevelDomainAgreements request.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsPreparer(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}/listAgreements", pathParameters),
		autorest.WithJSON(agreementOption),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListTopLevelDomainAgreementsSender sends the ListTopLevelDomainAgreements request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListTopLevelDomainAgreementsResponder handles the response to the ListTopLevelDomainAgreements request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsResponder(resp *http.Response) (result TldLegalAgreementCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listTopLevelDomainAgreementsNextResults retrieves the next set of results, if any.
func (client TopLevelDomainsClient) listTopLevelDomainAgreementsNextResults(ctx context.Context, lastResults TldLegalAgreementCollection) (result TldLegalAgreementCollection, err error) {
	req, err := lastResults.tldLegalAgreementCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listTopLevelDomainAgreementsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListTopLevelDomainAgreementsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listTopLevelDomainAgreementsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListTopLevelDomainAgreementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listTopLevelDomainAgreementsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListTopLevelDomainAgreementsComplete enumerates all values, automatically crossing page boundaries as required.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsComplete(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption) (result TldLegalAgreementCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.ListTopLevelDomainAgreements")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListTopLevelDomainAgreements(ctx, name, agreementOption)
	return
}
