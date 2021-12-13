// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBlobClient(t *testing.T) {
	// t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := getContainerClient(containerName, svcClient)

	blobName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(blobName, containerClient)

	blobURLParts := NewBlobURLParts(bbClient.URL())
	require.Equal(t, blobURLParts.BlobName, blobName)
	require.Equal(t, blobURLParts.ContainerName, containerName)

	accountName, err := getRequiredEnv(AccountNameEnvVar)
	require.NoError(t, err)
	correctURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)
	require.Equal(t, bbClient.URL(), correctURL)
}

//nolint
//func (s *azblobUnrecordedTestSuite) TestCreateBlobClientWithSnapshotAndSAS() {
//	_assert := assert.New(s.T())
//	// testName := s.T().Name()
//	_context := getTestContext(s.T().Name())
//	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
//	_assert.NoError(err)
//
//	containerName := generateContainerName(s.T().Name())
//	containerClient := getContainerClient(containerName, svcClient)
//
//	blobName := generateBlobName(s.T().Name())
//	bbClient := getBlockBlobClient(blobName, containerClient)
//
//	currentTime, err := time.Parse(time.UnixDate, "Fri Jun 11 20:00:00 UTC 2049")
//	_assert.NoError(err)
//
//	credential, err := getGenericCredential(nil, testAccountDefault)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//	sasQueryParams, err := AccountSASSignatureValues{
//		Protocol:      SASProtocolHTTPS,
//		ExpiryTime:    currentTime,
//		Permissions:   AccountSASPermissions{Read: true, List: true}.String(),
//		Services:      AccountSASServices{Blob: true}.String(),
//		ResourceTypes: AccountSASResourceTypes{Container: true, Object: true}.String(),
//	}.NewSASQueryParameters(credential)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//
//	parts := NewBlobURLParts(bbClient.URL())
//	parts.SAS = sasQueryParams
//	parts.Snapshot = currentTime.Format(SnapshotTimeFormat)
//	blobURLParts := parts.URL()
//
//	// The snapshot format string is taken from the snapshotTimeFormat value in parsing_urls.go. The field is not public, so
//	// it is copied here
//	accountName, err := getRequiredEnv(AccountNameEnvVar)
//	_assert.NoError(err)
//	correctURL := "https://" + accountName + DefaultBlobEndpointSuffix + containerName + "/" + blobName +
//		"?" + "snapshot=" + currentTime.Format("2006-01-02T15:04:05.0000000Z07:00") + "&" + sasQueryParams.Encode()
//	_assert.Equal(blobURLParts, correctURL)
//}

//nolint
//func (s *azblobUnrecordedTestSuite) TestCreateBlobClientWithSnapshotAndSASUsingConnectionString() {
//	_assert := assert.New(s.T())
//	// testName := s.T().Name()
//	svcClient, err := getServiceClientFromConnectionString(nil, testAccountDefault, nil)
//	_assert.NoError(err)
//
//	containerName := generateContainerName(s.T().Name())
//	containerClient := getContainerClient(containerName, svcClient)
//
//	blobName := generateBlobName(s.T().Name())
//	bbClient := getBlockBlobClient(blobName, containerClient)
//
//	currentTime, err := time.Parse(time.UnixDate, "Fri Jun 11 20:00:00 UTC 2049")
//	_assert.NoError(err)
//
//	credential, err := getGenericCredential(nil, testAccountDefault)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//	sasQueryParams, err := AccountSASSignatureValues{
//		Protocol:      SASProtocolHTTPS,
//		ExpiryTime:    currentTime,
//		Permissions:   AccountSASPermissions{Read: true, List: true}.String(),
//		Services:      AccountSASServices{Blob: true}.String(),
//		ResourceTypes: AccountSASResourceTypes{Container: true, Object: true}.String(),
//	}.NewSASQueryParameters(credential)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//
//	parts := NewBlobURLParts(bbClient.URL())
//	parts.SAS = sasQueryParams
//	parts.Snapshot = currentTime.Format(SnapshotTimeFormat)
//	blobURLParts := parts.URL()
//
//	// The snapshot format string is taken from the snapshotTimeFormat value in parsing_urls.go. The field is not public, so
//	// it is copied here
//	accountName, err := getRequiredEnv(AccountNameEnvVar)
//	_assert.NoError(err)
//	correctURL := "https://" + accountName + DefaultBlobEndpointSuffix + containerName + "/" + blobName +
//		"?" + "snapshot=" + currentTime.Format("2006-01-02T15:04:05.0000000Z07:00") + "&" + sasQueryParams.Encode()
//	_assert.Equal(blobURLParts, correctURL)
//}

func waitForCopy(_assert *assert.Assertions, copyBlobClient BlockBlobClient, blobCopyResponse BlobStartCopyFromURLResponse) {
	status := *blobCopyResponse.CopyStatus
	// Wait for the copy to finish. If the copy takes longer than a minute, we will fail
	start := time.Now()
	for status != CopyStatusTypeSuccess {
		props, _ := copyBlobClient.GetProperties(ctx, nil)
		status = *props.CopyStatus
		currentTime := time.Now()
		if currentTime.Sub(start) >= time.Minute {
			_assert.Fail("If the copy takes longer than a minute, we will fail")
		}
	}
}

func TestBlobStartCopyDestEmpty(t *testing.T) {
	// t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := getContainerClient(containerName, svcClient)

	_, err = containerClient.Create(ctx, nil)
	require.NoError(t, err)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blobName, containerClient)

	anotherBlobName := "copy" + blobName
	copyBlobClient := getBlockBlobClient(anotherBlobName, containerClient)

	blobCopyResponse, err := copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), nil)
	require.NoError(t, err)
	waitForCopy(assert.New(t), copyBlobClient, blobCopyResponse)

	resp, err := copyBlobClient.Download(ctx, nil)
	require.NoError(t, err)

	// Read the blob data to verify the copy
	data, err := ioutil.ReadAll(resp.RawResponse.Body)
	require.NoError(t, err)
	require.Equal(t, *resp.ContentLength, int64(len(blockBlobDefaultData)))
	require.Equal(t, string(data), blockBlobDefaultData)
	_ = resp.Body(RetryReaderOptions{}).Close()
}

func TestBlobStartCopyMetadata(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := getContainerClient(containerName, svcClient)

	_, err = containerClient.Create(ctx, nil)
	require.NoError(t, err)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blobName, containerClient)

	anotherBlobName := "copy" + blobName
	copyBlobClient := getBlockBlobClient(anotherBlobName, containerClient)

	metadata := make(map[string]string)
	metadata["Bla"] = "foo"
	options := StartCopyBlobOptions{
		Metadata: metadata,
	}
	resp, err := copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)
	waitForCopy(assert.New(t), copyBlobClient, resp)

	resp2, err := copyBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
	require.EqualValues(t, resp2.Metadata, metadata)
}

