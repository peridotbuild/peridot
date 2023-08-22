load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "tasks",
    srcs = [
        "fifo_scheduler.go",
        "interleaved_weighted_round_robin.go",
        "monitor.go",
        "priority.go",
        "scheduler.go",
        "scheduler_mock.go",
        "state.go",
        "task.go",
        "task_mock.go",
        "weighted_channel.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/tasks",
    importpath = "go.temporal.io/server/common/tasks",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/go.temporal.io/server/common",
        "//vendor/go.temporal.io/server/common/backoff",
        "//vendor/go.temporal.io/server/common/dynamicconfig",
        "//vendor/go.temporal.io/server/common/log",
        "//vendor/go.temporal.io/server/common/log/tag",
    ],
)
