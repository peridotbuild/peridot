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
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	mothershipadmin_rpc "go.resf.org/peridot/tools/mothership/admin/rpc"
	"os"
)

func run(ctx *cli.Context) error {
	oidcInterceptorDetails := base.FlagsToOidcInterceptorDetails(ctx)

	s, err := mothershipadmin_rpc.NewServer(
		base.GetDBFromFlags(ctx),
		oidcInterceptorDetails,
		base.FlagsToGRPCServerOptions(ctx)...,
	)
	if err != nil {
		return err
	}
	return s.Start()
}

func main() {
	base.ChangeDefaultDatabaseURL("mothership")
	base.ChangeDefaultForEnvVar(base.EnvVarGRPCPort, "6687")
	base.ChangeDefaultForEnvVar(base.EnvVarGatewayPort, "6688")

	app := &cli.App{
		Name:   "mship_admin_server",
		Action: run,
		Flags:  base.WithDefaultCliFlags(),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_admin_server: %v", err)
	}
}
