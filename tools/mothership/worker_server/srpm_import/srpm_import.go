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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	storage2 "github.com/go-git/go-git/v5/storage"
	"github.com/pkg/errors"
	"github.com/sassoftware/go-rpmutils"
	"go.resf.org/peridot/base/go/storage"
	"golang.org/x/crypto/openpgp"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	elDistRegex  = regexp.MustCompile(`el\d+`)
	releaseRegex = regexp.MustCompile(`.*release (\d+\.\d+).*`)
)

type State struct {
	// tempDir is the temporary directory where the SRPM is extracted to.
	tempDir string

	// rpm is the SRPM.
	rpm *rpmutils.Rpm

	// authorName is the name of the author of the commit.
	authorName string

	// authorEmail is the email of the author of the commit.
	authorEmail string

	// lookasideBlobs is a map of blob names to their SHA256 hashes.
	lookasideBlobs map[string]string

	// rolling determines how the branch is named.
	// if true, the branch is named "elX" where X is the major release
	// if false, the branch is named "el-X.Y" where X.Y is the full release
	rolling bool
}

// copyFromOS copies specified file from OS filesystem to target filesystem.
func copyFromOS(targetFS billy.Filesystem, path string, targetPath string) error {
	// Open file from OS filesystem.
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}
	defer f.Close()

	// Create file in target filesystem.
	targetFile, err := targetFS.Create(targetPath)
	if err != nil {
		return errors.Wrap(err, "failed to create file")
	}
	defer targetFile.Close()

	// Copy contents of file from OS filesystem to target filesystem.
	_, err = io.Copy(targetFile, f)
	if err != nil {
		return errors.Wrap(err, "failed to copy file")
	}

	return nil
}

// FromFile creates a new State from an SRPM file.
// The SRPM file is extracted to a temporary directory.
func FromFile(path string, rolling bool, keys ...*openpgp.Entity) (*State, error) {
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
		rpm:            rpm,
		authorName:     "Mship Bot",
		authorEmail:    "no-reply+mshipbot@resf.org",
		lookasideBlobs: make(map[string]string),
		rolling:        rolling,
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

func (s *State) SetAuthor(name, email string) {
	s.authorName = name
	s.authorEmail = email
}

// determineLookasideBlobs determines which blobs need to be uploaded to the
// lookaside cache.
// Currently, the rule is that if a file is larger than 5MB, and is binary,
// then it should be uploaded to the lookaside cache.
// If the file name contains ".tar", then it is assumed to be a tarball, and
// is ALWAYS uploaded to the lookaside cache.
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

		if info.Size() > 5*1024*1024 || strings.Contains(f.Name(), ".tar") {
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
		// First check if they exist, since it's a waste of time to upload
		// something that already exists.
		// They are uploaded by hash, so if the hash already exists, then the
		// file already exists.
		exists, err := lookaside.Exists(hash)
		if err != nil {
			return errors.Wrap(err, "failed to check if blob exists")
		}

		if exists {
			continue
		}

		_, err = lookaside.Put(hash, filepath.Join(s.tempDir, path))
		if err != nil {
			return errors.Wrap(err, "failed to upload file")
		}
	}

	return nil
}

// writeMetadata file writes the metadata map file.
// The metadata file contains lines of the format:
//
//	<path to download> <blob hash>
//
// For example:
//
//	1234567890abcdef SOURCES/bar
func (s *State) writeMetadataFile(targetFS billy.Filesystem) error {
	// Open metadata file for writing.
	name, err := s.rpm.Header.GetStrings(rpmutils.NAME)
	if err != nil {
		return errors.Wrap(err, "failed to get RPM name")
	}

	metadataFile := fmt.Sprintf(".%s.metadata", name[0])
	f, err := targetFS.Create(metadataFile)
	if err != nil {
		return errors.Wrap(err, "failed to open metadata file")
	}
	defer f.Close()

	// Write each line to the metadata file.
	for path, hash := range s.lookasideBlobs {
		// RPM sources MUST be in SOURCES/ directory
		_, err = f.Write([]byte(hash + " " + filepath.Join("SOURCES", path) + "\n"))
		if err != nil {
			return errors.Wrap(err, "failed to write line to metadata file")
		}
	}

	// Each file in metadata needs to be added to gitignore
	// Overwrite the gitignore file
	gitignoreFile := ".gitignore"
	f, err = targetFS.OpenFile(gitignoreFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to open gitignore file")
	}

	// Write each line to the gitignore file.
	for path, _ := range s.lookasideBlobs {
		_, err = f.Write([]byte(filepath.Join("SOURCES", path) + "\n"))
		if err != nil {
			return errors.Wrap(err, "failed to write line to gitignore file")
		}
	}

	return nil
}

