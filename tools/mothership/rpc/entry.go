package mothership_rpc

import (
	"context"
	"go.ciq.dev/pika"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetEntry(_ context.Context, req *mothershippb.GetEntryRequest) (*mothershippb.Entry, error) {
	entry, err := base.Q[mothership_db.Entry](s.db).F("name", req.Name).GetOrNil()
	if err != nil {
		base.LogErrorf("failed to get entry: %v", err)
		return nil, status.Error(codes.Internal, "failed to get entry")
	}

	if entry == nil {
		return nil, status.Error(codes.NotFound, "entry not found")
	}

	return entry.ToPB(), nil
}

func (s *Server) ListEntries(_ context.Context, req *mothershippb.ListEntriesRequest) (*mothershippb.ListEntriesResponse, error) {
	aipOptions := pika.ProtoReflect(&mothershippb.Entry{})

	page, nt, err := base.Q[mothership_db.Entry](s.db).GetPage(req, aipOptions)
	if err != nil {
		base.LogErrorf("failed to get entry page: %v", err)
		return nil, status.Error(codes.Internal, "failed to get entry page")
	}

	return &mothershippb.ListEntriesResponse{
		Entries:       base.SliceToPB[*mothershippb.Entry, *mothership_db.Entry](page),
		NextPageToken: nt,
	}, nil
}
