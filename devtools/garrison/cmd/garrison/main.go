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
