load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "replication",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/api/replication/v1",
    importpath = "go.temporal.io/api/replication/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/api/enums/v1:enums",
    ],
)
