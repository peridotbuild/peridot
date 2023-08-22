load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "executor",
    srcs = [
        "executor.go",
        "runq.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/scanner/executor",
    importpath = "go.temporal.io/server/service/worker/scanner/executor",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/metrics",
    ],
)
