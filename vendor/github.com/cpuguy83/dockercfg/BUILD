load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dockercfg",
    srcs = [
        "auth.go",
        "config.go",
        "load.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/cpuguy83/dockercfg",
    importpath = "github.com/cpuguy83/dockercfg",
    visibility = ["//visibility:public"],
)
