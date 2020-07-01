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

// LocationClient is the client for the Location methods of the Batch service.
type LocationClient struct {
	BaseClient
}

// NewLocationClient creates an instance of the LocationClient client.
func NewLocationClient(subscriptionID string) LocationClient {
	return NewLocationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationClientWithBaseURI creates an instance of the LocationClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationClientWithBaseURI(baseURI string, subscriptionID string) LocationClient {
	return LocationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks whether the Batch account name is available in the specified region.
// Parameters:
// locationName - the desired region for the name check.
// parameters - properties needed to check the availability of a name.
func (client LocationClient) CheckNameAvailability(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters) (result CheckNameAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.LocationClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, locationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.LocationClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.LocationClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.LocationClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client LocationClient) CheckNameAvailabilityPreparer(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/checkNameAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client LocationClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client LocationClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetQuotas gets the Batch service quotas for the specified subscription at the given location.
// Parameters:
// locationName - the region for which to retrieve Batch service quotas.
func (client LocationClient) GetQuotas(ctx context.Context, locationName string) (result LocationQuota, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationClient.GetQuotas")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetQuotasPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.LocationClient", "GetQuotas", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetQuotasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.LocationClient", "GetQuotas", resp, "Failure sending request")
		return
	}

	result, err = client.GetQuotasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.LocationClient", "GetQuotas", resp, "Failure responding to request")
	}

	return
}

// GetQuotasPreparer prepares the GetQuotas request.
func (client LocationClient) GetQuotasPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/quotas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetQuotasSender sends the GetQuotas request. The method will close the
// http.Response Body if it receives an error.
func (client LocationClient) GetQuotasSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetQuotasResponder handles the response to the GetQuotas request. The method always
// closes the http.Response Body.
func (client LocationClient) GetQuotasResponder(resp *http.Response) (result LocationQuota, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
