load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "mship_dev_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_dev",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//tools/mothership/admin/rpc",
        "//tools/mothership/admin/ui",
        "//tools/mothership/rpc",
        "//tools/mothership/ui",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/v2/runtime",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "mship_dev",
    embed = [":mship_dev_lib"],
    visibility = ["//visibility:public"],
)
