load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "jsonmessage",
    srcs = ["jsonmessage.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/docker/docker/pkg/jsonmessage",
    importpath = "github.com/docker/docker/pkg/jsonmessage",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/go-units",
        "//vendor/github.com/moby/term",
        "//vendor/github.com/morikuni/aec",
    ],
)
