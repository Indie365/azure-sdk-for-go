package costmanagement

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

// DimensionsClient is the client for the Dimensions methods of the Costmanagement service.
type DimensionsClient struct {
	BaseClient
}

// NewDimensionsClient creates an instance of the DimensionsClient client.
func NewDimensionsClient(subscriptionID string) DimensionsClient {
	return NewDimensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDimensionsClientWithBaseURI creates an instance of the DimensionsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDimensionsClientWithBaseURI(baseURI string, subscriptionID string) DimensionsClient {
	return DimensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ByExternalCloudProviderType lists the dimensions by the external cloud provider type.
// Parameters:
// externalCloudProviderType - the external cloud provider type associated with dimension/query operations.
// This includes 'externalSubscriptions' for linked account and 'externalBillingAccounts' for consolidated
// account.
// externalCloudProviderID - this can be '{externalSubscriptionId}' for linked account or
// '{externalBillingAccountId}' for consolidated account used with dimension/query operations.
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) ByExternalCloudProviderType(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ByExternalCloudProviderType")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "ByExternalCloudProviderType", err.Error())
	}

	req, err := client.ByExternalCloudProviderTypePreparer(ctx, externalCloudProviderType, externalCloudProviderID, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ByExternalCloudProviderType", nil, "Failure preparing request")
		return
	}

	resp, err := client.ByExternalCloudProviderTypeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ByExternalCloudProviderType", resp, "Failure sending request")
		return
	}

	result, err = client.ByExternalCloudProviderTypeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "ByExternalCloudProviderType", resp, "Failure responding to request")
		return
	}

	return
}

// ByExternalCloudProviderTypePreparer prepares the ByExternalCloudProviderType request.
func (client DimensionsClient) ByExternalCloudProviderTypePreparer(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"externalCloudProviderId":   autorest.Encode("path", externalCloudProviderID),
		"externalCloudProviderType": autorest.Encode("path", externalCloudProviderType),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ByExternalCloudProviderTypeSender sends the ByExternalCloudProviderType request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ByExternalCloudProviderTypeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ByExternalCloudProviderTypeResponder handles the response to the ByExternalCloudProviderType request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ByExternalCloudProviderTypeResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the dimensions by the defined scope.
// Parameters:
// scope - the scope associated with dimension operations. This includes '/subscriptions/{subscriptionId}/' for
// subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department
// scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for
// Management Group scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for
// billingProfile scope,
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for
// partners.
// filter - may be used to filter dimensions by properties/category, properties/usageStart,
// properties/usageEnd. Supported operators are 'eq','lt', 'gt', 'le', 'ge'.
// expand - may be used to expand the properties/data within a dimension category. By default, data is not
// included when listing dimensions.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N dimension data.
func (client DimensionsClient) List(ctx context.Context, scope string, filter string, expand string, skiptoken string, top *int32) (result DimensionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.DimensionsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, scope, filter, expand, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.DimensionsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DimensionsClient) ListPreparer(ctx context.Context, scope string, filter string, expand string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CostManagement/dimensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListResponder(resp *http.Response) (result DimensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
