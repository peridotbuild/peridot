package kernelmanager_worker

import (
	"context"
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/tools/kernelmanager/kernel_repack"
	"go.resf.org/peridot/tools/kernelmanager/kernel_repack/kernelorg"
	repack_v1 "go.resf.org/peridot/tools/kernelmanager/kernel_repack/v1"
	kernelmanagerpb "go.resf.org/peridot/tools/kernelmanager/pb"
	"golang.org/x/crypto/openpgp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func (w *Worker) GetKernel(ctx context.Context, name string) (*kernelmanagerpb.Kernel, error) {
	kernelBytes, err := w.kv.Get(ctx, fmt.Sprintf("/kernels/entries/%s", name))
	if err != nil {
		base.LogErrorf("failed to get kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to get kernel")
	}

	kernel := &kernelmanagerpb.Kernel{}
	err = proto.Unmarshal(kernelBytes.Value, kernel)
	if err != nil {
		base.LogErrorf("failed to unmarshal kernel: %v", err)
		return nil, status.Error(codes.Internal, "failed to unmarshal kernel")
	}

	return kernel, nil
}

func (w *Worker) KernelRepack(ctx context.Context, kernel *kernelmanagerpb.Kernel) (*kernelmanagerpb.Update, error) {
	gitForge := w.forge.WithNamespace(kernel.Config.ScmNamespace)
	gitRemote := gitForge.GetRemote(kernel.Pkg)
	gitAuth, err := gitForge.GetAuthenticator()
	if err != nil {
		return nil, err
	}

	err = gitForge.EnsureRepositoryExists(gitAuth, kernel.Pkg)
	if err != nil {
		return nil, err
	}

	// Clone the repository, to the target filesystem.
	// We do an init, then a fetch, then a checkout
	// If the repo doesn't exist, then we init only
	storer := memory.NewStorage()
	fs := memfs.New()
	repo, err := git.Init(storer, fs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init git repo")
	}
	wt, err := repo.Worktree()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get worktree")
	}

	// Create refspecs list
	var refspecs []config.RefSpec
	for _, branch := range kernel.Config.ScmBranches {
		refspecs = append(refspecs, config.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%[1]s", branch)))
	}

	// Create a new remote
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name:  "origin",
		URLs:  []string{gitRemote},
		Fetch: refspecs,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create remote")
	}

	// Fetch the remote
	_ = repo.Fetch(&git.FetchOptions{
		Auth:       gitAuth.AuthMethod,
		RemoteName: "origin",
		RefSpecs:   refspecs,
	})

	// Checkout the branch
	// refName := plumbing.NewBranchReferenceName(branch)

	// if err != nil {
	//     h := plumbing.NewSymbolicReference(plumbing.HEAD, refName)
	//     if err := repo.Storer.CheckAndSetReference(h, nil); err != nil {
	//         return nil, "", errors.Wrap(err, "failed to checkout branch")
	//     }
	// } else {
	//     err = wt.Checkout(&git.CheckoutOptions{
	//         Branch: plumbing.NewBranchReferenceName(branch),
	//         Force:  true,
	//     })
	// }

	var output *kernel_repack.Output
	var entity *openpgp.Entity
	var version string

	// BuildID should be YYYYMMDDHHMM
	buildID := time.Now().Format("200601021504")

	changelog := func(version string) []*repack_v1.ChangelogEntry {
		msg := fmt.Sprintf("Rebase to %s", version)
		return []*repack_v1.ChangelogEntry{
			{
				Date:    time.Now().Format("Mon Jan 02 2006"),
				Name:    "Mustafa Gezen",
				Version: version,
				BuildID: buildID,
				Text:    msg,
			},
		}
	}

	switch kernel.Config.RepackOptions.KernelOrgVariant {
	case kernelmanagerpb.RepackOptions_MAINLINE:
		mlVersion, mlTarball, _, err := kernelorg.GetLatestStable()
		if err != nil {
			return nil, err
		}
		version = mlVersion

		out, err := repack_v1.ML(&repack_v1.Input{
			Version:                mlVersion,
			BuildID:                buildID,
			KernelPackage:          kernel.Pkg,
			Tarball:                mlTarball,
			Changelog:              changelog(mlVersion),
			AdditionalKernelConfig: kernel.Config.RepackOptions.AdditionalKernelConfig,
		})
		if err != nil {
			return nil, err
		}

		output = out
	case kernelmanagerpb.RepackOptions_LONGTERM:
		repackVersion := kernel.Config.RepackOptions.KernelOrgVersion
		if !strings.HasSuffix(repackVersion, ".") {
			repackVersion = repackVersion + "."
		}
		ltVersion, ltTarball, ltEntity, err := kernelorg.GetLatestLT(repackVersion)
		if err != nil {
			return nil, err
		}
		version = ltVersion

		out, err := repack_v1.LT(&repack_v1.Input{
			Version:                ltVersion,
			BuildID:                buildID,
			KernelPackage:          kernel.Pkg,
			Tarball:                ltTarball,
			Changelog:              changelog(ltVersion),
			AdditionalKernelConfig: kernel.Config.RepackOptions.AdditionalKernelConfig,
		})
		if err != nil {
			return nil, err
		}

		output = out
		entity = ltEntity
	}

	// Upload tarball to S3

	// Check out each branch, delete all files in SOURCES, then extract to FS.
	for _, branch := range kernel.Config.ScmBranches {
		refName := plumbing.NewBranchReferenceName(branch)

		err := wt.Checkout(&git.CheckoutOptions{
			Branch: refName,
		})
		if err != nil {
			h := plumbing.NewSymbolicReference(plumbing.HEAD, refName)
			if err := repo.Storer.CheckAndSetReference(h, nil); err != nil {
				return nil, errors.Wrap(err, "failed to checkout branch")
			}
		}

		// Delete all files in SOURCES
		files, err := fs.ReadDir("SOURCES")
		if err != nil {
			return nil, errors.Wrap(err, "failed to read SOURCES directory")
		}

		for _, file := range files {
			err = fs.Remove(fmt.Sprintf("SOURCES/%s", file.Name()))
			if err != nil {
				return nil, errors.Wrap(err, "failed to remove file")
			}
		}

		// Extract to FS
		err = output.ToFS(fs)
		if err != nil {
			return nil, errors.Wrap(err, "failed to write output files")
		}

		// Commit changes
		_, err = wt.Add(".")
		if err != nil {
			return nil, errors.Wrap(err, "failed to add files to git")
		}

		msg := fmt.Sprintf("Repacking %s from kernel.org - %s - %s", kernel.Pkg, version, buildID)
		_, err = wt.Commit(msg, &git.CommitOptions{
			AllowEmptyCommits: true,
			Author: &object.Signature{
				Name:  gitAuth.AuthorName,
				Email: gitAuth.AuthorEmail,
				When:  time.Now(),
			},
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to commit changes")
		}
	}

	// Push changes
	err = repo.Push(&git.PushOptions{
		Auth:       gitAuth.AuthMethod,
		RemoteName: "origin",
		RefSpecs:   refspecs,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to push changes")
	}

	// Create update
	update := &kernelmanagerpb.Update{
		Kernel:                      kernel,
		KernelOrgTarballSha256:      output.TarballSha256,
		KernelOrgTarballPgpIdentity: "",
		KernelOrgVersion:            version,
		FinishedTime:                timestamppb.New(time.Now()),
	}
	if entity != nil {
		update.KernelOrgTarballPgpIdentity = entity.PrimaryKey.KeyIdShortString()
	}

	return update, nil
}
