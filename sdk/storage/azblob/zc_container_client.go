// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/uuid"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// A ContainerClient represents a URL to the Azure Storage container allowing you to manipulate its blobs.
type ContainerClient struct {
	client *containerClient
	cred   azcore.Credential
}

// URL returns the URL endpoint used by the ContainerClient object.
func (c ContainerClient) URL() string {
	return c.client.con.u
}

// NewContainerClient creates a ContainerClient object using the specified URL and request policy pipeline.
func NewContainerClient(containerURL string, cred azcore.Credential, options *ClientOptions) (ContainerClient, error) {
	return ContainerClient{client: &containerClient{
		con: newConnection(containerURL, cred, options.getConnectionOptions()),
	}, cred: cred}, nil
}

// NewContainerClientFromConnectionString creates a ContainerClient object using connection string of an account
func NewContainerClientFromConnectionString(connectionString string, containerName string, options *ClientOptions) (ContainerClient, error) {
	svcClient, err := NewServiceClientFromConnectionString(connectionString, options)
	if err != nil {
		return ContainerClient{}, err
	}
	return svcClient.NewContainerClient(containerName), nil
}

// NewBlobClient creates a new BlobClient object by concatenating blobName to the end of
// ContainerClient's URL. The new BlobClient uses the same request policy pipeline as the ContainerClient.
// To change the pipeline, create the BlobClient and then call its WithPipeline method passing in the
// desired pipeline object. Or, call this package's NewBlobClient instead of calling this object's
// NewBlobClient method.
func (c ContainerClient) NewBlobClient(blobName string) BlobClient {
	blobURL := appendToURLPath(c.URL(), blobName)
	newCon := &connection{u: blobURL, p: c.client.con.p}

	return BlobClient{
		client: &blobClient{newCon, nil},
	}
}

// NewAppendBlobClient creates a new AppendBlobURL object by concatenating blobName to the end of
// ContainerClient's URL. The new AppendBlobURL uses the same request policy pipeline as the ContainerClient.
// To change the pipeline, create the AppendBlobURL and then call its WithPipeline method passing in the
// desired pipeline object. Or, call this package's NewAppendBlobClient instead of calling this object's
// NewAppendBlobClient method.
func (c ContainerClient) NewAppendBlobClient(blobName string) AppendBlobClient {
	blobURL := appendToURLPath(c.URL(), blobName)
	newCon := &connection{blobURL, c.client.con.p}

	return AppendBlobClient{
		client:     &appendBlobClient{newCon},
		BlobClient: BlobClient{client: &blobClient{con: newCon}},
	}
}

// NewBlockBlobClient creates a new BlockBlobClient object by concatenating blobName to the end of
// ContainerClient's URL. The new BlockBlobClient uses the same request policy pipeline as the ContainerClient.
// To change the pipeline, create the BlockBlobClient and then call its WithPipeline method passing in the
// desired pipeline object. Or, call this package's NewBlockBlobClient instead of calling this object's
// NewBlockBlobClient method.
func (c ContainerClient) NewBlockBlobClient(blobName string) BlockBlobClient {
	blobURL := appendToURLPath(c.URL(), blobName)
	newCon := &connection{blobURL, c.client.con.p}

	return BlockBlobClient{
		client:     &blockBlobClient{newCon},
		BlobClient: BlobClient{client: &blobClient{con: newCon}},
	}
}

// NewPageBlobClient creates a new PageBlobURL object by concatenating blobName to the end of ContainerClient's URL. The new PageBlobURL uses the same request policy pipeline as the ContainerClient.
// To change the pipeline, create the PageBlobURL and then call its WithPipeline method passing in the
// desired pipeline object. Or, call this package's NewPageBlobClient instead of calling this object's
// NewPageBlobClient method.
func (c ContainerClient) NewPageBlobClient(blobName string) PageBlobClient {
	blobURL := appendToURLPath(c.URL(), blobName)
	newCon := &connection{blobURL, c.client.con.p}

	return PageBlobClient{
		client:     &pageBlobClient{newCon},
		BlobClient: BlobClient{client: &blobClient{con: newCon}},
	}
}

