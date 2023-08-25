package main

import (
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	mship_ui "go.resf.org/peridot/tools/mothership/ui"
	"os"
)

func run(ctx *cli.Context) error {
	info := base.FlagsToFrontendInfo(ctx)
	assets := mship_ui.InitFrontendInfo(info)

	return base.FrontendServer(info, assets)
}

func main() {
	app := &cli.App{
		Name:   "mship_ui",
		Action: run,
		Flags:  base.WithDefaultFrontendNoAuthCliFlags(),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_ui: %v", err)
	}
}
