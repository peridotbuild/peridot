load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "mship_ui_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/mothership/cmd/mship_ui",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//tools/mothership/ui",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "mship_ui",
    embed = [":mship_ui_lib"],
    visibility = ["//visibility:public"],
)