// DeleteBlob marks the specified blob or snapshot within a container for deletion.
// The blob is later deleted during garbage collection.
// Note that deleting a blob also deletes all its snapshots.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/delete-blob.
//nolint
func (c ContainerClient) DeleteBlob(ctx context.Context, blobName string, options *DeleteBlobOptions) (BlobDeleteResponse, error) {
	blobClient := c.NewBlobClient(blobName)
	deleteResp, err := blobClient.Delete(ctx, options)
	return deleteResp, err
}

func (c ContainerClient) NewContainerLeaseClient(leaseID *string) (ContainerLeaseClient, error) {
	if leaseID == nil {
		generatedUuid, err := uuid.New()
		if err != nil {
			return ContainerLeaseClient{}, err
		}
		leaseID = to.StringPtr(generatedUuid.String())
	}
	return ContainerLeaseClient{
		ContainerClient: c,
		LeaseID:         leaseID,
	}, nil
}

func (c ContainerClient) GetAccountInfo(ctx context.Context) (ContainerGetAccountInfoResponse, error) {
	resp, err := c.client.GetAccountInfo(ctx, nil)

	return resp, handleError(err)
}

// Create creates a new container within a storage account. If a container with the same name already exists, the operation fails.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/create-container.
func (c ContainerClient) Create(ctx context.Context, options *CreateContainerOptions) (ContainerCreateResponse, error) {
	basics, cpkInfo := options.pointers()
	resp, err := c.client.Create(ctx, basics, cpkInfo)

	return resp, handleError(err)
}

// CreateIfNotExists operation creates a new container under the specified account.
// If the container with the same name already exists, it is not changed.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/create-container.
func (c ContainerClient) CreateIfNotExists(ctx context.Context, options *CreateContainerOptions) (ContainerCreateResponse, error) {
	basics, cpkInfo := options.pointers()

	resp, err := c.client.Create(ctx, basics, cpkInfo)

	// If container already exists, return nil error
	err = handleError(err)
	if err != nil && strings.Contains(err.Error(), string(StorageErrorCodeContainerAlreadyExists)) {
		return ContainerCreateResponse{}, nil
	}

	return resp, err
}

// Delete marks the specified container for deletion. The container and any blobs contained within it are later deleted during garbage collection.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/delete-container.
func (c ContainerClient) Delete(ctx context.Context, options *DeleteContainerOptions) (ContainerDeleteResponse, error) {
	basics, leaseInfo, accessConditions := options.pointers()
	resp, err := c.client.Delete(ctx, basics, leaseInfo, accessConditions)

	return resp, handleError(err)
}

// DeleteIfExists operation marks the specified container for deletion if it exists.
// The container and any blobs contained within it are later deleted during garbage collection.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/delete-container
func (c ContainerClient) DeleteIfExists(ctx context.Context, options *DeleteContainerOptions) (ContainerDeleteResponse, error) {
	basics, leaseInfo, accessConditions := options.pointers()
	resp, err := c.client.Delete(ctx, basics, leaseInfo, accessConditions)

	err = handleError(err)
	// If container is already deleted, return nil error
	if err != nil && (strings.Contains(err.Error(), string(StorageErrorCodeContainerNotFound)) || strings.Contains(err.Error(), string(StorageErrorCodeBlobNotFound))) {
		return ContainerDeleteResponse{}, nil
	}
	return resp, err
}

func (c ContainerClient) GetMetadata(ctx context.Context, gpo *GetPropertiesOptionsContainer) (map[string]string, error) {
	resp, err := c.GetProperties(ctx, gpo)

	if err != nil {
		return nil, handleError(err)
	}

	return resp.Metadata, nil
}

// GetProperties returns the container's properties.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-container-metadata.
func (c ContainerClient) GetProperties(ctx context.Context, gpo *GetPropertiesOptionsContainer) (ContainerGetPropertiesResponse, error) {
	// NOTE: GetMetadata actually calls GetProperties internally because GetProperties returns the metadata AND the properties.
	// This allows us to not expose a GetProperties method at all simplifying the API.
	// The optionals are nil, like they were in track 1.5
	options, leaseAccess := gpo.pointers()

	resp, err := c.client.GetProperties(ctx, options, leaseAccess)

	return resp, handleError(err)
}

