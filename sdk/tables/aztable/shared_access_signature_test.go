// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztable

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

func TestSASServiceClient(t *testing.T) {
	recording.LiveOnly(t)
	accountName := os.Getenv("TABLES_PRIMARY_ACCOUNT_NAME")
	accountKey := os.Getenv("TABLES_PRIMARY_STORAGE_ACCOUNT_KEY")
	cred, err := NewSharedKeyCredential(accountName, accountKey)
	require.NoError(t, err)

	serviceClient, err := NewTableServiceClient(fmt.Sprintf("https://%s.table.core.windows.net/", accountName), cred, nil)
	require.NoError(t, err)

	tableName, err := createRandomName(t, "tableName")
	require.NoError(t, err)

	delete := func() {
		_, err := serviceClient.DeleteTable(context.Background(), tableName, nil)
		require.NoError(t, err)
	}
	defer delete()

	_, err = serviceClient.CreateTable(context.Background(), tableName)
	require.NoError(t, err)

	resources := AccountSASResourceTypes{
		Object:    true,
		Service:   true,
		Container: true,
	}
	permissions := AccountSASPermissions{
		Read:  true,
		Add:   true,
		Write: true,
	}
	start := time.Date(2021, time.August, 4, 1, 1, 0, 0, time.UTC)
	expiry := time.Date(2022, time.August, 4, 1, 1, 0, 0, time.UTC)

	accountSAS, err := serviceClient.GetAccountSASToken(resources, permissions, start, expiry)
	require.NoError(t, err)

	queryParams := accountSAS.Encode()

	sasUrl := fmt.Sprintf("https://%s.table.core.windows.net/?%s", accountName, queryParams)

	err = recording.StartRecording(t, pathToPackage, nil)
	require.NoError(t, err)
	client, err := createTableClientForRecording(t, tableName, sasUrl, azcore.NewAnonymousCredential())
	require.NoError(t, err)
	defer recording.StopRecording(t, nil)  //nolint

	entity := map[string]string{
		"PartitionKey": "pk001",
		"RowKey":       "rk001",
		"Value":        "5",
	}
	marshalled, err := json.Marshal(entity)
	require.NoError(t, err)

	_, err = client.AddEntity(context.Background(), marshalled)
	require.NoError(t, err)
}

func TestSASTableClient(t *testing.T) {
	recording.LiveOnly(t)
	accountName := os.Getenv("TABLES_PRIMARY_ACCOUNT_NAME")
	accountKey := os.Getenv("TABLES_PRIMARY_STORAGE_ACCOUNT_KEY")
	cred, err := NewSharedKeyCredential(accountName, accountKey)
	require.NoError(t, err)

	serviceClient, err := NewTableServiceClient(fmt.Sprintf("https://%s.table.core.windows.net/", accountName), cred, nil)
	require.NoError(t, err)

	tableName, err := createRandomName(t, "tableName")
	require.NoError(t, err)
	tableName = "tablename"

	delete := func() {
		_, err := serviceClient.DeleteTable(context.Background(), tableName, nil)
		require.NoError(t, err)
	}
	defer delete()

	_, err = serviceClient.CreateTable(context.Background(), tableName)
	require.NoError(t, err)

	permissions := TableSASPermissions{
		Read: true,
		Add:  true,
	}
	start := time.Date(2021, time.August, 4, 1, 1, 0, 0, time.UTC)
	expiry := time.Date(2022, time.August, 4, 1, 1, 0, 0, time.UTC)

	accountSAS, err := serviceClient.GetTableSASToken(tableName, permissions, start, expiry)
	require.NoError(t, err)

	queryParams := accountSAS.Encode()

	sasUrl := fmt.Sprintf("https://%s.table.core.windows.net/?%s", accountName, queryParams)

	err = recording.StartRecording(t, pathToPackage, nil)
	require.NoError(t, err)
	client, err := createTableClientForRecording(t, tableName, sasUrl, azcore.NewAnonymousCredential())
	require.NoError(t, err)
	defer recording.StopRecording(t, nil)  //nolint

	entity := map[string]string{
		"PartitionKey": "pk001",
		"RowKey":       "rk001",
		"Value":        "5",
	}
	marshalled, err := json.Marshal(entity)
	require.NoError(t, err)

	_, err = client.AddEntity(context.Background(), marshalled)
	require.NoError(t, err)
}

