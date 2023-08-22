load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = [
        "auth.go",
        "config.go",
        "config_provider.go",
        "config_provider_with_refresh.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/temporalio/ui-server/v2/server/config",
    importpath = "github.com/temporalio/ui-server/v2/server/config",
    visibility = ["//visibility:public"],
)
