//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

// RecordSetsClientCreateOrUpdateResponse contains the response from method RecordSetsClient.CreateOrUpdate.
type RecordSetsClientCreateOrUpdateResponse struct {
	RecordSet
}

// RecordSetsClientDeleteResponse contains the response from method RecordSetsClient.Delete.
type RecordSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// RecordSetsClientGetResponse contains the response from method RecordSetsClient.Get.
type RecordSetsClientGetResponse struct {
	RecordSet
}

// RecordSetsClientListAllByDNSZoneResponse contains the response from method RecordSetsClient.ListAllByDNSZone.
type RecordSetsClientListAllByDNSZoneResponse struct {
	RecordSetListResult
}

// RecordSetsClientListByDNSZoneResponse contains the response from method RecordSetsClient.ListByDNSZone.
type RecordSetsClientListByDNSZoneResponse struct {
	RecordSetListResult
}

// RecordSetsClientListByTypeResponse contains the response from method RecordSetsClient.ListByType.
type RecordSetsClientListByTypeResponse struct {
	RecordSetListResult
}

// RecordSetsClientUpdateResponse contains the response from method RecordSetsClient.Update.
type RecordSetsClientUpdateResponse struct {
	RecordSet
}

// ResourceReferenceClientGetByTargetResourcesResponse contains the response from method ResourceReferenceClient.GetByTargetResources.
type ResourceReferenceClientGetByTargetResourcesResponse struct {
	ResourceReferenceResult
}

// ZonesClientCreateOrUpdateResponse contains the response from method ZonesClient.CreateOrUpdate.
type ZonesClientCreateOrUpdateResponse struct {
	Zone
}

// ZonesClientDeleteResponse contains the response from method ZonesClient.Delete.
type ZonesClientDeleteResponse struct {
	// placeholder for future response values
}

// ZonesClientGetResponse contains the response from method ZonesClient.Get.
type ZonesClientGetResponse struct {
	Zone
}

// ZonesClientListByResourceGroupResponse contains the response from method ZonesClient.ListByResourceGroup.
type ZonesClientListByResourceGroupResponse struct {
	ZoneListResult
}

// ZonesClientListResponse contains the response from method ZonesClient.List.
type ZonesClientListResponse struct {
	ZoneListResult
}

// ZonesClientUpdateResponse contains the response from method ZonesClient.Update.
type ZonesClientUpdateResponse struct {
	Zone
}
