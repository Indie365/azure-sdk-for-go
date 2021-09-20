//go:build go1.16
// +build go1.16

package azsecrets

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

var pathToPackage = "sdk/keyvault/azsecrets"

func createRandomName(t *testing.T, prefix string) (string, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(t.Name()))
	return prefix + fmt.Sprint(h.Sum32()), err
}

type recordingPolicy struct {
	options recording.RecordingOptions
	t       *testing.T
}

func NewRecordingPolicy(t *testing.T, o *recording.RecordingOptions) policy.Policy {
	if o == nil {
		o = &recording.RecordingOptions{}
	}
	p := &recordingPolicy{options: *o, t: t}
	p.options.Init()
	return p
}

func (p *recordingPolicy) Do(req *policy.Request) (resp *http.Response, err error) {
	originalURLHost := req.Raw().URL.Host
	req.Raw().URL.Scheme = "https"
	req.Raw().URL.Host = p.options.Host
	req.Raw().Host = p.options.Host

	req.Raw().Header.Set(recording.UpstreamUriHeader, fmt.Sprintf("%v://%v", p.options.Scheme, originalURLHost))
	req.Raw().Header.Set(recording.ModeHeader, recording.GetRecordMode())
	req.Raw().Header.Set(recording.IdHeader, recording.GetRecordingId(p.t))

	return req.Next()
}

func createClient(t *testing.T) (*Client, error) {
	vaultUrl, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		t.Fatal("Could not find environment variable AZURE_KEYVAULT_URL")
	}

	p := NewRecordingPolicy(t, &recording.RecordingOptions{UseHTTPS: true})
	client, err := recording.GetHTTPClient(t)
	require.NoError(t, err)

	options := &ClientOptions{
		PerCallPolicies: []policy.Policy{p},
		HTTPClient:      client,
	}
	_ = options

	cred, err := azidentity.NewClientSecretCredential(
		os.Getenv("KEYVAULT_TENANT_ID"),
		os.Getenv("KEYVAULT_CLIENT_ID"),
		os.Getenv("KEYVAULT_CLIENT_SECRET"),
		nil,
	)
	require.NoError(t, err)

	return NewClient(vaultUrl, cred, options)
}

func cleanUpSecret(t *testing.T, client *Client, secret string) {
	resp, err := client.BeginDeleteSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	_, err = client.PurgeDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)
}

func TestSetGetSecret(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secret")
	require.NoError(t, err)
	value, err := createRandomName(t, "value")
	require.NoError(t, err)

	defer cleanUpSecret(t, client, secret)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	getResp, err := client.GetSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	require.Equal(t, *getResp.Value, value)
}

func TestListSecretVersionss(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secret")
	require.NoError(t, err)
	value, err := createRandomName(t, "value")
	require.NoError(t, err)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)
	_, err = client.SetSecret(context.Background(), secret, value+"1", nil)
	require.NoError(t, err)
	_, err = client.SetSecret(context.Background(), secret, value+"2", nil)
	require.NoError(t, err)

	count := 0
	pager := client.ListSecretVersions(secret, nil)
	for pager.NextPage(context.Background()) {
		page := pager.PageResponse()
		count += len(page.Secrets)
	}
	require.GreaterOrEqual(t, count, 3)
	require.NoError(t, pager.Err())

	// clean up test
	resp, err := client.BeginDeleteSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	_, err = client.PurgeDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)
}

func TestListSecrets(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	count := 0
	pager := client.ListSecrets(nil)
	for pager.NextPage(context.Background()) {
		page := pager.PageResponse()
		count += len(page.Secrets)
	}
	require.Greater(t, count, 0)
	require.NoError(t, pager.Err())
}

func TestListDeletedSecrets(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret1, err := createRandomName(t, "secret1")
	require.NoError(t, err)
	value1, err := createRandomName(t, "value1")
	require.NoError(t, err)
	secret2, err := createRandomName(t, "secret2")
	require.NoError(t, err)
	value2, err := createRandomName(t, "value2")
	require.NoError(t, err)

	f := func() {
		_, err := client.PurgeDeletedSecret(context.Background(), secret1, nil)
		require.NoError(t, err)
		_, err = client.PurgeDeletedSecret(context.Background(), secret2, nil)
		require.NoError(t, err)
	}
	defer f()

	// 1. Create 2 secrets
	_, err = client.SetSecret(context.Background(), secret1, value1, nil)
	require.NoError(t, err)

	_, err = client.SetSecret(context.Background(), secret2, value2, nil)
	require.NoError(t, err)

	// 2. Delete both secrets
	resp, err := client.BeginDeleteSecret(context.Background(), secret1, nil)
	require.NoError(t, err)
	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	resp, err = client.BeginDeleteSecret(context.Background(), secret2, nil)
	require.NoError(t, err)
	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	// Make sure both secrets show up in deleted secrets
	deletedSecrets := map[string]bool{
		secret1: false,
		secret2: false,
	}
	count := 0
	pager := client.ListDeletedSecrets(nil)
	for pager.NextPage(context.Background()) {
		page := pager.PageResponse()
		count += len(page.DeletedSecrets)
		for _, secret := range page.DeletedSecrets {
			for deleted := range deletedSecrets {
				if strings.Contains(*secret.ID, deleted) {
					deletedSecrets[deleted] = true
					break
				}
			}
		}
	}

	for _, deleted := range deletedSecrets {
		require.True(t, deleted)
	}
}

