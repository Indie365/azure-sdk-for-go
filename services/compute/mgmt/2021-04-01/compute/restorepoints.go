package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RestorePointsClient is the compute Client
type RestorePointsClient struct {
	BaseClient
}

// NewRestorePointsClient creates an instance of the RestorePointsClient client.
func NewRestorePointsClient(subscriptionID string) RestorePointsClient {
	return NewRestorePointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorePointsClientWithBaseURI creates an instance of the RestorePointsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRestorePointsClientWithBaseURI(baseURI string, subscriptionID string) RestorePointsClient {
	return RestorePointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create the operation to create the restore point. Updating properties of an existing restore point is not allowed
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the restore point collection.
// restorePointName - the name of the restore point.
// parameters - parameters supplied to the Create restore point operation.
func (client RestorePointsClient) Create(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, parameters RestorePoint) (result RestorePointsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorePointsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RestorePointProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey.SecretURL", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey.SourceVault", Name: validation.Null, Rule: true, Chain: nil},
									}},
									{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey.KeyURL", Name: validation.Null, Rule: true, Chain: nil},
											{Target: "parameters.RestorePointProperties.SourceMetadata.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey.SourceVault", Name: validation.Null, Rule: true, Chain: nil},
										}},
								}},
							}},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.RestorePointsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, restorePointCollectionName, restorePointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RestorePointsClient) CreatePreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, parameters RestorePoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"restorePointName":           autorest.Encode("path", restorePointName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{restorePointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RestorePointsClient) CreateSender(req *http.Request) (future RestorePointsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RestorePointsClient) CreateResponder(resp *http.Response) (result RestorePoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the restore point.
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the Restore Point Collection.
// restorePointName - the name of the restore point.
func (client RestorePointsClient) Delete(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string) (result RestorePointsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorePointsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, restorePointCollectionName, restorePointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RestorePointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"restorePointName":           autorest.Encode("path", restorePointName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{restorePointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RestorePointsClient) DeleteSender(req *http.Request) (future RestorePointsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RestorePointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get the restore point.
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the restore point collection.
// restorePointName - the name of the restore point.
func (client RestorePointsClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string) (result RestorePoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorePointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, restorePointCollectionName, restorePointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RestorePointsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RestorePointsClient) GetPreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"restorePointName":           autorest.Encode("path", restorePointName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{restorePointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RestorePointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RestorePointsClient) GetResponder(resp *http.Response) (result RestorePoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
