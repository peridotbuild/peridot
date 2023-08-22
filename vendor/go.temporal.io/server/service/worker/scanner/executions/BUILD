load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "executions",
    srcs = [
        "history_event_id_validator.go",
        "interfaces.go",
        "mutable_state_id_validator.go",
        "scavenger.go",
        "task.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/scanner/executions",
    importpath = "go.temporal.io/server/service/worker/scanner/executions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/collection",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/versionhistory",
        "//vendor/go.temporal.io/server/common/quotas",
        "//vendor/go.temporal.io/server/service/worker/scanner/executor",
    ],
)
