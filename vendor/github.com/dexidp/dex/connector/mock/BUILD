load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "mock",
    srcs = ["connectortest.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/mock",
    importpath = "github.com/dexidp/dex/connector/mock",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/log",
    ],
)
