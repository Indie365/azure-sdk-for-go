//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package pageblob

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/base"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/shared"
)

// ClientOptions adds additional client options while constructing connection
type ClientOptions = exported.ClientOptions

// SharedKeyCredential contains an account's name and its primary or secondary key.
type SharedKeyCredential = exported.SharedKeyCredential

// Client represents a client to an Azure Storage page blob;
type Client base.CompositeClient[generated.BlobClient, generated.PageBlobClient]

// NewClient creates a ServiceClient object using the specified URL, Azure AD credential, and options.
// Example of serviceURL: https://<your_storage_account>.blob.core.windows.net
func NewClient(blobURL string, cred azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	authPolicy := runtime.NewBearerTokenPolicy(cred, []string{shared.TokenScope}, nil)
	conOptions := exported.GetConnectionOptions(options)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(shared.ModuleName, shared.ModuleVersion, runtime.PipelineOptions{}, conOptions)

	return (*Client)(base.NewPageBlobClient(blobURL, pl)), nil
}

// NewClientWithNoCredential creates a ServiceClient object using the specified URL and options.
// Example of serviceURL: https://<your_storage_account>.blob.core.windows.net?<SAS token>
func NewClientWithNoCredential(blobURL string, options *ClientOptions) (*Client, error) {
	conOptions := exported.GetConnectionOptions(options)
	pl := runtime.NewPipeline(shared.ModuleName, shared.ModuleVersion, runtime.PipelineOptions{}, conOptions)

	return (*Client)(base.NewPageBlobClient(blobURL, pl)), nil
}

// NewClientWithSharedKey creates a ServiceClient object using the specified URL, shared key, and options.
// Example of serviceURL: https://<your_storage_account>.blob.core.windows.net
func NewClientWithSharedKey(blobURL string, cred *SharedKeyCredential, options *ClientOptions) (*Client, error) {
	authPolicy := exported.NewSharedKeyCredPolicy(cred)
	conOptions := exported.GetConnectionOptions(options)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(shared.ModuleName, shared.ModuleVersion, runtime.PipelineOptions{}, conOptions)

	return (*Client)(base.NewPageBlobClient(blobURL, pl)), nil
}

// NewClientFromConnectionString creates Client from a connection String
func NewClientFromConnectionString(connectionString, containerName, blobName string, options *ClientOptions) (*Client, error) {
	parsed, err := shared.ParseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	parsed.ServiceURL = runtime.JoinPaths(parsed.ServiceURL, containerName, blobName)

	if parsed.AccountKey != "" && parsed.AccountName != "" {
		credential, err := exported.NewSharedKeyCredential(parsed.AccountName, parsed.AccountKey)
		if err != nil {
			return nil, err
		}
		return NewClientWithSharedKey(parsed.ServiceURL, credential, options)
	}

	return NewClientWithNoCredential(parsed.ServiceURL, options)
}

func (pb *Client) generated() *generated.PageBlobClient {
	_, pageBlob := base.InnerClients((*base.CompositeClient[generated.BlobClient, generated.PageBlobClient])(pb))
	return pageBlob
}

// URL returns the URL endpoint used by the Client object.
func (pb *Client) URL() string {
	return pb.generated().Endpoint()
}

// Blob returns the blob client for this PageBlob client.
func (pb *Client) blobClient() *blob.Client {
	innerBlob, _ := base.InnerClients((*base.CompositeClient[generated.BlobClient, generated.PageBlobClient])(pb))
	return (*blob.Client)(innerBlob)
}

// WithSnapshot creates a new PageBlobURL object identical to the source but with the specified snapshot timestamp.
// Pass "" to remove the snapshot returning a URL to the base blob.
func (pb *Client) WithSnapshot(snapshot string) (*Client, error) {
	p, err := exported.ParseBlobURL(pb.URL())
	if err != nil {
		return nil, err
	}
	p.Snapshot = snapshot

	return (*Client)(base.NewPageBlobClient(p.URL(), pb.generated().Pipeline())), nil
}

// WithVersionID creates a new PageBlobURL object identical to the source but with the specified snapshot timestamp.
// Pass "" to remove the version returning a URL to the base blob.
func (pb *Client) WithVersionID(versionID string) (*Client, error) {
	p, err := exported.ParseBlobURL(pb.URL())
	if err != nil {
		return nil, err
	}
	p.VersionID = versionID

	return (*Client)(base.NewPageBlobClient(p.URL(), pb.generated().Pipeline())), nil
}

