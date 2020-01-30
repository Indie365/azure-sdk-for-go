package migrate

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

// DatabasesClient is the migrate your workloads to Azure.
type DatabasesClient struct {
	BaseClient
}

// NewDatabasesClient creates an instance of the DatabasesClient client.
func NewDatabasesClient(subscriptionID string, acceptLanguage string) DatabasesClient {
	return NewDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewDatabasesClientWithBaseURI creates an instance of the DatabasesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) DatabasesClient {
	return DatabasesClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// EnumerateDatabases sends the enumerate databases request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// continuationToken - the continuation token.
// pageSize - the number of items to be returned in a single page. This value is honored only if it is less
// than the 100.
func (client DatabasesClient) EnumerateDatabases(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result DatabaseCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.EnumerateDatabases")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnumerateDatabasesPreparer(ctx, resourceGroupName, migrateProjectName, continuationToken, pageSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabasesClient", "EnumerateDatabases", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnumerateDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.DatabasesClient", "EnumerateDatabases", resp, "Failure sending request")
		return
	}

	result, err = client.EnumerateDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabasesClient", "EnumerateDatabases", resp, "Failure responding to request")
	}

	return
}

// EnumerateDatabasesPreparer prepares the EnumerateDatabases request.
func (client DatabasesClient) EnumerateDatabasesPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(continuationToken) > 0 {
		queryParameters["continuationToken"] = autorest.Encode("query", continuationToken)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnumerateDatabasesSender sends the EnumerateDatabases request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) EnumerateDatabasesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// EnumerateDatabasesResponder handles the response to the EnumerateDatabases request. The method always
// closes the http.Response Body.
func (client DatabasesClient) EnumerateDatabasesResponder(resp *http.Response) (result DatabaseCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDatabase sends the get database request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// databaseName - unique name of a database in Azure migration hub.
func (client DatabasesClient) GetDatabase(ctx context.Context, resourceGroupName string, migrateProjectName string, databaseName string) (result Database, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabasesClient.GetDatabase")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDatabasePreparer(ctx, resourceGroupName, migrateProjectName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabasesClient", "GetDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.DatabasesClient", "GetDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.GetDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.DatabasesClient", "GetDatabase", resp, "Failure responding to request")
	}

	return
}

// GetDatabasePreparer prepares the GetDatabase request.
func (client DatabasesClient) GetDatabasePreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":       autorest.Encode("path", databaseName),
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databases/{databaseName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDatabaseSender sends the GetDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client DatabasesClient) GetDatabaseSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetDatabaseResponder handles the response to the GetDatabase request. The method always
// closes the http.Response Body.
func (client DatabasesClient) GetDatabaseResponder(resp *http.Response) (result Database, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
