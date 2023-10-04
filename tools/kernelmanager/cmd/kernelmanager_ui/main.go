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
	kernelmanager_ui "go.resf.org/peridot/tools/kernelmanager/ui"
	"os"
)

func run(ctx *cli.Context) error {
	info := base.FlagsToFrontendInfo(ctx)
	assets := kernelmanager_ui.InitFrontendInfo(info)

	return base.FrontendServer(info, assets)
}

func main() {
	app := &cli.App{
		Name:   "kernelmanager_ui",
		Action: run,
		Flags:  base.WithFrontendFlags(9112),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start kernelmanager_ui: %v", err)
	}
}
