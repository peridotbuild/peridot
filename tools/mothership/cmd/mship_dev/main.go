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
	err := base.FrontendServer(info, assets)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setupApi(ctx *cli.Context) (*runtime.ServeMux, error) {
	s, err := mothership_rpc.NewServer(
		base.GetDBFromFlags(ctx),
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
	oidcInterceptorDetails := base.FlagsToOidcInterceptorDetails(ctx)

	s, err := mothershipadmin_rpc.NewServer(
		base.GetDBFromFlags(ctx),
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
	base.ChangeDefaultDatabaseURL("mothership")

	app := &cli.App{
		Name:   "mship_dev",
		Action: run,
		Flags:  base.WithDefaultCliFlagsNoAuth(base.WithDefaultFrontendCliFlags()...),
	}

	if err := app.Run(os.Args); err != nil {
		base.LogFatalf("failed to start mship_ui: %v", err)
	}
}