// Create creates a page blob of the specified length. Call PutPage to upload data to a page blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/put-blob.
func (pb *Client) Create(ctx context.Context, size int64, o *CreateOptions) (CreateResponse, error) {
	var createOptions *generated.PageBlobClientCreateOptions
	var headers *generated.BlobHTTPHeaders
	var cpkInfo *generated.CpkInfo
	var cpkScope *generated.CpkScopeInfo
	var leaseConditions *generated.LeaseAccessConditions
	var accessConditions *generated.ModifiedAccessConditions

	if o != nil {
		createOptions = &generated.PageBlobClientCreateOptions{
			BlobSequenceNumber: o.SequenceNumber,
			BlobTagsString:     shared.SerializeBlobTagsToStrPtr(o.Tags),
			Metadata:           o.Metadata,
			Tier:               o.Tier,
		}
		leaseConditions, accessConditions = exported.FormatBlobAccessConditions(o.AccessConditions)
		cpkInfo = o.CpkInfo
		cpkScope = o.CpkScopeInfo
	}

	resp, err := pb.generated().Create(ctx, 0, size, createOptions, headers, leaseConditions, cpkInfo, cpkScope, accessConditions)
	return resp, err
}

// UploadPages writes 1 or more pages to the page blob. The start offset and the stream size must be a multiple of 512 bytes.
// This method panics if the stream is not at position 0.
// Note that the http client closes the body stream after the request is sent to the service.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/put-page.
/*func (pb *Client) UploadPages(ctx context.Context, body io.ReadSeekCloser, options *PageBlobUploadPagesOptions) (PageBlobUploadPagesResponse, error) {
	count, err := validateSeekableStreamAt0AndGetCount(body)

	if err != nil {
		return PageBlobUploadPagesResponse{}, err
	}

	uploadPagesOptions, leaseAccessConditions, cpkInfo, cpkScopeInfo, sequenceNumberAccessConditions, modifiedAccessConditions := options.format()

	resp, err := pb.client.UploadPages(ctx, count, body, uploadPagesOptions, leaseAccessConditions,
		cpkInfo, cpkScopeInfo, sequenceNumberAccessConditions, modifiedAccessConditions)

	return toPageBlobUploadPagesResponse(resp), handleError(err)
}

// UploadPagesFromURL copies 1 or more pages from a source URL to the page blob.
// The sourceOffset specifies the start offset of source data to copy from.
// The destOffset specifies the start offset of data in page blob will be written to.
// The count must be a multiple of 512 bytes.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/put-page-from-url.
func (pb *Client) UploadPagesFromURL(ctx context.Context, source string, sourceOffset, destOffset, count int64,
	options *PageBlobUploadPagesFromURLOptions) (PageBlobUploadPagesFromURLResponse, error) {

	uploadPagesFromURLOptions, cpkInfo, cpkScopeInfo, leaseAccessConditions, sequenceNumberAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions := options.format()

	resp, err := pb.client.UploadPagesFromURL(ctx, source, rangeToString(sourceOffset, count), 0,
		rangeToString(destOffset, count), uploadPagesFromURLOptions, cpkInfo, cpkScopeInfo, leaseAccessConditions,
		sequenceNumberAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)

	return toPageBlobUploadPagesFromURLResponse(resp), handleError(err)
}

// ClearPages frees the specified pages from the page blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/put-page.
func (pb *Client) ClearPages(ctx context.Context, pageRange HttpRange, options *PageBlobClearPagesOptions) (PageBlobClearPagesResponse, error) {
	clearOptions := &pageBlobClientClearPagesOptions{
		Range: pageRange.format(),
	}

	leaseAccessConditions, cpkInfo, cpkScopeInfo, sequenceNumberAccessConditions, modifiedAccessConditions := options.format()

	resp, err := pb.client.ClearPages(ctx, 0, clearOptions, leaseAccessConditions, cpkInfo,
		cpkScopeInfo, sequenceNumberAccessConditions, modifiedAccessConditions)

	return toPageBlobClearPagesResponse(resp), handleError(err)
}

// GetPageRanges returns the list of valid page ranges for a page blob or snapshot of a page blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-page-ranges.
func (pb *Client) GetPageRanges(options *PageBlobGetPageRangesOptions) *PageBlobGetPageRangesPager {
	getPageRangesOptions, leaseAccessConditions, modifiedAccessConditions := options.format()

	pageBlobGetPageRangesPager := pb.client.GetPageRanges(getPageRangesOptions, leaseAccessConditions, modifiedAccessConditions)

	// Fixing Advancer
	pageBlobGetPageRangesPager.advancer = func(ctx context.Context, response pageBlobClientGetPageRangesResponse) (*policy.Request, error) {
		getPageRangesOptions.Marker = response.NextMarker
		req, err := pb.client.getPageRangesCreateRequest(ctx, getPageRangesOptions, leaseAccessConditions, modifiedAccessConditions)
		if err != nil {
			return nil, handleError(err)
		}
		queryValues, err := url.ParseQuery(req.Raw().URL.RawQuery)
		if err != nil {
			return nil, handleError(err)
		}
		req.Raw().URL.RawQuery = queryValues.Encode()
		return req, nil
	}

	return toPageBlobGetPageRangesPager(pageBlobGetPageRangesPager)
}

// GetPageRangesDiff gets the collection of page ranges that differ between a specified snapshot and this page blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-page-ranges.
func (pb *Client) GetPageRangesDiff(options *PageBlobGetPageRangesDiffOptions) *PageBlobGetPageRangesDiffPager {
	getPageRangesDiffOptions, leaseAccessConditions, modifiedAccessConditions := options.format()

	getPageRangesDiffPager := pb.client.GetPageRangesDiff(getPageRangesDiffOptions, leaseAccessConditions, modifiedAccessConditions)

	// Fixing Advancer
	getPageRangesDiffPager.advancer = func(ctx context.Context, response pageBlobClientGetPageRangesDiffResponse) (*policy.Request, error) {
		getPageRangesDiffOptions.Marker = response.NextMarker
		req, err := pb.client.getPageRangesDiffCreateRequest(ctx, getPageRangesDiffOptions, leaseAccessConditions, modifiedAccessConditions)
		if err != nil {
			return nil, handleError(err)
		}
		queryValues, err := url.ParseQuery(req.Raw().URL.RawQuery)
		if err != nil {
			return nil, handleError(err)
		}
		req.Raw().URL.RawQuery = queryValues.Encode()
		return req, nil
	}

	return toPageBlobGetPageRangesDiffPager(getPageRangesDiffPager)
}

// Resize resizes the page blob to the specified size (which must be a multiple of 512).
// For more information, see https://docs.microsoft.com/rest/api/storageservices/set-blob-properties.
func (pb *Client) Resize(ctx context.Context, size int64, options *PageBlobResizeOptions) (PageBlobResizeResponse, error) {
	resizeOptions, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions := options.format()

	resp, err := pb.client.Resize(ctx, size, resizeOptions, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)

	return toPageBlobResizeResponse(resp), handleError(err)
}

// UpdateSequenceNumber sets the page blob's sequence number.
func (pb *Client) UpdateSequenceNumber(ctx context.Context, options *PageBlobUpdateSequenceNumberOptions) (PageBlobUpdateSequenceNumberResponse, error) {
	actionType, updateOptions, lac, mac := options.format()
	resp, err := pb.client.UpdateSequenceNumber(ctx, *actionType, updateOptions, lac, mac)

	return toPageBlobUpdateSequenceNumberResponse(resp), handleError(err)
}

// StartCopyIncremental begins an operation to start an incremental copy from one page blob's snapshot to this page blob.
// The snapshot is copied such that only the differential changes between the previously copied snapshot are transferred to the destination.
// The copied snapshots are complete copies of the original snapshot and can be read or copied from as usual.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/incremental-copy-blob and
// https://docs.microsoft.com/en-us/azure/virtual-machines/windows/incremental-snapshots.
func (pb *Client) StartCopyIncremental(ctx context.Context, copySource string, prevSnapshot string, options *PageBlobCopyIncrementalOptions) (PageBlobCopyIncrementalResponse, error) {
	copySourceURL, err := url.Parse(copySource)
	if err != nil {
		return PageBlobCopyIncrementalResponse{}, err
	}

	queryParams := copySourceURL.Query()
	queryParams.Set("snapshot", prevSnapshot)
	copySourceURL.RawQuery = queryParams.Encode()

	pageBlobCopyIncrementalOptions, modifiedAccessConditions := options.format()
	resp, err := pb.client.CopyIncremental(ctx, copySourceURL.String(), pageBlobCopyIncrementalOptions, modifiedAccessConditions)

	return toPageBlobCopyIncrementalResponse(resp), handleError(err)
}
*/

