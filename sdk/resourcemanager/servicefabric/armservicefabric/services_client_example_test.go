//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServiceGetOperation_example.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "resRg", "myCluster", "myApp", "myService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armservicefabric.ServiceResource{
	// 	Name: to.Ptr("myCluster"),
	// 	Type: to.Ptr("services"),
	// 	Etag: to.Ptr("W/\"636462502183671258\""),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp/services/myService"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabric.StatelessServiceProperties{
	// 		DefaultMoveCost: to.Ptr(armservicefabric.MoveCostMedium),
	// 		PlacementConstraints: to.Ptr("NodeType==frontend"),
	// 		ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
	// 			{
	// 				Name: to.Ptr("metric1"),
	// 				Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
	// 		}},
	// 		ServicePlacementPolicies: []armservicefabric.ServicePlacementPolicyDescriptionClassification{
	// 		},
	// 		PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
	// 			PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceKind: to.Ptr(armservicefabric.ServiceKindStateless),
	// 		ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
	// 		ServiceTypeName: to.Ptr("myServiceType"),
	// 		InstanceCount: to.Ptr[int32](5),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_max.json
func ExampleServicesClient_BeginCreateOrUpdate_putAServiceWithMaximumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabric.ServiceResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.StatelessServiceProperties{
			CorrelationScheme: []*armservicefabric.ServiceCorrelationDescription{
				{
					Scheme:      to.Ptr(armservicefabric.ServiceCorrelationSchemeAffinity),
					ServiceName: to.Ptr("fabric:/app1/app1~svc1"),
				}},
			DefaultMoveCost:      to.Ptr(armservicefabric.MoveCostMedium),
			PlacementConstraints: to.Ptr("NodeType==frontend"),
			ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
				{
					Name:   to.Ptr("metric1"),
					Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
				}},
			ServicePlacementPolicies: []armservicefabric.ServicePlacementPolicyDescriptionClassification{},
			PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
				PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
			},
			ServiceDNSName:               to.Ptr("my.service.dns"),
			ServiceKind:                  to.Ptr(armservicefabric.ServiceKindStateless),
			ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
			ServiceTypeName:              to.Ptr("myServiceType"),
			InstanceCount:                to.Ptr[int32](5),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_min.json
func ExampleServicesClient_BeginCreateOrUpdate_putAServiceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabric.ServiceResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.StatelessServiceProperties{
			PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
				PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
			},
			ServiceKind:     to.Ptr(armservicefabric.ServiceKindStateless),
			ServiceTypeName: to.Ptr("myServiceType"),
			InstanceCount:   to.Ptr[int32](1),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePatchOperation_example.json
func ExampleServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginUpdate(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabric.ServiceResourceUpdate{
		Tags: map[string]*string{},
		Properties: &armservicefabric.StatelessServiceUpdateProperties{
			ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
				{
					Name:   to.Ptr("metric1"),
					Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
				}},
			ServiceKind: to.Ptr(armservicefabric.ServiceKindStateless),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServiceDeleteOperation_example.json
func ExampleServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginDelete(ctx, "resRg", "myCluster", "myApp", "myService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServiceListOperation_example.json
func ExampleServicesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().List(ctx, "resRg", "myCluster", "myApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResourceList = armservicefabric.ServiceResourceList{
	// 	Value: []*armservicefabric.ServiceResource{
	// 		{
	// 			Name: to.Ptr("myCluster"),
	// 			Type: to.Ptr("services"),
	// 			Etag: to.Ptr("W/\"636462502183671257\""),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp/services/myService"),
	// 			Location: to.Ptr("eastus"),
	// 			Tags: map[string]*string{
	// 			},
	// 			Properties: &armservicefabric.StatelessServiceProperties{
	// 				ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
	// 					{
	// 						Name: to.Ptr("metric1"),
	// 						Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
	// 				}},
	// 				PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
	// 					PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
	// 				},
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				ServiceKind: to.Ptr(armservicefabric.ServiceKindStateless),
	// 				ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
	// 				ServiceTypeName: to.Ptr("myServiceType"),
	// 				InstanceCount: to.Ptr[int32](1),
	// 			},
	// 	}},
	// }
}
