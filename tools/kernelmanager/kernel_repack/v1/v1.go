package repack_v1

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"go.resf.org/peridot/tools/kernelmanager/kernel_repack"
	"io"
	"strings"
	"text/template"
)

//go:embed data/*
var Data embed.FS

type ChangelogEntry struct {
	Date    string
	Name    string
	Version string
	BuildID string
	Text    string
}

type Input struct {
	Version       string
	BuildID       string
	KernelPackage string
	Changelog     []*ChangelogEntry
	Tarball       []byte
}

func kernel(kernelType string, in *Input) (*kernel_repack.Output, error) {
	var spec *kernel_repack.File
	var files []*kernel_repack.File

	dir, err := Data.ReadDir("data")
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".spec") {
			if file.Name() != kernelType+".spec" {
				continue
			}

			// Read spec file
			f, err := Data.Open("data/" + file.Name())
			if err != nil {
				return nil, err
			}
			defer f.Close()

			specBytes, err := io.ReadAll(f)
			if err != nil {
				return nil, err
			}
			specName := fmt.Sprintf("%s.spec", in.KernelPackage)
			spec = &kernel_repack.File{
				Name: specName,
				Data: specBytes,
			}

			continue
		}

		// Read other files
		f, err := Data.Open("data/" + file.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}

		files = append(files, &kernel_repack.File{
			Name: file.Name(),
			Data: data,
		})
	}

	// Get sha256sum of tarball
	hash := sha256.New()
	_, err = hash.Write(in.Tarball)
	if err != nil {
		return nil, err
	}
	sum := hex.EncodeToString(hash.Sum(nil))

	// Create .metadata file
	suffix := "xz"
	if in.Tarball[0] == 0x1f && in.Tarball[1] == 0x8b {
		suffix = "gz"
	}
	metadata := fmt.Sprintf("%s SOURCES/linux-%s.tar.%s", sum, in.Version, suffix)

	// Create .[kernelpackage].metadata
	metadataName := fmt.Sprintf(".%s.metadata", in.KernelPackage)

	// Replace placeholders in spec file
	var buf bytes.Buffer
	txtTemplate, err := template.New("spec").Parse(string(spec.Data))
	if err != nil {
		return nil, err
	}
	err = txtTemplate.Execute(&buf, in)
	if err != nil {
		return nil, err
	}
	spec.Data = buf.Bytes()

	output := &kernel_repack.Output{
		Spec:          spec,
		Tarball:       in.Tarball,
		TarballSha256: sum,
		Metadata: &kernel_repack.File{
			Name: metadataName,
			Data: []byte(metadata),
		},
		OtherFiles: files,
	}

	return output, nil
}

// LT creates a new kernel package for the LT kernel
// Returns spec and SOURCE files
func LT(in *Input) (*kernel_repack.Output, error) {
	return kernel("lt", in)
}

// ML creates a new kernel package for the ML kernel
// Returns spec and SOURCE files
func ML(in *Input) (*kernel_repack.Output, error) {
	return kernel("ml", in)
}
