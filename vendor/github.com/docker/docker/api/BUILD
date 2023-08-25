load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "common.go",
        "common_unix.go",
        "common_windows.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api",
    importpath = "github.com/docker/docker/api",
    visibility = ["//visibility:public"],
)
