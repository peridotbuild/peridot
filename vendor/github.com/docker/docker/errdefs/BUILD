load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "errdefs",
    srcs = [
        "defs.go",
        "doc.go",
        "helpers.go",
        "http_helpers.go",
        "is.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/errdefs",
    importpath = "github.com/docker/docker/errdefs",
    visibility = ["//visibility:public"],
)
