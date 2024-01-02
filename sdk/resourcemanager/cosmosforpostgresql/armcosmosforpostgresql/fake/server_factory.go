//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armcosmosforpostgresql.ClientFactory type.
type ServerFactory struct {
	ClustersServer                   ClustersServer
	ConfigurationsServer             ConfigurationsServer
	FirewallRulesServer              FirewallRulesServer
	OperationsServer                 OperationsServer
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer       PrivateLinkResourcesServer
	RolesServer                      RolesServer
	ServersServer                    ServersServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcosmosforpostgresql.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcosmosforpostgresql.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trClustersServer                   *ClustersServerTransport
	trConfigurationsServer             *ConfigurationsServerTransport
	trFirewallRulesServer              *FirewallRulesServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
	trRolesServer                      *RolesServerTransport
	trServersServer                    *ServersServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ClustersClient":
		initServer(s, &s.trClustersServer, func() *ClustersServerTransport { return NewClustersServerTransport(&s.srv.ClustersServer) })
		resp, err = s.trClustersServer.Do(req)
	case "ConfigurationsClient":
		initServer(s, &s.trConfigurationsServer, func() *ConfigurationsServerTransport {
			return NewConfigurationsServerTransport(&s.srv.ConfigurationsServer)
		})
		resp, err = s.trConfigurationsServer.Do(req)
	case "FirewallRulesClient":
		initServer(s, &s.trFirewallRulesServer, func() *FirewallRulesServerTransport {
			return NewFirewallRulesServerTransport(&s.srv.FirewallRulesServer)
		})
		resp, err = s.trFirewallRulesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "RolesClient":
		initServer(s, &s.trRolesServer, func() *RolesServerTransport { return NewRolesServerTransport(&s.srv.RolesServer) })
		resp, err = s.trRolesServer.Do(req)
	case "ServersClient":
		initServer(s, &s.trServersServer, func() *ServersServerTransport { return NewServersServerTransport(&s.srv.ServersServer) })
		resp, err = s.trServersServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
