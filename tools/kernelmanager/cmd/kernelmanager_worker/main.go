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
	_ "embed"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/forge/gitlab"
	"go.resf.org/peridot/base/go/kv/dynamodb"
	storage_detector "go.resf.org/peridot/base/go/storage/detector"
	kernelmanager_worker "go.resf.org/peridot/tools/kernelmanager/worker"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"os"
)

func run(ctx *cli.Context) error {
	temporalClient, err := base.GetTemporalClientFromFlags(ctx, client.Options{})
	if err != nil {
		return err
	}

	kv, err := dynamodb.New(ctx.String("dynamodb-endpoint"), ctx.String("dynamodb-table"))
	if err != nil {
		return err
	}

	gitlabForge := gitlab.New(
		ctx.String("gitlab-host"),
		"",
		ctx.String("gitlab-username"),
		ctx.String("gitlab-password"),
		"RESF KernelManager",
		"releng+kernelmanager@rockylinux.org",
		true,
	)

	st, err := storage_detector.FromFlags(ctx)
	if err != nil {
		return err
	}

	w := worker.New(temporalClient, ctx.String("temporal-task-queue"), worker.Options{})
	workerServer := kernelmanager_worker.New(
		kv,
		gitlabForge,
		st,
	)

	// Register workflows
	w.RegisterWorkflow(kernelmanager_worker.TriggerKernelUpdateWorkflow)

	// Register activities
	w.RegisterActivity(workerServer)

	// Start worker
	return w.Run(worker.InterruptCh())
}

func main() {
	flags := base.WithFlags(
		base.WithTemporalFlags("", "kernelmanager_queue"),
		base.WithStorageFlags(),
		[]cli.Flag{
			&cli.StringFlag{
				Name:    "dynamodb-endpoint",
				Usage:   "DynamoDB endpoint",
				EnvVars: []string{"DYNAMODB_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:    "dynamodb-table",
				Usage:   "DynamoDB table name",
				EnvVars: []string{"DYNAMODB_TABLE"},
				Value:   "kernelmanager",
			},
			&cli.StringFlag{
				Name:    "gitlab-host",
				Usage:   "GitLab host",
				EnvVars: []string{"GITLAB_HOST"},
				Value:   "git.rockylinux.org",
			},
			&cli.StringFlag{
				Name:    "gitlab-username",
				Usage:   "GitLab username",
				EnvVars: []string{"GITLAB_USERNAME"},
			},
			&cli.StringFlag{
				Name:    "gitlab-password",
				Usage:   "GitLab password",
				EnvVars: []string{"GITLAB_PASSWORD"},
			},
		},
	)

	app := &cli.App{
		Name:   "kernelmanager_worker",
		Action: run,
		Flags:  flags,
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to run kernelmanager_worker: %v", err)
	}
}
