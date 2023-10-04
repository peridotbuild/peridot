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

// Package main implements the dev server for Mothership.
// This runs the services just like it would be structured in production. (The RESF way)
// This means:
//   - localhost:9111 serves the external UI
//   - localhost:9111/admin serves the admin UI
//   - localhost:9111/api serves the external API
//   - localhost:9111/admin/api serves the admin API
package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/urfave/cli/v2"
	base "go.resf.org/peridot/base/go"
	mothershipadmin_rpc "go.resf.org/peridot/tools/mothership/admin/rpc"
	mship_admin_ui "go.resf.org/peridot/tools/mothership/admin/ui"
	mothership_rpc "go.resf.org/peridot/tools/mothership/rpc"
	mship_ui "go.resf.org/peridot/tools/mothership/ui"
	"go.temporal.io/sdk/client"
	"net/http"
	"os"
	"strings"
)

var (
	apiGrpcPort      = 3334
	adminApiGrpcPort = 3335
)

func setupUi(ctx *cli.Context) error {

	info := base.FlagsToFrontendInfo(ctx)
	assets := mship_ui.InitFrontendInfo(info)
	info.NoRun = true
	info.Self = "http://localhost:9111"

	return base.FrontendServer(info, assets)
}

func setupAdminUi(ctx *cli.Context) (*base.FrontendInfo, error) {
	info := base.FlagsToFrontendInfo(ctx)
	assets := mship_admin_ui.InitFrontendInfo(info)
	info.NoRun = true
	info.Self = "http://localhost:9111/admin"
	info.OIDCGroup = "authors"
	err := base.FrontendServer(info, assets)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setupApi(ctx *cli.Context) (*runtime.ServeMux, error) {
	temporalClient, err := base.GetTemporalClientFromFlags(ctx, client.Options{})
	if err != nil {
		return nil, err
	}

	s, err := mothership_rpc.NewServer(
		base.GetDBFromFlags(ctx),
		temporalClient,
		base.WithGRPCPort(apiGrpcPort),
		base.WithNoGRPCGateway(),
		base.WithNoMetrics(),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		err := s.Start()
		if err != nil {
			base.LogFatalf("failed to start mship_api: %v", err)
		}
	}()

	return s.GatewayMux(), nil
}

func setupAdminApi(ctx *cli.Context) (*runtime.ServeMux, error) {
	oidcInterceptorDetails, err := base.FlagsToOidcInterceptorDetails(ctx)
	if err != nil {
		return nil, err
	}
	oidcInterceptorDetails.Group = "authors"

	temporalClient, err := base.GetTemporalClientFromFlags(ctx, client.Options{})
	if err != nil {
		return nil, err
	}

	s, err := mothershipadmin_rpc.NewServer(
		base.GetDBFromFlags(ctx),
		temporalClient,
		oidcInterceptorDetails,
		base.WithGRPCPort(adminApiGrpcPort),
		base.WithNoGRPCGateway(),
		base.WithNoMetrics(),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		err := s.Start()
		if err != nil {
			base.LogFatalf("failed to start mship_admin_api: %v", err)
		}
	}()

	return s.GatewayMux(), nil
}

func run(ctx *cli.Context) error {
	err := setupUi(ctx)
	if err != nil {
		return err
	}

	adminInfo, err := setupAdminUi(ctx)
	if err != nil {
		return err
	}

	apiMux, err := setupApi(ctx)
	if err != nil {
		return err
	}
	http.HandleFunc("/api/", http.StripPrefix("/api", apiMux).ServeHTTP)

	adminApiMux, err := setupAdminApi(ctx)
	if err != nil {
		return err
	}
	http.HandleFunc("/admin/api/", http.StripPrefix("/admin/api", adminApiMux).ServeHTTP)

	handler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("x-peridot-beta", "true")

			if strings.HasPrefix(r.URL.Path, "/admin") {
				adminInfo.MuxHandler.ServeHTTP(w, r)
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}

	// Start server
	port := 9111
	base.LogInfof("Starting server on port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler(http.DefaultServeMux))
}

func main() {
	app := &cli.App{
		Name:   "mship_dev",
		Action: run,
		Flags: base.WithFlags(
			base.WithDatabaseFlags("mothership"),
			base.WithTemporalFlags("", "mship_worker_server"),
			base.WithFrontendAuthFlags(""),
		),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_dev: %v", err)
	}
}