func TestBlobStartCopyMetadataNil(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	anotherBlobName := "copy" + blockBlobName
	copyBlobClient := getBlockBlobClient(anotherBlobName, containerClient)

	// Have the destination start with metadata, so we ensure the nil metadata passed later takes effect
	_, err = copyBlobClient.Upload(ctx, internal.NopCloser(bytes.NewReader([]byte("data"))), nil)
	require.NoError(t, err)

	resp, err := copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), nil)
	require.NoError(t, err)

	waitForCopy(assert.New(t), copyBlobClient, resp)

	resp2, err := copyBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
	require.Len(t, resp2.Metadata, 0)
}

func TestBlobStartCopyMetadataEmpty(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blobName, containerClient)

	anotherBlobName := "copy" + blobName
	copyBlobClient := getBlockBlobClient(anotherBlobName, containerClient)

	// Have the destination start with metadata, so we ensure the empty metadata passed later takes effect
	_, err = copyBlobClient.Upload(ctx, internal.NopCloser(bytes.NewReader([]byte("data"))), nil)
	require.NoError(t, err)

	metadata := make(map[string]string)
	options := StartCopyBlobOptions{
		Metadata: metadata,
	}
	resp, err := copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	waitForCopy(assert.New(t), copyBlobClient, resp)

	resp2, err := copyBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
	require.Len(t, resp2.Metadata, 0)
}

func TestBlobStartCopyMetadataInvalidField(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blobName, containerClient)

	anotherBlobName := "copy" + generateBlobName(t.Name())
	copyBlobClient := getBlockBlobClient(anotherBlobName, containerClient)

	metadata := make(map[string]string)
	metadata["I nvalid."] = "foo"
	options := StartCopyBlobOptions{
		Metadata: metadata,
	}
	_, err = copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
	require.Equal(t, strings.Contains(err.Error(), invalidHeaderErrorSubstring), true)
}

func TestBlobStartCopySourceNonExistent(t *testing.T) {
	// t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(blobName, containerClient)

	anotherBlobName := "copy" + blobName
	copyBlobClient := getBlockBlobClient(anotherBlobName, containerClient)

	_, err = copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), nil)
	require.Error(t, err)
	require.Equal(t, strings.Contains(err.Error(), "not exist"), true)
}

func TestBlobStartCopySourcePrivate(t *testing.T) {
	// t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	_, err = containerClient.SetAccessPolicy(ctx, nil)
	require.NoError(t, err)

	bbClient := createNewBlockBlob(assert.New(t), generateBlobName(t.Name()), containerClient)

	serviceClient2, err := createServiceClientWithSharedKeyForRecording(t, testAccountSecondary)
	if err != nil {
		t.Skip(err.Error())
		return
	}

	copyContainerClient := createNewContainer(assert.New(t), "cpyc"+containerName, serviceClient2)
	defer deleteContainer(assert.New(t), copyContainerClient)
	copyBlobName := "copyb" + generateBlobName(t.Name())
	copyBlobClient := getBlockBlobClient(copyBlobName, copyContainerClient)

	if svcClient.URL() == serviceClient2.URL() {
		t.Skip("Test not valid because primary and secondary accounts are the same")
	}
	_, err = copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), nil)
	validateStorageError(assert.New(t), err, StorageErrorCodeCannotVerifyCopySource)
}

func TestBlobStartCopyUsingSASSrc(t *testing.T) {
	// t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	_, err = containerClient.SetAccessPolicy(ctx, nil)
	require.NoError(t, err)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	// Create sas values for the source blob
	credential, err := getGenericCredential(nil, testAccountDefault)
	if err != nil {
		t.Fatal("Couldn't fetch credential because " + err.Error())
	}

	startTime, err := time.Parse(time.UnixDate, "Fri Jun 11 20:00:00 UTC 2021")
	require.NoError(t, err)
	endTime, err := time.Parse(time.UnixDate, "Fri Jun 11 20:00:00 UTC 2049")
	require.NoError(t, err)
	serviceSASValues := BlobSASSignatureValues{
		StartTime:     startTime,
		ExpiryTime:    endTime,
		Permissions:   BlobSASPermissions{Read: true, Write: true}.String(),
		ContainerName: containerName,
		BlobName:      blockBlobName}
	queryParams, err := serviceSASValues.NewSASQueryParameters(credential)
	if err != nil {
		t.Fatal(err)
	}

	// Create URLs to the destination blob with sas parameters
	sasURL := NewBlobURLParts(bbClient.URL())
	sasURL.SAS = queryParams

	// Create a new container for the destination
	serviceClient2, err := getServiceClient(nil, testAccountSecondary, nil)
	if err != nil {
		t.Skip(err.Error())
	}

	copyContainerName := "copy" + generateContainerName(t.Name())
	copyContainerClient := createNewContainer(assert.New(t), copyContainerName, serviceClient2)
	defer deleteContainer(assert.New(t), copyContainerClient)

	copyBlobName := "copy" + generateBlobName(t.Name())
	copyBlobClient := getBlockBlobClient(copyBlobName, copyContainerClient)

	resp, err := copyBlobClient.StartCopyFromURL(ctx, sasURL.URL(), nil)
	require.NoError(t, err)

	waitForCopy(assert.New(t), copyBlobClient, resp)

	downloadBlobOptions := DownloadBlobOptions{
		Offset: to.Int64Ptr(0),
		Count:  to.Int64Ptr(int64(len(blockBlobDefaultData))),
	}
	resp2, err := copyBlobClient.Download(ctx, &downloadBlobOptions)
	require.NoError(t, err)

	data, err := ioutil.ReadAll(resp2.RawResponse.Body)
	require.NoError(t, err)
	require.Equal(t, *resp2.ContentLength, int64(len(blockBlobDefaultData)))
	require.Equal(t, string(data), blockBlobDefaultData)
	_ = resp2.Body(RetryReaderOptions{}).Close()
}

func TestBlobStartCopyUsingSASDest(t *testing.T) {
	stop := start(t)
	defer stop()

	var svcClient ServiceClient
	var err error
	for i := 1; i <= 2; i++ {
		if i == 1 {
			svcClient, err = createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
			require.NoError(t, err)
		} else {
			svcClient, err = createServiceClientWithConnStrForRecording(t, testAccountDefault)
			require.NoError(t, err)
		}
		require.NoError(t, err)

		containerClient := createNewContainer(assert.New(t), generateContainerName(t.Name())+strconv.Itoa(i), svcClient)
		_, err := containerClient.SetAccessPolicy(ctx, nil)
		require.NoError(t, err)

		blobClient := createNewBlockBlob(assert.New(t), generateBlobName(t.Name()), containerClient)
		_, err = blobClient.Delete(ctx, nil)
		require.NoError(t, err)

		deleteContainer(assert.New(t), containerClient)
	}
}

func TestBlobStartCopySourceIfModifiedSinceTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	bbClient := getBlockBlobClient(generateBlobName(t.Name()), containerClient)
	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)
	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfModifiedSince: &currentTime,
		},
	}

	destBlobClient := getBlockBlobClient("dst"+generateBlobName(t.Name()), containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	_, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopySourceIfModifiedSinceFalse(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(blobName, containerClient)
	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)
	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfModifiedSince: &currentTime,
		},
	}

	destBlobClient := getBlockBlobClient("dst"+blobName, containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
}

