load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cache",
    srcs = [
        "cache.go",
        "lru.go",
        "simple.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/cache",
    importpath = "go.temporal.io/server/common/cache",
    visibility = ["//visibility:public"],
)
