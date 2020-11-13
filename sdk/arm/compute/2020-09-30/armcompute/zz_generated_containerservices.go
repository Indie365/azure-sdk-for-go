// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ContainerServicesOperations contains the methods for the ContainerServices group.
type ContainerServicesOperations interface {
// BeginCreateOrUpdate - Creates or updates a container service with the specified configuration of orchestrator, masters, and agents.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesCreateOrUpdateOptions) (*ContainerServicePollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ContainerServicePoller, error)
// BeginDelete - Deletes the specified container service in the specified subscription and resource group. The operation does not delete other resources
// created as part of creating a container service, including
// storage accounts, VMs, and availability sets. All the other resources created with the container service are part of the same resource group and can
// be deleted individually.
	BeginDelete(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
// Get - Gets the properties of the specified container service in the specified subscription and resource group. The operation returns the properties including
// state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
	Get(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesGetOptions) (*ContainerServiceResponse, error)
// List - Gets a list of container services in the specified subscription. The operation returns properties of each container service including state, orchestrator,
// number of masters and agents, and FQDNs of
// masters and agents.
	List(options *ContainerServicesListOptions) (ContainerServiceListResultPager)
// ListByResourceGroup - Gets a list of container services in the specified subscription and resource group. The operation returns properties of each container
// service including state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
	ListByResourceGroup(resourceGroupName string, options *ContainerServicesListByResourceGroupOptions) (ContainerServiceListResultPager)
}

// ContainerServicesClient implements the ContainerServicesOperations interface.
// Don't use this type directly, use NewContainerServicesClient() instead.
type ContainerServicesClient struct {
	con *armcore.Connection
	subscriptionID string
}

// NewContainerServicesClient creates a new instance of ContainerServicesClient with the specified values.
func NewContainerServicesClient(con *armcore.Connection, subscriptionID string) ContainerServicesOperations {
	return &ContainerServicesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *ContainerServicesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *ContainerServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesCreateOrUpdateOptions) (*ContainerServicePollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, containerServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ContainerServicePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ContainerServicesClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &containerServicePoller{
		pt: pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ContainerServiceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ContainerServicesClient) ResumeCreateOrUpdate(token string) (ContainerServicePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ContainerServicesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &containerServicePoller{
		pipeline: client.con.Pipeline(),
		pt: pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a container service with the specified configuration of orchestrator, masters, and agents.
func (client *ContainerServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, containerServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	 return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContainerServicesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService, options *ContainerServicesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2017-01-31")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContainerServicesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ContainerServiceResponse, error) {
	result := ContainerServiceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ContainerService)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ContainerServicesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
    }
    if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

func (client *ContainerServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ContainerServicesClient.Delete", "", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt: pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ContainerServicesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ContainerServicesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt: pt,
	}, nil
}

// Delete - Deletes the specified container service in the specified subscription and resource group. The operation does not delete other resources created
// as part of creating a container service, including
// storage accounts, VMs, and availability sets. All the other resources created with the container service are part of the same resource group and can
// be deleted individually.
func (client *ContainerServicesClient) Delete(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	 return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *ContainerServicesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2017-01-31")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *ContainerServicesClient) DeleteHandleError(resp *azcore.Response) error {
body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
    }
    if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

// Get - Gets the properties of the specified container service in the specified subscription and resource group. The operation returns the properties including
// state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
func (client *ContainerServicesClient) Get(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesGetOptions) (*ContainerServiceResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, containerServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *ContainerServicesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, containerServiceName string, options *ContainerServicesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{containerServiceName}", url.PathEscape(containerServiceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2017-01-31")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *ContainerServicesClient) GetHandleResponse(resp *azcore.Response) (*ContainerServiceResponse, error) {
	result := ContainerServiceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ContainerService)
}

// GetHandleError handles the Get error response.
func (client *ContainerServicesClient) GetHandleError(resp *azcore.Response) error {
body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
    }
    if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

// List - Gets a list of container services in the specified subscription. The operation returns properties of each container service including state, orchestrator,
// number of masters and agents, and FQDNs of
// masters and agents.
func (client *ContainerServicesClient) List(options *ContainerServicesListOptions) (ContainerServiceListResultPager) {
	return &containerServiceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ContainerServiceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ContainerServiceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *ContainerServicesClient) ListCreateRequest(ctx context.Context, options *ContainerServicesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2017-01-31")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *ContainerServicesClient) ListHandleResponse(resp *azcore.Response) (*ContainerServiceListResultResponse, error) {
	result := ContainerServiceListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ContainerServiceListResult)
}

// ListHandleError handles the List error response.
func (client *ContainerServicesClient) ListHandleError(resp *azcore.Response) error {
body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
    }
    if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

// ListByResourceGroup - Gets a list of container services in the specified subscription and resource group. The operation returns properties of each container
// service including state, orchestrator, number of masters and
// agents, and FQDNs of masters and agents.
func (client *ContainerServicesClient) ListByResourceGroup(resourceGroupName string, options *ContainerServicesListByResourceGroupOptions) (ContainerServiceListResultPager) {
	return &containerServiceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ContainerServiceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ContainerServiceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerServicesClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerServicesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2017-01-31")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerServicesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*ContainerServiceListResultResponse, error) {
	result := ContainerServiceListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ContainerServiceListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ContainerServicesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
    }
    if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

