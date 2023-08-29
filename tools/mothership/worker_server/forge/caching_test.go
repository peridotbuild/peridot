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
	transport_http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type testForge struct {
	Forge

	callToGetAuthenticator int
}

func (t *testForge) GetAuthenticator() (*Authenticator, error) {
	t.callToGetAuthenticator++
	return &Authenticator{
		AuthMethod: &transport_http.BasicAuth{
			Username: "test",
			Password: "test",
		},
		AuthorName:  "test",
		AuthorEmail: "test@resf.org",
		Expires:     time.Now().Add(time.Minute * 45),
	}, nil
}

func TestNewCacher(t *testing.T) {
	f := &testForge{}
	fAuth, err := f.GetAuthenticator()
	require.Nil(t, err)
	fAuthAuthMethod := fAuth.AuthMethod.(*transport_http.BasicAuth)
	require.Equal(t, "test", fAuthAuthMethod.Username)
	require.Equal(t, "test", fAuthAuthMethod.Password)

	c := NewCacher(f)
	cAuth, err := c.GetAuthenticator()
	require.Nil(t, err)
	cAuthAuthMethod := cAuth.AuthMethod.(*transport_http.BasicAuth)
	require.Equal(t, "test", cAuthAuthMethod.Username)
	require.Equal(t, "test", cAuthAuthMethod.Password)
}

func TestCacher_GetAuthenticator(t *testing.T) {
	f := &testForge{}
	c := NewCacher(f)

	cAuth, err := c.GetAuthenticator()
	require.Nil(t, err)
	cAuthAuthMethod := cAuth.AuthMethod.(*transport_http.BasicAuth)
	require.Equal(t, "test", cAuthAuthMethod.Username)
	require.Equal(t, "test", cAuthAuthMethod.Password)
}

func TestCacher_GetAuthenticator_Cached(t *testing.T) {
	f := &testForge{}
	c := NewCacher(f)

	cAuth, err := c.GetAuthenticator()
	require.Nil(t, err)
	require.Equal(t, "test", cAuth.AuthorName)

	cAuth, err = c.GetAuthenticator()
	require.Nil(t, err)
	require.Equal(t, "test", cAuth.AuthorName)

	require.Equal(t, 1, f.callToGetAuthenticator)
}

func TestCacher_GetAuthenticator_Expired(t *testing.T) {
	f := &testForge{}
	c := NewCacher(f)

	cAuth, err := c.GetAuthenticator()
	require.Nil(t, err)
	require.Equal(t, "test", cAuth.AuthorName)

	cAuth.Expires = time.Now().Add(-time.Minute * 10)

	cAuth, err = c.GetAuthenticator()
	require.Nil(t, err)
	require.Equal(t, "test", cAuth.AuthorName)

	require.Equal(t, 2, f.callToGetAuthenticator)
}
