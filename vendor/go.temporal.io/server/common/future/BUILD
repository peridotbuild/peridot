load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "future",
    srcs = [
        "future.go",
        "future_impl.go",
        "ready_future_impl.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/future",
    importpath = "go.temporal.io/server/common/future",
    visibility = ["//visibility:public"],
)
