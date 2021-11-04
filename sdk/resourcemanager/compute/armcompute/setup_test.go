//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcompute_test

import (
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var pathToPackage = "sdk/resourcemanager/compute/armcompute/testdata"

func TestMain(m *testing.M) {
	// Initialize
	err := recording.AddURISubscriptionIDSanitizer("00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		panic(err)
	}

	// Run
	exitVal := m.Run()

	// cleanup
	os.Exit(exitVal)
}

func startTest(t *testing.T) func() {
	err := recording.Start(t, pathToPackage, nil)
	require.NoError(t, err)
	return func() {
		err := recording.Stop(t, nil)
		require.NoError(t, err)
	}
}