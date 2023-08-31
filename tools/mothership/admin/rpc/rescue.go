package mothershipadmin_rpc

import (
	"context"
	base "go.resf.org/peridot/base/go"
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RescueEntryImport(ctx context.Context, req *mshipadminpb.RescueEntryImportRequest) (*emptypb.Empty, error) {
	// First make sure an entry with the given name exists.
	entry, err := base.Q[mothership_db.Entry](s.db).F("name", req.Name).GetOrNil()
	if err != nil {
		base.LogErrorf("failed to get entry: %v", err)
		return nil, status.Error(codes.Internal, "failed to get entry")
	}

	if entry == nil {
		return nil, status.Error(codes.NotFound, "entry not found")
	}

	// Make sure the entry is on hold.
	if entry.State != mothershippb.Entry_ON_HOLD {
		return nil, status.Error(codes.FailedPrecondition, "entry is not on hold")
	}

	// If on hold, then signal the workflow to continue.
	err = s.temporal.SignalWorkflow(ctx, "operations/"+entry.Sha256Sum, "", "rescue", true)
	if err != nil {
		base.LogErrorf("failed to signal workflow: %v", err)
		return nil, status.Error(codes.Internal, "failed to signal workflow")
	}

	return &emptypb.Empty{}, nil
}
