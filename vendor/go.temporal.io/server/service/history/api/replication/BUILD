load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "replication",
    srcs = [
        "generate_task.go",
        "get_dlq_tasks.go",
        "get_tasks.go",
        "status.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api/replication",
    importpath = "go.temporal.io/server/service/history/api/replication",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/service/history/api",
        "//vendor/go.temporal.io/server/service/history/replication",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/tasks",
    ],
)
