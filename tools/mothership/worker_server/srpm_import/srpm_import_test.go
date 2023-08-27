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

package srpm_import

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/stretchr/testify/require"
	storage_memory "go.resf.org/peridot/base/go/storage/memory"
	"golang.org/x/crypto/openpgp"
	"os"
	"testing"
)

func TestFromFile(t *testing.T) {
	s, err := FromFile("testdata/efi-rpm-macros-3-3.el8.src.rpm")
	require.Nil(t, err)
	require.NotNil(t, s)
	require.Nil(t, s.Close())
}

func TestFromFile_SignatureOK(t *testing.T) {
	keyF, err := os.Open("testdata/RPM-GPG-KEY-Rocky-8")
	require.Nil(t, err)

	testKey, err := openpgp.ReadArmoredKeyRing(keyF)
	require.Nil(t, err)

	s, err := FromFile("testdata/efi-rpm-macros-3-3.el8.src.rpm", testKey...)
	require.Nil(t, err)
	require.NotNil(t, s)
	require.Nil(t, s.Close())
}

func TestFromFile_SignatureFail(t *testing.T) {
	keyF, err := os.Open("testdata/RPM-GPG-KEY-Rocky-9")
	require.Nil(t, err)

	testKey, err := openpgp.ReadArmoredKeyRing(keyF)
	require.Nil(t, err)

	s, err := FromFile("testdata/efi-rpm-macros-3-3.el8.src.rpm", testKey...)
	require.NotNil(t, err)
	require.Nil(t, s)
	require.Equal(t, "failed to verify RPM: keyid 15af5dac6d745a60 not found", err.Error())
}

func TestDetermineLookasideBlobs_Empty(t *testing.T) {
	s, err := FromFile("testdata/efi-rpm-macros-3-3.el8.src.rpm")
	require.Nil(t, err)
	require.NotNil(t, s)
	defer func() {
		require.Nil(t, s.Close())
	}()
	require.Nil(t, s.determineLookasideBlobs())
	require.Equal(t, 0, len(s.lookasideBlobs))
}

func TestDetermineLookasideBlobs_NotEmpty_Over5MB(t *testing.T) {
	s, err := FromFile("testdata/bash-4.4.20-4.el8_6.src.rpm")
	require.Nil(t, err)
	require.NotNil(t, s)
	defer func() {
		require.Nil(t, s.Close())
	}()
	require.Nil(t, s.determineLookasideBlobs())
	require.Equal(t, 1, len(s.lookasideBlobs))
}

func TestUploadLookaside_Empty(t *testing.T) {
	s, err := FromFile("testdata/efi-rpm-macros-3-3.el8.src.rpm")
	require.Nil(t, err)
	require.NotNil(t, s)
	defer func() {
		require.Nil(t, s.Close())
	}()
	require.Nil(t, s.determineLookasideBlobs())

	// we can use memfs since we're not actually writing anything
	fs := memfs.New()
	lookaside := storage_memory.New(fs)
	require.Nil(t, s.uploadLookasideBlobs(lookaside))

	fi, err := fs.ReadDir(".")
	require.Nil(t, err)
	require.Equal(t, 0, len(fi))
}

func TestUploadLookaside_NotEmpty(t *testing.T) {
	s, err := FromFile("testdata/bash-4.4.20-4.el8_6.src.rpm")
	require.Nil(t, err)
	require.NotNil(t, s)
	defer func() {
		require.Nil(t, s.Close())
	}()
	require.Nil(t, s.determineLookasideBlobs())

	fs := osfs.New("/")
	lookaside := storage_memory.New(fs)
	require.Nil(t, s.uploadLookasideBlobs(lookaside))

	ok, err := lookaside.Exists("d86b3392c1202e8ff5a423b302e6284db7f8f435ea9f39b5b1b20fd3ac36dfcb")
	require.Nil(t, err)
	require.True(t, ok)
}
