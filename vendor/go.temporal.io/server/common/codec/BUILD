load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "codec",
    srcs = ["jsonpb.go"],
    importmap = "go.resf.org/peridot/vendor/go.temporal.io/server/common/codec",
    importpath = "go.temporal.io/server/common/codec",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/jsonpb",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/go.temporal.io/api/history/v1:history",
    ],
)
