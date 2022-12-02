//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DedicatedCloudNode.
func (d DedicatedCloudNode) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "sku", d.SKU)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DedicatedCloudNodeProperties.
func (d DedicatedCloudNodeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availabilityZoneId", d.AvailabilityZoneID)
	populate(objectMap, "availabilityZoneName", d.AvailabilityZoneName)
	populate(objectMap, "cloudRackName", d.CloudRackName)
	populateTimeRFC3339(objectMap, "created", d.Created)
	populate(objectMap, "nodesCount", d.NodesCount)
	populate(objectMap, "placementGroupId", d.PlacementGroupID)
	populate(objectMap, "placementGroupName", d.PlacementGroupName)
	populate(objectMap, "privateCloudId", d.PrivateCloudID)
	populate(objectMap, "privateCloudName", d.PrivateCloudName)
	populate(objectMap, "provisioningState", d.ProvisioningState)
	populate(objectMap, "purchaseId", d.PurchaseID)
	populate(objectMap, "skuDescription", d.SKUDescription)
	populate(objectMap, "status", d.Status)
	populate(objectMap, "vmwareClusterName", d.VmwareClusterName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DedicatedCloudNodeProperties.
func (d *DedicatedCloudNodeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "availabilityZoneId":
			err = unpopulate(val, "AvailabilityZoneID", &d.AvailabilityZoneID)
			delete(rawMsg, key)
		case "availabilityZoneName":
			err = unpopulate(val, "AvailabilityZoneName", &d.AvailabilityZoneName)
			delete(rawMsg, key)
		case "cloudRackName":
			err = unpopulate(val, "CloudRackName", &d.CloudRackName)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC3339(val, "Created", &d.Created)
			delete(rawMsg, key)
		case "nodesCount":
			err = unpopulate(val, "NodesCount", &d.NodesCount)
			delete(rawMsg, key)
		case "placementGroupId":
			err = unpopulate(val, "PlacementGroupID", &d.PlacementGroupID)
			delete(rawMsg, key)
		case "placementGroupName":
			err = unpopulate(val, "PlacementGroupName", &d.PlacementGroupName)
			delete(rawMsg, key)
		case "privateCloudId":
			err = unpopulate(val, "PrivateCloudID", &d.PrivateCloudID)
			delete(rawMsg, key)
		case "privateCloudName":
			err = unpopulate(val, "PrivateCloudName", &d.PrivateCloudName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &d.ProvisioningState)
			delete(rawMsg, key)
		case "purchaseId":
			err = unpopulate(val, "PurchaseID", &d.PurchaseID)
			delete(rawMsg, key)
		case "skuDescription":
			err = unpopulate(val, "SKUDescription", &d.SKUDescription)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &d.Status)
			delete(rawMsg, key)
		case "vmwareClusterName":
			err = unpopulate(val, "VmwareClusterName", &d.VmwareClusterName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DedicatedCloudService.
func (d DedicatedCloudService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GuestOSCustomization.
func (g GuestOSCustomization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dnsServers", g.DNSServers)
	populate(objectMap, "hostName", g.HostName)
	populate(objectMap, "password", g.Password)
	populate(objectMap, "policyId", g.PolicyID)
	populate(objectMap, "username", g.Username)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GuestOSNICCustomization.
func (g GuestOSNICCustomization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allocation", g.Allocation)
	populate(objectMap, "dnsServers", g.DNSServers)
	populate(objectMap, "gateway", g.Gateway)
	populate(objectMap, "ipAddress", g.IPAddress)
	populate(objectMap, "mask", g.Mask)
	populate(objectMap, "primaryWinsServer", g.PrimaryWinsServer)
	populate(objectMap, "secondaryWinsServer", g.SecondaryWinsServer)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResource.
func (o *OperationResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &o.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &o.Error)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &o.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &o.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PatchPayload.
func (p PatchPayload) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PrivateCloudProperties.
func (p *PrivateCloudProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "availabilityZoneId":
			err = unpopulate(val, "AvailabilityZoneID", &p.AvailabilityZoneID)
			delete(rawMsg, key)
		case "availabilityZoneName":
			err = unpopulate(val, "AvailabilityZoneName", &p.AvailabilityZoneName)
			delete(rawMsg, key)
		case "clustersNumber":
			err = unpopulate(val, "ClustersNumber", &p.ClustersNumber)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &p.CreatedBy)
			delete(rawMsg, key)
		case "createdOn":
			err = unpopulateTimeRFC3339(val, "CreatedOn", &p.CreatedOn)
			delete(rawMsg, key)
		case "dnsServers":
			err = unpopulate(val, "DNSServers", &p.DNSServers)
			delete(rawMsg, key)
		case "expires":
			err = unpopulate(val, "Expires", &p.Expires)
			delete(rawMsg, key)
		case "nsxType":
			err = unpopulate(val, "NsxType", &p.NsxType)
			delete(rawMsg, key)
		case "placementGroupId":
			err = unpopulate(val, "PlacementGroupID", &p.PlacementGroupID)
			delete(rawMsg, key)
		case "placementGroupName":
			err = unpopulate(val, "PlacementGroupName", &p.PlacementGroupName)
			delete(rawMsg, key)
		case "privateCloudId":
			err = unpopulate(val, "PrivateCloudID", &p.PrivateCloudID)
			delete(rawMsg, key)
		case "resourcePools":
			err = unpopulate(val, "ResourcePools", &p.ResourcePools)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &p.State)
			delete(rawMsg, key)
		case "totalCpuCores":
			err = unpopulate(val, "TotalCPUCores", &p.TotalCPUCores)
			delete(rawMsg, key)
		case "totalNodes":
			err = unpopulate(val, "TotalNodes", &p.TotalNodes)
			delete(rawMsg, key)
		case "totalRam":
			err = unpopulate(val, "TotalRAM", &p.TotalRAM)
			delete(rawMsg, key)
		case "totalStorage":
			err = unpopulate(val, "TotalStorage", &p.TotalStorage)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		case "vSphereVersion":
			err = unpopulate(val, "VSphereVersion", &p.VSphereVersion)
			delete(rawMsg, key)
		case "vcenterFqdn":
			err = unpopulate(val, "VcenterFqdn", &p.VcenterFqdn)
			delete(rawMsg, key)
		case "vcenterRefid":
			err = unpopulate(val, "VcenterRefid", &p.VcenterRefid)
			delete(rawMsg, key)
		case "virtualMachineTemplates":
			err = unpopulate(val, "VirtualMachineTemplates", &p.VirtualMachineTemplates)
			delete(rawMsg, key)
		case "virtualNetworks":
			err = unpopulate(val, "VirtualNetworks", &p.VirtualNetworks)
			delete(rawMsg, key)
		case "vrOpsEnabled":
			err = unpopulate(val, "VrOpsEnabled", &p.VrOpsEnabled)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VirtualMachine.
func (v VirtualMachine) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualMachineProperties.
func (v VirtualMachineProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "amountOfRam", v.AmountOfRAM)
	populate(objectMap, "controllers", v.Controllers)
	populate(objectMap, "customization", v.Customization)
	populate(objectMap, "disks", v.Disks)
	populate(objectMap, "dnsname", v.Dnsname)
	populate(objectMap, "exposeToGuestVM", v.ExposeToGuestVM)
	populate(objectMap, "folder", v.Folder)
	populate(objectMap, "guestOS", v.GuestOS)
	populate(objectMap, "guestOSType", v.GuestOSType)
	populate(objectMap, "nics", v.Nics)
	populate(objectMap, "numberOfCores", v.NumberOfCores)
	populate(objectMap, "password", v.Password)
	populate(objectMap, "privateCloudId", v.PrivateCloudID)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "publicIP", v.PublicIP)
	populate(objectMap, "resourcePool", v.ResourcePool)
	populate(objectMap, "status", v.Status)
	populate(objectMap, "templateId", v.TemplateID)
	populate(objectMap, "username", v.Username)
	populate(objectMap, "vmId", v.VMID)
	populate(objectMap, "vSphereNetworks", v.VSphereNetworks)
	populate(objectMap, "vmwaretools", v.Vmwaretools)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNic.
func (v VirtualNic) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "customization", v.Customization)
	populate(objectMap, "ipAddresses", v.IPAddresses)
	populate(objectMap, "macAddress", v.MacAddress)
	populate(objectMap, "network", v.Network)
	populate(objectMap, "nicType", v.NicType)
	populate(objectMap, "powerOnBoot", v.PowerOnBoot)
	populate(objectMap, "virtualNicId", v.VirtualNicID)
	populate(objectMap, "virtualNicName", v.VirtualNicName)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}