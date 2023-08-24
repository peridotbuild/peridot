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
)

type Worker struct {
	PikaTableName string `pika:"workers"`

	Name            string       `json:"name"`
	WorkerID        string       `json:"worker_id"`
	LastCheckinTime sql.NullTime `json:"last_check_in_time"`
}

func (w *Worker) GetID() string {
	return w.Name
}

func (w *Worker) ToPB() *mshipadminpb.Worker {
	return &mshipadminpb.Worker{
		Name:            w.Name,
		WorkerId:        w.WorkerID,
		LastCheckinTime: base.SqlNullTime(w.LastCheckinTime),
	}
}
