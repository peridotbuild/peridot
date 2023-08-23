load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "memory",
    srcs = ["memory.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/dexidp/dex/storage/memory",
    importpath = "github.com/dexidp/dex/storage/memory",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dexidp/dex/pkg/log",
        "//vendor/github.com/dexidp/dex/storage",
    ],
)
