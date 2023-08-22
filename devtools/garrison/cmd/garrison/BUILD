load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "garrison_lib",
    srcs = ["main.go"],
    importpath = "go.resf.org/peridot/devtools/garrison/cmd/garrison",
    visibility = ["//visibility:private"],
    deps = [
        "//base/go",
        "//devtools/garrison/internal/pkg/kubernetes",
        "//devtools/garrison/internal/pkg/starlark",
        "//vendor/github.com/urfave/cli/v2:cli",
    ],
)

go_binary(
    name = "garrison",
    embed = [":garrison_lib"],
    visibility = ["//visibility:public"],
)
