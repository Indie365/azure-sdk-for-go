package batch

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CertificateClient is the client for the Certificate methods of the Batch service.
type CertificateClient struct {
	BaseClient
}

// NewCertificateClient creates an instance of the CertificateClient client.
func NewCertificateClient(subscriptionID string) CertificateClient {
	return NewCertificateClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificateClientWithBaseURI creates an instance of the CertificateClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCertificateClientWithBaseURI(baseURI string, subscriptionID string) CertificateClient {
	return CertificateClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CancelDeletion if you try to delete a certificate that is being used by a pool or compute node, the status of the
// certificate changes to deleteFailed. If you decide that you want to continue using the certificate, you can use this
// operation to set the status of the certificate back to active. If you intend to delete the certificate, you do not
// need to run this operation after the deletion failed. You must make sure that the certificate is not being used by
// any resources, and then you can try again to delete the certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// certificateName - the identifier for the certificate. This must be made up of algorithm and thumbprint
// separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
func (client CertificateClient) CancelDeletion(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (result Certificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.CancelDeletion")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.MaxLength, Rule: 45, Chain: nil},
				{Target: "certificateName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "certificateName", Name: validation.Pattern, Rule: `^[\w]+-[\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "CancelDeletion", err.Error())
	}

	req, err := client.CancelDeletionPreparer(ctx, resourceGroupName, accountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "CancelDeletion", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelDeletionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "CancelDeletion", resp, "Failure sending request")
		return
	}

	result, err = client.CancelDeletionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "CancelDeletion", resp, "Failure responding to request")
	}

	return
}

// CancelDeletionPreparer prepares the CancelDeletion request.
func (client CertificateClient) CancelDeletionPreparer(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}/cancelDelete", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelDeletionSender sends the CancelDeletion request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) CancelDeletionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelDeletionResponder handles the response to the CancelDeletion request. The method always
// closes the http.Response Body.
func (client CertificateClient) CancelDeletionResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a new certificate inside the specified account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// certificateName - the identifier for the certificate. This must be made up of algorithm and thumbprint
// separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
// parameters - additional parameters for certificate creation.
// ifMatch - the entity state (ETag) version of the certificate to update. A value of "*" can be used to apply
// the operation only if the certificate already exists. If omitted, this operation will always be applied.
// ifNoneMatch - set to '*' to allow a new certificate to be created, but to prevent updating an existing
// certificate. Other values will be ignored.
func (client CertificateClient) Create(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, ifMatch string, ifNoneMatch string) (result CertificateCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.MaxLength, Rule: 45, Chain: nil},
				{Target: "certificateName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "certificateName", Name: validation.Pattern, Rule: `^[\w]+-[\w]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.CertificateCreateOrUpdateProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.CertificateCreateOrUpdateProperties.Data", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, certificateName, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client CertificateClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) CreateSender(req *http.Request) (future CertificateCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CertificateClient) CreateResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// certificateName - the identifier for the certificate. This must be made up of algorithm and thumbprint
// separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
func (client CertificateClient) Delete(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (result CertificateDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.MaxLength, Rule: 45, Chain: nil},
				{Target: "certificateName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "certificateName", Name: validation.Pattern, Rule: `^[\w]+-[\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CertificateClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) DeleteSender(req *http.Request) (future CertificateDeleteFuture, err error) {
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
func (client CertificateClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// certificateName - the identifier for the certificate. This must be made up of algorithm and thumbprint
// separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
func (client CertificateClient) Get(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (result Certificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.MaxLength, Rule: 45, Chain: nil},
				{Target: "certificateName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "certificateName", Name: validation.Pattern, Rule: `^[\w]+-[\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CertificateClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CertificateClient) GetResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBatchAccount lists all of the certificates in the specified account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// maxresults - the maximum number of items to return in the response.
// selectParameter - comma separated list of properties that should be returned. e.g.
// "properties/provisioningState". Only top level properties under properties/ are valid for selection.
// filter - oData filter expression. Valid properties for filtering are "properties/provisioningState",
// "properties/provisioningStateTransitionTime", "name".
func (client CertificateClient) ListByBatchAccount(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result ListCertificatesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.lcr.Response.Response != nil {
				sc = result.lcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "ListByBatchAccount", err.Error())
	}

	result.fn = client.listByBatchAccountNextResults
	req, err := client.ListByBatchAccountPreparer(ctx, resourceGroupName, accountName, maxresults, selectParameter, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "ListByBatchAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.lcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "ListByBatchAccount", resp, "Failure sending request")
		return
	}

	result.lcr, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "ListByBatchAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBatchAccountPreparer prepares the ListByBatchAccount request.
func (client CertificateClient) ListByBatchAccountPreparer(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBatchAccountSender sends the ListByBatchAccount request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) ListByBatchAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByBatchAccountResponder handles the response to the ListByBatchAccount request. The method always
// closes the http.Response Body.
func (client CertificateClient) ListByBatchAccountResponder(resp *http.Response) (result ListCertificatesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBatchAccountNextResults retrieves the next set of results, if any.
func (client CertificateClient) listByBatchAccountNextResults(ctx context.Context, lastResults ListCertificatesResult) (result ListCertificatesResult, err error) {
	req, err := lastResults.listCertificatesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.CertificateClient", "listByBatchAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.CertificateClient", "listByBatchAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "listByBatchAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBatchAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client CertificateClient) ListByBatchAccountComplete(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result ListCertificatesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBatchAccount(ctx, resourceGroupName, accountName, maxresults, selectParameter, filter)
	return
}

// Update updates the properties of an existing certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// certificateName - the identifier for the certificate. This must be made up of algorithm and thumbprint
// separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
// parameters - certificate entity to update.
// ifMatch - the entity state (ETag) version of the certificate to update. This value can be omitted or set to
// "*" to apply the operation unconditionally.
func (client CertificateClient) Update(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, ifMatch string) (result Certificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.MaxLength, Rule: 45, Chain: nil},
				{Target: "certificateName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "certificateName", Name: validation.Pattern, Rule: `^[\w]+-[\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.CertificateClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, certificateName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.CertificateClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client CertificateClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client CertificateClient) UpdateResponder(resp *http.Response) (result Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
