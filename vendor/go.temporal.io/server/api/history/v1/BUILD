load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "history",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/history/v1",
    importpath = "go.temporal.io/server/api/history/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/go.temporal.io/api/history/v1:history",
    ],
)
