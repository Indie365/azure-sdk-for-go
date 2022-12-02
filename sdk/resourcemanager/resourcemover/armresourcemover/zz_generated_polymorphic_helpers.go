//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import "encoding/json"

func unmarshalResourceSettingsClassification(rawMsg json.RawMessage) (ResourceSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceSettingsClassification
	switch m["resourceType"] {
	case "Microsoft.Compute/availabilitySets":
		b = &AvailabilitySetResourceSettings{}
	case "Microsoft.Compute/diskEncryptionSets":
		b = &DiskEncryptionSetResourceSettings{}
	case "Microsoft.Compute/virtualMachines":
		b = &VirtualMachineResourceSettings{}
	case "Microsoft.KeyVault/vaults":
		b = &KeyVaultResourceSettings{}
	case "Microsoft.Network/loadBalancers":
		b = &LoadBalancerResourceSettings{}
	case "Microsoft.Network/networkInterfaces":
		b = &NetworkInterfaceResourceSettings{}
	case "Microsoft.Network/networkSecurityGroups":
		b = &NetworkSecurityGroupResourceSettings{}
	case "Microsoft.Network/publicIPAddresses":
		b = &PublicIPAddressResourceSettings{}
	case "Microsoft.Network/virtualNetworks":
		b = &VirtualNetworkResourceSettings{}
	case "Microsoft.Sql/servers":
		b = &SQLServerResourceSettings{}
	case "Microsoft.Sql/servers/databases":
		b = &SQLDatabaseResourceSettings{}
	case "Microsoft.Sql/servers/elasticPools":
		b = &SQLElasticPoolResourceSettings{}
	case "resourceGroups":
		b = &ResourceGroupResourceSettings{}
	default:
		b = &ResourceSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}