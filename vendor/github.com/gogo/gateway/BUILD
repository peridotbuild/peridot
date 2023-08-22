load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gateway",
    srcs = [
        "jsonpb.go",
        "proto.go",
    ],
    importmap = "go.resf.org/peridot/vendor/github.com/gogo/gateway",
    importpath = "github.com/gogo/gateway",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/jsonpb",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime",
    ],
)
