package consumption

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
	"net/http"
)

// UsageDetailsByEnrollmentAccountClient is the consumption management client provides access to consumption resources
// for Azure Enterprise Subscriptions.
type UsageDetailsByEnrollmentAccountClient struct {
	BaseClient
}

// NewUsageDetailsByEnrollmentAccountClient creates an instance of the UsageDetailsByEnrollmentAccountClient client.
func NewUsageDetailsByEnrollmentAccountClient(subscriptionID string) UsageDetailsByEnrollmentAccountClient {
	return NewUsageDetailsByEnrollmentAccountClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageDetailsByEnrollmentAccountClientWithBaseURI creates an instance of the UsageDetailsByEnrollmentAccountClient
// client.
func NewUsageDetailsByEnrollmentAccountClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsByEnrollmentAccountClient {
	return UsageDetailsByEnrollmentAccountClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the usage details by enrollmentAccountId for a scope by current billing period. Usage details are
// available via this API only for May 1, 2014 or later.
//
// enrollmentAccountID is enrollmentAccount ID expand is may be used to expand the properties/additionalProperties
// or properties/meterDetails within a list of usage details. By default, these fields are not included when
// listing usage details. filter is may be used to filter usageDetails by properties/usageEnd (Utc time),
// properties/usageStart (Utc time), properties/resourceGroup, properties/instanceName, properties/instanceId or
// tags. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does not currently support 'ne', 'or', or
// 'not'. Tag filter is a key value pair string where key and value is separated by a colon (:). skiptoken is
// skiptoken is only used if a previous operation returned a partial result. If a previous response contains a
// nextLink element, the value of the nextLink element will include a skiptoken parameter that specifies a starting
// point to use for subsequent calls. top is may be used to limit the number of results to the most recent N
// usageDetails. apply is oData apply expression to aggregate usageDetails by tags or (tags and
// properties/usageStart)
func (client UsageDetailsByEnrollmentAccountClient) List(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result UsageDetailsListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.UsageDetailsByEnrollmentAccountClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, enrollmentAccountID, expand, filter, skiptoken, top, apply)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.udlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "List", resp, "Failure sending request")
		return
	}

	result.udlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UsageDetailsByEnrollmentAccountClient) ListPreparer(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.Consumption/usageDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsageDetailsByEnrollmentAccountClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsageDetailsByEnrollmentAccountClient) ListResponder(resp *http.Response) (result UsageDetailsListResult, err error) {
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
func (client UsageDetailsByEnrollmentAccountClient) listNextResults(lastResults UsageDetailsListResult) (result UsageDetailsListResult, err error) {
	req, err := lastResults.usageDetailsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsageDetailsByEnrollmentAccountClient) ListComplete(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result UsageDetailsListResultIterator, err error) {
	result.page, err = client.List(ctx, enrollmentAccountID, expand, filter, skiptoken, top, apply)
	return
}

// ListByBillingPeriod lists the usage details based on enrollmentAccountId for a scope by billing period. Usage
// details are available via this API only for May 1, 2014 or later.
//
// enrollmentAccountID is enrollmentAccount ID billingPeriodName is billing Period Name. expand is may be used to
// expand the properties/additionalProperties or properties/meterDetails within a list of usage details. By
// default, these fields are not included when listing usage details. filter is may be used to filter usageDetails
// by properties/usageEnd (Utc time), properties/usageStart (Utc time), properties/resourceGroup,
// properties/instanceName or properties/instanceId. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'.
// It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key and value is
// separated by a colon (:). apply is oData apply expression to aggregate usageDetails by tags or (tags and
// properties/usageStart) for specified billing period skiptoken is skiptoken is only used if a previous operation
// returned a partial result. If a previous response contains a nextLink element, the value of the nextLink element
// will include a skiptoken parameter that specifies a starting point to use for subsequent calls. top is may be
// used to limit the number of results to the most recent N usageDetails.
func (client UsageDetailsByEnrollmentAccountClient) ListByBillingPeriod(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result UsageDetailsListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.UsageDetailsByEnrollmentAccountClient", "ListByBillingPeriod", err.Error())
	}

	result.fn = client.listByBillingPeriodNextResults
	req, err := client.ListByBillingPeriodPreparer(ctx, enrollmentAccountID, billingPeriodName, expand, filter, apply, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "ListByBillingPeriod", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.udlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "ListByBillingPeriod", resp, "Failure sending request")
		return
	}

	result.udlr, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "ListByBillingPeriod", resp, "Failure responding to request")
	}

	return
}

// ListByBillingPeriodPreparer prepares the ListByBillingPeriod request.
func (client UsageDetailsByEnrollmentAccountClient) ListByBillingPeriodPreparer(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingPeriodName":   autorest.Encode("path", billingPeriodName),
		"enrollmentAccountId": autorest.Encode("path", enrollmentAccountID),
	}

	const APIVersion = "2018-03-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
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
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/enrollmentAccounts/{enrollmentAccountId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/usageDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingPeriodSender sends the ListByBillingPeriod request. The method will close the
// http.Response Body if it receives an error.
func (client UsageDetailsByEnrollmentAccountClient) ListByBillingPeriodSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingPeriodResponder handles the response to the ListByBillingPeriod request. The method always
// closes the http.Response Body.
func (client UsageDetailsByEnrollmentAccountClient) ListByBillingPeriodResponder(resp *http.Response) (result UsageDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingPeriodNextResults retrieves the next set of results, if any.
func (client UsageDetailsByEnrollmentAccountClient) listByBillingPeriodNextResults(lastResults UsageDetailsListResult) (result UsageDetailsListResult, err error) {
	req, err := lastResults.usageDetailsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "listByBillingPeriodNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingPeriodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "listByBillingPeriodNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingPeriodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsByEnrollmentAccountClient", "listByBillingPeriodNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingPeriodComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsageDetailsByEnrollmentAccountClient) ListByBillingPeriodComplete(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result UsageDetailsListResultIterator, err error) {
	result.page, err = client.ListByBillingPeriod(ctx, enrollmentAccountID, billingPeriodName, expand, filter, apply, skiptoken, top)
	return
}
