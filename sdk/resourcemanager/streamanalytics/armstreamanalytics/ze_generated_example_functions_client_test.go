//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Create_JavaScript.json
func ExampleFunctionsClient_CreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrReplace(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<function-name>",
		armstreamanalytics.Function{
			Properties: &armstreamanalytics.ScalarFunctionProperties{
				Type: to.Ptr("<type>"),
				Properties: &armstreamanalytics.FunctionConfiguration{
					Binding: &armstreamanalytics.JavaScriptFunctionBinding{
						Type: to.Ptr("<type>"),
						Properties: &armstreamanalytics.JavaScriptFunctionBindingProperties{
							Script: to.Ptr("<script>"),
						},
					},
					Inputs: []*armstreamanalytics.FunctionInput{
						{
							DataType: to.Ptr("<data-type>"),
						}},
					Output: &armstreamanalytics.FunctionOutput{
						DataType: to.Ptr("<data-type>"),
					},
				},
			},
		},
		&armstreamanalytics.FunctionsClientCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Update_JavaScript.json
func ExampleFunctionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<function-name>",
		armstreamanalytics.Function{
			Properties: &armstreamanalytics.ScalarFunctionProperties{
				Type: to.Ptr("<type>"),
				Properties: &armstreamanalytics.FunctionConfiguration{
					Binding: &armstreamanalytics.JavaScriptFunctionBinding{
						Type: to.Ptr("<type>"),
						Properties: &armstreamanalytics.JavaScriptFunctionBindingProperties{
							Script: to.Ptr("<script>"),
						},
					},
				},
			},
		},
		&armstreamanalytics.FunctionsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Delete.json
func ExampleFunctionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<function-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Get_JavaScript.json
func ExampleFunctionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<function-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_ListByStreamingJob.json
func ExampleFunctionsClient_NewListByStreamingJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByStreamingJobPager("<resource-group-name>",
		"<job-name>",
		&armstreamanalytics.FunctionsClientListByStreamingJobOptions{Select: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Test_JavaScript.json
func ExampleFunctionsClient_BeginTest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginTest(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<function-name>",
		&armstreamanalytics.FunctionsClientBeginTestOptions{Function: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_RetrieveDefaultDefinition_AzureML.json
func ExampleFunctionsClient_RetrieveDefaultDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.RetrieveDefaultDefinition(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<function-name>",
		&armstreamanalytics.FunctionsClientRetrieveDefaultDefinitionOptions{FunctionRetrieveDefaultDefinitionParameters: &armstreamanalytics.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters{
			BindingType: to.Ptr("<binding-type>"),
			BindingRetrievalProperties: &armstreamanalytics.AzureMachineLearningWebServiceFunctionBindingRetrievalProperties{
				ExecuteEndpoint: to.Ptr("<execute-endpoint>"),
				UdfType:         to.Ptr("<udf-type>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
