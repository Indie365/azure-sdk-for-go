//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

// Pipeline represents a primitive for sending HTTP requests and receiving responses.
// Its behavior can be extended by specifying policies during construction.
type Pipeline struct {
	policies []policy.Policy
}

// NewPipeline creates a new Pipeline object from the specified Transport and Policies.
// If no transport is provided then the default *http.Client transport will be used.
func NewPipeline(transport policy.Transporter, policies ...policy.Policy) Pipeline {
	if transport == nil {
		transport = defaultHTTPClient
	}
	// transport policy must always be the last in the slice
	policies = append(policies, shared.PolicyFunc(httpHeaderPolicy), shared.PolicyFunc(bodyDownloadPolicy), transportPolicy{trans: transport})
	return Pipeline{
		policies: policies,
	}
}

// Do is called for each and every HTTP request. It passes the request through all
// the Policy objects (which can transform the Request's URL/query parameters/headers)
// and ultimately sends the transformed HTTP request over the network.
func (p Pipeline) Do(req *policy.Request) (*http.Response, error) {
	shared.SetPolicies(req, p.policies)
	return req.Next()
}

// NewRequest creates a new policy.Request with the specified input.
func NewRequest(ctx context.Context, httpMethod string, endpoint string) (*policy.Request, error) {
	return shared.NewRequest(ctx, httpMethod, endpoint)
}

// AuthenticationOptions contains various options used to create a credential policy.
type AuthenticationOptions struct {
	// TokenRequest is a TokenRequestOptions that includes a scopes field which contains
	// the list of OAuth2 authentication scopes used when requesting a token.
	// This field is ignored for other forms of authentication (e.g. shared key).
	TokenRequest policy.TokenRequestOptions
	// AuxiliaryTenants contains a list of additional tenant IDs to be used to authenticate
	// in cross-tenant applications.
	AuxiliaryTenants []string
}
