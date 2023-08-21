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
	"fmt"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/devtools/garrison/internal/pkg/kubernetes"
	"go.resf.org/peridot/devtools/garrison/internal/pkg/starlark"
	"log"
	"os"
)

func run(ctx *cli.Context) error {
	stage := starlark.Stage(ctx.String("stage"))
	if !base.Contains(starlark.GetValidStages(), stage) {
		return cli.Exit("invalid stage", 1)
	}

	state, err := starlark.ExecFile(ctx.Args().First(), stage)
	if err != nil {
		return cli.Exit(err, 1)
	}

	k8sState := kubernetes.NewState(state)
	err = k8sState.GenerateManifests()
	if err != nil {
		return cli.Exit(err, 1)
	}

	out, err := k8sState.YAML()
	if err != nil {
		return cli.Exit(err, 1)
	}

	fmt.Println(out)

	return nil
}

func main() {
	app := &cli.App{
		Name:  "garrison",
		Usage: "Deployment framework for the RESF",
		Commands: []*cli.Command{
			{
				Name:   "build",
				Usage:  "Build deployment manifests from a starlark file",
				Action: run,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "stage",
				Aliases: []string{"s"},
				Usage:   "stage to deploy to",
				Value:   "development",
			},
			&cli.StringFlag{
				Name:  "output-mode",
				Usage: "output mode",
				Value: "kubernetes",
				Action: func(ctx *cli.Context, s string) error {
					// todo(mustafa): add helm support
					if s != "kubernetes" /*&& s != "helm"*/ {
						return cli.Exit("invalid output mode", 1)
					}

					return nil
				},
			},
			&cli.StringFlag{
				Name:  "output-dir",
				Usage: "output directory",
				Value: "output",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
