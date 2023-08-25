load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "nat",
    srcs = [
        "nat.go",
        "parse.go",
        "sort.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/go-connections/nat",
    importpath = "github.com/docker/go-connections/nat",
    visibility = ["//visibility:public"],
)
