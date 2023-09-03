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
	"context"
	base "go.resf.org/peridot/base/go"
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	mothership_worker_server "go.resf.org/peridot/tools/mothership/worker_server"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Server) RetractEntry(ctx context.Context, req *mshipadminpb.RetractEntryRequest) (*longrunning.Operation, error) {
	startWorkflowOpts := client.StartWorkflowOptions{
		ID:                                       "operations/retract/" + req.Name,
		WorkflowExecutionErrorWhenAlreadyStarted: true,
		WorkflowIDReusePolicy:                    enumspb.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
	}

	// Submit to Temporal
	run, err := s.temporal.ExecuteWorkflow(
		context.Background(),
		startWorkflowOpts,
		mothership_worker_server.RetractEntryWorkflow,
		req.Name,
	)
	if err != nil {
		if strings.Contains(err.Error(), "is already running") {
			return nil, status.Error(codes.AlreadyExists, "entry is already running")
		}
		base.LogErrorf("failed to start workflow: %v", err)
		return nil, status.Error(codes.Internal, "failed to start workflow")
	}

	return s.getOperation(ctx, run.GetID())
}
