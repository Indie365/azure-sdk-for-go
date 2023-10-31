// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azwebpubsub_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azwebpubsub"
	"github.com/stretchr/testify/require"
)

type clientWrapper struct {
	*azwebpubsub.Client
	TestVars testVars
}

var fakeTestVars = testVars{
	Key:          "key",
	Endpoint:     "https://fake.eastus-1.webpubsub.azure.net",
	Subscription: "subscription",
}

type testVars struct {
	Key              string
	Endpoint         string
	Subscription     string
	ConnectionString string

	// KeyLogPath is the value of environment "SSLKEYLOGFILE_TEST", which
	// points to a file on disk where we'll write the TLS pre-master-secret.
	// This is useful if you want to trace parts of this test using Wireshark.
	KeyLogPath string
}

func loadEnv() (testVars, error) {
	var missing []string

	get := func(n string) string {
		if v := os.Getenv(n); v == "" {
			missing = append(missing, n)
		}

		return os.Getenv(n)
	}

	tv := testVars{
		Key:              get("WEBPUBSUB_KEY"),
		Endpoint:         get("WEBPUBSUB_ENDPOINT"),
		ConnectionString: get("WEBPUBSUB_CONNECTIONSTRING"),
	}

	if len(missing) == 3 {
		return testVars{}, fmt.Errorf("Missing env variables: %s", strings.Join(missing, ","))
	}

	// Setting this variable will cause the test clients to dump out the pre-master-key
	// for your HTTP connection. This allows you decrypt a packet capture from wireshark.
	//
	// If you want to do this just set SSLKEYLOGFILE env var to a path on disk and
	// Go will write out the key.
	tv.KeyLogPath = os.Getenv("SSLKEYLOGFILE")
	return tv, nil
}

func newClientWrapper(t *testing.T) clientWrapper {
	var client *azwebpubsub.Client
	var tv testVars

	if recording.GetRecordMode() != recording.PlaybackMode {
		tmpTestVars, err := loadEnv()
		require.NoError(t, err)
		tv = tmpTestVars
	} else {
		tv = fakeTestVars
	}

	if recording.GetRecordMode() == recording.LiveMode {
		if tv.KeyLogPath != "" {
			keyLogWriter, err := os.OpenFile(tv.KeyLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
			require.NoError(t, err)

			t.Cleanup(func() { keyLogWriter.Close() })

			tp := http.DefaultTransport.(*http.Transport).Clone()
			tp.TLSClientConfig = &tls.Config{
				KeyLogWriter: keyLogWriter,
			}

			httpClient := &http.Client{Transport: tp}

			tmpClient, err := azwebpubsub.NewClientFromConnectionString(tv.ConnectionString, &azwebpubsub.ClientOptions{
				ClientOptions: azcore.ClientOptions{
					Transport: httpClient,
				},
			})
			require.NoError(t, err)
			client = tmpClient
		} else {
			tmpClient, err := azwebpubsub.NewClientFromConnectionString(tv.ConnectionString, nil)
			require.NoError(t, err)
			client = tmpClient
		}

	} else {
		tmpClient, err := azwebpubsub.NewClientFromConnectionString(tv.ConnectionString, &azwebpubsub.ClientOptions{
			ClientOptions: azcore.ClientOptions{
				Transport: newRecordingTransporter(t, tv),
			},
		})
		require.NoError(t, err)
		client = tmpClient
	}

	return clientWrapper{
		Client:   client,
		TestVars: tv,
	}
}

func newRecordingTransporter(t *testing.T, testVars testVars) policy.Transporter {
	transport, err := recording.NewRecordingHTTPClient(t, nil)
	require.NoError(t, err)

	err = recording.Start(t, "sdk/messaging/azwebpubsub/testdata", nil)
	require.NoError(t, err)

	// err = recording.ResetProxy(nil)
	// require.NoError(t, err)

	err = recording.AddURISanitizer(fakeTestVars.Endpoint, testVars.Endpoint, nil)
	require.NoError(t, err)

	err = recording.AddURISanitizer(fakeTestVars.Subscription, testVars.Subscription, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(`"time": "2023-10-31T00:33:32Z"`, `"time":".+?"`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"id":"00000000-0000-0000-0000-000000000000"`,
		`"id":"[^"]+"`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"lockToken":"fake-lock-token"`,
		`"lockToken":\s*"[^"]+"`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"lockTokens": ["fake-lock-token"]`,
		`"lockTokens":\s*\[\s*"[^"]+"\s*\]`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"succeededLockTokens": ["fake-lock-token"]`,
		`"succeededLockTokens":\s*\[\s*"[^"]+"\s*\]`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"succeededLockTokens": ["fake-lock-token", "fake-lock-token", "fake-lock-token"]`,
		`"succeededLockTokens":\s*`+
			`\[`+
			`(\s*"[^"]+"\s*\,){2}`+
			`\s*"[^"]+"\s*`+
			`\]`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"lockTokens": ["fake-lock-token", "fake-lock-token"]`,
		`"lockTokens":\s*\[\s*"[^"]+"\s*\,\s*"[^"]+"\s*\]`, nil)
	require.NoError(t, err)

	err = recording.AddGeneralRegexSanitizer(
		`"lockTokens": ["fake-lock-token", "fake-lock-token", "fake-lock-token"]`,
		`"lockTokens":\s*`+
			`\[`+
			`(\s*"[^"]+"\s*\,){2}`+
			`\s*"[^"]+"\s*`+
			`\]`, nil)
	require.NoError(t, err)

	t.Cleanup(func() {
		err := recording.Stop(t, nil)
		require.NoError(t, err)
	})

	return transport
}
