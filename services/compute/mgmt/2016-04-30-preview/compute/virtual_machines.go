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

// VirtualMachinesClient is the compute Client
type VirtualMachinesClient struct {
	ManagementClient
}

// NewVirtualMachinesClient creates an instance of the VirtualMachinesClient client.
func NewVirtualMachinesClient(p pipeline.Pipeline) VirtualMachinesClient {
	return VirtualMachinesClient{NewManagementClient(p)}
}

// Capture captures the VM by copying virtual hard disks of the VM and outputs a template that can be used to create
// similar VMs. This method may poll for completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine. parameters is
// parameters supplied to the Capture Virtual Machine operation.
func (client VirtualMachinesClient) Capture(ctx context.Context, resourceGroupName string, VMName string, parameters VirtualMachineCaptureParameters) (*VirtualMachineCaptureResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.VhdPrefix", name: null, rule: true, chain: nil},
				{target: "parameters.DestinationContainerName", name: null, rule: true, chain: nil},
				{target: "parameters.OverwriteVhds", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.capturePreparer(resourceGroupName, VMName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.captureResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualMachineCaptureResult), err
}

// capturePreparer prepares the Capture request.
func (client VirtualMachinesClient) capturePreparer(resourceGroupName string, VMName string, parameters VirtualMachineCaptureParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/capture"
	req, err := pipeline.NewRequest("POST", u, nil)
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

// captureResponder handles the response to the Capture request.
func (client VirtualMachinesClient) captureResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &VirtualMachineCaptureResult{rawResponse: resp.Response()}
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

// ConvertToManagedDisks converts virtual machine disks from blob-based to managed disks. Virtual machine must be
// stop-deallocated before invoking this operation. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) ConvertToManagedDisks(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.convertToManagedDisksPreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.convertToManagedDisksResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// convertToManagedDisksPreparer prepares the ConvertToManagedDisks request.
func (client VirtualMachinesClient) convertToManagedDisksPreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/convertToManagedDisks"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// convertToManagedDisksResponder handles the response to the ConvertToManagedDisks request.
func (client VirtualMachinesClient) convertToManagedDisksResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// CreateOrUpdate the operation to create or update a virtual machine. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine. parameters is
// parameters supplied to the Create Virtual Machine operation.
func (client VirtualMachinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, parameters VirtualMachine) (*VirtualMachine, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.VirtualMachineProperties", name: null, rule: false,
				chain: []constraint{{target: "parameters.VirtualMachineProperties.StorageProfile", name: null, rule: false,
					chain: []constraint{{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk", name: null, rule: false,
						chain: []constraint{{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings", name: null, rule: false,
							chain: []constraint{{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey", name: null, rule: false,
								chain: []constraint{{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey.SecretURL", name: null, rule: true, chain: nil},
									{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey.SourceVault", name: null, rule: true, chain: nil},
								}},
								{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey", name: null, rule: false,
									chain: []constraint{{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey.KeyURL", name: null, rule: true, chain: nil},
										{target: "parameters.VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey.SourceVault", name: null, rule: true, chain: nil},
									}},
							}},
						}},
					}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, VMName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualMachine), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachinesClient) createOrUpdatePreparer(resourceGroupName string, VMName string, parameters VirtualMachine) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"
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
func (client VirtualMachinesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &VirtualMachine{rawResponse: resp.Response()}
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

// Deallocate shuts down the virtual machine and releases the compute resources. You are not billed for the compute
// resources that this virtual machine uses. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) Deallocate(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.deallocatePreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deallocateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// deallocatePreparer prepares the Deallocate request.
func (client VirtualMachinesClient) deallocatePreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/deallocate"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deallocateResponder handles the response to the Deallocate request.
func (client VirtualMachinesClient) deallocateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// Delete the operation to delete a virtual machine. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) Delete(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.deletePreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// deletePreparer prepares the Delete request.
func (client VirtualMachinesClient) deletePreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"
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
func (client VirtualMachinesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// Generalize sets the state of the virtual machine to generalized.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) Generalize(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.generalizePreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.generalizeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// generalizePreparer prepares the Generalize request.
func (client VirtualMachinesClient) generalizePreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/generalize"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// generalizeResponder handles the response to the Generalize request.
func (client VirtualMachinesClient) generalizeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// Get retrieves information about the model view or the instance view of a virtual machine.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine. expand is the expand
// expression to apply on the operation.
func (client VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, VMName string, expand InstanceViewTypesType) (*VirtualMachine, error) {
	req, err := client.getPreparer(resourceGroupName, VMName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualMachine), err
}

// getPreparer prepares the Get request.
func (client VirtualMachinesClient) getPreparer(resourceGroupName string, VMName string, expand InstanceViewTypesType) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if expand != InstanceViewTypesNone {
		params.Set("$expand", fmt.Sprintf("%v", expand))
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client VirtualMachinesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualMachine{rawResponse: resp.Response()}
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

// List lists all of the virtual machines in the specified resource group. Use the nextLink property in the response to
// get the next page of virtual machines.
//
// resourceGroupName is the name of the resource group.
func (client VirtualMachinesClient) List(ctx context.Context, resourceGroupName string) (*VirtualMachineListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualMachineListResult), err
}

// listPreparer prepares the List request.
func (client VirtualMachinesClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines"
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
func (client VirtualMachinesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualMachineListResult{rawResponse: resp.Response()}
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

// ListAll lists all of the virtual machines in the specified subscription. Use the nextLink property in the response
// to get the next page of virtual machines.
func (client VirtualMachinesClient) ListAll(ctx context.Context) (*VirtualMachineListResult, error) {
	req, err := client.listAllPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAllResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualMachineListResult), err
}

// listAllPreparer prepares the ListAll request.
func (client VirtualMachinesClient) listAllPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines"
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
func (client VirtualMachinesClient) listAllResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualMachineListResult{rawResponse: resp.Response()}
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

// ListAvailableSizes lists all available virtual machine sizes to which the specified virtual machine can be resized.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) ListAvailableSizes(ctx context.Context, resourceGroupName string, VMName string) (*VirtualMachineSizeListResult, error) {
	req, err := client.listAvailableSizesPreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAvailableSizesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*VirtualMachineSizeListResult), err
}

// listAvailableSizesPreparer prepares the ListAvailableSizes request.
func (client VirtualMachinesClient) listAvailableSizesPreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/vmSizes"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listAvailableSizesResponder handles the response to the ListAvailableSizes request.
func (client VirtualMachinesClient) listAvailableSizesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &VirtualMachineSizeListResult{rawResponse: resp.Response()}
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

// PowerOff the operation to power off (stop) a virtual machine. The virtual machine can be restarted with the same
// provisioned resources. You are still charged for this virtual machine. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) PowerOff(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.powerOffPreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.powerOffResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// powerOffPreparer prepares the PowerOff request.
func (client VirtualMachinesClient) powerOffPreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/powerOff"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// powerOffResponder handles the response to the PowerOff request.
func (client VirtualMachinesClient) powerOffResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// Redeploy the operation to redeploy a virtual machine. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) Redeploy(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.redeployPreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.redeployResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// redeployPreparer prepares the Redeploy request.
func (client VirtualMachinesClient) redeployPreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/redeploy"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// redeployResponder handles the response to the Redeploy request.
func (client VirtualMachinesClient) redeployResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// Restart the operation to restart a virtual machine. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) Restart(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.restartPreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.restartResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// restartPreparer prepares the Restart request.
func (client VirtualMachinesClient) restartPreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/restart"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// restartResponder handles the response to the Restart request.
func (client VirtualMachinesClient) restartResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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

// Start the operation to start a virtual machine. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. VMName is the name of the virtual machine.
func (client VirtualMachinesClient) Start(ctx context.Context, resourceGroupName string, VMName string) (*OperationStatusResponse, error) {
	req, err := client.startPreparer(resourceGroupName, VMName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.startResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*OperationStatusResponse), err
}

// startPreparer prepares the Start request.
func (client VirtualMachinesClient) startPreparer(resourceGroupName string, VMName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/start"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// startResponder handles the response to the Start request.
func (client VirtualMachinesClient) startResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &OperationStatusResponse{rawResponse: resp.Response()}
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
