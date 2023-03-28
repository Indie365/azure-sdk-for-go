// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ../namespace.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	amqpwrap "github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/amqpwrap"
	gomock "github.com/golang/mock/gomock"
)

// MockNamespaceWithNewAMQPLinks is a mock of NamespaceWithNewAMQPLinks interface.
type MockNamespaceWithNewAMQPLinks struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceWithNewAMQPLinksMockRecorder
}

// MockNamespaceWithNewAMQPLinksMockRecorder is the mock recorder for MockNamespaceWithNewAMQPLinks.
type MockNamespaceWithNewAMQPLinksMockRecorder struct {
	mock *MockNamespaceWithNewAMQPLinks
}

// NewMockNamespaceWithNewAMQPLinks creates a new mock instance.
func NewMockNamespaceWithNewAMQPLinks(ctrl *gomock.Controller) *MockNamespaceWithNewAMQPLinks {
	mock := &MockNamespaceWithNewAMQPLinks{ctrl: ctrl}
	mock.recorder = &MockNamespaceWithNewAMQPLinksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceWithNewAMQPLinks) EXPECT() *MockNamespaceWithNewAMQPLinksMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockNamespaceWithNewAMQPLinks) Check() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockNamespaceWithNewAMQPLinksMockRecorder) Check() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockNamespaceWithNewAMQPLinks)(nil).Check))
}

// MockNamespaceForAMQPLinks is a mock of NamespaceForAMQPLinks interface.
type MockNamespaceForAMQPLinks struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceForAMQPLinksMockRecorder
}

// MockNamespaceForAMQPLinksMockRecorder is the mock recorder for MockNamespaceForAMQPLinks.
type MockNamespaceForAMQPLinksMockRecorder struct {
	mock *MockNamespaceForAMQPLinks
}

// NewMockNamespaceForAMQPLinks creates a new mock instance.
func NewMockNamespaceForAMQPLinks(ctrl *gomock.Controller) *MockNamespaceForAMQPLinks {
	mock := &MockNamespaceForAMQPLinks{ctrl: ctrl}
	mock.recorder = &MockNamespaceForAMQPLinksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceForAMQPLinks) EXPECT() *MockNamespaceForAMQPLinksMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockNamespaceForAMQPLinks) Close(ctx context.Context, permanently bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx, permanently)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockNamespaceForAMQPLinksMockRecorder) Close(ctx, permanently any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNamespaceForAMQPLinks)(nil).Close), ctx, permanently)
}

// GetEntityAudience mocks base method.
func (m *MockNamespaceForAMQPLinks) GetEntityAudience(entityPath string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntityAudience", entityPath)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEntityAudience indicates an expected call of GetEntityAudience.
func (mr *MockNamespaceForAMQPLinksMockRecorder) GetEntityAudience(entityPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntityAudience", reflect.TypeOf((*MockNamespaceForAMQPLinks)(nil).GetEntityAudience), entityPath)
}

// NegotiateClaim mocks base method.
func (m *MockNamespaceForAMQPLinks) NegotiateClaim(ctx context.Context, entityPath string) (context.CancelFunc, <-chan struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NegotiateClaim", ctx, entityPath)
	ret0, _ := ret[0].(context.CancelFunc)
	ret1, _ := ret[1].(<-chan struct{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NegotiateClaim indicates an expected call of NegotiateClaim.
func (mr *MockNamespaceForAMQPLinksMockRecorder) NegotiateClaim(ctx, entityPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NegotiateClaim", reflect.TypeOf((*MockNamespaceForAMQPLinks)(nil).NegotiateClaim), ctx, entityPath)
}

// NewAMQPSession mocks base method.
func (m *MockNamespaceForAMQPLinks) NewAMQPSession(ctx context.Context) (amqpwrap.AMQPSession, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAMQPSession", ctx)
	ret0, _ := ret[0].(amqpwrap.AMQPSession)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewAMQPSession indicates an expected call of NewAMQPSession.
func (mr *MockNamespaceForAMQPLinksMockRecorder) NewAMQPSession(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAMQPSession", reflect.TypeOf((*MockNamespaceForAMQPLinks)(nil).NewAMQPSession), ctx)
}

// NewRPCLink mocks base method.
func (m *MockNamespaceForAMQPLinks) NewRPCLink(ctx context.Context, managementPath string) (amqpwrap.RPCLink, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRPCLink", ctx, managementPath)
	ret0, _ := ret[0].(amqpwrap.RPCLink)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewRPCLink indicates an expected call of NewRPCLink.
func (mr *MockNamespaceForAMQPLinksMockRecorder) NewRPCLink(ctx, managementPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRPCLink", reflect.TypeOf((*MockNamespaceForAMQPLinks)(nil).NewRPCLink), ctx, managementPath)
}

// Recover mocks base method.
func (m *MockNamespaceForAMQPLinks) Recover(ctx context.Context, clientRevision uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recover", ctx, clientRevision)
	ret0, _ := ret[0].(error)
	return ret0
}

// Recover indicates an expected call of Recover.
func (mr *MockNamespaceForAMQPLinksMockRecorder) Recover(ctx, clientRevision any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recover", reflect.TypeOf((*MockNamespaceForAMQPLinks)(nil).Recover), ctx, clientRevision)
}
