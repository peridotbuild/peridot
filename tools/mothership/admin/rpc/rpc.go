package mothershipadmin_rpc

import (
	base "go.resf.org/peridot/base/go"
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	"google.golang.org/grpc"
)

type Server struct {
	base.GRPCServer
	mshipadminpb.UnimplementedMshipAdminServer

	db *base.DB
}

func NewServer(db *base.DB, oidcInterceptorDetails *base.OidcInterceptorDetails, opts ...base.GRPCServerOption) (*Server, error) {
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
		db:         db,
	}, nil
}

func (s *Server) Start() error {
	s.RegisterService(func(server *grpc.Server) {
		mshipadminpb.RegisterMshipAdminServer(server, s)
	})
	if err := s.GatewayEndpoints(
		mshipadminpb.RegisterMshipAdminHandler,
	); err != nil {
		return err
	}

	return s.GRPCServer.Start()
}
