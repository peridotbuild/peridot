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

package base

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"os"
)

type EnvVar string

const (
	EnvVarGRPCPort                     EnvVar = "GRPC_PORT"
	EnvVarGatewayPort                  EnvVar = "GATEWAY_PORT"
	EnvVarDatabaseURL                  EnvVar = "DATABASE_URL"
	EnvVarFrontendPort                 EnvVar = "FRONTEND_PORT"
	EnvVarFrontendOIDCIssuer           EnvVar = "FRONTEND_OIDC_ISSUER"
	EnvVarFrontendOIDCClientID         EnvVar = "FRONTEND_OIDC_CLIENT_ID"
	EnvVarFrontendOIDCClientSecret     EnvVar = "FRONTEND_OIDC_CLIENT_SECRET"
	EnvVarFrontendOIDCUserInfoOverride EnvVar = "FRONTEND_OIDC_USERINFO_OVERRIDE"
	EnvVarFrontendRequiredOIDCGroup    EnvVar = "FRONTEND_REQUIRED_OIDC_GROUP"
	EnvVarTemporalNamespace            EnvVar = "TEMPORAL_NAMESPACE"
	EnvVarTemporalAddress              EnvVar = "TEMPORAL_ADDRESS"
	EnvVarTemporalTaskQueue            EnvVar = "TEMPORAL_TASK_QUEUE"
	EnvVarFrontendSelf                 EnvVar = "FRONTEND_SELF"
	EnvVarStorageEndpoint              EnvVar = "STORAGE_ENDPOINT"
	EnvVarStorageConnectionString      EnvVar = "STORAGE_CONNECTION_STRING"
	EnvVarStorageRegion                EnvVar = "STORAGE_REGION"
	EnvVarStorageSecure                EnvVar = "STORAGE_SECURE"
	EnvVarStoragePathStyle             EnvVar = "STORAGE_PATH_STYLE"
)

var defaultCliFlagsDatabaseOnly = []cli.Flag{
	&cli.StringFlag{
		Name:    "database-url",
		Aliases: []string{"d"},
		Usage:   "database url",
		EnvVars: []string{string(EnvVarDatabaseURL)},
		Value:   "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	},
}

var defaultCliFlagsTemporal = append(defaultCliFlagsDatabaseOnly, []cli.Flag{
	&cli.StringFlag{
		Name:    "temporal-namespace",
		Aliases: []string{"n"},
		Usage:   "temporal namespace",
		EnvVars: []string{string(EnvVarTemporalNamespace)},
		Value:   "default",
	},
	&cli.StringFlag{
		Name:    "temporal-address",
		Aliases: []string{"a"},
		Usage:   "temporal address",
		EnvVars: []string{string(EnvVarTemporalAddress)},
		Value:   "localhost:7233",
	},
	&cli.StringFlag{
		Name:    "temporal-task-queue",
		Aliases: []string{"q"},
		Usage:   "temporal task queue",
		EnvVars: []string{string(EnvVarTemporalTaskQueue)},
	},
}...)

var defaultCliFlagsNoAuth = append(defaultCliFlagsDatabaseOnly, []cli.Flag{
	&cli.IntFlag{
		Name:    "grpc-port",
		Usage:   "gRPC port",
		EnvVars: []string{string(EnvVarGRPCPort)},
		Value:   8080,
	},
	&cli.IntFlag{
		Name:    "gateway-port",
		Usage:   "gRPC gateway port",
		EnvVars: []string{string(EnvVarGatewayPort)},
		Value:   8081,
	},
}...)

var defaultCliFlags = append(defaultCliFlagsNoAuth, []cli.Flag{
	&cli.StringFlag{
		Name:    "oidc-issuer",
		Usage:   "OIDC issuer",
		EnvVars: []string{string(EnvVarFrontendOIDCIssuer)},
		Value:   "https://accounts.rockylinux.org/auth/realms/rocky",
	},
	&cli.StringFlag{
		Name:    "required-oidc-group",
		Usage:   "OIDC group that is required to access the frontend",
		EnvVars: []string{string(EnvVarFrontendRequiredOIDCGroup)},
	},
}...)

var defaultFrontendNoAuthCliFlags = []cli.Flag{
	&cli.IntFlag{
		Name:    "port",
		Usage:   "frontend port",
		EnvVars: []string{string(EnvVarFrontendPort)},
		Value:   9111,
	},
}

var defaultFrontendCliFlags = append(defaultFrontendNoAuthCliFlags, []cli.Flag{
	&cli.StringFlag{
		Name:    "oidc-issuer",
		Usage:   "OIDC issuer",
		EnvVars: []string{string(EnvVarFrontendOIDCIssuer)},
		Value:   "https://accounts.rockylinux.org/auth/realms/rocky",
	},
	&cli.StringFlag{
		Name:    "oidc-client-id",
		Usage:   "OIDC client ID",
		EnvVars: []string{string(EnvVarFrontendOIDCClientID)},
	},
	&cli.StringFlag{
		Name:    "oidc-client-secret",
		Usage:   "OIDC client secret",
		EnvVars: []string{string(EnvVarFrontendOIDCClientSecret)},
	},
	&cli.StringFlag{
		Name:    "oidc-userinfo-override",
		Usage:   "OIDC userinfo override",
		EnvVars: []string{string(EnvVarFrontendOIDCUserInfoOverride)},
	},
	&cli.StringFlag{
		Name:    "required-oidc-group",
		Usage:   "OIDC group that is required to access the frontend",
		EnvVars: []string{string(EnvVarFrontendRequiredOIDCGroup)},
	},
	&cli.StringFlag{
		Name:    "self",
		Usage:   "Endpoint pointing to the frontend",
		EnvVars: []string{string(EnvVarFrontendSelf)},
	},
}...)

var storageFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "storage-endpoint",
		Usage:   "storage endpoint",
		EnvVars: []string{string(EnvVarStorageEndpoint)},
		Value:   "",
	},
	&cli.StringFlag{
		Name:    "storage-connection-string",
		Usage:   "storage connection string",
		EnvVars: []string{string(EnvVarStorageConnectionString)},
	},
	&cli.StringFlag{
		Name:    "storage-region",
		Usage:   "storage region",
		EnvVars: []string{string(EnvVarStorageRegion)},
		// RESF default region
		Value: "us-east-2",
	},
	&cli.BoolFlag{
		Name:    "storage-secure",
		Usage:   "storage secure",
		EnvVars: []string{string(EnvVarStorageSecure)},
		Value:   true,
	},
	&cli.BoolFlag{
		Name:    "storage-path-style",
		Usage:   "storage path style",
		EnvVars: []string{string(EnvVarStoragePathStyle)},
		Value:   false,
	},
}

