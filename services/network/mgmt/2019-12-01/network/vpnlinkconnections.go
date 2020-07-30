package network

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

// VpnLinkConnectionsClient is the network Client
type VpnLinkConnectionsClient struct {
	BaseClient
}

// NewVpnLinkConnectionsClient creates an instance of the VpnLinkConnectionsClient client.
func NewVpnLinkConnectionsClient(subscriptionID string) VpnLinkConnectionsClient {
	return NewVpnLinkConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnLinkConnectionsClientWithBaseURI creates an instance of the VpnLinkConnectionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewVpnLinkConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VpnLinkConnectionsClient {
	return VpnLinkConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByVpnConnection retrieves all vpn site link connections for a particular virtual wan vpn gateway vpn connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
func (client VpnLinkConnectionsClient) ListByVpnConnection(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result ListVpnSiteLinkConnectionsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnLinkConnectionsClient.ListByVpnConnection")
		defer func() {
			sc := -1
			if result.lvslcr.Response.Response != nil {
				sc = result.lvslcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnConnectionNextResults
	req, err := client.ListByVpnConnectionPreparer(ctx, resourceGroupName, gatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ListByVpnConnection", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnConnectionSender(req)
	if err != nil {
		result.lvslcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ListByVpnConnection", resp, "Failure sending request")
		return
	}

	result.lvslcr, err = client.ListByVpnConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "ListByVpnConnection", resp, "Failure responding to request")
	}
	if result.lvslcr.hasNextLink() && result.lvslcr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByVpnConnectionPreparer prepares the ListByVpnConnection request.
func (client VpnLinkConnectionsClient) ListByVpnConnectionPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnConnectionSender sends the ListByVpnConnection request. The method will close the
// http.Response Body if it receives an error.
func (client VpnLinkConnectionsClient) ListByVpnConnectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVpnConnectionResponder handles the response to the ListByVpnConnection request. The method always
// closes the http.Response Body.
func (client VpnLinkConnectionsClient) ListByVpnConnectionResponder(resp *http.Response) (result ListVpnSiteLinkConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnConnectionNextResults retrieves the next set of results, if any.
func (client VpnLinkConnectionsClient) listByVpnConnectionNextResults(ctx context.Context, lastResults ListVpnSiteLinkConnectionsResult) (result ListVpnSiteLinkConnectionsResult, err error) {
	req, err := lastResults.listVpnSiteLinkConnectionsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "listByVpnConnectionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnConnectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "listByVpnConnectionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnConnectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnLinkConnectionsClient", "listByVpnConnectionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnConnectionComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnLinkConnectionsClient) ListByVpnConnectionComplete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result ListVpnSiteLinkConnectionsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnLinkConnectionsClient.ListByVpnConnection")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnConnection(ctx, resourceGroupName, gatewayName, connectionName)
	return
}
