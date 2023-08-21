package main

import (
	_ "embed"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"os"
)

//go:embed bundle.js
var bundleJs string

func run(ctx *cli.Context) error {
	info := &base.FrontendInfo{
		Title: "Mship",
	}
	base.FrontendServer(info, "bundle.js", bundleJs)

	return nil
}

func main() {
	base.ChangeDefaultForEnvVar(base.EnvVarFrontendAdminOIDCGroup, "sig-core-maintainer")

	app := &cli.App{
		Name:   "mship_ui",
		Action: run,
		Flags:  base.WithDefaultFrontendAdminCliFlags(),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_ui: %v", err)
	}
}
