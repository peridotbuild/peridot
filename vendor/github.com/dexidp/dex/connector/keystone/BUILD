load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "keystone",
    srcs = ["keystone.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/keystone",
    importpath = "github.com/dexidp/dex/connector/keystone",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/log",
    ],
)
