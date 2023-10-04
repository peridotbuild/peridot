load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "kernelmanager_dev_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/cmd/kernelmanager_dev",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//base/go/kv/dynamodb",
        "//tools/kernelmanager/rpc",
        "//tools/kernelmanager/ui",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/go.temporal.io/sdk/client",
    ],
)

go_binary(
    name = "kernelmanager_dev",
    embed = [":kernelmanager_dev_lib"],
    visibility = ["//visibility:public"],
)
