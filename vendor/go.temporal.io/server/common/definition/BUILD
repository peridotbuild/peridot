load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "definition",
    srcs = [
        "resource_dedup.go",
        "workflow_key.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/definition",
    importpath = "go.temporal.io/server/common/definition",
    visibility = ["//visibility:public"],
)
