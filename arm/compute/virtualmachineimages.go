package compute

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// VirtualMachineImages Client
type VirtualMachineImagesClient struct {
	ComputeManagementClient
}

func NewVirtualMachineImagesClient(subscriptionId string) VirtualMachineImagesClient {
	return NewVirtualMachineImagesClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewVirtualMachineImagesClientWithBaseUri(baseUri string, subscriptionId string) VirtualMachineImagesClient {
	return VirtualMachineImagesClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// Get gets a virtual machine image.
//
func (client VirtualMachineImagesClient) Get(location string, publisherName string, offer string, skus string, version string) (result VirtualMachineImage, ae autorest.Error) {
	req, err := client.NewGetRequest(location, publisherName, offer, skus, version)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client VirtualMachineImagesClient) NewGetRequest(location string, publisherName string, offer string, skus string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"offer":          url.QueryEscape(offer),
		"publisherName":  url.QueryEscape(publisherName),
		"skus":           url.QueryEscape(skus),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
		"version":        url.QueryEscape(version),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client VirtualMachineImagesClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"))
}

// ListOffers gets a list of virtual machine image offers.
//
func (client VirtualMachineImagesClient) ListOffers(location string, publisherName string) (result VirtualMachineImageResourceList, ae autorest.Error) {
	req, err := client.NewListOffersRequest(location, publisherName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result.Value))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListOffers request.
func (client VirtualMachineImagesClient) NewListOffersRequest(location string, publisherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"publisherName":  url.QueryEscape(publisherName),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListOffersRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListOffers request.
func (client VirtualMachineImagesClient) ListOffersRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers"))
}

// ListPublishers gets a list of virtual machine image publishers.
//
func (client VirtualMachineImagesClient) ListPublishers(location string) (result VirtualMachineImageResourceList, ae autorest.Error) {
	req, err := client.NewListPublishersRequest(location)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result.Value))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListPublishers request.
func (client VirtualMachineImagesClient) NewListPublishersRequest(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListPublishersRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListPublishers request.
func (client VirtualMachineImagesClient) ListPublishersRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers"))
}

// ListSkus gets a list of virtual machine image skus.
//
func (client VirtualMachineImagesClient) ListSkus(location string, publisherName string, offer string) (result VirtualMachineImageResourceList, ae autorest.Error) {
	req, err := client.NewListSkusRequest(location, publisherName, offer)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result.Value))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the ListSkus request.
func (client VirtualMachineImagesClient) NewListSkusRequest(location string, publisherName string, offer string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"offer":          url.QueryEscape(offer),
		"publisherName":  url.QueryEscape(publisherName),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListSkusRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the ListSkus request.
func (client VirtualMachineImagesClient) ListSkusRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"))
}

// List gets a list of virtual machine images.
//
// filter is the filter to apply on the operation.
func (client VirtualMachineImagesClient) List(location string, publisherName string, offer string, skus string, filter string, top int, orderby string) (result VirtualMachineImageResourceList, ae autorest.Error) {
	req, err := client.NewListRequest(location, publisherName, offer, skus, filter, top, orderby)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result.Value))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client VirtualMachineImagesClient) NewListRequest(location string, publisherName string, offer string, skus string, filter string, top int, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"offer":          url.QueryEscape(offer),
		"publisherName":  url.QueryEscape(publisherName),
		"skus":           url.QueryEscape(skus),
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"$orderby":    orderby,
		"$top":        top,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client VirtualMachineImagesClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"))
}
