load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "taskqueue",
    srcs = [
        "db.go",
        "handler.go",
        "scavenger.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/scanner/taskqueue",
    importpath = "go.temporal.io/server/service/worker/scanner/taskqueue",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/service/worker/scanner/executor",
    ],
)