// SetMetadata sets the container's metadata.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/set-container-metadata.
func (c ContainerClient) SetMetadata(ctx context.Context, options *SetMetadataContainerOptions) (ContainerSetMetadataResponse, error) {
	metadataOptions, lac, mac := options.pointers()

	resp, err := c.client.SetMetadata(ctx, metadataOptions, lac, mac)

	return resp, handleError(err)
}

// GetAccessPolicy returns the container's access policy. The access policy indicates whether container's blobs may be accessed publicly.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-container-acl.
func (c ContainerClient) GetAccessPolicy(ctx context.Context, options *GetAccessPolicyOptions) (ContainerGetAccessPolicyResponse, error) {
	o, ac := options.pointers()

	resp, err := c.client.GetAccessPolicy(ctx, o, ac)

	return resp, handleError(err)
}

// SetAccessPolicy sets the container's permissions. The access policy indicates whether blobs in a container may be accessed publicly.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/set-container-acl.
func (c ContainerClient) SetAccessPolicy(ctx context.Context, options *SetAccessPolicyOptions) (ContainerSetAccessPolicyResponse, error) {
	accessPolicy, mac, lac := options.pointers()

	resp, err := c.client.SetAccessPolicy(ctx, &accessPolicy, mac, lac)

	return resp, handleError(err)
}

// ListBlobsFlatSegment returns a pager for blobs starting from the specified Marker. Use an empty
// Marker to start enumeration from the beginning. Blob names are returned in lexicographic order.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/list-blobs.
func (c ContainerClient) ListBlobsFlatSegment(listOptions *ContainerListBlobFlatSegmentOptions) *ContainerListBlobFlatSegmentPager {
	pager := c.client.ListBlobFlatSegment(listOptions)
	// override the generated pager to insert our handleError(error)
	if pager.Err() != nil {
		return pager
	}

	// TODO: Come Here
	//pager.err = func(response *azcore.Response) error {
	//	return handleError(c.client.listBlobFlatSegmentHandleError(response))
	//}

	return pager
}

// ListBlobsHierarchySegment returns a channel of blobs starting from the specified Marker. Use an empty
// Marker to start enumeration from the beginning. Blob names are returned in lexicographic order.
// After getting a segment, process it, and then call ListBlobsHierarchicalSegment again (passing the the
// previously-returned Marker) to get the next segment.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/list-blobs.
// AutoPagerTimeout specifies the amount of time with no read operations before the channel times out and closes. Specify no time and it will be ignored.
// AutoPagerBufferSize specifies the channel's buffer size.
// Both the blob item channel and error channel should be watched. Only one error will be released via this channel (or a nil error, to register a clean exit.)
func (c ContainerClient) ListBlobsHierarchySegment(delimiter string, listOptions *ContainerListBlobHierarchySegmentOptions) *ContainerListBlobHierarchySegmentPager {
	pager := c.client.ListBlobHierarchySegment(delimiter, listOptions)
	// override the generated pager to insert our handleError(error)
	if pager.Err() != nil {
		return pager
	}

	// TODO: Come here
	//p := pager.(*listBlobsHierarchySegmentResponsePager)
	//p.errorer = func(response *azcore.Response) error {
	//	return handleError(c.client.listBlobHierarchySegmentHandleError(response))
	//}

	return pager
}

func (c ContainerClient) CanGetContainerSASToken() bool {
	return c.cred != nil
}

// GetContainerSASToken is a convenience method for generating a SAS token for the currently pointed at container.
// It can only be used if the supplied azcore.Credential during creation was a SharedKeyCredential.
// This validity can be checked with CanGetContainerSASToken().
func (c ContainerClient) GetContainerSASToken(permissions BlobSASPermissions, validityTime time.Duration) (SASQueryParameters, error) {
	urlParts := NewBlobURLParts(c.URL())

	// Containers do not have snapshots, nor versions.

	return BlobSASSignatureValues{
		ContainerName: urlParts.ContainerName,

		Permissions: permissions.String(),

		StartTime:  time.Now().UTC(),
		ExpiryTime: time.Now().UTC().Add(validityTime),
	}.NewSASQueryParameters(c.cred.(*SharedKeyCredential))
}
