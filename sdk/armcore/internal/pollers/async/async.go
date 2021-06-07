// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package async

import (
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/armcore/internal/pollers"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const (
	finalStateAsync = "azure-async-operation"
	finalStateLoc   = "location"
	finalStateOrig  = "original-uri"
)

// Applicable returns true if the LRO is using Azure-AsyncOperation.
func Applicable(resp *azcore.Response) bool {
	return resp.Header.Get(pollers.HeaderAzureAsync) != ""
}

// Poller is an LRO poller that uses the Azure-AsyncOperation pattern.
type Poller struct {
	Type       string `json:"type"`
	AsyncURL   string `json:"asyncURL"`
	LocURL     string `json:"locURL"`
	OrigURL    string `json:"origURL"`
	Method     string `json:"method"`
	FinalState string `json:"finalState"`
	CurState   string `json:"state"`
}

// New creates a new Poller from the provided initial response and final-state type.
func New(resp *azcore.Response, finalState string, pollerID string) (*Poller, error) {
	azcore.Log().Write(azcore.LogLongRunningOperation, "Using Azure-AsyncOperation poller.")
	asyncURL := resp.Header.Get(pollers.HeaderAzureAsync)
	if asyncURL == "" {
		return nil, errors.New("response is missing Azure-AsyncOperation header")
	}
	p := &Poller{
		Type:       pollers.MakeID(pollerID, "async"),
		AsyncURL:   asyncURL,
		LocURL:     resp.Header.Get(pollers.HeaderLocation),
		OrigURL:    resp.Request.URL.String(),
		Method:     resp.Request.Method,
		FinalState: finalState,
	}
	// check for provisioning state
	state, err := pollers.GetProvisioningState(resp)
	if errors.Is(err, pollers.ErrNoProvisioningState) {
		if resp.Request.Method == http.MethodPut {
			// initial response for a PUT requires a provisioning state
			return nil, err
		}
		// for DELETE/PATCH/POST, provisioning state is optional
		state = "InProgress"
	} else if err != nil {
		return nil, err
	}
	p.CurState = state
	return p, nil
}

// Done returns true if the LRO has reached a terminal state.
func (p *Poller) Done() bool {
	return pollers.IsTerminalState(p.Status())
}

// Update updates the Poller from the polling response.
func (p *Poller) Update(resp *azcore.Response) error {
	state, err := pollers.GetStatus(resp)
	if err != nil {
		return err
	}
	p.CurState = state
	return nil
}

// FinalGetURL returns the URL to perform a final GET for the payload, or the empty string if not required.
func (p *Poller) FinalGetURL() string {
	if p.Method == http.MethodPatch || p.Method == http.MethodPut {
		// for PATCH and PUT, the final GET is on the original resource URL
		return p.OrigURL
	} else if p.Method == http.MethodPost {
		// for POST, we need to consult the final-state-via flag
		if p.FinalState == finalStateLoc && p.LocURL != "" {
			return p.LocURL
		} else if p.FinalState == finalStateOrig {
			return p.OrigURL
		}
		// finalStateAsync fall through
	}
	return ""
}

// URL returns the polling URL.
func (p *Poller) URL() string {
	return p.AsyncURL
}

// Status returns the status of the LRO.
func (p *Poller) Status() string {
	return p.CurState
}