func TestDeleteSecret(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secret1")
	require.NoError(t, err)
	value, err := createRandomName(t, "value1")
	require.NoError(t, err)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	resp, err := client.BeginDeleteSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	_, err = client.GetDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = client.PurgeDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)
}

func TestPurgeDeletedSecret(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secret1")
	require.NoError(t, err)
	value, err := createRandomName(t, "value1")
	require.NoError(t, err)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	resp, err := client.BeginDeleteSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	_, err = client.PurgeDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	pager := client.ListDeletedSecrets(nil)
	for pager.NextPage(context.Background()) {
		page := pager.PageResponse()
		for _, secret := range page.DeletedSecrets {
			require.NotEqual(t, *secret.ID, secret)
		}
	}
}

func TestUpdateSecretProperties(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secret")
	require.NoError(t, err)
	value, err := createRandomName(t, "value")
	require.NoError(t, err)

	defer cleanUpSecret(t, client, secret)

	resp, err := client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	secretVersion := strings.Split(*resp.ID, "/")

	getResp, err := client.GetSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	require.Equal(t, *getResp.Value, value)

	params := SecretProperties{
		ContentType: to.StringPtr("password"),
		Tags: map[string]*string{
			"Tag1": to.StringPtr("TagVal1"),
		},
	}

	_, err = client.UpdateSecretProperties(context.Background(), secret, secretVersion[len(secretVersion)-1], params, nil)
	require.NoError(t, err)

	getResp, err = client.GetSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	require.Equal(t, *getResp.Value, value)
	require.Equal(t, *getResp.Tags["Tag1"], "TagVal1")
	require.Equal(t, *getResp.ContentType, "password")
}

func TestBeginRecoverDeletedSecret(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secret")
	require.NoError(t, err)
	value, err := createRandomName(t, "value")
	require.NoError(t, err)

	defer cleanUpSecret(t, client, secret)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	pollerResp, err := client.BeginDeleteSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = pollerResp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	resp, err := client.BeginRecoverDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = resp.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	getResp, err := client.GetSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	require.Equal(t, *getResp.Value, value)
}

func TestBackupSecret(t *testing.T) {
	recording.StartRecording(t, pathToPackage, nil)
	defer recording.StopRecording(t, nil)

	client, err := createClient(t)
	require.NoError(t, err)

	secret, err := createRandomName(t, "secrets")
	require.NoError(t, err)
	value, err := createRandomName(t, "value")
	require.NoError(t, err)

	_, err = client.SetSecret(context.Background(), secret, value, nil)
	require.NoError(t, err)

	backupResp, err := client.BackupSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	require.Greater(t, len(backupResp.Value), 0)

	respPoller, err := client.BeginDeleteSecret(context.Background(), secret, nil)
	require.NoError(t, err)
	_, err = respPoller.PollUntilDone(context.Background(), 100*time.Millisecond)
	require.NoError(t, err)

	_, err = client.PurgeDeletedSecret(context.Background(), secret, nil)
	require.NoError(t, err)

	_, err = client.GetSecret(context.Background(), secret, nil)
	var httpErr azcore.HTTPResponse
	require.True(t, errors.As(err, &httpErr))
	require.Equal(t, httpErr.RawResponse().StatusCode, http.StatusNotFound)

	_, err = client.GetDeletedSecret(context.Background(), secret, nil)
	require.True(t, errors.As(err, &httpErr))
	require.Equal(t, httpErr.RawResponse().StatusCode, http.StatusNotFound)

	// Poll this operation manually
	var restoreResp RestoreSecretBackupResponse
	var i int
	for i = 0; i < 10; i++ {
		restoreResp, err = client.RestoreSecretBackup(context.Background(), backupResp.Value, nil)
		if err == nil {
			break
		}
	}
	require.NoError(t, err)
	require.Contains(t, *restoreResp.ID, secret)

	// Now the Secret should be Get-able
	_, err = client.GetSecret(context.Background(), secret, nil)
	require.NoError(t, err)
}
