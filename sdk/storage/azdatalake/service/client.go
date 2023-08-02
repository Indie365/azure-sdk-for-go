//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package service

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/filesystem"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/base"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/generated"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/generated_blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/sas"
	"net/http"
	"strings"
	"time"
)

// ClientOptions contains the optional parameters when creating a Client.
type ClientOptions base.ClientOptions

// Client represents a URL to the Azure Datalake Storage service.
type Client base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client]

// NewClient creates an instance of Client with the specified values.
//   - serviceURL - the URL of the blob e.g. https://<account>.dfs.core.windows.net/
//   - cred - an Azure AD credential, typically obtained via the azidentity module
//   - options - client options; pass nil to accept the default values
func NewClient(serviceURL string, cred azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	blobServiceURL, datalakeServiceURL := shared.GetURLs(serviceURL)
	authPolicy := runtime.NewBearerTokenPolicy(cred, []string{shared.TokenScope}, nil)
	conOptions := shared.GetClientOptions(options)
	plOpts := runtime.PipelineOptions{
		PerRetry: []policy.Policy{authPolicy},
	}
	base.SetPipelineOptions((*base.ClientOptions)(conOptions), &plOpts)

	azClient, err := azcore.NewClient(shared.ServiceClient, exported.ModuleVersion, plOpts, &conOptions.ClientOptions)
	if err != nil {
		return nil, err
	}

	if options == nil {
		options = &ClientOptions{}
	}
	blobServiceClientOpts := service.ClientOptions{
		ClientOptions: options.ClientOptions,
	}
	blobSvcClient, _ := service.NewClient(blobServiceURL, cred, &blobServiceClientOpts)
	svcClient := base.NewServiceClient(datalakeServiceURL, blobServiceURL, blobSvcClient, azClient, nil, &cred, (*base.ClientOptions)(conOptions))

	return (*Client)(svcClient), nil
}

// NewClientWithNoCredential creates an instance of Client with the specified values.
//   - serviceURL - the URL of the storage account e.g. https://<account>.dfs.core.windows.net/
//   - options - client options; pass nil to accept the default values.
func NewClientWithNoCredential(serviceURL string, options *ClientOptions) (*Client, error) {
	blobServiceURL, datalakeServiceURL := shared.GetURLs(serviceURL)
	conOptions := shared.GetClientOptions(options)
	plOpts := runtime.PipelineOptions{}
	base.SetPipelineOptions((*base.ClientOptions)(conOptions), &plOpts)

	azClient, err := azcore.NewClient(shared.ServiceClient, exported.ModuleVersion, plOpts, &conOptions.ClientOptions)
	if err != nil {
		return nil, err
	}

	if options == nil {
		options = &ClientOptions{}
	}
	blobServiceClientOpts := service.ClientOptions{
		ClientOptions: options.ClientOptions,
	}
	blobSvcClient, _ := service.NewClientWithNoCredential(blobServiceURL, &blobServiceClientOpts)
	svcClient := base.NewServiceClient(datalakeServiceURL, blobServiceURL, blobSvcClient, azClient, nil, nil, (*base.ClientOptions)(conOptions))

	return (*Client)(svcClient), nil
}

// NewClientWithSharedKeyCredential creates an instance of Client with the specified values.
//   - serviceURL - the URL of the storage account e.g. https://<account>.dfs.core.windows.net/
//   - cred - a SharedKeyCredential created with the matching storage account and access key
//   - options - client options; pass nil to accept the default values
func NewClientWithSharedKeyCredential(serviceURL string, cred *SharedKeyCredential, options *ClientOptions) (*Client, error) {
	blobServiceURL, datalakeServiceURL := shared.GetURLs(serviceURL)
	authPolicy := exported.NewSharedKeyCredPolicy(cred)
	conOptions := shared.GetClientOptions(options)
	plOpts := runtime.PipelineOptions{
		PerRetry: []policy.Policy{authPolicy},
	}
	base.SetPipelineOptions((*base.ClientOptions)(conOptions), &plOpts)

	azClient, err := azcore.NewClient(shared.ServiceClient, exported.ModuleVersion, plOpts, &conOptions.ClientOptions)
	if err != nil {
		return nil, err
	}

	if options == nil {
		options = &ClientOptions{}
	}
	blobServiceClientOpts := service.ClientOptions{
		ClientOptions: options.ClientOptions,
	}
	blobSharedKey, err := cred.ConvertToBlobSharedKey()
	if err != nil {
		return nil, err
	}
	blobSvcClient, _ := service.NewClientWithSharedKeyCredential(blobServiceURL, blobSharedKey, &blobServiceClientOpts)
	svcClient := base.NewServiceClient(datalakeServiceURL, blobServiceURL, blobSvcClient, azClient, cred, nil, (*base.ClientOptions)(conOptions))

	return (*Client)(svcClient), nil
}

