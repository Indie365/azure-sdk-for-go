package network

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
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// ExpressRouteCircuitsClient is the network Client
type ExpressRouteCircuitsClient struct {
	ManagementClient
}

// NewExpressRouteCircuitsClient creates an instance of the ExpressRouteCircuitsClient client.
func NewExpressRouteCircuitsClient(p pipeline.Pipeline) ExpressRouteCircuitsClient {
	return ExpressRouteCircuitsClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates an express route circuit. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the circuit. parameters is
// parameters supplied to the create or update express route circuit operation.
func (client ExpressRouteCircuitsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*ExpressRouteCircuit, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, circuitName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuit), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitsClient) createOrUpdatePreparer(resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client ExpressRouteCircuitsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuit{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes the specified express route circuit. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
func (client ExpressRouteCircuitsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, circuitName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client ExpressRouteCircuitsClient) deletePreparer(resourceGroupName string, circuitName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ExpressRouteCircuitsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// Get gets information about the specified express route circuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of express route circuit.
func (client ExpressRouteCircuitsClient) Get(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuit, error) {
	req, err := client.getPreparer(resourceGroupName, circuitName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuit), err
}

// getPreparer prepares the Get request.
func (client ExpressRouteCircuitsClient) getPreparer(resourceGroupName string, circuitName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client ExpressRouteCircuitsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuit{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetPeeringStats gets all stats from an express route circuit in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering.
func (client ExpressRouteCircuitsClient) GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*ExpressRouteCircuitStats, error) {
	req, err := client.getPeeringStatsPreparer(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPeeringStatsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitStats), err
}

// getPeeringStatsPreparer prepares the GetPeeringStats request.
func (client ExpressRouteCircuitsClient) getPeeringStatsPreparer(resourceGroupName string, circuitName string, peeringName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getPeeringStatsResponder handles the response to the GetPeeringStats request.
func (client ExpressRouteCircuitsClient) getPeeringStatsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitStats{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetStats gets all the stats from an express route circuit in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
func (client ExpressRouteCircuitsClient) GetStats(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuitStats, error) {
	req, err := client.getStatsPreparer(resourceGroupName, circuitName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getStatsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitStats), err
}

// getStatsPreparer prepares the GetStats request.
func (client ExpressRouteCircuitsClient) getStatsPreparer(resourceGroupName string, circuitName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getStatsResponder handles the response to the GetStats request.
func (client ExpressRouteCircuitsClient) getStatsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitStats{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List gets all the express route circuits in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ExpressRouteCircuitsClient) List(ctx context.Context, resourceGroupName string) (*ExpressRouteCircuitListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitListResult), err
}

// listPreparer prepares the List request.
func (client ExpressRouteCircuitsClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client ExpressRouteCircuitsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListAll gets all the express route circuits in a subscription.
func (client ExpressRouteCircuitsClient) ListAll(ctx context.Context) (*ExpressRouteCircuitListResult, error) {
	req, err := client.listAllPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAllResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitListResult), err
}

// listAllPreparer prepares the ListAll request.
func (client ExpressRouteCircuitsClient) listAllPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listAllResponder handles the response to the ListAll request.
func (client ExpressRouteCircuitsClient) listAllResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListArpTable gets the currently advertised ARP table associated with the express route circuit in a resource group.
// This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering. devicePath is the path of the device.
func (client ExpressRouteCircuitsClient) ListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsArpTableListResult, error) {
	req, err := client.listArpTablePreparer(resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listArpTableResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitsArpTableListResult), err
}

// listArpTablePreparer prepares the ListArpTable request.
func (client ExpressRouteCircuitsClient) listArpTablePreparer(resourceGroupName string, circuitName string, peeringName string, devicePath string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/arpTables/{devicePath}"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listArpTableResponder handles the response to the ListArpTable request.
func (client ExpressRouteCircuitsClient) listArpTableResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsArpTableListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListRoutesTable gets the currently advertised routes table associated with the express route circuit in a resource
// group. This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering. devicePath is the path of the device.
func (client ExpressRouteCircuitsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableListResult, error) {
	req, err := client.listRoutesTablePreparer(resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listRoutesTableResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitsRoutesTableListResult), err
}

// listRoutesTablePreparer prepares the ListRoutesTable request.
func (client ExpressRouteCircuitsClient) listRoutesTablePreparer(resourceGroupName string, circuitName string, peeringName string, devicePath string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTables/{devicePath}"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listRoutesTableResponder handles the response to the ListRoutesTable request.
func (client ExpressRouteCircuitsClient) listRoutesTableResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListRoutesTableSummary gets the currently advertised routes table summary associated with the express route circuit
// in a resource group. This method may poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering. devicePath is the path of the device.
func (client ExpressRouteCircuitsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (*ExpressRouteCircuitsRoutesTableSummaryListResult, error) {
	req, err := client.listRoutesTableSummaryPreparer(resourceGroupName, circuitName, peeringName, devicePath)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listRoutesTableSummaryResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitsRoutesTableSummaryListResult), err
}

// listRoutesTableSummaryPreparer prepares the ListRoutesTableSummary request.
func (client ExpressRouteCircuitsClient) listRoutesTableSummaryPreparer(resourceGroupName string, circuitName string, peeringName string, devicePath string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listRoutesTableSummaryResponder handles the response to the ListRoutesTableSummary request.
func (client ExpressRouteCircuitsClient) listRoutesTableSummaryResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableSummaryListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
