// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// UsernamePasswordCredential enables authentication to Azure Active Directory using a user's  username and password. If the user has MFA enabled this
// credential will fail to get a token returning an AuthenticationFailureError. Also, this credential requires a high degree of trust and is not
// recommended outside of prototyping when more secure credentials can be used.
type UsernamePasswordCredential struct {
	azcore.TokenCredential
	client   *aadIdentityClient
	tenantID string // Gets the Azure Active Directory tenant (directory) ID of the service principal
	clientID string // Gets the client (application) ID of the service principal
	username string // Gets the user account's user name
	password string // Gets the user account's password
}

// NewUsernamePasswordCredential constructs a new UsernamePasswordCredential with the details needed to authenticate against Azure Active Directory with
// a simple username and password.
// tenantID: The Azure Active Directory tenant (directory) ID of the service principal.
// clientID: The client (application) ID of the service principal.
// username: A user's account username
// password: A user's account password
// options: TokenCredentialOptions used to configure the pipeline for the requests sent to Azure Active Directory.
func NewUsernamePasswordCredential(tenantID string, clientID string, username string, password string, options *TokenCredentialOptions) (*UsernamePasswordCredential, error) {
	c, err := newAADIdentityClient(options)
	if err != nil {
		return nil, err
	}
	c.options.AuthorityHost.Path = path.Join(c.options.AuthorityHost.Path, tenantID+tokenEndpoint)
	return &UsernamePasswordCredential{tenantID: tenantID, clientID: clientID, username: username, password: password, client: c}, nil
}

// GetToken obtains a token from Azure Active Directory using the specified username and password.
// scopes: The list of scopes for which the token will have access.
// ctx: The context used to control the request lifetime.
// Returns an AccessToken which can be used to authenticate service client calls.
func (c *UsernamePasswordCredential) GetToken(ctx context.Context, opts azcore.TokenRequestOptions) (*azcore.AccessToken, error) {
	tk, err := c.client.authenticateUsernamePassword(ctx, c.tenantID, c.clientID, c.username, c.password, opts.Scopes)
	log := azcore.Log()
	if err != nil {
		msg := fmt.Sprintf("Azure Identity => ERROR in GetToken() call for %T: %s", c, err.Error())
		log.Write(azcore.LogError, msg)
	} else {
		msg := fmt.Sprintf("Azure Identity => GetToken() result for %T: SUCCESS", c)
		log.Write(LogCredential, msg)
		vmsg := fmt.Sprintf("Azure Identity => Scopes: [%s]", strings.Join(opts.Scopes, ", "))
		log.Write(LogCredential, vmsg)
	}
	return tk, err
}

// AuthenticationPolicy implements the azcore.Credential interface on UsernamePasswordCredential.
func (c *UsernamePasswordCredential) AuthenticationPolicy(options azcore.AuthenticationPolicyOptions) azcore.Policy {
	return newBearerTokenPolicy(c, options)
}
