load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "batch",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/api/batch/v1",
    importpath = "go.temporal.io/api/batch/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
    ],
)
