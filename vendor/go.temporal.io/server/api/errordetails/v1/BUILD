load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "errordetails",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/errordetails/v1",
    importpath = "go.temporal.io/server/api/errordetails/v1",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto"],
)
