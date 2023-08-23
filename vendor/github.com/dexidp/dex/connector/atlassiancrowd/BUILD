load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "atlassiancrowd",
    srcs = ["atlassiancrowd.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/connector/atlassiancrowd",
    importpath = "github.com/dexidp/dex/connector/atlassiancrowd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/connector",
        "//vendor/github.com/dexidp/dex/pkg/groups",
        "//vendor/github.com/dexidp/dex/pkg/log",
    ],
)
