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

type Batch struct {
	PikaTableName      string `pika:"batches"`
	PikaDefaultOrderBy string `pika:"-create_time"`

	Name          string         `db:"name"`
	BatchID       string         `db:"batch_id"`
	WorkerID      string         `db:"worker_id"`
	CreateTime    time.Time      `db:"create_time" pika:"omitempty"`
	UpdateTime    time.Time      `db:"create_time" pika:"omitempty"`
	SealTime      sql.NullTime   `db:"create_time"`
	BugtrackerURI sql.NullString `db:"bugtracker_uri"`
}

func (b *Batch) GetID() string {
	return b.Name
}

func (b *Batch) ToPB() *mothershippb.Batch {
	return &mothershippb.Batch{
		Name:          b.Name,
		BatchId:       b.BatchID,
		WorkerId:      b.WorkerID,
		CreateTime:    timestamppb.New(b.CreateTime),
		UpdateTime:    timestamppb.New(b.CreateTime),
		SealTime:      base.SqlNullTime(b.SealTime),
		BugtrackerUri: base.SqlNullString(b.BugtrackerURI),
	}
}
