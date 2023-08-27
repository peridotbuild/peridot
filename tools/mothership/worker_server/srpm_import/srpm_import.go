package srpm_import

import (
    "crypto/sha256"
    "encoding/hex"
    "github.com/pkg/errors"
    "github.com/sassoftware/go-rpmutils"
    "go.resf.org/peridot/base/go/storage"
    "golang.org/x/crypto/openpgp"
    "io"
    "os"
    "path/filepath"
    "strings"
)

type State struct {
    tempDir string

    // lookasideBlobs is a map of blob names to their SHA256 hashes.
    lookasideBlobs map[string]string
}

// FromFile creates a new State from an SRPM file.
// The SRPM file is extracted to a temporary directory.
func FromFile(path string, keys ...*openpgp.Entity) (*State, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, errors.Wrap(err, "failed to open file")
    }
    defer f.Close()

    // If keys is not empty, then verify the RPM signature.
    if len(keys) > 0 {
        _, _, err := rpmutils.Verify(f, keys)
        if err != nil {
            return nil, errors.Wrap(err, "failed to verify RPM")
        }

        // After verifying the RPM, seek back to the start of the file.
        _, err = f.Seek(0, io.SeekStart)
        if err != nil {
            return nil, errors.Wrap(err, "failed to seek to start of file")
        }
    }

    rpm, err := rpmutils.ReadRpm(f)
    if err != nil {
        return nil, errors.Wrap(err, "failed to read RPM")
    }

    state := &State{
        lookasideBlobs: make(map[string]string),
    }
    // Create a temporary directory.
    state.tempDir, err = os.MkdirTemp("", "srpm_import-*")
    if err != nil {
        return nil, errors.Wrap(err, "failed to create temporary directory")
    }

    // Extract the SRPM.
    err = rpm.ExpandPayload(state.tempDir)
    if err != nil {
        return nil, errors.Wrap(err, "failed to extract SRPM")
    }

    return state, nil
}

func (s *State) Close() error {
    return os.RemoveAll(s.tempDir)
}

func (s *State) GetDir() string {
    return s.tempDir
}

// determineLookasideBlobs determines which blobs need to be uploaded to the
// lookaside cache.
// Currently, the rule is that if a file is larger than 5MB, and is binary,
// then it should be uploaded to the lookaside cache.
func (s *State) determineLookasideBlobs() error {
    // Read all files in tempDir, except for the SPEC file
    // For each file, if it is larger than 5MB, and is binary, then add it to
    // the lookasideBlobs map.
    // If the file is not binary, then skip it.
    ls, err := os.ReadDir(s.tempDir)
    if err != nil {
        return errors.Wrap(err, "failed to read directory")
    }

    for _, f := range ls {
        // If file ends with ".spec", then skip it.
        if f.IsDir() || strings.HasSuffix(f.Name(), ".spec") {
            continue
        }

        // If file is larger than 5MB, then add it to the lookasideBlobs map.
        info, err := f.Info()
        if err != nil {
            return errors.Wrap(err, "failed to get file info")
        }

        if info.Size() > 5*1024*1024 {
            sum, err := func() (string, error) {
                hash := sha256.New()
                file, err := os.Open(filepath.Join(s.tempDir, f.Name()))
                if err != nil {
                    return "", errors.Wrap(err, "failed to open file")
                }
                defer file.Close()

                _, err = io.Copy(hash, file)
                if err != nil {
                    return "", errors.Wrap(err, "failed to copy file")
                }

                return hex.EncodeToString(hash.Sum(nil)), nil
            }()
            if err != nil {
                return err
            }

            s.lookasideBlobs[f.Name()] = sum
        }
    }

    return nil
}

// uploadLookasideBlobs uploads all blobs in the lookasideBlobs map to the
// lookaside cache.
func (s *State) uploadLookasideBlobs(lookaside storage.Storage) error {
    // The object name is the SHA256 hash of the file.
    for path, hash := range s.lookasideBlobs {
        _, err := lookaside.Put(hash, filepath.Join(s.tempDir, path))
        if err != nil {
            return errors.Wrap(err, "failed to upload file")
        }
    }

    return nil
}
