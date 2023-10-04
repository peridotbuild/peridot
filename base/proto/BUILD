load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "basepb_proto",
    srcs = ["kv_page_token.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "basepb_go_proto",
    importpath = "go.resf.org/peridot/base/go/pb",
    proto = ":basepb_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "pb",
    embed = [":basepb_go_proto"],
    importpath = "go.resf.org/peridot/base/go/pb",
    visibility = ["//visibility:public"],
)
