// Copyright 2023 Peridot Authors
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

package base

import (
	"context"
	"errors"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type GRPCServer struct {
	server     *grpc.Server
	gatewayMux *runtime.ServeMux

	gatewayClientConn *grpc.ClientConn

	serverOptions      []grpc.ServerOption
	dialOptions        []grpc.DialOption
	muxOptions         []runtime.ServeMuxOption
	outgoingHeaders    []string
	incomingHeaders    []string
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	timeout            time.Duration
	grpcPort           int
	gatewayPort        int
	noGrpcGateway      bool
}

type GrpcEndpointRegister func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
type GRPCServerOption func(*GRPCServer)

// WithServerOptions sets the gRPC server options. (Append)
func WithServerOptions(opts ...grpc.ServerOption) GRPCServerOption {
	return func(g *GRPCServer) {
		g.serverOptions = append(g.serverOptions, opts...)
	}
}

// WithDialOptions sets the gRPC dial options. (Append)
func WithDialOptions(opts ...grpc.DialOption) GRPCServerOption {
	return func(g *GRPCServer) {
		g.dialOptions = append(g.dialOptions, opts...)
	}
}

// WithMuxOptions sets the gRPC-gateway mux options. (Append)
func WithMuxOptions(opts ...runtime.ServeMuxOption) GRPCServerOption {
	return func(g *GRPCServer) {
		g.muxOptions = append(g.muxOptions, opts...)
	}
}

// WithOutgoingHeaders sets the outgoing headers for the gRPC-gateway. (Append)
func WithOutgoingHeaders(headers ...string) GRPCServerOption {
	return func(g *GRPCServer) {
		g.outgoingHeaders = append(g.outgoingHeaders, headers...)
	}
}

// WithIncomingHeaders sets the incoming headers for the gRPC-gateway. (Append)
func WithIncomingHeaders(headers ...string) GRPCServerOption {
	return func(g *GRPCServer) {
		g.incomingHeaders = append(g.incomingHeaders, headers...)
	}
}

// WithUnaryInterceptors sets the gRPC unary interceptors. (Append)
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) GRPCServerOption {
	return func(g *GRPCServer) {
		g.unaryInterceptors = append(g.unaryInterceptors, interceptors...)
	}
}

// WithStreamInterceptors sets the gRPC stream interceptors. (Append)
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) GRPCServerOption {
	return func(g *GRPCServer) {
		g.streamInterceptors = append(g.streamInterceptors, interceptors...)
	}
}

// WithTimeout sets the timeout for the gRPC server.
func WithTimeout(timeout time.Duration) GRPCServerOption {
	return func(g *GRPCServer) {
		g.timeout = timeout
	}
}

// WithGRPCPort sets the gRPC port for the gRPC server.
func WithGRPCPort(port int) GRPCServerOption {
	return func(g *GRPCServer) {
		g.grpcPort = port
	}
}

// WithGatewayPort sets the gRPC-gateway port for the gRPC server.
func WithGatewayPort(port int) GRPCServerOption {
	return func(g *GRPCServer) {
		g.gatewayPort = port
	}
}

// WithNoGRPCGateway disables the gRPC-gateway for the gRPC server.
func WithNoGRPCGateway() GRPCServerOption {
	return func(g *GRPCServer) {
		g.noGrpcGateway = true
	}
}

// NewGRPCServer creates a new gRPC-server with gRPC-gateway, default interceptors
// and exposed Prometheus metrics.
func NewGRPCServer(opts ...GRPCServerOption) (*GRPCServer, error) {
	g := &GRPCServer{
		serverOptions:      []grpc.ServerOption{},
		dialOptions:        []grpc.DialOption{},
		muxOptions:         []runtime.ServeMuxOption{},
		outgoingHeaders:    []string{},
		incomingHeaders:    []string{},
		unaryInterceptors:  []grpc.UnaryServerInterceptor{},
		streamInterceptors: []grpc.StreamServerInterceptor{},
	}

	// Apply options first
	for _, opt := range opts {
		opt(g)
	}

	// Set defaults
	if g.timeout == 0 {
		g.timeout = 10 * time.Second
	}
	if g.grpcPort == 0 {
		g.grpcPort = 8080
	}
	if g.gatewayPort == 0 {
		g.gatewayPort = g.grpcPort + 1
	}

	// Always prepend the insecure dial option
	// RESF deploys with Istio, which handles mTLS
	g.dialOptions = append([]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}, g.dialOptions...)

	// Set default interceptors
	if g.unaryInterceptors == nil {
		g.unaryInterceptors = []grpc.UnaryServerInterceptor{}
	}
	if g.streamInterceptors == nil {
		g.streamInterceptors = []grpc.StreamServerInterceptor{}
	}

	// Always prepend the prometheus interceptor
	g.unaryInterceptors = append([]grpc.UnaryServerInterceptor{grpc_prometheus.UnaryServerInterceptor}, g.unaryInterceptors...)
	g.streamInterceptors = append([]grpc.StreamServerInterceptor{grpc_prometheus.StreamServerInterceptor}, g.streamInterceptors...)

	// Chain the interceptors
	g.serverOptions = append(g.serverOptions, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(g.unaryInterceptors...)))
	g.serverOptions = append(g.serverOptions, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(g.streamInterceptors...)))

	g.server = grpc.NewServer(g.serverOptions...)

	if !g.noGrpcGateway {
		g.gatewayMux = runtime.NewServeMux(g.muxOptions...)

		// Create gateway client connection
		var err error
		g.gatewayClientConn, err = grpc.Dial("localhost:"+strconv.Itoa(g.grpcPort), g.dialOptions...)
		if err != nil {
			return nil, err
		}
	}

	return g, nil
}

func (g *GRPCServer) RegisterService(register func(*grpc.Server)) {
	register(g.server)
}

func (g *GRPCServer) GatewayEndpoints(registerEndpoints ...GrpcEndpointRegister) error {
	if g.noGrpcGateway {
		return errors.New("gRPC-gateway is disabled")
	}

	for _, register := range registerEndpoints {
		if err := register(context.Background(), g.gatewayMux, g.gatewayClientConn); err != nil {
			return err
		}
	}

	return nil
}

func (g *GRPCServer) Start() error {
	// Create gRPC listener
	grpcLis, err := net.Listen("tcp", ":"+strconv.Itoa(g.grpcPort))
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	// First start the gRPC server
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		LogInfof("gRPC server listening on port " + strconv.Itoa(g.grpcPort))
		grpc_prometheus.Register(g.server)

		err := g.server.Serve(grpcLis)
		if err != nil {
			LogFatalf("gRPC server failed to serve: %v", err.Error())
		}

		LogInfof("gRPC server stopped")
	}(&wg)

	// Then start the gRPC-gateway
	if !g.noGrpcGateway {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			LogInfof("gRPC-gateway listening on port " + strconv.Itoa(g.gatewayPort))
			err := http.ListenAndServe(":"+strconv.Itoa(g.gatewayPort), g.gatewayMux)
			if err != nil {
				LogFatalf("gRPC-gateway failed to serve: %v", err.Error())
			}

			LogInfof("gRPC-gateway stopped")
		}(&wg)
	}

	// Serve proxmux
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		promMux := http.NewServeMux()
		promMux.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":7332", promMux)
		if err != nil {
			LogFatalf("Prometheus mux failed to serve: %v", err.Error())
		}
	}(&wg)

	wg.Wait()

	return nil
}
