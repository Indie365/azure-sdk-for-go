// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/perf"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/spf13/pflag"
)

type uploadPerfTest struct {
	perf.PerfTestOptions
	containerName string
	blobName      string
	blobClient    azblob.BlockBlobClient
	data          io.ReadSeekCloser
}

func (m *uploadPerfTest) GlobalSetup(ctx context.Context) error {
	connStr, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		return fmt.Errorf("the environment variable 'AZURE_STORAGE_CONNECTION_STRING' could not be found")
	}

	containerClient, err := azblob.NewContainerClientFromConnectionString(connStr, m.containerName, nil)
	if err != nil {
		fmt.Println("Error creating the container client: ")
		return err
	}
	_, err = containerClient.Create(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error creating the container: '%s'\n", m.containerName)
		return err
	}

	return nil
}

func (m *uploadPerfTest) Setup(ctx context.Context) error {
	connStr, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		return fmt.Errorf("the environment variable 'AZURE_STORAGE_CONNECTION_STRING' could not be found")
	}

	containerClient, err := azblob.NewContainerClientFromConnectionString(connStr, m.containerName, nil)
	if err != nil {
		return err
	}

	m.blobClient = containerClient.NewBlockBlobClient(m.blobName)
	return nil
}

func (m *uploadPerfTest) Run(ctx context.Context) error {
	fmt.Println(1)
	_, err := m.blobClient.Upload(ctx, m.data, &azblob.UploadBlockBlobOptions{})
	fmt.Println(2)
	return err
}

func (m *uploadPerfTest) Cleanup(ctx context.Context) error {
	return nil
}

func (m *uploadPerfTest) GlobalCleanup(ctx context.Context) error {
	connStr, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		return fmt.Errorf("the environment variable 'AZURE_STORAGE_CONNECTION_STRING' could not be found")
	}

	containerClient, err := azblob.NewContainerClientFromConnectionString(connStr, m.containerName, nil)
	if err != nil {
		return err
	}

	_, err = containerClient.Delete(context.Background(), nil)
	return err
}

func (m *uploadPerfTest) GetMetadata() perf.PerfTestOptions {
	return m.PerfTestOptions
}

func (*uploadPerfTest) RegisterArguments() error {
	count = pflag.Int64("num-blobs", 100, "Number of blobs to list. Defaults to 100.")
	size = pflag.Int64("size", 10240, "Size in bytes of data to be transferred in upload or download tests. Default is 10240.")

	return nil
}

func NewUploadTest(options *perf.PerfTestOptions) perf.PerfTest {
	if options == nil {
		options = &perf.PerfTestOptions{}
	}
	options.Name = "BlobUploadTest"

	if size == nil {
		size = to.Int64Ptr(10240)
	}

	if count == nil {
		count = to.Int64Ptr(100)
	}
	data, err := perf.NewRandomStream(int(*size))
	if err != nil {
		panic(err)
	}
	return &uploadPerfTest{
		PerfTestOptions: *options,
		blobName:        "uploadtest",
		containerName:   "uploadcontainer",
		data:            data,
		// data:            "This is all placeholder random data for now. This is all placeholder random data for now. This is all placeholder random data for now.",
	}
}
