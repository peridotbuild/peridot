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
