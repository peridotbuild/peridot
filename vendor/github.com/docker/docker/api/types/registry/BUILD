load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "registry",
    srcs = [
        "authconfig.go",
        "authenticate.go",
        "registry.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/api/types/registry",
    importpath = "github.com/docker/docker/api/types/registry",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opencontainers/image-spec/specs-go/v1:specs-go",
        "//vendor/github.com/pkg/errors",
    ],
)
