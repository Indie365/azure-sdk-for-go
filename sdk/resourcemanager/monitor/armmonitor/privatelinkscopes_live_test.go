//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/stretchr/testify/suite"
)

type PrivatelinkscopesTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *PrivatelinkscopesTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.location = testutil.GetEnv("LOCATION", "eastus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/monitor/armmonitor/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
}

func (testsuite *PrivatelinkscopesTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestPrivatelinkscopesTestSuite(t *testing.T) {
	suite.Run(t, new(PrivatelinkscopesTestSuite))
}

// microsoft.insights/privateLinkScopes
func (testsuite *PrivatelinkscopesTestSuite) TestPrivatelinkscope() {
	scopeName := testutil.GenerateAlphaNumericID(testsuite.T(), "linkscopena", 6)
	var err error
	// From step PrivateLinkScopes_Create
	privateLinkScopesClient, err := armmonitor.NewPrivateLinkScopesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = privateLinkScopesClient.CreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, scopeName, armmonitor.AzureMonitorPrivateLinkScope{
		Location: to.Ptr("Global"),
	}, nil)
	testsuite.Require().NoError(err)

	// From step PrivateLinkScopes_List
	privateLinkScopesClientNewListPager := privateLinkScopesClient.NewListPager(nil)
	for privateLinkScopesClientNewListPager.More() {
		_, err := privateLinkScopesClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step PrivateLinkScopes_ListByResourceGroup
	privateLinkScopesClientNewListByResourceGroupPager := privateLinkScopesClient.NewListByResourceGroupPager(testsuite.resourceGroupName, nil)
	for privateLinkScopesClientNewListByResourceGroupPager.More() {
		_, err := privateLinkScopesClientNewListByResourceGroupPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step PrivateLinkScopes_Get
	_, err = privateLinkScopesClient.Get(testsuite.ctx, testsuite.resourceGroupName, scopeName, nil)
	testsuite.Require().NoError(err)

	// From step PrivateLinkScopes_UpdateTags
	_, err = privateLinkScopesClient.UpdateTags(testsuite.ctx, testsuite.resourceGroupName, scopeName, armmonitor.TagsResource{
		Tags: map[string]*string{
			"Tag1": to.Ptr("Value1"),
			"Tag2": to.Ptr("Value2"),
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step PrivateLinkScopes_Delete
	privateLinkScopesClientDeleteResponsePoller, err := privateLinkScopesClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, scopeName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, privateLinkScopesClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}