func TestBlobStartCopySourceIfUnmodifiedSinceTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(blobName, containerClient)
	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)
	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfUnmodifiedSince: &currentTime,
		},
	}

	destBlobClient := getBlockBlobClient("dst"+generateBlobName(t.Name()), containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	_, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopySourceIfUnmodifiedSinceFalse(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blobName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(blobName, containerClient)
	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)
	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfUnmodifiedSince: &currentTime,
		},
	}
	destBlobClient := getBlockBlobClient("dst"+blobName, containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
}

func  TestBlobStartCopySourceIfMatchTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	require.NoError(t, err)

	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfMatch: resp.ETag,
		},
	}

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := getBlockBlobClient(destBlobName, containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	_, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopySourceIfMatchFalse(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	randomEtag := "a"
	accessConditions := SourceModifiedAccessConditions{
		SourceIfMatch: &randomEtag,
	}
	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &accessConditions,
	}

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := getBlockBlobClient(destBlobName, containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
	validateStorageError(assert.New(t), err, StorageErrorCodeSourceConditionNotMet)
}

func TestBlobStartCopySourceIfNoneMatchTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	_, err = bbClient.GetProperties(ctx, nil)
	require.NoError(t, err)

	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfNoneMatch: to.StringPtr("a"),
		},
	}

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := getBlockBlobClient(destBlobName, containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	_, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopySourceIfNoneMatchFalse(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	require.NoError(t, err)

	options := StartCopyBlobOptions{
		SourceModifiedAccessConditions: &SourceModifiedAccessConditions{
			SourceIfNoneMatch: resp.ETag,
		},
	}

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := getBlockBlobClient(destBlobName, containerClient)
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
	validateStorageError(assert.New(t), err, StorageErrorCodeSourceConditionNotMet)
}

func TestBlobStartCopyDestIfModifiedSinceTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	bbName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfModifiedSince: &currentTime,
		},
	}
	destBlobClient := createNewBlockBlob(assert.New(t), "dst"+bbName, containerClient) // The blob must exist to have a last-modified time
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	_, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopyDestIfModifiedSinceFalse(t *testing.T) {
	// t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	bbName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	destBlobClient := createNewBlockBlob(assert.New(t), "dst"+bbName, containerClient) // The blob must exist to have a last-modified time

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)
	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfModifiedSince: &currentTime,
		},
	}
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	validateStorageError(assert.New(t), err, StorageErrorCodeTargetConditionNotMet)
}

func TestBlobStartCopyDestIfUnmodifiedSinceTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	bbName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	destBlobClient := createNewBlockBlob(assert.New(t), "dst"+bbName, containerClient)

	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfUnmodifiedSince: &currentTime,
		},
	}
	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	_, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopyDestIfUnmodifiedSinceFalse(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	bbName := generateBlobName(t.Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	require.NoError(t, err)
	require.Equal(t, cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	destBlobClient := createNewBlockBlob(assert.New(t), "dst"+bbName, containerClient)
	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfUnmodifiedSince: &currentTime,
		},
	}

	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
}

func TestBlobStartCopyDestIfMatchTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := createNewBlockBlob(assert.New(t), destBlobName, containerClient)
	resp, err := destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)

	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfMatch: resp.ETag,
		},
	}

	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	resp, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func TestBlobStartCopyDestIfMatchFalse(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := createNewBlockBlob(assert.New(t), destBlobName, containerClient)
	resp, err := destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)

	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfMatch: resp.ETag,
		},
	}
	metadata := make(map[string]string)
	metadata["bla"] = "bla"
	_, err = destBlobClient.SetMetadata(ctx, metadata, nil)
	require.NoError(t, err)

	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.Error(t, err)
	validateStorageError(assert.New(t), err, StorageErrorCodeTargetConditionNotMet)
}

func TestBlobStartCopyDestIfNoneMatchTrue(t *testing.T) {
	t.Skip("Error: 'System.InvalidCastException: Unable to cast object of type 'System.Net.Http.EmptyReadStream' to type 'System.IO.MemoryStream'.'")
	stop := start(t)
	defer stop()

	svcClient, err := createServiceClientWithSharedKeyForRecording(t, testAccountDefault)
	require.NoError(t, err)

	containerName := generateContainerName(t.Name())
	containerClient := createNewContainer(assert.New(t), containerName, svcClient)
	defer deleteContainer(assert.New(t), containerClient)

	blockBlobName := generateBlobName(t.Name())
	bbClient := createNewBlockBlob(assert.New(t), blockBlobName, containerClient)

	destBlobName := "dest" + generateBlobName(t.Name())
	destBlobClient := createNewBlockBlob(assert.New(t), destBlobName, containerClient)
	resp, err := destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)

	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfNoneMatch: resp.ETag,
		},
	}

	_, err = destBlobClient.SetMetadata(ctx, nil, nil) // SetMetadata chances the blob's etag
	require.NoError(t, err)

	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	require.NoError(t, err)

	resp, err = destBlobClient.GetProperties(ctx, nil)
	require.NoError(t, err)
}

func (s *azblobTestSuite) TestBlobStartCopyDestIfNoneMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	destBlobName := "dest" + generateBlobName(s.T().Name())
	destBlobClient := createNewBlockBlob(assert.New(s.T()), destBlobName, containerClient)
	resp, err := destBlobClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	options := StartCopyBlobOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfNoneMatch: resp.ETag,
		},
	}

	_, err = destBlobClient.StartCopyFromURL(ctx, bbClient.URL(), &options)
	_assert.Error(err)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeTargetConditionNotMet)
}

//nolint
func (s *azblobUnrecordedTestSuite) TestBlobAbortCopyInProgress() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	svcClient, err := getServiceClient(nil, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(blockBlobName, containerClient)

	// Create a large blob that takes time to copy
	blobSize := 8 * 1024 * 1024
	blobReader, _ := getRandomDataAndReader(blobSize)
	_, err = bbClient.Upload(ctx, internal.NopCloser(blobReader), nil)
	_assert.NoError(err)

	access := PublicAccessTypeBlob
	setAccessPolicyOptions := SetAccessPolicyOptions{
		ContainerSetAccessPolicyOptions: ContainerSetAccessPolicyOptions{
			Access: &access,
		},
	}
	_, err = containerClient.SetAccessPolicy(ctx, &setAccessPolicyOptions) // So that we don't have to create a SAS
	_assert.NoError(err)

	// Must copy across accounts so it takes time to copy
	serviceClient2, err := getServiceClient(nil, testAccountSecondary, nil)
	if err != nil {
		s.T().Skip(err.Error())
	}

	copyContainerName := "copy" + generateContainerName(s.T().Name())
	copyContainerClient := createNewContainer(assert.New(s.T()), copyContainerName, serviceClient2)

	copyBlobName := "copy" + generateBlobName(s.T().Name())
	copyBlobClient := getBlockBlobClient(copyBlobName, copyContainerClient)

	defer deleteContainer(assert.New(s.T()), copyContainerClient)

	resp, err := copyBlobClient.StartCopyFromURL(ctx, bbClient.URL(), nil)
	_assert.NoError(err)
	_assert.Equal(*resp.CopyStatus, CopyStatusTypePending)
	_assert.NotNil(resp.CopyID)

	_, err = copyBlobClient.AbortCopyFromURL(ctx, *resp.CopyID, nil)
	if err != nil {
		// If the error is nil, the test continues as normal.
		// If the error is not nil, we want to check if it's because the copy is finished and send a message indicating this.
		validateStorageError(assert.New(s.T()), err, StorageErrorCodeNoPendingCopyOperation)
		_assert.Error(errors.New("the test failed because the copy completed because it was aborted"))
	}

	resp2, _ := copyBlobClient.GetProperties(ctx, nil)
	_assert.Equal(*resp2.CopyStatus, CopyStatusTypeAborted)
}

