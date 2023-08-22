load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "mship_api_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_api",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//tools/mothership/rpc",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "mship_api",
    embed = [":mship_api_lib"],
    visibility = ["//visibility:public"],
)
