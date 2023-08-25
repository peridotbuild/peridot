load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "refreshworkflow",
    srcs = ["api.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api/refreshworkflow",
    importpath = "go.temporal.io/server/service/history/api/refreshworkflow",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence",
        "//vendor/go.temporal.io/server/service/history/api",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/workflow",
    ],
)