func (s *azblobTestSuite) TestBlobAbortCopyNoCopyStarted() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	copyBlobClient := getBlockBlobClient(blockBlobName, containerClient)

	_, err = copyBlobClient.AbortCopyFromURL(ctx, "copynotstarted", nil)
	_assert.Error(err)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeInvalidQueryParameterValue)
}

func (s *azblobTestSuite) TestBlobSnapshotMetadata() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	createBlobSnapshotOptions := CreateBlobSnapshotOptions{
		Metadata: basicMetadata,
	}
	resp, err := bbClient.CreateSnapshot(ctx, &createBlobSnapshotOptions)
	_assert.NoError(err)
	_assert.NotNil(resp.Snapshot)

	// Since metadata is specified on the snapshot, the snapshot should have its own metadata different from the (empty) metadata on the source
	snapshotURL := bbClient.WithSnapshot(*resp.Snapshot)
	resp2, err := snapshotURL.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.EqualValues(resp2.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobSnapshotMetadataEmpty() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	resp, err := bbClient.CreateSnapshot(ctx, nil)
	_assert.NoError(err)
	_assert.NotNil(resp.Snapshot)

	// In this case, because no metadata was specified, it should copy the basicMetadata from the source
	snapshotURL := bbClient.WithSnapshot(*resp.Snapshot)
	resp2, err := snapshotURL.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.EqualValues(resp2.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobSnapshotMetadataNil() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	resp, err := bbClient.CreateSnapshot(ctx, nil)
	_assert.NoError(err)
	_assert.NotNil(resp.Snapshot)

	snapshotURL := bbClient.WithSnapshot(*resp.Snapshot)
	resp2, err := snapshotURL.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.EqualValues(resp2.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobSnapshotMetadataInvalid() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	createBlobSnapshotOptions := CreateBlobSnapshotOptions{
		Metadata: map[string]string{"Invalid Field!": "value"},
	}
	_, err = bbClient.CreateSnapshot(ctx, &createBlobSnapshotOptions)
	_assert.Error(err)
	_assert.Contains(err.Error(), invalidHeaderErrorSubstring)
}

func (s *azblobTestSuite) TestBlobSnapshotBlobNotExist() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(blockBlobName, containerClient)

	_, err = bbClient.CreateSnapshot(ctx, nil)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSnapshotOfSnapshot() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	snapshotString, err := time.Parse(SnapshotTimeFormat, "2021-01-01T01:01:01.0000000Z")
	_assert.NoError(err)
	snapshotURL := bbClient.WithSnapshot(snapshotString.String())
	// The library allows the server to handle the snapshot of snapshot error
	_, err = snapshotURL.CreateSnapshot(ctx, nil)
	_assert.Error(err)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeInvalidQueryParameterValue)
}

func (s *azblobTestSuite) TestBlobSnapshotIfModifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	access := ModifiedAccessConditions{
		IfModifiedSince: &currentTime,
	}
	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &access,
	}
	resp, err := bbClient.CreateSnapshot(ctx, &options)
	_assert.NoError(err)
	_assert.NotEqual(*resp.Snapshot, "") // i.e. The snapshot time is not zero. If the service gives us back a snapshot time, it successfully created a snapshot
}

func (s *azblobTestSuite) TestBlobSnapshotIfModifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	access := ModifiedAccessConditions{
		IfModifiedSince: &currentTime,
	}
	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &access,
	}
	_, err = bbClient.CreateSnapshot(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSnapshotIfUnmodifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)
	access := ModifiedAccessConditions{
		IfUnmodifiedSince: &currentTime,
	}
	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &access,
	}
	resp, err := bbClient.CreateSnapshot(ctx, &options)
	_assert.NoError(err)
	_assert.NotEqual(*resp.Snapshot, "")
}

func (s *azblobTestSuite) TestBlobSnapshotIfUnmodifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)
	access := ModifiedAccessConditions{
		IfUnmodifiedSince: &currentTime,
	}
	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &access,
	}
	_, err = bbClient.CreateSnapshot(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSnapshotIfMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfMatch: resp.ETag,
		},
	}
	resp2, err := bbClient.CreateSnapshot(ctx, &options)
	_assert.NoError(err)
	_assert.NotEqual(*resp2.Snapshot, "")
}

func (s *azblobTestSuite) TestBlobSnapshotIfMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	randomEtag := "garbage"
	access := ModifiedAccessConditions{
		IfMatch: &randomEtag,
	}
	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &access,
	}
	_, err = bbClient.CreateSnapshot(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSnapshotIfNoneMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	randomEtag := "garbage"
	access := ModifiedAccessConditions{
		IfNoneMatch: &randomEtag,
	}
	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &access,
	}
	resp, err := bbClient.CreateSnapshot(ctx, &options)
	_assert.NoError(err)
	_assert.NotEqual(*resp.Snapshot, "")
}

func (s *azblobTestSuite) TestBlobSnapshotIfNoneMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	options := CreateBlobSnapshotOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{
			IfNoneMatch: resp.ETag,
		},
	}
	_, err = bbClient.CreateSnapshot(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataNonExistentBlob() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blobName := generateBlobName(s.T().Name())
	bbClient := containerClient.NewBlobClient(blobName)

	_, err = bbClient.Download(ctx, nil)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataNegativeOffset() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Offset: to.Int64Ptr(-1),
	}
	_, err = bbClient.Download(ctx, &options)
	_assert.NoError(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataOffsetOutOfRange() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Offset: to.Int64Ptr(int64(len(blockBlobDefaultData))),
	}
	_, err = bbClient.Download(ctx, &options)
	_assert.Error(err)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeInvalidRange)
}

func (s *azblobTestSuite) TestBlobDownloadDataCountNegative() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Count: to.Int64Ptr(-2),
	}
	_, err = bbClient.Download(ctx, &options)
	_assert.NoError(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataCountZero() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Count: to.Int64Ptr(0),
	}
	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)

	// Specifying a count of 0 results in the value being ignored
	data, err := ioutil.ReadAll(resp.RawResponse.Body)
	_assert.NoError(err)
	_assert.Equal(string(data), blockBlobDefaultData)
}

