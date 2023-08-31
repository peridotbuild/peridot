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
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Worker struct {
	PikaTableName      string `pika:"workers"`
	PikaDefaultOrderBy string `pika:"-create_time"`

	Name            string       `db:"name"`
	CreateTime      time.Time    `db:"create_time" pika:"omitempty"`
	WorkerID        string       `db:"worker_id"`
	LastCheckinTime sql.NullTime `db:"last_checkin_time"`
	ApiSecret       string       `db:"api_secret"`
}

func (w *Worker) GetID() string {
	return w.Name
}

func (w *Worker) ToPB() *mshipadminpb.Worker {
	return &mshipadminpb.Worker{
		Name:            w.Name,
		WorkerId:        w.WorkerID,
		CreateTime:      timestamppb.New(w.CreateTime),
		LastCheckinTime: base.SqlNullTime(w.LastCheckinTime),
	}
}
