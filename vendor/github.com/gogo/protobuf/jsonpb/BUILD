load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "jsonpb",
    srcs = ["jsonpb.go"],
    importmap = "go.resf.org/peridot/vendor/github.com/gogo/protobuf/jsonpb",
    importpath = "github.com/gogo/protobuf/jsonpb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
    ],
)