func (s *azblobTestSuite) TestBlobDownloadDataCountExact() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	count := int64(len(blockBlobDefaultData))
	options := DownloadBlobOptions{
		Count: &count,
	}
	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)

	data, err := ioutil.ReadAll(resp.RawResponse.Body)
	_assert.NoError(err)
	_assert.Equal(string(data), blockBlobDefaultData)
}

func (s *azblobTestSuite) TestBlobDownloadDataCountOutOfRange() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Count: to.Int64Ptr(int64((len(blockBlobDefaultData)) * 2)),
	}
	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)

	data, err := ioutil.ReadAll(resp.RawResponse.Body)
	_assert.NoError(err)
	_assert.Equal(string(data), blockBlobDefaultData)
}

func (s *azblobTestSuite) TestBlobDownloadDataEmptyRangeStruct() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Count:  to.Int64Ptr(0),
		Offset: to.Int64Ptr(0),
	}
	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)

	data, err := ioutil.ReadAll(resp.RawResponse.Body)
	_assert.NoError(err)
	_assert.Equal(string(data), blockBlobDefaultData)
}

func (s *azblobTestSuite) TestBlobDownloadDataContentMD5() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	options := DownloadBlobOptions{
		Count:              to.Int64Ptr(3),
		Offset:             to.Int64Ptr(10),
		RangeGetContentMD5: to.BoolPtr(true),
	}
	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)
	mdf := md5.Sum([]byte(blockBlobDefaultData)[10:13])
	_assert.Equal(resp.ContentMD5, mdf[:])
}

func (s *azblobTestSuite) TestBlobDownloadDataIfModifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	access := ModifiedAccessConditions{
		IfModifiedSince: &currentTime,
	}
	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{ModifiedAccessConditions: &access},
	}

	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)
	_assert.Equal(*resp.ContentLength, int64(len(blockBlobDefaultData)))
}

func (s *azblobTestSuite) TestBlobDownloadDataIfModifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	access := ModifiedAccessConditions{
		IfModifiedSince: &currentTime,
	}
	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{ModifiedAccessConditions: &access},
	}
	_, err = bbClient.Download(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataIfUnmodifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
		},
	}
	resp, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)
	_assert.Equal(*resp.ContentLength, int64(len(blockBlobDefaultData)))
}

func (s *azblobTestSuite) TestBlobDownloadDataIfUnmodifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)
	access := ModifiedAccessConditions{
		IfUnmodifiedSince: &currentTime,
	}
	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{ModifiedAccessConditions: &access},
	}
	_, err = bbClient.Download(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataIfMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: resp.ETag},
		},
	}
	resp2, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)
	_assert.Equal(*resp2.ContentLength, int64(len(blockBlobDefaultData)))
}

func (s *azblobTestSuite) TestBlobDownloadDataIfMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: resp.ETag},
		},
	}

	_, err = bbClient.SetMetadata(ctx, nil, nil)
	_assert.NoError(err)

	_, err = bbClient.Download(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDownloadDataIfNoneMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	access := ModifiedAccessConditions{
		IfNoneMatch: resp.ETag,
	}
	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{ModifiedAccessConditions: &access},
	}

	_, err = bbClient.SetMetadata(ctx, nil, nil)
	_assert.NoError(err)

	resp2, err := bbClient.Download(ctx, &options)
	_assert.NoError(err)
	_assert.Equal(*resp2.ContentLength, int64(len(blockBlobDefaultData)))
}

func (s *azblobTestSuite) TestBlobDownloadDataIfNoneMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	options := DownloadBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: resp.ETag},
		},
	}

	_, err = bbClient.Download(ctx, &options)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDeleteNonExistent() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := containerClient.NewBlockBlobClient(blockBlobName)

	_, err = bbClient.Delete(ctx, nil)
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobDeleteSnapshot() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.CreateSnapshot(ctx, nil)
	_assert.NoError(err)
	snapshotURL := bbClient.WithSnapshot(*resp.Snapshot)

	_, err = snapshotURL.Delete(ctx, nil)
	_assert.NoError(err)

	validateBlobDeleted(assert.New(s.T()), snapshotURL.BlobClient)
}

//
////func (s *azblobTestSuite) TestBlobDeleteSnapshotsInclude() {
////	svcClient := getServiceClient()
////	containerClient, _ := createNewContainer(c, svcClient)
////	defer deleteContainer(assert.New(s.T()), containerClient)
////	bbClient, _ := createNewBlockBlob(c, containerClient)
////
////	_, err := bbClient.CreateSnapshot(ctx, nil)
////	_assert.NoError(err)
////
////	deleteSnapshots := DeleteSnapshotsOptionInclude
////	_, err = bbClient.Delete(ctx, &DeleteBlobOptions{
////		DeleteSnapshots: &deleteSnapshots,
////	})
////	_assert.NoError(err)
////
////	include := []ListBlobsIncludeItem{ListBlobsIncludeItemSnapshots}
////	containerListBlobFlatSegmentOptions := ContainerListBlobFlatSegmentOptions{
////		Include: include,
////	}
////	blobs, errChan := containerClient.ListBlobsFlat(ctx, 3, 0, &containerListBlobFlatSegmentOptions)
////	_assert(<- errChan, chk.IsNil)
////	_assert(<- blobs, chk.HasLen, 0)
////}
//
////func (s *azblobTestSuite) TestBlobDeleteSnapshotsOnly() {
////	svcClient := getServiceClient()
////	containerClient, _ := createNewContainer(c, svcClient)
////	defer deleteContainer(assert.New(s.T()), containerClient)
////	bbClient, _ := createNewBlockBlob(c, containerClient)
////
////	_, err := bbClient.CreateSnapshot(ctx, nil)
////	_assert.NoError(err)
////	deleteSnapshot := DeleteSnapshotsOptionOnly
////	deleteBlobOptions := DeleteBlobOptions{
////		DeleteSnapshots: &deleteSnapshot,
////	}
////	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
////	_assert.NoError(err)
////
////	include := []ListBlobsIncludeItem{ListBlobsIncludeItemSnapshots}
////	containerListBlobFlatSegmentOptions := ContainerListBlobFlatSegmentOptions{
////		Include: include,
////	}
////	blobs, errChan := containerClient.ListBlobsFlat(ctx, 3, 0, &containerListBlobFlatSegmentOptions)
////	_assert(<- errChan, chk.IsNil)
////	_assert(blobs, chk.HasLen, 1)
////	_assert(*(<-blobs).Snapshot == "", chk.Equals, true)
////}

func (s *azblobTestSuite) TestBlobDeleteSnapshotsNoneWithSnapshots() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.CreateSnapshot(ctx, nil)
	_assert.NoError(err)
	_, err = bbClient.Delete(ctx, nil)
	_assert.Error(err)
}

func validateBlobDeleted(_assert *assert.Assertions, bbClient BlobClient) {
	_, err := bbClient.GetProperties(ctx, nil)
	_assert.Error(err)

	var storageError *StorageError
	_assert.Equal(errors.As(err, &storageError), true)
	_assert.Equal(storageError.ErrorCode, StorageErrorCodeBlobNotFound)
}

