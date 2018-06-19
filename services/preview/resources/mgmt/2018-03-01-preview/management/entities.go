package managementgroups

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

// EntitiesClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type EntitiesClient struct {
	BaseClient
}

// NewEntitiesClient creates an instance of the EntitiesClient client.
func NewEntitiesClient(operationResultID string, skip *int32, top *int32, skiptoken string) EntitiesClient {
	return NewEntitiesClientWithBaseURI(DefaultBaseURI, operationResultID, skip, top, skiptoken)
}

// NewEntitiesClientWithBaseURI creates an instance of the EntitiesClient client.
func NewEntitiesClientWithBaseURI(baseURI string, operationResultID string, skip *int32, top *int32, skiptoken string) EntitiesClient {
	return EntitiesClient{NewWithBaseURI(baseURI, operationResultID, skip, top, skiptoken)}
}

// List list all entities (Management Groups, Subscriptions, etc.) for the authenticated user.
// Parameters:
// selectParameter - this parameter specifies the fields to include in the response. Can include any
// combination of Name,DisplayName,Type,ParentDisplayNameChain,ParentChain, e.g.
// '$select=Name,DisplayName,Type,ParentDisplayNameChain,ParentNameChain'. When specified the $select parameter
// can override select in $skipToken.
// search - the $search parameter is used in conjunction with the $filter parameter to return three different
// outputs depending on the parameter passed in.
// With $search=AllowedParents the API will return the entity info of all groups that the requested entity will
// be able to reparent to as determined by the user's permissions.
// With $search=AllowedChildren the API will return the entity info of all entities that can be added as
// children of the requested entity.
// With $search=ParentAndFirstLevelChildren the API will return the parent and  first level of children that
// the user has either direct access to or indirect access via one of their descendants.
// With $search=ParentOnly the API will return only the group if the user has access to at least one of the
// descendants of the group.
// With $search=ChildrenOnly the API will return only the first level of children of the group entity info
// specified in $filter.  The user must have direct access to the children entities or one of it's descendants
// for it to show up in the results.
// filter - the filter parameter allows you to filter on the the name or display name fields. You can check for
// equality on the name field (e.g. name eq '{entityName}')  and you can check for substrings on either the
// name or display name fields(e.g. contains(name, '{substringToSearch}'), contains(displayName,
// '{substringToSearch')). Note that the '{entityName}' and '{substringToSearch}' fields are checked case
// insensitively.
// view - the view parameter allows clients to filter the type of data that is returned by the getEntities
// call.
// groupName - a filter which allows the get entities call to focus on a particular group (i.e. "$filter=name
// eq 'groupName'")
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client EntitiesClient) List(ctx context.Context, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (result EntitiesListFuture, err error) {
	req, err := client.ListPreparer(ctx, selectParameter, search, filter, view, groupName, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", nil, "Failure preparing request")
		return
	}

	result, err = client.ListSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", result.Response(), "Failure sending request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EntitiesClient) ListPreparer(ctx context.Context, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (*http.Request, error) {
	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(client.Skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", client.Skiptoken)
	}
	if client.Skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *client.Skip)
	}
	if client.Top != nil {
		queryParameters["$top"] = autorest.Encode("query", *client.Top)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(string(search)) > 0 {
		queryParameters["$search"] = autorest.Encode("query", search)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(string(view)) > 0 {
		queryParameters["$view"] = autorest.Encode("query", view)
	}
	if len(groupName) > 0 {
		queryParameters["groupName"] = autorest.Encode("query", groupName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/getEntities"),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EntitiesClient) ListSender(req *http.Request) (future EntitiesListFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EntitiesClient) ListResponder(resp *http.Response) (result EntityListResultPage, err error) {
	result.elr, err = client.listResponder(resp)
	result.fn = client.listNextResults
	return
}

func (client EntitiesClient) listResponder(resp *http.Response) (result EntityListResult, err error) {
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
func (client EntitiesClient) listNextResults(lastResults EntityListResult) (result EntityListResult, err error) {
	req, err := lastResults.entityListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", resp, "Failure sending next results request")
	}
	return client.listResponder(resp)
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EntitiesClient) ListComplete(ctx context.Context, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (result EntitiesListAllFuture, err error) {
	var future EntitiesListFuture
	future, err = client.List(ctx, selectParameter, search, filter, view, groupName, cacheControl)
	result.Future = future.Future
	return
}
