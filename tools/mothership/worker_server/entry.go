package mothership_worker_server

import (
	"database/sql"
	"github.com/pkg/errors"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
)

func (w *Worker) CreateEntry(args *mothershippb.ProcessRPMArgs, importRpmRes *mothershippb.ImportRPMResponse) (*mothershippb.Entry, error) {
	req := args.Request
	internalReq := args.InternalRequest
	entry := mothership_db.Entry{
		Name:           base.NameGen("entries"),
		EntryID:        importRpmRes.Nevra,
		OSRelease:      req.OsRelease,
		Sha256Sum:      req.Checksum,
		RepositoryName: req.Repository,
		WorkerID: sql.NullString{
			String: internalReq.WorkerId,
			Valid:  true,
		},
		CommitURI:  importRpmRes.CommitUri,
		CommitHash: importRpmRes.CommitHash,
	}
	if req.Batch != "" {
		entry.BatchName = sql.NullString{
			String: req.Batch,
			Valid:  true,
		}
	}

	err := base.Q[mothership_db.Entry](w.db).Create(&entry)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create entry")
	}

	return entry.ToPB(), nil
}
