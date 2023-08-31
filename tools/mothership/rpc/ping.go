package mothership_rpc

import (
	"context"
	"database/sql"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (s *Server) WorkerPing(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	worker, err := s.getWorkerIdentity(ctx)
	if err != nil {
		return nil, err
	}

	worker.LastCheckinTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if err := base.Q[mothership_db.Worker](s.db).U(worker); err != nil {
		return nil, status.Error(codes.Internal, "failed to update worker")
	}

	return &emptypb.Empty{}, nil
}
