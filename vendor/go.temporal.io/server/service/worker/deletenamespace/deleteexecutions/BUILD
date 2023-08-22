load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "deleteexecutions",
    srcs = [
        "activities.go",
        "config.go",
        "workflow.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/deletenamespace/deleteexecutions",
    importpath = "go.temporal.io/server/service/worker/deletenamespace/deleteexecutions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/sdk/activity",
        "//vendor/go.temporal.io/sdk/temporal",
        "//vendor/go.temporal.io/sdk/workflow",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
        "//vendor/go.temporal.io/server/common/quotas",
        "//vendor/go.temporal.io/server/service/worker/deletenamespace/errors",
    ],
)
