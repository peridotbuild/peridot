package mothershipadmin_rpc

import (
	base "go.resf.org/peridot/base/go"
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	"google.golang.org/grpc"
)

type Server struct {
	mshipadminpb.UnimplementedMshipAdminServer

	db *base.DB
}

func NewServer(db *base.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Start(opts ...base.GRPCServerOption) error {
	grpcServer, err := base.NewGRPCServer(opts...)
	if err != nil {
		return err
	}

	grpcServer.RegisterService(func(server *grpc.Server) {
		mshipadminpb.RegisterMshipAdminServer(server, s)
	})
	if err := grpcServer.GatewayEndpoints(
		mshipadminpb.RegisterMshipAdminHandler,
	); err != nil {
		return err
	}

	return grpcServer.Start()
}
