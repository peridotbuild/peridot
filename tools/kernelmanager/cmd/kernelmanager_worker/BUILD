load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "kernelmanager_worker_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/cmd/kernelmanager_worker",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//base/go/forge/gitlab",
        "//base/go/kv/dynamodb",
        "//base/go/storage/detector",
        "//tools/kernelmanager/worker",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/sdk/worker",
    ],
)

go_binary(
    name = "kernelmanager_worker",
    embed = [":kernelmanager_worker_lib"],
    visibility = ["//visibility:public"],
)
