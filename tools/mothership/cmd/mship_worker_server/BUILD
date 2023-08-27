load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "mship_worker_server_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_worker_server",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//base/go/storage/detector",
        "//tools/mothership/worker_server",
        "//vendor/github.com/urfave/cli/v2:cli",
        "//vendor/go.temporal.io/sdk/client",
        "//vendor/go.temporal.io/sdk/worker",
    ],
)

go_binary(
    name = "mship_worker_server",
    embed = [":mship_worker_server_lib"],
    visibility = ["//visibility:public"],
)
