// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztables

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	generated "github.com/Azure/azure-sdk-for-go/sdk/data/aztables/internal"
)

type ClientOptions struct {
	// Transporter sets the transport for making HTTP requests.
	Transporter policy.Transporter

	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions

	// PerCallPolicies are policies that execute once per operation
	PerCallPolicies []policy.Policy

	// PerTryPolicies are options to run on every retry
	PerTryPolicies []policy.Policy
}

func (o *ClientOptions) getConnectionOptions() *generated.ConnectionOptions {
	if o == nil {
		return nil
	}

	return &generated.ConnectionOptions{
		HTTPClient:       o.Transporter,
		Retry:            o.Retry,
		Telemetry:        o.Telemetry,
		PerCallPolicies:  []policy.Policy{},
		PerRetryPolicies: []policy.Policy{},
	}
}
