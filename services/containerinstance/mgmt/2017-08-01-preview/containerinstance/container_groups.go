package containerinstance

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

// ContainerGroupsClient is the client for the ContainerGroups methods of the Containerinstance service.
type ContainerGroupsClient struct {
	ManagementClient
}

// NewContainerGroupsClient creates an instance of the ContainerGroupsClient client.
func NewContainerGroupsClient(p pipeline.Pipeline) ContainerGroupsClient {
	return ContainerGroupsClient{NewManagementClient(p)}
}

// CreateOrUpdate create or update container groups with specified configurations.
//
// resourceGroupName is the name of the resource group to contain the container group to be created or updated.
// containerGroupName is the name of the container group to be created or updated. containerGroup is the properties of
// the container group to be created or updated.
func (client ContainerGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup ContainerGroup) (*ContainerGroup, error) {
	if err := validate([]validation{
		{targetValue: containerGroup,
			constraints: []constraint{{target: "containerGroup.ContainerGroupProperties", name: null, rule: false,
				chain: []constraint{{target: "containerGroup.ContainerGroupProperties.IPAddress", name: null, rule: false,
					chain: []constraint{{target: "containerGroup.ContainerGroupProperties.IPAddress.Ports", name: null, rule: true, chain: nil},
						{target: "containerGroup.ContainerGroupProperties.IPAddress.Type", name: null, rule: true, chain: nil},
					}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, containerGroupName, containerGroup)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerGroup), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContainerGroupsClient) createOrUpdatePreparer(resourceGroupName string, containerGroupName string, containerGroup ContainerGroup) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(containerGroup)
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
func (client ContainerGroupsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &ContainerGroup{rawResponse: resp.Response()}
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

// Delete delete the specified container group in the specified subscription and resource group. The operation does not
// delete other resources provided by the user, such as volumes.
//
// resourceGroupName is the name of the resource group that contains the container group. containerGroupName is the
// name of the container group to be deleted.
func (client ContainerGroupsClient) Delete(ctx context.Context, resourceGroupName string, containerGroupName string) (*ContainerGroup, error) {
	req, err := client.deletePreparer(resourceGroupName, containerGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerGroup), err
}

// deletePreparer prepares the Delete request.
func (client ContainerGroupsClient) deletePreparer(resourceGroupName string, containerGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
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
func (client ContainerGroupsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	result := &ContainerGroup{rawResponse: resp.Response()}
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

// Get gets the properties of the specified container group in the specified subscription and resource group. The
// operation returns the properties of each container group including containers, image registry credentials, restart
// policy, IP address type, OS type, state, and volumes.
//
// resourceGroupName is the name of the resource group that contains the container group. containerGroupName is the
// name of the container group.
func (client ContainerGroupsClient) Get(ctx context.Context, resourceGroupName string, containerGroupName string) (*ContainerGroup, error) {
	req, err := client.getPreparer(resourceGroupName, containerGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerGroup), err
}

// getPreparer prepares the Get request.
func (client ContainerGroupsClient) getPreparer(resourceGroupName string, containerGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
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
func (client ContainerGroupsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ContainerGroup{rawResponse: resp.Response()}
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

// List get a list of container groups in the specified subscription. This operation returns properties of each
// container group including containers, image registry credentials, restart policy, IP address type, OS type, state,
// and volumes.
func (client ContainerGroupsClient) List(ctx context.Context) (*ContainerGroupListResult, error) {
	req, err := client.listPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerGroupListResult), err
}

// listPreparer prepares the List request.
func (client ContainerGroupsClient) listPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups"
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
func (client ContainerGroupsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ContainerGroupListResult{rawResponse: resp.Response()}
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

// ListByResourceGroup get a list of container groups in a specified subscription and resource group. This operation
// returns properties of each container group including containers, image registry credentials, restart policy, IP
// address type, OS type, state, and volumes.
//
// resourceGroupName is the name of the resource group that contains the container group.
func (client ContainerGroupsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (*ContainerGroupListResult, error) {
	req, err := client.listByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ContainerGroupListResult), err
}

// listByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ContainerGroupsClient) listByResourceGroupPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByResourceGroupResponder handles the response to the ListByResourceGroup request.
func (client ContainerGroupsClient) listByResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ContainerGroupListResult{rawResponse: resp.Response()}
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