func (s *azblobTestSuite) TestBlobDeleteIfModifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	_assert.NoError(err)

	validateBlobDeleted(assert.New(s.T()), bbClient.BlobClient)
}

func (s *azblobTestSuite) TestBlobDeleteIfModifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobDeleteIfUnmodifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	_assert.NoError(err)

	validateBlobDeleted(assert.New(s.T()), bbClient.BlobClient)
}

func (s *azblobTestSuite) TestBlobDeleteIfUnmodifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobDeleteIfMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, _ := bbClient.GetProperties(ctx, nil)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: resp.ETag},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	_assert.NoError(err)

	validateBlobDeleted(assert.New(s.T()), bbClient.BlobClient)
}

func (s *azblobTestSuite) TestBlobDeleteIfMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	etag := resp.ETag

	_, err = bbClient.SetMetadata(ctx, nil, nil)
	_assert.NoError(err)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: etag},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobDeleteIfNoneMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, _ := bbClient.GetProperties(ctx, nil)
	etag := resp.ETag
	_, err = bbClient.SetMetadata(ctx, nil, nil)
	_assert.NoError(err)

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: etag},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	_assert.NoError(err)

	validateBlobDeleted(assert.New(s.T()), bbClient.BlobClient)
}

func (s *azblobTestSuite) TestBlobDeleteIfNoneMatchFalse() {
	// _assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, _ := bbClient.GetProperties(ctx, nil)
	etag := resp.ETag

	deleteBlobOptions := DeleteBlobOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: etag},
		},
	}
	_, err = bbClient.Delete(ctx, &deleteBlobOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfModifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime},
		},
	}
	resp, err := bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.NoError(err)
	_assert.EqualValues(resp.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfModifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime},
		},
	}
	_, err = bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.Error(err)
	var storageError *StorageError
	_assert.Equal(errors.As(err, &storageError), true)
	_assert.Equal(storageError.response.StatusCode, 304) // No service code returned for a HEAD
}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfUnmodifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
		},
	}
	resp, err := bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.NoError(err)
	_assert.EqualValues(resp.Metadata, basicMetadata)
}

//func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfUnmodifiedSinceFalse() {
//	// TODO: Not Working
//	_assert := assert.New(s.T())
//	// testName := s.T().Name()
//	_context := getTestContext(s.T().Name())
//	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
//	if err != nil {
//		s.Fail("Unable to fetch service client because " + err.Error())
//	}
//
//	containerName := generateContainerName(s.T().Name())
//	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
//	defer deleteContainer(assert.New(s.T()), containerClient)
//
//	blockBlobName := generateBlobName(s.T().Name())
//	bbClient := getBlockBlobClient(blockBlobName, containerClient)
//
//	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
//
//	_assert.NoError(err)
//	_assert.Equal(cResp.RawResponse.StatusCode, 201)
//
//	currentTime := getRelativeTimeFromAnchor(cResp.Date,-10)
//
//	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
//	_assert.NoError(err)
//
//	getBlobPropertiesOptions := GetBlobPropertiesOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
//	}
//	_, err = bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
//	_assert.Error(err)
//}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: resp.ETag},
		},
	}
	resp2, err := bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.NoError(err)
	_assert.EqualValues(resp2.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobGetPropsOnMissingBlob() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbClient := containerClient.NewBlobClient("MISSING")

	_, err = bbClient.GetProperties(ctx, nil)
	_assert.Error(err)
	var storageError *StorageError
	_assert.Equal(errors.As(err, &storageError), true)
	_assert.Equal(storageError.response.StatusCode, 404)
	_assert.Equal(storageError.response.Header.Get("x-ms-error-code"), string(StorageErrorCodeBlobNotFound))
}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	eTag := "garbage"
	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: &eTag},
		},
	}
	_, err = bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.Error(err)
	var storageError *StorageError
	_assert.Equal(errors.As(err, &storageError), true)
	_assert.Equal(storageError.response.StatusCode, 412)
}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfNoneMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())

	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetMetadata(ctx, basicMetadata, nil)
	_assert.NoError(err)

	eTag := "garbage"
	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: &eTag},
		},
	}
	resp, err := bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.NoError(err)
	_assert.EqualValues(resp.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobGetPropsAndMetadataIfNoneMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())

	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.SetMetadata(ctx, nil, nil)
	_assert.NoError(err)

	getBlobPropertiesOptions := GetBlobPropertiesOptions{
		BlobAccessConditions: &BlobAccessConditions{
			ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: resp.ETag},
		},
	}
	_, err = bbClient.GetProperties(ctx, &getBlobPropertiesOptions)
	_assert.Error(err)
	var storageError *StorageError
	_assert.Equal(errors.As(err, &storageError), true)
	_assert.Equal(storageError.response.StatusCode, 304)
}

func (s *azblobTestSuite) TestBlobSetPropertiesBasic() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetHTTPHeaders(ctx, basicHeaders, nil)
	_assert.NoError(err)

	resp, _ := bbClient.GetProperties(ctx, nil)
	h := resp.GetHTTPHeaders()
	_assert.EqualValues(h, basicHeaders)
}

func (s *azblobTestSuite) TestBlobSetPropertiesEmptyValue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	contentType := to.StringPtr("my_type")
	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentType: contentType}, nil)
	_assert.NoError(err)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.EqualValues(resp.ContentType, contentType)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{}, nil)
	_assert.NoError(err)

	resp, err = bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.Nil(resp.ContentType)
}

func validatePropertiesSet(_assert *assert.Assertions, bbClient BlockBlobClient, disposition string) {
	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.Equal(*resp.ContentDisposition, disposition)
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfModifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime}})
	_assert.NoError(err)

	validatePropertiesSet(assert.New(s.T()), bbClient, "my_disposition")
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfModifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime}})
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfUnmodifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime}})
	_assert.NoError(err)

	validatePropertiesSet(assert.New(s.T()), bbClient, "my_disposition")
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfUnmodifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime}})
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: resp.ETag}})
	_assert.NoError(err)

	validatePropertiesSet(assert.New(s.T()), bbClient, "my_disposition")
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: to.StringPtr("garbage")}})
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfNoneMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: to.StringPtr("garbage")}})
	_assert.NoError(err)

	validatePropertiesSet(assert.New(s.T()), bbClient, "my_disposition")
}

func (s *azblobTestSuite) TestBlobSetPropertiesIfNoneMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	_, err = bbClient.SetHTTPHeaders(ctx, BlobHTTPHeaders{BlobContentDisposition: to.StringPtr("my_disposition")},
		&SetBlobHTTPHeadersOptions{ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: resp.ETag}})
	_assert.Error(err)
}

func (s *azblobTestSuite) TestBlobSetMetadataNil() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetMetadata(ctx, map[string]string{"not": "nil"}, nil)
	_assert.NoError(err)

	_, err = bbClient.SetMetadata(ctx, nil, nil)
	_assert.NoError(err)

	blobGetResp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.Len(blobGetResp.Metadata, 0)
}

