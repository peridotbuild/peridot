load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "peridotpb_proto",
    srcs = [
        "batch.proto",
        "build.proto",
        "catalog.proto",
        "import.proto",
        "module.proto",
        "package.proto",
        "project.proto",
        "rpm.proto",
        "search.proto",
        "task.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//third_party/peridot/proto/v1/yumrepofs:yumrepofspb_proto",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_protobuf//:wrappers_proto",
        "@go_googleapis//google/api:httpbody_proto",
        "@go_googleapis//google/rpc:errdetails_proto",
        "@googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "peridotpb_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "//:go_gen_grpc_gateway",
    ],
    importpath = "peridot.resf.org/peridot/pb",
    proto = ":peridotpb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//third_party/peridot/proto/v1/yumrepofs:pb",
        "@go_googleapis//google/api:httpbody_go_proto",
        "@go_googleapis//google/rpc:errdetails_go_proto",
        "@org_golang_google_genproto//googleapis/api/annotations",
    ],
)

go_library(
    name = "pb",
    embed = [":peridotpb_go_proto"],
    importpath = "peridot.resf.org/peridot/pb",
    visibility = ["//visibility:public"],
)