// NewClientFromConnectionString creates an instance of Client with the specified values.
//   - connectionString - a connection string for the desired storage account
//   - options - client options; pass nil to accept the default values
func NewClientFromConnectionString(connectionString string, options *ClientOptions) (*Client, error) {
	parsed, err := shared.ParseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}

	if parsed.AccountKey != "" && parsed.AccountName != "" {
		credential, err := exported.NewSharedKeyCredential(parsed.AccountName, parsed.AccountKey)
		if err != nil {
			return nil, err
		}
		return NewClientWithSharedKeyCredential(parsed.ServiceURL, credential, options)
	}

	return NewClientWithNoCredential(parsed.ServiceURL, options)
}

func (s *Client) getClientOptions() *base.ClientOptions {
	return base.GetCompositeClientOptions((*base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client])(s))
}

// NewFilesystemClient creates a new filesystem.Client object by concatenating filesystemName to the end of this Client's URL.
// The new filesystem.Client uses the same request policy pipeline as the Client.
func (s *Client) NewFilesystemClient(filesystemName string) *filesystem.Client {
	filesystemURL := runtime.JoinPaths(s.generatedServiceClientWithDFS().Endpoint(), filesystemName)
	containerURL, filesystemURL := shared.GetURLs(filesystemURL)
	return (*filesystem.Client)(base.NewFilesystemClient(filesystemURL, containerURL, s.serviceClient().NewContainerClient(filesystemName), s.generatedServiceClientWithDFS().InternalClient().WithClientName(shared.FilesystemClient), s.sharedKey(), s.identityCredential(), s.getClientOptions()))
}

// GetUserDelegationCredential obtains a UserDelegationKey object using the base ServiceURL object.
// OAuth is required for this call, as well as any role that can delegate access to the storage account.
func (s *Client) GetUserDelegationCredential(ctx context.Context, info KeyInfo, o *GetUserDelegationCredentialOptions) (*UserDelegationCredential, error) {
	url, err := azdatalake.ParseURL(s.BlobURL())
	if err != nil {
		return nil, err
	}

	getUserDelegationKeyOptions := o.format()
	udk, err := s.generatedServiceClientWithBlob().GetUserDelegationKey(ctx, info, getUserDelegationKeyOptions)
	if err != nil {
		return nil, exported.ConvertToDFSError(err)
	}

	return exported.NewUserDelegationCredential(strings.Split(url.Host, ".")[0], udk.UserDelegationKey), nil
}

func (s *Client) generatedServiceClientWithDFS() *generated.ServiceClient {
	svcClientWithDFS, _, _ := base.InnerClients((*base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client])(s))
	return svcClientWithDFS
}

func (s *Client) generatedServiceClientWithBlob() *generated_blob.ServiceClient {
	_, svcClientWithBlob, _ := base.InnerClients((*base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client])(s))
	return svcClientWithBlob
}

func (s *Client) serviceClient() *service.Client {
	_, _, serviceClient := base.InnerClients((*base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client])(s))
	return serviceClient
}

func (s *Client) sharedKey() *exported.SharedKeyCredential {
	return base.SharedKeyComposite((*base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client])(s))
}

func (s *Client) identityCredential() *azcore.TokenCredential {
	return base.IdentityCredentialComposite((*base.CompositeClient[generated.ServiceClient, generated_blob.ServiceClient, service.Client])(s))
}

// DFSURL returns the URL endpoint used by the Client object.
func (s *Client) DFSURL() string {
	return s.generatedServiceClientWithDFS().Endpoint()
}

// BlobURL returns the URL endpoint used by the Client object.
func (s *Client) BlobURL() string {
	return s.generatedServiceClientWithBlob().Endpoint()
}

