//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/GetSettings_example.json
func ExampleSettingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSettingsClient().NewListPager(nil)
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
		// page.SettingsList = armsecurity.SettingsList{
		// 	Value: []armsecurity.SettingClassification{
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("MCAS"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/MCAS"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(true),
		// 			},
		// 		},
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("WDATP"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/WDATP"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 		},
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("WDATP_EXCLUDE_LINUX_PUBLIC_PREVIEW"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/WDATP_EXCLUDE_LINUX_PUBLIC_PREVIEW"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 		},
		// 		&armsecurity.DataExportSettings{
		// 			Name: to.Ptr("WDATP_UNIFIED_SOLUTION"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/WDATP_UNIFIED_SOLUTION"),
		// 			Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		// 			Properties: &armsecurity.DataExportSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 		},
		// 		&armsecurity.AlertSyncSettings{
		// 			Name: to.Ptr("Sentinel"),
		// 			Type: to.Ptr("Microsoft.Security/settings"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/Sentinel"),
		// 			Kind: to.Ptr(armsecurity.SettingKindAlertSyncSettings),
		// 			Properties: &armsecurity.AlertSyncSettingProperties{
		// 				Enabled: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/GetSetting_example.json
func ExampleSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSettingsClient().Get(ctx, armsecurity.SettingNameMCAS, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.SettingsClientGetResponse{
	// 	                            SettingClassification: &armsecurity.DataExportSettings{
	// 		Name: to.Ptr("MCAS"),
	// 		Type: to.Ptr("Microsoft.Security/settings"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/MCAS"),
	// 		Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
	// 		Properties: &armsecurity.DataExportSettingProperties{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/UpdateSetting_example.json
func ExampleSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSettingsClient().Update(ctx, armsecurity.SettingNameMCAS, &armsecurity.DataExportSettings{
		Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
		Properties: &armsecurity.DataExportSettingProperties{
			Enabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.SettingsClientUpdateResponse{
	// 	                            SettingClassification: &armsecurity.DataExportSettings{
	// 		Name: to.Ptr("MCAS"),
	// 		Type: to.Ptr("Microsoft.Security/settings"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/settings/MCAS"),
	// 		Kind: to.Ptr(armsecurity.SettingKindDataExportSettings),
	// 		Properties: &armsecurity.DataExportSettingProperties{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 	},
	// 	                        }
}
