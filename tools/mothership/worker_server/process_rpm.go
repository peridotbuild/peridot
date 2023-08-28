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
	"encoding/hex"
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
	"github.com/sassoftware/go-rpmutils"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"go.resf.org/peridot/tools/mothership/worker_server/srpm_import"
	"go.temporal.io/sdk/temporal"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// VerifyResourceExists verifies that the resource exists.
// This is a Temporal activity.
func (w *Worker) VerifyResourceExists(uri string) error {
	canRead, err := w.storage.CanReadURI(uri)
	if err != nil {
		return errors.Wrap(err, "failed to check if resource URI can be read")
	}

	if !canRead {
		return temporal.NewNonRetryableApplicationError(
			"cannot read resource URI",
			"cannotReadResourceURI",
			errors.New("client submitted a resource URI that cannot be read by server"),
		)
	}

	// Get object name from URI.
	// Check if object exists.
	// If not, return error.
	parsed, err := url.Parse(uri)
	if err != nil {
		return temporal.NewNonRetryableApplicationError(
			"could not parse resource URI",
			"couldNotParseResourceURI",
			errors.Wrap(err, "failed to parse resource URI"),
		)
	}

	split := strings.SplitN(parsed.Path, "/", 2)
	if len(split) < 2 {
		return temporal.NewNonRetryableApplicationError(
			"invalid resource URI",
			"invalidResourceURI",
			errors.New("client submitted an invalid resource URI"),
		)
	}

	object := split[1]
	exists, err := w.storage.Exists(object)
	if err != nil {
		return errors.Wrap(err, "failed to check if resource exists")
	}

	if !exists {
		// Since the client can trigger this activity before uploading the resource,
		// we should not return a non-retryable error.
		// The parent workflow should handle the retry arrangements up to 2 hours
		// per the spec.
		return errors.New("resource does not exist")
	}

	return nil
}

// ImportRPM imports an RPM into the database.
// This is a Temporal activity.
func (w *Worker) ImportRPM(uri string, checksumSha256 string) (*mothershippb.ImportRPMResponse, error) {
	tempDir, err := os.MkdirTemp("", "mothership-worker-server-import-rpm-*")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	// Parse uri
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse resource URI")
	}

	// Download the resource to the temporary directory
	err = w.storage.Download(parsed.Path, filepath.Join(tempDir, "resource.rpm"))
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

	// Ensure repository exists
	repoName := nevra.Name

	// First ensure that the repo exists.
	authenticator, err := w.forge.GetAuthenticator()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get forge authenticator")
	}

	err = w.forge.EnsureRepositoryExists(authenticator, repoName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ensure repository exists")
	}

	// Then do an import
	srpmState, err := srpm_import.FromFile(filepath.Join(tempDir, "resource.rpm"), w.gpgKeys...)
	if err != nil {
		if strings.Contains(err.Error(), "failed to verify RPM") {
			return nil, temporal.NewNonRetryableApplicationError(
				"failed to verify RPM",
				"failedToVerifyRPM",
				err,
			)
		}
		return nil, errors.Wrap(err, "failed to import SRPM")
	}
	srpmState.SetAuthor(authenticator.AuthorName, authenticator.AuthorEmail)

	cloneOpts := &git.CloneOptions{
		URL:  fmt.Sprintf("%s/%s", w.forge.GetRemote(), repoName),
		Auth: authenticator.AuthMethod,
	}
	storer := memory.NewStorage()
	fs := memfs.New()
	commit, err := srpmState.Import(cloneOpts, storer, fs, w.storage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to import SRPM")
	}

	return &mothershippb.ImportRPMResponse{
		CommitHash: commit.Hash.String(),
	}, nil
}
