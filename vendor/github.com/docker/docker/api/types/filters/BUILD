load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "filters",
    srcs = [
        "errors.go",
        "parse.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/filters",
    importpath = "github.com/docker/docker/api/types/filters",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/api/types/versions"],
)
