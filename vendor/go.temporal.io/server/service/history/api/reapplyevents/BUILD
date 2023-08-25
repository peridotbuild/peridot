load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "reapplyevents",
    srcs = ["api.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api/reapplyevents",
    importpath = "go.temporal.io/server/service/history/api/reapplyevents",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/uuid",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence/versionhistory",
        "//vendor/go.temporal.io/server/service/history/api",
        "//vendor/go.temporal.io/server/service/history/ndc",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/workflow",
    ],
)
