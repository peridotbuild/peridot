load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "history",
    srcs = ["scavenger.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/worker/scanner/history",
    importpath = "go.temporal.io/server/service/worker/scanner/history",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/sdk/activity",
        "//vendor/go.temporal.io/server/api/adminservice/v1:adminservice",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/collection",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/common/quotas",
    ],
)
