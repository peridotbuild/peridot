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

package storage

import "errors"

var ErrNotFound = errors.New("not found")

// UploadInfo is the information about an upload.
type UploadInfo struct {
	// Location is the location of the uploaded file.
	Location string

	// VersionID is the version ID of the uploaded file.
	VersionID *string
}

// Storage is an interface for storage backends.
// Usually S3, but can be anything.
type Storage interface {
	// Download downloads a file from the storage backend to the given path.
	Download(object string, toPath string) error

	// Get returns the contents of a file from the storage backend.
	// Returns ErrNotFound if the file does not exist.
	Get(object string) ([]byte, error)

	// Put uploads a file to the storage backend.
	Put(object string, fromPath string) (*UploadInfo, error)

	// PutBytes uploads a file to the storage backend.
	PutBytes(object string, data []byte) (*UploadInfo, error)

	// Delete deletes a file from the storage backend.
	// Returns ErrNotFound if the file does not exist.
	Delete(object string) error

	// Exists checks if a file exists in the storage backend.
	// Returns false if the file does not exist.
	Exists(object string) (bool, error)

	// CanReadURI checks if a URI can be read by the storage backend.
	// Returns false if the URI cannot be read.
	CanReadURI(uri string) (bool, error)
}
