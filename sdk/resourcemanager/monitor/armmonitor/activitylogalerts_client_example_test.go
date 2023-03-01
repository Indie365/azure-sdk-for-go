//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdate.json
func ExampleActivityLogAlertsClient_CreateOrUpdate_createOrUpdateAnActivityLogAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert rule."),
			Actions: &armmonitor.ActionList{
				ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
					{
						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					}},
			},
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("Administrative"),
						Field:  to.Ptr("category"),
					},
					{
						Equals: to.Ptr("Error"),
						Field:  to.Ptr("level"),
					}},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActivityLogAlertResource = armmonitor.ActivityLogAlertResource{
	// 	Name: to.Ptr("SampleActivityLogAlertRule"),
	// 	Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
	// 	ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmonitor.AlertRuleProperties{
	// 		Description: to.Ptr("Description of sample Activity Log Alert rule."),
	// 		Actions: &armmonitor.ActionList{
	// 			ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
	// 					WebhookProperties: map[string]*string{
	// 						"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 					},
	// 			}},
	// 		},
	// 		Condition: &armmonitor.AlertRuleAllOfCondition{
	// 			AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
	// 				{
	// 					Equals: to.Ptr("Administrative"),
	// 					Field: to.Ptr("Category"),
	// 				},
	// 				{
	// 					Equals: to.Ptr("Error"),
	// 					Field: to.Ptr("Level"),
	// 			}},
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		Scopes: []*string{
	// 			to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithAnyOfCondition.json
func ExampleActivityLogAlertsClient_CreateOrUpdate_createOrUpdateAnActivityLogAlertRuleWithAnyOfCondition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyResourceGroup", "SampleActivityLogAlertRuleWithAnyOfCondition", armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert rule with 'anyOf' condition."),
			Actions: &armmonitor.ActionList{
				ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
					{
						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					}},
			},
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("ServiceHealth"),
						Field:  to.Ptr("category"),
					},
					{
						AnyOf: []*armmonitor.AlertRuleLeafCondition{
							{
								Equals: to.Ptr("Incident"),
								Field:  to.Ptr("properties.incidentType"),
							},
							{
								Equals: to.Ptr("Maintenance"),
								Field:  to.Ptr("properties.incidentType"),
							}},
					}},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActivityLogAlertResource = armmonitor.ActivityLogAlertResource{
	// 	Name: to.Ptr("SampleActivityLogAlertRuleWithAnyOfCondition"),
	// 	Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
	// 	ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRuleWithAnyOfCondition"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmonitor.AlertRuleProperties{
	// 		Description: to.Ptr("Description of sample Activity Log Alert rule with 'anyOf' condition."),
	// 		Actions: &armmonitor.ActionList{
	// 			ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
	// 					WebhookProperties: map[string]*string{
	// 						"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 					},
	// 			}},
	// 		},
	// 		Condition: &armmonitor.AlertRuleAllOfCondition{
	// 			AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
	// 				{
	// 					Equals: to.Ptr("ServiceHealth"),
	// 					Field: to.Ptr("category"),
	// 				},
	// 				{
	// 					AnyOf: []*armmonitor.AlertRuleLeafCondition{
	// 						{
	// 							Equals: to.Ptr("Incident"),
	// 							Field: to.Ptr("properties.incidentType"),
	// 						},
	// 						{
	// 							Equals: to.Ptr("Maintenance"),
	// 							Field: to.Ptr("properties.incidentType"),
	// 					}},
	// 			}},
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		Scopes: []*string{
	// 			to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithContainsAny.json
func ExampleActivityLogAlertsClient_CreateOrUpdate_createOrUpdateAnActivityLogAlertRuleWithContainsAny() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyResourceGroup", "SampleActivityLogAlertRuleWithContainsAny", armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert rule with 'containsAny'."),
			Actions: &armmonitor.ActionList{
				ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
					{
						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					}},
			},
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("ServiceHealth"),
						Field:  to.Ptr("category"),
					},
					{
						ContainsAny: []*string{
							to.Ptr("North Europe"),
							to.Ptr("West Europe")},
						Field: to.Ptr("properties.impactedServices[*].ImpactedRegions[*].RegionName"),
					}},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActivityLogAlertResource = armmonitor.ActivityLogAlertResource{
	// 	Name: to.Ptr("SampleActivityLogAlertRuleWithContainsAny"),
	// 	Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
	// 	ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRuleWithContainsAny"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmonitor.AlertRuleProperties{
	// 		Description: to.Ptr("Description of sample Activity Log Alert rule with 'containsAny'."),
	// 		Actions: &armmonitor.ActionList{
	// 			ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
	// 					WebhookProperties: map[string]*string{
	// 						"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 					},
	// 			}},
	// 		},
	// 		Condition: &armmonitor.AlertRuleAllOfCondition{
	// 			AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
	// 				{
	// 					Equals: to.Ptr("ServiceHealth"),
	// 					Field: to.Ptr("category"),
	// 				},
	// 				{
	// 					ContainsAny: []*string{
	// 						to.Ptr("North Europe"),
	// 						to.Ptr("West Europe")},
	// 						Field: to.Ptr("properties.impactedServices[*].ImpactedRegions[*].RegionName"),
	// 				}},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Scopes: []*string{
	// 				to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Get.json
func ExampleActivityLogAlertsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActivityLogAlertResource = armmonitor.ActivityLogAlertResource{
	// 	Name: to.Ptr("SampleActivityLogAlertRule"),
	// 	Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
	// 	ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmonitor.AlertRuleProperties{
	// 		Description: to.Ptr("Description of sample Activity Log Alert rule."),
	// 		Actions: &armmonitor.ActionList{
	// 			ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
	// 					WebhookProperties: map[string]*string{
	// 						"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 					},
	// 			}},
	// 		},
	// 		Condition: &armmonitor.AlertRuleAllOfCondition{
	// 			AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
	// 				{
	// 					Equals: to.Ptr("Administrative"),
	// 					Field: to.Ptr("category"),
	// 				},
	// 				{
	// 					Equals: to.Ptr("Error"),
	// 					Field: to.Ptr("level"),
	// 			}},
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		Scopes: []*string{
	// 			to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Delete.json
func ExampleActivityLogAlertsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Update.json
func ExampleActivityLogAlertsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", armmonitor.AlertRulePatchObject{
		Properties: &armmonitor.AlertRulePatchProperties{
			Enabled: to.Ptr(false),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActivityLogAlertResource = armmonitor.ActivityLogAlertResource{
	// 	Name: to.Ptr("SampleActivityLogAlertRule"),
	// 	Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
	// 	ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmonitor.AlertRuleProperties{
	// 		Description: to.Ptr("Description of sample Activity Log Alert rule."),
	// 		Actions: &armmonitor.ActionList{
	// 			ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
	// 					WebhookProperties: map[string]*string{
	// 						"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
	// 					},
	// 			}},
	// 		},
	// 		Condition: &armmonitor.AlertRuleAllOfCondition{
	// 			AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
	// 				{
	// 					Equals: to.Ptr("Administrative"),
	// 					Field: to.Ptr("category"),
	// 				},
	// 				{
	// 					Equals: to.Ptr("Error"),
	// 					Field: to.Ptr("level"),
	// 			}},
	// 		},
	// 		Enabled: to.Ptr(false),
	// 		Scopes: []*string{
	// 			to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_ListBySubscriptionId.json
func ExampleActivityLogAlertsClient_NewListBySubscriptionIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionIDPager(nil)
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
		// page.AlertRuleList = armmonitor.AlertRuleList{
		// 	Value: []*armmonitor.ActivityLogAlertResource{
		// 		{
		// 			Name: to.Ptr("SampleActivityLogAlertRule1"),
		// 			Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
		// 			ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup1/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule1"),
		// 			Location: to.Ptr("Global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmonitor.AlertRuleProperties{
		// 				Description: to.Ptr("Description of sample Activity Log Alert rule."),
		// 				Actions: &armmonitor.ActionList{
		// 					ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
		// 						{
		// 							ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup1/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 							WebhookProperties: map[string]*string{
		// 								"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
		// 							},
		// 					}},
		// 				},
		// 				Condition: &armmonitor.AlertRuleAllOfCondition{
		// 					AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
		// 						{
		// 							Equals: to.Ptr("Administrative"),
		// 							Field: to.Ptr("category"),
		// 						},
		// 						{
		// 							Equals: to.Ptr("Error"),
		// 							Field: to.Ptr("level"),
		// 					}},
		// 				},
		// 				Enabled: to.Ptr(true),
		// 				Scopes: []*string{
		// 					to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SampleActivityLogAlertRule2"),
		// 				Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
		// 				ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup2/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule2"),
		// 				Location: to.Ptr("Global"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armmonitor.AlertRuleProperties{
		// 					Description: to.Ptr("Description of sample Activity Log Alert rule."),
		// 					Actions: &armmonitor.ActionList{
		// 						ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
		// 							{
		// 								ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup2/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 								WebhookProperties: map[string]*string{
		// 								},
		// 						}},
		// 					},
		// 					Condition: &armmonitor.AlertRuleAllOfCondition{
		// 						AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
		// 							{
		// 								Equals: to.Ptr("Administrative"),
		// 								Field: to.Ptr("category"),
		// 							},
		// 							{
		// 								Equals: to.Ptr("Succeeded"),
		// 								Field: to.Ptr("status"),
		// 						}},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Scopes: []*string{
		// 						to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup2")},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_ListByResourceGroupName.json
func ExampleActivityLogAlertsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page.AlertRuleList = armmonitor.AlertRuleList{
		// 	Value: []*armmonitor.ActivityLogAlertResource{
		// 		{
		// 			Name: to.Ptr("SampleActivityLogAlertRule1"),
		// 			Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
		// 			ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule1"),
		// 			Location: to.Ptr("Global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmonitor.AlertRuleProperties{
		// 				Description: to.Ptr("Description of sample Activity Log Alert rule."),
		// 				Actions: &armmonitor.ActionList{
		// 					ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
		// 						{
		// 							ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 							WebhookProperties: map[string]*string{
		// 								"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
		// 							},
		// 					}},
		// 				},
		// 				Condition: &armmonitor.AlertRuleAllOfCondition{
		// 					AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
		// 						{
		// 							Equals: to.Ptr("Administrative"),
		// 							Field: to.Ptr("category"),
		// 						},
		// 						{
		// 							Equals: to.Ptr("Error"),
		// 							Field: to.Ptr("level"),
		// 					}},
		// 				},
		// 				Enabled: to.Ptr(true),
		// 				Scopes: []*string{
		// 					to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SampleActivityLogAlertRule2"),
		// 				Type: to.Ptr("Microsoft.Insights/ActivityLogAlerts"),
		// 				ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/activityLogAlerts/SampleActivityLogAlertRule2"),
		// 				Location: to.Ptr("Global"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armmonitor.AlertRuleProperties{
		// 					Description: to.Ptr("Description of sample Activity Log Alert rule."),
		// 					Actions: &armmonitor.ActionList{
		// 						ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
		// 							{
		// 								ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 								WebhookProperties: map[string]*string{
		// 								},
		// 						}},
		// 					},
		// 					Condition: &armmonitor.AlertRuleAllOfCondition{
		// 						AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
		// 							{
		// 								Equals: to.Ptr("Administrative"),
		// 								Field: to.Ptr("category"),
		// 							},
		// 							{
		// 								Equals: to.Ptr("Succeeded"),
		// 								Field: to.Ptr("status"),
		// 						}},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Scopes: []*string{
		// 						to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup")},
		// 					},
		// 			}},
		// 		}
	}
}
