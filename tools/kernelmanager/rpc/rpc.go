package kernelmanager_rpc

import (
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/kv"
	kernelmanagerpb "go.resf.org/peridot/tools/kernelmanager/pb"
	"go.temporal.io/sdk/client"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	base.GRPCServer

	longrunning.UnimplementedOperationsServer
	kernelmanagerpb.UnimplementedKernelManagerServer

	kv       kv.KV
	temporal client.Client
}

func NewServer(kv kv.KV, temporalClient client.Client, oidcInterceptorDetails *base.OidcInterceptorDetails, opts ...base.GRPCServerOption) (*Server, error) {
	oidcInterceptor, err := base.OidcGrpcInterceptor(oidcInterceptorDetails)
	if err != nil {
		return nil, err
	}

	opts = append(opts, base.WithUnaryInterceptors(oidcInterceptor))

	grpcServer, err := base.NewGRPCServer(opts...)
	if err != nil {
		return nil, err
	}

	return &Server{
		GRPCServer: *grpcServer,
		kv:         kv,
		temporal:   temporalClient,
	}, nil
}

func (s *Server) Start() error {
	s.RegisterService(func(server *grpc.Server) {
		reflection.Register(server)
		longrunning.RegisterOperationsServer(server, s)
		kernelmanagerpb.RegisterKernelManagerServer(server, s)
	})
	if err := s.GatewayEndpoints(
		longrunning.RegisterOperationsHandler,
		kernelmanagerpb.RegisterKernelManagerHandler,
	); err != nil {
		return err
	}

	return s.GRPCServer.Start()
}
