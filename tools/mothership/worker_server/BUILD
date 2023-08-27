load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "worker_server",
    srcs = [
        "process_rpm.go",
        "worker.go",
        "workflows.go",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/worker_server",
    visibility = ["//visibility:public"],
    deps = [
        "//base/go",
        "//base/go/storage",
        "//tools/mothership/proto/v1:pb",
        "//vendor/github.com/pkg/errors",
        "//vendor/go.temporal.io/sdk/temporal",
        "//vendor/go.temporal.io/sdk/workflow",
    ],
)
