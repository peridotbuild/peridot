load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "userns",
    srcs = [
        "userns_linux.go",
        "userns_unsupported.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/containerd/containerd/pkg/userns",
    importpath = "github.com/containerd/containerd/pkg/userns",
    visibility = ["//visibility:public"],
)
