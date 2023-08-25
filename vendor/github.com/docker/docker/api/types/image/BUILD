load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "image",
    srcs = [
        "image_history.go",
        "opts.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/image",
    importpath = "github.com/docker/docker/api/types/image",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opencontainers/image-spec/specs-go/v1:specs-go"],
)