// WithDefaultCliFlags adds the default cli flags to the app.
func WithDefaultCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultCliFlags, flags...)
}

// WithDefaultCliFlagsNoAuth adds the default cli flags to the app.
func WithDefaultCliFlagsNoAuth(flags ...cli.Flag) []cli.Flag {
	return append(defaultCliFlagsNoAuth, flags...)
}

// WithDefaultCliFlagsTemporal adds the default cli flags to the app.
func WithDefaultCliFlagsTemporal(flags ...cli.Flag) []cli.Flag {
	return append(defaultCliFlagsTemporal, flags...)
}

// WithDefaultCliFlagsDatabaseOnly adds the default cli flags to the app.
func WithDefaultCliFlagsDatabaseOnly(flags ...cli.Flag) []cli.Flag {
	return append(defaultCliFlagsDatabaseOnly, flags...)
}

// WithDefaultFrontendNoAuthCliFlags adds the default frontend cli flags to the app.
func WithDefaultFrontendNoAuthCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultFrontendNoAuthCliFlags, flags...)
}

// WithDefaultFrontendCliFlags adds the default frontend cli flags to the app.
func WithDefaultFrontendCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultFrontendCliFlags, flags...)
}

// WithStorageFlags adds the storage flags to the app.
func WithStorageFlags(flags ...cli.Flag) []cli.Flag {
	return append(storageFlags, flags...)
}

// FlagsToGRPCServerOptions converts the cli flags to gRPC server options.
func FlagsToGRPCServerOptions(ctx *cli.Context) []GRPCServerOption {
	return []GRPCServerOption{
		WithGRPCPort(ctx.Int("grpc-port")),
		WithGatewayPort(ctx.Int("gateway-port")),
	}
}

// FlagsToFrontendInfo converts the cli flags to frontend info.
func FlagsToFrontendInfo(ctx *cli.Context) *FrontendInfo {
	return &FrontendInfo{
		Title:                ctx.App.Name,
		Port:                 ctx.Int("port"),
		Self:                 ctx.String("self"),
		OIDCIssuer:           ctx.String("oidc-issuer"),
		OIDCClientID:         ctx.String("oidc-client-id"),
		OIDCClientSecret:     ctx.String("oidc-client-secret"),
		OIDCGroup:            ctx.String("required-oidc-group"),
		OIDCUserInfoOverride: ctx.String("oidc-userinfo-override"),
	}
}

// FlagsToOidcInterceptorDetails converts the cli flags to oidc interceptor details.
func FlagsToOidcInterceptorDetails(ctx *cli.Context) (*OidcInterceptorDetails, error) {
	provider, err := oidc.NewProvider(ctx.Context, ctx.String("oidc-issuer"))
	if err != nil {
		return nil, err
	}

	return &OidcInterceptorDetails{
		Provider: &OidcProviderImpl{provider},
		Group:    ctx.String("required-oidc-group"),
	}, nil
}

// GetDBFromFlags gets the database from the cli flags.
func GetDBFromFlags(ctx *cli.Context) *DB {
	db, err := NewDB(ctx.String("database-url"))
	if err != nil {
		LogFatalf("failed to create database: %v", err)
	}

	return db
}

// GetTemporalClientFromFlags gets the temporal client from the cli flags.
func GetTemporalClientFromFlags(ctx *cli.Context, opts client.Options) (client.Client, error) {
	return NewTemporalClient(ctx.String("temporal-address"), ctx.String("temporal-namespace"), opts)
}

// ChangeDefaultForEnvVar changes the default value of a flag based on an environment variable.
func ChangeDefaultForEnvVar(envVar EnvVar, newDefault string) {
	// Check if the environment variable is set.
	if _, ok := os.LookupEnv(string(envVar)); ok {
		return
	}

	// Change the default value.
	if err := os.Setenv(string(envVar), newDefault); err != nil {
		LogFatalf("failed to set environment variable %s: %v", envVar, err)
	}
}

// RareUseChangeDefault changes the default value of an arbitrary environment variable.
func RareUseChangeDefault(envVar string, newDefault string) {
	// Check if the environment variable is set.
	if _, ok := os.LookupEnv(envVar); ok {
		return
	}

	// Change the default value.
	if err := os.Setenv(envVar, newDefault); err != nil {
		LogFatalf("failed to set environment variable %s: %v", envVar, err)
	}
}

// ChangeDefaultDatabaseURL changes the default value of the database url based on an environment variable.
func ChangeDefaultDatabaseURL(appName string) {
	ChangeDefaultForEnvVar(EnvVarDatabaseURL, "postgres://postgres:postgres@localhost:5432/"+appName+"?sslmode=disable")
}
