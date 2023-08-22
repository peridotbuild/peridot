load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "workflow",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/workflow/v1",
    importpath = "go.temporal.io/server/api/workflow/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/api/common/v1:common",
        "//vendor/go.temporal.io/server/api/clock/v1:clock",
    ],
)
