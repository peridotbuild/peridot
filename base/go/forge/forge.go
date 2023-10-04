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

package forge

import (
	"github.com/go-git/go-git/v5/plumbing/transport"
	"time"
)

type Authenticator struct {
	transport.AuthMethod

	AuthorName  string
	AuthorEmail string
	// Expires is the time when the token expires.
	// So it can be used to cache the token.
	Expires time.Time
}

type Forge interface {
	GetAuthenticator() (*Authenticator, error)
	GetRemote(repo string) string
	GetCommitViewerURL(repo string, commit string) string
	EnsureRepositoryExists(auth *Authenticator, repo string) error
	WithNamespace(namespace string) Forge
}
