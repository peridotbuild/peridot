load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "startworkflow",
    srcs = ["api.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api/startworkflow",
    importpath = "go.temporal.io/server/service/history/api/startworkflow",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/uuid",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/metrics",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/service/history/api",
        "//vendor/go.temporal.io/server/service/history/consts",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/workflow",
    ],
)
