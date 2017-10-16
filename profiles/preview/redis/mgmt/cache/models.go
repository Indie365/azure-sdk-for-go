// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 714052a3359963ba46703fe033cf9dd4838bea11

package cache

import original "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2017-02-01/cache"

type DayOfWeek = original.DayOfWeek

const (
	Everyday	DayOfWeek	= original.Everyday
	Friday		DayOfWeek	= original.Friday
	Monday		DayOfWeek	= original.Monday
	Saturday	DayOfWeek	= original.Saturday
	Sunday		DayOfWeek	= original.Sunday
	Thursday	DayOfWeek	= original.Thursday
	Tuesday		DayOfWeek	= original.Tuesday
	Wednesday	DayOfWeek	= original.Wednesday
	Weekend		DayOfWeek	= original.Weekend
)

type RebootType = original.RebootType

const (
	AllNodes	RebootType	= original.AllNodes
	PrimaryNode	RebootType	= original.PrimaryNode
	SecondaryNode	RebootType	= original.SecondaryNode
)

type RedisKeyType = original.RedisKeyType

const (
	Primary		RedisKeyType	= original.Primary
	Secondary	RedisKeyType	= original.Secondary
)

type ReplicationRole = original.ReplicationRole

const (
	ReplicationRolePrimary		ReplicationRole	= original.ReplicationRolePrimary
	ReplicationRoleSecondary	ReplicationRole	= original.ReplicationRoleSecondary
)

type SkuFamily = original.SkuFamily

const (
	C	SkuFamily	= original.C
	P	SkuFamily	= original.P
)

type SkuName = original.SkuName

const (
	Basic		SkuName	= original.Basic
	Premium		SkuName	= original.Premium
	Standard	SkuName	= original.Standard
)

type ExportRDBParameters = original.ExportRDBParameters
type ImportRDBParameters = original.ImportRDBParameters
type RedisAccessKeys = original.RedisAccessKeys
type RedisCreateParameters = original.RedisCreateParameters
type RedisCreateProperties = original.RedisCreateProperties
type RedisForceRebootResponse = original.RedisForceRebootResponse
type RedisLinkedServer = original.RedisLinkedServer
type RedisLinkedServerCreateParameters = original.RedisLinkedServerCreateParameters
type RedisLinkedServerCreateProperties = original.RedisLinkedServerCreateProperties
type RedisLinkedServerList = original.RedisLinkedServerList
type RedisLinkedServerProperties = original.RedisLinkedServerProperties
type RedisLinkedServerWithProperties = original.RedisLinkedServerWithProperties
type RedisLinkedServerWithPropertiesList = original.RedisLinkedServerWithPropertiesList
type RedisListResult = original.RedisListResult
type RedisPatchSchedule = original.RedisPatchSchedule
type RedisProperties = original.RedisProperties
type RedisRebootParameters = original.RedisRebootParameters
type RedisRegenerateKeyParameters = original.RedisRegenerateKeyParameters
type RedisResource = original.RedisResource
type RedisResourceProperties = original.RedisResourceProperties
type RedisUpdateParameters = original.RedisUpdateParameters
type RedisUpdateProperties = original.RedisUpdateProperties
type Resource = original.Resource
type ScheduleEntries = original.ScheduleEntries
type ScheduleEntry = original.ScheduleEntry
type Sku = original.Sku
type PatchSchedulesClient = original.PatchSchedulesClient
type RedisClient = original.RedisClient
type RedisLinkedServerClient = original.RedisLinkedServerClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewPatchSchedulesClient(subscriptionID string) PatchSchedulesClient {
	return original.NewPatchSchedulesClient(subscriptionID)
}
func NewPatchSchedulesClientWithBaseURI(baseURI string, subscriptionID string) PatchSchedulesClient {
	return original.NewPatchSchedulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRedisClient(subscriptionID string) RedisClient {
	return original.NewRedisClient(subscriptionID)
}
func NewRedisClientWithBaseURI(baseURI string, subscriptionID string) RedisClient {
	return original.NewRedisClientWithBaseURI(baseURI, subscriptionID)
}
func NewRedisLinkedServerClient(subscriptionID string) RedisLinkedServerClient {
	return original.NewRedisLinkedServerClient(subscriptionID)
}
func NewRedisLinkedServerClientWithBaseURI(baseURI string, subscriptionID string) RedisLinkedServerClient {
	return original.NewRedisLinkedServerClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
