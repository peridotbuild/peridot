load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "temporalite",
    srcs = [
        "options.go",
        "server.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/temporalite",
    importpath = "github.com/temporalio/temporalite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/temporalio/temporalite/internal/liteconfig",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/server/common/authorization",
        "//vendor/go.temporal.io/server/common/config",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/schema/sqlite",
        "//vendor/go.temporal.io/server/temporal",
    ],
)
