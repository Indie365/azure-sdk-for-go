// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// BlobRestoreStatusPoller provides polling facilities until the operation reaches a terminal state.
type BlobRestoreStatusPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final BlobRestoreStatusResponse will be returned.
	FinalResponse(ctx context.Context) (BlobRestoreStatusResponse, error)
}

type blobRestoreStatusPoller struct {
	pt *armcore.LROPoller
}

func (p *blobRestoreStatusPoller) Done() bool {
	return p.pt.Done()
}

func (p *blobRestoreStatusPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *blobRestoreStatusPoller) FinalResponse(ctx context.Context) (BlobRestoreStatusResponse, error) {
	respType := BlobRestoreStatusResponse{BlobRestoreStatus: &BlobRestoreStatus{}}
	resp, err := p.pt.FinalResponse(ctx, respType.BlobRestoreStatus)
	if err != nil {
		return BlobRestoreStatusResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *blobRestoreStatusPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *blobRestoreStatusPoller) pollUntilDone(ctx context.Context, freq time.Duration) (BlobRestoreStatusResponse, error) {
	respType := BlobRestoreStatusResponse{BlobRestoreStatus: &BlobRestoreStatus{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.BlobRestoreStatus)
	if err != nil {
		return BlobRestoreStatusResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// HTTPPoller provides polling facilities until the operation reaches a terminal state.
type HTTPPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final *http.Response will be returned.
	FinalResponse(ctx context.Context) (*http.Response, error)
}

type httpPoller struct {
	pt *armcore.LROPoller
}

func (p *httpPoller) Done() bool {
	return p.pt.Done()
}

func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *httpPoller) FinalResponse(ctx context.Context) (*http.Response, error) {
	return p.pt.FinalResponse(ctx, nil)
}

func (p *httpPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpPoller) pollUntilDone(ctx context.Context, freq time.Duration) (*http.Response, error) {
	return p.pt.PollUntilDone(ctx, freq, nil)
}

// StorageAccountPoller provides polling facilities until the operation reaches a terminal state.
type StorageAccountPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final StorageAccountResponse will be returned.
	FinalResponse(ctx context.Context) (StorageAccountResponse, error)
}

type storageAccountPoller struct {
	pt *armcore.LROPoller
}

func (p *storageAccountPoller) Done() bool {
	return p.pt.Done()
}

func (p *storageAccountPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *storageAccountPoller) FinalResponse(ctx context.Context) (StorageAccountResponse, error) {
	respType := StorageAccountResponse{StorageAccount: &StorageAccount{}}
	resp, err := p.pt.FinalResponse(ctx, respType.StorageAccount)
	if err != nil {
		return StorageAccountResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *storageAccountPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *storageAccountPoller) pollUntilDone(ctx context.Context, freq time.Duration) (StorageAccountResponse, error) {
	respType := StorageAccountResponse{StorageAccount: &StorageAccount{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.StorageAccount)
	if err != nil {
		return StorageAccountResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
