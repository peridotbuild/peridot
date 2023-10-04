package main

import (
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/kv/dynamodb"
	kernelmanager_rpc "go.resf.org/peridot/tools/kernelmanager/rpc"
	"go.temporal.io/sdk/client"
	"os"
)

func run(ctx *cli.Context) error {
	oidcInterceptorDetails, err := base.FlagsToOidcInterceptorDetails(ctx)
	if err != nil {
		return err
	}
	oidcInterceptorDetails.AllowUnauthenticated = true

	temporalClient, err := base.GetTemporalClientFromFlags(ctx, client.Options{})
	if err != nil {
		return err
	}

	kv, err := dynamodb.New(ctx.String("dynamodb-table"))
	if err != nil {
		return err
	}

	s, err := kernelmanager_rpc.NewServer(
		kv,
		temporalClient,
		oidcInterceptorDetails,
		base.FlagsToGRPCServerOptions(ctx)...,
	)
	if err != nil {
		return err
	}
	return s.Start()
}

func main() {
	app := &cli.App{
		Name:   "kernelmanager_server",
		Action: run,
		Flags: base.WithFlags(
			base.WithDatabaseFlags("kernelmanager"),
			base.WithTemporalFlags("", "kernelmanager_queue"),
			base.WithGrpcFlags(6677),
			base.WithGatewayFlags(6678),
			base.WithOidcFlags("", "releng"),
			[]cli.Flag{
				&cli.StringFlag{
					Name:    "dynamodb-table",
					Usage:   "DynamoDB table name",
					EnvVars: []string{"DYNAMODB_TABLE"},
					Value:   "kernelmanager",
				},
			},
		),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_api: %v", err)
	}
}
