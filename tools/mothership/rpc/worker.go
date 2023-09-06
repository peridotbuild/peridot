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
	"context"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// getWorkerIdentity returns the identity of the worker that the request is
// coming from. Returns an error if the worker is not found or unauthenticated.
func (s *Server) getWorkerIdentity(ctx context.Context) (*mothership_db.Worker, error) {
	// Get x-mship-worker-secret
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	secrets := md["x-mship-worker-secret"]
	if len(secrets) != 1 {
		return nil, status.Error(codes.Unauthenticated, "missing worker secret")
	}

	secret := secrets[0]
	worker, err := base.Q[mothership_db.Worker](s.db).F("api_secret", secret).GetOrNil()
	if err != nil {
		base.LogErrorf("failed to get worker: %v", err)
		return nil, status.Error(codes.Internal, "failed to get worker")
	}

	if worker == nil {
		return nil, status.Error(codes.Unauthenticated, "invalid worker secret")
	}

	return worker, nil
}
