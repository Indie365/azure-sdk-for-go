// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package route

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"strings"
)

var scopes = []string{"https://atlas.microsoft.com/.default"}
// connectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type connectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging azcore.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []azcore.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []azcore.Policy
}

func (c *connectionOptions) telemetryOptions() *azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return &to
}

// connection - Azure Maps Route REST APIs
type connection struct {
	u string
	p azcore.Pipeline
}

// newConnection creates an instance of the connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newConnection(geography *Geography, cred azcore.Credential, options *connectionOptions) *connection {
	if options == nil {
		options = &connectionOptions{}
	}
	policies := []azcore.Policy{
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, azcore.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
		policies = append(policies, cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: scopes}}))
	policies = append(policies, azcore.NewLogPolicy(&options.Logging))
	hostURL := "https://{geography}.atlas.microsoft.com"
	if geography == nil {
		defaultValue := "us"
		geography = &defaultValue
	}
	hostURL = strings.ReplaceAll(hostURL, "{geography}", *geography)
	return &connection{u: hostURL, p: azcore.NewPipeline(options.HTTPClient, policies...)}
}

// Endpoint returns the connection's endpoint.
func (c *connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *connection) Pipeline() (azcore.Pipeline) {
	return c.p
}