// expandLayout expands the layout of the SRPM into the target filesystem.
// Moves all sources into SOURCES/ directory.
// Spec file is moved to SPECS/ directory.
func (s *State) expandLayout(targetFS billy.Filesystem) error {
	// Create SOURCES/ directory.
	err := targetFS.MkdirAll("SOURCES", 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create SOURCES directory")
	}

	// Copy all files from OS filesystem to target filesystem.
	ls, err := os.ReadDir(s.tempDir)
	if err != nil {
		return errors.Wrap(err, "failed to read directory")
	}

	for _, f := range ls {
		baseName := filepath.Base(f.Name())
		// If file ends with ".spec", then copy to SPECS/ directory.
		if strings.HasSuffix(f.Name(), ".spec") {
			err := copyFromOS(targetFS, filepath.Join(s.tempDir, f.Name()), filepath.Join("SPECS", baseName))
			if err != nil {
				return errors.Wrap(err, "failed to copy spec file")
			}
		} else {
			// Copy all other files to SOURCES/ directory.
			// Only if they are not present in lookasideBlobs
			_, ok := s.lookasideBlobs[f.Name()]
			if ok {
				continue
			}
			err := copyFromOS(targetFS, filepath.Join(s.tempDir, f.Name()), filepath.Join("SOURCES", baseName))
			if err != nil {
				return errors.Wrap(err, "failed to copy file")
			}
		}
	}

	return nil
}

// getRepo returns the target repository for the SRPM.
// This is where the payload is uploaded to.
func (s *State) getRepo(opts *git.CloneOptions, storer storage2.Storer, targetFS billy.Filesystem, osRelease string) (*git.Repository, string, error) {
	// Determine branch
	// If the OS release is not specified, then we use the dist tag
	var branch string
	if osRelease == "" {
		// Determine dist tag
		nevra, err := s.rpm.Header.GetNEVRA()
		if err != nil {
			return nil, "", errors.Wrap(err, "failed to get NEVRA")
		}

		// The dist tag will be used as the branch
		dist := elDistRegex.FindString(nevra.Release)
		if dist == "" {
			return nil, "", errors.Wrap(err, "failed to determine dist tag")
		}

		if s.rolling {
			branch = dist
		} else {
			branch = "el-" + dist[2:]
		}
	} else {
		// Determine branch from OS release
		if !releaseRegex.MatchString(osRelease) {
			return nil, "", fmt.Errorf("invalid OS release %s", osRelease)
		}
		ver := releaseRegex.FindStringSubmatch(osRelease)[1]

		if s.rolling {
			dist := elDistRegex.FindString("el" + ver)
			if dist == "" {
				return nil, "", errors.New("failed to determine dist tag")
			}
			branch = dist
		} else {
			branch = "el-" + ver
		}
	}

	// Set branch to dist tag
	opts.ReferenceName = plumbing.NewBranchReferenceName(branch)
	opts.SingleBranch = true

	// Clone the repository, to the target filesystem.
	// We do an init, then a fetch, then a checkout
	// If the repo doesn't exist, then we init only
	repo, err := git.Init(storer, targetFS)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to init repo")
	}
	wt, err := repo.Worktree()
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to get worktree")
	}

	// Create a new remote
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{opts.URL},
		Fetch: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%[1]s", branch)),
		},
	})
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to create remote")
	}

	// Fetch the remote
	err = repo.Fetch(&git.FetchOptions{
		Auth:       opts.Auth,
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%[1]s", branch)),
		},
	})

	// Checkout the branch
	refName := plumbing.NewBranchReferenceName(branch)

	if err != nil {
		h := plumbing.NewSymbolicReference(plumbing.HEAD, refName)
		if err := repo.Storer.CheckAndSetReference(h, nil); err != nil {
			return nil, "", errors.Wrap(err, "failed to checkout branch")
		}
	} else {
		err = wt.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(branch),
			Force:  true,
		})
	}

	return repo, branch, nil
}

