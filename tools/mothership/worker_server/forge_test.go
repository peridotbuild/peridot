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
	"fmt"
	transport_http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"go.resf.org/peridot/tools/mothership/worker_server/forge"
	"time"
)

type inMemoryForge struct {
	localTempDir  string
	repos         map[string]bool
	remoteBaseURL string
}

func (f *inMemoryForge) GetAuthenticator() (*forge.Authenticator, error) {
	return &forge.Authenticator{
		AuthMethod: &transport_http.BasicAuth{
			Username: "user",
			Password: "pass",
		},
		AuthorName:  "Test User",
		AuthorEmail: "test@resf.org",
		Expires:     time.Now().Add(time.Hour),
	}, nil
}

func (f *inMemoryForge) GetRemote(repo string) string {
	return fmt.Sprintf("file://%s/%s", f.localTempDir, repo)
}

func (f *inMemoryForge) GetCommitViewerURL(repo string, commit string) string {
	return f.remoteBaseURL + "/" + repo + "/commit/" + commit
}

func (f *inMemoryForge) EnsureRepositoryExists(auth *forge.Authenticator, repo string) error {
	f.repos[repo] = true
	return nil
}
