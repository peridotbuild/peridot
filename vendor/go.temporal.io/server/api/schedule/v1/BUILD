load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schedule",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/schedule/v1",
    importpath = "go.temporal.io/server/api/schedule/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/api/enums/v1:enums",
        "//vendor/go.temporal.io/api/failure/v1:failure",
        "//vendor/go.temporal.io/api/schedule/v1:schedule",
        "//vendor/go.temporal.io/api/workflowservice/v1:workflowservice",
    ],
)
