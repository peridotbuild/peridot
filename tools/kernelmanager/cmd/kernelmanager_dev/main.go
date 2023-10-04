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

// Package main implements the dev server for KernelManager.
// This runs the services just like it would be structured in production. (The RESF way)
// This means:
//   - localhost:9111 serves the UI
//   - localhost:9111/api serves the API
package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	"go.resf.org/peridot/base/go/kv/dynamodb"
	kernelmanager_rpc "go.resf.org/peridot/tools/kernelmanager/rpc"
	kernelmanager_ui "go.resf.org/peridot/tools/kernelmanager/ui"
	"go.temporal.io/sdk/client"
	"net/http"
	"os"
)

var (
	apiGrpcPort = 3734
)

func setupUi(ctx *cli.Context) (*base.FrontendInfo, error) {
	info := base.FlagsToFrontendInfo(ctx)
	assets := kernelmanager_ui.InitFrontendInfo(info)
	info.NoRun = true
	info.Self = "http://localhost:9111"

	err := base.FrontendServer(info, assets)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setupApi(ctx *cli.Context) (*runtime.ServeMux, error) {
	oidcInterceptorDetails, err := base.FlagsToOidcInterceptorDetails(ctx)
	if err != nil {
		return nil, err
	}
	oidcInterceptorDetails.AllowUnauthenticated = true

	temporalClient, err := base.GetTemporalClientFromFlags(ctx, client.Options{})
	if err != nil {
		return nil, err
	}

	kv, err := dynamodb.New("http://localhost:4566", "kernelmanager")
	if err != nil {
		return nil, err
	}

	s, err := kernelmanager_rpc.NewServer(
		kv,
		temporalClient,
		oidcInterceptorDetails,
		base.WithGRPCPort(apiGrpcPort),
		base.WithNoGRPCGateway(),
		base.WithNoMetrics(),
	)
	go func() {
		err := s.Start()
		if err != nil {
			base.LogFatalf("failed to start kernelmanager_api: %v", err)
		}
	}()

	return s.GatewayMux(), nil
}

func run(ctx *cli.Context) error {
	info, err := setupUi(ctx)
	if err != nil {
		return err
	}

	apiMux, err := setupApi(ctx)
	if err != nil {
		return err
	}
	http.HandleFunc("/api/", http.StripPrefix("/api", apiMux).ServeHTTP)

	handler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("x-peridot-beta", "true")

			h.ServeHTTP(w, r)
		})
	}

	// Start server
	port := 9111
	base.LogInfof("Starting server on port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler(info.MuxHandler))
}

func main() {
	app := &cli.App{
		Name:   "kernelmanager_dev",
		Action: run,
		Flags: base.WithFlags(
			base.WithDatabaseFlags("kernelmanager"),
			base.WithTemporalFlags("", "kernelmanager_queue"),
			base.WithFrontendAuthFlags(""),
		),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start kernelmanager_dev: %v", err)
	}
}
