load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "kernelmanager_server_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/cmd/kernelmanager_server",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//base/go/kv/dynamodb",
        "//tools/kernelmanager/rpc",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/go.temporal.io/sdk/client",
    ],
)

go_binary(
    name = "kernelmanager_server",
    embed = [":kernelmanager_server_lib"],
    visibility = ["//visibility:public"],
)
