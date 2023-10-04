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

package kernelmanager_worker

import (
	"go.resf.org/peridot/base/go/forge"
	"go.resf.org/peridot/base/go/kv"
	"go.resf.org/peridot/base/go/storage"
)

type Worker struct {
	kv      kv.KV
	forge   forge.Forge
	storage storage.Storage
}

func New(kv kv.KV, forge forge.Forge, st storage.Storage) *Worker {
	return &Worker{
		kv:      kv,
		forge:   forge,
		storage: st,
	}
}
