load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "namespace",
    srcs = ["message.pb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/api/namespace/v1",
    importpath = "go.temporal.io/server/api/namespace/v1",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gogo/protobuf/proto"],
)
