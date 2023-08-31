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
	"go.temporal.io/sdk/client"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
)

type Server struct {
	base.GRPCServer

	mothershippb.UnimplementedSrpmArchiverServer
	longrunning.UnimplementedOperationsServer

	db       *base.DB
	temporal client.Client
}

func NewServer(db *base.DB, temporalClient client.Client, opts ...base.GRPCServerOption) (*Server, error) {
	opts = append(opts, base.WithServeMuxAdditionalHeaders("x-mship-worker-secret"))
	grpcServer, err := base.NewGRPCServer(opts...)
	if err != nil {
		return nil, err
	}

	return &Server{
		GRPCServer: *grpcServer,
		db:         db,
		temporal:   temporalClient,
	}, nil
}

func (s *Server) Start() error {
	s.RegisterService(func(server *grpc.Server) {
		longrunning.RegisterOperationsServer(server, s)
		mothershippb.RegisterSrpmArchiverServer(server, s)
	})
	if err := s.GatewayEndpoints(
		longrunning.RegisterOperationsHandler,
		mothershippb.RegisterSrpmArchiverHandler,
	); err != nil {
		return err
	}

	return s.GRPCServer.Start()
}
