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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-02-01-preview/examples/HealthReports/GetHealthReport_example.json
func ExampleHealthReportClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHealthReportClient().Get(ctx, "subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2/resourcegroups/E2E-IBB0WX/providers/Microsoft.Security/securityconnectors/AwsConnectorAllOfferings", "909c629a-bf39-4521-8e4f-10b443a0bc02", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HealthReport = armsecurity.HealthReport{
	// 	Name: to.Ptr("909c629a-bf39-4521-8e4f-10b443a0bc02"),
	// 	Type: to.Ptr("Microsoft.Security/healthReports"),
	// 	ID: to.Ptr("/subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2/resourcegroups/E2E-IBB0WX/providers/Microsoft.Security/securityconnectors/AwsConnectorAllOfferings/providers/Microsoft.Security/healthReports/909c629a-bf39-4521-8e4f-10b443a0bc02"),
	// 	Properties: &armsecurity.HealthReportProperties{
	// 		AffectedDefendersPlans: []*string{
	// 		},
	// 		EnvironmentDetails: &armsecurity.EnvironmentDetails{
	// 			EnvironmentHierarchyID: to.Ptr("a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2"),
	// 			NativeResourceID: to.Ptr("arn:aws:iam::827098768879"),
	// 			OrganizationalHierarchyID: to.Ptr("e81b978c-11be-449f-a392-42c0ed96bb91"),
	// 			SubscriptionID: to.Ptr("a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2"),
	// 			TenantID: to.Ptr("a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2"),
	// 		},
	// 		HealthDataClassification: &armsecurity.HealthDataClassification{
	// 			Component: to.Ptr("Connectivity"),
	// 			Scope: to.Ptr(armsecurity.ScopeNameConnectors),
	// 		},
	// 		Issues: []*armsecurity.Issue{
	// 			{
	// 				IssueAdditionalData: map[string]*string{
	// 					"StacksetName": to.Ptr("ProdStackSet"),
	// 				},
	// 				IssueDescription: to.Ptr("A problem was identified with the AWS CloudFormation StackSet. The StackSet is used to create stacks across multiple accounts. To grant Defender for Cloud access to your member accounts, there is a need to run the StackSet on the member accounts."),
	// 				IssueKey: to.Ptr("414af15d-207e-4c63-b8eb-624d1b652e45"),
	// 				IssueName: to.Ptr("AWS CloudFormation StackSet name invalid or does not exist"),
	// 				RemediationScript: to.Ptr(""),
	// 				RemediationSteps: to.Ptr("Validate that the StackSet name in AWS matches the name provided in the onboarding set up: StackSet name can be found in AWS Management Console  -> CloudFormation -> StackSets -> StackSet name In case the names do not match, update the StackSet name to match the StackSet name provided in the onboarding set up.  In case the StackSet does not exist, re-run the CloudFormation template only for StackSet. Navigate to CloudFormation 'StackSets' in AWS Management Console -> Click 'Create StackSet' -> Choose 'Upload a template file', `Choose file` and select the downloaded template. Make sure to enter the exact StackSet name as it was provided in the onboarding set up. Download template link "),
	// 				SecurityValues: []*string{
	// 					to.Ptr("Connectivity to AWS member accounts")},
	// 			}},
	// 			ResourceDetails: &armsecurity.ResourceDetailsAutoGenerated{
	// 				ConnectorID: to.Ptr("bb7ad9cc-26b6-48ec-a5b4-23fc23be2733"),
	// 				ID: to.Ptr("/subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2/resourcegroups/E2E-IBB0WX/providers/Microsoft.Security/securityconnectors/AwsConnectorAllOfferings"),
	// 				Source: to.Ptr(armsecurity.Source("Aws")),
	// 			},
	// 			Status: &armsecurity.StatusAutoGenerated{
	// 				Code: to.Ptr(armsecurity.StatusNameNotHealthy),
	// 				FirstEvaluationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T09:07:18.6759138Z"); return t}()),
	// 				StatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T09:07:18.6759138Z"); return t}()),
	// 			},
	// 		},
	// 	}
}
