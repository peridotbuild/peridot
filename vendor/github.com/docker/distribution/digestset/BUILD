load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "digestset",
    srcs = ["set.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/distribution/digestset",
    importpath = "github.com/docker/distribution/digestset",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opencontainers/go-digest"],
)
