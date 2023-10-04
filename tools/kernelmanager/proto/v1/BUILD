load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("//tools/build_rules/oapi_gen:defs.bzl", "oapi_gen_ts")

proto_library(
    name = "kernelmanagerpb_proto",
    srcs = [
        "changelog.proto",
        "kernel.proto",
        "kernel_state.proto",
        "kernelmanager.proto",
        "update.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:field_mask_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@go_googleapis//google/api:annotations_proto",
        "@go_googleapis//google/longrunning:longrunning_proto",
        "@googleapis//google/api:annotations_proto",
    ],
)

oapi_gen_ts(
    name = "kernelmanagerpb_ts_proto",
    proto = ":kernelmanagerpb_proto",
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "kernelmanagerpb_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "//:go_gen_grpc_gateway",
    ],
    importpath = "go.resf.org/peridot/tools/kernelmanager/pb",
    proto = ":kernelmanagerpb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//third_party/googleapis/google/longrunning:longrunning_go_proto",
        "@go_googleapis//google/api:annotations_go_proto",
        "@org_golang_google_genproto//googleapis/api/annotations",
    ],
)

go_library(
    name = "pb",
    embed = [":kernelmanagerpb_go_proto"],
    importpath = "go.resf.org/peridot/tools/kernelmanager/pb",
    visibility = ["//visibility:public"],
)
