//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package arm

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pipeline"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// Endpoint is the base URL for Azure Resource Manager.
type Endpoint string

const (
	// AzureChina is the Azure Resource Manager China cloud endpoint.
	AzureChina Endpoint = "https://management.chinacloudapi.cn/"
	// AzureGermany is the Azure Resource Manager Germany cloud endpoint.
	AzureGermany Endpoint = "https://management.microsoftazure.de/"
	// AzureGovernment is the Azure Resource Manager US government cloud endpoint.
	AzureGovernment Endpoint = "https://management.usgovcloudapi.net/"
	// AzurePublicCloud is the Azure Resource Manager public cloud endpoint.
	AzurePublicCloud Endpoint = "https://management.azure.com/"
)

// ClientOptions contains configuration settings for a client's pipeline.
// All zero-value fields will be initialized with their default values.
type ClientOptions struct {
	// AuxiliaryTenants contains a list of additional tenants to be used to authenticate
	// across multiple tenants.
	AuxiliaryTenants []string

	// Host is the base URL for Azure Resource Manager
	Host Endpoint

	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient policy.Transporter

	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions

	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions

	// DisableRPRegistration disables the auto-RP registration policy.
	// The default value is false.
	DisableRPRegistration bool

	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy

	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []policy.Policy
}

// NewPipeline creates a pipeline from connection options.
// The telemetry policy, when enabled, will use the specified module and version info.
func NewPipeline(module, version string, cred azcore.TokenCredential, options *ClientOptions) pipeline.Pipeline {
	if options == nil {
		options = &ClientOptions{}
	}
	ep := options.Host
	if len(ep) == 0 {
		ep = AzurePublicCloud
	}
	policies := []policy.Policy{}
	if !options.Telemetry.Disabled {
		policies = append(policies, azruntime.NewTelemetryPolicy(module, version, &options.Telemetry))
	}
	if !options.DisableRPRegistration {
		regRPOpts := armruntime.RegistrationOptions{
			HTTPClient: options.HTTPClient,
			Logging:    options.Logging,
			Retry:      options.Retry,
			Telemetry:  options.Telemetry,
		}
		policies = append(policies, armruntime.NewRPRegistrationPolicy(string(ep), cred, &regRPOpts))
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, azruntime.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies,
		azruntime.NewBearerTokenPolicy(cred, azruntime.AuthenticationOptions{
			TokenRequest: policy.TokenRequestOptions{
				Scopes: []string{shared.EndpointToScope(string(ep))},
			},
			AuxiliaryTenants: options.AuxiliaryTenants,
		},
		),
		azruntime.NewLogPolicy(&options.Logging))
	return azruntime.NewPipeline(options.HTTPClient, policies...)
}
