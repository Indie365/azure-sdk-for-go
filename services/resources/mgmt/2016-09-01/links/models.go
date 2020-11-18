package links

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"

// Operation microsoft.Resources operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Resources
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list Microsoft.Resources operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Microsoft.Resources operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// ResourceLink the resource link.
type ResourceLink struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The fully qualified ID of the resource link.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource link.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource link object.
	Type interface{} `json:"type,omitempty"`
	// Properties - Properties for resource link.
	Properties *ResourceLinkProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceLink.
func (rl ResourceLink) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rl.Properties != nil {
		objectMap["properties"] = rl.Properties
	}
	return json.Marshal(objectMap)
}

// ResourceLinkFilter resource link filter.
type ResourceLinkFilter struct {
	// TargetID - The ID of the target resource.
	TargetID *string `json:"targetId,omitempty"`
}

// ResourceLinkProperties the resource link properties.
type ResourceLinkProperties struct {
	// SourceID - READ-ONLY; The fully qualified ID of the source resource in the link.
	SourceID *string `json:"sourceId,omitempty"`
	// TargetID - The fully qualified ID of the target resource in the link.
	TargetID *string `json:"targetId,omitempty"`
	// Notes - Notes about the resource link.
	Notes *string `json:"notes,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceLinkProperties.
func (rlp ResourceLinkProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rlp.TargetID != nil {
		objectMap["targetId"] = rlp.TargetID
	}
	if rlp.Notes != nil {
		objectMap["notes"] = rlp.Notes
	}
	return json.Marshal(objectMap)
}

// ResourceLinkResult list of resource links.
type ResourceLinkResult struct {
	autorest.Response `json:"-"`
	// Value - An array of resource links.
	Value *[]ResourceLink `json:"value,omitempty"`
	// NextLink - READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceLinkResult.
func (rlr ResourceLinkResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rlr.Value != nil {
		objectMap["value"] = rlr.Value
	}
	return json.Marshal(objectMap)
}

// ResourceLinkResultIterator provides access to a complete listing of ResourceLink values.
type ResourceLinkResultIterator struct {
	i    int
	page ResourceLinkResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceLinkResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceLinkResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ResourceLinkResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceLinkResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceLinkResultIterator) Response() ResourceLinkResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceLinkResultIterator) Value() ResourceLink {
	if !iter.page.NotDone() {
		return ResourceLink{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceLinkResultIterator type.
func NewResourceLinkResultIterator(page ResourceLinkResultPage) ResourceLinkResultIterator {
	return ResourceLinkResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rlr ResourceLinkResult) IsEmpty() bool {
	return rlr.Value == nil || len(*rlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rlr ResourceLinkResult) hasNextLink() bool {
	return rlr.NextLink != nil && len(*rlr.NextLink) != 0
}

// resourceLinkResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rlr ResourceLinkResult) resourceLinkResultPreparer(ctx context.Context) (*http.Request, error) {
	if !rlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rlr.NextLink)))
}

// ResourceLinkResultPage contains a page of ResourceLink values.
type ResourceLinkResultPage struct {
	fn  func(context.Context, ResourceLinkResult) (ResourceLinkResult, error)
	rlr ResourceLinkResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceLinkResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceLinkResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rlr)
		if err != nil {
			return err
		}
		page.rlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceLinkResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceLinkResultPage) NotDone() bool {
	return !page.rlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceLinkResultPage) Response() ResourceLinkResult {
	return page.rlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceLinkResultPage) Values() []ResourceLink {
	if page.rlr.IsEmpty() {
		return nil
	}
	return *page.rlr.Value
}

// Creates a new instance of the ResourceLinkResultPage type.
