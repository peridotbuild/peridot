package srpm_import

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/rocky-linux/srpmproc/pkg/data"
	"github.com/stretchr/testify/require"
	storage_memory "go.resf.org/peridot/base/go/storage/memory"
	"testing"
)

func TestSrpmprocBlobCompat_Write(t *testing.T) {
	lookaside := storage_memory.New(memfs.New())
	s := &srpmprocBlobCompat{lookaside}
	require.Nil(t, s.Write("test", []byte("test")))
	x, err := lookaside.Get("test")
	require.Nil(t, err)
	require.Equal(t, []byte("test"), x)
}

func TestSrpmprocBlobCompat_Read(t *testing.T) {
	lookaside := storage_memory.New(memfs.New())
	s := &srpmprocBlobCompat{lookaside}
	_, err := lookaside.PutBytes("test", []byte("test"))
	require.Nil(t, err)
	x, err := s.Read("test")
	require.Nil(t, err)
	require.Equal(t, []byte("test"), x)
}

func TestSrpmprocImportModeCompat_ImportName(t *testing.T) {
	s := &srpmprocImportModeCompat{}
	pd := &data.ProcessData{
		ImportBranchPrefix: "r",
		Version:            9,
		RpmLocation:        "bash",
	}
	md := &data.ModeData{
		SourcesToIgnore: []*data.IgnoredSource{},
		TagBranch:       "refs/tags/imports/r9/bash-5.1.8-4.el9",
	}
	require.Equal(t, "bash-5.1.8-4.el9", s.ImportName(pd, md))
}

// todo(mustafa): actually recall what this mode was useful for in srpmproc. like what is this??
func TestSrpmprocImportModeCompat_ImportName_NoTag(t *testing.T) {
	s := &srpmprocImportModeCompat{}
	pd := &data.ProcessData{
		ImportBranchPrefix: "el",
		Version:            9,
		RpmLocation:        "bash",
	}
	md := &data.ModeData{
		SourcesToIgnore: []*data.IgnoredSource{},
		TagBranch:       "refs/heads/el9",
	}
	require.Equal(t, "el9", s.ImportName(pd, md))
}
