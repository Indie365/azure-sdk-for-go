// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// Server is a fake server for instances of the armedgezones.Client type.
type Server struct {
	// ExtendedZonesServer contains the fakes for client ExtendedZonesClient
	ExtendedZonesServer ExtendedZonesServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armedgezones.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{srv: srv}
}

// ServerTransport connects instances of armedgezones.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                   *Server
	trMu                  sync.Mutex
	trExtendedZonesServer *ExtendedZonesServerTransport
	trOperationsServer    *OperationsServerTransport
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (s *ServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "ExtendedZonesClient":
		initServer(&s.trMu, &s.trExtendedZonesServer, func() *ExtendedZonesServerTransport {
			return NewExtendedZonesServerTransport(&s.srv.ExtendedZonesServer)
		})
		resp, err = s.trExtendedZonesServer.Do(req)
	case "OperationsClient":
		initServer(&s.trMu, &s.trOperationsServer, func() *OperationsServerTransport {
			return NewOperationsServerTransport(&s.srv.OperationsServer)
		})
		resp, err = s.trOperationsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}
