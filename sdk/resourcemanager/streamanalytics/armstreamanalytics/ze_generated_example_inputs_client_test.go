//go:build go1.16
// +build go1.16

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

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Reference_Blob_CSV.json
func ExampleInputsClient_CreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrReplace(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		armstreamanalytics.Input{
			Properties: &armstreamanalytics.ReferenceInputProperties{
				Type: to.StringPtr("<type>"),
				Serialization: &armstreamanalytics.CSVSerialization{
					Type: armstreamanalytics.EventSerializationType("Csv").ToPtr(),
					Properties: &armstreamanalytics.CSVSerializationProperties{
						Encoding:       armstreamanalytics.Encoding("UTF8").ToPtr(),
						FieldDelimiter: to.StringPtr("<field-delimiter>"),
					},
				},
				Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
					Type: to.StringPtr("<type>"),
					Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
						Container:   to.StringPtr("<container>"),
						DateFormat:  to.StringPtr("<date-format>"),
						PathPattern: to.StringPtr("<path-pattern>"),
						StorageAccounts: []*armstreamanalytics.StorageAccount{
							{
								AccountKey:  to.StringPtr("<account-key>"),
								AccountName: to.StringPtr("<account-name>"),
							}},
						TimeFormat: to.StringPtr("<time-format>"),
					},
				},
			},
		},
		&armstreamanalytics.InputsClientCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.InputsClientCreateOrReplaceResult)
}

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Reference_Blob.json
func ExampleInputsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		armstreamanalytics.Input{
			Properties: &armstreamanalytics.ReferenceInputProperties{
				Type: to.StringPtr("<type>"),
				Serialization: &armstreamanalytics.CSVSerialization{
					Type: armstreamanalytics.EventSerializationType("Csv").ToPtr(),
					Properties: &armstreamanalytics.CSVSerializationProperties{
						Encoding:       armstreamanalytics.Encoding("UTF8").ToPtr(),
						FieldDelimiter: to.StringPtr("<field-delimiter>"),
					},
				},
				Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
					Type: to.StringPtr("<type>"),
					Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
						Container: to.StringPtr("<container>"),
					},
				},
			},
		},
		&armstreamanalytics.InputsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.InputsClientUpdateResult)
}

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Delete.json
func ExampleInputsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Get_Reference_Blob_CSV.json
func ExampleInputsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.InputsClientGetResult)
}

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_ListByStreamingJob.json
func ExampleInputsClient_ListByStreamingJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	pager := client.ListByStreamingJob("<resource-group-name>",
		"<job-name>",
		&armstreamanalytics.InputsClientListByStreamingJobOptions{Select: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Test.json
func ExampleInputsClient_BeginTest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginTest(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		&armstreamanalytics.InputsClientBeginTestOptions{Input: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.InputsClientTestResult)
}
