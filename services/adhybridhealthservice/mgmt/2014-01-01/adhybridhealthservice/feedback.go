package adhybridhealthservice

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

// FeedbackClient is the REST APIs for Azure Active Drectory Connect Health
type FeedbackClient struct {
	BaseClient
}

// NewFeedbackClient creates an instance of the FeedbackClient client.
func NewFeedbackClient() FeedbackClient {
	return NewFeedbackClientWithBaseURI(DefaultBaseURI)
}

// NewFeedbackClientWithBaseURI creates an instance of the FeedbackClient client.
func NewFeedbackClientWithBaseURI(baseURI string) FeedbackClient {
	return FeedbackClient{NewWithBaseURI(baseURI)}
}

// AddAlertFeedback adds an alert feedback submitted by customer.
// Parameters:
// serviceName - the name of the service.
// alertFeedback - the alert feedback.
func (client FeedbackClient) AddAlertFeedback(ctx context.Context, serviceName string, alertFeedback AlertFeedback) (result AlertFeedback, err error) {
	req, err := client.AddAlertFeedbackPreparer(ctx, serviceName, alertFeedback)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.FeedbackClient", "AddAlertFeedback", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddAlertFeedbackSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.FeedbackClient", "AddAlertFeedback", resp, "Failure sending request")
		return
	}

	result, err = client.AddAlertFeedbackResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.FeedbackClient", "AddAlertFeedback", resp, "Failure responding to request")
	}

	return
}

// AddAlertFeedbackPreparer prepares the AddAlertFeedback request.
func (client FeedbackClient) AddAlertFeedbackPreparer(ctx context.Context, serviceName string, alertFeedback AlertFeedback) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/feedbacktype/alerts/feedback", pathParameters),
		autorest.WithJSON(alertFeedback),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddAlertFeedbackSender sends the AddAlertFeedback request. The method will close the
// http.Response Body if it receives an error.
func (client FeedbackClient) AddAlertFeedbackSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddAlertFeedbackResponder handles the response to the AddAlertFeedback request. The method always
// closes the http.Response Body.
func (client FeedbackClient) AddAlertFeedbackResponder(resp *http.Response) (result AlertFeedback, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
