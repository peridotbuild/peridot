package base

import (
	"github.com/urfave/cli/v2"
	"os"
)

type EnvVar string

const (
	EnvVarGRPCPort                  EnvVar = "GRPC_PORT"
	EnvVarGatewayPort               EnvVar = "GATEWAY_PORT"
	EnvVarDatabaseURL               EnvVar = "DATABASE_URL"
	EnvVarFrontendPort              EnvVar = "FRONTEND_PORT"
	EnvVarFrontendOIDCIssuer        EnvVar = "FRONTEND_OIDC_ISSUER"
	EnvVarFrontendOIDCClientID      EnvVar = "FRONTEND_OIDC_CLIENT_ID"
	EnvVarFrontendOIDCClientSecret  EnvVar = "FRONTEND_OIDC_CLIENT_SECRET"
	EnvVarFrontendRequiredOIDCGroup EnvVar = "FRONTEND_REQUIRED_OIDC_GROUP"
	EnvVarFrontendAdminOIDCGroup    EnvVar = "FRONTEND_ADMIN_OIDC_GROUP"
)

var defaultCliFlags = []cli.Flag{
	&cli.IntFlag{
		Name:    "grpc-port",
		Aliases: []string{"p"},
		Usage:   "gRPC port",
		EnvVars: []string{string(EnvVarGRPCPort)},
		Value:   8080,
	},
	&cli.IntFlag{
		Name:    "gateway-port",
		Aliases: []string{"g"},
		Usage:   "gRPC gateway port",
		EnvVars: []string{string(EnvVarGatewayPort)},
		Value:   8081,
	},
	&cli.StringFlag{
		Name:    "database-url",
		Aliases: []string{"d"},
		Usage:   "database url",
		EnvVars: []string{string(EnvVarDatabaseURL)},
		Value:   "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	},
}

var defaultFrontendNoAuthCliFlags = []cli.Flag{
	&cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
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
		Name:    "required-oidc-group",
		Usage:   "OIDC group that is required to access the frontend",
		EnvVars: []string{string(EnvVarFrontendRequiredOIDCGroup)},
	},
}...)

var defaultFrontendAdminCliFlags = append(defaultFrontendCliFlags, []cli.Flag{
	&cli.StringFlag{
		Name:    "admin-oidc-group",
		Usage:   "OIDC group that is allowed to access the admin page",
		EnvVars: []string{string(EnvVarFrontendAdminOIDCGroup)},
	},
}...)

// WithDefaultCliFlags adds the default cli flags to the app.
func WithDefaultCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultCliFlags, flags...)
}

// WithDefaultFrontendNoAuthCliFlags adds the default frontend cli flags to the app.
func WithDefaultFrontendNoAuthCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultFrontendNoAuthCliFlags, flags...)
}

// WithDefaultFrontendCliFlags adds the default frontend cli flags to the app.
func WithDefaultFrontendCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultFrontendCliFlags, flags...)
}

// WithDefaultFrontendAdminCliFlags adds the default frontend cli flags to the app.
func WithDefaultFrontendAdminCliFlags(flags ...cli.Flag) []cli.Flag {
	return append(defaultFrontendAdminCliFlags, flags...)
}

// FlagsToGRPCServerOptions converts the cli flags to gRPC server options.
func FlagsToGRPCServerOptions(ctx *cli.Context) []GRPCServerOption {
	return []GRPCServerOption{
		WithGRPCPort(ctx.Int("grpc-port")),
		WithGatewayPort(ctx.Int("gateway-port")),
	}
}

// GetDBFromFlags gets the database from the cli flags.
func GetDBFromFlags(ctx *cli.Context) *DB {
	db, err := NewDB(ctx.String("database-url"))
	if err != nil {
		LogFatalf("failed to create database: %v", err)
	}

	return db
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

// ChangeDefaultDatabaseURL changes the default value of the database url based on an environment variable.
func ChangeDefaultDatabaseURL(appName string) {
	ChangeDefaultForEnvVar(EnvVarDatabaseURL, "postgres://postgres:postgres@localhost:5432/"+appName+"?sslmode=disable")
}