// cleanTargetRepo deletes all files in the target repository.
func (s *State) cleanTargetRepo(wt *git.Worktree, root string) error {
	// Delete all files in the target repository.
	ls, err := wt.Filesystem.ReadDir(root)
	if err != nil {
		return errors.Wrap(err, "failed to read directory")
	}

	for _, f := range ls {
		// If it's a directory, then recurse into it.
		if f.IsDir() {
			err := s.cleanTargetRepo(wt, filepath.Join(root, f.Name()))
			if err != nil {
				return errors.Wrap(err, "failed to clean target repo")
			}
		} else {
			// Otherwise, delete the file.
			_, err := wt.Remove(filepath.Join(root, f.Name()))
			if err != nil {
				return errors.Wrap(err, "failed to remove file")
			}
		}
	}

	return nil
}

// populateTargetRepo runs the following steps:
// 1. Clean the target repository.
// 2. Determine which blobs need to be uploaded to the lookaside cache.
// 3. Upload blobs to the lookaside cache.
// 4. Write the metadata file.
// 5. Expand the layout of the SRPM.
// 6. Commit the changes to the target repository.
func (s *State) populateTargetRepo(repo *git.Repository, targetFS billy.Filesystem, lookaside storage.Storage, branch string) error {
	// Clean the target repository.
	wt, err := repo.Worktree()
	if err != nil {
		return errors.Wrap(err, "failed to get worktree")
	}

	err = s.cleanTargetRepo(wt, ".")
	if err != nil {
		return errors.Wrap(err, "failed to clean target repo")
	}

	// Determine which blobs need to be uploaded to the lookaside cache.
	err = s.determineLookasideBlobs()
	if err != nil {
		return errors.Wrap(err, "failed to determine lookaside blobs")
	}

	// Upload blobs to the lookaside cache.
	err = s.uploadLookasideBlobs(lookaside)
	if err != nil {
		return errors.Wrap(err, "failed to upload lookaside blobs")
	}

	// Write the metadata file.
	err = s.writeMetadataFile(targetFS)
	if err != nil {
		return errors.Wrap(err, "failed to write metadata file")
	}

	// Expand the layout of the SRPM.
	err = s.expandLayout(targetFS)
	if err != nil {
		return errors.Wrap(err, "failed to expand layout")
	}

	// Commit the changes to the target repository.
	_, err = wt.Add(".")
	if err != nil {
		return errors.Wrap(err, "failed to add files")
	}

	nevra, err := s.rpm.Header.GetNEVRA()
	if err != nil {
		return errors.Wrap(err, "failed to get NEVRA")
	}
	importStr := fmt.Sprintf("import %s-%s-%s", nevra.Name, nevra.Version, nevra.Release)
	hash, err := wt.Commit(importStr, &git.CommitOptions{
		Author: &object.Signature{
			Name:  s.authorName,
			Email: s.authorEmail,
			When:  time.Now(),
		},
		AllowEmptyCommits: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to commit changes")
	}

	// Create a tag
	// The tag should follow the following format:
	//   imports/<branch>/<nvra>
	tag := fmt.Sprintf("imports/%s/%s-%s-%s", branch, nevra.Name, nevra.Version, nevra.Release)
	_, err = repo.CreateTag(tag, hash, &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  s.authorName,
			Email: s.authorEmail,
			When:  time.Now(),
		},
		Message: tag,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create tag")
	}

	return nil
}

// pushTargetRepo pushes the target repository to the upstream repository.
func (s *State) pushTargetRepo(repo *git.Repository, opts *git.PushOptions) error {
	// Push the target repository to the upstream repository.
	err := repo.Push(opts)
	if err != nil {
		return errors.Wrap(err, "failed to push repo")
	}

	return nil
}

// Import imports the SRPM into the target repository.
func (s *State) Import(opts *git.CloneOptions, storer storage2.Storer, targetFS billy.Filesystem, lookaside storage.Storage, osRelease string) (*object.Commit, error) {
	// Get the target repository.
	repo, branch, err := s.getRepo(opts, storer, targetFS, osRelease)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get repo")
	}

	// Populate the target repository.
	err = s.populateTargetRepo(repo, targetFS, lookaside, branch)
	if err != nil {
		return nil, errors.Wrap(err, "failed to populate target repo")
	}

	// Push the target repository.
	err = s.pushTargetRepo(repo, &git.PushOptions{
		Force: true,
		Auth:  opts.Auth,
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%[1]s", branch)),
			config.RefSpec(fmt.Sprintf("refs/tags/imports/%s/*:refs/tags/imports/%[1]s/*", branch)),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to push target repo")
	}

	// Get latest commit
	head, err := repo.Head()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get HEAD")
	}

	// Get commit object
	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get commit object")
	}

	return commit, nil
}
