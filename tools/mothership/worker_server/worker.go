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
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/storage"
	"go.resf.org/peridot/tools/mothership/worker_server/forge"
	"golang.org/x/crypto/openpgp"
)

type Worker struct {
	db      *base.DB
	storage storage.Storage
	gpgKeys openpgp.EntityList
	forge   forge.Forge
}

func New(db *base.DB, storage storage.Storage, gpgKeys openpgp.EntityList, forge forge.Forge) *Worker {
	return &Worker{
		db:      db,
		storage: storage,
		gpgKeys: gpgKeys,
		forge:   forge,
	}
}
