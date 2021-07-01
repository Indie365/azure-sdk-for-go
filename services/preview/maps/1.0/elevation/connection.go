// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package elevation

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

var scopes = []string{"https://atlas.microsoft.com/.default"}

// ConnectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type ConnectionOptions struct {
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

func (c *ConnectionOptions) telemetryOptions() *azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return &to
}

// Connection - The Azure Maps Elevation API provides an HTTP interface to query elevation data on the  surface of the Earth. Elevation data can be retrieved at specific locations by sending  lat/lon coordinates, by defining an ordered set of vertices that form a Polyline and a  number of sample points along the length of a Polyline, or by defining a bounding box  that consists of equally spaced vertices as rows and columns. The vertical datum is EPSG:3855.  This datum uses the EGM2008 geoid model applied to the WGS84 ellipsoid as its zero height  reference surface. The vertical unit is measured in meters, the spatial resolution of the  elevation data is 0.8 arc-second for global coverage (~24 meters).
type Connection struct {
	u string
	p azcore.Pipeline
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewConnection(geography *Geography, cred azcore.Credential, options *ConnectionOptions) *Connection {
	if options == nil {
		options = &ConnectionOptions{}
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
		geography = ((*Geography)(&defaultValue))
	}
	hostURL = strings.ReplaceAll(hostURL, "{geography}", (string)(*geography))
	return &Connection{u: hostURL, p: azcore.NewPipeline(options.HTTPClient, policies...)}
}

// Endpoint returns the connection's endpoint.
func (c *Connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() azcore.Pipeline {
	return c.p
}
