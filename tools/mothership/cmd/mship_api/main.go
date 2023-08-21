package main

import (
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	mothership_rpc "go.resf.org/peridot/tools/mothership/rpc"
	"os"
)

func run(ctx *cli.Context) error {
	s := mothership_rpc.NewServer(base.GetDBFromFlags(ctx))
	return s.Start(base.FlagsToGRPCServerOptions(ctx)...)
}

func main() {
	base.ChangeDefaultDatabaseURL("mothership")
	base.ChangeDefaultForEnvVar(base.EnvVarGRPCPort, "6677")
	base.ChangeDefaultForEnvVar(base.EnvVarGatewayPort, "6678")

	app := &cli.App{
		Name:   "mship_api",
		Action: run,
		Flags:  base.WithDefaultCliFlags(),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_api: %v", err)
	}
}
