load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "kernelmanager_ui_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/cmd/kernelmanager_ui",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//tools/kernelmanager/ui",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "kernelmanager_ui",
    embed = [":kernelmanager_ui_lib"],
    visibility = ["//visibility:public"],
)
