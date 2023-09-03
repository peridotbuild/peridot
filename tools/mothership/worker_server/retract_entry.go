package mothership_worker_server

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	base "go.resf.org/peridot/base/go"
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	"go.temporal.io/sdk/temporal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"regexp"
)

var pkgNameRegexp = regexp.MustCompile(`^([a-zA-Z0-9\-]+)-\d+.*$`)

// getRepo gets a git repository from a remote
// It clones into an in-memory filesystem
func getRepo(remote string) (*git.Repository, error) {
	// Just use in memory storage for all repos
	storer := memory.NewStorage()
	fs := memfs.New()
	repo, err := git.Init(storer, fs)
	if err != nil {
		return nil, err
	}

	// Add a new remote
	refspec := config.RefSpec("refs/heads/*:refs/remotes/origin/*")
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name:  "origin",
		URLs:  []string{remote},
		Fetch: []config.RefSpec{refspec},
	})
	if err != nil {
		return nil, err
	}

	// Fetch all the refs from the remote
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
	})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// clonePathToFS clones a path from one filesystem to another
func clonePathToFS(fromFS billy.Filesystem, toFS billy.Filesystem, rootPath string) error {
	// check if root directory exists
	_, err := fromFS.Stat(rootPath)
	if err != nil {
		// we don't care if the directory doesn't exist
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	// read the root directory
	rootDir, err := fromFS.ReadDir(rootPath)
	if err != nil {
		return err
	}

	// iterate over the files
	for _, file := range rootDir {
		// get the file path
		filePath := rootPath + "/" + file.Name()

		// check if the file is a directory
		if file.IsDir() {
			// create the directory in the toFS
			err = toFS.MkdirAll(filePath, 0755)
			if err != nil {
				return err
			}

			// recursively call this function
			err = clonePathToFS(fromFS, toFS, filePath)
			if err != nil {
				return err
			}
		} else {
			// open the file
			f, err := fromFS.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				return err
			}
			defer f.Close()

			// create the file in the toFS
			toFile, err := toFS.Create(filePath)
			if err != nil {
				return err
			}
			defer toFile.Close()

			// copy the file contents
			_, err = io.Copy(toFile, f)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// clonePatchesToTemporaryFS clones the PATCHES directory to a temporary filesystem
// PATCHES directory is the only directory that should survive a retraction
func clonePatchesToTemporaryFS(currentFS billy.Filesystem) (billy.Filesystem, error) {
	// create a new in-memory filesystem
	fs := memfs.New()

	// clone the current filesystem to the new filesystem
	err := clonePathToFS(currentFS, fs, "PATCHES")
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func resetRepoToPoint(repo *git.Repository, commit string) error {
	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	// reset the repo
	err = wt.Reset(&git.ResetOptions{
		Commit: plumbing.NewHash(commit),
		Mode:   git.HardReset,
	})
	if err != nil {
		return err
	}

	// clean the repo
	return wt.Clean(&git.CleanOptions{
		Dir: true,
	})
}

func (w *Worker) RetractEntry(name string) (*mshipadminpb.RetractEntryResponse, error) {
	entry, err := base.Q[mothership_db.Entry](w.db).F("name", name).GetOrNil()
	if err != nil {
		base.LogErrorf("failed to get entry: %v", err)
		return nil, status.Error(codes.Internal, "failed to get entry")
	}

	if entry == nil {
		return nil, temporal.NewNonRetryableApplicationError(
			"entry not found",
			"entryNotFound",
			nil,
		)
	}

	// Get the repo
	remote := w.forge.GetRemote(entry.PackageName)
	repo, err := getRepo(remote)
	if err != nil {
		base.LogErrorf("failed to get repo: %v", err)
		return nil, status.Error(codes.Internal, "failed to get repo")
	}

	err = repo.DeleteObject(plumbing.NewHash(entry.CommitHash))
	if err != nil {
		base.LogErrorf("failed to eject commit: %v", err)
		return nil, status.Error(codes.Internal, "failed to eject commit")
	}

	// If there's a tag, delete it
	_ = repo.DeleteTag(entry.CommitTag)

	// Push the changes
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Force:      true,
	})
	if err != nil {
		base.LogErrorf("failed to push changes: %v", err)
		return nil, status.Error(codes.Internal, "failed to push changes")
	}

	return &mshipadminpb.RetractEntryResponse{
		Name: entry.Name,
	}, nil
}
