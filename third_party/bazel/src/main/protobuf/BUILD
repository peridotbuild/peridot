load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "blaze_query_proto",
    srcs = ["build.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "blaze_query_go_proto",
    importpath = "go.resf.org/peridot/third_party/bazel/src/main/protobuf",
    proto = ":blaze_query_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "protobuf",
    embed = [":blaze_query_go_proto"],
    importpath = "go.resf.org/peridot/third_party/bazel/src/main/protobuf",
    visibility = ["//visibility:public"],
)
