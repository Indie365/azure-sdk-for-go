package kubernetesconfiguration

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

// SourceControlConfigurationsClient is the use these APIs to create Source Control Configuration resources through
// ARM, for Kubernetes Clusters.
type SourceControlConfigurationsClient struct {
	BaseClient
}

// NewSourceControlConfigurationsClient creates an instance of the SourceControlConfigurationsClient client.
func NewSourceControlConfigurationsClient(subscriptionID string) SourceControlConfigurationsClient {
	return NewSourceControlConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSourceControlConfigurationsClientWithBaseURI creates an instance of the SourceControlConfigurationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSourceControlConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SourceControlConfigurationsClient {
	return SourceControlConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new Kubernetes Source Control Configuration.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterRp - the Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or
// Microsoft.Kubernetes (for OnPrem K8S clusters).
// clusterResourceName - the Kubernetes cluster resource name - either managedClusters (for AKS clusters) or
// connectedClusters (for OnPrem K8S clusters).
// clusterName - the name of the kubernetes cluster.
// sourceControlConfigurationName - name of the Source Control Configuration.
// sourceControlConfiguration - properties necessary to Create KubernetesConfiguration.
func (client SourceControlConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, sourceControlConfiguration SourceControlConfiguration) (result SourceControlConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName, sourceControlConfiguration)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SourceControlConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, sourceControlConfiguration SourceControlConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":                    autorest.Encode("path", clusterName),
		"clusterResourceName":            autorest.Encode("path", clusterResourceName),
		"clusterRp":                      autorest.Encode("path", clusterRp),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"sourceControlConfigurationName": autorest.Encode("path", sourceControlConfigurationName),
		"subscriptionId":                 autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}", pathParameters),
		autorest.WithJSON(sourceControlConfiguration),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlConfigurationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SourceControlConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result SourceControlConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete this will delete the YAML file used to set up the Source control configuration, thus stopping future sync
// from the source repo.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterRp - the Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or
// Microsoft.Kubernetes (for OnPrem K8S clusters).
// clusterResourceName - the Kubernetes cluster resource name - either managedClusters (for AKS clusters) or
// connectedClusters (for OnPrem K8S clusters).
// clusterName - the name of the kubernetes cluster.
// sourceControlConfigurationName - name of the Source Control Configuration.
func (client SourceControlConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) (result SourceControlConfigurationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SourceControlConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":                    autorest.Encode("path", clusterName),
		"clusterResourceName":            autorest.Encode("path", clusterResourceName),
		"clusterRp":                      autorest.Encode("path", clusterRp),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"sourceControlConfigurationName": autorest.Encode("path", sourceControlConfigurationName),
		"subscriptionId":                 autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlConfigurationsClient) DeleteSender(req *http.Request) (future SourceControlConfigurationsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SourceControlConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details of the Source Control Configuration.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterRp - the Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or
// Microsoft.Kubernetes (for OnPrem K8S clusters).
// clusterResourceName - the Kubernetes cluster resource name - either managedClusters (for AKS clusters) or
// connectedClusters (for OnPrem K8S clusters).
// clusterName - the name of the kubernetes cluster.
// sourceControlConfigurationName - name of the Source Control Configuration.
func (client SourceControlConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) (result SourceControlConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SourceControlConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":                    autorest.Encode("path", clusterName),
		"clusterResourceName":            autorest.Encode("path", clusterResourceName),
		"clusterRp":                      autorest.Encode("path", clusterRp),
		"resourceGroupName":              autorest.Encode("path", resourceGroupName),
		"sourceControlConfigurationName": autorest.Encode("path", sourceControlConfigurationName),
		"subscriptionId":                 autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SourceControlConfigurationsClient) GetResponder(resp *http.Response) (result SourceControlConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all Source Control Configurations.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterRp - the Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or
// Microsoft.Kubernetes (for OnPrem K8S clusters).
// clusterResourceName - the Kubernetes cluster resource name - either managedClusters (for AKS clusters) or
// connectedClusters (for OnPrem K8S clusters).
// clusterName - the name of the kubernetes cluster.
func (client SourceControlConfigurationsClient) List(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result SourceControlConfigurationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.sccl.Response.Response != nil {
				sc = result.sccl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sccl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.sccl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SourceControlConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":         autorest.Encode("path", clusterName),
		"clusterResourceName": autorest.Encode("path", clusterResourceName),
		"clusterRp":           autorest.Encode("path", clusterRp),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SourceControlConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SourceControlConfigurationsClient) ListResponder(resp *http.Response) (result SourceControlConfigurationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SourceControlConfigurationsClient) listNextResults(ctx context.Context, lastResults SourceControlConfigurationList) (result SourceControlConfigurationList, err error) {
	req, err := lastResults.sourceControlConfigurationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SourceControlConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string) (result SourceControlConfigurationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName)
	return
}
