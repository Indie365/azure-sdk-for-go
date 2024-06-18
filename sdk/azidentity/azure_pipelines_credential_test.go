// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

func TestAzurePipelinesCredential(t *testing.T) {
	t.Run("getAssertion", func(t *testing.T) {
		srv, close := mock.NewServer()
		defer close()
		t.Setenv(systemOIDCRequestURI, srv.URL())
		connectionID := "connection"
		expected, err := url.Parse(fmt.Sprintf(
			"%s/?api-version=%s&serviceConnectionId=%s",
			srv.URL(), oidcAPIVersion, connectionID,
		))
		require.NoError(t, err, "test bug: expected URL should parse")
		srv.AppendResponse(
			mock.WithBody([]byte(fmt.Sprintf(`{"oidcToken":%q}`, tokenValue))),
			mock.WithPredicate(func(r *http.Request) bool {
				require.Equal(t, http.MethodPost, r.Method)
				require.Equal(t, expected.Host, r.Host)
				require.Equal(t, expected.Path, r.URL.Path)
				require.Equal(t, expected.RawQuery, r.URL.RawQuery)
				return true
			}),
		)
		srv.AppendResponse()
		o := AzurePipelinesCredentialOptions{
			ClientOptions: azcore.ClientOptions{
				Transport: srv,
			},
		}
		cred, err := NewAzurePipelinesCredential(fakeTenantID, fakeClientID, connectionID, tokenValue, &o)
		require.NoError(t, err)
		actual, err := cred.getAssertion(ctx)
		require.NoError(t, err)
		require.Equal(t, tokenValue, actual)
	})
	t.Run("Live", func(t *testing.T) {
		if recording.GetRecordMode() != recording.LiveMode {
			t.Skip("this test runs only live in an Azure Pipeline with a configured service connection")
		}
		clientID := os.Getenv("AZURESUBSCRIPTION_CLIENT_ID")
		connectionID := os.Getenv("AZURESUBSCRIPTION_SERVICE_CONNECTION_ID")
		systemAccessToken := os.Getenv("SYSTEM_ACCESSTOKEN")
		tenantID := os.Getenv("AZURESUBSCRIPTION_TENANT_ID")
		for _, s := range []string{clientID, connectionID, systemAccessToken, tenantID} {
			if s == "" {
				t.Skip("set AZURESUBSCRIPTION_CLIENT_ID, AZURESUBSCRIPTION_SERVICE_CONNECTION_ID, AZURESUBSCRIPTION_TENANT_ID and SYSTEM_ACCESSTOKEN to run this test")
			}
		}
		cred, err := NewAzurePipelinesCredential(tenantID, clientID, connectionID, systemAccessToken, nil)
		require.NoError(t, err)
		testGetTokenSuccess(t, cred, "https://vault.azure.net/.default")
	})
}
