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

package mothership_db

import (
	"database/sql"
	base "go.resf.org/peridot/base/go"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Entry struct {
	PikaTableName      string `pika:"entries"`
	PikaDefaultOrderBy string `pika:"-create_time"`

	Name           string                   `db:"name"`
	EntryID        string                   `db:"entry_id"`
	CreateTime     time.Time                `db:"create_time" pika:"omitempty"`
	OSRelease      string                   `db:"os_release"`
	Sha256Sum      string                   `db:"sha256_sum"`
	RepositoryName string                   `db:"repository_name"`
	WorkerID       sql.NullString           `db:"worker_id"`
	BatchName      sql.NullString           `db:"batch_name"`
	UserEmail      sql.NullString           `db:"user_email"`
	CommitURI      string                   `db:"commit_uri"`
	CommitHash     string                   `db:"commit_hash"`
	State          mothershippb.Entry_State `db:"state"`
}

func (e *Entry) GetID() string {
	return e.Name
}

func (e *Entry) ToPB() *mothershippb.Entry {
	return &mothershippb.Entry{
		Name:       e.Name,
		EntryId:    e.EntryID,
		CreateTime: timestamppb.New(e.CreateTime),
		OsRelease:  e.OSRelease,
		Sha256Sum:  e.Sha256Sum,
		Repository: e.RepositoryName,
		WorkerId:   base.SqlNullString(e.WorkerID),
		Batch:      base.SqlNullString(e.BatchName),
		UserEmail:  base.SqlNullString(e.UserEmail),
		CommitUri:  e.CommitURI,
		CommitHash: e.CommitHash,
		State:      e.State,
	}
}
