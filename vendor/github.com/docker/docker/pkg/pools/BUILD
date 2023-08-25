load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "pools",
    srcs = ["pools.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/pools",
    importpath = "github.com/docker/docker/pkg/pools",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/ioutils"],
)
