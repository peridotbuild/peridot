package main

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/pkg/errors"
	"go.resf.org/peridot/tools/kernelmanager/kernel_repack/kernelorg"
	repack_v1 "go.resf.org/peridot/tools/kernelmanager/kernel_repack/v1"
	"os"
	"time"
)

const exportDir = "/tmp/kernel_repack_test"

func main() {
	ltVersion, ltTarball, _, err := kernelorg.GetLatestLT("6.1.")
	if err != nil {
		panic(errors.Wrap(err, "failed to get latest LT version"))
	}

	// BuildID should be YYYYMMDDHHMM
	buildID := time.Now().Format("200601021504")

	out, err := repack_v1.LT(&repack_v1.Input{
		Version:       ltVersion,
		BuildID:       buildID,
		KernelPackage: "kernel-lt",
		Tarball:       ltTarball,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to repack LT kernel"))
	}

	// Create output directory, but first remove it if it already exists
	err = os.RemoveAll(exportDir)
	if err != nil {
		panic(errors.Wrap(err, "failed to remove output directory"))
	}

	err = os.MkdirAll(exportDir, 0755)
	if err != nil {
		panic(errors.Wrap(err, "failed to create output directory"))
	}

	err = out.ToFS(osfs.New(exportDir))
	if err != nil {
		panic(errors.Wrap(err, "failed to write output files"))
	}
}
