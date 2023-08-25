load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "recordchildworkflowcompleted",
    srcs = ["api.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/api/recordchildworkflowcompleted",
    importpath = "go.temporal.io/server/service/history/api/recordchildworkflowcompleted",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/historyservice/v1:historyservice",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/service/history/api",
        "//vendor/go.temporal.io/server/service/history/consts",
        "//vendor/go.temporal.io/server/service/history/shard",
        "//vendor/go.temporal.io/server/service/history/workflow",
    ],
)
