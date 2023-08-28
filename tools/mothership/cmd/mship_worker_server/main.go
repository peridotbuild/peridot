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

package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	storage_detector "go.resf.org/peridot/base/go/storage/detector"
	mothership_worker_server "go.resf.org/peridot/tools/mothership/worker_server"
	"go.resf.org/peridot/tools/mothership/worker_server/forge"
	github_forge "go.resf.org/peridot/tools/mothership/worker_server/forge/github"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"golang.org/x/crypto/openpgp"
	"os"
)

//go:embed rh_public_key.asc
var defaultGpgKey []byte

func run(ctx *cli.Context) error {
	temporalClient, err := base.GetTemporalClientFromFlags(ctx, client.Options{})
	if err != nil {
		return err
	}

	db := base.GetDBFromFlags(ctx)
	storage, err := storage_detector.FromFlags(ctx)
	if err != nil {
		return err
	}

	// Create pgp keys
	var gpgKeys openpgp.EntityList
	for _, key := range ctx.StringSlice("allowed-gpg-keys") {
		decoded, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return err
		}
		keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader([]byte(decoded)))
		if err != nil {
			return err
		}

		gpgKeys = append(gpgKeys, keyRing...)
	}

	// Create forge based on git provider
	var remoteForge forge.Forge
	switch ctx.String("git-provider") {
	case "github":
		var appPrivateKey []byte
		if ctx.Bool("github-app-private-key-base64") {
			appPrivateKey, err = base64.StdEncoding.DecodeString(ctx.String("github-app-private-key"))
			if err != nil {
				return err
			}
		} else {
			appPrivateKey = []byte(ctx.String("github-app-private-key"))
		}

		remoteForge, err = github_forge.New(
			ctx.String("github-org"),
			ctx.String("github-app-id"),
			appPrivateKey,
			ctx.Bool("github-make-repo-public"),
		)
		if err != nil {
			return err
		}
		remoteForge = forge.NewCacher(remoteForge)
	default:
		return cli.Exit("git-provider must be github", 1)
	}

	w := worker.New(temporalClient, ctx.String("temporal-task-queue"), worker.Options{})
	workerServer := mothership_worker_server.New(db, storage, gpgKeys, remoteForge)

	// Register workflows
	w.RegisterWorkflow(mothership_worker_server.ProcessRPMWorkflow)

	// Register activities
	w.RegisterActivity(workerServer)

	// Start worker
	return w.Run(worker.InterruptCh())
}

func main() {
	base.ChangeDefaultForEnvVar(base.EnvVarTemporalTaskQueue, "mship_worker_server")

	flags := base.WithDefaultCliFlagsTemporal(base.WithStorageFlags()...)
	flags = append(flags, &cli.StringSliceFlag{
		Name:    "allowed-gpg-keys",
		Usage:   "Armored GPG keys that we verify SRPMs with. Must be base64 encoded",
		EnvVars: []string{"ALLOWED_GPG_KEYS"},
	})
	flags = append(flags, []cli.Flag{
		&cli.StringFlag{
			Name: "git-provider",
			Action: func(ctx *cli.Context, s string) error {
				// Can only be github for now
				if s != "github" {
					return cli.Exit("git-provider must be github", 1)
				}

				return nil
			},
			Usage:   "Git provider to use. Currently only github is supported",
			EnvVars: []string{"GIT_PROVIDER"},
		},
		// Github only
		&cli.StringFlag{
			Name:    "github-org",
			Usage:   "Github organization to use",
			EnvVars: []string{"GITHUB_ORG"},
			Action: func(ctx *cli.Context, s string) error {
				// Required for github
				if ctx.String("git-provider") == "github" && s == "" {
					return cli.Exit("github-org is required for github", 1)
				}

				return nil
			},
		},
		&cli.StringFlag{
			Name:    "github-app-id",
			Usage:   "Github app ID",
			EnvVars: []string{"GITHUB_APP_ID"},
			Action: func(ctx *cli.Context, s string) error {
				// Required for github
				if ctx.String("git-provider") == "github" && s == "" {
					return cli.Exit("github-org is required for github", 1)
				}

				return nil
			},
		},
		&cli.StringFlag{
			Name:    "github-app-private-key",
			Usage:   "Github app private key",
			EnvVars: []string{"GITHUB_APP_PRIVATE_KEY"},
			Action: func(ctx *cli.Context, s string) error {
				// Required for github
				if ctx.String("git-provider") == "github" && s == "" {
					return cli.Exit("github-org is required for github", 1)
				}

				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "github-app-private-key-base64",
			Usage:   "Whether the Github app private key is base64 encoded",
			EnvVars: []string{"GITHUB_APP_PRIVATE_KEY_BASE64"},
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "github-make-repo-public",
			Usage:   "Whether to make the Github repository public",
			EnvVars: []string{"GITHUB_MAKE_REPO_PUBLIC"},
			Value:   false,
		},
	}...)

	base64EncodedDefaultGpgKey := base64.StdEncoding.EncodeToString(defaultGpgKey)
	base.RareUseChangeDefault("ALLOWED_GPG_KEYS", base64EncodedDefaultGpgKey)

	app := &cli.App{
		Name:   "mship_worker_server",
		Action: run,
		Flags:  flags,
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to run mship_worker_server: %v", err)
	}
}
