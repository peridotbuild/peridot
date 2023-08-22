load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "configs",
    srcs = [
        "config.go",
        "quotas.go",
        "task.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/service/history/configs",
    importpath = "go.temporal.io/server/service/history/configs",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
        "//vendor/go.temporal.io/server/common/namespace",
        "//vendor/go.temporal.io/server/common/persistence/visibility",
        "//vendor/go.temporal.io/server/common/quotas",
        "//vendor/go.temporal.io/server/common/tasks",
    ],
)