// Redeclared APIs ----- Copy over to Append blob and Page blob as well.

// Download reads a range of bytes from a blob. The response also includes the blob's properties and metadata.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-blob.
func (pb *Client) Download(ctx context.Context, o *blob.DownloadOptions) (blob.DownloadResponse, error) {
	return pb.blobClient().Download(ctx, o)
}

// Delete marks the specified blob or snapshot for deletion. The blob is later deleted during garbage collection.
// Note that deleting a blob also deletes all its snapshots.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/delete-blob.
func (pb *Client) Delete(ctx context.Context, o *blob.DeleteOptions) (blob.DeleteResponse, error) {
	return pb.blobClient().Delete(ctx, o)
}

// Undelete restores the contents and metadata of a soft-deleted blob and any associated soft-deleted snapshots.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/undelete-blob.
func (pb *Client) Undelete(ctx context.Context, o *blob.UndeleteOptions) (blob.UndeleteResponse, error) {
	return pb.blobClient().Undelete(ctx, o)
}

// SetTier operation sets the tier on a blob. The operation is allowed on a page
// blob in a premium storage account and on a block blob in a blob storage account (locally
// redundant storage only). A premium page blob's tier determines the allowed size, IOPS, and
// bandwidth of the blob. A block blob's tier determines Hot/Cool/Archive storage type. This operation
// does not update the blob's ETag.
// For detailed information about block blob level tiering see https://docs.microsoft.com/en-us/azure/storage/blobs/storage-blob-storage-tiers.
func (pb *Client) SetTier(ctx context.Context, tier AccessTier, o *blob.SetTierOptions) (blob.SetTierResponse, error) {
	return pb.blobClient().SetTier(ctx, tier, o)
}

