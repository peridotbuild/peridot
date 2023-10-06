package kernel_repack

import (
	"github.com/go-git/go-billy/v5"
	"os"
)

type File struct {
	Name        string
	Data        []byte
	Permissions os.FileMode
}

type Output struct {
	Spec          *File
	Tarball       []byte
	TarballSha256 string
	Metadata      *File
	OtherFiles    []*File
}

func (o *Output) ToFS(fs billy.Filesystem) error {
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
