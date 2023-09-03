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

package storage_memory

import (
	"github.com/go-git/go-billy/v5"
	"github.com/pkg/errors"
	"go.resf.org/peridot/base/go/storage"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type InMemory struct {
	storage.Storage

	rootPath string
	fs       billy.Filesystem
	blobs    map[string][]byte
}

// New creates a new InMemory storage.
// rootPath is a spread operator because it's optional and it looks nicer.
func New(fs billy.Filesystem, rootPath ...string) *InMemory {
	if len(rootPath) > 1 {
		panic("too many arguments")
	}

	inm := &InMemory{
		fs:    fs,
		blobs: make(map[string][]byte),
	}
	if len(rootPath) == 1 {
		inm.rootPath = rootPath[0]
	}
	return inm
}

func (im *InMemory) getBlob(object string) ([]byte, error) {
	blob, ok := im.blobs[object]
	if !ok {
		// If not in memory, check if it's on disk
		path := object
		if im.rootPath != "" {
			path = filepath.Join(im.rootPath, object)
		}
		f, err := im.fs.Open(path)
		if err != nil {
			return nil, storage.ErrNotFound
		}

		// Read file into blob
		blob, err := io.ReadAll(f)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read file")
		}

		// Store blob
		im.blobs[object] = blob

		return blob, nil
	}
	return blob, nil
}

func (im *InMemory) Download(object string, toPath string) error {
	blob, err := im.getBlob(object)
	if err != nil {
		return err
	}

	// Open file
	f, err := im.fs.OpenFile(toPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}

	// Write blob to file
	_, err = f.Write(blob)
	if err != nil {
		return errors.Wrap(err, "failed to write blob to file")
	}

	return nil
}

func (im *InMemory) Get(object string) ([]byte, error) {
	blob, err := im.getBlob(object)
	if err != nil {
		return nil, err
	}
	return blob, nil
}

func (im *InMemory) Put(object string, fromPath string) (*storage.UploadInfo, error) {
	// Open file
	f, err := im.fs.Open(fromPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	// Read file into blob
	blob, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	// Store blob
	im.blobs[object] = blob

	return &storage.UploadInfo{
		Location:  "memory://" + object,
		VersionID: nil,
	}, nil
}

func (im *InMemory) PutBytes(object string, blob []byte) (*storage.UploadInfo, error) {
	// Store blob
	im.blobs[object] = blob

	return &storage.UploadInfo{
		Location:  "memory://" + object,
		VersionID: nil,
	}, nil
}

func (im *InMemory) Delete(object string) error {
	delete(im.blobs, object)
	return nil
}

func (im *InMemory) Exists(object string) (bool, error) {
	_, err := im.getBlob(object)
	if err != nil {
		if err == storage.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (im *InMemory) CanReadURI(uri string) (bool, error) {
	return strings.HasPrefix(uri, "memory://"), nil
}
