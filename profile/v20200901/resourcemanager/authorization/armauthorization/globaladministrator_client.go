//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// GlobalAdministratorClient contains the methods for the GlobalAdministrator group.
// Don't use this type directly, use NewGlobalAdministratorClient() instead.
type GlobalAdministratorClient struct {
	internal *arm.Client
}

// NewGlobalAdministratorClient creates a new instance of GlobalAdministratorClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGlobalAdministratorClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GlobalAdministratorClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armauthorization.GlobalAdministratorClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GlobalAdministratorClient{
		internal: cl,
	}
	return client, nil
}

// ElevateAccess - Elevates access for a Global Administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-07-01
//   - options - GlobalAdministratorClientElevateAccessOptions contains the optional parameters for the GlobalAdministratorClient.ElevateAccess
//     method.
func (client *GlobalAdministratorClient) ElevateAccess(ctx context.Context, options *GlobalAdministratorClientElevateAccessOptions) (GlobalAdministratorClientElevateAccessResponse, error) {
	req, err := client.elevateAccessCreateRequest(ctx, options)
	if err != nil {
		return GlobalAdministratorClientElevateAccessResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GlobalAdministratorClientElevateAccessResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalAdministratorClientElevateAccessResponse{}, runtime.NewResponseError(resp)
	}
	return GlobalAdministratorClientElevateAccessResponse{}, nil
}

// elevateAccessCreateRequest creates the ElevateAccess request.
func (client *GlobalAdministratorClient) elevateAccessCreateRequest(ctx context.Context, options *GlobalAdministratorClientElevateAccessOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/elevateAccess"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
