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
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sassoftware/go-rpmutils"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"go.temporal.io/sdk/temporal"
	"io"
	"os"
	"path/filepath"
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

// SetEntryIDFromRPM sets the entry ID from the RPM.
// This is a Temporal activity.
func (w *Worker) SetEntryIDFromRPM(entry string, uri string, checksumSha256 string) (*mothershippb.Entry, error) {
	ent, err := base.Q[mothership_db.Entry](w.db).F("name", entry).GetOrNil()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get entry")
	}
	if ent == nil {
		return nil, errors.New("entry does not exist")
	}

	tempDir, err := os.MkdirTemp("", "mothership-worker-server-import-rpm-*")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	object, err := getObjectPath(uri)
	if err != nil {
		return nil, err
	}

	err = w.storage.Download(object, filepath.Join(tempDir, "resource.rpm"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to download resource")
	}

	// Verify checksum
	hash := sha256.New()
	f, err := os.Open(filepath.Join(tempDir, "resource.rpm"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open resource")
	}
	defer f.Close()
	if _, err := io.Copy(hash, f); err != nil {
		return nil, errors.Wrap(err, "failed to hash resource")
	}
	if hex.EncodeToString(hash.Sum(nil)) != checksumSha256 {
		return nil, errors.New("checksum does not match")
	}

	// Read the RPM headers
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, errors.Wrap(err, "failed to seek resource")
	}
	rpm, err := rpmutils.ReadRpm(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read RPM headers")
	}

	nevra, err := rpm.Header.GetNEVRA()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get RPM NEVRA")
	}

	// Set entry ID
	ent.EntryID = fmt.Sprintf("%s-%s-%s.src", nevra.Name, nevra.Version, nevra.Release)
	ent.Sha256Sum = checksumSha256

	// Update entry
	if err := base.Q[mothership_db.Entry](w.db).U(ent); err != nil {
		return nil, errors.Wrap(err, "failed to update entry")
	}

	return ent.ToPB(), nil
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
		ent.CommitBranch = importRpmRes.CommitBranch
		ent.CommitTag = importRpmRes.CommitTag
		ent.PackageName = importRpmRes.Pkg
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
