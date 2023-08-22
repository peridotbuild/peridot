load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "replicator",
    srcs = [
        "namespace_replication_message_processor.go",
        "replicator.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/replicator",
    importpath = "go.temporal.io/server/service/worker/replicator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "//vendor/go.temporal.io/server/client",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/membership",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/rpc",
    ],
)
