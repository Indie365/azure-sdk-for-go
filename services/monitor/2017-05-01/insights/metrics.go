package insights

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// MetricsClient is the monitor Client
type MetricsClient struct {
	ManagementClient
}

// NewMetricsClient creates an instance of the MetricsClient client.
func NewMetricsClient(subscriptionID string) MetricsClient {
	return NewMetricsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMetricsClientWithBaseURI creates an instance of the MetricsClient client.
func NewMetricsClientWithBaseURI(baseURI string, subscriptionID string) MetricsClient {
	return MetricsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List **Lists the metric values for a resource**.
//
// resourceURI is the identifier of the resource. timespan is the timespan of the query. It is a string with the
// following format 'startDateTime_ISO/endDateTime_ISO'. interval is the interval (i.e. timegrain) of the query. metric
// is the name of the metric to retrieve. aggregation is the list of aggregation types (comma separated) to retrieve.
// top is the maximum number of records to retrieve.
// Valid only if $filter is specified.
// Defaults to 10. orderby is the aggregation to use for sorting results and the direction of the sort.
// Only one order can be specified.
// Examples: sum asc. filter is the **$filter** is used to reduce the set of metric data
// returned.<br>Example:<br>Metric contains metadata A, B and C.<br>- Return all time series of C where A = a1 and B =
// b1 or b2<br>**$filter=A eq ‘a1’ and B eq ‘b1’ or B eq ‘b2’ and C eq ‘*’**<br>- Invalid variant:<br>**$filter=A eq
// ‘a1’ and B eq ‘b1’ and C eq ‘*’ or B = ‘b2’**<br>This is invalid because the logical or operator cannot separate two
// different metadata names.<br>- Return all time series where A = a1, B = b1 and C = c1:<br>**$filter=A eq ‘a1’ and B
// eq ‘b1’ and C eq ‘c1’**<br>- Return all time series where A = a1<br>**$filter=A eq ‘a1’ and B eq ‘*’ and C eq ‘*’**.
// resultType is reduces the set of data collected. The syntax allowed depends on the operation. See the operation's
// description for details.
func (client MetricsClient) List(resourceURI string, timespan string, interval *string, metric string, aggregation string, top *float64, orderby string, filter string, resultType ResultType) (result Response, err error) {
	req, err := client.ListPreparer(resourceURI, timespan, interval, metric, aggregation, top, orderby, filter, resultType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MetricsClient) ListPreparer(resourceURI string, timespan string, interval *string, metric string, aggregation string, top *float64, orderby string, filter string, resultType ResultType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2017-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(timespan) > 0 {
		queryParameters["timespan"] = autorest.Encode("query", timespan)
	}
	if interval != nil {
		queryParameters["interval"] = autorest.Encode("query", *interval)
	}
	if len(metric) > 0 {
		queryParameters["metric"] = autorest.Encode("query", metric)
	}
	if len(aggregation) > 0 {
		queryParameters["aggregation"] = autorest.Encode("query", aggregation)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(string(resultType)) > 0 {
		queryParameters["resultType"] = autorest.Encode("query", resultType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MetricsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MetricsClient) ListResponder(resp *http.Response) (result Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
