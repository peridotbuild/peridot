package storage_memory

import (
	"github.com/go-git/go-billy/v5"
	"github.com/pkg/errors"
	"go.resf.org/peridot/base/go/storage"
	"io"
	"os"
	"strings"
)

type InMemory struct {
	storage.Storage

	fs    billy.Filesystem
	blobs map[string][]byte
}

func New(fs billy.Filesystem) *InMemory {
	return &InMemory{
		fs:    fs,
		blobs: make(map[string][]byte),
	}
}

func (im *InMemory) Download(object string, toPath string) error {
	blob, ok := im.blobs[object]
	if !ok {
		return storage.ErrNotFound
	}

	// Open file
	f, err := im.fs.OpenFile(toPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}

	// Write blob to file
	_, err = f.Write(blob)
	if err != nil {
		return errors.Wrap(err, "failed to write blob to file")
	}

	return nil
}

func (im *InMemory) Get(object string) ([]byte, error) {
	blob, ok := im.blobs[object]
	if !ok {
		return nil, storage.ErrNotFound
	}
	return blob, nil
}

func (im *InMemory) Put(object string, fromPath string) (*storage.UploadInfo, error) {
	// Open file
	f, err := im.fs.Open(fromPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	// Read file into blob
	blob, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	// Store blob
	im.blobs[object] = blob

	return &storage.UploadInfo{
		Location:  "memory://" + object,
		VersionID: nil,
	}, nil
}

func (im *InMemory) PutBytes(object string, blob []byte) (*storage.UploadInfo, error) {
	// Store blob
	im.blobs[object] = blob

	return &storage.UploadInfo{
		Location:  "memory://" + object,
		VersionID: nil,
	}, nil
}

func (im *InMemory) Delete(object string) error {
	delete(im.blobs, object)
	return nil
}

func (im *InMemory) Exists(object string) (bool, error) {
	_, ok := im.blobs[object]
	return ok, nil
}

func (im *InMemory) CanReadURI(uri string) (bool, error) {
	return strings.HasPrefix(uri, "memory://"), nil
}
