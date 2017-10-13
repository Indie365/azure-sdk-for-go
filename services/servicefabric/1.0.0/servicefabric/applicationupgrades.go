package servicefabric

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ApplicationUpgradesClient is the client for the ApplicationUpgrades methods of the Servicefabric service.
type ApplicationUpgradesClient struct {
	ManagementClient
}

// NewApplicationUpgradesClient creates an instance of the ApplicationUpgradesClient client.
func NewApplicationUpgradesClient(timeout *int32) ApplicationUpgradesClient {
	return NewApplicationUpgradesClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewApplicationUpgradesClientWithBaseURI creates an instance of the ApplicationUpgradesClient client.
func NewApplicationUpgradesClientWithBaseURI(baseURI string, timeout *int32) ApplicationUpgradesClient {
	return ApplicationUpgradesClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get application upgrades
//
// applicationName is the name of the application
func (client ApplicationUpgradesClient) Get(applicationName string) (result ApplicationUpgrade, err error) {
	req, err := client.GetPreparer(applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationUpgradesClient) GetPreparer(applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/GetUpgradeProgress", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationUpgradesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationUpgradesClient) GetResponder(resp *http.Response) (result ApplicationUpgrade, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Resume resume application upgrades
//
// applicationName is the name of the application resumeApplicationUpgrade is the upgrade of the resume application
func (client ApplicationUpgradesClient) Resume(applicationName string, resumeApplicationUpgrade ResumeApplicationUpgrade) (result String, err error) {
	req, err := client.ResumePreparer(applicationName, resumeApplicationUpgrade)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Resume", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResumeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Resume", resp, "Failure sending request")
		return
	}

	result, err = client.ResumeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Resume", resp, "Failure responding to request")
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client ApplicationUpgradesClient) ResumePreparer(applicationName string, resumeApplicationUpgrade ResumeApplicationUpgrade) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/MoveNextUpgradeDomain", pathParameters),
		autorest.WithJSON(resumeApplicationUpgrade),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationUpgradesClient) ResumeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client ApplicationUpgradesClient) ResumeResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start start application upgrades
//
// applicationName is the name of the application startApplicationUpgrade is the description of the start application
// upgrade
func (client ApplicationUpgradesClient) Start(applicationName string, startApplicationUpgrade StartApplicationUpgrade) (result String, err error) {
	req, err := client.StartPreparer(applicationName, startApplicationUpgrade)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Start", resp, "Failure sending request")
		return
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client ApplicationUpgradesClient) StartPreparer(applicationName string, startApplicationUpgrade StartApplicationUpgrade) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/Upgrade", pathParameters),
		autorest.WithJSON(startApplicationUpgrade),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationUpgradesClient) StartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ApplicationUpgradesClient) StartResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update application upgrades
//
// applicationName is the name of the application updateApplicationUpgrade is the description of the update application
// upgrade
func (client ApplicationUpgradesClient) Update(applicationName string, updateApplicationUpgrade UpdateApplicationUpgrade) (result String, err error) {
	req, err := client.UpdatePreparer(applicationName, updateApplicationUpgrade)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationUpgradesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ApplicationUpgradesClient) UpdatePreparer(applicationName string, updateApplicationUpgrade UpdateApplicationUpgrade) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Applications/{applicationName}/$/UpdateUpgrade", pathParameters),
		autorest.WithJSON(updateApplicationUpgrade),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationUpgradesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ApplicationUpgradesClient) UpdateResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
