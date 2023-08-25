load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "network",
    srcs = ["network.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/network",
    importpath = "github.com/docker/docker/api/types/network",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/api/types/filters"],
)
