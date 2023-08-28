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

package storage_detector

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.resf.org/peridot/base/go/storage"
	storage_memory "go.resf.org/peridot/base/go/storage/memory"
	storage_s3 "go.resf.org/peridot/base/go/storage/s3"
	"net/url"
)

func FromFlags(ctx *cli.Context) (storage.Storage, error) {
	parsedURI, err := url.Parse(ctx.String("storage-connection-string"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse storage connection string")
	}

	switch parsedURI.Scheme {
	case "s3":
		return storage_s3.FromFlags(ctx)
	case "memory":
		return storage_memory.New(osfs.New("/")), nil
	default:
		return nil, errors.Errorf("unknown storage scheme: %s", parsedURI.Scheme)
	}
}
