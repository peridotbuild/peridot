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

package mothership_rpc

import (
	base "go.resf.org/peridot/base/go"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"google.golang.org/grpc"
)

type Server struct {
	base.GRPCServer
	mothershippb.UnimplementedSrpmArchiverServer

	db *base.DB
}

func NewServer(db *base.DB, opts ...base.GRPCServerOption) (*Server, error) {
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
		mothershippb.RegisterSrpmArchiverServer(server, s)
	})
	if err := s.GatewayEndpoints(
		mothershippb.RegisterSrpmArchiverHandler,
	); err != nil {
		return err
	}

	return s.GRPCServer.Start()
}