func (s *azblobTestSuite) TestBlobSetMetadataEmpty() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetMetadata(ctx, map[string]string{"not": "nil"}, nil)
	_assert.NoError(err)

	_, err = bbClient.SetMetadata(ctx, map[string]string{}, nil)
	_assert.NoError(err)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.Len(resp.Metadata, 0)
}

func (s *azblobTestSuite) TestBlobSetMetadataInvalidField() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	_, err = bbClient.SetMetadata(ctx, map[string]string{"Invalid field!": "value"}, nil)
	_assert.Error(err)
	_assert.Contains(err.Error(), invalidHeaderErrorSubstring)
	//_assert.Equal(strings.Contains(err.Error(), invalidHeaderErrorSubstring), true)
}

func validateMetadataSet(_assert *assert.Assertions, bbClient BlockBlobClient) {
	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.EqualValues(resp.Metadata, basicMetadata)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfModifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	_assert.NoError(err)

	validateMetadataSet(assert.New(s.T()), bbClient)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfModifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfModifiedSince: &currentTime},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfUnmodifiedSinceTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, 10)

	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	_assert.NoError(err)

	validateMetadataSet(assert.New(s.T()), bbClient)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfUnmodifiedSinceFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	bbName := generateBlobName(s.T().Name())
	bbClient := getBlockBlobClient(bbName, containerClient)

	cResp, err := bbClient.Upload(ctx, internal.NopCloser(strings.NewReader(blockBlobDefaultData)), nil)
	_assert.NoError(err)
	_assert.Equal(cResp.RawResponse.StatusCode, 201)

	currentTime := getRelativeTimeFromAnchor(cResp.Date, -10)

	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfUnmodifiedSince: &currentTime},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: resp.ETag},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	_assert.NoError(err)

	validateMetadataSet(assert.New(s.T()), bbClient)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfMatchFalse() {
	// _assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	eTag := "garbage"
	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfMatch: &eTag},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfNoneMatchTrue() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	eTag := "garbage"
	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: &eTag},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	_assert.NoError(err)

	validateMetadataSet(assert.New(s.T()), bbClient)
}

func (s *azblobTestSuite) TestBlobSetMetadataIfNoneMatchFalse() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)

	setBlobMetadataOptions := SetBlobMetadataOptions{
		ModifiedAccessConditions: &ModifiedAccessConditions{IfNoneMatch: resp.ETag},
	}
	_, err = bbClient.SetMetadata(ctx, basicMetadata, &setBlobMetadataOptions)
	validateStorageError(assert.New(s.T()), err, StorageErrorCodeConditionNotMet)
}

//nolint
func testBlobServiceClientDeleteImpl(_ *assert.Assertions, _ ServiceClient) error {
	//containerClient := createNewContainer(assert.New(s.T()), "gocblobserviceclientdeleteimpl", svcClient)
	//defer deleteContainer(assert.New(s.T()), containerClient)
	//bbClient := createNewBlockBlob(assert.New(s.T()), "goblobserviceclientdeleteimpl", containerClient)
	//
	//_, err := bbClient.Delete(ctx, nil)
	//_assert.NoError(err) // This call will not have errors related to slow update of service properties, so we assert.
	//
	//_, err = bbClient.Undelete(ctx)
	//if err != nil { // We want to give the wrapper method a chance to check if it was an error related to the service properties update.
	//	return err
	//}
	//
	//resp, err := bbClient.GetProperties(ctx, nil)
	//if err != nil {
	//	return errors.New(string(err.(*StorageError).ErrorCode))
	//}
	//_assert.Equal(resp.BlobType, BlobTypeBlockBlob) // We could check any property. This is just to double check it was undeleted.
	return nil
}

func (s *azblobTestSuite) TestBlobServiceClientDelete() {
	// _assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	code := 404
	runTestRequiringServiceProperties(assert.New(s.T()), svcClient, string(rune(code)), enableSoftDelete, testBlobServiceClientDeleteImpl, disableSoftDelete)
}

func setAndCheckBlobTier(_assert *assert.Assertions, bbClient BlobClient, tier AccessTier) {
	_, err := bbClient.SetTier(ctx, tier, nil)
	_assert.NoError(err)

	resp, err := bbClient.GetProperties(ctx, nil)
	_assert.NoError(err)
	_assert.Equal(*resp.AccessTier, string(tier))
}

func (s *azblobTestSuite) TestBlobSetTierAllTiers() {
	// _assert := assert.New(s.T())
	// testName := s.T().Name()
	_context := getTestContext(s.T().Name())
	svcClient, err := getServiceClient(_context.recording, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	setAndCheckBlobTier(assert.New(s.T()), bbClient.BlobClient, AccessTierHot)
	setAndCheckBlobTier(assert.New(s.T()), bbClient.BlobClient, AccessTierCool)
	setAndCheckBlobTier(assert.New(s.T()), bbClient.BlobClient, AccessTierArchive)

	premiumServiceClient, err := getServiceClient(_context.recording, testAccountPremium, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	premContainerName := "prem" + generateContainerName(s.T().Name())
	premContainerClient := createNewContainer(assert.New(s.T()), premContainerName, premiumServiceClient)
	defer deleteContainer(assert.New(s.T()), premContainerClient)

	pbClient := createNewPageBlob(assert.New(s.T()), blockBlobName, premContainerClient)

	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP4)
	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP6)
	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP10)
	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP20)
	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP30)
	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP40)
	setAndCheckBlobTier(assert.New(s.T()), pbClient.BlobClient, AccessTierP50)
}

