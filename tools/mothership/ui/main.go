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
