//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// CustomizationPoliciesGetResponse contains the response from method CustomizationPolicies.Get.
type CustomizationPoliciesGetResponse struct {
	CustomizationPoliciesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomizationPoliciesGetResult contains the result from method CustomizationPolicies.Get.
type CustomizationPoliciesGetResult struct {
	CustomizationPolicy
}

// CustomizationPoliciesListResponseEnvelope contains the response from method CustomizationPolicies.List.
type CustomizationPoliciesListResponseEnvelope struct {
	CustomizationPoliciesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomizationPoliciesListResult contains the result from method CustomizationPolicies.List.
type CustomizationPoliciesListResult struct {
	CustomizationPoliciesListResponse
}

// DedicatedCloudNodesCreateOrUpdatePollerResponse contains the response from method DedicatedCloudNodes.CreateOrUpdate.
type DedicatedCloudNodesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DedicatedCloudNodesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DedicatedCloudNodesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DedicatedCloudNodesCreateOrUpdateResponse, error) {
	respType := DedicatedCloudNodesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DedicatedCloudNode)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DedicatedCloudNodesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *DedicatedCloudNodesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *DedicatedCloudNodesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DedicatedCloudNodesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &DedicatedCloudNodesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DedicatedCloudNodesCreateOrUpdateResponse contains the response from method DedicatedCloudNodes.CreateOrUpdate.
type DedicatedCloudNodesCreateOrUpdateResponse struct {
	DedicatedCloudNodesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudNodesCreateOrUpdateResult contains the result from method DedicatedCloudNodes.CreateOrUpdate.
type DedicatedCloudNodesCreateOrUpdateResult struct {
	DedicatedCloudNode
}

// DedicatedCloudNodesDeleteResponse contains the response from method DedicatedCloudNodes.Delete.
type DedicatedCloudNodesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudNodesGetResponse contains the response from method DedicatedCloudNodes.Get.
type DedicatedCloudNodesGetResponse struct {
	DedicatedCloudNodesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudNodesGetResult contains the result from method DedicatedCloudNodes.Get.
type DedicatedCloudNodesGetResult struct {
	DedicatedCloudNode
}

// DedicatedCloudNodesListByResourceGroupResponse contains the response from method DedicatedCloudNodes.ListByResourceGroup.
type DedicatedCloudNodesListByResourceGroupResponse struct {
	DedicatedCloudNodesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudNodesListByResourceGroupResult contains the result from method DedicatedCloudNodes.ListByResourceGroup.
type DedicatedCloudNodesListByResourceGroupResult struct {
	DedicatedCloudNodeListResponse
}

// DedicatedCloudNodesListBySubscriptionResponse contains the response from method DedicatedCloudNodes.ListBySubscription.
type DedicatedCloudNodesListBySubscriptionResponse struct {
	DedicatedCloudNodesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudNodesListBySubscriptionResult contains the result from method DedicatedCloudNodes.ListBySubscription.
type DedicatedCloudNodesListBySubscriptionResult struct {
	DedicatedCloudNodeListResponse
}

// DedicatedCloudNodesUpdateResponse contains the response from method DedicatedCloudNodes.Update.
type DedicatedCloudNodesUpdateResponse struct {
	DedicatedCloudNodesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudNodesUpdateResult contains the result from method DedicatedCloudNodes.Update.
type DedicatedCloudNodesUpdateResult struct {
	DedicatedCloudNode
}

// DedicatedCloudServicesCreateOrUpdateResponse contains the response from method DedicatedCloudServices.CreateOrUpdate.
type DedicatedCloudServicesCreateOrUpdateResponse struct {
	DedicatedCloudServicesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudServicesCreateOrUpdateResult contains the result from method DedicatedCloudServices.CreateOrUpdate.
type DedicatedCloudServicesCreateOrUpdateResult struct {
	DedicatedCloudService
}

// DedicatedCloudServicesDeletePollerResponse contains the response from method DedicatedCloudServices.Delete.
type DedicatedCloudServicesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DedicatedCloudServicesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DedicatedCloudServicesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DedicatedCloudServicesDeleteResponse, error) {
	respType := DedicatedCloudServicesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DedicatedCloudServicesDeletePollerResponse from the provided client and resume token.
func (l *DedicatedCloudServicesDeletePollerResponse) Resume(ctx context.Context, client *DedicatedCloudServicesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DedicatedCloudServicesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &DedicatedCloudServicesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DedicatedCloudServicesDeleteResponse contains the response from method DedicatedCloudServices.Delete.
type DedicatedCloudServicesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudServicesGetResponse contains the response from method DedicatedCloudServices.Get.
type DedicatedCloudServicesGetResponse struct {
	DedicatedCloudServicesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudServicesGetResult contains the result from method DedicatedCloudServices.Get.
type DedicatedCloudServicesGetResult struct {
	DedicatedCloudService
}

// DedicatedCloudServicesListByResourceGroupResponse contains the response from method DedicatedCloudServices.ListByResourceGroup.
type DedicatedCloudServicesListByResourceGroupResponse struct {
	DedicatedCloudServicesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudServicesListByResourceGroupResult contains the result from method DedicatedCloudServices.ListByResourceGroup.
type DedicatedCloudServicesListByResourceGroupResult struct {
	DedicatedCloudServiceListResponse
}

// DedicatedCloudServicesListBySubscriptionResponse contains the response from method DedicatedCloudServices.ListBySubscription.
type DedicatedCloudServicesListBySubscriptionResponse struct {
	DedicatedCloudServicesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudServicesListBySubscriptionResult contains the result from method DedicatedCloudServices.ListBySubscription.
type DedicatedCloudServicesListBySubscriptionResult struct {
	DedicatedCloudServiceListResponse
}

// DedicatedCloudServicesUpdateResponse contains the response from method DedicatedCloudServices.Update.
type DedicatedCloudServicesUpdateResponse struct {
	DedicatedCloudServicesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DedicatedCloudServicesUpdateResult contains the result from method DedicatedCloudServices.Update.
type DedicatedCloudServicesUpdateResult struct {
	DedicatedCloudService
}

// OperationsGetResponse contains the response from method Operations.Get.
type OperationsGetResponse struct {
	OperationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsGetResult contains the result from method Operations.Get.
type OperationsGetResult struct {
	OperationResource
	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	AvailableOperationsListResponse
}

// PrivateCloudsGetResponse contains the response from method PrivateClouds.Get.
type PrivateCloudsGetResponse struct {
	PrivateCloudsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateCloudsGetResult contains the result from method PrivateClouds.Get.
type PrivateCloudsGetResult struct {
	PrivateCloud
}

// PrivateCloudsListResponse contains the response from method PrivateClouds.List.
type PrivateCloudsListResponse struct {
	PrivateCloudsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateCloudsListResult contains the result from method PrivateClouds.List.
type PrivateCloudsListResult struct {
	PrivateCloudList
}

// ResourcePoolsGetResponse contains the response from method ResourcePools.Get.
type ResourcePoolsGetResponse struct {
	ResourcePoolsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourcePoolsGetResult contains the result from method ResourcePools.Get.
type ResourcePoolsGetResult struct {
	ResourcePool
}

// ResourcePoolsListResponseEnvelope contains the response from method ResourcePools.List.
type ResourcePoolsListResponseEnvelope struct {
	ResourcePoolsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourcePoolsListResult contains the result from method ResourcePools.List.
type ResourcePoolsListResult struct {
	ResourcePoolsListResponse
}

// SKUsAvailabilityListResponse contains the response from method SKUsAvailability.List.
type SKUsAvailabilityListResponse struct {
	SKUsAvailabilityListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SKUsAvailabilityListResult contains the result from method SKUsAvailability.List.
type SKUsAvailabilityListResult struct {
	SKUAvailabilityListResponse
}

// UsagesListResponse contains the response from method Usages.List.
type UsagesListResponse struct {
	UsagesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UsagesListResult contains the result from method Usages.List.
type UsagesListResult struct {
	UsageListResponse
}

// VirtualMachineTemplatesGetResponse contains the response from method VirtualMachineTemplates.Get.
type VirtualMachineTemplatesGetResponse struct {
	VirtualMachineTemplatesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachineTemplatesGetResult contains the result from method VirtualMachineTemplates.Get.
type VirtualMachineTemplatesGetResult struct {
	VirtualMachineTemplate
}

// VirtualMachineTemplatesListResponse contains the response from method VirtualMachineTemplates.List.
type VirtualMachineTemplatesListResponse struct {
	VirtualMachineTemplatesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachineTemplatesListResult contains the result from method VirtualMachineTemplates.List.
type VirtualMachineTemplatesListResult struct {
	VirtualMachineTemplateListResponse
}

// VirtualMachinesCreateOrUpdatePollerResponse contains the response from method VirtualMachines.CreateOrUpdate.
type VirtualMachinesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualMachinesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VirtualMachinesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualMachinesCreateOrUpdateResponse, error) {
	respType := VirtualMachinesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.VirtualMachine)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualMachinesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *VirtualMachinesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *VirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualMachinesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualMachinesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualMachinesCreateOrUpdateResponse contains the response from method VirtualMachines.CreateOrUpdate.
type VirtualMachinesCreateOrUpdateResponse struct {
	VirtualMachinesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesCreateOrUpdateResult contains the result from method VirtualMachines.CreateOrUpdate.
type VirtualMachinesCreateOrUpdateResult struct {
	VirtualMachine
}

// VirtualMachinesDeletePollerResponse contains the response from method VirtualMachines.Delete.
type VirtualMachinesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualMachinesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VirtualMachinesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualMachinesDeleteResponse, error) {
	respType := VirtualMachinesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualMachinesDeletePollerResponse from the provided client and resume token.
func (l *VirtualMachinesDeletePollerResponse) Resume(ctx context.Context, client *VirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualMachinesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualMachinesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualMachinesDeleteResponse contains the response from method VirtualMachines.Delete.
type VirtualMachinesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesGetResponse contains the response from method VirtualMachines.Get.
type VirtualMachinesGetResponse struct {
	VirtualMachinesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesGetResult contains the result from method VirtualMachines.Get.
type VirtualMachinesGetResult struct {
	VirtualMachine
}

// VirtualMachinesListByResourceGroupResponse contains the response from method VirtualMachines.ListByResourceGroup.
type VirtualMachinesListByResourceGroupResponse struct {
	VirtualMachinesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesListByResourceGroupResult contains the result from method VirtualMachines.ListByResourceGroup.
type VirtualMachinesListByResourceGroupResult struct {
	VirtualMachineListResponse
}

// VirtualMachinesListBySubscriptionResponse contains the response from method VirtualMachines.ListBySubscription.
type VirtualMachinesListBySubscriptionResponse struct {
	VirtualMachinesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesListBySubscriptionResult contains the result from method VirtualMachines.ListBySubscription.
type VirtualMachinesListBySubscriptionResult struct {
	VirtualMachineListResponse
}

// VirtualMachinesStartPollerResponse contains the response from method VirtualMachines.Start.
type VirtualMachinesStartPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualMachinesStartPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VirtualMachinesStartPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualMachinesStartResponse, error) {
	respType := VirtualMachinesStartResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualMachinesStartPollerResponse from the provided client and resume token.
func (l *VirtualMachinesStartPollerResponse) Resume(ctx context.Context, client *VirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualMachinesClient.Start", token, client.pl, client.startHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualMachinesStartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualMachinesStartResponse contains the response from method VirtualMachines.Start.
type VirtualMachinesStartResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesStopPollerResponse contains the response from method VirtualMachines.Stop.
type VirtualMachinesStopPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualMachinesStopPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VirtualMachinesStopPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualMachinesStopResponse, error) {
	respType := VirtualMachinesStopResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualMachinesStopPollerResponse from the provided client and resume token.
func (l *VirtualMachinesStopPollerResponse) Resume(ctx context.Context, client *VirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualMachinesClient.Stop", token, client.pl, client.stopHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualMachinesStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualMachinesStopResponse contains the response from method VirtualMachines.Stop.
type VirtualMachinesStopResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesUpdatePollerResponse contains the response from method VirtualMachines.Update.
type VirtualMachinesUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VirtualMachinesUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VirtualMachinesUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VirtualMachinesUpdateResponse, error) {
	respType := VirtualMachinesUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.VirtualMachine)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VirtualMachinesUpdatePollerResponse from the provided client and resume token.
func (l *VirtualMachinesUpdatePollerResponse) Resume(ctx context.Context, client *VirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VirtualMachinesClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &VirtualMachinesUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VirtualMachinesUpdateResponse contains the response from method VirtualMachines.Update.
type VirtualMachinesUpdateResponse struct {
	VirtualMachinesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualMachinesUpdateResult contains the result from method VirtualMachines.Update.
type VirtualMachinesUpdateResult struct {
	VirtualMachine
}

// VirtualNetworksGetResponse contains the response from method VirtualNetworks.Get.
type VirtualNetworksGetResponse struct {
	VirtualNetworksGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworksGetResult contains the result from method VirtualNetworks.Get.
type VirtualNetworksGetResult struct {
	VirtualNetwork
}

// VirtualNetworksListResponse contains the response from method VirtualNetworks.List.
type VirtualNetworksListResponse struct {
	VirtualNetworksListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VirtualNetworksListResult contains the result from method VirtualNetworks.List.
type VirtualNetworksListResult struct {
	VirtualNetworkListResponse
}
