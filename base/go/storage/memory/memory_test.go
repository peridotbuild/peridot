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
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/require"
	"go.resf.org/peridot/base/go/storage"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	require.NotNil(t, New(memfs.New()))
}

func TestInMemory_Download_Found(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	im.blobs["foo"] = []byte("bar")
	err := im.Download("foo", "foo")
	require.Nil(t, err)

	_, err = fs.Stat("foo")
	require.Nil(t, err)

	f, err := fs.Open("foo")
	require.Nil(t, err)

	buf := make([]byte, 3)
	_, err = f.Read(buf)
	require.Nil(t, err)
	require.Equal(t, []byte("bar"), buf)
}

func TestInMemory_Download_Found_OnFS(t *testing.T) {
	fs := memfs.New()
	{
		f, _ := fs.OpenFile("foo", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		_, err := f.Write([]byte("bar"))
		require.Nil(t, err)
		require.Nil(t, f.Close())
	}
	im := New(fs)

	err := im.Download("foo", "foo")
	require.Nil(t, err)

	_, err = fs.Stat("foo")
	require.Nil(t, err)

	f, err := fs.Open("foo")
	require.Nil(t, err)

	buf := make([]byte, 3)
	_, err = f.Read(buf)
	require.Nil(t, err)
	require.Equal(t, []byte("bar"), buf)
}

func TestInMemory_Download_NotFound(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	err := im.Download("foo", "foo")
	require.Equal(t, storage.ErrNotFound, err)
}

func TestInMemory_Get_Found(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	im.blobs["foo"] = []byte("bar")
	blob, err := im.Get("foo")
	require.Nil(t, err)
	require.Equal(t, []byte("bar"), blob)
}

func TestInMemory_Get_Found_OnFS(t *testing.T) {
	fs := memfs.New()
	{
		f, _ := fs.OpenFile("foo", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		_, err := f.Write([]byte("bar"))
		require.Nil(t, err)
		require.Nil(t, f.Close())
	}
	im := New(fs)
	blob, err := im.Get("foo")
	require.Nil(t, err)
	require.Equal(t, []byte("bar"), blob)
}

func TestInMemory_Get_NotFound(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	_, err := im.Get("foo")
	require.Equal(t, storage.ErrNotFound, err)
}

func TestInMemory_Put(t *testing.T) {
	fs := memfs.New()

	f, err := fs.Create("foo")
	require.Nil(t, err)

	_, err = f.Write([]byte("bar"))
	require.Nil(t, err)

	err = f.Close()
	require.Nil(t, err)

	im := New(fs)
	_, err = im.Put("foo", "foo")
	require.Nil(t, err)
	require.Equal(t, []byte("bar"), im.blobs["foo"])
}

func TestInMemory_Put_NotFound(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	_, err := im.Put("foo", "testdata/bar")
	require.NotNil(t, err)
	require.Equal(t, "failed to open file: file does not exist", err.Error())
}

func TestInMemory_PutBytes(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	_, err := im.PutBytes("foo", []byte("bar"))
	require.Nil(t, err)
	require.Equal(t, []byte("bar"), im.blobs["foo"])
}

func TestInMemory_Delete(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	im.blobs["foo"] = []byte("bar")
	err := im.Delete("foo")
	require.Nil(t, err)
	_, ok := im.blobs["foo"]
	require.False(t, ok)
}

func TestInMemory_Exists_Found(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	im.blobs["foo"] = []byte("bar")
	ok, err := im.Exists("foo")
	require.Nil(t, err)
	require.True(t, ok)
}

func TestInMemory_Exists_Found_OnFS(t *testing.T) {
	fs := memfs.New()
	{
		f, _ := fs.OpenFile("foo", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		_, err := f.Write([]byte("bar"))
		require.Nil(t, err)
		require.Nil(t, f.Close())
	}
	im := New(fs)
	ok, err := im.Exists("foo")
	require.Nil(t, err)
	require.True(t, ok)
}

func TestInMemory_Exists_NotFound(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	ok, err := im.Exists("foo")
	require.Nil(t, err)
	require.False(t, ok)
}

func TestInMemory_CanReadURI(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	ok, err := im.CanReadURI("memory://foo")
	require.Nil(t, err)
	require.True(t, ok)
}

func TestInMemory_CanReadURI_No(t *testing.T) {
	fs := memfs.New()
	im := New(fs)
	ok, err := im.CanReadURI("file://foo")
	require.Nil(t, err)
	require.False(t, ok)
}
