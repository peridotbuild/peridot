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
