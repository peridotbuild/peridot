package main

import (
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	storage_detector "go.resf.org/peridot/base/go/storage/detector"
	mothership_worker_server "go.resf.org/peridot/tools/mothership/worker_server"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"os"
)

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

	w := worker.New(temporalClient, ctx.String(string(base.EnvVarTemporalTaskQueue)), worker.Options{})
	workerServer := mothership_worker_server.New(db, storage)

	// Register workflows
	w.RegisterWorkflow(mothership_worker_server.ProcessRPMWorkflow)

	// Register activities
	w.RegisterActivity(workerServer)

	// Start worker
	return w.Run(worker.InterruptCh())
}

func main() {
	base.ChangeDefaultForEnvVar(base.EnvVarTemporalTaskQueue, "mship_worker_server")

	app := &cli.App{
		Name:   "mship_worker_server",
		Action: run,
		Flags:  base.WithDefaultCliFlagsTemporal(base.WithStorageFlags()...),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to run mship_worker_server: %v", err)
	}
}
