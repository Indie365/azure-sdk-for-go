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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/GetAutomationsSubscription_example.json
func ExampleAutomationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutomationsClient().NewListPager(nil)
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
		// page.AutomationList = armsecurity.AutomationList{
		// 	Value: []*armsecurity.Automation{
		// 		{
		// 			Location: to.Ptr("Central US"),
		// 			Etag: to.Ptr("etag value"),
		// 			Name: to.Ptr("exampleAutomation"),
		// 			Type: to.Ptr("Microsoft.Security/automations"),
		// 			ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armsecurity.AutomationProperties{
		// 				Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
		// 				Actions: []armsecurity.AutomationActionClassification{
		// 					&armsecurity.AutomationActionLogicApp{
		// 						ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
		// 						LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
		// 				}},
		// 				IsEnabled: to.Ptr(true),
		// 				Scopes: []*armsecurity.AutomationScope{
		// 					{
		// 						Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
		// 						ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
		// 				}},
		// 				Sources: []*armsecurity.AutomationSource{
		// 					{
		// 						EventSource: to.Ptr(armsecurity.EventSourceAssessments),
		// 						RuleSets: []*armsecurity.AutomationRuleSet{
		// 							{
		// 								Rules: []*armsecurity.AutomationTriggeringRule{
		// 									{
		// 										ExpectedValue: to.Ptr("customAssessment"),
		// 										Operator: to.Ptr(armsecurity.OperatorEquals),
		// 										PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
		// 										PropertyType: to.Ptr(armsecurity.PropertyTypeString),
		// 								}},
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/GetAutomationsResourceGroup_example.json
func ExampleAutomationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutomationsClient().NewListByResourceGroupPager("exampleResourceGroup", nil)
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
		// page.AutomationList = armsecurity.AutomationList{
		// 	Value: []*armsecurity.Automation{
		// 		{
		// 			Location: to.Ptr("Central US"),
		// 			Etag: to.Ptr("etag value"),
		// 			Name: to.Ptr("exampleAutomation"),
		// 			Type: to.Ptr("Microsoft.Security/automations"),
		// 			ID: to.Ptr("/subscriptions/e4272367-5645-4c4e-9c67-3b74b59a6982/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armsecurity.AutomationProperties{
		// 				Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
		// 				Actions: []armsecurity.AutomationActionClassification{
		// 					&armsecurity.AutomationActionLogicApp{
		// 						ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
		// 						LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
		// 				}},
		// 				IsEnabled: to.Ptr(true),
		// 				Scopes: []*armsecurity.AutomationScope{
		// 					{
		// 						Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
		// 						ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
		// 				}},
		// 				Sources: []*armsecurity.AutomationSource{
		// 					{
		// 						EventSource: to.Ptr(armsecurity.EventSourceAssessments),
		// 						RuleSets: []*armsecurity.AutomationRuleSet{
		// 							{
		// 								Rules: []*armsecurity.AutomationTriggeringRule{
		// 									{
		// 										ExpectedValue: to.Ptr("customAssessment"),
		// 										Operator: to.Ptr(armsecurity.OperatorEquals),
		// 										PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
		// 										PropertyType: to.Ptr(armsecurity.PropertyTypeString),
		// 								}},
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/GetAutomationResourceGroup_example.json
func ExampleAutomationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationsClient().Get(ctx, "exampleResourceGroup", "exampleAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Automation = armsecurity.Automation{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("etag value"),
	// 	Name: to.Ptr("exampleAutomation"),
	// 	Type: to.Ptr("Microsoft.Security/automations"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.AutomationProperties{
	// 		Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
	// 		Actions: []armsecurity.AutomationActionClassification{
	// 			&armsecurity.AutomationActionLogicApp{
	// 				ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
	// 				LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
	// 		}},
	// 		IsEnabled: to.Ptr(true),
	// 		Scopes: []*armsecurity.AutomationScope{
	// 			{
	// 				Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
	// 				ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
	// 		}},
	// 		Sources: []*armsecurity.AutomationSource{
	// 			{
	// 				EventSource: to.Ptr(armsecurity.EventSourceAssessments),
	// 				RuleSets: []*armsecurity.AutomationRuleSet{
	// 					{
	// 						Rules: []*armsecurity.AutomationTriggeringRule{
	// 							{
	// 								ExpectedValue: to.Ptr("customAssessment"),
	// 								Operator: to.Ptr(armsecurity.OperatorEquals),
	// 								PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
	// 								PropertyType: to.Ptr(armsecurity.PropertyTypeString),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutAutomationAllAssessments_example.json
func ExampleAutomationsClient_CreateOrUpdate_createOrUpdateASecurityAutomationForAllAssessmentsIncludingAllSeverities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleAutomation", armsecurity.Automation{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.AutomationProperties{
			Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment"),
			Actions: []armsecurity.AutomationActionClassification{
				&armsecurity.AutomationActionLogicApp{
					ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
					LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
					URI:                to.Ptr("https://exampleTriggerUri1.com"),
				}},
			IsEnabled: to.Ptr(true),
			Scopes: []*armsecurity.AutomationScope{
				{
					Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
					ScopePath:   to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
				}},
			Sources: []*armsecurity.AutomationSource{
				{
					EventSource: to.Ptr(armsecurity.EventSourceAssessments),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Automation = armsecurity.Automation{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("new etag value"),
	// 	Name: to.Ptr("exampleAutomation"),
	// 	Type: to.Ptr("Microsoft.Security/automations"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.AutomationProperties{
	// 		Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment"),
	// 		Actions: []armsecurity.AutomationActionClassification{
	// 			&armsecurity.AutomationActionLogicApp{
	// 				ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
	// 				LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
	// 		}},
	// 		IsEnabled: to.Ptr(true),
	// 		Scopes: []*armsecurity.AutomationScope{
	// 			{
	// 				Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
	// 				ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
	// 		}},
	// 		Sources: []*armsecurity.AutomationSource{
	// 			{
	// 				EventSource: to.Ptr(armsecurity.EventSourceAssessments),
	// 				RuleSets: []*armsecurity.AutomationRuleSet{
	// 				},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutAutomationHighSeverityAssessments_example.json
func ExampleAutomationsClient_CreateOrUpdate_createOrUpdateASecurityAutomationForAllHighSeverityAssessments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleAutomation", armsecurity.Automation{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.AutomationProperties{
			Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any high severity security assessment"),
			Actions: []armsecurity.AutomationActionClassification{
				&armsecurity.AutomationActionLogicApp{
					ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
					LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
					URI:                to.Ptr("https://exampleTriggerUri1.com"),
				}},
			IsEnabled: to.Ptr(true),
			Scopes: []*armsecurity.AutomationScope{
				{
					Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
					ScopePath:   to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
				}},
			Sources: []*armsecurity.AutomationSource{
				{
					EventSource: to.Ptr(armsecurity.EventSourceAssessments),
					RuleSets: []*armsecurity.AutomationRuleSet{
						{
							Rules: []*armsecurity.AutomationTriggeringRule{
								{
									ExpectedValue: to.Ptr("High"),
									Operator:      to.Ptr(armsecurity.OperatorEquals),
									PropertyJPath: to.Ptr("properties.metadata.severity"),
									PropertyType:  to.Ptr(armsecurity.PropertyTypeString),
								}},
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Automation = armsecurity.Automation{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("new etag value"),
	// 	Name: to.Ptr("exampleAutomation"),
	// 	Type: to.Ptr("Microsoft.Security/automations"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.AutomationProperties{
	// 		Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any high severity security assessment"),
	// 		Actions: []armsecurity.AutomationActionClassification{
	// 			&armsecurity.AutomationActionLogicApp{
	// 				ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
	// 				LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
	// 		}},
	// 		IsEnabled: to.Ptr(true),
	// 		Scopes: []*armsecurity.AutomationScope{
	// 			{
	// 				Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
	// 				ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
	// 		}},
	// 		Sources: []*armsecurity.AutomationSource{
	// 			{
	// 				EventSource: to.Ptr(armsecurity.EventSourceAssessments),
	// 				RuleSets: []*armsecurity.AutomationRuleSet{
	// 					{
	// 						Rules: []*armsecurity.AutomationTriggeringRule{
	// 							{
	// 								ExpectedValue: to.Ptr("High"),
	// 								Operator: to.Ptr(armsecurity.OperatorEquals),
	// 								PropertyJPath: to.Ptr("properties.metadata.severity"),
	// 								PropertyType: to.Ptr(armsecurity.PropertyTypeString),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutDisableAutomation_example.json
func ExampleAutomationsClient_CreateOrUpdate_disableOrEnableASecurityAutomation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleAutomation", armsecurity.Automation{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.AutomationProperties{
			Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
			Actions: []armsecurity.AutomationActionClassification{
				&armsecurity.AutomationActionLogicApp{
					ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
					LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
					URI:                to.Ptr("https://exampleTriggerUri1.com"),
				}},
			IsEnabled: to.Ptr(false),
			Scopes: []*armsecurity.AutomationScope{
				{
					Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
					ScopePath:   to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
				}},
			Sources: []*armsecurity.AutomationSource{
				{
					EventSource: to.Ptr(armsecurity.EventSourceAssessments),
					RuleSets: []*armsecurity.AutomationRuleSet{
						{
							Rules: []*armsecurity.AutomationTriggeringRule{
								{
									ExpectedValue: to.Ptr("customAssessment"),
									Operator:      to.Ptr(armsecurity.OperatorEquals),
									PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
									PropertyType:  to.Ptr(armsecurity.PropertyTypeString),
								}},
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Automation = armsecurity.Automation{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("new etag value"),
	// 	Name: to.Ptr("exampleAutomation"),
	// 	Type: to.Ptr("Microsoft.Security/automations"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.AutomationProperties{
	// 		Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
	// 		Actions: []armsecurity.AutomationActionClassification{
	// 			&armsecurity.AutomationActionLogicApp{
	// 				ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
	// 				LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
	// 		}},
	// 		IsEnabled: to.Ptr(false),
	// 		Scopes: []*armsecurity.AutomationScope{
	// 			{
	// 				Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
	// 				ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
	// 		}},
	// 		Sources: []*armsecurity.AutomationSource{
	// 			{
	// 				EventSource: to.Ptr(armsecurity.EventSourceAssessments),
	// 				RuleSets: []*armsecurity.AutomationRuleSet{
	// 					{
	// 						Rules: []*armsecurity.AutomationTriggeringRule{
	// 							{
	// 								ExpectedValue: to.Ptr("customAssessment"),
	// 								Operator: to.Ptr(armsecurity.OperatorEquals),
	// 								PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
	// 								PropertyType: to.Ptr(armsecurity.PropertyTypeString),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/DeleteAutomation_example.json
func ExampleAutomationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAutomationsClient().Delete(ctx, "myRg", "myAutomationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/ValidateAutomation_example.json
func ExampleAutomationsClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationsClient().Validate(ctx, "exampleResourceGroup", "exampleAutomation", armsecurity.Automation{
		Location: to.Ptr("Central US"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.AutomationProperties{
			Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
			Actions: []armsecurity.AutomationActionClassification{
				&armsecurity.AutomationActionLogicApp{
					ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
					LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
					URI:                to.Ptr("https://exampleTriggerUri1.com"),
				}},
			IsEnabled: to.Ptr(true),
			Scopes: []*armsecurity.AutomationScope{
				{
					Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
					ScopePath:   to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
				}},
			Sources: []*armsecurity.AutomationSource{
				{
					EventSource: to.Ptr(armsecurity.EventSourceAssessments),
					RuleSets: []*armsecurity.AutomationRuleSet{
						{
							Rules: []*armsecurity.AutomationTriggeringRule{
								{
									ExpectedValue: to.Ptr("customAssessment"),
									Operator:      to.Ptr(armsecurity.OperatorEquals),
									PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
									PropertyType:  to.Ptr(armsecurity.PropertyTypeString),
								}},
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AutomationValidationStatus = armsecurity.AutomationValidationStatus{
	// 	IsValid: to.Ptr(true),
	// 	Message: to.Ptr("Validation Successful"),
	// }
}
