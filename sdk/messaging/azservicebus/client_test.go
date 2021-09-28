// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/stretchr/testify/require"
)

func TestNewClientWithAzureIdentity(t *testing.T) {
	queue, cleanup := createQueue(t, getConnectionString(t), nil)
	defer cleanup()

	// test with azure identity support
	ns := os.Getenv("SERVICEBUS_FQDN")
	dac, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil || ns == "" {
		t.Skip("Azure Identity compatible credentials not configured")
	}

	client, err := NewClient(ns, dac)
	require.NoError(t, err)

	sender, err := client.NewSender(queue)
	require.NoError(t, err)

	err = sender.SendMessage(context.TODO(), &Message{Body: []byte("hello - authenticating with a TokenCredential")})
	require.NoError(t, err)

	receiver, err := client.NewReceiver(ReceiverWithQueue(queue))
	require.NoError(t, err)
	receiver.settler.onlyDoBackupSettlement = true // this'll also exercise the management link

	messages, err := receiver.ReceiveMessages(context.TODO(), 1)
	require.NoError(t, err)

	require.EqualValues(t, []string{"hello - authenticating with a TokenCredential"}, getSortedBodies(messages))

	for _, m := range messages {
		err = receiver.CompleteMessage(context.TODO(), m)
		require.NoError(t, err)
	}

	client.Close(context.TODO())
}

func TestNewClientUnitTests(t *testing.T) {
	t.Run("WithTokenCredential", func(t *testing.T) {
		fakeTokenCredential := struct{ azcore.TokenCredential }{}

		client, err := NewClient("fake.something", fakeTokenCredential)
		require.NoError(t, err)

		require.NoError(t, err)
		require.EqualValues(t, fakeTokenCredential, client.config.tokenCredential)
		require.EqualValues(t, "fake.something", client.config.fullyQualifiedNamespace)

		client, err = NewClient("mysb.windows.servicebus.net", fakeTokenCredential)
		require.NoError(t, err)
		require.EqualValues(t, fakeTokenCredential, client.config.tokenCredential)
		require.EqualValues(t, "mysb.windows.servicebus.net", client.config.fullyQualifiedNamespace)

		// (really all part of the same functionality)
		ns := &internal.Namespace{}
		require.NoError(t, internal.NamespacesWithTokenCredential("mysb.windows.servicebus.net",
			fakeTokenCredential)(ns))

		require.EqualValues(t, ns.Name, "mysb")
		require.EqualValues(t, ns.Suffix, "windows.servicebus.net")
	})
}
