load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "exec",
    srcs = ["processor.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/testcontainers/testcontainers-go/exec",
    importpath = "github.com/testcontainers/testcontainers-go/exec",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/stdcopy"],
)