func TestSASTableClientReadOnly(t *testing.T) {
	recording.LiveOnly(t)
	accountName := os.Getenv("TABLES_PRIMARY_ACCOUNT_NAME")
	accountKey := os.Getenv("TABLES_PRIMARY_STORAGE_ACCOUNT_KEY")
	cred, err := NewSharedKeyCredential(accountName, accountKey)
	require.NoError(t, err)

	serviceClient, err := NewTableServiceClient(fmt.Sprintf("https://%s.table.core.windows.net/", accountName), cred, nil)
	require.NoError(t, err)

	tableName, err := createRandomName(t, "tableName")
	require.NoError(t, err)

	delete := func() {
		_, err := serviceClient.DeleteTable(context.Background(), tableName, nil)
		require.NoError(t, err)
	}
	defer delete()

	_, err = serviceClient.CreateTable(context.Background(), tableName)
	require.NoError(t, err)

	client := serviceClient.NewTableClient(tableName)
	err = insertNEntities("pk001", 4, client)
	require.NoError(t, err)

	permissions := TableSASPermissions{
		Read: true,
	}
	start := time.Date(2021, time.August, 4, 1, 1, 0, 0, time.UTC)
	expiry := time.Date(2022, time.August, 4, 1, 1, 0, 0, time.UTC)

	accountSAS, err := serviceClient.GetTableSASToken(tableName, permissions, start, expiry)
	require.NoError(t, err)

	queryParams := accountSAS.Encode()

	sasUrl := fmt.Sprintf("https://%s.table.core.windows.net/?%s", accountName, queryParams)

	err = recording.StartRecording(t, pathToPackage, nil)
	require.NoError(t, err)
	client, err = createTableClientForRecording(t, tableName, sasUrl, azcore.NewAnonymousCredential())
	require.NoError(t, err)
	defer recording.StopRecording(t, nil)  //nolint

	entity := map[string]string{
		"PartitionKey": "pk001",
		"RowKey":       "rk001",
		"Value":        "5",
	}
	marshalled, err := json.Marshal(entity)
	require.NoError(t, err)

	// Failure on a read
	_, err = client.AddEntity(context.Background(), marshalled)
	require.Error(t, err)

	// Success on a list
	pager := client.List(nil)
	count := 0
	for pager.NextPage(context.Background()) {
		count += len(pager.PageResponse().TableEntityQueryResponse.Value)
	}

	require.Equal(t, 4, count)
}

func TestSASCosmosTableClientReadOnly(t *testing.T) {
	recording.LiveOnly(t)
	accountName := os.Getenv("TABLES_COSMOS_ACCOUNT_NAME")
	accountKey := os.Getenv("TABLES_PRIMARY_COSMOS_ACCOUNT_KEY")
	cred, err := NewSharedKeyCredential(accountName, accountKey)
	require.NoError(t, err)

	serviceClient, err := NewTableServiceClient(fmt.Sprintf("https://%s.table.cosmos.azure.com/", accountName), cred, nil)
	require.NoError(t, err)

	tableName, err := createRandomName(t, "tableName")
	require.NoError(t, err)

	delete := func() {
		_, err := serviceClient.DeleteTable(context.Background(), tableName, nil)
		require.NoError(t, err)
	}
	defer delete()

	_, err = serviceClient.CreateTable(context.Background(), tableName)
	require.NoError(t, err)

	client := serviceClient.NewTableClient(tableName)
	err = insertNEntities("pk001", 4, client)
	require.NoError(t, err)

	permissions := TableSASPermissions{
		Read: true,
	}
	start := time.Date(2021, time.August, 4, 1, 1, 0, 0, time.UTC)
	expiry := time.Date(2022, time.August, 4, 1, 1, 0, 0, time.UTC)
	accountSAS, err := serviceClient.GetTableSASToken(tableName, permissions, start, expiry)
	require.NoError(t, err)

	queryParams := accountSAS.Encode()

	sasUrl := fmt.Sprintf("https://%s.table.cosmos.azure.com/?%s", accountName, queryParams)

	err = recording.StartRecording(t, pathToPackage, nil)
	require.NoError(t, err)
	client, err = createTableClientForRecording(t, tableName, sasUrl, azcore.NewAnonymousCredential())
	require.NoError(t, err)
	defer recording.StopRecording(t, nil)  //nolint

	entity := map[string]string{
		"PartitionKey": "pk001",
		"RowKey":       "rk001",
		"Value":        "5",
	}
	marshalled, err := json.Marshal(entity)
	require.NoError(t, err)

	// Failure on a read
	_, err = client.AddEntity(context.Background(), marshalled)
	require.Error(t, err)

	// Success on a list
	pager := client.List(nil)
	count := 0
	for pager.NextPage(context.Background()) {
		count += len(pager.PageResponse().TableEntityQueryResponse.Value)
	}

	require.Equal(t, 4, count)
}
