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
	"errors"
	"fmt"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	transport_http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"go.resf.org/peridot/base/go/forge"
	"path/filepath"
	"time"
)

type inMemoryForge struct {
	localTempDir        string
	repos               map[string]bool
	remoteBaseURL       string
	invalidUsernamePass bool
	noAuthMethod        bool
}

func (f *inMemoryForge) GetAuthenticator() (*forge.Authenticator, error) {
	ret := &forge.Authenticator{
		AuthMethod: &transport_http.BasicAuth{
			Username: "user",
			Password: "pass",
		},
		AuthorName:  "Test User",
		AuthorEmail: "test@resf.org",
		Expires:     time.Now().Add(time.Hour),
	}

	if f.noAuthMethod {
		ret.AuthMethod = nil
	} else if f.invalidUsernamePass {
		ret.AuthMethod = &transport_http.BasicAuth{
			Username: "invalid",
			Password: "invalid",
		}
	}

	return ret, nil
}

func (f *inMemoryForge) GetRemote(repo string) string {
	return fmt.Sprintf("file://%s/%s", f.localTempDir, repo)
}

func (f *inMemoryForge) GetCommitViewerURL(repo string, commit string) string {
	return f.remoteBaseURL + "/" + repo + "/commit/" + commit
}

func (f *inMemoryForge) EnsureRepositoryExists(auth *forge.Authenticator, repo string) error {
	// Try casting auth.AuthMethod to *transport_http.BasicAuth
	// If it fails, return an error
	authx, ok := auth.AuthMethod.(*transport_http.BasicAuth)
	if !ok {
		return errors.New("auth failed")
	}
	if authx.Username != "user" || authx.Password != "pass" {
		return errors.New("username or password incorrect")
	}

	if f.repos[repo] {
		return nil
	}

	osfsTemp := osfs.New(filepath.Join(f.localTempDir, repo))
	dot, err := osfsTemp.Chroot(".git")
	if err != nil {
		return err
	}

	filesystemTemp := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
	err = filesystemTemp.Init()
	if err != nil {
		return err
	}

	_, err = git.Init(filesystemTemp, nil)
	if err != nil {
		return err
	}

	f.repos[repo] = true
	return nil
}
