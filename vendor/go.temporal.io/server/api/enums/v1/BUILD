load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "enums",
    srcs = [
        "cluster.pb.go",
        "common.pb.go",
        "predicate.pb.go",
        "replication.pb.go",
        "task.pb.go",
        "workflow.pb.go",
    ],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/enums/v1",
    importpath = "go.temporal.io/server/api/enums/v1",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto"],
)
