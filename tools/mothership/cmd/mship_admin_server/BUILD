load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "mship_admin_server_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_admin_server",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//tools/mothership/admin/rpc",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "mship_admin_server",
    embed = [":mship_admin_server_lib"],
    visibility = ["//visibility:public"],
)
