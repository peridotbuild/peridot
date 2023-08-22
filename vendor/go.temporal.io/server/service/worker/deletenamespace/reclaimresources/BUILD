load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "reclaimresources",
    srcs = [
        "activities.go",
        "workflow.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/deletenamespace/reclaimresources",
    importpath = "go.temporal.io/server/service/worker/deletenamespace/reclaimresources",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/sdk/activity",
        "//vendor/go.temporal.io/sdk/temporal",
        "//vendor/go.temporal.io/sdk/workflow",
        "//vendor/go.temporal.io/server/common/headers",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/persistence/visibility/manager",
        "//vendor/go.temporal.io/server/service/worker/deletenamespace/deleteexecutions",
        "//vendor/go.temporal.io/server/service/worker/deletenamespace/errors",
    ],
)
