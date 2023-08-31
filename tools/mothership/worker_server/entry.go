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

package mothership_worker_server

import (
	"database/sql"
	"github.com/pkg/errors"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"go.temporal.io/sdk/temporal"
	"time"
)

func (w *Worker) CreateEntry(args *mothershippb.ProcessRPMArgs) (*mothershippb.Entry, error) {
	req := args.Request
	internalReq := args.InternalRequest
	entry := mothership_db.Entry{
		Name:           base.NameGen("entries"),
		OSRelease:      req.OsRelease,
		Sha256Sum:      req.Checksum,
		RepositoryName: req.Repository,
		WorkerID: sql.NullString{
			String: internalReq.WorkerId,
			Valid:  true,
		},
		State: mothershippb.Entry_ARCHIVING,
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

func (w *Worker) SetEntryState(entry string, state mothershippb.Entry_State, importRpmRes *mothershippb.ImportRPMResponse) (*mothershippb.Entry, error) {
	ent, err := base.Q[mothership_db.Entry](w.db).F("name", entry).GetOrNil()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get entry")
	}
	if ent == nil {
		return nil, temporal.NewNonRetryableApplicationError(
			"entry does not exist",
			"entryDoesNotExist",
			errors.New("entry does not exist"),
		)
	}

	ent.State = state
	if importRpmRes != nil {
		ent.CommitURI = importRpmRes.CommitUri
		ent.CommitHash = importRpmRes.CommitHash
	}

	if err := base.Q[mothership_db.Entry](w.db).U(ent); err != nil {
		return nil, errors.Wrap(err, "failed to update entry")
	}

	return ent.ToPB(), nil
}

func (w *Worker) SetWorkerLastCheckinTime(workerID string) error {
	wrk, err := base.Q[mothership_db.Worker](w.db).F("worker_id", workerID).GetOrNil()
	if err != nil {
		return errors.Wrap(err, "failed to get worker")
	}
	if wrk == nil {
		return temporal.NewNonRetryableApplicationError(
			"worker does not exist",
			"workerDoesNotExist",
			errors.New("worker does not exist"),
		)
	}

	wrk.LastCheckinTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	return base.Q[mothership_db.Worker](w.db).U(wrk)
}
