load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("//tools/build_rules/oapi_gen:defs.bzl", "oapi_gen_ts")

proto_library(
    name = "mshipadminpb_proto",
    srcs = [
        "bugtracker.proto",
        "mship_admin.proto",
        "worker.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@go_googleapis//google/api:annotations_proto",
        "@googleapis//google/api:annotations_proto",
    ],
)

oapi_gen_ts(
    name = "mshipadminpb_ts_proto",
    proto = ":mshipadminpb_proto",
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "mshipadminpb_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "//:go_gen_grpc_gateway",
    ],
    importpath = "go.resf.org/peridot/tools/mothership/admin/pb",
    proto = ":mshipadminpb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_genproto//googleapis/api/annotations",
    ],
)

go_library(
    name = "pb",
    embed = [":mshipadminpb_go_proto"],
    importpath = "go.resf.org/peridot/tools/mothership/admin/pb",
    visibility = ["//visibility:public"],
)
