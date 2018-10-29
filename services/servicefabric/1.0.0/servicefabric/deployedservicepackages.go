package servicefabric

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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// DeployedServicePackagesClient is the client for the DeployedServicePackages methods of the Servicefabric service.
type DeployedServicePackagesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewDeployedServicePackagesClient creates an instance of the DeployedServicePackagesClient client.
func NewDeployedServicePackagesClient(timeout *int32) DeployedServicePackagesClient {
	return NewDeployedServicePackagesClientWithBaseURI(DefaultBaseURI, timeout)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewDeployedServicePackagesClientWithBaseURI creates an instance of the DeployedServicePackagesClient client.
func NewDeployedServicePackagesClientWithBaseURI(baseURI string, timeout *int32) DeployedServicePackagesClient {
	return DeployedServicePackagesClient{NewWithBaseURI(baseURI, timeout)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Get get deployed service packages
// Parameters:
// nodeName - the name of the node
// applicationName - the name of the application
func (client DeployedServicePackagesClient) Get(ctx context.Context, nodeName string, applicationName string) (result ListDeployedServicePackage, err error) {
	req, err := client.GetPreparer(ctx, nodeName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedServicePackagesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedServicePackagesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedServicePackagesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// GetPreparer prepares the Get request.
func (client DeployedServicePackagesClient) GetPreparer(ctx context.Context, nodeName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":	applicationName,
		"nodeName":		autorest.Encode("path", nodeName),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Nodes/{nodeName}/$/GetApplications/{applicationName}/$/GetServicePackages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeployedServicePackagesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeployedServicePackagesClient) GetResponder(resp *http.Response) (result ListDeployedServicePackage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
