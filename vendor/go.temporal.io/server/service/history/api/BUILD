load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "api",
    srcs = [
        "activity_util.go",
        "consistency_checker.go",
        "create_workflow_util.go",
        "get_workflow_util.go",
        "namespace.go",
        "replication_util.go",
        "retry_util.go",
        "signal_workflow_util.go",
        "update_workflow_util.go",
        "workflow_context.go",
        "workflow_id_reuse_policy.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api",
    importpath = "go.temporal.io/server/service/history/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pborman/uuid",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/api/taskqueue/v1:taskqueue",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/token/v1:token",
        "//vendor/go.temporal.io/server/api/workflow/v1:workflow",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/cluster",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/payloads",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/versionhistory",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/common/rpc/interceptor",
        "//vendor/go.temporal.io/server/common/serviceerror",
        "//vendor/go.temporal.io/server/service/history/consts",
        "//vendor/go.temporal.io/server/service/history/events",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/vclock",
        "//vendor/go.temporal.io/server/service/history/workflow",
    ],
)
