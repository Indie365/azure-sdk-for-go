package containerinstance

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/azure-pipeline-go/pipeline"
	"net/url"
)

const (
	// ServiceVersion specifies the version of the operations used in this package.
	ServiceVersion = "2017-08-01-preview"
	// DefaultBaseURL is the default URL used for the service Containerinstance
	DefaultBaseURL = "https://management.azure.com"
)

// ManagementClient is the base client for Containerinstance.
type ManagementClient struct {
	url url.URL
	p   pipeline.Pipeline
}

// NewManagementClient creates an instance of the ManagementClient client.
func NewManagementClient(p pipeline.Pipeline) ManagementClient {
	u, err := url.Parse(DefaultBaseURL)
	if err != nil {
		panic(err)
	}
	return NewManagementClientWithURL(*u, p)
}

// NewManagementClientWithURL creates an instance of the ManagementClient client.
func NewManagementClientWithURL(url url.URL, p pipeline.Pipeline) ManagementClient {
	return ManagementClient{
		url: url,
		p:   p,
	}
}

// URL returns a copy of the URL for this client.
func (mc ManagementClient) URL() url.URL {
	return mc.url
}

// Pipeline returns the pipeline for this client.
func (mc ManagementClient) Pipeline() pipeline.Pipeline {
	return mc.p
}
