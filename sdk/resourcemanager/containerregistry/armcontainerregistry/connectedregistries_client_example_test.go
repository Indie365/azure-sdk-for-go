//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ConnectedRegistryList.json
func ExampleConnectedRegistriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedRegistriesClient().NewListPager("myResourceGroup", "myRegistry", &armcontainerregistry.ConnectedRegistriesClientListOptions{Filter: nil})
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
		// page.ConnectedRegistryListResult = armcontainerregistry.ConnectedRegistryListResult{
		// 	Value: []*armcontainerregistry.ConnectedRegistry{
		// 		{
		// 			Name: to.Ptr("myConnectedRegistry"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
		// 			Properties: &armcontainerregistry.ConnectedRegistryProperties{
		// 				Activation: &armcontainerregistry.ActivationProperties{
		// 					Status: to.Ptr(armcontainerregistry.ActivationStatusInactive),
		// 				},
		// 				ClientTokenIDs: []*string{
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
		// 					Logging: &armcontainerregistry.LoggingProperties{
		// 						AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusDisabled),
		// 						LogLevel: to.Ptr(armcontainerregistry.LogLevelInformation),
		// 					},
		// 					Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
		// 					NotificationsList: []*string{
		// 						to.Ptr("hello-world:*:*"),
		// 						to.Ptr("sample/repo/*:1.0:*")},
		// 						Parent: &armcontainerregistry.ParentProperties{
		// 							SyncProperties: &armcontainerregistry.SyncProperties{
		// 								MessageTTL: to.Ptr("P2D"),
		// 								Schedule: to.Ptr("0 9 * * *"),
		// 								SyncWindow: to.Ptr("PT3H"),
		// 								TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
		// 							},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ConnectedRegistryGet.json
func ExampleConnectedRegistriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedRegistriesClient().Get(ctx, "myResourceGroup", "myRegistry", "myConnectedRegistry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedRegistry = armcontainerregistry.ConnectedRegistry{
	// 	Name: to.Ptr("myConnectedRegistry"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
	// 	Properties: &armcontainerregistry.ConnectedRegistryProperties{
	// 		Activation: &armcontainerregistry.ActivationProperties{
	// 			Status: to.Ptr(armcontainerregistry.ActivationStatusInactive),
	// 		},
	// 		ClientTokenIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusDisabled),
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelInformation),
	// 			},
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*")},
	// 				Parent: &armcontainerregistry.ParentProperties{
	// 					SyncProperties: &armcontainerregistry.SyncProperties{
	// 						MessageTTL: to.Ptr("P2D"),
	// 						Schedule: to.Ptr("0 9 * * *"),
	// 						SyncWindow: to.Ptr("PT3H"),
	// 						TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					},
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ConnectedRegistryCreate.json
func ExampleConnectedRegistriesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myConnectedRegistry", armcontainerregistry.ConnectedRegistry{
		Properties: &armcontainerregistry.ConnectedRegistryProperties{
			ClientTokenIDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
			NotificationsList: []*string{
				to.Ptr("hello-world:*:*"),
				to.Ptr("sample/repo/*:1.0:*")},
			Parent: &armcontainerregistry.ParentProperties{
				SyncProperties: &armcontainerregistry.SyncProperties{
					MessageTTL: to.Ptr("P2D"),
					Schedule:   to.Ptr("0 9 * * *"),
					SyncWindow: to.Ptr("PT3H"),
					TokenID:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
				},
			},
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
	// res.ConnectedRegistry = armcontainerregistry.ConnectedRegistry{
	// 	Name: to.Ptr("myConnectedRegistry"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
	// 	Properties: &armcontainerregistry.ConnectedRegistryProperties{
	// 		Activation: &armcontainerregistry.ActivationProperties{
	// 			Status: to.Ptr(armcontainerregistry.ActivationStatusInactive),
	// 		},
	// 		ClientTokenIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token")},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusDisabled),
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelInformation),
	// 			},
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*")},
	// 				Parent: &armcontainerregistry.ParentProperties{
	// 					SyncProperties: &armcontainerregistry.SyncProperties{
	// 						MessageTTL: to.Ptr("P2D"),
	// 						Schedule: to.Ptr("0 9 * * *"),
	// 						SyncWindow: to.Ptr("PT3H"),
	// 						TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					},
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ConnectedRegistryDelete.json
func ExampleConnectedRegistriesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginDelete(ctx, "myResourceGroup", "myRegistry", "myConnectedRegistry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ConnectedRegistryUpdate.json
func ExampleConnectedRegistriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myScopeMap", armcontainerregistry.ConnectedRegistryUpdateParameters{
		Properties: &armcontainerregistry.ConnectedRegistryUpdateProperties{
			ClientTokenIDs: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")},
			Logging: &armcontainerregistry.LoggingProperties{
				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
				LogLevel:       to.Ptr(armcontainerregistry.LogLevelDebug),
			},
			NotificationsList: []*string{
				to.Ptr("hello-world:*:*"),
				to.Ptr("sample/repo/*:1.0:*")},
			SyncProperties: &armcontainerregistry.SyncUpdateProperties{
				MessageTTL: to.Ptr("P30D"),
				Schedule:   to.Ptr("0 0 */10 * *"),
				SyncWindow: to.Ptr("P2D"),
			},
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
	// res.ConnectedRegistry = armcontainerregistry.ConnectedRegistry{
	// 	Name: to.Ptr("myConnectedRegistry"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/connectedRegistries"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/connectedRegistries/myConnectedRegistry"),
	// 	Properties: &armcontainerregistry.ConnectedRegistryProperties{
	// 		ClientTokenIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")},
	// 			Logging: &armcontainerregistry.LoggingProperties{
	// 				AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
	// 				LogLevel: to.Ptr(armcontainerregistry.LogLevelDebug),
	// 			},
	// 			Mode: to.Ptr(armcontainerregistry.ConnectedRegistryModeReadWrite),
	// 			NotificationsList: []*string{
	// 				to.Ptr("hello-world:*:*"),
	// 				to.Ptr("sample/repo/*:1.0:*")},
	// 				Parent: &armcontainerregistry.ParentProperties{
	// 					SyncProperties: &armcontainerregistry.SyncProperties{
	// 						MessageTTL: to.Ptr("P30D"),
	// 						Schedule: to.Ptr("0 0 */10 * *"),
	// 						SyncWindow: to.Ptr("P2D"),
	// 						TokenID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken"),
	// 					},
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/ConnectedRegistryDeactivate.json
func ExampleConnectedRegistriesClient_BeginDeactivate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectedRegistriesClient().BeginDeactivate(ctx, "myResourceGroup", "myRegistry", "myConnectedRegistry", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