// CreateFilesystem creates a new filesystem under the specified account. (blob3)
func (s *Client) CreateFilesystem(ctx context.Context, filesystem string, options *CreateFilesystemOptions) (CreateFilesystemResponse, error) {
	filesystemClient := s.NewFilesystemClient(filesystem)
	resp, err := filesystemClient.Create(ctx, options)
	err = exported.ConvertToDFSError(err)
	return resp, err
}

// DeleteFilesystem deletes the specified filesystem. (blob3)
func (s *Client) DeleteFilesystem(ctx context.Context, filesystem string, options *DeleteFilesystemOptions) (DeleteFilesystemResponse, error) {
	filesystemClient := s.NewFilesystemClient(filesystem)
	resp, err := filesystemClient.Delete(ctx, options)
	err = exported.ConvertToDFSError(err)
	return resp, err
}

// SetProperties sets properties for a storage account's File service endpoint. (blob3)
func (s *Client) SetProperties(ctx context.Context, options *SetPropertiesOptions) (SetPropertiesResponse, error) {
	opts := options.format()
	resp, err := s.serviceClient().SetProperties(ctx, opts)
	err = exported.ConvertToDFSError(err)
	return resp, err
}

// GetProperties gets properties for a storage account's File service endpoint. (blob3)
func (s *Client) GetProperties(ctx context.Context, options *GetPropertiesOptions) (GetPropertiesResponse, error) {
	opts := options.format()
	resp, err := s.serviceClient().GetProperties(ctx, opts)
	err = exported.ConvertToDFSError(err)
	return resp, err

}

// NewListFilesystemsPager operation returns a pager of the shares under the specified account. (blob3)
// For more information, see https://learn.microsoft.com/en-us/rest/api/storageservices/list-shares
func (s *Client) NewListFilesystemsPager(o *ListFilesystemsOptions) *runtime.Pager[ListFilesystemsResponse] {
	listOptions := generated_blob.ServiceClientListContainersSegmentOptions{}
	if o != nil {
		if o.Include.Deleted {
			listOptions.Include = append(listOptions.Include, generated_blob.ListContainersIncludeTypeDeleted)
		}
		if o.Include.Metadata {
			listOptions.Include = append(listOptions.Include, generated_blob.ListContainersIncludeTypeMetadata)
		}
		if o.Include.System {
			listOptions.Include = append(listOptions.Include, generated_blob.ListContainersIncludeTypeSystem)
		}
		listOptions.Marker = o.Marker
		listOptions.Maxresults = o.MaxResults
		listOptions.Prefix = o.Prefix
	}
	return runtime.NewPager(runtime.PagingHandler[ListFilesystemsResponse]{
		More: func(page ListFilesystemsResponse) bool {
			return page.NextMarker != nil && len(*page.NextMarker) > 0
		},
		Fetcher: func(ctx context.Context, page *ListFilesystemsResponse) (ListFilesystemsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = s.generatedServiceClientWithBlob().ListContainersSegmentCreateRequest(ctx, &listOptions)
			} else {
				listOptions.Marker = page.NextMarker
				req, err = s.generatedServiceClientWithBlob().ListContainersSegmentCreateRequest(ctx, &listOptions)
			}
			if err != nil {
				return ListFilesystemsResponse{}, err
			}
			resp, err := s.generatedServiceClientWithBlob().InternalClient().Pipeline().Do(req)
			if err != nil {
				return ListFilesystemsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ListFilesystemsResponse{}, runtime.NewResponseError(resp)
			}
			return s.generatedServiceClientWithBlob().ListContainersSegmentHandleResponse(resp)
		},
	})
}

// GetSASURL is a convenience method for generating a SAS token for the currently pointed at account.
// It can only be used if the credential supplied during creation was a SharedKeyCredential.
func (s *Client) GetSASURL(resources sas.AccountResourceTypes, permissions sas.AccountPermissions, expiry time.Time, o *GetSASURLOptions) (string, error) {
	// format all options to blob service options
	res, perms, opts := o.format(resources, permissions)
	resp, err := s.serviceClient().GetSASURL(res, perms, expiry, opts)
	err = exported.ConvertToDFSError(err)
	return resp, err
}