// GetProperties returns the blob's properties.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-blob-properties.
func (pb *Client) GetProperties(ctx context.Context, o *blob.GetPropertiesOptions) (blob.GetPropertiesResponse, error) {
	return pb.blobClient().GetProperties(ctx, o)
}

// SetHTTPHeaders changes a blob's HTTP headers.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/set-blob-properties.
func (pb *Client) SetHTTPHeaders(ctx context.Context, HTTPHeaders blob.HTTPHeaders, o *blob.SetHTTPHeadersOptions) (blob.SetHTTPHeadersResponse, error) {
	return pb.blobClient().SetHTTPHeaders(ctx, HTTPHeaders, o)
}

// SetMetadata changes a blob's metadata.
// https://docs.microsoft.com/rest/api/storageservices/set-blob-metadata.
func (pb *Client) SetMetadata(ctx context.Context, metadata map[string]string, o *blob.SetMetadataOptions) (blob.SetMetadataResponse, error) {
	return pb.blobClient().SetMetadata(ctx, metadata, o)
}

// CreateSnapshot creates a read-only snapshot of a blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/snapshot-blob.
func (pb *Client) CreateSnapshot(ctx context.Context, o *blob.CreateSnapshotOptions) (blob.CreateSnapshotResponse, error) {
	return pb.blobClient().CreateSnapshot(ctx, o)
}

// StartCopyFromURL copies the data at the source URL to a blob.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/copy-blob.
func (pb *Client) StartCopyFromURL(ctx context.Context, copySource string, o *blob.StartCopyFromURLOptions) (blob.StartCopyFromURLResponse, error) {
	return pb.blobClient().StartCopyFromURL(ctx, copySource, o)
}

// AbortCopyFromURL stops a pending copy that was previously started and leaves a destination blob with 0 length and metadata.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/abort-copy-blob.
func (pb *Client) AbortCopyFromURL(ctx context.Context, copyID string, o *blob.AbortCopyFromURLOptions) (blob.AbortCopyFromURLResponse, error) {
	return pb.blobClient().AbortCopyFromURL(ctx, copyID, o)
}

// SetTags operation enables users to set tags on a blob or specific blob version, but not snapshot.
// Each call to this operation replaces all existing tags attached to the blob.
// To remove all tags from the blob, call this operation with no tags set.
// https://docs.microsoft.com/en-us/rest/api/storageservices/set-blob-tags
func (pb *Client) SetTags(ctx context.Context, o *blob.SetTagsOptions) (blob.SetTagsResponse, error) {
	return pb.blobClient().SetTags(ctx, o)
}

// GetTags operation enables users to get tags on a blob or specific blob version, or snapshot.
// https://docs.microsoft.com/en-us/rest/api/storageservices/get-blob-tags
func (pb *Client) GetTags(ctx context.Context, o *blob.GetTagsOptions) (blob.GetTagsResponse, error) {
	return pb.blobClient().GetTags(ctx, o)
}

// CopyFromURL synchronously copies the data at the source URL to a block blob, with sizes up to 256 MB.
// For more information, see https://docs.microsoft.com/en-us/rest/api/storageservices/copy-blob-from-url.
func (pb *Client) CopyFromURL(ctx context.Context, copySource string, o *blob.CopyFromURLOptions) (blob.CopyFromURLResponse, error) {
	return pb.blobClient().CopyFromURL(ctx, copySource, o)
}