//
////func (s *azblobTestSuite) TestBlobTierInferred() {
////	svcClient, err := getPremiumserviceClient()
////	if err != nil {
////		c.Skip(err.Error())
////	}
////
////	containerClient, _ := createNewContainer(c, svcClient)
////	defer deleteContainer(assert.New(s.T()), containerClient)
////	bbClient, _ := createNewPageBlob(c, containerClient)
////
////	resp, err := bbClient.GetProperties(ctx, nil)
////	_assert.NoError(err)
////	_assert(resp.AccessTierInferred(), chk.Equals, "true")
////
////	resp2, err := containerClient.ListBlobsFlat(ctx, Marker{}, ListBlobsSegmentOptions{})
////	_assert.NoError(err)
////	_assert(resp2.Segment.BlobItems[0].Properties.AccessTierInferred, chk.NotNil)
////	_assert(resp2.Segment.BlobItems[0].Properties.AccessTier, chk.Not(chk.Equals), "")
////
////	_, err = bbClient.SetTier(ctx, AccessTierP4, LeaseAccessConditions{})
////	_assert.NoError(err)
////
////	resp, err = bbClient.GetProperties(ctx, nil)
////	_assert.NoError(err)
////	_assert(resp.AccessTierInferred(), chk.Equals, "")
////
////	resp2, err = containerClient.ListBlobsFlat(ctx, Marker{}, ListBlobsSegmentOptions{})
////	_assert.NoError(err)
////	_assert(resp2.Segment.BlobItems[0].Properties.AccessTierInferred, chk.IsNil) // AccessTierInferred never returned if false
////}
////
////func (s *azblobTestSuite) TestBlobArchiveStatus() {
////	svcClient, err := getBlobStorageserviceClient()
////	if err != nil {
////		c.Skip(err.Error())
////	}
////
////	containerClient, _ := createNewContainer(c, svcClient)
////	defer deleteContainer(assert.New(s.T()), containerClient)
////	bbClient, _ := createNewBlockBlob(c, containerClient)
////
////	_, err = bbClient.SetTier(ctx, AccessTierArchive, LeaseAccessConditions{})
////	_assert.NoError(err)
////	_, err = bbClient.SetTier(ctx, AccessTierCool, LeaseAccessConditions{})
////	_assert.NoError(err)
////
////	resp, err := bbClient.GetProperties(ctx, nil)
////	_assert.NoError(err)
////	_assert(resp.ArchiveStatus(), chk.Equals, string(ArchiveStatusRehydratePendingToCool))
////
////	resp2, err := containerClient.ListBlobsFlat(ctx, Marker{}, ListBlobsSegmentOptions{})
////	_assert.NoError(err)
////	_assert(resp2.Segment.BlobItems[0].Properties.ArchiveStatus, chk.Equals, ArchiveStatusRehydratePendingToCool)
////
////	// delete first blob
////	_, err = bbClient.Delete(ctx, DeleteSnapshotsOptionNone, nil)
////	_assert.NoError(err)
////
////	bbClient, _ = createNewBlockBlob(c, containerClient)
////
////	_, err = bbClient.SetTier(ctx, AccessTierArchive, LeaseAccessConditions{})
////	_assert.NoError(err)
////	_, err = bbClient.SetTier(ctx, AccessTierHot, LeaseAccessConditions{})
////	_assert.NoError(err)
////
////	resp, err = bbClient.GetProperties(ctx, nil)
////	_assert.NoError(err)
////	_assert(resp.ArchiveStatus(), chk.Equals, string(ArchiveStatusRehydratePendingToHot))
////
////	resp2, err = containerClient.ListBlobsFlat(ctx, Marker{}, ListBlobsSegmentOptions{})
////	_assert.NoError(err)
////	_assert(resp2.Segment.BlobItems[0].Properties.ArchiveStatus, chk.Equals, ArchiveStatusRehydratePendingToHot)
////}
////
////func (s *azblobTestSuite) TestBlobTierInvalidValue() {
////	svcClient, err := getBlobStorageserviceClient()
////	if err != nil {
////		c.Skip(err.Error())
////	}
////
////	containerClient, _ := createNewContainer(c, svcClient)
////	defer deleteContainer(assert.New(s.T()), containerClient)
////	bbClient, _ := createNewBlockBlob(c, containerClient)
////
////	_, err = bbClient.SetTier(ctx, AccessTierType("garbage"), LeaseAccessConditions{})
////	validateStorageError(c, err, StorageErrorCodeInvalidHeaderValue)
////}
////

func (s *azblobTestSuite) TestBlobClientPartsSASQueryTimes() {
	_assert := assert.New(s.T())
	StartTimesInputs := []string{
		"2020-04-20",
		"2020-04-20T07:00Z",
		"2020-04-20T07:15:00Z",
		"2020-04-20T07:30:00.1234567Z",
	}
	StartTimesExpected := []time.Time{
		time.Date(2020, time.April, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2020, time.April, 20, 7, 0, 0, 0, time.UTC),
		time.Date(2020, time.April, 20, 7, 15, 0, 0, time.UTC),
		time.Date(2020, time.April, 20, 7, 30, 0, 123456700, time.UTC),
	}
	ExpiryTimesInputs := []string{
		"2020-04-21",
		"2020-04-20T08:00Z",
		"2020-04-20T08:15:00Z",
		"2020-04-20T08:30:00.2345678Z",
	}
	ExpiryTimesExpected := []time.Time{
		time.Date(2020, time.April, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2020, time.April, 20, 8, 0, 0, 0, time.UTC),
		time.Date(2020, time.April, 20, 8, 15, 0, 0, time.UTC),
		time.Date(2020, time.April, 20, 8, 30, 0, 234567800, time.UTC),
	}

	for i := 0; i < len(StartTimesInputs); i++ {
		urlString :=
			"https://myaccount.blob.core.windows.net/mycontainer/mydirectory/myfile.txt?" +
				"se=" + url.QueryEscape(ExpiryTimesInputs[i]) + "&" +
				"sig=NotASignature&" +
				"sp=r&" +
				"spr=https&" +
				"sr=b&" +
				"st=" + url.QueryEscape(StartTimesInputs[i]) + "&" +
				"sv=2019-10-10"

		parts := NewBlobURLParts(urlString)
		_assert.Equal(parts.Scheme, "https")
		_assert.Equal(parts.Host, "myaccount.blob.core.windows.net")
		_assert.Equal(parts.ContainerName, "mycontainer")
		_assert.Equal(parts.BlobName, "mydirectory/myfile.txt")

		sas := parts.SAS
		_assert.Equal(sas.StartTime(), StartTimesExpected[i])
		_assert.Equal(sas.ExpiryTime(), ExpiryTimesExpected[i])

		_assert.Equal(parts.URL(), urlString)
	}
}

//nolint
func (s *azblobUnrecordedTestSuite) TestDownloadBlockBlobUnexpectedEOF() {
	_assert := assert.New(s.T())
	// testName := s.T().Name()
	svcClient, err := getServiceClient(nil, testAccountDefault, nil)
	if err != nil {
		s.Fail("Unable to fetch service client because " + err.Error())
	}

	containerName := generateContainerName(s.T().Name())
	containerClient := createNewContainer(assert.New(s.T()), containerName, svcClient)
	defer deleteContainer(assert.New(s.T()), containerClient)

	blockBlobName := generateBlobName(s.T().Name())
	bbClient := createNewBlockBlob(assert.New(s.T()), blockBlobName, containerClient)

	resp, err := bbClient.Download(ctx, nil)
	_assert.NoError(err)

	// Verify that we can inject errors first.
	reader := resp.Body(InjectErrorInRetryReaderOptions(errors.New("unrecoverable error")))

	_, err = ioutil.ReadAll(reader)
	_assert.Error(err)
	_assert.Equal(err.Error(), "unrecoverable error")

	// Then inject the retryable error.
	reader = resp.Body(InjectErrorInRetryReaderOptions(io.ErrUnexpectedEOF))

	buf, err := ioutil.ReadAll(reader)
	_assert.NoError(err)
	_assert.EqualValues(buf, []byte(blockBlobDefaultData))
}

//nolint
func InjectErrorInRetryReaderOptions(err error) RetryReaderOptions {
	return RetryReaderOptions{
		MaxRetryRequests:       1,
		doInjectError:          true,
		doInjectErrorRound:     0,
		injectedError:          err,
		NotifyFailedRead:       nil,
		TreatEarlyCloseAsError: false,
		CpkInfo:                nil,
		CpkScopeInfo:           nil,
	}
}
