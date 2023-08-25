load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "enums",
    srcs = [
        "batch_operation.pb.go",
        "command_type.pb.go",
        "common.pb.go",
        "event_type.pb.go",
        "failed_cause.pb.go",
        "interaction_type.pb.go",
        "namespace.pb.go",
        "query.pb.go",
        "reset.pb.go",
        "schedule.pb.go",
        "task_queue.pb.go",
        "update.pb.go",
        "workflow.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/api/enums/v1",
    importpath = "go.temporal.io/api/enums/v1",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto"],
)
