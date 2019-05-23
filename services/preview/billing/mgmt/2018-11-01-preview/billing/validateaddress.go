package billing

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

// ValidateAddressClient is the billing client provides access to billing resources for Azure subscriptions.
type ValidateAddressClient struct {
	BaseClient
}

// NewValidateAddressClient creates an instance of the ValidateAddressClient client.
func NewValidateAddressClient(subscriptionID string) ValidateAddressClient {
	return NewValidateAddressClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewValidateAddressClientWithBaseURI creates an instance of the ValidateAddressClient client.
func NewValidateAddressClientWithBaseURI(baseURI string, subscriptionID string) ValidateAddressClient {
	return ValidateAddressClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Validate validates the address.
func (client ValidateAddressClient) Validate(ctx context.Context, address Address) (result ValidateAddressResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ValidateAddressClient.Validate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidatePreparer(ctx, address)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ValidateAddressClient", "Validate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ValidateAddressClient", "Validate", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ValidateAddressClient", "Validate", resp, "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client ValidateAddressClient) ValidatePreparer(ctx context.Context, address Address) (*http.Request, error) {
	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Billing/validateAddress"),
		autorest.WithJSON(address),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client ValidateAddressClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client ValidateAddressClient) ValidateResponder(resp *http.Response) (result ValidateAddressResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
