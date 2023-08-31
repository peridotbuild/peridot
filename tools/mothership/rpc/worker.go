package mothership_rpc

import (
	"context"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	mothership_worker_server "go.resf.org/peridot/tools/mothership/worker_server"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
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

// SubmitEntry handles the RPC request for submitting an entry. This is usually
// called by the worker. The worker must be authenticated. The checksum will "lease"
// the entry for the worker, so that other workers will not submit the same entry.
// This "lease" is enforced using Temporal
func (s *Server) SubmitEntry(ctx context.Context, req *mothershippb.SubmitEntryRequest) (*longrunning.Operation, error) {
	worker, err := s.getWorkerIdentity(ctx)
	if err != nil {
		return nil, err
	}

	startWorkflowOpts := client.StartWorkflowOptions{
		ID:                                       "operations/" + req.ProcessRpmRequest.Checksum,
		WorkflowExecutionErrorWhenAlreadyStarted: true,
		WorkflowIDReusePolicy:                    enumspb.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}

	// Submit to Temporal
	run, err := s.temporal.ExecuteWorkflow(
		context.Background(),
		startWorkflowOpts,
		mothership_worker_server.ProcessRPMWorkflow,
		&mothershippb.ProcessRPMArgs{
			Request: req.ProcessRpmRequest,
			InternalRequest: &mothershippb.ProcessRPMInternalRequest{
				WorkerId: worker.WorkerID,
			},
		},
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
