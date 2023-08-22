load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "temporalite_lib",
    srcs = [
        "main.go",
        "ui.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/temporalite/cmd/temporalite",
    importpath = "github.com/temporalio/temporalite/cmd/temporalite",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/temporalio/temporalite",
        "//vendor/github.com/temporalio/temporalite/internal/liteconfig",
        "//vendor/github.com/temporalio/ui-server/v2/plugins/fs_config_provider",
        "//vendor/github.com/temporalio/ui-server/v2/server",
        "//vendor/github.com/temporalio/ui-server/v2/server/config",
        "//vendor/github.com/temporalio/ui-server/v2/server/server_options",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/persistence/sql/sqlplugin/sqlite",
        "//vendor/go.temporal.io/server/temporal",
        "//vendor/go.uber.org/zap",
        "//vendor/go.uber.org/zap/zapcore",
    ],
)

go_binary(
    name = "temporalite",
    embed = [":temporalite_lib"],
    visibility = ["//visibility:public"],
)
