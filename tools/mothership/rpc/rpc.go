package mothership_rpc

import (
	base "go.resf.org/peridot/base/go"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"google.golang.org/grpc"
)

type Server struct {
	mothershippb.UnimplementedSrpmArchiverServer

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
		mothershippb.RegisterSrpmArchiverServer(server, s)
	})
	if err := grpcServer.GatewayEndpoints(
		mothershippb.RegisterSrpmArchiverHandler,
	); err != nil {
		return err
	}

	return grpcServer.Start()
}
