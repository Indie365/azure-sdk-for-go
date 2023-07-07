//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armazuresphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuresphere/armazuresphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetDeviceGroups.json
func ExampleDeviceGroupsClient_NewListByProductPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeviceGroupsClient().NewListByProductPager("MyResourceGroup1", "MyCatalog1", "MyProduct1", &armazuresphere.DeviceGroupsClientListByProductOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DeviceGroupListResult = armazuresphere.DeviceGroupListResult{
		// 	Value: []*armazuresphere.DeviceGroup{
		// 		{
		// 			Name: to.Ptr("MyDeviceGroup1"),
		// 			Type: to.Ptr("microsoft.azureSphere/catalogs/products/devicegroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct1/devicegroups/MyDeviceGroup1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MyDeviceGroup2"),
		// 			Type: to.Ptr("microsoft.azureSphere/catalogs/products/devicegroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/Products/MyProduct2/devicegroups/MyDeviceGroup2"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetDeviceGroup.json
func ExampleDeviceGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceGroupsClient().Get(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceGroup = armazuresphere.DeviceGroup{
	// 	Type: to.Ptr("microsoft.azureSphere/catalogs/products/devicegroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct1/devicegroups/MyDeviceGroup1"),
	// 	Properties: &armazuresphere.DeviceGroupProperties{
	// 		Description: to.Ptr("The description of MyDeviceGroup1"),
	// 		OSFeedType: to.Ptr(armazuresphere.OSFeedTypeRetail),
	// 		UpdatePolicy: to.Ptr(armazuresphere.UpdatePolicyUpdateAll),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutDeviceGroup.json
func ExampleDeviceGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceGroupsClient().BeginCreateOrUpdate(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", armazuresphere.DeviceGroup{
		Properties: &armazuresphere.DeviceGroupProperties{
			Description:  to.Ptr("Description for MyDeviceGroup1"),
			OSFeedType:   to.Ptr(armazuresphere.OSFeedTypeRetail),
			UpdatePolicy: to.Ptr(armazuresphere.UpdatePolicyUpdateAll),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceGroup = armazuresphere.DeviceGroup{
	// 	Type: to.Ptr("microsoft.AzureSphere/catalogs/products/devicegroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct1/devicegroups/MyDeviceGroup1"),
	// 	Properties: &armazuresphere.DeviceGroupProperties{
	// 		Description: to.Ptr("Description of MyDeviceGroup1"),
	// 		OSFeedType: to.Ptr(armazuresphere.OSFeedTypeRetail),
	// 		UpdatePolicy: to.Ptr(armazuresphere.UpdatePolicyUpdateAll),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteDeviceGroup.json
func ExampleDeviceGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceGroupsClient().BeginDelete(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PatchDeviceGroup.json
func ExampleDeviceGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceGroupsClient().BeginUpdate(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", armazuresphere.DeviceGroupUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceGroup = armazuresphere.DeviceGroup{
	// 	Type: to.Ptr("microsoft.azureSphere/catalogs/products/devicegroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct1/devicegroups/MyDeviceGroup1"),
	// 	Properties: &armazuresphere.DeviceGroupProperties{
	// 		Description: to.Ptr("The description of MyDeviceGroup1"),
	// 		OSFeedType: to.Ptr(armazuresphere.OSFeedTypeRetail),
	// 		UpdatePolicy: to.Ptr(armazuresphere.UpdatePolicyUpdateAll),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostClaimDevices.json
func ExampleDeviceGroupsClient_BeginClaimDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceGroupsClient().BeginClaimDevices(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", armazuresphere.ClaimDevicesRequest{
		DeviceIdentifiers: []*string{
			to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostCountDevicesDeviceGroup.json
func ExampleDeviceGroupsClient_CountDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceGroupsClient().CountDevices(ctx, "MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CountDeviceResponse = armazuresphere.CountDeviceResponse{
	// 	Value: to.Ptr[int32](3),
	// }
}
