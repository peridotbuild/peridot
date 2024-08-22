package packager

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/pkg/errors"
)

var (
	ErrTarballEmpty        = errors.New("tarball is empty")
	ErrTarballHashMismatch = errors.New("tarball hash mismatch")
)

type File struct {
	Name        string
	Data        []byte
	Permissions os.FileMode
}

type Output struct {
	Spec          *File
	Tarball       []byte
	TarballName   string
	TarballSha256 string
	Metadata      *File
	OtherFiles    []*File
}

func (o *Output) ToFS(fs billy.Filesystem) error {
	// Verify that tarball is not empty
	if len(o.Tarball) == 0 {
		return ErrTarballEmpty
	}

	// Verify that tarball hash matches
	hash := sha256.New()
	_, err := hash.Write(o.Tarball)
	if err != nil {
		return err
	}
	sum := hash.Sum(nil)
	if o.TarballSha256 != hex.EncodeToString(sum) {
		return errors.Wrap(ErrTarballHashMismatch, fmt.Sprintf("expected %s, got %s", o.TarballSha256, hex.EncodeToString(sum)))
	}

	// Verify that spec file is not empty
	if o.Spec == nil || len(o.Spec.Data) == 0 {
		return errors.New("spec file is empty")
	}

	// Verify that the metadata file has a name
	if o.Metadata == nil || len(o.Metadata.Name) == 0 {
		return errors.New("metadata file has no name")
	}

	// Create directories first
	dirs := []string{"SPECS", "SOURCES"}
	for _, dir := range dirs {
		err := fs.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// Create SOURCES files
	for _, file := range o.OtherFiles {
		f, err := fs.OpenFile("SOURCES/"+file.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Permissions)
		if err != nil {
			return err
		}

		_, err = f.Write(file.Data)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}
	}

	// Create SPEC file
	f, err := fs.OpenFile("SPECS/"+o.Spec.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(o.Spec.Data)
	if err != nil {
		return err
	}

	// Create metadata file
	f, err = fs.OpenFile(o.Metadata.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(o.Metadata.Data)
	if err != nil {
		return err
	}

	return nil
}
