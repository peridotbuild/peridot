load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "serialization",
    srcs = [
        "blob.go",
        "serializer.go",
        "task_serializer.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/persistence/serialization",
    importpath = "go.temporal.io/server/common/persistence/serialization",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/history/v1:history",
        "//vendor/go.temporal.io/api/serviceerror",
        "//vendor/go.temporal.io/server/api/enums/v1:enums",
        "//vendor/go.temporal.io/server/api/persistence/v1:persistence",
        "//vendor/go.temporal.io/server/api/replication/v1:replication",
        "//vendor/go.temporal.io/server/common/codec",
        "//vendor/go.temporal.io/server/common/definition",
        "//vendor/go.temporal.io/server/common/primitives/timestamp",
        "//vendor/go.temporal.io/server/service/history/tasks",
    ],
)
