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
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWorker_VerifyResourceExists(t *testing.T) {
	require.Nil(t, testW.VerifyResourceExists("memory://efi-rpm-macros-3-3.el8.src.rpm"))
}

func TestWorker_VerifyResourceExists_NotFound(t *testing.T) {
	err := testW.VerifyResourceExists("memory://not-found.rpm")
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "resource does not exist")
}

func TestWorker_VerifyResourceExists_CannotRead(t *testing.T) {
	err := testW.VerifyResourceExists("bad-protocol://not-found.rpm")
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "client submitted a resource URI that cannot be read by server")
}
