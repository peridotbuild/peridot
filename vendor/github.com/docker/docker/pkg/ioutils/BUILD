load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "ioutils",
    srcs = [
        "buffer.go",
        "bytespipe.go",
        "fswriters.go",
        "readers.go",
        "tempdir_deprecated.go",
        "writeflusher.go",
        "writers.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/ioutils",
    importpath = "github.com/docker/docker/pkg/ioutils",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/longpath"],
)
