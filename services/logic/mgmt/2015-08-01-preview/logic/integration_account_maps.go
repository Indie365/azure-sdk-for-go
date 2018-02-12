package logic

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
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// IntegrationAccountMapsClient is the REST API for Azure Logic Apps.
type IntegrationAccountMapsClient struct {
	ManagementClient
}

// NewIntegrationAccountMapsClient creates an instance of the IntegrationAccountMapsClient client.
func NewIntegrationAccountMapsClient(p pipeline.Pipeline) IntegrationAccountMapsClient {
	return IntegrationAccountMapsClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates an integration account map.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. mapName is the
// integration account map name. mapParameter is the integration account map.
func (client IntegrationAccountMapsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, mapParameter IntegrationAccountMap) (*IntegrationAccountMap, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, integrationAccountName, mapName, mapParameter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*IntegrationAccountMap), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountMapsClient) createOrUpdatePreparer(resourceGroupName string, integrationAccountName string, mapName string, mapParameter IntegrationAccountMap) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(mapParameter)
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
func (client IntegrationAccountMapsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &IntegrationAccountMap{rawResponse: resp.Response()}
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

// Delete deletes an integration account map.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. mapName is the
// integration account map name.
func (client IntegrationAccountMapsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, integrationAccountName, mapName)
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
func (client IntegrationAccountMapsClient) deletePreparer(resourceGroupName string, integrationAccountName string, mapName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
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
func (client IntegrationAccountMapsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets an integration account map.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. mapName is the
// integration account map name.
func (client IntegrationAccountMapsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string) (*IntegrationAccountMap, error) {
	req, err := client.getPreparer(resourceGroupName, integrationAccountName, mapName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*IntegrationAccountMap), err
}

// getPreparer prepares the Get request.
func (client IntegrationAccountMapsClient) getPreparer(resourceGroupName string, integrationAccountName string, mapName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
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
func (client IntegrationAccountMapsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &IntegrationAccountMap{rawResponse: resp.Response()}
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

// List gets a list of integration account maps.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. top is the
// number of items to be included in the result. filter is the filter to apply on the operation.
func (client IntegrationAccountMapsClient) List(ctx context.Context, resourceGroupName string, integrationAccountName string, top *int32, filter *string) (*IntegrationAccountMapListResult, error) {
	req, err := client.listPreparer(resourceGroupName, integrationAccountName, top, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*IntegrationAccountMapListResult), err
}

// listPreparer prepares the List request.
func (client IntegrationAccountMapsClient) listPreparer(resourceGroupName string, integrationAccountName string, top *int32, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	if top != nil {
		params.Set("$top", fmt.Sprintf("%v", *top))
	}
	if filter != nil {
		params.Set("$filter", *filter)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client IntegrationAccountMapsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &IntegrationAccountMapListResult{rawResponse: resp.Response()}
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
