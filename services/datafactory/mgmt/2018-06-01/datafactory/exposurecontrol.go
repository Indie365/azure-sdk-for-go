package datafactory

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

// ExposureControlClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that
// interact with Azure Data Factory V2 services.
type ExposureControlClient struct {
	BaseClient
}

// NewExposureControlClient creates an instance of the ExposureControlClient client.
func NewExposureControlClient(subscriptionID string) ExposureControlClient {
	return NewExposureControlClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExposureControlClientWithBaseURI creates an instance of the ExposureControlClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExposureControlClientWithBaseURI(baseURI string, subscriptionID string) ExposureControlClient {
	return ExposureControlClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetFeatureValue get exposure control feature for specific location.
// Parameters:
// locationID - the location identifier.
// exposureControlRequest - the exposure control request.
func (client ExposureControlClient) GetFeatureValue(ctx context.Context, locationID string, exposureControlRequest ExposureControlRequest) (result ExposureControlResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExposureControlClient.GetFeatureValue")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetFeatureValuePreparer(ctx, locationID, exposureControlRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ExposureControlClient", "GetFeatureValue", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFeatureValueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ExposureControlClient", "GetFeatureValue", resp, "Failure sending request")
		return
	}

	result, err = client.GetFeatureValueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ExposureControlClient", "GetFeatureValue", resp, "Failure responding to request")
	}

	return
}

// GetFeatureValuePreparer prepares the GetFeatureValue request.
func (client ExposureControlClient) GetFeatureValuePreparer(ctx context.Context, locationID string, exposureControlRequest ExposureControlRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationId":     autorest.Encode("path", locationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/locations/{locationId}/getFeatureValue", pathParameters),
		autorest.WithJSON(exposureControlRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFeatureValueSender sends the GetFeatureValue request. The method will close the
// http.Response Body if it receives an error.
func (client ExposureControlClient) GetFeatureValueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetFeatureValueResponder handles the response to the GetFeatureValue request. The method always
// closes the http.Response Body.
func (client ExposureControlClient) GetFeatureValueResponder(resp *http.Response) (result ExposureControlResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFeatureValueByFactory get exposure control feature for specific factory.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// exposureControlRequest - the exposure control request.
func (client ExposureControlClient) GetFeatureValueByFactory(ctx context.Context, resourceGroupName string, factoryName string, exposureControlRequest ExposureControlRequest) (result ExposureControlResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExposureControlClient.GetFeatureValueByFactory")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.ExposureControlClient", "GetFeatureValueByFactory", err.Error())
	}

	req, err := client.GetFeatureValueByFactoryPreparer(ctx, resourceGroupName, factoryName, exposureControlRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ExposureControlClient", "GetFeatureValueByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFeatureValueByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.ExposureControlClient", "GetFeatureValueByFactory", resp, "Failure sending request")
		return
	}

	result, err = client.GetFeatureValueByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.ExposureControlClient", "GetFeatureValueByFactory", resp, "Failure responding to request")
	}

	return
}

// GetFeatureValueByFactoryPreparer prepares the GetFeatureValueByFactory request.
func (client ExposureControlClient) GetFeatureValueByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string, exposureControlRequest ExposureControlRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/getFeatureValue", pathParameters),
		autorest.WithJSON(exposureControlRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFeatureValueByFactorySender sends the GetFeatureValueByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client ExposureControlClient) GetFeatureValueByFactorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetFeatureValueByFactoryResponder handles the response to the GetFeatureValueByFactory request. The method always
// closes the http.Response Body.
func (client ExposureControlClient) GetFeatureValueByFactoryResponder(resp *http.Response) (result ExposureControlResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
